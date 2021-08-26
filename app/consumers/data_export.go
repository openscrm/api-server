package consumers

import (
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/gogf/gf/os/gfile"
	"openscrm/common/app"
	"path"
	"path/filepath"

	//"github.com/tealeg/xlsx"
	"net/http"
	"openscrm/app/constants"
	"openscrm/app/models"
	"openscrm/app/requests"
	"openscrm/common/delay_queue"
	"openscrm/common/log"
	"openscrm/common/storage"
	"openscrm/conf"
	"strconv"
	"strings"
	"time"
)

/*
  The following procedures present as a backup for async task, not in use.
*/

type DataExporter struct {
	dataExportRepo    models.DataExportRepo
	groupChatRepo     models.GroupChat
	customerStaffRepo models.CustomerStaff
	relationHistory   models.CustomerStaffRelationHistory
}

func NewDataExporter() *DataExporter {
	return &DataExporter{
		dataExportRepo:    models.DataExport{},
		groupChatRepo:     models.GroupChat{},
		customerStaffRepo: models.CustomerStaff{},
		relationHistory:   models.CustomerStaffRelationHistory{},
	}
}

func (o DataExporter) DataExport(job delay_queue.Job) error {
	exportJob := requests.DataExportJob{}
	err := json.Unmarshal([]byte(job.Body), &exportJob)
	if err != nil {
		return err
	}
	switch exportJob.Type {
	case string(constants.DataExportTypeGroupChat):
		//return o.ExportGroupChatList(job.ID, exportJob, )
	case string(constants.DataExportTypeDeleteStaffWarning):
		return o.ExportDeleteStaffWarningList(job.ID, exportJob)
	case string(constants.DataExportTypeDeleteCustomerWarning):
		return o.ExportDeleteCustomerList(job.ID, exportJob)
	case string(constants.DataExportTypeCustomer):
		//return o.ExportCustomer(job.ID, exportJob)
	default:
	}

	return nil
}

func (o DataExporter) ExportGroupChatList(jobID string, extCorpID string, req requests.QueryGroupChatReq) error {
	groupChats, err := o.groupChatRepo.GetAll(req, extCorpID)
	if err != nil {
		log.Sugar.Errorw("groupChatRepo.GetAll", "err", err)
		return err
	}

	exportTime := time.Now().Format(constants.DateTimeLayout)
	filename := "/" + path.Join(
		string(constants.DataExportTypeGroupChat),
		time.Now().Format(constants.DateLayout),
		extCorpID,
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
		return err
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
			return err
		}
	}
	url, err := storage.FileStorage.SignURL(filename, http.MethodGet, int64(time.Second*3600*7))
	if err != nil {
		return err
	}

	dataExportRes := models.DataExport{
		ExtCorpModel: models.ExtCorpModel{ID: jobID, ExtCorpID: extCorpID},
		ExportTime:   exportTime,
		Status:       string(constants.AsyncTaskStatusCreating),
		URL:          url,
		Type:         string(constants.DataExportTypeGroupChat),
	}

	err = file.SaveAs(fullPath)
	if err != nil {
		log.Sugar.Errorw("save export data failed", "req", req, "err", err)
		dataExportRes.Status = string(constants.AsyncTaskStatusFailed)
		return o.dataExportRepo.Create(dataExportRes)
	}

	dataExportRes.Status = string(constants.AsyncTaskStatusSuccess)
	return o.dataExportRepo.Create(dataExportRes)
}

func (o DataExporter) ExportDeleteStaffWarningList(jobID string, req requests.DataExportJob) error {
	log.Sugar.Infow("ExportDeleteStaffWarningList", "req", req)
	customerLossInfos, _, err := o.relationHistory.QueryCustomerDeleteStaff(req.DataFilter.QueryCustomerLossesReq, "", &app.Sorter{}, &app.Pager{})
	if err != nil {
		log.Sugar.Errorw("QueryCustomerDeleteStaff failed", "err", err)
		return err
	}

	exportTime := time.Now().Format(constants.DateTimeLayout)
	filename := "/" + path.Join(
		string(constants.DataExportTypeDeleteStaffWarning),
		time.Now().Format(constants.DateLayout),
		req.ExtCorpID,
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
		return err
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
			return err
		}
	}
	url, err := storage.FileStorage.SignURL(filename, http.MethodGet, int64(time.Second*3600*7))
	if err != nil {
		return err
	}

	dataExportRes := models.DataExport{
		ExtCorpModel: models.ExtCorpModel{ID: jobID, ExtCorpID: req.ExtCorpID},
		ExportTime:   exportTime,
		Status:       string(constants.AsyncTaskStatusCreating),
		URL:          url,
		Type:         string(constants.DataExportTypeDeleteStaffWarning),
	}

	err = file.SaveAs(fullPath)
	if err != nil {
		log.Sugar.Errorw("save export data failed", "req", req, "err", err)
		dataExportRes.Status = string(constants.AsyncTaskStatusFailed)
		return o.dataExportRepo.Create(dataExportRes)
	}

	dataExportRes.Status = string(constants.AsyncTaskStatusSuccess)
	return o.dataExportRepo.Create(dataExportRes)
}

func (o DataExporter) ExportDeleteCustomerList(jobID string, req requests.DataExportJob) error {
	log.Sugar.Infow("ExportDeleteCustomerList", "req", req)
	customerLossInfos, _, err := o.relationHistory.QueryStaffDeleteCustomer(req.DataFilter.QueryStaffDeleteCustomerHistoryReq, "", &app.Pager{}, &app.Sorter{})
	if err != nil {
		log.Sugar.Errorw("QueryCustomerDeleteStaff failed", "err", err)
		return err
	}

	exportTime := time.Now().Format(constants.DateTimeLayout)
	filename := "/" + path.Join(
		string(constants.DataExportTypeDeleteCustomerWarning),
		time.Now().Format(constants.DateLayout),
		req.ExtCorpID,
		constants.DataExportDeleteCustomerFilenamePrefix+".xlsx",
	)
	fullPath := filepath.Join(
		conf.Settings.Storage.LocalRootPath,
		filename,
	)
	if !gfile.Exists(filepath.Dir(fullPath)) {
		gfile.Mkdir(filepath.Dir(fullPath))
	}

	file := excelize.NewFile()
	file.NewSheet(constants.DataExportDeleteCustomerListSheetName)
	file.DeleteSheet("Sheet1")
	titles := []string{"删除客户", "操作人", "删除时间", "添加好友时间", "unionid"}
	err = o.PrettifySheet(constants.DataExportDeleteCustomerListSheetName, file, exportTime, titles)
	if err != nil {
		log.Sugar.Error(err)
		return err
	}

	for k, v := range customerLossInfos {
		values := []string{
			v.ExtCustomerName,
			v.StaffName,
			v.RelationDeleteAt.Format(constants.DateTimeLayout),
			v.RelationCreateAt.Format(constants.DateTimeLayout),
			"",
		}
		err = file.SetSheetRow(constants.DataExportDeleteCustomerListSheetName, fmt.Sprint("A", k+3), &values)
		if err != nil {
			log.Sugar.Errorw("write excel failed", "err", err)
			return err
		}
	}
	url, err := storage.FileStorage.SignURL(filename, http.MethodGet, int64(time.Second*3600*7))
	if err != nil {
		return err
	}

	dataExportRes := models.DataExport{
		ExtCorpModel: models.ExtCorpModel{ID: jobID, ExtCorpID: req.ExtCorpID},
		ExportTime:   exportTime,
		Status:       string(constants.AsyncTaskStatusCreating),
		URL:          url,
		Type:         string(constants.DataExportTypeDeleteCustomerWarning),
	}

	err = file.SaveAs(fullPath)
	if err != nil {
		log.Sugar.Errorw("save export data failed", "req", req, "err", err)
		dataExportRes.Status = string(constants.AsyncTaskStatusFailed)
		return o.dataExportRepo.Create(dataExportRes)
	}

	dataExportRes.Status = string(constants.AsyncTaskStatusSuccess)
	return o.dataExportRepo.Create(dataExportRes)
}

func (o DataExporter) PrettifySheet(sheetName string, file *excelize.File, exportTime string, titles []string) error {
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
