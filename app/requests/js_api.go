package requests

type GetJSConfigReq struct {
	URL string `json:"url" form:"url" validate:"required"`
}

type GetJSAgentConfigReq struct {
	URL string `json:"url" form:"url" validate:"required"`
}
