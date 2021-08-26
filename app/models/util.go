package models

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CountField(db *gorm.DB, table interface{}, query string, args ...interface{}) (count int64, err error) {
	err = DB.Model(table).Where(query, args...).Count(&count).Error
	if err != nil {
		err = errors.Wrap(err, "CountField failed")
		return
	}

	return
}

func BatchUpsert(db *gorm.DB, table string, items interface{}, uniqueColumns []clause.Column) (err error) {
	err = db.Table(table).Clauses(clause.OnConflict{
		Columns:   uniqueColumns,
		UpdateAll: true,
	}).CreateInBatches(&items, 500).Error
	if err != nil {
		err = errors.Wrap(err, "BatchUpsert failed")
		return
	}
	return
}

// DeleteRefRecord 用于简化处理关联记录移除，从两个对象集合中pluck主键，再取差集，然后删除
//func DeleteRefRecord(tx *gorm.DB, model interface{}, objs1 interface{}, objs2 interface{}, primaryKey string) (err error) {
//	ids1, err := collection.NewObjCollection(objs1).Pluck(primaryKey).ToStrings()
//	if err != nil {
//		err = errors.Wrap(err, "Pluck failed")
//		return
//	}
//
//	ids2, err := collection.NewObjCollection(objs2).Pluck(primaryKey).ToStrings()
//	if err != nil {
//		err = errors.Wrap(err, "Pluck failed")
//		return
//	}
//
//	diffIDs, _ := funk.DifferenceString(ids1, ids2)
//	if len(diffIDs) > 0 {
//		// 会删除全部关联
//		err = tx.Select(clause.Associations).Where(strcase.ToSnake(primaryKey)+" in (?)", diffIDs).Delete(model).Error
//		if err != nil {
//			err = errors.Wrap(err, "Delete failed")
//			return
//		}
//	}
//	return
//}
