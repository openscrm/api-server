package workwx

import (
	"strconv"
	"time"
)

// mediaUploadResp 临时素材上传响应
type mediaUploadResp struct {
	CommonResp

	Type      string `json:"type"`
	MediaID   string `json:"media_id"`
	CreatedAt string `json:"created_at"`
}

func (x mediaUploadResp) intoMediaUploadResult() (MediaUploadResult, error) {
	createdAtInt, err := strconv.ParseInt(x.CreatedAt, 10, 64)
	if err != nil {
		return MediaUploadResult{}, err
	}
	createdAt := time.Unix(createdAtInt, 0)

	return MediaUploadResult{
		Type:      x.Type,
		MediaID:   x.MediaID,
		CreatedAt: createdAt,
	}, nil
}

// MediaUploadResult 临时素材上传结果
type MediaUploadResult struct {
	// Type 媒体文件类型，分别有图片（image）、语音（voice）、视频（video），普通文件(file)
	Type string
	// MediaID 媒体文件上传后获取的唯一标识，3天内有效
	MediaID string
	// CreatedAt 媒体文件上传时间戳
	CreatedAt time.Time
}
