package workwx

// 从2022年8月15日10点开始，“企业管理后台 - 管理工具 - 通讯录同步”的新增IP将不能再调用此接口
// 企业可通过「获取部门ID列表」接口获取部门ID列表。查看调整详情。
// 参考：https://developer.work.weixin.qq.com/document/path/96079#40802

// SimpleListAllDepartments 获取全量部门ID。
func (c *App) SimpleListAllDepartments() ([]*DeptSimpleInfo, error) {
	resp, err := c.execDeptSimpleList(deptSimpleListReq{
		ID: 0,
	})
	if err != nil {
		return nil, err
	}

	return resp.DepartmentId, nil
}

// ListAllDepartments 获取全量组织架构。
func (c *App) ListAllDepartments() ([]*DeptInfo, error) {
	resp, err := c.execDeptList(deptListReq{
		HaveID: false,
		ID:     0,
	})
	if err != nil {
		return nil, err
	}

	return resp.Department, nil
}

// ListDepartments 获取指定部门及其下的子部门。
func (c *App) ListDepartments(id int64) ([]*DeptInfo, error) {
	resp, err := c.execDeptList(deptListReq{
		HaveID: true,
		ID:     id,
	})
	if err != nil {
		return nil, err
	}

	return resp.Department, nil
}
