package storage

import (
	"github.com/pkg/errors"
	"github.com/tencentyun/cos-go-sdk-v5"
	"golang.org/x/net/context"
	"io"
	"mime"
	"net/http"
	"net/url"
	"openscrm/app/constants"
	setting "openscrm/conf"
	"path/filepath"
	"strings"
	"time"
)

type COSStorage struct {
	Client *cos.Client
	Config setting.StorageConfig
}

func NewCOS(conf setting.StorageConfig) (storage COSStorage, err error) {
	u, err := url.Parse(conf.BucketURL)
	if err != nil {
		err = errors.Wrap(err, "invalid BucketURL")
		return
	}

	b := &cos.BaseURL{BucketURL: u}
	storage.Client = cos.NewClient(b, &http.Client{
		//设置超时时间
		Timeout: 100 * time.Second,
		Transport: &cos.AuthorizationTransport{
			//如实填写账号和密钥，也可以设置为环境变量
			SecretID:  conf.SecretID,
			SecretKey: conf.SecretKey,
		},
	})

	storage.Config = conf

	return
}

func (o COSStorage) SignURL(objectKey string, method constants.HTTPMethod, expiredInSec int64) (signedURL string, err error) {
	contentType, err := GetContentType(objectKey)
	if err != nil {
		err = errors.Wrap(err, "GetContentType failed")
		return
	}

	opt := &cos.PresignedURLOptions{
		Header: &http.Header{},
	}
	opt.Header.Set("Content-Type", contentType)

	u, err := o.Client.Object.GetPresignedURL(
		context.Background(),
		string(method),
		objectKey,
		o.Config.SecretID,
		o.Config.SecretKey,
		time.Duration(expiredInSec)*time.Second,
		nil,
	)
	if err != nil {
		err = errors.Wrap(err, "GetPresignedURL failed")
		return
	}

	if o.Config.CdnURL != "" {
		cdnURL, err := url.Parse(o.Config.CdnURL)
		if err != nil {
			err = errors.Wrap(err, "url.ParseLink failed")
			return signedURL, err
		}

		u.Host = cdnURL.Host
		u.Scheme = cdnURL.Scheme
	}

	signedURL = u.String()

	return
}

func (o COSStorage) Get(objectKey string) (content io.ReadCloser, err error) {
	resp, err := o.Client.Object.Get(context.Background(), objectKey, nil)
	if err != nil {
		err = errors.Wrap(err, "GetObject failed")
		return
	}

	return resp.Body, nil
}

func (o COSStorage) Put(objectKey string, reader io.Reader) (err error) {
	contentType, err := GetContentType(objectKey)
	if err != nil {
		err = errors.Wrap(err, "GetContentType failed")
		return
	}

	opt := &cos.ObjectPutOptions{
		ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{
			ContentType: contentType,
		},
		ACLHeaderOptions: &cos.ACLHeaderOptions{
			XCosACL: "private",
		},
	}
	_, err = o.Client.Object.Put(context.Background(), objectKey, reader, opt)
	if err != nil {
		err = errors.Wrap(err, "PutObject failed")
		return
	}

	return
}

func (o COSStorage) IsExist(objectKey string) (ok bool, err error) {
	_, err = o.Client.Object.Head(context.Background(), objectKey, nil)
	if err != nil {
		err = errors.Wrap(err, "Head failed")
		return
	}
	return
}

func (o COSStorage) PutFromFile(objectKey string, filePath string) (err error) {
	ext := strings.ToLower(filepath.Ext(filePath))
	if ext == "" {
		err = errors.New("file ext is required")
		return
	}

	contentType := mime.TypeByExtension(ext)
	if contentType == "" {
		err = errors.New("invalid file ext")
		return
	}

	opt := &cos.ObjectPutOptions{
		ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{
			ContentType: contentType,
		},
		ACLHeaderOptions: &cos.ACLHeaderOptions{
			XCosACL: "private",
		},
	}

	_, err = o.Client.Object.PutFromFile(context.Background(), objectKey, filePath, opt)
	if err != nil {
		err = errors.Wrap(err, "PutFromFile failed")
		return
	}

	return
}

func (o COSStorage) Delete(objectKeys ...string) (deletedObjects []string, err error) {
	objects := make([]cos.Object, 0)
	for _, key := range objectKeys {
		objects = append(objects, cos.Object{
			Key: key,
		})
	}
	opt := &cos.ObjectDeleteMultiOptions{
		Objects: objects,
	}

	result, _, err := o.Client.Object.DeleteMulti(context.Background(), opt)
	if err != nil {
		err = errors.Wrap(err, "DeleteMulti failed")
		return
	}

	for _, object := range result.DeletedObjects {
		deletedObjects = append(deletedObjects, object.Key)
	}

	return
}
