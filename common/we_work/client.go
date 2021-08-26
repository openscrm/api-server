package we_work

import (
	"github.com/pkg/errors"
	"openscrm/common/log"
	"openscrm/pkg"
	"openscrm/pkg/easywework"
	"sync"
)

var App *workwx.App

var (
	Clients *WeWorkClients
	once    sync.Once
)

type CorpConf struct {
	// ExtCorpID 外部企业ID
	ExtCorpID string `json:"ext_corp_id" validate:"ext_corp_id" gorm:"unique;type:char(18);comment:外部企业ID"`
	// ContactSecret 通讯录secret
	ContactSecret string `json:"contact_secret" gorm:"comment:'通讯录secret'"`
	// CustomerSecret 客户联系secret
	CustomerSecret string `json:"customer_secret" gorm:"comment:'客户联系secret'"`
	// MainAgentID 企业主应用AgentID
	MainAgentID int64 `json:"main_agent_id" gorm:"comment:'企业主应用AgentID'" validate:"number,gte=0"`
	// MainAgentSecret 企业主应用secret
	MainAgentSecret string `json:"main_agent_secret" gorm:"comment:'企业主应用secret'"`
}

type Client struct {
	MainApp  *workwx.App
	Contact  *workwx.App
	Customer *workwx.App
}

type WeWorkClients struct {
	mutex sync.Mutex
	//todo concurrent.Map
	//clients map[string]*gowx.WorkwxApp
	clients map[string]Client
}

// SetupClient 在连接db后init
func SetupClient(conf CorpConf) {
	once.Do(func() {
		var err error
		client, err := NewClient(conf)
		if err != nil {
			log.Sugar.Error("NewClient failed", err)
			//panic(err)
		}
		Clients = &WeWorkClients{mutex: sync.Mutex{}, clients: map[string]Client{conf.ExtCorpID: client}}
	})
}

func NewClient(conf CorpConf) (client Client, err error) {
	client.Contact, err = NewWxApp(conf.ExtCorpID, conf.ContactSecret, 0)
	if err != nil {
		err = errors.Wrap(err, "new Contact app failed")
		return
	}

	client.Customer, err = NewWxApp(conf.ExtCorpID, conf.CustomerSecret, 0)
	if err != nil {
		err = errors.Wrap(err, "new Customer app failed")
		return
	}

	client.MainApp, err = NewWxApp(conf.ExtCorpID, conf.MainAgentSecret, conf.MainAgentID)
	if err != nil {
		err = errors.Wrap(err, "new main app failed")
		return
	}

	return
}

func (clients *WeWorkClients) Get(extCorpID string) (client Client, err error) {
	var ok bool
	clients.mutex.Lock()
	defer clients.mutex.Unlock()
	client, ok = clients.clients[extCorpID]
	if !ok {
		err = errors.New("invalid extCorpID")
		return
	}

	return
}

func (clients *WeWorkClients) Remove(extCorpID string) {
	clients.mutex.Lock()
	defer clients.mutex.Unlock()

	delete(clients.clients, extCorpID)
}

func (clients *WeWorkClients) Add(conf CorpConf) (err error) {
	clients.mutex.Lock()
	defer clients.mutex.Unlock()
	client, err := NewClient(conf)
	if err != nil {
		err = errors.Wrap(err, "NewClient failed")
		return
	}
	Clients.clients[conf.ExtCorpID] = client

	return
}

func NewWxApp(extCorpID string, secret string, agentID int64) (wxApp *workwx.App, err error) {
	cliOptions := pkg.CliOptions{
		CorpID:            extCorpID,
		CorpSecret:        secret,
		AgentID:           agentID,
		QYAPIHostOverride: "",
		TLSKeyLogFile:     "",
	}

	wxApp = cliOptions.MakeWorkwxApp().WithApp(cliOptions.CorpSecret, cliOptions.AgentID)
	//_, err = wxApp.GetToken()
	//if err != nil {
	//	err = errors.Wrap(err, "get token failed")
	//	return
	//}

	return
}
