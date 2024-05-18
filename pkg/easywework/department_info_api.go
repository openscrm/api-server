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

// deptSimpleListResp 部门ID列表响应
type deptSimpleListResp struct {
	CommonResp
	DepartmentId []*DeptSimpleInfo `json:"department_id"`
}

// execDeptSimpleList 获取子部门ID列表
// https://developer.work.weixin.qq.com/document/path/95350
func (c *App) execDeptSimpleList(req deptSimpleListReq) (deptSimpleListResp, error) {
	var resp deptSimpleListResp
	err := c.executeWXApiGet("/cgi-bin/department/simplelist", req, &resp, true)
	if err != nil {
		return deptSimpleListResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return deptSimpleListResp{}, bizErr
	}

	return resp, nil
}
