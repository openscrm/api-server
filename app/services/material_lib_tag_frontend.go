package services

import (
	"github.com/pkg/errors"
	"openscrm/app/models"
	"openscrm/common/app"
)

type MaterialLibTagFrontend struct {
	model models.MaterialLibTag
}

func NewMaterialLibTagFrontend() *MaterialLibTagFrontend {
	return &MaterialLibTagFrontend{model: models.MaterialLibTag{}}
}

func (o MaterialLibTagFrontend) Query(
	name string, creator models.Staff, pager *app.Pager, sorter *app.Sorter) (tags []models.MaterialLibTag, total int64, err error) {

	tags = make([]models.MaterialLibTag, 0)

	tag := models.MaterialLibTag{
		ExtCorpModel: models.ExtCorpModel{ExtCorpID: creator.ExtCorpID},
		Name:         name,
	}

	tags, total, err = o.model.Query(tag, sorter, pager)
	if err != nil {
		err = errors.WithStack(err)
		return nil, 0, err
	}

	return
}
