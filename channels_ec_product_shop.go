package wxshop_go_sdk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	// ProductDataTypeOnline 获取线上数据
	ProductDataTypeOnline = 1
	// ProductDataTypeDraft 获取草稿数据
	ProductDataTypeDraft = 2
	// ProductDataTypeOnlineDraft 同时获取线上与草稿数据
	ProductDataTypeOnlineDraft = 3
)

// GetProduct 获取商品详情
func (c *Client) GetProduct(ctx context.Context, reqData *GetProductRequest) (*GetProductResponse, error) {
	if c.AccessToken == "" {
		return nil, fmt.Errorf("access_token 不能为空")
	}
	if reqData == nil {
		return nil, fmt.Errorf("请求参数不能为空")
	}
	if reqData.ProductID == "" {
		return nil, fmt.Errorf("product_id 不能为空")
	}
	if reqData.DataType != 0 && reqData.DataType != ProductDataTypeOnline && reqData.DataType != ProductDataTypeDraft && reqData.DataType != ProductDataTypeOnlineDraft {
		return nil, fmt.Errorf("data_type 仅支持 1/2/3")
	}

	query := url.Values{}
	query.Set("access_token", c.AccessToken)

	bodyData := getProductBody{
		ProductID: reqData.ProductID,
	}
	if reqData.DataType != 0 {
		bodyData.DataType = reqData.DataType
	}
	bodyBytes, err := json.Marshal(bodyData)
	if err != nil {
		return nil, fmt.Errorf("序列化请求参数失败：%w", err)
	}

	reqURL := c.Env + ProductGetApi + "?" + query.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("构建获取商品请求失败：%w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求获取商品失败：%w", err)
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败：%w", err)
	}

	var result GetProductResponse
	if err = json.Unmarshal(respBytes, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败：%w", err)
	}
	return &result, nil
}

// GetProductList 获取商品列表
func (c *Client) GetProductList(ctx context.Context, reqData *GetProductListRequest) (*GetProductListResponse, error) {
	if c.AccessToken == "" {
		return nil, fmt.Errorf("access_token 不能为空")
	}
	if reqData == nil {
		return nil, fmt.Errorf("请求参数不能为空")
	}
	if reqData.PageSize <= 0 {
		return nil, fmt.Errorf("page_size 不能为空")
	}
	if reqData.PageSize > 30 {
		return nil, fmt.Errorf("page_size 不能大于 30")
	}

	query := url.Values{}
	query.Set("access_token", c.AccessToken)

	bodyData := getProductListBody{
		PageSize: reqData.PageSize,
		NextKey:  reqData.NextKey,
	}
	if reqData.Status != nil {
		bodyData.Status = reqData.Status
	}
	bodyBytes, err := json.Marshal(bodyData)
	if err != nil {
		return nil, fmt.Errorf("序列化请求参数失败：%w", err)
	}

	reqURL := c.Env + ProductListGetApi + "?" + query.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("构建获取商品列表请求失败：%w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求获取商品列表失败：%w", err)
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败：%w", err)
	}

	var result GetProductListResponse
	if err = json.Unmarshal(respBytes, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败：%w", err)
	}
	return &result, nil
}

// GetProductH5URL 获取商品 H5 短链
func (c *Client) GetProductH5URL(ctx context.Context, reqData *GetProductH5URLRequest) (*GetProductH5URLResponse, error) {
	if c.AccessToken == "" {
		return nil, fmt.Errorf("access_token 不能为空")
	}
	if reqData == nil {
		return nil, fmt.Errorf("请求参数不能为空")
	}
	if reqData.ProductID == "" {
		return nil, fmt.Errorf("product_id 不能为空")
	}

	query := url.Values{}
	query.Set("access_token", c.AccessToken)

	bodyBytes, err := json.Marshal(reqData)
	if err != nil {
		return nil, fmt.Errorf("序列化请求参数失败：%w", err)
	}

	reqURL := c.Env + ProductH5URLGetApi + "?" + query.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("构建获取商品 H5 短链请求失败：%w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求获取商品 H5 短链失败：%w", err)
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败：%w", err)
	}

	var result GetProductH5URLResponse
	if err = json.Unmarshal(respBytes, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败：%w", err)
	}
	return &result, nil
}

// GetProductTagLink 获取商品微信口令
func (c *Client) GetProductTagLink(ctx context.Context, reqData *GetProductTagLinkRequest) (*GetProductTagLinkResponse, error) {
	if c.AccessToken == "" {
		return nil, fmt.Errorf("access_token 不能为空")
	}
	if reqData == nil {
		return nil, fmt.Errorf("请求参数不能为空")
	}
	if reqData.ProductID == "" {
		return nil, fmt.Errorf("product_id 不能为空")
	}

	query := url.Values{}
	query.Set("access_token", c.AccessToken)

	bodyBytes, err := json.Marshal(reqData)
	if err != nil {
		return nil, fmt.Errorf("序列化请求参数失败：%w", err)
	}

	reqURL := c.Env + ProductTagLinkGetApi + "?" + query.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("构建获取商品口令请求失败：%w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求获取商品口令失败：%w", err)
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败：%w", err)
	}

	var result GetProductTagLinkResponse
	if err = json.Unmarshal(respBytes, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败：%w", err)
	}
	return &result, nil
}

// GetProductQRCode 获取商品二维码
func (c *Client) GetProductQRCode(ctx context.Context, reqData *GetProductQRCodeRequest) (*GetProductQRCodeResponse, error) {
	if c.AccessToken == "" {
		return nil, fmt.Errorf("access_token 不能为空")
	}
	if reqData == nil {
		return nil, fmt.Errorf("请求参数不能为空")
	}
	if reqData.ProductID == "" {
		return nil, fmt.Errorf("product_id 不能为空")
	}
	if reqData.QRCodeType != 0 && (reqData.QRCodeType < 1 || reqData.QRCodeType > 3) {
		return nil, fmt.Errorf("qrcode_type 仅支持 1/2/3")
	}

	query := url.Values{}
	query.Set("access_token", c.AccessToken)

	bodyData := *reqData
	if bodyData.QRCodeType == 0 {
		bodyData.QRCodeType = 1
	}
	bodyBytes, err := json.Marshal(bodyData)
	if err != nil {
		return nil, fmt.Errorf("序列化请求参数失败：%w", err)
	}

	reqURL := c.Env + ProductQRCodeGetApi + "?" + query.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("构建获取商品二维码请求失败：%w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求获取商品二维码失败：%w", err)
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败：%w", err)
	}

	var result GetProductQRCodeResponse
	if err = json.Unmarshal(respBytes, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败：%w", err)
	}
	return &result, nil
}

// GetProductScheme 获取商品移动应用跳转 scheme 码
func (c *Client) GetProductScheme(ctx context.Context, reqData *GetProductSchemeRequest) (*GetProductSchemeResponse, error) {
	if c.AccessToken == "" {
		return nil, fmt.Errorf("access_token 不能为空")
	}
	if reqData == nil {
		return nil, fmt.Errorf("请求参数不能为空")
	}
	if reqData.ProductID == "" {
		return nil, fmt.Errorf("product_id 不能为空")
	}
	if reqData.FromAppID == "" {
		return nil, fmt.Errorf("from_appid 不能为空")
	}
	if reqData.Expire <= 0 {
		return nil, fmt.Errorf("expire 不能为空")
	}

	query := url.Values{}
	query.Set("access_token", c.AccessToken)

	bodyBytes, err := json.Marshal(reqData)
	if err != nil {
		return nil, fmt.Errorf("序列化请求参数失败：%w", err)
	}

	reqURL := c.Env + ProductSchemeGetApi + "?" + query.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("构建获取商品 scheme 码请求失败：%w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求获取商品 scheme 码失败：%w", err)
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败：%w", err)
	}

	var result GetProductSchemeResponse
	if err = json.Unmarshal(respBytes, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败：%w", err)
	}
	return &result, nil
}

// GetProductAuditStrategy 获取商品上架策略
func (c *Client) GetProductAuditStrategy(ctx context.Context) (*GetProductAuditStrategyResponse, error) {
	if c.AccessToken == "" {
		return nil, fmt.Errorf("access_token 不能为空")
	}

	query := url.Values{}
	query.Set("access_token", c.AccessToken)

	reqURL := c.Env + ProductAuditStrategyGetApi + "?" + query.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, bytes.NewReader([]byte("{}")))
	if err != nil {
		return nil, fmt.Errorf("构建获取商品上架策略请求失败：%w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求获取商品上架策略失败：%w", err)
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败：%w", err)
	}

	var result GetProductAuditStrategyResponse
	if err = json.Unmarshal(respBytes, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败：%w", err)
	}
	return &result, nil
}

// GetProductRequest 获取商品请求参数
type GetProductRequest struct {
	ProductID string // 商品 ID
	DataType  int    // 数据类型，1 线上，2 草稿，3 线上+草稿
}

type getProductBody struct {
	ProductID string `json:"product_id"`          // 商品 ID
	DataType  int    `json:"data_type,omitempty"` // 数据类型
}

// GetProductResponse 获取商品响应
type GetProductResponse struct {
	ErrCode       int             `json:"errcode"`         // 错误码
	ErrMsg        string          `json:"errmsg"`          // 错误信息
	Product       json.RawMessage `json:"product"`         // 商品线上数据
	EditProduct   json.RawMessage `json:"edit_product"`    // 商品草稿数据
	SaleLimitInfo json.RawMessage `json:"sale_limit_info"` // 售卖上限提醒
	InfoScore     json.RawMessage `json:"info_score"`      // 商品信息质量
	CmpPriceInfo  json.RawMessage `json:"cmp_price_info"`  // 商品高价预警
	AuditInfo     json.RawMessage `json:"audit_info"`      // 审核信息
}

// GetProductListRequest 获取商品列表请求参数
type GetProductListRequest struct {
	Status   *int   // 商品状态
	PageSize int    // 每页数量，最大 30
	NextKey  string // 翻页游标
}

type getProductListBody struct {
	Status   *int   `json:"status,omitempty"`   // 商品状态
	PageSize int    `json:"page_size"`          // 每页数量
	NextKey  string `json:"next_key,omitempty"` // 翻页游标
}

// GetProductListResponse 获取商品列表响应
type GetProductListResponse struct {
	ErrCode    int     `json:"errcode"`     // 错误码
	ErrMsg     string  `json:"errmsg"`      // 错误信息
	ProductIDs []int64 `json:"product_ids"` // 商品 ID 列表
	NextKey    string  `json:"next_key"`    // 下一页游标
	TotalNum   int     `json:"total_num"`   // 商品总数
}

// GetProductH5URLRequest 获取商品 H5 短链请求参数
type GetProductH5URLRequest struct {
	ProductID   string `json:"product_id"`              // 商品 ID
	WecomCorpID string `json:"wecom_corp_id,omitempty"` // 企业微信 ID
	WecomUserID string `json:"wecom_user_id,omitempty"` // 企业微信成员 ID
}

// GetProductH5URLResponse 获取商品 H5 短链响应
type GetProductH5URLResponse struct {
	ErrCode      int    `json:"errcode"`       // 错误码
	ErrMsg       string `json:"errmsg"`        // 错误信息
	ProductH5URL string `json:"product_h5url"` // 商品 H5 短链
}

// GetProductTagLinkRequest 获取商品口令请求参数
type GetProductTagLinkRequest struct {
	ProductID   string `json:"product_id"`              // 商品 ID
	WecomCorpID string `json:"wecom_corp_id,omitempty"` // 企业微信 ID
	WecomUserID string `json:"wecom_user_id,omitempty"` // 企业微信成员 ID
}

// GetProductTagLinkResponse 获取商品口令响应
type GetProductTagLinkResponse struct {
	ErrCode        int    `json:"errcode"`         // 错误码
	ErrMsg         string `json:"errmsg"`          // 错误信息
	ProductTagLink string `json:"product_taglink"` // 商品口令
}

// GetProductQRCodeRequest 获取商品二维码请求参数
type GetProductQRCodeRequest struct {
	ProductID   string `json:"product_id"`              // 商品 ID
	QRCodeType  int    `json:"qrcode_type,omitempty"`   // 二维码类型
	WecomCorpID string `json:"wecom_corp_id,omitempty"` // 企业微信 ID
	WecomUserID string `json:"wecom_user_id,omitempty"` // 企业微信成员 ID
}

// GetProductQRCodeResponse 获取商品二维码响应
type GetProductQRCodeResponse struct {
	ErrCode       int    `json:"errcode"`        // 错误码
	ErrMsg        string `json:"errmsg"`         // 错误信息
	ProductQRCode string `json:"product_qrcode"` // 商品二维码链接
}

// GetProductSchemeRequest 获取商品 scheme 码请求参数
type GetProductSchemeRequest struct {
	ProductID string `json:"product_id"`         // 商品 ID
	FromAppID string `json:"from_appid"`         // 来源移动应用 AppID
	Expire    int    `json:"expire"`             // 过期时间，单位秒
	ExtInfo   string `json:"ext_info,omitempty"` // 附加信息
}

// GetProductSchemeResponse 获取商品 scheme 码响应
type GetProductSchemeResponse struct {
	ErrCode  int    `json:"errcode"`  // 错误码
	ErrMsg   string `json:"errmsg"`   // 错误信息
	OpenLink string `json:"openlink"` // 商品跳转 scheme 码
}

// GetProductAuditStrategyResponse 获取商品上架策略响应
type GetProductAuditStrategyResponse struct {
	ErrCode       int                  `json:"errcode"`        // 错误码
	ErrMsg        string               `json:"errmsg"`         // 错误信息
	AuditStrategy ProductAuditStrategy `json:"audit_strategy"` // 上架策略
}

// ProductAuditStrategy 商品上架策略
type ProductAuditStrategy struct {
	HideErrFieldFlag   int `json:"hide_err_field_flag"`    // 隐藏商品信息上架开关
	HitDuplicatedFlag  int `json:"hit_duplicated_flag"`    // 相似品可上架开关
	HitLowRiskRuleFlag int `json:"hit_low_risk_rule_flag"` // 低风险命中可上架开关
}
