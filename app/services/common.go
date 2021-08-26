package services

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"openscrm/common/log"
)

// PrettifySheet
// Description: 格式化xlsx 表头
func PrettifySheet(sheetName string, file *excelize.File, exportTime string, titles []string) error {
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
