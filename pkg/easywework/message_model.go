package workwx

// TaskCardBtn 任务卡片消息按钮
type TaskCardBtn struct {
	// Key 按钮key值，用户点击后，会产生任务卡片回调事件，回调事件会带上该key值，只能由数字、字母和“_-@”组成，最长支持128字节
	Key string `json:"key"`
	// Name 按钮名称
	Name string `json:"name"`
	// ReplaceName 点击按钮后显示的名称，默认为“已处理”
	ReplaceName string `json:"replace_name"`
	// Color 按钮字体颜色，可选“red”或者“blue”,默认为“blue”
	Color string `json:"color"`
	// IsBold 按钮字体是否加粗，默认false
	IsBold bool `json:"is_bold"`
}
