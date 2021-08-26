package workwx

import (
	"encoding/json"
)

type GetGroupChatReq struct {
	ChatId string `json:"chat_id"`
}

// GetGroupChatReq 获取客户群详情请求
// 文档：https://work.weixin.qq.com/api/doc/90000/90135/92122#获取客户群详情
var _ bodyer = GetGroupChatReq{}

func (x GetGroupChatReq) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetGroupChatResp 获取客户群详情响应
// 文档：https://work.weixin.qq.com/api/doc/90000/90135/92122#获取客户群详情
type GetGroupChatResp struct {
	CommonResp
	GroupChat `json:"group_chat"`
}

type GroupChat struct {
	AdminList []struct {
		Userid string `json:"userid"`
	} `json:"admin_list"`
	ChatID     string `json:"chat_id"`
	CreateTime int    `json:"create_time"`
	MemberList []struct {
		Invitor struct {
			Userid string `json:"userid"`
		} `json:"invitor"`
		JoinScene int    `json:"join_scene"`
		JoinTime  int    `json:"join_time"`
		Type      int    `json:"type"`
		Unionid   string `json:"unionid"`
		Userid    string `json:"userid"`
	} `json:"member_list"`
	Name   string `json:"name"`
	Notice string `json:"notice"`
	Owner  string `json:"owner"`
}

var _ bodyer = GetGroupChatResp{}

func (x GetGroupChatResp) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}
