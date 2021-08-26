package services

import (
	"fmt"
	"github.com/gogf/gf/encoding/gurl"
	"github.com/pkg/errors"
	"net/url"
	"openscrm/app/entities"
	"openscrm/app/models"
	"openscrm/common/ecode"
	"openscrm/common/we_work"
	"openscrm/conf"
)

type Login struct {
	staffModel models.Staff
}

func NewDefaultLogin() *Login {
	return &Login{
		staffModel: models.Staff{},
	}
}

func NewLogin(staff models.Staff) *Login {
	return &Login{
		staffModel: staff,
	}
}

func (o *Login) StaffAdminLogin(extCorpID string, state string, sourceURL string) (item entities.StaffAdminLoginResp, err error) {

	sourceURL, err = gurl.Decode(sourceURL)
	if err != nil {
		err = errors.Wrap(err, "gurl.Decode failed")
		return
	}

	sourceURLInfo, err := url.Parse(sourceURL)
	if err != nil {
		err = errors.Wrap(err, "url.Parse failed")
		return
	}

	item = entities.StaffAdminLoginResp{
		AppID:   extCorpID,
		AgentID: conf.Settings.WeWork.MainAgentID,
		RedirectURI: fmt.Sprintf(
			"%s://%s/api/v1/staff-admin/action/login-callback",
			sourceURLInfo.Scheme,
			sourceURLInfo.Host,
		),
		State: state,
	}

	item.LocationURL = fmt.Sprintf(
		"https://open.work.weixin.qq.com/wwopen/sso/qrConnect?appid=%s&agentid=%d&redirect_uri=%s&state=%s",
		item.AppID, item.AgentID, gurl.Encode(item.RedirectURI), item.State,
	)

	return
}

func (o *Login) StaffAdminLoginCallback(extCorpID string, code string) (item *models.Staff, err error) {
	client, err := we_work.Clients.Get(conf.Settings.WeWork.ExtCorpID)
	if err != nil {
		err = errors.Wrap(err, "Get Client failed")
		return
	}

	userInfo, err := client.MainApp.GetUserInfoByCode(code)
	if err != nil {
		err = errors.Wrap(err, "GetUserInfoByCode failed")
		return
	}

	if userInfo.UserID == "" {
		err = errors.WithStack(ecode.ForbiddenError)
		return
	}

	item, err = (models.Staff{}).Get(userInfo.UserID, "", false)
	if err != nil {
		err = errors.Wrap(err, "GetStaffByUserID failed")
		return
	}

	if item.ExtCorpID != extCorpID {
		err = errors.WithStack(ecode.ForbiddenError)
		return
	}

	return
}

func (o *Login) StaffLogin(extCorpID string, state string, sourceURL string) (item entities.StaffLoginResp, err error) {
	sourceURL, err = gurl.Decode(sourceURL)
	if err != nil {
		err = errors.Wrap(err, "gurl.Decode failed")
		return
	}

	sourceURLInfo, err := url.Parse(sourceURL)
	if err != nil {
		err = errors.Wrap(err, "url.Parse failed")
		return
	}

	item = entities.StaffLoginResp{
		AppID: extCorpID,
		RedirectURI: fmt.Sprintf(
			"%s://%s/api/v1/staff-frontend/action/login-callback?appid=%s&source_url=%s",
			sourceURLInfo.Scheme,
			sourceURLInfo.Host,
			extCorpID,
			gurl.Encode(sourceURL),
		),
		SourceURL: sourceURL,
		State:     state,
	}

	item.LocationURL = fmt.Sprintf(
		"https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=snsapi_base&state=%s#wechat_redirect",
		item.AppID, gurl.Encode(item.RedirectURI), item.State,
	)

	return
}

func (o *Login) StaffLoginCallback(extCorpID string, code string) (item *models.Staff, err error) {
	client, err := we_work.Clients.Get(conf.Settings.WeWork.ExtCorpID)
	if err != nil {
		err = errors.Wrap(err, "Get Client failed")
		return
	}

	userInfo, err := client.MainApp.GetUserInfoByCode(code)
	if err != nil {
		err = errors.Wrap(err, "GetUserInfoByCode failed")
		return
	}

	if userInfo.UserID == "" {
		err = ecode.ForbiddenError
		return
	}

	item, err = (models.Staff{}).Get(userInfo.UserID, "", false)
	if err != nil {
		err = errors.Wrap(err, "GetStaffByUserID failed")
		return
	}

	if item.ExtCorpID != extCorpID {
		err = ecode.ForbiddenError
		return
	}

	return
}

//func (o *Login) CustomerLogin(extCorpID string, state string, sourceURL string) (item entities.CustomerLoginResp, err error) {
//
//	sourceURL, err = gurl.Decode(sourceURL)
//	if err != nil {
//		err = errors.Wrap(err, "gurl.Decode failed")
//		return
//	}
//	sourceURL = gurl.Encode(sourceURL)
//
//	item = entities.CustomerLoginResp{
//		AppID:       extCorpID,
//		RedirectURI: fmt.Sprintf("%s/api/v1/customer-frontend/action/login-callback?appid=%s&source_url=%s", conf.Settings.App.URL, extCorpID, sourceURL),
//		SourceURL:   sourceURL,
//		State:       state,
//	}
//
//	item.LocationURL = fmt.Sprintf(
//		"https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=snsapi_base&state=%s#wechat_redirect",
//		item.AppID, gurl.Encode(item.RedirectURI), item.State,
//	)
//
//	return
//}

func (o *Login) CustomerLoginCallback(extCorpID string, code string) (item *models.Customer, err error) {
	client, err := we_work.Clients.Get(conf.Settings.WeWork.ExtCorpID)
	if err != nil {
		err = errors.Wrap(err, "Get Client failed")
		return
	}

	userInfo, err := client.MainApp.GetUserInfoByCode(code)
	if err != nil {
		err = errors.Wrap(err, "GetUserInfoByCode failed")
		return
	}

	if userInfo.UserID == "" {
		err = ecode.ForbiddenError
		return
	}

	item, err = (models.Customer{}).Get(userInfo.ExternalUserID, conf.Settings.WeWork.ExtCorpID, false)
	if err != nil {
		err = errors.Wrap(err, "Get Customer failed")
		return
	}

	if item.ExtCorpID != extCorpID {
		err = ecode.ForbiddenError
		return
	}

	return
}
