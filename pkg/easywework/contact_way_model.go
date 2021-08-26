package workwx

// ContactWayScene 渠道码场景
// 1-在小程序中联系
// 2-通过二维码联系
type ContactWayScene int

const (
	// ContactWaySceneMicroApp 在小程序中联系
	ContactWaySceneMicroApp ContactWayScene = 1
	// ContactWaySceneQrcode 通过二维码联系
	ContactWaySceneQrcode ContactWayScene = 2
)

// ContactWayType 渠道码联系方式
// 1-单人
// 2-多人
type ContactWayType int

const (
	// ContactWayTypeSingle 单人
	ContactWayTypeSingle ContactWayType = 1
	// ContactWayTypeMultiple 多人
	ContactWayTypeMultiple ContactWayType = 2
)

// ConclusionsReq 请求时的结束语配置，会话结束时自动发送给客户，可参考“<a href="#15645/结束语定义">结束语定义</a>”，仅在is_temp为true时有效
type ConclusionsReq struct {
	Image struct {
		MediaID string `json:"media_id"`
	} `json:"image"`
	Link struct {
		Desc   string `json:"desc"`
		Picurl string `json:"picurl"`
		Title  string `json:"title"`
		URL    string `json:"url"`
	} `json:"link"`
	Miniprogram struct {
		Appid      string `json:"appid"`
		Page       string `json:"page"`
		PicMediaID string `json:"pic_media_id"`
		Title      string `json:"title"`
	} `json:"miniprogram"`
	Text struct {
		Content string `json:"content"`
	} `json:"text"`
}

// AddContactWay 配置客户联系「联系我」方式
// 文档：https://work.weixin.qq.com/api/doc/90000/90135/92572#配置客户联系「联系我」方式
type AddContactWay struct {
	// ChatExpiresIn 临时会话有效期，以秒为单位。该参数仅在is_temp为true时有效，默认为添加好友后24小时
	ChatExpiresIn int `json:"chat_expires_in,omitempty"`
	// Conclusions 结束语，会话结束时自动发送给客户，可参考“<a href="#15645/结束语定义">结束语定义</a>”，仅在is_temp为true时有效
	Conclusions ConclusionsReq `json:"conclusions,omitempty"`
	// ExpiresIn 临时会话二维码有效期，以秒为单位。该参数仅在is_temp为true时有效，默认7天
	ExpiresIn int `json:"expires_in,omitempty"`
	// IsTemp 是否临时会话模式，true表示使用临时会话模式，默认为false
	IsTemp bool `json:"is_temp,omitempty"`
	// Party 使用该联系方式的部门id列表，只在type为2时有效
	Party []int `json:"party,omitempty"`
	// Remark 联系方式的备注信息，用于助记，不超过30个字符
	Remark string `json:"remark,omitempty"`
	// Scene 场景，1-在小程序中联系，2-通过二维码联系，必填
	Scene ContactWayScene `json:"scene"`
	// SkipVerify 外部客户添加时是否无需验证，默认为true
	SkipVerify bool `json:"skip_verify,omitempty"`
	// State 企业自定义的state参数，用于区分不同的添加渠道，在调用“获取外部联系人详情”时会返回该参数值，不超过30个字符
	State string `json:"state,omitempty"`
	// Style 在小程序中联系时使用的控件样式，详见附表
	Style int `json:"style,omitempty"`
	// Type 联系方式类型,1-单人, 2-多人，必填
	Type ContactWayType `json:"type"`
	// Unionid 可进行临时会话的客户unionid，该参数仅在is_temp为true时有效，如不指定则不进行限制
	Unionid string `json:"unionid,omitempty"`
	// User 使用该联系方式的用户userID列表，在type为1时为必填，且只能有一个
	User []string `json:"user,omitempty"`
}

// Conclusions 结束语配置，会话结束时自动发送给客户，可参考“<a href="#15645/结束语定义">结束语定义</a>”，仅在is_temp为true时有效
type Conclusions struct {
	Image struct {
		PicURL string `json:"pic_url"`
	} `json:"image"`
	Link struct {
		Desc   string `json:"desc"`
		Picurl string `json:"picurl"`
		Title  string `json:"title"`
		URL    string `json:"url"`
	} `json:"link"`
	Miniprogram struct {
		Appid      string `json:"appid"`
		Page       string `json:"page"`
		PicMediaID string `json:"pic_media_id"`
		Title      string `json:"title"`
	} `json:"miniprogram"`
	Text struct {
		Content string `json:"content"`
	} `json:"text"`
}

// ContactWay 渠道码
type ContactWay struct {
	// ChatExpiresIn 临时会话有效期，以秒为单位
	ChatExpiresIn int `json:"chat_expires_in"`
	// Conclusions 结束语，可参考“<a href="#15645/结束语定义">结束语定义</a>”
	Conclusions Conclusions `json:"conclusions"`
	// ConfigID 新增联系方式的配置id
	ConfigID string `json:"config_id"`
	// ExpiresIn 临时会话二维码有效期，以秒为单位
	ExpiresIn int `json:"expires_in"`
	// IsTemp 是否临时会话模式，默认为false，true表示使用临时会话模式
	IsTemp bool `json:"is_temp"`
	// Party 使用该联系方式的部门id列表
	Party []int `json:"party"`
	// QrCode 联系二维码的URL，仅在scene为2时返回
	QrCode string `json:"qr_code"`
	// Remark 联系方式的备注信息，用于助记
	Remark string `json:"remark"`
	// Scene 场景，1-在小程序中联系，2-通过二维码联系
	Scene ContactWayScene `json:"scene"`
	// SkipVerify 外部客户添加时是否无需验证
	SkipVerify bool `json:"skip_verify"`
	// State 企业自定义的state参数，用于区分不同的添加渠道，在调用“<a href="#13878">获取外部联系人详情</a>”时会返回该参数值
	State string `json:"state"`
	// Style 小程序中联系按钮的样式，仅在scene为1时返回，详见附录
	Style int `json:"style"`
	// Type 联系方式类型，1-单人，2-多人
	Type ContactWayType `json:"type"`
	// Unionid 可进行临时会话的客户unionid
	Unionid string `json:"unionid"`
	// User 使用该联系方式的用户userID列表
	User []string `json:"user"`
}

// UpdateContactWay 更新企业已配置的「联系我」方式
// 文档：https://work.weixin.qq.com/api/doc/90000/90135/92572#更新企业已配置的「联系我」方式
type UpdateContactWay struct {
	// ChatExpiresIn 临时会话有效期，以秒为单位，该参数仅在临时会话模式下有效
	ChatExpiresIn int `json:"chat_expires_in,omitempty"`
	// Conclusions 结束语，会话结束时自动发送给客户，可参考“<a href="#15645/结束语定义">结束语定义</a>”，仅临时会话模式（is_temp为true）可设置
	Conclusions ConclusionsReq `json:"conclusions,omitempty"`
	// ConfigID 企业联系方式的配置id，必填
	ConfigID string `json:"config_id"`
	// ExpiresIn 临时会话二维码有效期，以秒为单位，该参数仅在临时会话模式下有效
	ExpiresIn int `json:"expires_in,omitempty"`
	// Party 使用该联系方式的部门列表，将覆盖原有部门列表，只在配置的type为2时有效
	Party []int `json:"party,omitempty"`
	// Remark 联系方式的备注信息，不超过30个字符，将覆盖之前的备注
	Remark string `json:"remark,omitempty"`
	// SkipVerify 外部客户添加时是否无需验证
	SkipVerify bool `json:"skip_verify,omitempty"`
	// State 企业自定义的state参数，用于区分不同的添加渠道，在调用“<a href="#13878">获取外部联系人详情</a>”时会返回该参数值
	State string `json:"state,omitempty"`
	// Style 样式，只针对“在小程序中联系”的配置生效
	Style int `json:"style,omitempty"`
	// Unionid 可进行临时会话的客户unionid，该参数仅在临时会话模式有效，如不指定则不进行限制
	Unionid string `json:"unionid,omitempty"`
	// User 使用该联系方式的用户列表，将覆盖原有用户列表
	User []string `json:"user,omitempty"`
}
