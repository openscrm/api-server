package app

type JSONResult struct {
	// Code 响应状态码
	Code int `json:"code"`
	// Message 响应消息
	Message string `json:"message"`
	// Data 响应数据
	Data interface{} `json:"data"`
}

type ItemsData struct {
	// Items 数据列表
	Items interface{} `json:"items"`
	// Pager 列表分页信息
	Pager Pager `json:"pager"`
}
