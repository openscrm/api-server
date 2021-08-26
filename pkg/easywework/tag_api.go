package workwx

import (
	"encoding/json"
	"fmt"
	"net/url"
)

var _ bodyer = Tag{}

func (x Tag) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// createTagResp 创建标签
// 文档：https://open.work.weixin.qq.com/api/doc/90000/90135/90210#创建标签
type createTagResp struct {
	CommonResp
	TagID int `json:"tagid"`
}

var _ bodyer = createTagResp{}

func (x createTagResp) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// execCreateTag 创建标签
// 文档：https://open.work.weixin.qq.com/api/doc/90000/90135/90210#创建标签
func (c *App) execCreateTag(req Tag) (createTagResp, error) {
	var resp createTagResp
	err := c.executeWXApiJSONPost("/cgi-bin/tag/create", req, &resp, true)
	if err != nil {
		return createTagResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return createTagResp{}, bizErr
	}

	return resp, nil
}

// updateTagResp 更新标签名字
// 文档：https://open.work.weixin.qq.com/api/doc/90000/90135/90211#更新标签名字
type updateTagResp struct {
	CommonResp
}

var _ bodyer = updateTagResp{}

func (x updateTagResp) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// execUpdateTag 更新标签名字
// 文档：https://open.work.weixin.qq.com/api/doc/90000/90135/90211#更新标签名字
func (c *App) execUpdateTag(req Tag) (updateTagResp, error) {
	var resp updateTagResp
	err := c.executeWXApiJSONPost("/cgi-bin/tag/update", req, &resp, true)
	if err != nil {
		return updateTagResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return updateTagResp{}, bizErr
	}

	return resp, nil
}

// listTagResp 获取标签列表
// 文档：https://open.work.weixin.qq.com/api/doc/90000/90135/90216#获取标签列表
type listTagResp struct {
	CommonResp
	TagList []Tag `json:"taglist"` //标签列表
}

var _ bodyer = listTagResp{}

func (x listTagResp) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// execListTag 获取标签列表
// 文档：https://open.work.weixin.qq.com/api/doc/90000/90135/90216#获取标签列表
func (c *App) execListTag() (listTagResp, error) {
	var resp listTagResp
	err := c.executeWXApiGet("/cgi-bin/tag/list", nil, &resp, true)
	if err != nil {
		return listTagResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return listTagResp{}, bizErr
	}

	return resp, nil
}

// deleteTagReq 获取群聊会话请求
type deleteTagReq struct {
	TagID int
}

var _ urlValuer = deleteTagReq{}

func (x deleteTagReq) intoURLValues() url.Values {
	return url.Values{
		"tagid": {fmt.Sprintf("%d", x.TagID)},
	}
}

// deleteTagResp 删除标签
// 文档：https://open.work.weixin.qq.com/api/doc/90000/90135/90212#删除标签
type deleteTagResp struct {
	CommonResp
}

var _ bodyer = deleteTagResp{}

func (x deleteTagResp) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// execDeleteTag 删除标签
// 文档：https://open.work.weixin.qq.com/api/doc/90000/90135/90212#删除标签
func (c *App) execDeleteTag(req deleteTagReq) (deleteTagResp, error) {
	resp := deleteTagResp{}
	err := c.executeWXApiGet("/cgi-bin/tag/delete", req, &resp, true)
	if err != nil {
		return deleteTagResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return deleteTagResp{}, bizErr
	}

	return resp, nil
}

// getTagReq 获取标签成员
// 文档：https://open.work.weixin.qq.com/api/doc/90000/90135/90213#获取标签成员
type getTagReq struct {
	TagID int
}

var _ urlValuer = getTagReq{}

func (x getTagReq) intoURLValues() url.Values {
	return url.Values{
		"tagid": {fmt.Sprintf("%d", x.TagID)},
	}
}

// getTagResp 获取标签成员
// 文档：https://open.work.weixin.qq.com/api/doc/90000/90135/90213#获取标签成员
type getTagResp struct {
	CommonResp
	TagDetail
}

var _ bodyer = getTagResp{}

func (x getTagResp) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// execGetTag 获取标签成员
// 文档：https://open.work.weixin.qq.com/api/doc/90000/90135/90213#获取标签成员
func (c *App) execGetTag(req getTagReq) (getTagResp, error) {
	var resp getTagResp
	err := c.executeWXApiGet("/cgi-bin/tag/get", req, &resp, true)
	if err != nil {
		return getTagResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return getTagResp{}, bizErr
	}

	return resp, nil
}

// AddTagUsersReq 增加标签成员
// 文档：https://open.work.weixin.qq.com/api/doc/90000/90135/90214#增加标签成员

var _ bodyer = AddTagUsersReq{}

func (x AddTagUsersReq) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// addTagUsersResp 增加标签成员
// 文档：https://open.work.weixin.qq.com/api/doc/90000/90135/90214#增加标签成员
type addTagUsersResp struct {
	CommonResp
}

var _ bodyer = addTagUsersResp{}

func (x addTagUsersResp) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// execAddTagUsers 增加标签成员
// 文档：https://open.work.weixin.qq.com/api/doc/90000/90135/90214#增加标签成员
func (c *App) execAddTagUsers(req AddTagUsersReq) (addTagUsersResp, error) {
	var resp addTagUsersResp
	err := c.executeWXApiJSONPost("/cgi-bin/tag/addtagusers", req, &resp, true)
	if err != nil {
		return addTagUsersResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return addTagUsersResp{}, bizErr
	}

	return resp, nil
}

// 删除标签成员请求
// 文档：https://open.work.weixin.qq.com/api/doc/90000/90135/90215#删除标签成员
var _ bodyer = DelTagUsersReq{}

func (x DelTagUsersReq) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// delTagUsersResp 删除标签成员响应
// 文档：https://open.work.weixin.qq.com/api/doc/90000/90135/90215#删除标签成员
type delTagUsersResp struct {
	CommonResp
}

var _ bodyer = delTagUsersResp{}

func (x delTagUsersResp) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// execDelTagUsers 删除标签成员
// 文档：https://open.work.weixin.qq.com/api/doc/90000/90135/90215#删除标签成员
func (c *App) execDelTagUsers(req DelTagUsersReq) (delTagUsersResp, error) {
	var resp delTagUsersResp
	err := c.executeWXApiJSONPost("/cgi-bin/tag/deltagusers", req, &resp, true)
	if err != nil {
		return delTagUsersResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return delTagUsersResp{}, bizErr
	}

	return resp, nil
}
