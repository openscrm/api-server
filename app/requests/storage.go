package requests

type GetUploadURLReq struct {
	// 文件名
	Filename string `json:"file_name" validate:"required"`
}

// GetUploadURLResp 上传文件地址
type GetUploadURLResp struct {
	// 上传地址
	UploadURL string `json:"upload_url"`
	// 下载地址
	DownloadURL string `json:"download_url"`
}
