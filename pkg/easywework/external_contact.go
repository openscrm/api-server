package workwx

import (
	"time"
)

// ListExternalContact 获取客户列表
func (c *App) ListExternalContact(userID string) ([]string, error) {
	resp, err := c.execExternalContactList(externalContactListReq{
		UserID: userID,
	})
	if err != nil {
		return nil, err
	}
	return resp.ExternalUserID, nil
}

// GetExternalContact 获取客户详情
func (c *App) GetExternalContact(externalUserid string) (*ExternalContactInfo, error) {
	resp, err := c.execExternalContactGet(externalContactGetReq{
		ExternalUserID: externalUserid,
	})
	if err != nil {
		return nil, err
	}
	return &resp.ExternalContactInfo, nil
}

// BatchListExternalContact 批量获取客户详情
func (c *App) BatchListExternalContact(userID string, cursor string, limit int) (*BatchListExternalContactsResp, error) {
	resp, err := c.execExternalContactBatchList(externalContactBatchListReq{
		UserID: userID,
		Cursor: cursor,
		Limit:  limit,
	})
	if err != nil {
		return nil, err
	}
	return &BatchListExternalContactsResp{Result: resp.ExternalContactList, NextCursor: resp.NextCursor}, nil
}

// RemarkExternalContact 修改客户备注信息
func (c *App) RemarkExternalContact(req *ExternalContactRemark) error {
	_, err := c.execExternalContactRemark(externalContactRemarkReq{
		Remark: req,
	})
	return err
}

// ListExternalContactCorpTags 获取企业标签库
func (c *App) ListExternalContactCorpTags(tagIDs ...string) ([]ExternalContactCorpTagGroup, error) {
	resp, err := c.execExternalContactListCorpTags(externalContactListCorpTagsReq{
		TagIDs: tagIDs,
	})
	if err != nil {
		return nil, err
	}
	return resp.TagGroup, nil
}

// AddExternalContactCorpTag 添加企业客户标签
func (c *App) AddExternalContactCorpTag(req ExternalContactCorpTagGroup) (ExternalContactCorpTagGroup, error) {
	resp, err := c.execExternalContactAddCorpTag(externalContactAddCorpTagReq{
		ExternalContactCorpTagGroup: req,
	})
	if err != nil {
		return ExternalContactCorpTagGroup{}, err
	}
	return resp.TagGroup, nil
}

// EditExternalContactCorpTag 编辑企业客户标签
func (c *App) EditExternalContactCorpTag(id, name string, order uint32) error {
	_, err := c.execExternalContactEditCorpTag(externalContactEditCorpTagReq{
		ID:    id,
		Name:  name,
		Order: order,
	})
	return err
}

// DelExternalContactCorpTag 删除企业客户标签
func (c *App) DelExternalContactCorpTag(tagID, groupID []string) error {
	_, err := c.execExternalContactDelCorpTag(externalContactDelCorpTagReq{
		TagID:   tagID,
		GroupID: groupID,
	})
	return err
}

// MarkExternalContactTag 标记客户企业标签
func (c *App) MarkExternalContactTag(userID, externalUserID string, addTag, removeTag []string) error {
	_, err := c.execExternalContactMarkTag(externalContactMarkTagReq{
		UserID:         userID,
		ExternalUserID: externalUserID,
		AddTag:         addTag,
		RemoveTag:      removeTag,
	})
	return err
}

// ExternalContactUnassigned 离职成员的客户
type ExternalContactUnassigned struct {
	// HandoverUserID 离职成员的userid
	HandoverUserID string
	// ExternalUserID 外部联系人userid
	ExternalUserID string
	// DemissionTime 成员离职时间
	DemissionTime time.Time
}

// ListUnassignedExternalContact 获取离职成员的客户列表
func (c *App) ListUnassignedExternalContact(pageID, pageSize uint32, cursor string) (*ExternalContactUnassignedList, error) {
	resp, err := c.execListUnassignedExternalContact(listUnassignedExternalContactReq{
		PageID:   pageID,
		PageSize: pageSize,
		Cursor:   cursor,
	})
	if err != nil {
		return nil, err
	}
	externalContactUnassignedList := resp.intoExternalContactUnassignedList()
	return &externalContactUnassignedList, nil
}

// TransferExternalContact 分配成员的客户
func (c *App) TransferExternalContact(externalUserID, handoverUserID, takeoverUserID, transferSuccessMsg string) error {
	_, err := c.execTransferExternalContact(transferExternalContactReq{
		ExternalUserID:     externalUserID,
		HandoverUserID:     handoverUserID,
		TakeoverUserID:     takeoverUserID,
		TransferSuccessMsg: transferSuccessMsg,
	})
	return err
}

// ExternalContactTransferResult 客户接替结果
type ExternalContactTransferResult struct {
	// Status 接替状态， 1-接替完毕 2-等待接替 3-客户拒绝 4-接替成员客户达到上限 5-无接替记录
	Status ExternalContactTransferStatus
	// TakeoverTime 接替客户的时间，如果是等待接替状态，则为未来的自动接替时间
	TakeoverTime time.Time
}

// GetTransferExternalContactResult 查询客户接替结果
func (c *App) GetTransferExternalContactResult(externalUserID, handoverUserID, takeoverUserID string) (*ExternalContactTransferResult, error) {
	resp, err := c.execGetTransferExternalContactResult(getTransferExternalContactResultReq{
		ExternalUserID: externalUserID,
		HandoverUserID: handoverUserID,
		TakeoverUserID: takeoverUserID,
	})
	if err != nil {
		return nil, err
	}
	externalContactTransferResult := resp.intoExternalContactTransferResult()
	return &externalContactTransferResult, nil
}

// TransferGroupChatExternalContact 离职成员的群再分配
func (c *App) TransferGroupChatExternalContact(chatIDList []string, newOwner string) ([]ExternalContactGroupChatTransferFailed, error) {
	resp, err := c.execTransferGroupChatExternalContact(transferGroupChatExternalContactReq{
		ChatIDList: chatIDList,
		NewOwner:   newOwner,
	})
	if err != nil {
		return nil, err
	}
	return resp.FailedChatList, nil
}
