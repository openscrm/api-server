package workwx

// MsgAuditAgreeStatus 会话中外部成员的同意状态
type MsgAuditAgreeStatus string

const (
	// MsgAuditAgreeStatusAgree 同意
	MsgAuditAgreeStatusAgree = "Agree"
	// MsgAuditAgreeStatusDisagree 不同意
	MsgAuditAgreeStatusDisagree = "Disagree"
	// MsgAuditAgreeStatusDefaultAgree 默认同意
	MsgAuditAgreeStatusDefaultAgree = "Default_Agree"
)

// CheckMsgAuditSingleAgree 获取会话同意情况（单聊）
func (c *App) CheckMsgAuditSingleAgree(infos []CheckMsgAuditSingleAgreeUserInfo) ([]CheckMsgAuditSingleAgreeInfo, error) {
	resp, err := c.execMsgAuditCheckSingleAgree(msgAuditCheckSingleAgreeReq{
		Infos: infos,
	})
	if err != nil {
		return nil, err
	}
	return resp.intoCheckSingleAgreeInfoList(), nil
}

// CheckMsgAuditRoomAgree 获取会话同意情况（群聊）
func (c *App) CheckMsgAuditRoomAgree(roomId string) ([]CheckMsgAuditRoomAgreeInfo, error) {
	resp, err := c.execMsgAuditCheckRoomAgree(msgAuditCheckRoomAgreeReq{
		RoomID: roomId,
	})
	if err != nil {
		return nil, err
	}
	return resp.intoCheckRoomAgreeInfoList(), nil
}

// MsgAuditEdition 会话内容存档版本
type MsgAuditEdition uint8

const (
	// MsgAuditEditionOffice 会话内容存档办公版
	MsgAuditEditionOffice MsgAuditEdition = 1
	// MsgAuditEditionService 会话内容存档服务版
	MsgAuditEditionService MsgAuditEdition = 2
	// MsgAuditEditionEnterprise 会话内容存档企业版
	MsgAuditEditionEnterprise MsgAuditEdition = 3
)

// ListMsgAuditPermitUser 获取会话内容存档开启成员列表
func (c *App) ListMsgAuditPermitUser(msgAuditEdition MsgAuditEdition) ([]string, error) {
	resp, err := c.execMsgAuditListPermitUser(msgAuditListPermitUserReq{
		MsgAuditEdition: msgAuditEdition,
	})
	if err != nil {
		return nil, err
	}
	return resp.IDs, nil
}

// GetMsgAuditGroupChat 获取会话内容存档内部群信息
func (c *App) GetMsgAuditGroupChat(roomID string) (*MsgAuditGroupChat, error) {
	resp, err := c.execMsgAuditGetGroupChat(msgAuditGetGroupChatReq{
		RoomID: roomID,
	})
	if err != nil {
		return nil, err
	}
	groupChat := resp.intoGroupChat()
	return &groupChat, nil
}
