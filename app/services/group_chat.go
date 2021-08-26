package services

import (
	"bytes"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/gogf/gf/os/gfile"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"openscrm/app/constants"
	"openscrm/app/models"
	"openscrm/app/requests"
	"openscrm/common/app"
	"openscrm/common/ecode"
	"openscrm/common/id_generator"
	"openscrm/common/log"
	"openscrm/common/we_work"
	"openscrm/conf"
	gowx "openscrm/pkg/easywework"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type GroupChatService struct {
	GroupChatRepo models.GroupChat
	StaffRepo     models.Staff
	customerRepo  models.Customer
}

func NewGroupChatService() *GroupChatService {
	return &GroupChatService{
		GroupChatRepo: models.GroupChat{},
		StaffRepo:     models.Staff{},
		customerRepo:  models.Customer{},
	}
}

// Syncs 同步群聊列表
func (o GroupChatService) Syncs(chatID string, extCorpID string) error {

	log.Sugar.Debugw("sync group chatInfo ", "chatID", chatID, "extCorpID", extCorpID)
	client, err := we_work.Clients.Get(extCorpID)
	if err != nil {
		err = errors.WithStack(err)
		return err
	}

	groupChatInfo, err := client.Customer.GetGroupChat(gowx.GetGroupChatReq{ChatId: chatID})
	if err != nil {
		err = errors.WithStack(err)
		return err
	}

	// 拼接群名
	chatInfo := groupChatInfo.GroupChat
	chatName, err := o.CreateChatName(chatInfo, extCorpID)
	if err != nil {
		err = errors.WithStack(err)
		return err
	}

	// 群主信息
	owner, err := o.StaffRepo.Get(chatInfo.Owner, extCorpID, false)
	if err != nil {
		err = errors.WithStack(err)
		return err
	}

	groupChat := models.GroupChat{
		ExtCorpModel: models.ExtCorpModel{ID: id_generator.StringID(), ExtCorpID: extCorpID, ExtCreatorID: chatInfo.Owner},
		ExtChatID:    chatID,
		Name:         chatName,
		Owner:        chatInfo.Owner,
		OwnerName:    owner.Name,
		CreateTime:   time.Unix(int64(chatInfo.CreateTime), 0),
		Notice:       chatInfo.Notice,
		Status:       constants.GroupChatStatusNotDismissed,
	}

	// 管理员
	for _, s := range chatInfo.AdminList {
		groupChat.AdminList = append(groupChat.AdminList, s.Userid)
	}

	// 群成员
	groupChat.MemberList = make([]models.GroupChatMember, 0)
	for _, m := range chatInfo.MemberList {
		member := models.GroupChatMember{}
		err = copier.Copy(&member, m)
		if err != nil {
			return err
		}
		member.Invitor = m.Invitor.Userid
		member.ID = id_generator.StringID()
		member.ExtCorpID = extCorpID
		groupChat.MemberList = append(groupChat.MemberList, member)
	}

	return o.GroupChatRepo.Upsert(groupChat)
}

// Query
// Description: 查询客户群列表
func (o GroupChatService) Query(
	req requests.QueryGroupChatReq, extCorpID string, pager *app.Pager, sorter *app.Sorter) (res []models.GroupChat, total int64, err error) {
	res = make([]models.GroupChat, 0)
	res, total, err = o.GroupChatRepo.Query(req, extCorpID, pager, sorter)
	return
}

// SyncAll 同步企业所有群聊
func (o GroupChatService) SyncAll(extCorpID string) error {
	client, err := we_work.Clients.Get(extCorpID)
	if err != nil {
		err = errors.WithStack(err)
		return err
	}
	pager := &app.Pager{}
	staffs, n, err := o.StaffRepo.Query(models.Staff{}, extCorpID, &app.Sorter{}, pager)
	if err != nil {
		err = errors.WithStack(err)
		return err
	}

	for n > 0 && len(staffs) > 0 {
		for _, staff := range staffs {
			req := gowx.ListGroupChatReq{
				Limit: 100,
				OwnerFilter: struct {
					UseridList []string `json:"userid_list"`
				}{
					UseridList: []string{staff.ExtID},
				},
			}
			groupChat, err := client.Customer.ListGroupChat(req)
			if err != nil && errors.Is(err, ecode.ErrCode40003) {
				err = errors.WithStack(err)
				return err
			}
			for _, chat := range groupChat.GroupChatList {
				err = o.Syncs(chat.ChatID, extCorpID)
				if err != nil {
					err = errors.WithStack(err)
					return err
				}
			}
		}

		pager.Page += 1
		staffs, n, err = o.StaffRepo.Query(models.Staff{}, extCorpID, &app.Sorter{}, pager)
		if err != nil {
			err = errors.WithStack(err)
			return err
		}
	}
	return nil
}

// UpdateTags
// Description: 群聊批量打标签
func (o GroupChatService) UpdateTags(req requests.UpdateCustomerGroupReq) (err error) {
	for _, groupChatID := range req.GroupChatIDs {
		err = o.GroupChatRepo.UpdateTags(groupChatID, req.AddTagIDs, req.RemoveTagIDs)
		if err != nil {
			err = errors.WithStack(err)
			return err
		}
	}
	return
}

func (o GroupChatService) Export(req requests.QueryGroupChatReq, extCorpID string) (*bytes.Buffer, string, error) {
	log.Sugar.Infow("ExportDeleteStaffWarningList", "req", req)
	groupChats, err := o.GroupChatRepo.GetAll(req, extCorpID)
	if err != nil {
		log.Sugar.Errorw("groupChatRepo.GetAll", "err", err)
		return nil, "", err
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

	err = PrettifySheet(constants.DataExportGroupChatListSheetName, file, exportTime, titles)
	if err != nil {
		log.Sugar.Error(err)
		return nil, "", err
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
			return nil, "", err
		}
	}

	buf, err := file.WriteToBuffer()
	if err != nil {
		return nil, "", err
	}

	return buf, fullPath, nil
}

func (o GroupChatService) GetAllOwners(extCorpId string) (staffs []models.StaffMainInfo, err error) {
	return o.GroupChatRepo.GetAllOwners(extCorpId)
}

// CreateChatName
// Description: 没有群名的群聊拼接群名
// Detail: 群主加两个管理员,管理员不足两个,就用群成员补齐
func (o GroupChatService) CreateChatName(chatInfo gowx.GroupChat, extCorpID string) (chatName string, err error) {
	sb := strings.Builder{}
	if chatInfo.Name == "" {
		staff, err := o.StaffRepo.Get(chatInfo.Owner, extCorpID, false)
		if err != nil {
			err = errors.WithStack(err)
			return chatName, err
		}
		sb.WriteString(staff.Name)

		// 最多再拼接两个管理员名字
		const NameLenLimit = 2
		var k int = 0
		for i := 0; i < len(chatInfo.AdminList) && k < NameLenLimit; i, k = i+1, k+1 {
			staff, err := o.StaffRepo.Get(chatInfo.Owner, extCorpID, false)
			if err != nil {
				err = errors.WithStack(err)
				return chatName, err
			}
			sb.WriteString(",")
			sb.WriteString(staff.Name)
		}
		// 还需要添加一个群成员名字
		for j := 0; j < len(chatInfo.MemberList) && k < NameLenLimit; j, k = j+1, k+1 {
			customer, err := o.customerRepo.Get(chatInfo.MemberList[j].Userid, extCorpID, false)
			if err != nil {
				err = errors.WithStack(err)
				return chatName, err
			}
			sb.WriteString(",")
			sb.WriteString(customer.Name)
		}
	} else {
		sb.WriteString(chatInfo.Name)
	}
	chatName = sb.String()
	return
}

func (o GroupChatService) GetAll(
	req requests.QueryGroupChatReq, extCorpID string) (items []models.GroupChatMainInfo, total int64, err error) {

	return o.GroupChatRepo.GetAllChatsMainInfo(&req.Pager, &req.Sorter, extCorpID)
}
