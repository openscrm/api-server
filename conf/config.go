package conf

import (
	"github.com/spf13/viper"
	"log"
	"time"
)

var Settings *config

type config struct {
	App        AppConfig
	Server     serverConfig
	Redis      redisConfig
	DB         DBConfig
	DelayQueue delayQueueConfig
	Storage    StorageConfig
	WeWork     weWorkConfig
}

// delayQueueConfig  基于redis延迟队列的配置
type delayQueueConfig struct {
	// bucket数量
	BucketSize int `validate:"required"`
	// bucket在redis中的键名
	BucketName string `validate:"required"`
	// ready queue在redis中的键名
	QueueName string `validate:"required"`
	//调用blpop阻塞超时时间, 单位秒, 修改此项, redis.read_timeout必须做相应调整
	QueueBlockTimeout int `validate:"required"`
}

type AppConfig struct {
	Name               string `validate:"required"`
	Key                string `validate:"required,base64"` //应用秘钥 64位，生成命令：openssl rand 64 -base64
	Env                string `validate:"required,oneof=PROD DEV TEST"`
	AutoMigration      bool
	AutoSyncWeWorkData bool // 启动时同步微信数据
	// SuperAdminPhone 此处手机号对应员工的赋予超级管理员权限
	SuperAdminPhone []string `validate:"required,dive,phone"`
	InnerSrvAppCode string   // 内部服务调用key
}

type serverConfig struct {
	RunMode         string        `validate:"required,oneof=debug test release"`
	HttpPort        int           `validate:"required,gt=0"`
	ReadTimeout     time.Duration `validate:"required,gt=0"`
	WriteTimeout    time.Duration `validate:"required,gt=0"`
	MsgArchHttpPort int
	MsgArchSrvHost  string
}

// weWorkConfig 企业微信配置
type weWorkConfig struct {
	// ExtCorpID 外部企业ID
	ExtCorpID string `json:"ext_corp_id" validate:"required,corp_id"`
	// ContactSecret 通讯录secret
	ContactSecret string `json:"contact_secret" validate:"required"`
	// CustomerSecret 客户联系secret
	CustomerSecret string `json:"customer_secret" validate:"required"`
	// MainAgentID 企业主应用AgentID
	MainAgentID int64 `json:"main_agent_id" validate:"number,gt=0"`
	// MainAgentSecret 企业主应用secret
	MainAgentSecret string `json:"main_agent_secret" validate:"required"`
	// CallbackToken 企业微信事件回调Token
	CallbackToken string `json:"callback_token" validate:"required"`
	// CallbackAesKey 企业微信事件回调AesKey
	CallbackAesKey string `json:"callback_aes_key" validate:"required"`
	// PriKeyPath 会话存档解密私钥
	PriKeyPath string `json:"pri_key_path"`
	// MsgArchBatchSize 会话存档拉取，每次拉取的条数
	MsgArchBatchSize int `json:"msg_arch_batch_size"`
	// MsgArchTimeout 会话存档拉取，超时时间
	MsgArchTimeout int `json:"msg_arch_timeout"`
	// MsgArchProxy  会话存档拉取代理地址
	MsgArchProxy string `json:"msg_arch_proxy"`
	// MsgArchProxyPasswd  会话存档拉取代理密码
	MsgArchProxyPasswd string `json:"msg_arch_proxy_passwd"`
}

type DBConfig struct {
	User     string `validate:"required"`
	Password string `validate:"required"`
	Host     string `validate:"required"`
	Name     string `validate:"required"`
}

// redisConfig redis config
type redisConfig struct {
	Host        string        `validate:"required"`
	Password    string        `validate:"required"`
	IdleTimeout time.Duration `validate:"required"`
	DBNumber    int           `validate:"gte=0"`
	DialTimeout time.Duration `validate:"required"`
	ReadTimeout time.Duration `validate:"required"`
}

type StorageConfig struct {
	// Type 存储类型, 可配置aliyun, qcloud；分别对应阿里云OSS, 腾讯云COS
	Type string `validate:"required,oneof=aliyun qcloud"`
	// CdnURL CDN绑定域名，可选配置，本地存储必填
	CdnURL string `validate:"omitempty,url"`

	// 阿里云OSS相关配置，请使用子账户凭据，且仅授权oss访问权限
	AccessKeyId     string `validate:"required_if=Type aliyun"`
	AccessKeySecret string `validate:"required_if=Type aliyun"`
	EndPoint        string `validate:"required_if=Type aliyun"`
	Bucket          string `validate:"required_if=Type aliyun"`

	// 腾讯云OSS相关配置，请使用子账户凭据，且仅授权cos访问权限
	SecretID  string `validate:"required_if=Type qcloud"`
	SecretKey string `validate:"required_if=Type qcloud"`
	BucketURL string `validate:"required_if=Type qcloud"`

	// 本地存储相关配置
	// LocalRootPath 本地存储文件的根目录，必须是绝对路径
	LocalRootPath string `validate:"required_if=Type local"`
	// ServerRootPath 文件服务的根目录，http服务中的文件根目录，相对路径，用于识别文件服务请求的路径标识
	ServerRootPath string `validate:"required_if=Type local"`
}

// SetupSetting Setup initialize the configuration instance
func SetupSetting() error {
	var err error
	viper.SetConfigName("config")     // name of config file (without extension)
	viper.AddConfigPath("conf")       // optionally look for config in the working directory
	viper.AddConfigPath("../conf")    // optionally look for config in the working directory
	viper.AddConfigPath("../../conf") // optionally look for config in the working directory
	viper.AddConfigPath("/srv")       // optionally look for config in the working directory
	err = viper.ReadInConfig()        // Find and read the config file
	if err != nil {                   // Handle errors reading the config file
		log.Printf("missing config.yaml : %s\n", err.Error())
		return err
	}
	Settings = &config{}
	err = viper.Unmarshal(Settings)
	if err != nil {
		log.Printf("parse config.yaml failed : %s", err.Error())
		return err
	}
	viper.WatchConfig()
	Settings.Server.ReadTimeout = Settings.Server.ReadTimeout * time.Second
	Settings.Server.WriteTimeout = Settings.Server.WriteTimeout * time.Second

	Settings.Redis.DialTimeout = Settings.Redis.DialTimeout * time.Second
	Settings.Redis.IdleTimeout = Settings.Redis.IdleTimeout * time.Second
	Settings.Redis.ReadTimeout = Settings.Redis.ReadTimeout * time.Second
	return nil
}

//SetupTestSetting 初始化单元测试的配置
func SetupTestSetting() error {
	var err error
	viper.SetConfigName("config.test") // name of config file (without extension)
	viper.AddConfigPath("conf")        // optionally look for config in the working directory
	viper.AddConfigPath("../conf")     // optionally look for config in the working directory
	viper.AddConfigPath("../../conf")  // optionally look for config in the working directory
	err = viper.ReadInConfig()         // Find and read the config file
	if err != nil {                    // Handle errors reading the config file
		log.Printf("missing config.yaml : %s\n", err.Error())
		return err
	}
	Settings = &config{}
	err = viper.Unmarshal(Settings)
	if err != nil {
		log.Printf("parse config.yaml failed : %s", err.Error())
		return err
	}
	viper.WatchConfig()
	Settings.Server.ReadTimeout = Settings.Server.ReadTimeout * time.Second
	Settings.Server.WriteTimeout = Settings.Server.WriteTimeout * time.Second

	Settings.Redis.DialTimeout = Settings.Redis.DialTimeout * time.Second
	Settings.Redis.IdleTimeout = Settings.Redis.IdleTimeout * time.Second
	Settings.Redis.ReadTimeout = Settings.Redis.ReadTimeout * time.Second
	return nil
}
