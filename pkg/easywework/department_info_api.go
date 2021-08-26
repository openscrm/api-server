package workwx

// deptListResp 部门列表响应
type deptListResp struct {
	CommonResp
	Department []*DeptInfo `json:"department"`
}

// execDeptList 获取部门列表
func (c *App) execDeptList(req deptListReq) (deptListResp, error) {
	var resp deptListResp
	err := c.executeWXApiGet("/cgi-bin/department/list", req, &resp, true)
	if err != nil {
		return deptListResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return deptListResp{}, bizErr
	}

	return resp, nil
}
