package services

import (
	"fmt"
	"github.com/gogf/gf/crypto/gsha1"
	"github.com/gogf/gf/util/grand"
	"github.com/pkg/errors"
	"openscrm/app/responses"
	log2 "openscrm/common/log"
	"openscrm/common/we_work"
	"openscrm/conf"
	"strings"
	"time"
)

type JsApiService struct {
}

func NewJsApiService() *JsApiService {
	return &JsApiService{}
}

// GetJSConfig 获取企业级别的JS SDK config参数
func (o JsApiService) GetJSConfig(url string, extCorpID string) (res responses.GetJSConfigResult, err error) {
	res = responses.GetJSConfigResult{
		Timestamp: time.Now().Unix(),
		NonceStr:  grand.Letters(20),
		AppID:     conf.Settings.WeWork.ExtCorpID,
		URL:       strings.Split(url, "#")[0],
	}

	client, err := we_work.Clients.Get(extCorpID)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	ticket, err := client.Customer.GetJSAPITicket()
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	rawStr := fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%d&url=%s", ticket, res.NonceStr, res.Timestamp, res.URL)
	log2.Sugar.Debugw("rawStr", "rawStr", rawStr)
	res.Signature = gsha1.Encrypt(rawStr)

	return
}

// GetJSAgentConfig 获取应用级别的JS SDK agentConfig参数
func (o JsApiService) GetJSAgentConfig(url string, extCorpID string) (res responses.GetJSAgentConfigResult, err error) {
	res = responses.GetJSAgentConfigResult{
		CorpID:    conf.Settings.WeWork.ExtCorpID,
		AgentID:   conf.Settings.WeWork.MainAgentID,
		Timestamp: time.Now().Unix(),
		NonceStr:  grand.Letters(20),
		URL:       strings.Split(url, "#")[0],
	}

	client, err := we_work.Clients.Get(extCorpID)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	ticket, err := client.MainApp.GetJSAPIAgentTicket()
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	rawStr := fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%d&url=%s", ticket, res.NonceStr, res.Timestamp, res.URL)
	log2.Sugar.Debugw("rawStr", "rawStr", rawStr)
	res.Signature = gsha1.Encrypt(rawStr)

	return
}
