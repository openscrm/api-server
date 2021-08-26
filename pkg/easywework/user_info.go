package workwx

//企业成员（员工） 相关接口

// GetUser 读取成员
func (c *App) GetUser(userid string) (*UserInfo, error) {
	resp, err := c.execUserGet(userGetReq{
		UserID: userid,
	})
	if err != nil {
		return nil, err
	}

	obj := resp.intoUserInfo()
	return &obj, nil
}

// ListUsersByDeptID 获取部门成员详情
func (c *App) ListUsersByDeptID(deptID int64, fetchChild bool) ([]*UserInfo, error) {
	resp, err := c.execUserList(userListReq{
		DeptID:     deptID,
		FetchChild: fetchChild,
	})
	if err != nil {
		return nil, err
	}
	users := make([]*UserInfo, len(resp.Users))
	for index, user := range resp.Users {
		userInfo := user.intoUserInfo()
		users[index] = &userInfo
	}
	return users, nil
}

// GetUserIDByMobile 通过手机号获取 userid
func (c *App) GetUserIDByMobile(mobile string) (string, error) {
	resp, err := c.execUserIDByMobile(userIDByMobileReq{
		Mobile: mobile,
	})
	if err != nil {
		return "", err
	}
	return resp.UserID, nil
}

// GetUserInfoByCode 获取访问用户身份，根据code获取成员信息
func (c *App) GetUserInfoByCode(code string) (*UserIdentityInfo, error) {
	resp, err := c.execUserInfoGet(userInfoGetReq{
		Code: code,
	})
	if err != nil {
		return nil, err
	}
	return &resp.UserIdentityInfo, nil
}

// UpdateUser 更新成员
// 文档：https://work.weixin.qq.com/api/doc/90000/90135/90197#更新成员
func (c *App) UpdateUser(req UpdateUserReq) (ok bool, err error) {
	var resp updateUserResp
	resp, err = c.execUpdateUser(req)
	if err != nil {
		return false, err
	}
	ok = resp.IsOK()
	return
}
