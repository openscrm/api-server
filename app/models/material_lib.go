package models

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"openscrm/app/constants"
	"openscrm/app/requests"
)

type Material struct {
	ExtCorpModel
	// 素材类型
	MaterialType string `json:"material_type" gorm:"char(12);index;"`
	// 标题
	Title string `json:"title" gorm:"type:char(64);index;comment:素材标题"`
	// 文件大小
	FileSize string `json:"file_size" gorm:"type:int;comment:素材文件大小"`
	// 素材地址
	FileUrl string `json:"url" gorm:"type:varchar(512);comment:素材下载地址"`
	// Link
	Link string `json:"link" gorm:"type:varchar(255)"`
	// 链接摘要
	Digest          string                     `json:"digest" gorm:"type:text;comment:链接类型素材的摘要"`
	MaterialTagList constants.StringArrayField `json:"material_tag_list" gorm:"type:json"`
	Timestamp
}

type MaterialWithTags struct {
	Material
	Tags []MaterialLibTag `json:"tags"`
}

func (m Material) Update(material Material) error {
	return DB.Model(&Material{}).
		Omit("id").
		Where("id = ?", material.ID).
		Updates(&material).Error
}

func (m Material) Create(material Material) error {
	return DB.Create(&material).Error
}

func (m Material) Delete(ids []string, extCorpID string) (int64, error) {
	result := DB.Where("ext_corp_id = ?", extCorpID).Where("id in (?)", ids).Delete(&Material{})
	err := result.Error
	if err != nil {
		err = errors.Wrap(err, "Delete Material failed")
		return 0, err
	}

	return result.RowsAffected, result.Error
}

func (m Material) Query(req requests.QueryMaterialReq, extCorpID string) (items []Material, total int64, err error) {
	db := DB.Model(&Material{}).Preload(clause.Associations).Where("ext_corp_id = ?", extCorpID)
	if req.MaterialType != "" {
		db = db.Where("material_type = ?", req.MaterialType)
	}
	if req.Title != "" {
		db = db.Where("title like ?", req.Title+"%")
	}
	if len(req.MaterialTagList) > 0 {
		db = db.Where(func(db *gorm.DB) *gorm.DB {
			for _, tagID := range req.MaterialTagList {
				db = db.Or("json_contains(material_tag_list, json_array(?))", tagID)
			}
			return db
		}(DB))
	}

	err = db.Count(&total).Error
	if err != nil || total == 0 {
		err = errors.Wrap(err, "Count material failed")
		return nil, 0, err
	}

	req.Sorter.SetDefault()
	db = db.Order(clause.OrderByColumn{Column: clause.Column{Name: string(req.Sorter.SortField)}, Desc: req.Sorter.SortType == constants.SortTypeDesc})

	req.Pager.SetDefault()
	db = db.Offset(req.Pager.GetOffset()).Limit(req.Pager.GetLimit())

	err = db.Find(&items).Error
	if err != nil {
		err = errors.Wrap(err, "Find material failed")
		return nil, 0, err
	}
	return items, total, nil
}
