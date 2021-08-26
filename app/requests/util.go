package requests

type ParseLinkReq struct {
	// 链接
	URL string `json:"url" form:"url" validate:"required,url"`
}

type UploadMediaReq struct {
	// 文件类型
	Type string `json:"type" form:"type" validate:"oneof=image voice video file"`
	// 文件链接
	URL string `json:"url" form:"url" validate:"required,url"`
}
