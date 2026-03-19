package wxshop_go_sdk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// ============================================================================
// access_token 相关接口
// ============================================================================

// GetAccessTokenDirect 直接获取 access_token（不缓存）
func (c *Client) GetAccessTokenDirect(ctx context.Context) (*AccTokenRes, error) {
	params := url.Values{}
	params.Set("grant_type", "client_credential")
	params.Set("appid", c.AppID)
	params.Set("secret", c.AppSecret)

	reqURL := c.Env + AccessTokenApi + "?" + params.Encode()

	resp, err := http.Get(reqURL)
	if err != nil {
		return nil, fmt.Errorf("请求 access_token 失败：%w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败：%w", err)
	}

	var result AccTokenRes
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败：%w", err)
	}

	return &result, nil
}

// GetStableAccessTokenDirect 直接获取稳定的 access_token
func (c *Client) GetStableAccessTokenDirect(ctx context.Context, forceRefresh bool) (*AccTokenRes, error) {
	payload := map[string]any{
		"grant_type": "client_credential",
		"appid":      c.AppID,
		"secret":     c.AppSecret,
	}
	if forceRefresh {
		payload["force_refresh"] = true
	}

	bodyBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("序列化请求参数失败：%w", err)
	}

	reqURL := c.Env + AccessStableTokenApi
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("构建 stable access_token 请求失败：%w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求 stable access_token 失败：%w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败：%w", err)
	}

	var result AccTokenRes
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败：%w", err)
	}

	return &result, nil
}

// AccTokenRes access_token 接口响应
type AccTokenRes struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

// ============================================================================
// API 调用额度管理相关接口
// ============================================================================

// GetApiQuota 查询指定接口的调用额度和频率限制
func (c *Client) GetApiQuota(ctx context.Context, reqData *GetApiQuotaRequest) (*GetApiQuotaResponse, error) {
	if c.AccessToken == "" {
		return nil, fmt.Errorf("access_token 不能为空")
	}
	if reqData == nil {
		return nil, fmt.Errorf("请求参数不能为空")
	}
	if reqData.CGIPath == "" {
		return nil, fmt.Errorf("cgi_path 不能为空")
	}

	query := url.Values{}
	query.Set("access_token", c.AccessToken)

	bodyData := GetApiQuotaBody{
		CGIPath: reqData.CGIPath,
	}
	bodyBytes, err := json.Marshal(bodyData)
	if err != nil {
		return nil, fmt.Errorf("序列化请求参数失败：%w", err)
	}

	reqURL := c.Env + OpenApiQuotaGetApi + "?" + query.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("构建查询 API 调用额度请求失败：%w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求查询 API 调用额度失败：%w", err)
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败：%w", err)
	}

	var result GetApiQuotaResponse
	if err = json.Unmarshal(respBytes, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败：%w", err)
	}

	return &result, nil
}

// ClearApiQuota 重置指定 API 的调用次数
func (c *Client) ClearApiQuota(ctx context.Context, reqData *ClearApiQuotaRequest) (*ClearApiQuotaResponse, error) {
	if c.AccessToken == "" {
		return nil, fmt.Errorf("access_token 不能为空")
	}
	if reqData == nil {
		return nil, fmt.Errorf("请求参数不能为空")
	}
	if reqData.CGIPath == "" {
		return nil, fmt.Errorf("cgi_path 不能为空")
	}
	if !strings.HasPrefix(reqData.CGIPath, "/channels/ec/") {
		return nil, fmt.Errorf("cgi_path 必须以 /channels/ec/ 开头")
	}

	query := url.Values{}
	query.Set("access_token", c.AccessToken)

	bodyData := ClearApiQuotaBody{
		CGIPath: reqData.CGIPath,
	}
	bodyBytes, err := json.Marshal(bodyData)
	if err != nil {
		return nil, fmt.Errorf("序列化请求参数失败：%w", err)
	}

	reqURL := c.Env + OpenApiQuotaClearApi + "?" + query.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("构建重置 API 调用次数请求失败：%w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求重置 API 调用次数失败：%w", err)
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败：%w", err)
	}

	var result ClearApiQuotaResponse
	if err = json.Unmarshal(respBytes, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败：%w", err)
	}

	return &result, nil
}

// ClearQuota 重置每日 API 调用次数（使用 access_token）
func (c *Client) ClearQuota(ctx context.Context, reqData *ClearQuotaRequest) (*ClearQuotaResponse, error) {
	if c.AccessToken == "" {
		return nil, fmt.Errorf("access_token 不能为空")
	}
	if reqData == nil {
		reqData = &ClearQuotaRequest{}
	}
	appID := reqData.AppID
	if appID == "" {
		appID = c.AppID
	}
	if appID == "" {
		return nil, fmt.Errorf("appid 不能为空")
	}

	query := url.Values{}
	query.Set("access_token", c.AccessToken)

	bodyData := clearQuotaBody{
		AppID: appID,
	}
	bodyBytes, err := json.Marshal(bodyData)
	if err != nil {
		return nil, fmt.Errorf("序列化请求参数失败：%w", err)
	}

	reqURL := c.Env + ClearQuotaApi + "?" + query.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("构建重置每日 API 调用次数请求失败：%w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求重置每日 API 调用次数失败：%w", err)
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败：%w", err)
	}

	var result ClearQuotaResponse
	if err = json.Unmarshal(respBytes, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败：%w", err)
	}

	return &result, nil
}

// ClearQuotaByAppSecret 使用 AppSecret 重置每日 API 调用次数
func (c *Client) ClearQuotaByAppSecret(ctx context.Context, reqData *ClearQuotaByAppSecretRequest) (*ClearQuotaByAppSecretResponse, error) {
	if reqData == nil {
		reqData = &ClearQuotaByAppSecretRequest{}
	}
	appID := reqData.AppID
	if appID == "" {
		appID = c.AppID
	}
	if appID == "" {
		return nil, fmt.Errorf("appid 不能为空")
	}
	appSecret := reqData.AppSecret
	if appSecret == "" {
		appSecret = c.AppSecret
	}
	if appSecret == "" {
		return nil, fmt.Errorf("appsecret 不能为空")
	}

	bodyData := clearQuotaByAppSecretBody{
		AppID:     appID,
		AppSecret: appSecret,
	}
	bodyBytes, err := json.Marshal(bodyData)
	if err != nil {
		return nil, fmt.Errorf("序列化请求参数失败：%w", err)
	}

	reqURL := c.Env + ClearQuotaBySecretApi
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("构建使用 AppSecret 重置每日 API 调用次数请求失败：%w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求使用 AppSecret 重置每日 API 调用次数失败：%w", err)
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败：%w", err)
	}

	var result ClearQuotaByAppSecretResponse
	if err = json.Unmarshal(respBytes, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败：%w", err)
	}

	return &result, nil
}

// GetApiQuotaRequest 查询 API 调用额度请求参数
type GetApiQuotaRequest struct {
	// CGIPath 待查询接口路径，示例：/cgi-bin/message/custom/send
	CGIPath string
}

// GetApiQuotaBody 查询 API 调用额度请求体
type GetApiQuotaBody struct {
	// CGIPath 待查询接口路径
	CGIPath string `json:"cgi_path"`
}

// GetApiQuotaResponse 查询 API 调用额度响应
type GetApiQuotaResponse struct {
	// ErrCode 返回码
	ErrCode int `json:"errcode"`
	// ErrMsg 错误信息
	ErrMsg string `json:"errmsg"`
	// Quota 当日调用额度详情
	Quota GetApiQuotaQuota `json:"quota"`
	// RateLimit 普通调用频率限制
	RateLimit GetApiQuotaRateLimit `json:"rate_limit"`
	// ComponentRateLimit 代调用频率限制
	ComponentRateLimit GetApiQuotaRateLimit `json:"component_rate_limit"`
}

// GetApiQuotaQuota 当日调用额度详情
type GetApiQuotaQuota struct {
	// DailyLimit 当日可调用总次数
	DailyLimit int `json:"daily_limit"`
	// Used 当日已调用次数
	Used int `json:"used"`
	// Remain 当日剩余调用次数
	Remain int `json:"remain"`
}

// GetApiQuotaRateLimit 调用频率限制
type GetApiQuotaRateLimit struct {
	// CallCount 周期内可调用次数
	CallCount int `json:"call_count"`
	// RefreshSecond 刷新周期，单位秒
	RefreshSecond int `json:"refresh_second"`
}

// ClearApiQuotaRequest 重置 API 调用次数请求参数
type ClearApiQuotaRequest struct {
	CGIPath string
}

// ClearApiQuotaBody 重置 API 调用次数请求体
type ClearApiQuotaBody struct {
	CGIPath string `json:"cgi_path"`
}

// ClearApiQuotaResponse 重置 API 调用次数响应
type ClearApiQuotaResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// ClearQuotaRequest 重置每日 API 调用次数请求参数
type ClearQuotaRequest struct {
	AppID string
}

// clearQuotaBody 重置每日 API 调用次数请求体
type clearQuotaBody struct {
	AppID string `json:"appid"`
}

// ClearQuotaResponse 重置每日 API 调用次数响应
type ClearQuotaResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// ClearQuotaByAppSecretRequest 使用 AppSecret 重置每日 API 调用次数请求参数
type ClearQuotaByAppSecretRequest struct {
	AppID     string
	AppSecret string
}

// clearQuotaByAppSecretBody 使用 AppSecret 重置每日 API 调用次数请求体
type clearQuotaByAppSecretBody struct {
	AppID     string `json:"appid"`
	AppSecret string `json:"appsecret"`
}

// ClearQuotaByAppSecretResponse 使用 AppSecret 重置每日 API 调用次数响应
type ClearQuotaByAppSecretResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// ============================================================================
// RID 管理相关接口
// ============================================================================

// GetRidInfo 查询 rid 信息
func (c *Client) GetRidInfo(ctx context.Context, reqData *GetRidInfoRequest) (*GetRidInfoResponse, error) {
	if c.AccessToken == "" {
		return nil, fmt.Errorf("access_token 不能为空")
	}
	if reqData == nil {
		return nil, fmt.Errorf("请求参数不能为空")
	}
	if reqData.RID == "" {
		return nil, fmt.Errorf("rid 不能为空")
	}

	query := url.Values{}
	query.Set("access_token", c.AccessToken)

	bodyData := getRidInfoBody{
		RID: reqData.RID,
	}
	bodyBytes, err := json.Marshal(bodyData)
	if err != nil {
		return nil, fmt.Errorf("序列化请求参数失败：%w", err)
	}

	reqURL := c.Env + GetRidInfoApi + "?" + query.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("构建查询 rid 信息请求失败：%w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求查询 rid 信息失败：%w", err)
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败：%w", err)
	}

	var result GetRidInfoResponse
	if err = json.Unmarshal(respBytes, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败：%w", err)
	}

	return &result, nil
}

// GetRidInfoRequest 查询 rid 信息请求参数
type GetRidInfoRequest struct {
	RID string
}

// getRidInfoBody 查询 rid 信息请求体
type getRidInfoBody struct {
	RID string `json:"rid"`
}

// GetRidInfoResponse 查询 rid 信息响应
type GetRidInfoResponse struct {
	ErrCode int               `json:"errcode"`
	ErrMsg  string            `json:"errmsg"`
	Request GetRidInfoPayload `json:"request"`
}

// GetRidInfoPayload rid 信息详情
type GetRidInfoPayload struct {
	InvokeTime   int64  `json:"invoke_time"`
	CostInMS     int64  `json:"cost_in_ms"`
	RequestURL   string `json:"request_url"`
	RequestBody  string `json:"request_body"`
	ResponseBody string `json:"response_body"`
	ClientIP     string `json:"client_ip"`
}

// ============================================================================
// 回调检测相关接口
// ============================================================================

const (
	CallbackCheckActionDNS  = "dns"
	CallbackCheckActionPing = "ping"
	CallbackCheckActionAll  = "all"

	CallbackCheckOperatorChinaNet = "CHINANET"
	CallbackCheckOperatorUnicom   = "UNICOM"
	CallbackCheckOperatorCap      = "CAP"
	CallbackCheckOperatorDefault  = "DEFAULT"
)

// CallbackCheck 网络通信检测
func (c *Client) CallbackCheck(ctx context.Context, reqData *CallbackCheckRequest) (*CallbackCheckResponse, error) {
	if c.AccessToken == "" {
		return nil, fmt.Errorf("access_token 不能为空")
	}
	if reqData == nil {
		reqData = &CallbackCheckRequest{}
	}

	action := reqData.Action
	if action == "" {
		action = CallbackCheckActionAll
	}
	checkOperator := reqData.CheckOperator
	if checkOperator == "" {
		checkOperator = CallbackCheckOperatorDefault
	}

	query := url.Values{}
	query.Set("access_token", c.AccessToken)

	bodyData := callbackCheckBody{
		Action:        action,
		CheckOperator: checkOperator,
	}
	bodyBytes, err := json.Marshal(bodyData)
	if err != nil {
		return nil, fmt.Errorf("序列化请求参数失败：%w", err)
	}

	reqURL := c.Env + CallbackCheckApi + "?" + query.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("构建网络通信检测请求失败：%w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求网络通信检测失败：%w", err)
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败：%w", err)
	}

	var result CallbackCheckResponse
	if err = json.Unmarshal(respBytes, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败：%w", err)
	}

	return &result, nil
}

// CallbackCheckRequest 网络通信检测请求参数
type CallbackCheckRequest struct {
	Action        string
	CheckOperator string
}

// callbackCheckBody 网络通信检测请求体
type callbackCheckBody struct {
	Action        string `json:"action"`
	CheckOperator string `json:"check_operator"`
}

// CallbackCheckResponse 网络通信检测响应
type CallbackCheckResponse struct {
	ErrCode int                     `json:"errcode"`
	ErrMsg  string                  `json:"errmsg"`
	DNS     []CallbackCheckDNSItem  `json:"dns"`
	Ping    []CallbackCheckPingItem `json:"ping"`
}

// CallbackCheckDNSItem DNS 检测结果
type CallbackCheckDNSItem struct {
	IP           string `json:"ip"`
	RealOperator string `json:"real_operator"`
}

// CallbackCheckPingItem Ping 检测结果
type CallbackCheckPingItem struct {
	IP           string `json:"ip"`
	FromOperator string `json:"from_operator"`
	PackageLoss  string `json:"package_loss"`
}
