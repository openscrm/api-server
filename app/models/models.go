package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"openscrm/app/constants"
	log "openscrm/common/log"
	"openscrm/conf"
	"os"
	"time"
)

var DB *gorm.DB // 默认数据库

type Timestamp struct {
	CreatedAt time.Time      `sql:"index" gorm:"comment:'创建时间'" json:"created_at"`
	UpdatedAt time.Time      `sql:"index" gorm:"comment:'更新时间'" json:"updated_at"`
	DeletedAt gorm.DeletedAt `sql:"index" gorm:"comment:'删除时间'" json:"deleted_at" swaggerignore:"true"`
}

type Model struct {
	ID string `gorm:"primaryKey;type:bigint AUTO_INCREMENT;comment:'ID'" json:"id" validate:"int64"`
}

type ExtCorpModel struct {
	// ID
	ID string `json:"id" gorm:"primaryKey;type:bigint;comment:'ID'" validate:"int64"`
	// ExtCorpID 外部企业ID
	ExtCorpID string `json:"ext_corp_id" gorm:"index;type:char(18);comment:外部企业ID" validate:ext_corp_id"`
	// ExtCreatorID 创建者外部员工ID
	ExtCreatorID string `json:"ext_creator_id" gorm:"index;type:char(32);comment:创建者外部员工ID" validate:"word"`
}

// RefModel 关联表基本模型，ID仅用做唯一键，使用组合字段作为主键，方便去重，可实现Association replace保留原纪录
type RefModel struct {
	// ID
	ID string `json:"id" gorm:"unique;type:bigint;comment:'ID'" validate:"int64"`
	// ExtCorpID 外部企业ID
	ExtCorpID string `json:"ext_corp_id" gorm:"index;type:char(18);comment:外部企业ID" validate:ext_corp_id"`
}

// SetupDB Setup 用于初始化数据库连接
func SetupDB() {
	DB = initDB(conf.Settings.DB)
	SetupPermissions()     // 初始化的权限
	SetupRoles()           // 初始化默认角色
	SetupStaffRole()       // 初始化超级管理员权限
	SetupContactWayGroup() // 初始化默认渠道码分组
	SetupCustomerInfoDisplayRule()
}

//AutoMigrate 用于根据结构体定义自动迁移表结构,只新增字段
func AutoMigrate(db *gorm.DB) error {
	if !conf.Settings.App.AutoMigration {
		return nil
	}

	db = db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci")
	err := db.AutoMigrate(
		// 此处顺序不可改动，以免处罚外键约束
		&Customer{},
		&CustomerInfo{},
		&CustomerInfoDisplayRule{},
		&CustomerStaff{},
		&CustomerStaffRelationHistory{},
		&Staff{},
		&Department{},
		&MaterialLibTag{},
		&StaffDepartment{},
		&CustomerStaffTag{},
		&GroupChatTagGroup{},
		&CorpSetting{},
		&GroupChatTag{},
		&TagGroup{},
		&Tag{},
		&ChatMsg{},
		&ChatMsgContent{},
		&QuickReplyGroup{},
		&QuickReply{},
		&QuickReplyDetail{},
		&CustomerRemark{},
		&RemarkOption{},
		&CustomerEvent{},
		&InternalTag{},
		&Material{},
		&Remainder{},
		&MassMsg{},
		&MassMsgStaff{},
		&WelcomeMsg{},
		&EventNotify{},
		&GroupChat{},
		&GroupChatMember{},
		&GroupChatGroup{},
		&GroupChatAutoJoinCode{},
		&GroupChatQRCode{},
		&GroupChatAutoJoinCodeStaff{},
		&GroupChatWelcomeMsg{},
		&GroupChatMassMsg{},
		&CustomerStatistic{},
		&DataExport{},
		&ContactWayGroup{},
		&ContactWay{},
		&ContactWaySchedule{},
		&ContactWayScheduleStaff{},
		&ContactWayBackupStaff{},
		&ContactWayStaff{},
		&Permission{},
		&CustomerInfoDisplayRule{},
		&Tag{},
		&TagGroup{},
		&Staff{},
		&Role{},
	)
	if err != nil {
		log.Sugar.Errorw(err.Error())
		return err
	}
	// 修改 Staff 的 Departments 字段的连接表为 StaffDepartment
	err = db.SetupJoinTable(&Staff{}, "Departments", &StaffDepartment{})
	if err != nil {
		return err
	}

	return nil
}

//initDB 初始化数据库连接
func initDB(c conf.DBConfig) (db *gorm.DB) {
	var err error

	gormLogLevel := logger.Warn
	if conf.Settings.App.Env == constants.DEV {
		gormLogLevel = logger.Info
	}

	db, err = gorm.Open(
		mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Asia%%2FShanghai",
			c.User,
			c.Password,
			c.Host,
			c.Name)),
		&gorm.Config{
			SkipDefaultTransaction:                   false,
			DisableForeignKeyConstraintWhenMigrating: true,
			Logger:                                   logger.Default.LogMode(gormLogLevel),
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
			}},
	)
	if err != nil {
		log.Sugar.Errorw("models.Setup failed", "err", err, "conf", c)
		os.Exit(1)
	}

	err = AutoMigrate(db)
	if err != nil {
		log.Sugar.Errorw("model auto migrate failed", "err", err, "conf", c)
		os.Exit(1)
	}

	return db
}
