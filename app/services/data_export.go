package services

import (
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/gogf/gf/os/gfile"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"openscrm/app/constants"
	"openscrm/app/models"
	"openscrm/app/requests"
	"openscrm/common/app"
	"openscrm/common/delay_queue"
	"openscrm/common/id_generator"
	"openscrm/common/log"
	"openscrm/conf"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type DataExport struct {
	dataExportRepo    models.DataExportRepo
	groupChatRepo     models.GroupChat
	customerStaffRepo models.CustomerStaff
	repo              models.DataExportRepo
	relationHistory   models.CustomerStaffRelationHistory
}

func NewDataExport() *DataExport {
	return &DataExport{
		repo:              models.DataExport{},
		dataExportRepo:    models.DataExport{},
		groupChatRepo:     models.GroupChat{},
		customerStaffRepo: models.CustomerStaff{},
		relationHistory:   models.CustomerStaffRelationHistory{},
	}
}

func (o DataExport) ExportGroupChatList(req requests.QueryGroupChatReq, ExtCorpID string) (string, error) {
	groupChats, err := o.groupChatRepo.GetAll(req, ExtCorpID)
	if err != nil {
		log.Sugar.Errorw("groupChatRepo.GetAll", "err", err)
		return "", err
	}

	exportTime := time.Now().Format(constants.DateTimeLayout)
	filename := "/" + path.Join(
		string(constants.DataExportTypeGroupChat),
		time.Now().Format(constants.DateLayout),
		ExtCorpID,
		constants.DataExportGroupChatListPrefix+".xlsx",
	)
	fullPath := filepath.Join(
		conf.Settings.Storage.LocalRootPath,
		filename,
	)
	if !gfile.Exists(filepath.Dir(fullPath)) {
		gfile.Mkdir(filepath.Dir(fullPath))
	}

	file := excelize.NewFile()
	titles := []string{"客户群名称", "群主", "群标签", "群人数", "当日群人数", "当日入群", "当日退群", "创群时间", "群ID"}
	sheetIndex := file.NewSheet(constants.DataExportGroupChatListSheetName)
	file.DeleteSheet("Sheet1")
	file.SetActiveSheet(sheetIndex)

	err = o.PrettifySheet(constants.DataExportGroupChatListSheetName, file, exportTime, titles)
	if err != nil {
		log.Sugar.Error(err)
		return "", err
	}

	for k, v := range groupChats {
		values := []string{
			v.Name,
			v.Owner,
			"", //strings.Join(v.TagList, ","),
			strconv.Itoa(len(v.MemberList)),
			strconv.Itoa(int(v.Total)),
			strconv.Itoa(int(v.TodayJoinMemberNum)),
			strconv.Itoa(int(v.TodayQuitMemberNum)),
			v.CreateTime.Format(constants.DateTimeLayout),
			v.ExtChatID,
		}
		err = file.SetSheetRow(constants.DataExportGroupChatListSheetName, fmt.Sprint("A", k+3), &values)
		if err != nil {
			log.Sugar.Errorw("create new sheet in excel failed", "err", err)
			return "", err
		}
	}

	err = file.SaveAs(fullPath)
	if err != nil {
		log.Sugar.Errorw("save export data failed", "req", req, "err", err)
	}
	return fullPath, nil
}

func (o DataExport) ExportDeleteStaffWarningList(req requests.DataExportReq, extCorpID string) (string, error) {
	log.Sugar.Infow("ExportDeleteStaffWarningList", "req", req)
	customerLossInfos, _, err := o.relationHistory.QueryCustomerDeleteStaff(req.DataFilter.QueryCustomerLossesReq, "", &app.Sorter{}, &app.Pager{})
	if err != nil {
		log.Sugar.Errorw("QueryCustomerDeleteStaff failed", "err", err)
		return "", err
	}

	exportTime := time.Now().Format(constants.DateTimeLayout)
	filename := "/" + path.Join(
		string(constants.DataExportTypeDeleteStaffWarning),
		time.Now().Format(constants.DateLayout),
		extCorpID,
		constants.DataExportDeleteStaffFilenamePrefix+".xlsx",
	)
	fullPath := filepath.Join(
		conf.Settings.Storage.LocalRootPath,
		filename,
	)
	if !gfile.Exists(filepath.Dir(fullPath)) {
		gfile.Mkdir(filepath.Dir(fullPath))
	}

	file := excelize.NewFile()
	file.NewSheet(constants.DataExportDeleteStaffListSheetName)
	file.DeleteSheet("Sheet1")

	titles := []string{"流失客户", "所属客服", "标签", "流失时间", "添加时间", "添加员工时长/天"}
	err = o.PrettifySheet(constants.DataExportDeleteStaffListSheetName, file, exportTime, titles)
	if err != nil {
		log.Sugar.Error(err)
		return "", err
	}

	for k, v := range customerLossInfos {
		values := []string{
			v.ExtStaffID,
			v.StaffName,
			strings.Join(v.ExtTagIDs, ","),
			v.CustomerDeleteStaffAt.Format(constants.DateTimeLayout),
			v.RelationCreateAt.Format(constants.DateTimeLayout),
			strconv.Itoa(int(v.InConnectionTimeRange)),
		}
		err = file.SetSheetRow(constants.DataExportDeleteStaffListSheetName, fmt.Sprint("A", k+3), &values)
		if err != nil {
			log.Sugar.Errorw("write excel failed", "err", err)
			return "", err
		}
	}

	err = file.SaveAs(fullPath)
	if err != nil {
		log.Sugar.Errorw("save export data failed", "req", req, "err", err)
		return "", err
	}

	return fullPath, nil
}

func (o DataExport) PrettifySheet(sheetName string, file *excelize.File, exportTime string, titles []string) error {
	colCnt := string(rune(int('A') + len(titles)))
	err := file.SetColWidth(sheetName, "A", "H", 18)
	if err != nil {
		log.Sugar.Error("set excel col width failed", err)
		return err
	}

	err = file.SetColWidth(sheetName, "I", "I", 38)
	if err != nil {
		log.Sugar.Errorw("et excel col width failed", "err", err)
		return err
	}

	titleName := []string{fmt.Sprintf("%s(导出时间: %s)", sheetName, exportTime)}
	err = file.SetSheetRow(sheetName, fmt.Sprint("A", 1), &titleName)
	if err != nil {
		log.Sugar.Errorw("set row  failed", "err", err)
		return err
	}

	styleCenter, err := file.NewStyle(`{"font":{"bold":true},"alignment":{"horizontal":"center","vertical":"center"}}`)
	if err != nil {
		log.Sugar.Errorw("create new style  failed", "err", err)
		return err
	}

	err = file.SetCellStyle(sheetName, "A1", colCnt+"1", styleCenter)
	if err != nil {
		log.Sugar.Errorw("create new sheet style failed", "err", err)
		return err
	}
	err = file.SetRowHeight(sheetName, 1, 30)
	if err != nil {
		log.Sugar.Errorw("set new sheet height failed", "err", err)
		return err
	}
	err = file.MergeCell(sheetName, "A1", colCnt+"1")
	if err != nil {
		log.Sugar.Errorw("merge excel cell failed", "err", err)
		return err
	}

	err = file.SetSheetRow(sheetName, fmt.Sprint("A", 2), &titles)
	if err != nil {
		log.Sugar.Error("set sheet row in excel failed", "err", err)
		return err
	}
	return nil
}

// ---------------- backup for async task ------------------

func (o DataExport) GenDataExportTask(req requests.DataExportReq, staffAdmin models.Staff) (string, error) {
	exportJob := requests.DataExportJob{
		DataExportReq: req,
		ExtCorpID:     staffAdmin.ExtCorpID,
	}
	reqBytes, err := json.Marshal(exportJob)
	if err != nil {
		log.Sugar.Errorw("marshal job failed", "err", err)
		return "", err
	}
	job := delay_queue.Job{
		Topic:     constants.DataExportTopic,
		ID:        id_generator.StringID(),
		ExecuteAt: time.Now().Unix(),
		TTR:       10,
		Body:      string(reqBytes),
	}
	err = delay_queue.Add(job)
	if err != nil {
		return "", err
	}
	return job.ID, err
}

func (o DataExport) Get(id string, staff models.Staff) (requests.DataExportTaskExeResult, error) {
	var res requests.DataExportTaskExeResult
	data, err := o.repo.Get(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			res.Status = string(constants.AsyncTaskStatusCreating)
			return res, nil
		}
		return res, err
	}

	err = copier.Copy(&res, data)
	return res, err
}
