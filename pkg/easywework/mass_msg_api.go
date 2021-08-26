package workwx

import (
	"encoding/json"
)

var _ bodyer = AddMsgTemplateReq{}

func (x AddMsgTemplateReq) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

var _ bodyer = addMsgTemplateResp{}

func (x addMsgTemplateResp) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// execAddMsgTemplate 创建企业群发
// 文档：https://open.work.weixin.qq.com/api/doc/90000/90135/92135#创建企业群发
func (c *App) execAddMsgTemplate(req AddMsgTemplateReq) (addMsgTemplateResp, error) {
	var resp addMsgTemplateResp
	err := c.executeWXApiJSONPost("/cgi-bin/externalcontact/add_msg_template", req, &resp, true)
	if err != nil {
		return addMsgTemplateResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return addMsgTemplateResp{}, bizErr
	}

	return resp, nil
}

var _ bodyer = GetGroupMsgSendResultExternalContactReq{}

func (x GetGroupMsgSendResultExternalContactReq) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

var _ bodyer = GetGroupMsgSendResultExternalContactResp{}

func (x GetGroupMsgSendResultExternalContactResp) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// execGetGroupMsgSendResultExternalContact 获取企业群发成员执行结果
// 文档：https://work.weixin.qq.com/api/doc/90000/90135/93338#获取企业群发成员执行结果
func (c *App) execGetGroupMsgSendResultExternalContact(req GetGroupMsgSendResultExternalContactReq) (GetGroupMsgSendResultExternalContactResp, error) {
	var resp GetGroupMsgSendResultExternalContactResp
	err := c.executeWXApiJSONPost("/cgi-bin/externalcontact/get_groupmsg_send_result", req, &resp, true)
	if err != nil {
		return GetGroupMsgSendResultExternalContactResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return GetGroupMsgSendResultExternalContactResp{}, bizErr
	}

	return resp, nil
}

var _ bodyer = reqGetGroupmsgTaskExternalcontact{}

func (x reqGetGroupmsgTaskExternalcontact) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

var _ bodyer = getGroupMsgTaskExternalContactResp{}

func (x getGroupMsgTaskExternalContactResp) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// execGetGroupMsgTaskExternalContact 获取群发成员发送任务列表
// 文档：https://work.weixin.qq.com/api/doc/90000/90135/93338#获取群发成员发送任务列表
func (c *App) execGetGroupMsgTaskExternalContact(req reqGetGroupmsgTaskExternalcontact) (getGroupMsgTaskExternalContactResp, error) {
	var resp getGroupMsgTaskExternalContactResp
	err := c.executeWXApiJSONPost("/cgi-bin/externalcontact/get_groupmsg_task", req, &resp, true)
	if err != nil {
		return getGroupMsgTaskExternalContactResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return getGroupMsgTaskExternalContactResp{}, bizErr
	}

	return resp, nil
}

var _ bodyer = getGroupMsgListV2ExternalContactReq{}

func (x getGroupMsgListV2ExternalContactReq) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

var _ bodyer = getGroupMsgListV2ExternalContactResp{}

func (x getGroupMsgListV2ExternalContactResp) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// execGetGroupmsgListV2Externalcontact 获取群发记录列表
// 文档：https://work.weixin.qq.com/api/doc/90000/90135/93338#获取群发记录列表
func (c *App) execGetGroupmsgListV2Externalcontact(req getGroupMsgListV2ExternalContactReq) (getGroupMsgListV2ExternalContactResp, error) {
	var resp getGroupMsgListV2ExternalContactResp
	err := c.executeWXApiJSONPost("/cgi-bin/externalcontact/get_groupmsg_list_v2", req, &resp, true)
	if err != nil {
		return getGroupMsgListV2ExternalContactResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return getGroupMsgListV2ExternalContactResp{}, bizErr
	}

	return resp, nil
}

// execGetGroupmsgTaskExternalcontact 获取群发成员发送任务列表
// 文档：https://work.weixin.qq.com/api/doc/90000/90135/93338#获取群发成员发送任务列表
func (c *App) execGetGroupmsgTaskExternalcontact(req reqGetGroupmsgTaskExternalcontact) (getGroupMsgTaskExternalContactResp, error) {
	var resp getGroupMsgTaskExternalContactResp
	err := c.executeWXApiJSONPost("/cgi-bin/externalcontact/get_groupmsg_task", req, &resp, true)
	if err != nil {
		return getGroupMsgTaskExternalContactResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return getGroupMsgTaskExternalContactResp{}, bizErr
	}

	return resp, nil
}
