package apis

import (
    "context"
    "errors"
    "github.com/cenkalti/backoff/v4"
    "log"
    "sync"
    "time"
)

// access_token等token信息直接存储在内存中
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

/*** access_token 模块 -begin- ***/
func (c *ApiClient) GetToken() (token string, err error) {
    token = c.accessToken.getToken()
    if token == "" {
        err = errors.New("access_token获取失败，配置有误或没有在服务商后台设置IP白名单或企业取消授权")
        return
    }
    return
}

// 获取服务商的provider_access_token
func (c *ApiClient) getProviderAccessToken() (tokenInfo, error) {
    get, err := c.ExecGetProviderTokenService(ReqGetProviderTokenService{
        Corpid:         c.CorpId,
        ProviderSecret: c.CorpProviderSecret,
    })
    if err != nil {
        return tokenInfo{}, err
    }
    return tokenInfo{token: get.ProviderAccessToken, expiresIn: time.Duration(get.ExpiresIn) * time.Second}, nil
}

// 第三方应用的suite_access_token
func (c *ApiClient) getSuiteToken() (tokenInfo, error) {
    if c.AppSuiteTicket == "" {
        return tokenInfo{}, errors.New("服务商的suite_ticket缺失，app_suite_id:" + c.AppSuiteId)
    }
    get, err := c.ExecGetSuiteTokenService(ReqGetSuiteTokenService{
        SuiteID:     c.AppSuiteId,
        SuiteSecret: c.AppSuiteSecret,
        SuiteTicket: c.AppSuiteTicket,
    })
    if err != nil {
        return tokenInfo{}, err
    }
    return tokenInfo{token: get.SuiteAccessToken, expiresIn: time.Duration(get.ExpiresIn) * time.Second}, nil
}

// 获取授权企业access_token
func (c *ApiClient) getAuthCorpToken() (tokenInfo, error) {
    if c.CompanyPermanentCode == "" {
        return tokenInfo{}, errors.New("永久授权码不存在，corp_id:" + c.CorpId)
    }
    get, err := c.ExecGetCorpTokenService(ReqGetCorpTokenService{
        AuthCorpid:    c.CorpId,
        PermanentCode: c.CompanyPermanentCode,
    })
    if err != nil {
        return tokenInfo{}, err
    }
    return tokenInfo{token: get.AccessToken, expiresIn: time.Duration(get.ExpiresIn) * time.Second}, nil
}

// SpawnAccessTokenRefresher 启动该 app 的 access token 刷新 goroutine
//
// NOTE: 该 goroutine 本身没有 keep-alive 逻辑，需要自助保活
func (c *ApiClient) SpawnAccessTokenRefresher() {
    ctx := context.Background()
    c.SpawnAccessTokenRefresherWithContext(ctx)
}

// SpawnAccessTokenRefresherWithContext 启动该 app 的 access token 刷新 goroutine
// 可以通过 context cancellation 停止此 goroutine
//
// NOTE: 该 goroutine 本身没有 keep-alive 逻辑，需要自助保活
func (c *ApiClient) SpawnAccessTokenRefresherWithContext(ctx context.Context) {
    go c.accessToken.tokenRefresher(ctx)
}

/*** access_token 模块 -end- ***/

/*** jsapi_ticket 模块 -begin- ***/

// GetJSAPITicket 获取 JSAPI_ticket
func (c *ApiClient) GetJSAPITicket() (string, error) {
    return c.jsapiTicket.getToken(), nil
}

// getJSAPITicket 获取 JSAPI_ticket
func (c *ApiClient) getJSAPITicket() (tokenInfo, error) {
    get, err := c.ExecGetJSAPITicket(JsAPITicketReq{})
    if err != nil {
        return tokenInfo{}, err
    }
    return tokenInfo{token: get.Ticket, expiresIn: time.Duration(get.ExpiresInSecs) * time.Second}, nil
}

// SpawnJSAPITicketRefresher 启动该 app 的 JSAPI_ticket 刷新 goroutine
//
// NOTE: 该 goroutine 本身没有 keep-alive 逻辑，需要自助保活
func (c *ApiClient) SpawnJSAPITicketRefresher() {
    ctx := context.Background()
    c.SpawnJSAPITicketRefresherWithContext(ctx)
}

// SpawnJSAPITicketRefresherWithContext 启动该 app 的 JSAPI_ticket 刷新 goroutine
// 可以通过 context cancellation 停止此 goroutine
//
// NOTE: 该 goroutine 本身没有 keep-alive 逻辑，需要自助保活
func (c *ApiClient) SpawnJSAPITicketRefresherWithContext(ctx context.Context) {
    go c.jsapiTicket.tokenRefresher(ctx)
}

/*** jsapi_ticket 模块 -end- ***/

/*** jsapi_ticket_agent_config 模块 -begin- ***/

// GetJSAPITicketAgentConfig 获取 JSAPI_ticket_agent_config
func (c *ApiClient) GetJSAPITicketAgentConfig() (string, error) {
    return c.jsapiTicketAgentConfig.getToken(), nil
}

// getJSAPITicketAgentConfig 获取 JSAPI_ticket_agent_config
func (c *ApiClient) getJSAPITicketAgentConfig() (tokenInfo, error) {
    get, err := c.ExecGetJSAPITicketAgentConfig(JsAPITicketAgentConfigReq{})
    if err != nil {
        return tokenInfo{}, err
    }
    return tokenInfo{token: get.Ticket, expiresIn: time.Duration(get.ExpiresInSecs) * time.Second}, nil
}

// SpawnJSAPITicketAgentConfigRefresher 启动该 app 的 JSAPI_ticket_agent_config 刷新 goroutine
//
// NOTE: 该 goroutine 本身没有 keep-alive 逻辑，需要自助保活
func (c *ApiClient) SpawnJSAPITicketAgentConfigRefresher() {
    ctx := context.Background()
    c.SpawnJSAPITicketAgentConfigRefresherWithContext(ctx)
}

// SpawnJSAPITicketAgentConfigRefresherWithContext 启动该 app 的 JSAPI_ticket_agent_config 刷新 goroutine
// 可以通过 context cancellation 停止此 goroutine
//
// NOTE: 该 goroutine 本身没有 keep-alive 逻辑，需要自助保活
func (c *ApiClient) SpawnJSAPITicketAgentConfigRefresherWithContext(ctx context.Context) {
    go c.jsapiTicketAgentConfig.tokenRefresher(ctx)
}

/*** jsapi_ticket_agent_config 模块 -end- ***/

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
    t.expiresIn = get.expiresIn
    t.lastRefresh = time.Now()
    return nil
}

// 每60分钟刷新一次token（每次获取token时，该token都有2小时有效期）
func (t *token) tokenRefresher(ctx context.Context) {
    // refresh per 30m
    const refreshTimeWindow = 60 * time.Minute
    const minRefreshDuration = 5 * time.Second

    var nextRefreshDuration time.Duration = 0
    for {
        select {
        case <-time.After(nextRefreshDuration):
            retryer := backoff.WithContext(backoff.NewExponentialBackOff(), ctx)
            if err := backoff.Retry(t.syncToken, retryer); err != nil {
                log.Println("retry getting access token failed", "err", err)
                _ = err
            }

            nextRefreshDuration = t.lastRefresh.Add(t.expiresIn - refreshTimeWindow).Sub(t.lastRefresh)
            if nextRefreshDuration < minRefreshDuration {
                nextRefreshDuration = minRefreshDuration
            }
        case <-ctx.Done():
            return
        }
    }
}