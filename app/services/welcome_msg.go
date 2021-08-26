package services

import (
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"io"
	"io/ioutil"
	"net/http"
	url2 "net/url"
	"openscrm/app/constants"
	"openscrm/app/models"
	"openscrm/app/requests"
	"openscrm/common/app"
	"openscrm/common/ecode"
	"openscrm/common/id_generator"
	"openscrm/common/log"
	"openscrm/common/storage"
	"openscrm/common/we_work"
	work_wx "openscrm/pkg/easywework"
	"strings"
)

type WelcomeMsgService struct {
	departmentRepo models.Department
	staffRepo      models.Staff
	msgRepo        models.WelcomeMsg
}

func NewWelcomeMsgService() *WelcomeMsgService {
	return &WelcomeMsgService{
		msgRepo:        models.WelcomeMsg{},
		staffRepo:      models.Staff{},
		departmentRepo: models.Department{},
	}
}

func GenGetURLFromPutURL(url string) (getURL string, err error) {
	unescape, err := url2.QueryUnescape(url)
	if err != nil {
		return
	}
	parse, err := url2.Parse(unescape)
	if err != nil {
		return
	}

	expiredSec := int64(3600)
	obj := strings.TrimPrefix(parse.Fragment, "/")
	getURL, err = storage.FileStorage.SignURL(obj, http.MethodGet, expiredSec)
	if err != nil {
		return
	}
	return
}

func GetObjFromSignedURL(url string) (obj string, err error) {
	unescape, err := url2.QueryUnescape(url)
	if err != nil {
		return
	}
	parse, err := url2.Parse(unescape)
	if err != nil {
		return
	}
	obj = strings.TrimPrefix(parse.Path, "/")
	return
}

// Create
// Description 创建员工添加客户时自动发送给客户的欢迎语
// Detail
//	支持文字、图片、链接和小程序
// 	支持设置分时段欢迎语
// 	欢迎语可用员工如与可用部门重复，以员工为准
// 	部门和员工表都有 welcome_msg_id，使用时先找员工的welcome_msg_id，若未找到，则找所属部门的welcome_msg_id，按部门级别从下往上找
// 	企微不支持ossURL，需将文件上传到指定位置再使用
// Param req requests.CreateWelcomeMsgReq 创建欢迎语请求
// Param extCorpID string 企业外部ID
// Param creator string 创建者的外部ID
// return msg models.WelcomeMsg 欢迎语model
func (o WelcomeMsgService) Create(req requests.CreateWelcomeMsgReq, extCorpID string, extCreatorID string) (msg models.WelcomeMsg, err error) {
	// 欢迎语主要内容
	mainMsgID := id_generator.StringID()
	mainMsg := models.WelcomeMsg{
		ExtCorpModel:        models.ExtCorpModel{ID: mainMsgID, ExtCorpID: extCorpID, ExtCreatorID: extCreatorID},
		WelcomeMsg:          req.WelcomeMsg,
		EnableTimePeriodMsg: req.EnableTimePeriodMsg,
		Name:                req.Name,
	}

	// 上传图片到微信
	err = o.createImageWxURL(&req, extCorpID)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	// 按微信要求限制附件
	err = o.checkAttachments(&req, extCorpID)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	// 分时段欢迎语
	if req.TimePeriodMsg != nil {
		var timePeriodMsgs []models.WelcomeMsg
		for _, msg := range req.TimePeriodMsg {
			welcomeMsg := models.WelcomeMsg{}
			err := copier.Copy(&welcomeMsg, msg)
			if err != nil {
				err = errors.WithStack(err)
				return mainMsg, err
			}
			welcomeMsg.ID = id_generator.StringID()
			welcomeMsg.MainWelcomeMsgID = &mainMsgID
			timePeriodMsgs = append(timePeriodMsgs, welcomeMsg)
		}
		mainMsg.TimePeriodMsg = timePeriodMsgs
	}

	// 更新员工可用欢迎语
	// 没有指定员工id和部门id，则全部部门可用
	if len(req.ExtDepartmentIDs) == 0 && len(req.ExtStaffIds) == 0 {
		req.ExtDepartmentIDs = []int64{1}
	}

	err = models.DB.Transaction(func(tx *gorm.DB) error {
		err = o.msgRepo.Create(tx, mainMsg)
		if err != nil {
			err = errors.WithStack(err)
			return err
		}

		if req.ExtDepartmentIDs != nil {
			err = o.departmentRepo.UpdateWelcomeMsg(tx, mainMsgID, extCorpID, req.ExtDepartmentIDs)
			if err != nil {
				err = errors.WithStack(err)
				return err
			}
		}

		if req.ExtStaffIds != nil {
			err = o.staffRepo.UpdateWelcomeMsg(tx, extCorpID, req.ExtStaffIds, mainMsgID)
			if err != nil {
				err = errors.WithStack(err)
				return err
			}
		}
		return nil
	})

	return mainMsg, err
}

func (o WelcomeMsgService) Update(ID string, req requests.UpdateWelcomeMsgReq, extCorpID string) (models.WelcomeMsg, error) {
	mainMsg := models.WelcomeMsg{
		ExtCorpModel:        models.ExtCorpModel{ID: ID, ExtCorpID: extCorpID},
		WelcomeMsg:          req.WelcomeMsg,
		EnableTimePeriodMsg: req.EnableTimePeriodMsg,
		Name:                req.Name,
	}
	if req.TimePeriodMsg != nil {
		var timePeriodMsgs []models.WelcomeMsg
		for _, msg := range req.TimePeriodMsg {
			welcomeMsg := models.WelcomeMsg{}
			err := copier.Copy(&welcomeMsg, msg)
			if err != nil {
				return mainMsg, err
			}
			if welcomeMsg.ID == "" {
				welcomeMsg.ID = id_generator.StringID()
			}
			welcomeMsg.MainWelcomeMsgID = &ID
			timePeriodMsgs = append(timePeriodMsgs, welcomeMsg)
		}
		mainMsg.TimePeriodMsg = timePeriodMsgs
	}
	err := models.DB.Transaction(func(tx *gorm.DB) error {
		if req.EnableTimePeriodMsg == constants.False {
			err := o.msgRepo.DeleteTimePeriodMsg(tx, ID)
			if err != nil {
				return err
			}
		}
		err := o.msgRepo.Update(tx, mainMsg)
		if err != nil {
			return err
		}

		if req.ExtDepartmentIDs != nil {
			err = o.departmentRepo.RemoveOriginalWelcomeMsg(tx, extCorpID, ID)
			if err != nil {
				return err
			}
			err = o.departmentRepo.UpdateWelcomeMsg(tx, ID, extCorpID, req.ExtDepartmentIDs)
			if err != nil {
				return err
			}
		}

		if req.ExtStaffIds != nil {
			err = o.staffRepo.RemoveOriginalWelcomeMsg(tx, ID)
			if err != nil {
				return err
			}
			err = o.staffRepo.UpdateWelcomeMsg(tx, extCorpID, req.ExtStaffIds, ID)
			if err != nil {
				return err
			}
		}
		return nil
	})

	return mainMsg, err
}

// Delete
// Description: 删除欢迎语
// Detail:
//	附件已经上传至微信，不用删除
// 	员工表中有欢迎语ID, 没有将其置空，取用时查不到已删除的欢迎语即可。
func (o WelcomeMsgService) Delete(ids []string, extCorpID string) error {
	return o.msgRepo.Delete(ids, extCorpID)
}

func (o WelcomeMsgService) Query(req requests.QueryWelcomeMsgReq, extCorpID string, sorter *app.Sorter, pager *app.Pager) (welcomeMsgWithStaffAndDept []models.WelcomeMsgWithDeptAndStaff, total int64, err error) {
	welcomeMsgWithStaffAndDept = make([]models.WelcomeMsgWithDeptAndStaff, 0)
	msgs, total, err := o.msgRepo.Query(req, extCorpID, sorter, pager)
	if err != nil {
		return nil, 0, err
	}

	log.Sugar.Debugw("msgs", "msgs", msgs)

	for _, msg := range msgs {
		staffs, err := o.staffRepo.GetMainInfoByMsgID(msg.ID)
		if err != nil {
			return nil, 0, err
		}

		depts, err := o.departmentRepo.GetMainInfoByMsgID(msg.ID)
		if err != nil {
			return nil, 0, err
		}

		welcomeMsgWithStaffAndDept = append(welcomeMsgWithStaffAndDept,
			models.WelcomeMsgWithDeptAndStaff{WelcomeMsg: msg, Department: depts, Staffs: staffs})
	}

	return welcomeMsgWithStaffAndDept, total, nil
}

func (o WelcomeMsgService) Get(id string, extCorpID string) (msg models.WelcomeMsgWithDeptAndStaff, err error) {
	welcomeMsg, err := o.msgRepo.Get(id, extCorpID)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	staffs, err := o.staffRepo.GetMainInfoByMsgID(id)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	departments, err := o.departmentRepo.GetMainInfoByMsgID(id)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	msg = models.WelcomeMsgWithDeptAndStaff{WelcomeMsg: welcomeMsg, Department: departments, Staffs: staffs}
	return
}

func (o WelcomeMsgService) UploadImg(body io.ReadCloser, filename, extCorpID string) (url string, err error) {

	data, err := ioutil.ReadAll(body)
	if err != nil {
		err = errors.Wrap(err, "ioutil.ReadAll failed")
		return
	}

	if len(data) == 0 {
		err = errors.WithStack(ecode.BadRequest)
		return
	}

	contentType := http.DetectContentType(data)
	n := strings.IndexByte(contentType, '/')
	if contentType[0:n] != "image" {
		err = ecode.NotImageFile
		return
	}

	media, err := work_wx.NewMediaFromBuffer(filename, data)
	if err != nil {
		return
	}

	client, err := we_work.Clients.Get(extCorpID)
	if err != nil {
		return
	}
	url, err = client.Customer.UploadPermanentImageMedia(media)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}

func (o WelcomeMsgService) createImageWxURL(req *requests.CreateWelcomeMsgReq, extCorpID string) (err error) {
	for i, _ := range req.WelcomeMsg.Attachments {
		if req.WelcomeMsg.Attachments[i].MsgType == string(constants.ImageMsgType) {
			uploadURL := req.WelcomeMsg.Attachments[i].Image.PicURL
			// 从oss下载，上传到wx
			var obj string
			obj, err = GetObjFromSignedURL(uploadURL)
			if err != nil {
				err = errors.WithStack(err)
				return
			}
			log.Sugar.Debugw("obj", "obj", obj)

			var readCloser io.ReadCloser
			readCloser, err = storage.FileStorage.Get(obj)
			if err != nil {
				return
			}

			var data []byte
			data, err = ioutil.ReadAll(readCloser)
			if err != nil {
				err = errors.Wrap(err, "ioutil.ReadAll failed")
				return
			}

			var media *work_wx.Media
			media, err = work_wx.NewMediaFromBuffer(obj, data)
			if err != nil {
				return
			}

			var client we_work.Client
			client, err = we_work.Clients.Get(extCorpID)
			if err != nil {
				return
			}

			var url string
			url, err = client.Customer.UploadPermanentImageMedia(media)
			if err != nil {
				err = errors.WithStack(err)
				return
			}
			req.WelcomeMsg.Attachments[i].Image.PicURL = url
		}
	}
	return
}

// checkAttachments
// Description: 检查attachments参数是否符合wx要求。
// Detail: wx 限制链接的描述不大于512，链接的标题不大于128字节
// Param: req *requests.CreateWelcomeMsgReq 创建欢迎语请求
// Param: extCorpID string 外部企业ID
// return
func (o WelcomeMsgService) checkAttachments(req *requests.CreateWelcomeMsgReq, extCorpID string) (err error) {
	if req == nil {
		return
	}
	for i, attachment := range req.WelcomeMsg.Attachments {
		switch attachment.MsgType {
		case string(constants.LinkMsgType):
			if len(attachment.Link.Desc) > 512 {
				req.WelcomeMsg.Attachments[i].Link.Desc = req.WelcomeMsg.Attachments[i].Link.Desc[:256]
				//req.WelcomeMsg.Attachments[i].Link.Desc = string([]byte(desc)[:256])
			}
			if len(attachment.Link.Title) > 128 {
				req.WelcomeMsg.Attachments[i].Link.Title = req.WelcomeMsg.Attachments[i].Link.Title[:128]
			}
		}
	}
	return
}
