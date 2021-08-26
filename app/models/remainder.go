package models

import (
	"openscrm/common/log"
	"time"
)

type Remainder struct {
	ExtCorpModel
	Content string    `json:"content" gorm:"comment:提醒内容"`
	SendAt  time.Time `json:"send_at" gorm:"comment:提醒时间"`
	Timestamp
}

func (o Remainder) Create(remainder Remainder) error {
	return DB.Create(&remainder).Error
}

func (o Remainder) Delete(id string) (int64, error) {
	res := DB.Model(&Remainder{}).Where("id = ?", id).Delete(&Remainder{})
	return res.RowsAffected, res.Error
}

func (o Remainder) Update(id string, r Remainder) error {
	return DB.Model(&Remainder{}).Where("id = ?", id).Omit("id").Updates(&r).Error
}

func (o Remainder) Get(id string) (Remainder, error) {
	remainder := Remainder{}
	err := DB.Model(&Remainder{}).Where("id = ?", id).First(&remainder).Error
	if err != nil {
		log.Sugar.Errorw("get remainder failed", "id", id)
	}
	return remainder, err
}
