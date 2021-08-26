package workwx

import (
	"strconv"
	"time"
)

// GetOATemplateDetail 获取审批模板详情
func (c *App) GetOATemplateDetail(templateID string) (*OATemplateDetail, error) {
	resp, err := c.execOAGetTemplateDetail(oaGetTemplateDetailReq{
		TemplateID: templateID,
	})
	if err != nil {
		return nil, err
	}
	return &resp.OATemplateDetail, nil
}

// ApplyOAEvent 提交审批申请
func (c *App) ApplyOAEvent(applyInfo OAApplyEvent) (string, error) {
	resp, err := c.execOAApplyEvent(oaApplyEventReq{
		OAApplyEvent: applyInfo,
	})
	if err != nil {
		return "", err
	}
	return resp.SpNo, nil
}

// GetOAApprovalInfo 批量获取审批单号
func (c *App) GetOAApprovalInfo(req GetOAApprovalInfoReq) ([]string, error) {
	resp, err := c.execOAGetApprovalInfo(oaGetApprovalInfoReq{
		StartTime: strconv.FormatInt(req.StartTime.Unix(), 10),
		EndTime:   strconv.FormatInt(req.EndTime.Unix(), 10),
		Cursor:    req.Cursor,
		Size:      req.Size,
		Filters:   req.Filters,
	})
	if err != nil {
		return nil, err
	}
	return resp.SpNoList, nil
}

// GetOAApprovalDetail 提交审批申请
func (c *App) GetOAApprovalDetail(spNo string) (*OAApprovalDetail, error) {
	resp, err := c.execOAGetApprovalDetail(oaGetApprovalDetailReq{
		SpNo: spNo,
	})
	if err != nil {
		return nil, err
	}
	return &resp.Info, nil
}

// GetOAApprovalInfoReq 批量获取审批单号请求
type GetOAApprovalInfoReq struct {
	// StartTime 审批单提交的时间范围，开始时间，UNix时间戳
	StartTime time.Time
	// EndTime 审批单提交的时间范围，结束时间，Unix时间戳
	EndTime time.Time
	// Cursor 分页查询游标，默认为0，后续使用返回的next_cursor进行分页拉取
	Cursor int
	// Size 一次请求拉取审批单数量，默认值为100，上限值为100
	Size uint32
	// Filters 筛选条件，可对批量拉取的审批申请设置约束条件，支持设置多个条件
	Filters []OAApprovalInfoFilter
}

// OAApprovalInfo 审批申请状态变化回调通知
type OAApprovalInfo struct {
	// SpNo 审批编号
	SpNo string `xml:"SpNo"`
	// SpName 审批申请类型名称（审批模板名称）
	SpName string `xml:"SpName"`
	// SpStatus 申请单状态：1-审批中；2-已通过；3-已驳回；4-已撤销；6-通过后撤销；7-已删除；10-已支付
	SpStatus string `xml:"SpStatus"`
	// TemplateID 审批模板id。可在“获取审批申请详情”、“审批状态变化回调通知”中获得，也可在审批模板的模板编辑页面链接中获得。
	TemplateID string `xml:"TemplateId"`
	// ApplyTime 审批申请提交时间,Unix时间戳
	ApplyTime string `xml:"ApplyTime"`
	// Applicant 申请人信息
	Applicant OAApprovalInfoApplicant `xml:"Applyer"`
	// SpRecord 审批流程信息，可能有多个审批节点。
	SpRecord []OAApprovalInfoSpRecord `xml:"SpRecord"`
	// Notifier 抄送信息，可能有多个抄送节点
	Notifier OAApprovalInfoNotifier `xml:"Notifyer"`
	// Comments 审批申请备注信息，可能有多个备注节点
	Comments []OAApprovalInfoComment `xml:"Comments"`
	// StatusChangeEvent 审批申请状态变化类型：1-提单；2-同意；3-驳回；4-转审；5-催办；6-撤销；8-通过后撤销；10-添加备注
	StatusChangeEvent string `xml:"StatuChangeEvent"`
}

// OAApprovalInfoApplicant 申请人信息
type OAApprovalInfoApplicant struct {
	// UserID 申请人userid
	UserID string `xml:"UserId"`
	// Party 申请人所在部门pid
	Party string `xml:"Party"`
}

// OAApprovalInfoSpRecord 审批流程信息，可能有多个审批节点。
type OAApprovalInfoSpRecord struct {
	// SpStatus 审批节点状态：1-审批中；2-已同意；3-已驳回；4-已转审
	SpStatus string `xml:"SpStatus"`
	// ApproverAttr 节点审批方式：1-或签；2-会签
	ApproverAttr string `xml:"ApproverAttr"`
	// Details 审批节点详情。当节点为标签或上级时，一个节点可能有多个分支
	Details []OAApprovalInfoSpRecordDetail `xml:"Details"`
}

// OAApprovalInfoSpRecordDetail 审批节点详情。当节点为标签或上级时，一个节点可能有多个分支
type OAApprovalInfoSpRecordDetail struct {
	// Approver 分支审批人
	Approver OAApprovalInfoSpRecordDetailApprover `xml:"Approver"`
	// Speech 审批意见字段
	Speech string `xml:"Speech"`
	// SpStatus 分支审批人审批状态：1-审批中；2-已同意；3-已驳回；4-已转审
	SpStatus string `xml:"SpStatus"`
	// SpTime 节点分支审批人审批操作时间，0为尚未操作
	SpTime string `xml:"SpTime"`
	// Attach 节点分支审批人审批意见附件，赋值为media_id具体使用请参考：文档-获取临时素材
	Attach []string `xml:"Attach"`
}

// OAApprovalInfoSpRecordDetailApprover 分支审批人
type OAApprovalInfoSpRecordDetailApprover struct {
	// UserID 分支审批人userid
	UserID string `xml:"UserId"`
}

// OAApprovalInfoNotifier 抄送信息，可能有多个抄送节点
type OAApprovalInfoNotifier struct {
	// UserID 节点抄送人userid
	UserID string `xml:"UserId"`
}

// OAApprovalInfoComment 审批申请备注信息，可能有多个备注节点
type OAApprovalInfoComment struct {
	// CommentUserInfo 备注人信息
	CommentUserInfo OAApprovalInfoCommentUserInfo `xml:"CommentUserInfo"`
	// CommentTime 备注提交时间
	CommentTime string `xml:"CommentTime"`
	// CommentContent 备注文本内容
	CommentContent string `xml:"CommentContent"`
	// CommentID 备注id
	CommentID string `xml:"CommentId"`
	// Attach 备注意见附件，值是附件media_id具体使用请参考：文档-获取临时素材
	Attach []string `xml:"Attach"`
}

// OAApprovalInfoCommentUserInfo 备注人信息
type OAApprovalInfoCommentUserInfo struct {
	// UserID 备注人userid
	UserID string `xml:"UserId"`
}
