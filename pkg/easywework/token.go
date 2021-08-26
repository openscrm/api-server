package workwx

import (
	"context"
	"fmt"
	"github.com/cenkalti/backoff/v4"
	"sync"
	"time"
)

type tokenInfo struct {
	token     string
	expiresIn time.Duration
}

type token struct {
	mutex *sync.RWMutex
	tokenInfo
	lastRefresh  time.Time
	getTokenFunc func() (tokenInfo, error)
}

// getAccessToken 获取 access token
func (c *App) getAccessToken() (tokenInfo, error) {
	//fmt.Println("fetch access_token")
	get, err := c.execGetAccessToken(accessTokenReq{
		CorpID:     c.CorpID,
		CorpSecret: c.CorpSecret,
	})
	if err != nil {
		return tokenInfo{}, err
	}
	return tokenInfo{token: get.AccessToken, expiresIn: time.Duration(get.ExpiresInSecs)}, nil
}

// SpawnAccessTokenRefresher 启动该 app 的 access token 刷新 goroutine
//
// NOTE: 该 goroutine 本身没有 keep-alive 逻辑，需要自助保活
func (c *App) SpawnAccessTokenRefresher() {
	ctx := context.Background()
	c.SpawnAccessTokenRefresherWithContext(ctx)
}

// SpawnAccessTokenRefresherWithContext 启动该 app 的 access token 刷新 goroutine
// 可以通过 context cancellation 停止此 goroutine
//
// NOTE: 该 goroutine 本身没有 keep-alive 逻辑，需要自助保活
func (c *App) SpawnAccessTokenRefresherWithContext(ctx context.Context) {
	go c.accessToken.tokenRefresher(ctx)
}

// GetJSAPITicket 获取 JSAPI_ticket
func (c *App) GetJSAPITicket() (string, error) {
	return c.jsapiTicket.getToken(), nil
}

// getJSAPITicket 获取 JSAPI_ticket
func (c *App) getJSAPITicket() (tokenInfo, error) {
	get, err := c.execGetJSAPITicket(jsAPITicketReq{})
	if err != nil {
		return tokenInfo{}, err
	}
	return tokenInfo{token: get.Ticket, expiresIn: time.Duration(get.ExpiresInSecs)}, nil
}

// GetJSAPIAgentTicket 获取 JSAPI_agent_ticket
func (c *App) GetJSAPIAgentTicket() (string, error) {
	return c.jsapiTicketAgentConfig.getToken(), nil
}

// getJSAPIAgentTicket 获取 JSAPI_agent_ticket
func (c *App) getJSAPIAgentTicket() (tokenInfo, error) {
	get, err := c.execGetJSAPITicketAgentConfig(jsAPITicketAgentConfigReq{})
	if err != nil {
		return tokenInfo{}, err
	}
	return tokenInfo{token: get.Ticket, expiresIn: time.Duration(get.ExpiresInSecs)}, nil
}

// SpawnJSAPITicketRefresher 启动该 app 的 JSAPI_ticket 刷新 goroutine
//
// NOTE: 该 goroutine 本身没有 keep-alive 逻辑，需要自助保活
func (c *App) SpawnJSAPITicketRefresher() {
	ctx := context.Background()
	c.SpawnJSAPITicketRefresherWithContext(ctx)
}

// SpawnJSAPITicketRefresherWithContext 启动该 app 的 JSAPI_ticket 刷新 goroutine
// 可以通过 context cancellation 停止此 goroutine
//
// NOTE: 该 goroutine 本身没有 keep-alive 逻辑，需要自助保活
func (c *App) SpawnJSAPITicketRefresherWithContext(ctx context.Context) {
	go c.jsapiTicket.tokenRefresher(ctx)
}

// GetJSAPITicketAgentConfig 获取 JSAPI_ticket_agent_config
func (c *App) GetJSAPITicketAgentConfig() (string, error) {
	return c.jsapiTicketAgentConfig.getToken(), nil
}

// getJSAPITicketAgentConfig 获取 JSAPI_ticket_agent_config
func (c *App) getJSAPITicketAgentConfig() (tokenInfo, error) {
	get, err := c.execGetJSAPITicketAgentConfig(jsAPITicketAgentConfigReq{})
	if err != nil {
		return tokenInfo{}, err
	}
	return tokenInfo{token: get.Ticket, expiresIn: time.Duration(get.ExpiresInSecs)}, nil
}

// SpawnJSAPITicketAgentConfigRefresher 启动该 app 的 JSAPI_ticket_agent_config 刷新 goroutine
//
// NOTE: 该 goroutine 本身没有 keep-alive 逻辑，需要自助保活
func (c *App) SpawnJSAPITicketAgentConfigRefresher() {
	ctx := context.Background()
	c.SpawnJSAPITicketAgentConfigRefresherWithContext(ctx)
}

// SpawnJSAPITicketAgentConfigRefresherWithContext 启动该 app 的 JSAPI_ticket_agent_config 刷新 goroutine
// 可以通过 context cancellation 停止此 goroutine
//
// NOTE: 该 goroutine 本身没有 keep-alive 逻辑，需要自助保活
func (c *App) SpawnJSAPITicketAgentConfigRefresherWithContext(ctx context.Context) {
	go c.jsapiTicketAgentConfig.tokenRefresher(ctx)
}

func (t *token) setGetTokenFunc(f func() (tokenInfo, error)) {
	t.getTokenFunc = f
}

func (t *token) getToken() string {
	// intensive mutex juggling action
	t.mutex.RLock()
	if t.token == "" {
		t.mutex.RUnlock() // RWMutex doesn't like recursive locking
		_ = t.syncToken()
		t.mutex.RLock()
	}
	tokenToUse := t.token
	t.mutex.RUnlock()
	return tokenToUse
}

func (t *token) syncToken() error {
	get, err := t.getTokenFunc()
	if err != nil {
		return err
	}
	t.mutex.Lock()
	defer t.mutex.Unlock()

	t.token = get.token
	t.expiresIn = get.expiresIn * time.Second
	t.lastRefresh = time.Now()
	return nil
}

func (t *token) tokenRefresher(ctx context.Context) {
	// refresh per 30m
	const refreshTimeWindow = (2*60 - 30) * time.Minute
	const minRefreshDuration = 5 * time.Second

	var waitDuration time.Duration = 0
	for {
		select {
		case <-time.After(waitDuration):
			retryer := backoff.WithContext(backoff.NewExponentialBackOff(), ctx)
			if err := backoff.Retry(t.syncToken, retryer); err != nil {
				fmt.Println("retry getting access toke failed", "err", err)
				_ = err
			}

			waitDuration = t.lastRefresh.Add(t.expiresIn).Add(-refreshTimeWindow).Sub(t.lastRefresh)
			fmt.Println("access_token", "token", t.token, "expiresIn", t.expiresIn, "nextRefreshTime", waitDuration)
			if waitDuration < minRefreshDuration {
				waitDuration = minRefreshDuration
			}
		case <-ctx.Done():
			return
		}
	}
}

// JSCode2Session 临时登录凭证校验
func (c *App) JSCode2Session(jscode string) (*JSCodeSession, error) {
	resp, err := c.execJSCode2Session(jsCode2SessionReq{JSCode: jscode})
	if err != nil {
		return nil, err
	}
	return &resp.JSCodeSession, nil
}
