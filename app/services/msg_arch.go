package services

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"net/http"
	"openscrm/app/constants"
	"openscrm/app/requests"
	"openscrm/common/ecode"
	log2 "openscrm/common/log"
	"openscrm/common/util"
	"openscrm/conf"
	"path"
	"strconv"
)

type MsgArch struct {
	httpClient *resty.Client
}

// QuerySessions
// Description: 查会话列表
func (a *MsgArch) QuerySessions(req requests.QuerySessionReq, extCorpID string) (res interface{}, err error) {
	h := hmac.New(sha256.New, []byte(conf.Settings.App.InnerSrvAppCode))
	bytes, err := util.GenBytesOrderByColumn(req)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	h.Write(bytes)
	signedStr := hex.EncodeToString(h.Sum(nil))
	reqParams := map[string]string{
		"name":         req.Name,
		"ext_corp_id":  extCorpID,
		"ext_staff_id": req.ExtStaffID,
		"session_type": req.SessionType,
		"sort_field":   string(req.SortField),
		"sort_type":    string(req.SortType),
		"page":         strconv.Itoa(req.Page),
		"page_size":    strconv.Itoa(req.PageSize),
		"signature":    signedStr,
	}

	var result constants.JsonResult
	resp, err := a.httpClient.R().
		SetQueryParams(reqParams).
		SetResult(&result).
		SetHeader("Accept", "application/json").
		Get(path.Join("/api/v1", constants.MsgArchSrvPathSessions))
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	if resp.StatusCode() != http.StatusOK {
		err = ecode.InternalError
		return
	}

	log2.Sugar.Debugw("resp", "resp", result.Data)

	return result.Data, nil
}

// Sync
// Description: 同步存档会话数据
func (a *MsgArch) Sync(extCorpID string) error {
	h := hmac.New(sha256.New, []byte(conf.Settings.App.InnerSrvAppCode))
	h.Write([]byte(extCorpID))

	body := requests.SyncReq{
		ExtCorpID: extCorpID,
		Signature: hex.EncodeToString(h.Sum(nil)),
	}
	log2.Sugar.Debugf("signature: %v", body.Signature)
	reqBytes, err := json.Marshal(body)
	if err != nil {
		return err
	}

	resp, err := a.httpClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(reqBytes).
		Post(path.Join("/api/v1", constants.MsgArchSrvPathSync))
	if err != nil {
		err = errors.WithStack(err)
		return err
	}
	if resp.StatusCode() != http.StatusOK {
		err = ecode.InternalError
		return err

	}
	return err
}

// QueryMsgs
// Description: 查消息列表
func (a *MsgArch) QueryMsgs(req requests.QueryChatMsgReq, extCorpID string) (res interface{}, err error) {
	key := []byte(conf.Settings.App.InnerSrvAppCode)
	message, err := util.GenBytesOrderByColumn(req)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	MAC := mac.Sum(nil)
	signedStr := hex.EncodeToString(MAC)

	limit := ""
	if req.Limit > 0 {
		limit = strconv.Itoa(req.Limit)
	}
	reqParams := map[string]string{
		"ext_corp_id":   extCorpID,
		"msg_type":      req.MsgType,
		"receiver_id":   req.ReceiverID,
		"ext_staff_id":  req.ExtStaffID,
		"sort_field":    string(req.SortField),
		"sort_type":     string(req.SortType),
		"page":          strconv.Itoa(req.Page),
		"page_size":     strconv.Itoa(req.PageSize),
		"send_at_start": string(req.SendAtStart),
		"send_at_end":   string(req.SendAtEnd),
		"min_id":        req.MinID,
		"max_id":        req.MaxID,
		"limit":         limit,
		"signature":     signedStr,
	}

	if req.SendAtStart != "" {
		reqParams["send_at_start"] = string(req.SendAtStart)
	}

	if req.SendAtEnd != "" {
		reqParams["send_at_end"] = string(req.SendAtEnd)
	}

	var result constants.JsonResult
	resp, err := a.httpClient.R().
		SetQueryParams(reqParams).
		SetHeader("Accept", "application/json").
		SetResult(&result).
		Get(path.Join("/api/v1", constants.MsgArchSrvPathMsgs))
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	if resp.StatusCode() != http.StatusOK {
		err = ecode.InternalError
		return
	}

	log2.Sugar.Debugw("resp", "resp", result.Data)

	return result.Data, nil
}

// SearchMsgs
// Description: 关键字搜索会话
func (a *MsgArch) SearchMsgs(req requests.SearchMsgReq, extCorpID string) (res interface{}, err error) {
	key := []byte(conf.Settings.App.InnerSrvAppCode)
	message, err := util.GenBytesOrderByColumn(req)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	MAC := mac.Sum(nil)
	signedStr := hex.EncodeToString(MAC)
	reqParams := map[string]string{
		"ext_corp_id":  extCorpID,
		"ext_peer_id":  req.ExtPeerID,
		"ext_staff_id": req.ExtStaffID,
		"keyword":      req.Keyword,
		"page":         strconv.Itoa(req.Page),
		"page_size":    strconv.Itoa(req.PageSize),
		"signature":    signedStr,
	}

	var result constants.JsonResult
	resp, err := a.httpClient.R().
		SetQueryParams(reqParams).
		SetHeader("Accept", "application/json").
		SetResult(&result).
		Get(path.Join("/api/v1", constants.MsgArchSrvSearchMsgs))
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	if resp.StatusCode() != http.StatusOK {
		err = ecode.InternalError
		return
	}

	log2.Sugar.Debugw("resp", "resp", result.Data)

	return result.Data, nil
}

func NewMsgArch() *MsgArch {
	hostUrl := fmt.Sprintf("http://%s:%d", conf.Settings.Server.MsgArchSrvHost, conf.Settings.Server.MsgArchHttpPort)
	return &MsgArch{httpClient: resty.New().SetHostURL(hostUrl)}
}
