package workwx

// AddContactWay 配置客户联系「联系我」方式
// 文档：https://work.weixin.qq.com/api/doc/90000/90135/92572#配置客户联系「联系我」方式
func (c *App) AddContactWay(req AddContactWay) (configID string, err error) {
	var resp addContactWayResp
	resp, err = c.execAddContactWay(req)
	if err != nil {
		return configID, err
	}
	configID = resp.ConfigID
	return
}

// GetContactWay 获取企业已配置的「联系我」方式
// 文档：https://work.weixin.qq.com/api/doc/90000/90135/92572#获取企业已配置的「联系我」方式
func (c *App) GetContactWay(configID string) (contactWay ContactWay, err error) {
	var resp getContactWayResp
	resp, err = c.execGetContactWay(getContactWayReq{ConfigID: configID})
	if err != nil {
		return
	}
	contactWay = resp.ContactWay
	return
}

// UpdateContactWay 更新企业已配置的「联系我」方式
// 文档：https://work.weixin.qq.com/api/doc/90000/90135/92572#更新企业已配置的「联系我」方式
func (c *App) UpdateContactWay(req UpdateContactWay) (ok bool, err error) {
	var resp updateContactWayResp
	resp, err = c.execUpdateContactWay(req)
	if err != nil {
		return false, err
	}
	ok = resp.IsOK()
	return ok, err
}

// DelContactWay 删除企业已配置的「联系我」方式
// 文档：https://work.weixin.qq.com/api/doc/90000/90135/92572#删除企业已配置的「联系我」方式
func (c *App) DelContactWay(configID string) (ok bool, err error) {
	var resp delContactWayResp
	resp, err = c.execDelContactWay(delContactWayReq{
		ConfigID: configID,
	})
	if err != nil {
		return false, err
	}
	ok = resp.IsOK()
	return ok, err
}

// CloseTempChat 结束临时会话
// 文档：https://work.weixin.qq.com/api/doc/90000/90135/92572#结束临时会话
func (c *App) CloseTempChat(externalUserid string, userid string) (ok bool, err error) {
	var resp closeTempChatResp
	resp, err = c.execCloseTempChat(closeTempChatReq{
		ExternalUserid: externalUserid,
		Userid:         userid,
	})
	if err != nil {
		return false, err
	}
	ok = resp.IsOK()
	return ok, err
}
