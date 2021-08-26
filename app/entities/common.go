package entities

// ActionReq 根据ID进行操作公共的请求
type ActionReq struct {
	// ID
	IDs []string `json:"ids" validate:"div,int64"`
}
