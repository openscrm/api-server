package workwx

// Tag 创建标签
// 文档：https://open.work.weixin.qq.com/api/doc/90000/90135/90210#创建标签
type Tag struct {
	TagID   int    `json:"tagid"`   // 非必填，标签id，非负整型，指定此参数时新增的标签会生成对应的标签id，不指定时则以目前最大的id自增。
	TagName string `json:"tagname"` // 必填，标签名称，长度限制为32个字以内（汉字或英文字母），标签名不可与其他标签重名
}

// TagUser 标签关联成员
type TagUser struct {
	// Name 成员名称，此字段从2019年12月30日起，对新创建第三方应用不再返回，2020年6月30日起，对所有历史第三方应用不再返回，后续第三方仅通讯录应用可获取，第三方页面需要通过<a href="#17172">通讯录展示组件</a>来展示名字
	Name string `json:"name"`
	// Userid 成员帐号
	Userid string `json:"userid"`
}

// TagDetail 标签详情
// 文档：https://open.work.weixin.qq.com/api/doc/90000/90135/90213#获取标签成员
type TagDetail struct {
	// PartyList 标签中包含的部门id列表
	PartyList []int `json:"partylist"`
	// TagName 标签名
	TagName  string    `json:"tagname"`
	UserList []TagUser `json:"userlist"` //标签中包含的成员列表
}

// AddTagUsersReq 增加标签成员
// 文档：https://open.work.weixin.qq.com/api/doc/90000/90135/90214#增加标签成员
type AddTagUsersReq struct {
	// PartyList 企业部门ID列表，注意:userlist、partylist不能同时为空，单次请求个数不超过100
	PartyList []int `json:"partylist,omitempty"`
	// TagID 标签ID，必填
	TagID int `json:"tagid"`
	// UserList 企业成员ID列表，注意:userlist、partylist不能同时为空，单次请求个数不超过1000
	UserList []string `json:"userlist,omitempty"`
}

// DelTagUsersReq 删除标签成员请求
// 文档：https://open.work.weixin.qq.com/api/doc/90000/90135/90215#删除标签成员
type DelTagUsersReq struct {
	// PartyList 企业部门ID列表，注意:userlist、partylist不能同时为空，单次请求个数不超过100
	PartyList []int `json:"partylist,omitempty"`
	// TagID 标签ID，必填
	TagID int `json:"tagid"`
	// UserList 企业成员ID列表，注意:userlist、partylist不能同时为空，单次请求个数不超过1000
	UserList []string `json:"userlist,omitempty"`
}
