package workwx

// ListGroupChat 获取客户群列表
// 文档：https://work.weixin.qq.com/api/doc/90000/90135/92120#获取客户群列表
func (c *App) ListGroupChat(req ListGroupChatReq) (ListGroupChatResp, error) {
	var resp ListGroupChatResp
	resp, err := c.execListGroupChat(req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// GetGroupChat 获取客户群详情
// 文档：https://work.weixin.qq.com/api/doc/90000/90135/92122#获取客户群详情
func (c *App) GetGroupChat(req GetGroupChatReq) (GetGroupChatResp, error) {
	resp, err := c.execGetGroupChat(req)
	if err != nil {
		return GetGroupChatResp{}, err
	}
	return resp, err
}
