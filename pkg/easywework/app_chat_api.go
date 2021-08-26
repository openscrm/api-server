package workwx

import (
	"encoding/json"
	"net/url"
)

// appChatGetResp 获取群聊会话响应
type appChatGetResp struct {
	CommonResp

	ChatInfo *ChatInfo `json:"chat_info"`
}

// appChatGetReq 获取群聊会话请求
type appChatGetReq struct {
	ChatID string
}

var _ urlValuer = appChatGetReq{}

func (x appChatGetReq) intoURLValues() url.Values {
	return url.Values{
		"chatid": {x.ChatID},
	}
}

// appChatCreateReq 创建群聊会话请求
type appChatCreateReq struct {
	ChatInfo *ChatInfo
}

var _ bodyer = appChatCreateReq{}

func (x appChatCreateReq) intoBody() ([]byte, error) {
	result, err := json.Marshal(x.ChatInfo)
	if err != nil {
		// should never happen unless OOM or similar bad things
		// TODO: error_chain
		return nil, err
	}

	return result, nil
}

// appChatCreateResp 创建群聊会话响应
type appChatCreateResp struct {
	CommonResp
	ChatID string `json:"chatid"`
}

// execAppChatCreate 创建群聊会话
func (c *App) execAppChatCreate(req appChatCreateReq) (appChatCreateResp, error) {
	var resp appChatCreateResp
	err := c.executeWXApiJSONPost("/cgi-bin/appchat/create", req, &resp, true)
	if err != nil {
		return appChatCreateResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return appChatCreateResp{}, bizErr
	}

	return resp, nil
}

// execAppChatGet 获取群聊会话
func (c *App) execAppChatGet(req appChatGetReq) (appChatGetResp, error) {
	var resp appChatGetResp
	err := c.executeWXApiGet("/cgi-bin/appchat/get", req, &resp, true)
	if err != nil {
		return appChatGetResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return appChatGetResp{}, bizErr
	}

	return resp, nil
}
