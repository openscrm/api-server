package services

import (
	"openscrm/app/models"
	"openscrm/common/app"
	"openscrm/common/id_generator"
)

type MaterialLibTag struct {
	model models.MaterialLibTag
}

func NewMaterialLibTag() *MaterialLibTag {
	return &MaterialLibTag{model: models.MaterialLibTag{}}
}

func (o MaterialLibTag) Create(names []string, creator models.Staff) ([]models.MaterialLibTag, error) {
	tags := make([]models.MaterialLibTag, 0)
	for _, name := range names {
		tag := models.MaterialLibTag{
			ExtCorpModel: models.ExtCorpModel{ID: id_generator.StringID(), ExtCorpID: creator.ExtCorpID, ExtCreatorID: creator.ExtID},
			Name:         name,
		}
		tags = append(tags, tag)
	}

	err := o.model.Create(tags)
	return tags, err
}

func (o MaterialLibTag) Delete(ids []string) (int64, error) {

	return o.model.Delete(ids)
}

func (o MaterialLibTag) Query(name string, creator models.Staff, pager *app.Pager, sorter *app.Sorter) ([]models.MaterialLibTag, int64, error) {

	tag := models.MaterialLibTag{
		ExtCorpModel: models.ExtCorpModel{ExtCorpID: creator.ExtCorpID},
		Name:         name,
	}

	tags := make([]models.MaterialLibTag, 0)
	tags, total, err := o.model.Query(tag, sorter, pager)
	if err != nil {
		return nil, 0, err
	}
	return tags, total, nil
}
