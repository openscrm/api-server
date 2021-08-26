package test

import (
	"github.com/gin-gonic/gin/binding"
	"log"
	"openscrm/app/consumers"
	"openscrm/app/models"
	"openscrm/common/delay_queue"
	"openscrm/common/id_generator"
	log2 "openscrm/common/log"
	"openscrm/common/redis"
	"openscrm/common/session"
	"openscrm/common/storage"
	"openscrm/common/validator"
	"openscrm/common/we_work"
	"openscrm/conf"
)

func init() {
	err := conf.SetupTestSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	setupValidator()
	//conf.Validate(conf.Settings)
	log2.SetupLogger(conf.Settings.App.Env)
	id_generator.SetupIDGenerator()
	models.SetupDB()
	storage.Setup(conf.Settings.Storage)
	redis.Setup(conf.Settings.Redis.Host, conf.Settings.Redis.Password, conf.Settings.Redis.DBNumber)
	session.Setup(conf.Settings.Redis.Host, conf.Settings.Redis.Password, conf.Settings.App.Key)
	we_work.SetupWXCallback()
	delay_queue.SetupDelayQueue()
	consumers.Start()
	models.SetupIDConverter()
}

func setupValidator() {
	binding.Validator = validator.NewCustomValidator()
}
