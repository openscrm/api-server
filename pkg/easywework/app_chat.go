package workwx

// CreateAppChat 创建群聊会话
func (c *App) CreateAppChat(chatInfo *ChatInfo) (chatid string, err error) {
	resp, err := c.execAppChatCreate(appChatCreateReq{
		ChatInfo: chatInfo,
	})
	if err != nil {
		return "", err
	}
	return resp.ChatID, nil
}

// GetAppChat 获取群聊会话
func (c *App) GetAppChat(chatID string) (*ChatInfo, error) {
	resp, err := c.execAppChatGet(appChatGetReq{
		ChatID: chatID,
	})
	if err != nil {
		return nil, err
	}
	obj := resp.ChatInfo
	return obj, nil
}
