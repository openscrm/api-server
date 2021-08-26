package workwx

import "net/url"

// mediaUploadReq 临时素材上传请求
type mediaUploadReq struct {
	Type  string
	Media *Media
}

var _ urlValuer = mediaUploadReq{}
var _ mediaUploader = mediaUploadReq{}

func (x mediaUploadReq) intoURLValues() url.Values {
	return url.Values{
		"type": {x.Type},
	}
}

func (x mediaUploadReq) getMedia() *Media {
	return x.Media
}

// mediaUploadImgReq 永久图片素材上传请求
type mediaUploadImgReq struct {
	Media *Media
}

var _ urlValuer = mediaUploadImgReq{}
var _ mediaUploader = mediaUploadImgReq{}

func (x mediaUploadImgReq) intoURLValues() url.Values {
	return url.Values{}
}

func (x mediaUploadImgReq) getMedia() *Media {
	return x.Media
}

// mediaUploadImgResp 永久图片素材上传响应
type mediaUploadImgResp struct {
	CommonResp
	URL string `json:"url"`
}

// execMediaUploadImg 上传永久图片
func (c *App) execMediaUploadImg(req mediaUploadImgReq) (mediaUploadImgResp, error) {
	var resp mediaUploadImgResp
	err := c.executeWXApiMediaUpload("/cgi-bin/media/uploadimg", req, &resp, true)
	if err != nil {
		return mediaUploadImgResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return mediaUploadImgResp{}, bizErr
	}

	return resp, nil
}

// mediaUploadImg 上传永久图片
func (c *App) mediaUploadImg(media *Media) (url string, err error) {
	resp, err := c.execMediaUploadImg(mediaUploadImgReq{
		Media: media,
	})
	if err != nil {
		return "", err
	}

	return resp.URL, nil
}

// mediaUpload 上传临时素材
func (c *App) mediaUpload(typ string, media *Media) (*MediaUploadResult, error) {
	resp, err := c.execMediaUpload(mediaUploadReq{
		Type:  typ,
		Media: media,
	})
	if err != nil {
		return nil, err
	}

	obj, err := resp.intoMediaUploadResult()
	if err != nil {
		return nil, err
	}

	return &obj, nil
}

// execMediaUpload 上传临时素材
func (c *App) execMediaUpload(req mediaUploadReq) (mediaUploadResp, error) {
	var resp mediaUploadResp
	err := c.executeWXApiMediaUpload("/cgi-bin/media/upload", req, &resp, true)
	if err != nil {
		return mediaUploadResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return mediaUploadResp{}, bizErr
	}

	return resp, nil
}
