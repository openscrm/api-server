package workwx

import (
	"encoding/json"
	"net/url"
	"strconv"
)

// UpdateUserReq 更新成员请求
// 文档：https://work.weixin.qq.com/api/doc/90000/90135/90197#更新成员
type UpdateUserReq struct {
	// Address 地址。长度最大128个字符
	Address string `json:"address,omitempty"`
	// Alias 别名。长度为1-32个utf8字符
	Alias string `json:"alias,omitempty"`
	// AvatarMediaid 成员头像的mediaid，通过<a href="#10112">素材管理</a>接口上传图片获得的mediaid
	AvatarMediaid string `json:"avatar_mediaid,omitempty"`
	// Department 成员所属部门id列表，不超过100个
	Department []int `json:"department,omitempty"`
	// Email 邮箱。长度不超过64个字节，且为有效的email格式。企业内必须唯一。<font color="red">若是绑定了腾讯企业邮箱的企业微信，则需要在腾讯企业邮箱中修改邮箱（此情况下该参数被忽略，但不会报错）</font>
	Email string `json:"email,omitempty"`
	// Enable 启用/禁用成员。1表示启用成员，0表示禁用成员
	Enable  int `json:"enable,omitempty"`
	Extattr struct {
		Attrs []struct {
			// Name 成员名称。长度为1~64个utf8字符
			Name string `json:"name,omitempty"`
			Text struct {
				Value string `json:"value"`
			} `json:"text"`
			Type int `json:"type"`
			Web  struct {
				Title string `json:"title"`
				URL   string `json:"url"`
			} `json:"web"`
		} `json:"attrs"`
	} `json:"extattr,omitempty"` //自定义字段。自定义字段需要先在WEB管理端添加，见<a href="#10016/扩展属性的添加方法">扩展属性添加方法</a>，否则忽略未知属性的赋值。与对外属性一致，不过只支持type=0的文本和type=1的网页类型，详细描述查看<a href="#13450">对外属性</a>
	// ExternalPosition 对外职务，如果设置了该值，则以此作为对外展示的职务，否则以position来展示。不超过12个汉字
	ExternalPosition string `json:"external_position,omitempty"`
	ExternalProfile  struct {
		ExternalAttr []struct {
			Miniprogram struct {
				Appid    string `json:"appid"`
				Pagepath string `json:"pagepath"`
				Title    string `json:"title"`
			} `json:"miniprogram"`
			// Name 成员名称。长度为1~64个utf8字符
			Name string `json:"name,omitempty"`
			Text struct {
				Value string `json:"value"`
			} `json:"text"`
			Type int `json:"type"`
			Web  struct {
				Title string `json:"title"`
				URL   string `json:"url"`
			} `json:"web"`
		} `json:"external_attr"`
		ExternalCorpName string `json:"external_corp_name"`
	} `json:"external_profile,omitempty"` //成员对外属性，字段详情见<a href="#13450">对外属性</a>
	// Gender 性别。1表示男性，2表示女性
	Gender string `json:"gender,omitempty"`
	// IsLeaderInDept 上级字段，个数必须和department一致，表示在所在的部门内是否为上级。
	IsLeaderInDept []int `json:"is_leader_in_dept,omitempty"`
	// MainDepartment 主部门
	MainDepartment int `json:"main_department,omitempty"`
	// Mobile 手机号码。企业内必须唯一。<font color="red">若成员已激活企业微信，则需成员自行修改（此情况下该参数被忽略，但不会报错） </font>
	Mobile string `json:"mobile,omitempty"`
	// Name 成员名称。长度为1~64个utf8字符
	Name string `json:"name,omitempty"`
	// Order 部门内的排序值，默认为0。当有传入department时有效。数量必须和department一致，数值越大排序越前面。有效的值范围是[0, 2^32)
	Order []int `json:"order,omitempty"`
	// Position 职务信息。长度为0~128个字符
	Position string `json:"position,omitempty"`
	// Telephone 座机。由1-32位的纯数字或’-‘号组成
	Telephone string `json:"telephone,omitempty"`
	// Userid 成员UserID。对应管理端的帐号，企业内必须唯一。不区分大小写，长度为1~64个字节，必填
	Userid string `json:"userid"`
}

var _ bodyer = UpdateUserReq{}

func (x UpdateUserReq) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// updateUserResp 更新成员响应
// 文档：https://work.weixin.qq.com/api/doc/90000/90135/90197#更新成员
type updateUserResp struct {
	CommonResp
}

var _ bodyer = updateUserResp{}

func (x updateUserResp) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// execUpdateUser 更新成员
// 文档：https://work.weixin.qq.com/api/doc/90000/90135/90197#更新成员
func (c *App) execUpdateUser(req UpdateUserReq) (updateUserResp, error) {
	var resp updateUserResp
	err := c.executeWXApiJSONPost("/cgi-bin/user/update", req, &resp, true)
	if err != nil {
		return updateUserResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return updateUserResp{}, bizErr
	}

	return resp, nil
}

// userGetResp 读取成员响应
type userGetResp struct {
	CommonResp

	userDetailResp
}

// userListReq 部门成员请求
type userListReq struct {
	DeptID     int64
	FetchChild bool
}

var _ urlValuer = userListReq{}

func (x userListReq) intoURLValues() url.Values {
	var fetchChild int64
	if x.FetchChild {
		fetchChild = 1
	}

	return url.Values{
		"department_id": {strconv.FormatInt(x.DeptID, 10)},
		"fetch_child":   {strconv.FormatInt(fetchChild, 10)},
	}
}

// usersByDeptIDResp 部门成员详情响应
type userListResp struct {
	CommonResp

	Users []*userDetailResp `json:"userlist"`
}

// userIDByMobileReq 手机号获取 userid 请求
type userIDByMobileReq struct {
	Mobile string `json:"mobile"`
}

var _ bodyer = userIDByMobileReq{}

func (x userIDByMobileReq) intoBody() ([]byte, error) {
	body, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// userIDByMobileResp 手机号获取 userid 响应
type userIDByMobileResp struct {
	CommonResp

	UserID string `json:"userid"`
}

var _ urlValuer = userInfoGetReq{}

func (x userInfoGetReq) intoURLValues() url.Values {
	return url.Values{
		"code": {x.Code},
	}
}

// userInfoGetResp 部门列表响应
type userInfoGetResp struct {
	CommonResp
	UserIdentityInfo
}

// userDetailResp 成员详细信息的公共字段
type userDetailResp struct {
	UserID         string   `json:"userid"`
	Name           string   `json:"name"`
	DeptIDs        []int64  `json:"department"`
	DeptOrder      []uint32 `json:"order"`
	IsLeaderInDept []int    `json:"is_leader_in_dept"`
	Position       string   `json:"position"`
	Mobile         string   `json:"mobile"`
	Gender         string   `json:"gender"`
	Email          string   `json:"email"`
	AvatarURL      string   `json:"avatar"`
	Telephone      string   `json:"telephone"`
	IsEnabled      int      `json:"enable"`
	Alias          string   `json:"alias"`
	Status         int      `json:"status"`
	QRCodeURL      string   `json:"qr_code"`
	// TODO: extattr external_profile external_position
	Extattr Extattr `json:"extattr"`
}

type Extattr struct {
	Attrs []Attrs `json:"attrs"`
}

type Attrs struct {
	Type int    `json:"type"`
	Name string `json:"name"`
	Text Texts  `json:"text"`
	Web  Web    `json:"web"`
}

type Web struct {
	Url   string `json:"url"`
	Title string `json:"title"`
}

type Texts struct {
	Value string `json:"value"`
}

type deptListReq struct {
	HaveID bool
	ID     int64
}

var _ urlValuer = deptListReq{}

func (x deptListReq) intoURLValues() url.Values {
	if !x.HaveID {
		return url.Values{}
	}

	return url.Values{
		"id": {strconv.FormatInt(x.ID, 10)},
	}
}

type userGetReq struct {
	UserID string
}

var _ urlValuer = userGetReq{}

func (x userGetReq) intoURLValues() url.Values {
	return url.Values{
		"userid": {x.UserID},
	}
}

// execUserGet 读取成员
func (c *App) execUserGet(req userGetReq) (userGetResp, error) {
	var resp userGetResp
	err := c.executeWXApiGet("/cgi-bin/user/get", req, &resp, true)
	if err != nil {
		return userGetResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return userGetResp{}, bizErr
	}

	return resp, nil
}

// execUserList 获取部门成员详情
func (c *App) execUserList(req userListReq) (userListResp, error) {
	var resp userListResp
	err := c.executeWXApiGet("/cgi-bin/user/list", req, &resp, true)
	if err != nil {
		return userListResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return userListResp{}, bizErr
	}

	return resp, nil
}

// execUserIDByMobile 手机号获取userid
func (c *App) execUserIDByMobile(req userIDByMobileReq) (userIDByMobileResp, error) {
	var resp userIDByMobileResp
	err := c.executeWXApiJSONPost("/cgi-bin/user/getuserid", req, &resp, true)
	if err != nil {
		return userIDByMobileResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return userIDByMobileResp{}, bizErr
	}

	return resp, nil
}

// execUserInfoGet 获取访问用户身份
func (c *App) execUserInfoGet(req userInfoGetReq) (userInfoGetResp, error) {
	var resp userInfoGetResp
	err := c.executeWXApiGet("/cgi-bin/user/getuserinfo", req, &resp, true)
	if err != nil {
		return userInfoGetResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return userInfoGetResp{}, bizErr
	}

	return resp, nil
}
