package workwx

import "encoding/json"

type oaGetTemplateDetailReq struct {
	TemplateID string `json:"template_id"`
}

var _ bodyer = oaGetTemplateDetailReq{}

func (x oaGetTemplateDetailReq) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type oaGetTemplateDetailResp struct {
	CommonResp
	OATemplateDetail
}

type oaApplyEventReq struct {
	OAApplyEvent
}

var _ bodyer = oaApplyEventReq{}

func (x oaApplyEventReq) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type oaApplyEventResp struct {
	CommonResp
	// SpNo 表单提交成功后，返回的表单编号
	SpNo string `json:"sp_no"`
}

type oaGetApprovalInfoReq struct {
	StartTime string                 `json:"starttime"`
	EndTime   string                 `json:"endtime"`
	Cursor    int                    `json:"cursor"`
	Size      uint32                 `json:"size"`
	Filters   []OAApprovalInfoFilter `json:"filters"`
}

var _ bodyer = oaGetApprovalInfoReq{}

func (x oaGetApprovalInfoReq) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type oaGetApprovalInfoResp struct {
	CommonResp
	// SpNoList 审批单号列表，包含满足条件的审批申请
	SpNoList []string `json:"sp_no_list"`
}

type oaGetApprovalDetailReq struct {
	// SpNo 审批单编号。
	SpNo string `json:"sp_no"`
}

var _ bodyer = oaGetApprovalDetailReq{}

func (x oaGetApprovalDetailReq) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type oaGetApprovalDetailResp struct {
	CommonResp
	// Info 审批申请详情
	Info OAApprovalDetail `json:"info"`
}

// execOAGetTemplateDetail 获取审批模板详情
func (c *App) execOAGetTemplateDetail(req oaGetTemplateDetailReq) (oaGetTemplateDetailResp, error) {
	var resp oaGetTemplateDetailResp
	err := c.executeWXApiJSONPost("/cgi-bin/oa/gettemplatedetail", req, &resp, true)
	if err != nil {
		return oaGetTemplateDetailResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return oaGetTemplateDetailResp{}, bizErr
	}

	return resp, nil
}

// execOAApplyEvent 提交审批申请
func (c *App) execOAApplyEvent(req oaApplyEventReq) (oaApplyEventResp, error) {
	var resp oaApplyEventResp
	err := c.executeWXApiJSONPost("/cgi-bin/oa/applyevent", req, &resp, true)
	if err != nil {
		return oaApplyEventResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return oaApplyEventResp{}, bizErr
	}

	return resp, nil
}

// execOAGetApprovalInfo 批量获取审批单号
func (c *App) execOAGetApprovalInfo(req oaGetApprovalInfoReq) (oaGetApprovalInfoResp, error) {
	var resp oaGetApprovalInfoResp
	err := c.executeWXApiJSONPost("/cgi-bin/oa/getapprovalinfo", req, &resp, true)
	if err != nil {
		return oaGetApprovalInfoResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return oaGetApprovalInfoResp{}, bizErr
	}

	return resp, nil
}

// execOAGetApprovalDetail 获取审批申请详情
func (c *App) execOAGetApprovalDetail(req oaGetApprovalDetailReq) (oaGetApprovalDetailResp, error) {
	var resp oaGetApprovalDetailResp
	err := c.executeWXApiJSONPost("/cgi-bin/oa/getapprovaldetail", req, &resp, true)
	if err != nil {
		return oaGetApprovalDetailResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return oaGetApprovalDetailResp{}, bizErr
	}

	return resp, nil
}
