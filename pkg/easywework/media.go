package workwx

import (
	"bytes"
	"io"
	"mime/multipart"
	"os"
)

const mediaFieldName = "media"

// Media 欲上传的素材

type Media struct {
	filename string
	filesize int64
	stream   io.Reader
}

// NewMediaFromFile 从操作系统级文件创建一个欲上传的素材对象
func NewMediaFromFile(f *os.File) (*Media, error) {
	stat, err := f.Stat()
	if err != nil {
		return nil, err
	}

	return &Media{
		filename: stat.Name(),
		filesize: stat.Size(),
		stream:   f,
	}, nil
}

// NewMediaFromBuffer 从内存创建一个欲上传的素材对象
func NewMediaFromBuffer(filename string, buf []byte) (*Media, error) {
	stream := bytes.NewReader(buf)
	return &Media{
		filename: filename,
		filesize: int64(len(buf)),
		stream:   stream,
	}, nil
}

func (m *Media) writeTo(w *multipart.Writer) error {
	wr, err := w.CreateFormFile(mediaFieldName, m.filename)
	if err != nil {
		return err
	}

	_, err = io.Copy(wr, m.stream)
	if err != nil {
		return err
	}

	return nil
}

// UploadPermanentImageMedia 上传永久图片素材
func (c *App) UploadPermanentImageMedia(media *Media) (url string, err error) {
	url, err = c.mediaUploadImg(media)
	if err != nil {
		return "", err
	}

	return url, nil
}

const (
	tempMediaTypeImage = "image"
	tempMediaTypeVoice = "voice"
	tempMediaTypeVideo = "video"
	tempMediaTypeFile  = "file"
)

// UploadTempImageMedia 上传临时图片素材
func (c *App) UploadTempImageMedia(media *Media) (*MediaUploadResult, error) {
	result, err := c.mediaUpload(tempMediaTypeImage, media)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// UploadTempVoiceMedia 上传临时语音素材
func (c *App) UploadTempVoiceMedia(media *Media) (*MediaUploadResult, error) {
	result, err := c.mediaUpload(tempMediaTypeVoice, media)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// UploadTempVideoMedia 上传临时视频素材
func (c *App) UploadTempVideoMedia(media *Media) (*MediaUploadResult, error) {
	result, err := c.mediaUpload(tempMediaTypeVideo, media)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// UploadTempFileMedia 上传临时文件素材
func (c *App) UploadTempFileMedia(media *Media) (*MediaUploadResult, error) {
	result, err := c.mediaUpload(tempMediaTypeFile, media)
	if err != nil {
		return nil, err
	}

	return result, nil
}
