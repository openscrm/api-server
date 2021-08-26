package workwx

// CreateTag 创建标签
// 文档：https://open.work.weixin.qq.com/api/doc/90000/90135/90210#创建标签
func (c *App) CreateTag(req Tag) (tagID int, err error) {
	var resp createTagResp
	resp, err = c.execCreateTag(req)
	if err != nil {
		return
	}

	tagID = resp.TagID
	return
}

// UpdateTag 更新标签名字
// 文档：https://open.work.weixin.qq.com/api/doc/90000/90135/90211#更新标签名字
func (c *App) UpdateTag(req Tag) (ok bool, err error) {
	var resp updateTagResp
	resp, err = c.execUpdateTag(req)
	if err != nil {
		return
	}
	ok = resp.IsOK()
	return
}

// ListTag 获取标签列表
// 文档：https://open.work.weixin.qq.com/api/doc/90000/90135/90216#获取标签列表
func (c *App) ListTag() (tags []Tag, err error) {
	var resp listTagResp
	resp, err = c.execListTag()
	if err != nil {
		return
	}

	tags = resp.TagList
	return
}

// DeleteTag 删除标签
// 文档：https://open.work.weixin.qq.com/api/doc/90000/90135/90212#删除标签
func (c *App) DeleteTag(tagID int) (ok bool, err error) {
	var resp deleteTagResp
	resp, err = c.execDeleteTag(deleteTagReq{TagID: tagID})
	if err != nil {
		return false, err
	}
	ok = resp.IsOK()
	return
}

// GetTagDetail 获取标签成员
// 文档：https://open.work.weixin.qq.com/api/doc/90000/90135/90213#获取标签成员
func (c *App) GetTagDetail(tagID int) (tagDetail TagDetail, err error) {
	var resp getTagResp
	resp, err = c.execGetTag(getTagReq{
		tagID,
	})
	if err != nil {
		return TagDetail{}, err
	}

	tagDetail = resp.TagDetail
	return
}

// AddTagUsers 增加标签成员
// 文档：https://open.work.weixin.qq.com/api/doc/90000/90135/90214#增加标签成员
func (c *App) AddTagUsers(req AddTagUsersReq) (ok bool, err error) {
	var resp addTagUsersResp
	resp, err = c.execAddTagUsers(req)
	if err != nil {
		return false, err
	}
	ok = resp.IsOK()
	return
}

// DelTagUsers 删除标签成员
// 文档：https://open.work.weixin.qq.com/api/doc/90000/90135/90215#删除标签成员
func (c *App) DelTagUsers(req DelTagUsersReq) (ok bool, err error) {
	var resp delTagUsersResp
	resp, err = c.execDelTagUsers(req)
	if err != nil {
		return false, err
	}
	ok = resp.IsOK()
	return
}
