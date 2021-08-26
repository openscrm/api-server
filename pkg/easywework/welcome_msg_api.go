package workwx

import "encoding/json"

type SendWelcomeMsgReq struct {
	//附件，最多可添加9个附件
	Attachments []Attachments `json:"attachments,omitempty"`
	Text        Text          `json:"text"`
	// WelcomeCode 通过 添加外部联系人事件 推送给企业的发送欢迎语的凭证，有效期为 20秒，必填
	WelcomeCode string `json:"welcome_code"`
}

var _ bodyer = SendWelcomeMsgReq{}

func (x SendWelcomeMsgReq) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// sendWelcomeMsgResp 发送新客户欢迎语响应
// 文档：https://open.work.weixin.qq.com/api/doc/90000/90135/92137#发送新客户欢迎语
type sendWelcomeMsgResp struct {
	CommonResp
}

var _ bodyer = sendWelcomeMsgResp{}

func (x sendWelcomeMsgResp) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// execSendWelcomeMsg 发送新客户欢迎语
// 文档：https://open.work.weixin.qq.com/api/doc/90000/90135/92137#发送新客户欢迎语
func (c *App) execSendWelcomeMsg(req SendWelcomeMsgReq) (sendWelcomeMsgResp, error) {
	var resp sendWelcomeMsgResp
	err := c.executeWXApiJSONPost("/cgi-bin/externalcontact/send_welcome_msg", req, &resp, true)
	if err != nil {
		return sendWelcomeMsgResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return sendWelcomeMsgResp{}, bizErr
	}

	return resp, nil
}
