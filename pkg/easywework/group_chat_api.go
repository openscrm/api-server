package workwx

import (
	"encoding/json"
)

// ListGroupChatReq
type ListGroupChatReq struct {
	StatusFilter int `json:"status_filter"`
	OwnerFilter  struct {
		UseridList []string `json:"userid_list"`
	} `json:"owner_filter"`
	Cursor string `json:"cursor"`
	Limit  int    `json:"limit"`
}

// ListGroupChatReq 获取客户群列表请求
// 文档：https://work.weixin.qq.com/api/doc/90000/90135/92120#获取客户群列表
var _ bodyer = ListGroupChatReq{}

func (x ListGroupChatReq) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// ListGroupChatResp 获取客户群列表响应
// 文档：https://work.weixin.qq.com/api/doc/90000/90135/92120#获取客户群列表
type ListGroupChatResp struct {
	CommonResp
	GroupChatList []struct {
		ChatID string `json:"chat_id"`
		Status int    `json:"status"`
	} `json:"group_chat_list"`
	NextCursor string `json:"next_cursor"`
}

var _ bodyer = ListGroupChatResp{}

func (x ListGroupChatResp) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// execListGroupChat 获取客户群列表
// 文档：https://work.weixin.qq.com/api/doc/90000/90135/92120#获取客户群列表
func (c *App) execListGroupChat(req ListGroupChatReq) (ListGroupChatResp, error) {
	var resp ListGroupChatResp
	err := c.executeWXApiJSONPost("/cgi-bin/externalcontact/groupchat/list", req, &resp, true)
	if err != nil {
		return ListGroupChatResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return ListGroupChatResp{}, bizErr
	}

	return resp, nil
}

// execGetGroupChat 获取客户群详情
// 文档：https://work.weixin.qq.com/api/doc/90000/90135/92122#获取客户群详情
func (c *App) execGetGroupChat(req GetGroupChatReq) (GetGroupChatResp, error) {
	var resp GetGroupChatResp
	err := c.executeWXApiJSONPost("/cgi-bin/externalcontact/groupchat/get", req, &resp, true)
	if err != nil {
		return GetGroupChatResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return GetGroupChatResp{}, bizErr
	}

	return resp, nil
}
