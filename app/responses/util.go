package responses

import "time"

type ParseLinkResp struct {
	// 标题
	Title string `json:"title"`
	// 描述
	Desc string `json:"desc"`
	// 封面图片url
	ImgURL string `json:"img_url"`
	// 链接url
	LinkURL string `json:"link_url"`
}

type UploadMediaResult struct {
	// Type 媒体文件类型，分别有图片（image）、语音（voice）、视频（video），普通文件(file)
	Type string `json:"type"`
	// MediaID 媒体文件上传后获取的唯一标识，3天内有效
	MediaID string `json:"media_id"`
	// CreatedAt 媒体文件上传时间戳
	CreatedAt time.Time `json:"created_at"`
}
