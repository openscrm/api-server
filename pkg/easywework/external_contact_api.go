package workwx

import (
	"encoding/json"
	"net/url"
	"time"
)

// execExternalContactList 获取客户列表
func (c *App) execExternalContactList(req externalContactListReq) (externalContactListResp, error) {
	var resp externalContactListResp
	err := c.executeWXApiGet("/cgi-bin/externalcontact/list", req, &resp, true)
	if err != nil {
		return externalContactListResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return externalContactListResp{}, bizErr
	}

	return resp, nil
}

// execExternalContactGet 获取客户详情
func (c *App) execExternalContactGet(req externalContactGetReq) (externalContactGetResp, error) {
	var resp externalContactGetResp
	err := c.executeWXApiGet("/cgi-bin/externalcontact/get", req, &resp, true)
	if err != nil {
		return externalContactGetResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return externalContactGetResp{}, bizErr
	}

	return resp, nil
}

// execExternalContactBatchList 批量获取客户详情
func (c *App) execExternalContactBatchList(req externalContactBatchListReq) (externalContactBatchListResp, error) {
	var resp externalContactBatchListResp
	err := c.executeWXApiJSONPost("/cgi-bin/externalcontact/batch/get_by_user", req, &resp, true)
	if err != nil {
		return externalContactBatchListResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return externalContactBatchListResp{}, bizErr
	}

	return resp, nil
}

// execExternalContactRemark 修改客户备注信息
func (c *App) execExternalContactRemark(req externalContactRemarkReq) (externalContactRemarkResp, error) {
	var resp externalContactRemarkResp
	err := c.executeWXApiJSONPost("/cgi-bin/externalcontact/remark", req, &resp, true)
	if err != nil {
		return externalContactRemarkResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return externalContactRemarkResp{}, bizErr
	}

	return resp, nil
}

// execExternalContactListCorpTags 获取企业标签库
func (c *App) execExternalContactListCorpTags(req externalContactListCorpTagsReq) (externalContactListCorpTagsResp, error) {
	var resp externalContactListCorpTagsResp
	err := c.executeWXApiJSONPost("/cgi-bin/externalcontact/get_corp_tag_list", req, &resp, true)
	if err != nil {
		return externalContactListCorpTagsResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return externalContactListCorpTagsResp{}, bizErr
	}

	return resp, nil
}

// execExternalContactAddCorpTag 添加企业客户标签
func (c *App) execExternalContactAddCorpTag(req externalContactAddCorpTagReq) (externalContactAddCorpTagResp, error) {
	var resp externalContactAddCorpTagResp
	err := c.executeWXApiJSONPost("/cgi-bin/externalcontact/add_corp_tag", req, &resp, true)
	if err != nil {
		return externalContactAddCorpTagResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return externalContactAddCorpTagResp{}, bizErr
	}

	return resp, nil
}

// execExternalContactEditCorpTag 编辑企业客户标签
func (c *App) execExternalContactEditCorpTag(req externalContactEditCorpTagReq) (externalContactEditCorpTagResp, error) {
	var resp externalContactEditCorpTagResp
	err := c.executeWXApiJSONPost("/cgi-bin/externalcontact/edit_corp_tag", req, &resp, true)
	if err != nil {
		return externalContactEditCorpTagResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return externalContactEditCorpTagResp{}, bizErr
	}

	return resp, nil
}

// execExternalContactDelCorpTag 删除企业客户标签
func (c *App) execExternalContactDelCorpTag(req externalContactDelCorpTagReq) (externalContactDelCorpTagResp, error) {
	var resp externalContactDelCorpTagResp
	err := c.executeWXApiJSONPost("/cgi-bin/externalcontact/del_corp_tag", req, &resp, true)
	if err != nil {
		return externalContactDelCorpTagResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return externalContactDelCorpTagResp{}, bizErr
	}

	return resp, nil
}

// execExternalContactMarkTag 标记客户企业标签
func (c *App) execExternalContactMarkTag(req externalContactMarkTagReq) (externalContactMarkTagResp, error) {
	var resp externalContactMarkTagResp
	err := c.executeWXApiJSONPost("/cgi-bin/externalcontact/mark_tag", req, &resp, true)
	if err != nil {
		return externalContactMarkTagResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return externalContactMarkTagResp{}, bizErr
	}

	return resp, nil
}

// execListUnassignedExternalContact 获取离职成员的客户列表
func (c *App) execListUnassignedExternalContact(req listUnassignedExternalContactReq) (listUnassignedExternalContactResp, error) {
	var resp listUnassignedExternalContactResp
	err := c.executeWXApiJSONPost("/cgi-bin/externalcontact/get_unassigned_list", req, &resp, true)
	if err != nil {
		return listUnassignedExternalContactResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return listUnassignedExternalContactResp{}, bizErr
	}

	return resp, nil
}

// execTransferExternalContact 分配成员的客户
func (c *App) execTransferExternalContact(req transferExternalContactReq) (transferExternalContactResp, error) {
	var resp transferExternalContactResp
	err := c.executeWXApiJSONPost("/cgi-bin/externalcontact/transfer", req, &resp, true)
	if err != nil {
		return transferExternalContactResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return transferExternalContactResp{}, bizErr
	}

	return resp, nil
}

// execGetTransferExternalContactResult 查询客户接替结果
func (c *App) execGetTransferExternalContactResult(req getTransferExternalContactResultReq) (getTransferExternalContactResultResp, error) {
	var resp getTransferExternalContactResultResp
	err := c.executeWXApiJSONPost("/cgi-bin/externalcontact/get_transfer_result", req, &resp, true)
	if err != nil {
		return getTransferExternalContactResultResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return getTransferExternalContactResultResp{}, bizErr
	}

	return resp, nil
}

// execTransferGroupChatExternalContact 离职成员的群再分配
func (c *App) execTransferGroupChatExternalContact(req transferGroupChatExternalContactReq) (transferGroupChatExternalContactResp, error) {
	var resp transferGroupChatExternalContactResp
	err := c.executeWXApiJSONPost("/cgi-bin/externalcontact/groupchat/transfer", req, &resp, true)
	if err != nil {
		return transferGroupChatExternalContactResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return transferGroupChatExternalContactResp{}, bizErr
	}

	return resp, nil
}

type listUnassignedExternalContactReq struct {
	// PageID 分页查询，要查询页号，从0开始
	PageID uint32 `json:"page_id"`
	// PageSize 每次返回的最大记录数，默认为1000，最大值为1000
	PageSize uint32 `json:"page_size"`
	// Cursor 分页查询游标，字符串类型，适用于数据量较大的情况，如果使用该参数则无需填写page_id，该参数由上一次调用返回
	Cursor string `json:"cursor"`
}

var _ bodyer = listUnassignedExternalContactReq{}

func (x listUnassignedExternalContactReq) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type listUnassignedExternalContactResp struct {
	CommonResp
	Info []struct {
		HandoverUserid string `json:"handover_userid"`
		ExternalUserid string `json:"external_userid"`
		DemissionTime  int    `json:"dimission_time"`
	} `json:"info"`
	IsLast     bool   `json:"is_last"`
	NextCursor string `json:"next_cursor"`
}

func (x listUnassignedExternalContactResp) intoExternalContactUnassignedList() (resp ExternalContactUnassignedList) {
	list := make([]ExternalContactUnassigned, 0, len(x.Info))
	for _, info := range x.Info {
		list = append(list, ExternalContactUnassigned{
			HandoverUserID: info.HandoverUserid,
			ExternalUserID: info.ExternalUserid,
			DemissionTime:  time.Unix(int64(info.DemissionTime), 0),
		})
	}
	resp.Info = list
	resp.IsLast = x.IsLast
	resp.NextCursor = x.NextCursor
	return resp
}

type transferExternalContactReq struct {
	// ExternalUserID 外部联系人的userid，注意不是企业成员的帐号
	ExternalUserID string `json:"external_userid"`
	// HandoverUserID 原跟进成员的userid
	HandoverUserID string `json:"handover_userid"`
	// TakeoverUserID 接替成员的userid
	TakeoverUserID string `json:"takeover_userid"`
	// TransferSuccessMsg 转移成功后发给客户的消息，最多200个字符，不填则使用默认文案，目前只对在职成员分配客户的情况生效
	TransferSuccessMsg string `json:"transfer_success_msg"`
}

var _ bodyer = transferExternalContactReq{}

func (x transferExternalContactReq) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type transferExternalContactResp struct {
	CommonResp
}

type getTransferExternalContactResultReq struct {
	// ExternalUserID 外部联系人的userid，注意不是企业成员的帐号
	ExternalUserID string `json:"external_userid"`
	// HandoverUserID 原跟进成员的userid
	HandoverUserID string `json:"handover_userid"`
	// TakeoverUserID 接替成员的userid
	TakeoverUserID string `json:"takeover_userid"`
}

var _ bodyer = getTransferExternalContactResultReq{}

func (x getTransferExternalContactResultReq) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type getTransferExternalContactResultResp struct {
	CommonResp
	Status       uint8 `json:"status"`
	TakeoverTime int   `json:"takeover_time"`
}

func (x getTransferExternalContactResultResp) intoExternalContactTransferResult() ExternalContactTransferResult {
	return ExternalContactTransferResult{
		Status:       ExternalContactTransferStatus(x.Status),
		TakeoverTime: time.Unix(int64(x.TakeoverTime), 0),
	}
}

type transferGroupChatExternalContactReq struct {
	// ChatIDList 需要转群主的客户群ID列表。取值范围： 1 ~ 100
	ChatIDList []string `json:"chat_id_list"`
	// NewOwner 新群主ID
	NewOwner string `json:"new_owner"`
}

var _ bodyer = transferGroupChatExternalContactReq{}

func (x transferGroupChatExternalContactReq) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type transferGroupChatExternalContactResp struct {
	CommonResp
	FailedChatList []ExternalContactGroupChatTransferFailed `json:"failed_chat_list"`
}

// externalContactBatchListResp 批量获取客户详情
type externalContactBatchListResp struct {
	CommonResp
	NextCursor          string                     `json:"next_cursor"`
	ExternalContactList []ExternalContactBatchInfo `json:"external_contact_list"`
}

// externalContactRemarkReq 获取客户详情
type externalContactRemarkReq struct {
	Remark *ExternalContactRemark
}

var _ bodyer = externalContactRemarkReq{}

func (x externalContactRemarkReq) intoBody() ([]byte, error) {
	result, err := json.Marshal(x.Remark)
	if err != nil {
		// should never happen unless OOM or similar bad things
		// TODO: error_chain
		return nil, err
	}

	return result, nil
}

// externalContactRemarkResp 获取客户详情
type externalContactRemarkResp struct {
	CommonResp
}

// userInfoGetReq 获取访问用户身份
type userInfoGetReq struct {
	// 通过成员授权获取到的code，最大为512字节。每次成员授权带上的code将不一样，code只能使用一次，5分钟未被使用自动过期。
	Code string
}

// externalContactListCorpTagsReq 获取企业标签库
type externalContactListCorpTagsReq struct {
	// 要查询的标签id，如果不填则获取该企业的所有客户标签，目前暂不支持标签组id
	TagIDs []string `json:"tag_id"`
}

var _ bodyer = externalContactListCorpTagsReq{}

func (x externalContactListCorpTagsReq) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		// should never happen unless OOM or similar bad things
		// TODO: error_chain
		return nil, err
	}

	return result, nil
}

// externalContactListCorpTagsResp 获取企业标签库
type externalContactListCorpTagsResp struct {
	CommonResp
	// 标签组列表
	TagGroup []ExternalContactCorpTagGroup `json:"tag_group"`
}

// externalContactAddCorpTagReq 添加企业客户标签
type externalContactAddCorpTagReq struct {
	ExternalContactCorpTagGroup
}

var _ bodyer = externalContactAddCorpTagReq{}

func (x externalContactAddCorpTagReq) intoBody() ([]byte, error) {
	result, err := json.Marshal(x.ExternalContactCorpTagGroup)
	if err != nil {
		// should never happen unless OOM or similar bad things
		// TODO: error_chain
		return nil, err
	}

	return result, nil
}

// externalContactAddCorpTagResp 添加企业客户标签
type externalContactAddCorpTagResp struct {
	CommonResp
	// 标签组列表
	TagGroup ExternalContactCorpTagGroup `json:"tag_group"`
}

// externalContactEditCorpTagReq 编辑企业客户标签
type externalContactEditCorpTagReq struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Order uint32 `json:"order"`
}

var _ bodyer = externalContactEditCorpTagReq{}

func (x externalContactEditCorpTagReq) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		// should never happen unless OOM or similar bad things
		// TODO: error_chain
		return nil, err
	}

	return result, nil
}

// externalContactEditCorpTagResp 编辑企业客户标签
type externalContactEditCorpTagResp struct {
	CommonResp
}

// externalContactDelCorpTagReq 删除企业客户标签
type externalContactDelCorpTagReq struct {
	TagID   []string `json:"tag_id"`
	GroupID []string `json:"group_id"`
}

var _ bodyer = externalContactDelCorpTagReq{}

func (x externalContactDelCorpTagReq) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		// should never happen unless OOM or similar bad things
		// TODO: error_chain
		return nil, err
	}

	return result, nil
}

// externalContactDelCorpTagResp 删除企业客户标签
type externalContactDelCorpTagResp struct {
	CommonResp
}

// externalContactMarkTagReq 编辑企业客户标签
type externalContactMarkTagReq struct {
	UserID         string   `json:"userid"`
	ExternalUserID string   `json:"external_userid"`
	AddTag         []string `json:"add_tag"`
	RemoveTag      []string `json:"remove_tag"`
}

var _ bodyer = externalContactMarkTagReq{}

func (x externalContactMarkTagReq) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		// should never happen unless OOM or similar bad things
		// TODO: error_chain
		return nil, err
	}

	return result, nil
}

// externalContactMarkTagResp 编辑企业客户标签
type externalContactMarkTagResp struct {
	CommonResp
}

// jsCode2SessionReq 临时登录凭证校验
type jsCode2SessionReq struct {
	JSCode string
}

var _ urlValuer = jsCode2SessionReq{}

func (x jsCode2SessionReq) intoURLValues() url.Values {
	return url.Values{
		"js_code":    {x.JSCode},
		"grant_type": {"authorization_code"},
	}
}

// jsCode2SessionResp 临时登录凭证校验
type jsCode2SessionResp struct {
	CommonResp
	JSCodeSession
}

// JSCodeSession 临时登录凭证
type JSCodeSession struct {
	CorpID     string `json:"corpid"`
	UserID     string `json:"userid"`
	SessionKey string `json:"session_key"`
}

type msgAuditListPermitUserReq struct {
	MsgAuditEdition MsgAuditEdition `json:"type"`
}

var _ bodyer = msgAuditListPermitUserReq{}

func (x msgAuditListPermitUserReq) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type msgAuditListPermitUserResp struct {
	CommonResp
	IDs []string `json:"ids"`
}

type msgAuditCheckSingleAgreeReq struct {
	Infos []CheckMsgAuditSingleAgreeUserInfo `json:"info"`
}

var _ bodyer = msgAuditCheckSingleAgreeReq{}

func (x msgAuditCheckSingleAgreeReq) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type msgAuditCheckSingleAgreeResp struct {
	CommonResp
	AgreeInfo []struct {
		UserID           string              `json:"userid"`
		ExternalOpenID   string              `json:"exteranalopenid"`
		AgreeStatus      MsgAuditAgreeStatus `json:"agree_status"`
		StatusChangeTime int                 `json:"status_change_time"`
	} `json:"agreeinfo"`
}

func (x msgAuditCheckSingleAgreeResp) intoCheckSingleAgreeInfoList() (resp []CheckMsgAuditSingleAgreeInfo) {
	for _, agreeInfo := range x.AgreeInfo {
		resp = append(resp, CheckMsgAuditSingleAgreeInfo{
			CheckMsgAuditSingleAgreeUserInfo: CheckMsgAuditSingleAgreeUserInfo{
				UserID:         agreeInfo.UserID,
				ExternalOpenID: agreeInfo.ExternalOpenID,
			},
			AgreeStatus:      agreeInfo.AgreeStatus,
			StatusChangeTime: time.Unix(int64(agreeInfo.StatusChangeTime), 0),
		})
	}
	return resp
}

type msgAuditCheckRoomAgreeReq struct {
	RoomID string `json:"roomid"`
}

var _ bodyer = msgAuditCheckRoomAgreeReq{}

func (x msgAuditCheckRoomAgreeReq) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type msgAuditCheckRoomAgreeResp struct {
	CommonResp
	AgreeInfo []struct {
		StatusChangeTime int                 `json:"status_change_time"`
		AgreeStatus      MsgAuditAgreeStatus `json:"agree_status"`
		ExternalOpenID   string              `json:"exteranalopenid"`
	} `json:"agreeinfo"`
}

func (x msgAuditCheckRoomAgreeResp) intoCheckRoomAgreeInfoList() (resp []CheckMsgAuditRoomAgreeInfo) {
	for _, agreeInfo := range x.AgreeInfo {
		resp = append(resp, CheckMsgAuditRoomAgreeInfo{
			StatusChangeTime: time.Unix(int64(agreeInfo.StatusChangeTime), 0),
			AgreeStatus:      agreeInfo.AgreeStatus,
			ExternalOpenID:   agreeInfo.ExternalOpenID,
		})
	}
	return resp
}

type msgAuditGetGroupChatReq struct {
	RoomID string `json:"roomid"`
}

var _ bodyer = msgAuditGetGroupChatReq{}

func (x msgAuditGetGroupChatReq) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type msgAuditGetGroupChatResp struct {
	CommonResp
	Members []struct {
		MemberID int `json:"memberid"`
		JoinTime int `json:"jointime"`
	} `json:"members"`
	RoomName       string `json:"roomname"`
	Creator        string `json:"creator"`
	RoomCreateTime int    `json:"room_create_time"`
	Notice         string `json:"notice"`
}

func (x msgAuditGetGroupChatResp) intoGroupChat() (resp MsgAuditGroupChat) {
	resp.Creator = x.Creator
	resp.Notice = x.Notice
	resp.RoomName = x.RoomName
	resp.RoomCreateTime = time.Unix(int64(x.RoomCreateTime), 0)
	for _, member := range x.Members {
		resp.Members = append(resp.Members, MsgAuditGroupChatMember{
			MemberID: member.MemberID,
			JoinTime: time.Unix(int64(member.JoinTime), 0),
		})
	}
	return resp
}

// externalContactListReq 获取客户列表
type externalContactListReq struct {
	UserID string `json:"userid"`
}

var _ urlValuer = externalContactListReq{}

func (x externalContactListReq) intoURLValues() url.Values {
	return url.Values{
		"userid": {x.UserID},
	}
}

// externalContactListResp 获取客户列表
type externalContactListResp struct {
	CommonResp

	ExternalUserID []string `json:"external_userid"`
}

// externalContactGetReq 获取客户详情
type externalContactGetReq struct {
	ExternalUserID string `json:"external_userid"`
}

var _ urlValuer = externalContactGetReq{}

func (x externalContactGetReq) intoURLValues() url.Values {
	return url.Values{
		"external_userid": {x.ExternalUserID},
	}
}

// externalContactGetResp 获取客户详情
type externalContactGetResp struct {
	CommonResp
	ExternalContactInfo
}

// ExternalContactInfo 外部联系人信息
type ExternalContactInfo struct {
	ExternalContact ExternalContact `json:"external_contact"`
	FollowUser      []FollowUser    `json:"follow_user"`
}

// ExternalContactBatchInfo 外部联系人信息
type ExternalContactBatchInfo struct {
	ExternalContact ExternalContact `json:"external_contact"`
	FollowInfo      FollowInfo      `json:"follow_info"`
}

// BatchListExternalContactsResp 外部联系人信息
type BatchListExternalContactsResp struct {
	Result     []ExternalContactBatchInfo
	NextCursor string
}

// externalContactBatchListReq 批量获取客户详情
type externalContactBatchListReq struct {
	UserID string `json:"userid"`
	Cursor string `json:"cursor"`
	Limit  int    `json:"limit"`
}

var _ bodyer = externalContactBatchListReq{}

func (x externalContactBatchListReq) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		// should never happen unless OOM or similar bad things
		// TODO: error_chain
		return nil, err
	}

	return result, nil
}
