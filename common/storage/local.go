package storage

import (
	"fmt"
	"github.com/gogf/gf/crypto/gsha1"
	"github.com/gogf/gf/os/gfile"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"net/url"
	"openscrm/common/ecode"
	"openscrm/common/log"
	"openscrm/conf"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

type LocalStorage struct {
	Config conf.StorageConfig
}

func NewLocalStorage(conf conf.StorageConfig) (localStorage LocalStorage, err error) {
	localStorage.Config = conf
	localStorage.Config.LocalRootPath = filepath.Clean(localStorage.Config.LocalRootPath)
	localStorage.Config.ServerRootPath = path.Clean(localStorage.Config.ServerRootPath)
	if conf.LocalRootPath == "" || conf.ServerRootPath == "" {
		err = errors.New("invalid LocalRootPath or ServerRootPath")
		return
	}
	return
}

func (o LocalStorage) CheckSignedURL(signedURL string, method string, expireAt int64, signature string) (err error) {
	fileURL, err := url.Parse(signedURL)
	if err != nil {
		err = errors.Wrap(err, "invalid signedURL failed")
		return
	}

	if time.Now().Unix() > expireAt {
		err = errors.WithStack(ecode.ExpiredSignError)
		return
	}

	fileURL.RawQuery = fmt.Sprintf("expire_at=%d",
		expireAt,
	)

	signData := fmt.Sprintf("path=%s;method=%s;query=%s;sign_key=%s", fileURL.Path, method, fileURL.RawQuery, conf.Settings.App.Key)
	log.Sugar.Debug("signData", signData)
	if signature != gsha1.Encrypt(signData) {
		err = errors.WithStack(ecode.InvalidSignError)
		return
	}

	return
}

//func (o LocalStorage) SignURL(objectKey string, method constants.HTTPMethod, expiredInSec int64) (signedURL string, err error) {
//	if !IsValidObjectKey(objectKey) {
//		err = errors.WithStack(ecode.InvalidPathError)
//		return
//	}
//
//	//fileURL, err := url.Parse(conf.Settings.App.URL)
//	//if err != nil {
//	//	err = errors.Wrap(err, "invalid MainApp URL")
//	//	return
//	//}
//
//	fileURL.Path = path.Join(o.Config.ServerRootPath, objectKey)
//
//	fileURL.RawQuery = fmt.Sprintf("expire_at=%d",
//		time.Now().Unix()+expiredInSec,
//	)
//
//	signData := fmt.Sprintf("path=%s;method=%s;query=%s;sign_key=%s", fileURL.Path, method, fileURL.RawQuery, conf.Settings.App.Key)
//	log.Sugar.Debug("signData", signData)
//	signature := gsha1.Encrypt(signData)
//	fileURL.RawQuery += fmt.Sprintf("&signature=%s", signature)
//	signedURL = fileURL.String()
//
//	return
//}

func (o LocalStorage) Get(objectKey string) (content io.ReadCloser, err error) {
	filePath, err := o.AbsPath(objectKey)
	if err != nil {
		err = errors.Wrap(err, "AbsPath failed")
		return
	}

	if !o.IsFilePathContains(filePath, o.Config.LocalRootPath) {
		err = errors.New("invalid filePath")
		return
	}

	if !gfile.Exists(filePath) {
		err = errors.WithStack(ecode.BadRequest)
		return
	}

	f, err := os.Open(filePath)
	if err != nil {
		err = errors.Wrap(err, "os.Open failed")
		return
	}

	return f, nil
}

func (o LocalStorage) Put(objectKey string, reader io.Reader) (err error) {
	filePath, err := o.AbsPath(objectKey)
	if err != nil {
		err = errors.Wrap(err, "AbsPath failed")
		return
	}

	if !o.IsFilePathContains(filePath, o.Config.LocalRootPath) {
		err = errors.New("invalid filePath")
		return
	}

	data, err := ioutil.ReadAll(reader)
	if err != nil {
		err = errors.Wrap(err, "ioutil.ReadAll failed")
		return
	}

	if len(data) == 0 {
		err = errors.WithStack(ecode.BadRequest)
		return
	}

	err = os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
	if err != nil {
		err = errors.Wrap(err, "os.MkdirAll failed")
		return
	}

	err = ioutil.WriteFile(filePath, data, os.ModePerm)
	if err != nil {
		err = errors.Wrap(err, "ioutil.WriteFile failed")
		return
	}

	return
}

func (o LocalStorage) IsExist(objectKey string) (ok bool, err error) {
	filePath, err := o.AbsPath(objectKey)
	if err != nil {
		err = errors.Wrap(err, "AbsPath failed")
		return
	}

	ok = gfile.Exists(filePath)

	return
}

func (o LocalStorage) PutFromFile(objectKey string, filePath string) (err error) {
	if !gfile.Exists(filePath) {
		err = errors.New("target file is not exists")
		return
	}

	f, err := os.Open(filePath)
	if err != nil {
		err = errors.Wrap(err, "os.Open failed")
		return
	}
	defer f.Close()

	err = o.Put(objectKey, f)
	if err != nil {
		err = errors.Wrap(err, "Put failed")
		return
	}

	return
}

func (o LocalStorage) Delete(objectKeys ...string) (deletedObjects []string, err error) {
	deletedObjects = make([]string, 0)
	for _, key := range objectKeys {
		filePath, err := o.AbsPath(key)
		if err != nil {
			log.TracedError("AbsPath failed", err)
			continue
		}

		err = os.Remove(filePath)
		if err != nil {
			log.TracedError("os.Remove failed", err)
			continue
		}

		deletedObjects = append(deletedObjects, key)
	}

	return
}

func (o LocalStorage) AbsPath(objectKey string) (absPath string, err error) {
	if !IsValidObjectKey(objectKey) {
		err = errors.WithStack(ecode.InvalidPathError)
		return
	}

	absPath = filepath.Join(o.Config.LocalRootPath, objectKey)
	if !o.IsFilePathContains(absPath, o.Config.LocalRootPath) {
		err = errors.New("invalid absPath")
		return
	}

	return
}

func (o LocalStorage) IsFilePathContains(targetPath string, parentPath string) bool {
	targetPath = filepath.Clean(targetPath)
	parentPath = filepath.Clean(parentPath)
	return strings.HasPrefix(targetPath, parentPath)
}
