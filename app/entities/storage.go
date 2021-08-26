package entities

// LocalStorageFileReq 本地存储请求
type LocalStorageFileReq struct {
	// ExpireAt 过期时间
	ExpireAt int64 `form:"expire_at" json:"expire_at" validate:"required"`
	// Signature hmac签名
	Signature string `form:"signature" json:"signature" validate:"required,alphanum"`
}

type GetSignedURLReq struct {
	ObjectKey    string `json:"object_key" validate:"required"`
	Method       string `json:"method" validate:"required,oneof=GET PUT"`
	ExpiredInSec int64  `json:"expired_in_sec" validate:"required"`
}
