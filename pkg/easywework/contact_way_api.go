package workwx

import (
	"encoding/json"
)

var _ bodyer = AddContactWay{}

func (x AddContactWay) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// addContactWayResp 配置客户联系「联系我」方式
// 文档：https://work.weixin.qq.com/api/doc/90000/90135/92572#配置客户联系「联系我」方式
type addContactWayResp struct {
	// ConfigID 新增联系方式的配置id
	ConfigID string `json:"config_id"`
	CommonResp
	// QrCode 联系我二维码链接，仅在scene为2时返回
	QrCode string `json:"qr_code"`
}

var _ bodyer = addContactWayResp{}

func (x addContactWayResp) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// execAddContactWay 配置客户联系「联系我」方式
// 文档：https://work.weixin.qq.com/api/doc/90000/90135/92572#配置客户联系「联系我」方式
func (c *App) execAddContactWay(req AddContactWay) (addContactWayResp, error) {
	var resp addContactWayResp
	err := c.executeWXApiJSONPost("/cgi-bin/externalcontact/add_contact_way", req, &resp, true)
	if err != nil {
		return addContactWayResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return addContactWayResp{}, bizErr
	}

	return resp, nil
}

// getContactWayReq 获取企业已配置的「联系我」方式
// 文档：https://work.weixin.qq.com/api/doc/90000/90135/92572#获取企业已配置的「联系我」方式
type getContactWayReq struct {
	// ConfigID 联系方式的配置id，必填
	ConfigID string `json:"config_id"`
}

var _ bodyer = getContactWayReq{}

func (x getContactWayReq) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// getContactWayResp 获取企业已配置的「联系我」方式
// 文档：https://work.weixin.qq.com/api/doc/90000/90135/92572#获取企业已配置的「联系我」方式
type getContactWayResp struct {
	ContactWay ContactWay `json:"contact_way"`
	CommonResp
}

var _ bodyer = getContactWayResp{}

func (x getContactWayResp) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// execGetContactWay 获取企业已配置的「联系我」方式
// 文档：https://work.weixin.qq.com/api/doc/90000/90135/92572#获取企业已配置的「联系我」方式
func (c *App) execGetContactWay(req getContactWayReq) (getContactWayResp, error) {
	var resp getContactWayResp
	err := c.executeWXApiJSONPost("/cgi-bin/externalcontact/get_contact_way", req, &resp, true)
	if err != nil {
		return getContactWayResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return getContactWayResp{}, bizErr
	}

	return resp, nil
}

var _ bodyer = UpdateContactWay{}

func (x UpdateContactWay) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// updateContactWayResp 更新企业已配置的「联系我」方式
// 文档：https://work.weixin.qq.com/api/doc/90000/90135/92572#更新企业已配置的「联系我」方式
type updateContactWayResp struct {
	CommonResp
}

var _ bodyer = updateContactWayResp{}

func (x updateContactWayResp) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// execUpdateContactWay 更新企业已配置的「联系我」方式
// 文档：https://work.weixin.qq.com/api/doc/90000/90135/92572#更新企业已配置的「联系我」方式
func (c *App) execUpdateContactWay(req UpdateContactWay) (updateContactWayResp, error) {
	var resp updateContactWayResp
	err := c.executeWXApiJSONPost("/cgi-bin/externalcontact/update_contact_way", req, &resp, true)
	if err != nil {
		return updateContactWayResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return updateContactWayResp{}, bizErr
	}

	return resp, nil
}

// delContactWayReq 删除企业已配置的「联系我」方式
// 文档：https://work.weixin.qq.com/api/doc/90000/90135/92572#删除企业已配置的「联系我」方式
type delContactWayReq struct {
	// ConfigID 企业联系方式的配置id，必填
	ConfigID string `json:"config_id"`
}

var _ bodyer = delContactWayReq{}

func (x delContactWayReq) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// delContactWayResp 删除企业已配置的「联系我」方式
// 文档：https://work.weixin.qq.com/api/doc/90000/90135/92572#删除企业已配置的「联系我」方式
type delContactWayResp struct {
	CommonResp
}

var _ bodyer = delContactWayResp{}

func (x delContactWayResp) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// execDelContactWay 删除企业已配置的「联系我」方式
// 文档：https://work.weixin.qq.com/api/doc/90000/90135/92572#删除企业已配置的「联系我」方式
func (c *App) execDelContactWay(req delContactWayReq) (delContactWayResp, error) {
	var resp delContactWayResp
	err := c.executeWXApiJSONPost("/cgi-bin/externalcontact/del_contact_way", req, &resp, true)
	if err != nil {
		return delContactWayResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return delContactWayResp{}, bizErr
	}

	return resp, nil
}

// closeTempChatReq 结束临时会话
// 文档：https://work.weixin.qq.com/api/doc/90000/90135/92572#结束临时会话
type closeTempChatReq struct {
	// ExternalUserid 客户的外部联系人userid，必填
	ExternalUserid string `json:"external_userid"`
	// Userid 企业成员的userid，必填
	Userid string `json:"userid"`
}

var _ bodyer = closeTempChatReq{}

func (x closeTempChatReq) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// closeTempChatResp 结束临时会话
// 文档：https://work.weixin.qq.com/api/doc/90000/90135/92572#结束临时会话
type closeTempChatResp struct {
	CommonResp
}

var _ bodyer = closeTempChatResp{}

func (x closeTempChatResp) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// execCloseTempChat 结束临时会话
// 文档：https://work.weixin.qq.com/api/doc/90000/90135/92572#结束临时会话
func (c *App) execCloseTempChat(req closeTempChatReq) (closeTempChatResp, error) {
	var resp closeTempChatResp
	err := c.executeWXApiJSONPost("/cgi-bin/externalcontact/close_temp_chat", req, &resp, true)
	if err != nil {
		return closeTempChatResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return closeTempChatResp{}, bizErr
	}

	return resp, nil
}
