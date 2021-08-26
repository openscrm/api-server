package app

type Pager struct {
	// Page 页码
	Page int `json:"page" form:"page" validate:"gte=0" gorm:"-"`
	// PageSize 每页数量
	PageSize int `json:"page_size" form:"page_size" validate:"gte=0" gorm:"-"`
	// TotalRows 总行数
	TotalRows int64 `json:"total_rows" gorm:"-"`
}

func (o *Pager) SetDefault() *Pager {
	if o.Page <= 0 || o.Page > 1000000 {
		o.Page = 1
	}
	if o.PageSize <= 0 || o.PageSize > 1000000 {
		o.PageSize = 10
	}
	return o
}

func (o *Pager) GetOffset() int {
	return (o.Page - 1) * o.PageSize
}

func (o *Pager) GetLimit() int {
	return o.PageSize
}
