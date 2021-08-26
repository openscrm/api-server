package services

import (
	"github.com/gogf/gf/os/gfile"
	"github.com/pkg/errors"
	"net/http"
	url2 "net/url"
	"openscrm/app/entities"
	"openscrm/common/ecode"
	"openscrm/common/storage"
	"openscrm/common/we_work"
	"openscrm/conf"
	"openscrm/pkg/easywework"
	"os"
	"strconv"
	"strings"
)

type TempMaterial struct {
}

func NewTempMaterial() *TempMaterial {
	return &TempMaterial{}
}

func (o TempMaterial) UploadMedia(req entities.UploadMaterialReq, extCorpID string) (*workwx.MediaUploadResult, error) {
	client, err := we_work.Clients.Get(extCorpID)
	if err != nil {
		return nil, err
	}

	filename, err := getStorageFile(req.FileURL)
	if err != nil {
		return nil, err
	}

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	media, err := workwx.NewMediaFromFile(f)
	if err != nil {
		return nil, err
	}

	tempImageMedia, err := client.Customer.UploadTempImageMedia(media)
	if err != nil {
		return nil, err
	}

	return tempImageMedia, err
}

func getStorageFile(fileURL string) (string, error) {
	localStorage, err := storage.NewLocalStorage(conf.Settings.Storage)
	if err != nil {
		return "", err
	}

	URL, err := url2.Parse(fileURL)
	if err != nil {
		return "", err
	}
	expireAt, err := strconv.Atoi(URL.Query().Get("expire_at"))
	if err != nil {
		return "", err
	}
	signature := URL.Query().Get("signature")

	err = localStorage.CheckSignedURL(URL.Path, http.MethodGet, int64(expireAt), signature)
	if err != nil {
		return "", err
	}

	objectKey := strings.TrimPrefix(URL.Path, conf.Settings.Storage.ServerRootPath)
	filePath, err := localStorage.AbsPath(objectKey)
	if err != nil {
		err = errors.Wrap(err, "AbsPath failed")
		return "", err
	}

	if !localStorage.IsFilePathContains(filePath, localStorage.Config.LocalRootPath) {
		err = errors.New("invalid filePath")
		return "", err
	}

	if !gfile.Exists(filePath) {
		err = errors.WithStack(ecode.BadRequest)
		return "", err
	}
	return filePath, nil
}
