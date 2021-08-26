package workwx

import (
	"encoding/json"
	"errors"
	"strings"
)

// sendMessage 发送消息底层接口
//
// 收件人参数如果仅设置了 `ChatID` 字段，则为【发送消息到群聊会话】接口调用；
// 否则为单纯的【发送应用消息】接口调用。
func (c *App) sendMessage(
	recipient *Recipient,
	msgtype string,
	content map[string]interface{},
	isSafe bool,
) error {
	isApichatSendRequest := false
	if !recipient.isValidForMessageSend() {
		if !recipient.isValidForAppChatSend() {
			// TODO: better error
			return errors.New("recipient invalid for message sending")
		}

		// 发送给群聊
		isApichatSendRequest = true
	}

	req := MessageReq{
		ToUser:  recipient.UserIDs,
		ToParty: recipient.PartyIDs,
		ToTag:   recipient.TagIDs,
		ChatID:  recipient.ChatID,
		AgentID: c.AgentID,
		MsgType: msgtype,
		Content: content,
		IsSafe:  isSafe,
	}

	var resp messageSendResp
	var err error
	if isApichatSendRequest {
		resp, err = c.execAppChatSend(req)
	} else {
		resp, err = c.execMessageSend(req)
	}

	if err != nil {
		return err
	}
	_ = resp
	return nil
}

// execMessageSend 发送应用消息
func (c *App) execMessageSend(req MessageReq) (messageSendResp, error) {
	var resp messageSendResp
	err := c.executeWXApiJSONPost("/cgi-bin/message/send", req, &resp, true)
	if err != nil {
		return messageSendResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return messageSendResp{}, bizErr
	}

	return resp, nil
}

// execAppChatSend 应用推送消息
func (c *App) execAppChatSend(req MessageReq) (messageSendResp, error) {
	var resp messageSendResp
	err := c.executeWXApiJSONPost("/cgi-bin/appchat/send", req, &resp, true)
	if err != nil {
		return messageSendResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return messageSendResp{}, bizErr
	}

	return resp, nil
}

// MessageReq 消息发送请求
type MessageReq struct {
	ToUser  []string
	ToParty []string
	ToTag   []string
	ChatID  string
	AgentID int64
	MsgType string
	Content map[string]interface{}
	IsSafe  bool
}

var _ bodyer = MessageReq{}

func (x MessageReq) intoBody() ([]byte, error) {
	// fuck
	safeInt := 0
	if x.IsSafe {
		safeInt = 1
	}

	obj := map[string]interface{}{
		"msgtype": x.MsgType,
		"agentid": x.AgentID,
		"safe":    safeInt,
	}

	// msgtype polymorphism
	obj[x.MsgType] = x.Content

	// 复用这个结构体，因为是 package-private 的所以这么做没风险
	if x.ChatID != "" {
		obj["chatid"] = x.ChatID
	} else {
		obj["touser"] = strings.Join(x.ToUser, "|")
		obj["toparty"] = strings.Join(x.ToParty, "|")
		obj["totag"] = strings.Join(x.ToTag, "|")
	}

	result, err := json.Marshal(obj)
	if err != nil {
		// should never happen unless OOM or similar bad things
		// TODO: error_chain
		return nil, err
	}

	return result, nil
}

// messageSendResp 消息发送响应
type messageSendResp struct {
	CommonResp

	InvalidUsers   string `json:"invaliduser"`
	InvalidParties string `json:"invalidparty"`
	InvalidTags    string `json:"invalidtag"`
}
