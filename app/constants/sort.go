package constants

// SortField 排序字段
// id
// created_at
// updated_at
// sort_weight
type SortField string

const (
	SortFieldID         SortField = "id"
	SortFieldCreatedAt  SortField = "created_at"
	SortFieldUpdatedAt  SortField = "updated_at"
	SortFieldSortWeight SortField = "sort_weight"
)

// SortType 排序方式，升序，降序，默认降序
// asc
// desc
type SortType string

const (
	SortTypeAsc  SortType = "asc"
	SortTypeDesc SortType = "desc"
)
