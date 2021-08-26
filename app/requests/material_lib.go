package requests

import (
	"openscrm/app/constants"
	"openscrm/common/app"
)

type UploadMaterialReq struct {
	// 文件名
	Title string `json:"title"`
	// link-链接 poster-海报 video pdf ppt excel-表格 word-文档
	MaterialType string `json:"material_type" validate:"required,oneof=link poster video pdf ppt excel word" form:"material_type"`
	// 链接封面/图片/pdf 等地址
	FileUrl string `json:"file_url" validate:"required_if=Type poster,url"`
	// link
	Link string `json:"link" form:"link"`
	// 文件大小
	FileSize string `json:"file_size" validate:"omitempty"`
	// 摘要,连接类型素材的描述
	Digest string `json:"digest" validate:"omitempty"`
	// 素材标签
	MaterialTagList []string `json:"material_tag_list"`
}

type QueryMaterialReq struct {
	Title           string                     `json:"title" validate:"omitempty" form:"title"`
	MaterialType    string                     `json:"material_type" validate:"omitempty,oneof=link poster video pdf ppt excel word" form:"material_type"`
	MaterialTagList constants.StringArrayField `json:"material_tag_list" form:"material_tag_list" validate:"omitempty,gt=0"`
	app.Sorter      `form:"app_sorter"`
	app.Pager       `form:"app_pager"`
}

type UpdateMaterialReq struct {
	// link
	Link string `json:"link" form:"link"`
	// 内容
	Content string `json:"content"`
	// 文件名
	Title string `json:"title"`
	// 链接封面/图片/pdf 等地址
	FileUrl string `json:"file_url" validate:"required_if=Type poster,url"`
	// 文件大小
	FileSize string `json:"file_size" validate:"omitempty"`
	Digest   string `json:"digest" validate:"omitempty"`
	// 素材标签
	MaterialTagList constants.StringArrayField `form:"material_tag_list" json:"material_tag_list"`
}

type UpdateGetSidebarStatusReq struct {
	// 侧边栏开关 1-开 2-关
	Status constants.Boolean `json:"status" validate:"oneof=1 2"`
}
type GetSidebarStatusResp struct {
	// 侧边栏开关 1-开 2-关
	Status constants.Boolean `json:"status" validate:"oneof=1 2"`
}
