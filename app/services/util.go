package services

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-resty/resty/v2"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/crypto/gsha1"
	"github.com/gogf/gf/util/grand"
	"github.com/pkg/errors"
	"github.com/thoas/go-funk"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"openscrm/app/requests"
	"openscrm/app/responses"
	"openscrm/common/ecode"
	"openscrm/common/log"
	"openscrm/common/redis"
	"openscrm/common/storage"
	"openscrm/common/we_work"
	"openscrm/pkg/easywework"
	"path"
	"strings"
	"time"
)

type Util struct {
}

// CachedUploadMedia  上传临时素材(带缓存)
func (o *Util) CachedUploadMedia(req requests.UploadMediaReq, extCorpID string) (result responses.UploadMediaResult, err error) {
	key := gsha1.Encrypt(req.URL)
	err = redis.GetOrSetFunc(key, func() (interface{}, error) {
		return o.UploadMedia(req, extCorpID)
	}, time.Hour*24*2, &result)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	return
}

// UploadMedia  上传临时素材
func (o *Util) UploadMedia(req requests.UploadMediaReq, extCorpID string) (result responses.UploadMediaResult, err error) {
	c := resty.New()

	urlInfo, err := url.Parse(req.URL)
	if err != nil {
		err = errors.WithStack(ecode.IllegalURL)
		return
	}

	client, err := we_work.Clients.Get(extCorpID)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	// 下载文件
	res, err := c.R().Get(req.URL)
	if err != nil {
		err = errors.WithStack(ecode.IllegalURL)
		return
	}

	filename := fmt.Sprintf("%s%s", gsha1.Encrypt(time.Now().UnixNano()+int64(grand.Intn(math.MaxInt32))), path.Ext(urlInfo.Path))

	media, err := workwx.NewMediaFromBuffer(filename, res.Body())
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	uploadRes := &workwx.MediaUploadResult{}
	if req.Type == "file" {
		uploadRes, err = client.Customer.UploadTempFileMedia(media)
		if err != nil {
			err = errors.WithStack(err)
			return
		}
	}

	if req.Type == "image" {
		uploadRes, err = client.Customer.UploadTempImageMedia(media)
		if err != nil {
			err = errors.WithStack(err)
			return
		}
	}

	if req.Type == "video" {
		uploadRes, err = client.Customer.UploadTempVideoMedia(media)
		if err != nil {
			err = errors.WithStack(err)
			return
		}
	}

	if req.Type == "voice" {
		uploadRes, err = client.Customer.UploadTempVoiceMedia(media)
		if err != nil {
			err = errors.WithStack(err)
			return
		}
	}

	result.MediaID = uploadRes.MediaID
	result.CreatedAt = uploadRes.CreatedAt
	result.Type = uploadRes.Type

	return
}

// ParseLink  解析指定URL页面
func (o *Util) ParseLink(req requests.ParseLinkReq) (result responses.ParseLinkResp, err error) {
	client := resty.New().
		SetRetryCount(0).
		SetTimeout(time.Second * 15).
		SetHeaders(map[string]string{
			http.CanonicalHeaderKey("User-Agent"): "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36",
		}).
		SetRedirectPolicy(resty.FlexibleRedirectPolicy(3))

	res, err := client.R().Get(req.URL)
	if err != nil {
		err = errors.WithStack(ecode.IllegalURL)
		return
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(res.Body()))
	if err != nil {
		err = errors.Wrap(err, "parse html failed")
		return
	}

	result.LinkURL = req.URL
	result.Title = doc.Find("head > title").First().Text()
	result.Desc, _ = doc.Find("head > meta[name=description]").First().Attr("content")
	result.ImgURL, _ = doc.Find("body img").First().Attr("src")

	// 处理相对路径
	if strings.HasPrefix(result.ImgURL, "/") || strings.HasPrefix(result.ImgURL, ".") {
		result.ImgURL = req.URL + result.ImgURL
	}

	// 下载图片
	if result.ImgURL != "" {
		urlInfo, ignoredErr := url.Parse(result.ImgURL)
		if ignoredErr != nil {
			log.Sugar.Warnw("parse url failed", "err", ignoredErr)
			return
		}

		ext := strings.ToLower(path.Ext(urlInfo.Path))
		if !funk.ContainsString([]string{".jpg", ".png", ".jpeg"}, ext) {
			log.Sugar.Warnw("invalid file ext", "ext", ext)
			return
		}

		imgRes, err := client.R().Get(result.ImgURL)
		if err != nil {
			err = errors.Wrap(err, "save image failed")
			return result, err
		}

		rand.Seed(time.Now().UnixNano())
		object := `public/link-info/assets/images/` + gmd5.MustEncryptString(result.ImgURL) + fmt.Sprintf("%d.%s", rand.Int63(), ext)

		err = storage.FileStorage.Put(object, bytes.NewReader(imgRes.Body()))
		if err != nil {
			err = errors.Wrap(err, "put File failed")
			return result, err
		}

		result.ImgURL, err = storage.FileStorage.SignURL(object, http.MethodGet, 86400*365*10)
		if err != nil {
			err = errors.Wrap(err, "SignURL failed")
			return result, err
		}
	}

	return
}
