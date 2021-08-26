package workwx

// execMsgAuditListPermitUser 获取会话内容存档开启成员列表
func (c *App) execMsgAuditListPermitUser(req msgAuditListPermitUserReq) (msgAuditListPermitUserResp, error) {
	var resp msgAuditListPermitUserResp
	err := c.executeWXApiJSONPost("/cgi-bin/msgaudit/get_permit_user_list", req, &resp, true)
	if err != nil {
		return msgAuditListPermitUserResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return msgAuditListPermitUserResp{}, bizErr
	}

	return resp, nil
}

// execMsgAuditCheckSingleAgree 获取会话同意情况（单聊）
func (c *App) execMsgAuditCheckSingleAgree(req msgAuditCheckSingleAgreeReq) (msgAuditCheckSingleAgreeResp, error) {
	var resp msgAuditCheckSingleAgreeResp
	err := c.executeWXApiJSONPost("/cgi-bin/msgaudit/check_single_agree", req, &resp, true)
	if err != nil {
		return msgAuditCheckSingleAgreeResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return msgAuditCheckSingleAgreeResp{}, bizErr
	}

	return resp, nil
}

// execMsgAuditCheckRoomAgree 获取会话同意情况（群聊）
func (c *App) execMsgAuditCheckRoomAgree(req msgAuditCheckRoomAgreeReq) (msgAuditCheckRoomAgreeResp, error) {
	var resp msgAuditCheckRoomAgreeResp
	err := c.executeWXApiJSONPost("/cgi-bin/msgaudit/check_room_agree", req, &resp, true)
	if err != nil {
		return msgAuditCheckRoomAgreeResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return msgAuditCheckRoomAgreeResp{}, bizErr
	}

	return resp, nil
}

// execMsgAuditGetGroupChat 获取会话内容存档内部群信息
func (c *App) execMsgAuditGetGroupChat(req msgAuditGetGroupChatReq) (msgAuditGetGroupChatResp, error) {
	var resp msgAuditGetGroupChatResp
	err := c.executeWXApiJSONPost("/cgi-bin/msgaudit/groupchat/get", req, &resp, true)
	if err != nil {
		return msgAuditGetGroupChatResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return msgAuditGetGroupChatResp{}, bizErr
	}

	return resp, nil
}
