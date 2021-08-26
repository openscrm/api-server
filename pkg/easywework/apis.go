package workwx

// execGetAccessToken 获取access_token
func (c *App) execGetAccessToken(req accessTokenReq) (accessTokenResp, error) {
	var resp accessTokenResp
	err := c.executeWXApiGet("/cgi-bin/gettoken", req, &resp, false)
	if err != nil {
		return accessTokenResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return accessTokenResp{}, bizErr
	}

	return resp, nil
}

// execGetJSAPITicket 获取企业的jsapi_ticket
func (c *App) execGetJSAPITicket(req jsAPITicketReq) (jsAPITicketResp, error) {
	var resp jsAPITicketResp
	err := c.executeWXApiGet("/cgi-bin/get_jsapi_ticket", req, &resp, true)
	if err != nil {
		return jsAPITicketResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return jsAPITicketResp{}, bizErr
	}

	return resp, nil
}

// execGetJSAPITicketAgentConfig 获取应用的jsapi_ticket
func (c *App) execGetJSAPITicketAgentConfig(req jsAPITicketAgentConfigReq) (jsAPITicketResp, error) {
	var resp jsAPITicketResp
	err := c.executeWXApiGet("/cgi-bin/ticket/get", req, &resp, true)
	if err != nil {
		return jsAPITicketResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return jsAPITicketResp{}, bizErr
	}

	return resp, nil
}

// execJSCode2Session 临时登录凭证校验code2Session
func (c *App) execJSCode2Session(req jsCode2SessionReq) (jsCode2SessionResp, error) {
	var resp jsCode2SessionResp
	err := c.executeWXApiGet("/cgi-bin/miniprogram/jscode2session", req, &resp, true)
	if err != nil {
		return jsCode2SessionResp{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return jsCode2SessionResp{}, bizErr
	}

	return resp, nil
}
