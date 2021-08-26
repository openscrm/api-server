package services

import (
	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"openscrm/app/models"
	"openscrm/app/requests"
	"openscrm/common/ecode"
	"openscrm/common/id_generator"
)

type CustomerRemarkService struct {
	infoRepo     models.CustomerInfo
	remarkRepo   models.CustomerRemark
	optionRepo   models.RemarkOption
	displayRules models.CustomerInfoDisplayRule
}

func NewCustomerRemarkService() *CustomerRemarkService {
	return &CustomerRemarkService{
		remarkRepo:   models.CustomerRemark{},
		optionRepo:   models.RemarkOption{},
		infoRepo:     models.CustomerInfo{},
		displayRules: models.CustomerInfoDisplayRule{},
	}
}

type InfoRemark struct {
	DisplayRules *models.CustomerInfoDisplayRule `json:"display_rules"`
	Remark       []*models.CustomerRemark        `json:"remark"`
}

func (cs CustomerRemarkService) Get(extCorpID string) (InfoRemark, error) {
	remarks, err := cs.remarkRepo.Get(extCorpID)
	if err != nil {
		return InfoRemark{}, err
	}

	displayRules := &models.CustomerInfoDisplayRule{ExtCorpID: extCorpID}
	infoDisplayRules, err := cs.displayRules.Get(displayRules)
	if err != nil {
		return InfoRemark{}, err
	}

	return InfoRemark{DisplayRules: infoDisplayRules, Remark: remarks}, nil
}

func (cs CustomerRemarkService) Create(req *requests.AddCustomerRemarkReq, extCorpID string) (remark models.CustomerRemark, err error) {
	remark = models.CustomerRemark{
		ExtCorpModel: models.ExtCorpModel{ID: id_generator.StringID()},
		FieldType:    req.FieldType,
		Name:         req.FieldName,
	}
	options := make([]models.RemarkOption, 0)
	if req.FieldType == "option_text" && len(req.OptionNameList) != 0 {
		for _, optionName := range req.OptionNameList {
			options = append(options, models.RemarkOption{Model: models.Model{ID: id_generator.StringID()}, Name: optionName})
		}
	}
	remark.Options = options
	err = cs.remarkRepo.Create(remark)
	mysqlErr := &mysql.MySQLError{}
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		err = ecode.DuplicateRemarkNameError
		return
	}

	return
}

func (cs CustomerRemarkService) Delete(ids []string, ExtCorpID string) error {
	return cs.remarkRepo.Delete(ids, ExtCorpID)
}

func (cs CustomerRemarkService) Update(req *requests.UpdateRemarkReq) (remark models.CustomerRemark, err error) {
	remark = models.CustomerRemark{ExtCorpModel: models.ExtCorpModel{ID: req.ID}, Name: req.Name}
	err = cs.remarkRepo.Update(remark)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}

// --------------------------------------

func (cs CustomerRemarkService) AddRemarkOption(req *requests.AddRemarkOptionReq) error {
	return cs.optionRepo.Create(models.RemarkOption{
		Model:    models.Model{ID: id_generator.StringID()},
		RemarkID: req.RemarkID,
		Name:     req.Name,
	})
}

func (cs CustomerRemarkService) UpdateRemarkOption(req *requests.UpdateRemarkOptionReq) error {
	textOption := &models.RemarkOption{
		Model: models.Model{ID: req.RemarkOptionID},
		Name:  req.Name,
	}
	err := cs.optionRepo.Update(textOption)
	if err != nil {
		mysqlErr := &mysql.MySQLError{}
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return ecode.DuplicateRemarkNameError
		}
		return err
	}
	return nil
}

func (cs CustomerRemarkService) DeleteRemarkOption(ids []string) error {
	return cs.optionRepo.Delete(ids)
}

func (cs CustomerRemarkService) ExchangeOrder(req *requests.ExchangeOrderReq) error {
	return cs.remarkRepo.ExchangeOrder(req.ID, req.ExchangeOrderID)
}
