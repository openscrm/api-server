package models

type DataExport struct {
	ExtCorpModel
	// 导出时间
	ExportTime string `json:"export_time"`
	// 导出状态 creating-导出中 success-成功 failed-失败
	Status string `json:"status"`
	// 下载链接
	URL string `json:"url"`
	// 数据类型 group_chat_list
	Type string `json:"type"`
	Timestamp
}

func (d DataExport) Create(res DataExport) error {
	return DB.Create(&res).Error
}

func (d DataExport) Get(id string) (DataExport, error) {
	var data DataExport
	err := DB.Model(&DataExport{}).Where("id =?", id).First(&data).Error
	return data, err
}

type DataExportRepo interface {
	Get(id string) (DataExport, error)
	Create(res DataExport) error
}
