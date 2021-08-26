package storage

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/pkg/errors"
	"io"
	"net/url"
	"openscrm/app/constants"
	setting "openscrm/conf"
)

type OSSStorage struct {
	Client *oss.Client
	Bucket *oss.Bucket
	Config setting.StorageConfig
}

func NewOSS(conf setting.StorageConfig) (ossStorage OSSStorage, err error) {
	ossStorage.Client, err = oss.New(conf.EndPoint, conf.AccessKeyId, conf.AccessKeySecret)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	ossStorage.Bucket, err = ossStorage.Client.Bucket(conf.Bucket)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	ossStorage.Config = conf

	return
}

func (o OSSStorage) SignURL(objectKey string, method constants.HTTPMethod, expiredInSec int64) (signedURL string, err error) {
	contentType, err := GetContentType(objectKey)
	if err != nil {
		err = errors.Wrap(err, "GetContentType failed")
		return
	}
	options := []oss.Option{
		oss.ContentType(contentType),
	}
	if method == constants.HTTPGet {
		options = nil
	}

	signedURL, err = o.Bucket.SignURL(objectKey, oss.HTTPMethod(method), expiredInSec, options...)
	if err != nil {
		err = errors.Wrap(err, "Bucket.SignURL failed")
		return
	}

	if o.Config.CdnURL != "" {
		fileURL, err := url.Parse(signedURL)
		if err != nil {
			err = errors.Wrap(err, "url.ParseLink failed")
			return signedURL, err
		}

		cdnURL, err := url.Parse(o.Config.CdnURL)
		if err != nil {
			err = errors.Wrap(err, "url.ParseLink failed")
			return signedURL, err
		}

		fileURL.Host = cdnURL.Host
		fileURL.Scheme = cdnURL.Scheme
		signedURL = fileURL.String()
	}

	return
}

func (o OSSStorage) Get(objectKey string) (content io.ReadCloser, err error) {
	content, err = o.Bucket.GetObject(objectKey)
	if err != nil {
		err = errors.Wrap(err, "GetObject failed")
		return
	}

	return
}

func (o OSSStorage) Put(objectKey string, reader io.Reader) (err error) {
	contentType, err := GetContentType(objectKey)
	if err != nil {
		err = errors.Wrap(err, "GetContentType failed")
		return
	}

	options := []oss.Option{
		oss.ContentType(contentType),
	}

	err = o.Bucket.PutObject(objectKey, reader, options...)
	if err != nil {
		err = errors.Wrap(err, "PutObject failed")
		return
	}

	return
}

func (o OSSStorage) IsExist(objectKey string) (ok bool, err error) {
	ok, err = o.Bucket.IsObjectExist(objectKey)
	if err != nil {
		err = errors.Wrap(err, "IsObjectExist failed")
		return
	}

	return
}

func (o OSSStorage) PutFromFile(objectKey string, filePath string) (err error) {
	err = o.Bucket.PutObjectFromFile(objectKey, filePath)
	if err != nil {
		err = errors.Wrap(err, "PutObjectFromFile failed")
		return
	}

	return
}

func (o OSSStorage) Delete(objectKeys ...string) (deletedObjects []string, err error) {
	result, err := o.Bucket.DeleteObjects(objectKeys)
	if err != nil {
		err = errors.Wrap(err, "DeleteObjects failed")
		return
	}

	deletedObjects = result.DeletedObjects
	return
}
