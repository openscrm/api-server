package workwx

// SendWelcomeMsg 发送新客户欢迎语
// 文档：https://open.work.weixin.qq.com/api/doc/90000/90135/92137#发送新客户欢迎语
func (c *App) SendWelcomeMsg(req SendWelcomeMsgReq) (ok bool, err error) {
	var resp sendWelcomeMsgResp
	resp, err = c.execSendWelcomeMsg(req)
	if err != nil {
		return false, err
	}
	ok = resp.IsOK()
	return
}
