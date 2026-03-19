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
	// ShopQRCodeTypeNormal 普通二维码
	ShopQRCodeTypeNormal = 1
	// ShopQRCodeTypeAsset 标准物料
	ShopQRCodeTypeAsset = 2
	// ShopQRCodeTypeGift 送礼物物料
	ShopQRCodeTypeGift = 3
)

// GetShopBasicInfo 获取店铺基本信息
func (c *Client) GetShopBasicInfo(ctx context.Context) (*GetShopBasicInfoResponse, error) {
	if c.AccessToken == "" {
		return nil, fmt.Errorf("access_token 不能为空")
	}

	query := url.Values{}
	query.Set("access_token", c.AccessToken)

	reqURL := c.Env + BasicsInfoGetApi + "?" + query.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, fmt.Errorf("构建获取店铺基本信息请求失败：%w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求获取店铺基本信息失败：%w", err)
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败：%w", err)
	}

	var result GetShopBasicInfoResponse
	if err = json.Unmarshal(respBytes, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败：%w", err)
	}

	return &result, nil
}

// GetShopH5URL 获取店铺 H5 链接
func (c *Client) GetShopH5URL(ctx context.Context, reqData *GetShopH5URLRequest) (*GetShopH5URLResponse, error) {
	if c.AccessToken == "" {
		return nil, fmt.Errorf("access_token 不能为空")
	}
	if reqData == nil {
		reqData = &GetShopH5URLRequest{}
	}

	query := url.Values{}
	query.Set("access_token", c.AccessToken)

	bodyBytes, err := json.Marshal(reqData)
	if err != nil {
		return nil, fmt.Errorf("序列化请求参数失败：%w", err)
	}

	reqURL := c.Env + ShopH5URLGetApi + "?" + query.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("构建获取店铺 H5 链接请求失败：%w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求获取店铺 H5 链接失败：%w", err)
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败：%w", err)
	}

	var result GetShopH5URLResponse
	if err = json.Unmarshal(respBytes, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败：%w", err)
	}

	return &result, nil
}

// GetShopQRCode 获取店铺二维码
func (c *Client) GetShopQRCode(ctx context.Context, reqData *GetShopQRCodeRequest) (*GetShopQRCodeResponse, error) {
	if c.AccessToken == "" {
		return nil, fmt.Errorf("access_token 不能为空")
	}
	if reqData == nil {
		return nil, fmt.Errorf("请求参数不能为空")
	}
	if reqData.QRCodeType == 0 {
		return nil, fmt.Errorf("qrcode_type 不能为空")
	}

	query := url.Values{}
	query.Set("access_token", c.AccessToken)

	bodyBytes, err := json.Marshal(reqData)
	if err != nil {
		return nil, fmt.Errorf("序列化请求参数失败：%w", err)
	}

	reqURL := c.Env + ShopQRCodeGetApi + "?" + query.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("构建获取店铺二维码请求失败：%w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求获取店铺二维码失败：%w", err)
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败：%w", err)
	}

	var result GetShopQRCodeResponse
	if err = json.Unmarshal(respBytes, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败：%w", err)
	}

	return &result, nil
}

// GetShopTagLink 获取店铺口令
func (c *Client) GetShopTagLink(ctx context.Context, reqData *GetShopTagLinkRequest) (*GetShopTagLinkResponse, error) {
	if c.AccessToken == "" {
		return nil, fmt.Errorf("access_token 不能为空")
	}
	if reqData == nil {
		reqData = &GetShopTagLinkRequest{}
	}

	query := url.Values{}
	query.Set("access_token", c.AccessToken)

	bodyBytes, err := json.Marshal(reqData)
	if err != nil {
		return nil, fmt.Errorf("序列化请求参数失败：%w", err)
	}

	reqURL := c.Env + ShopTagLinkGetApi + "?" + query.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("构建获取店铺口令请求失败：%w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求获取店铺口令失败：%w", err)
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败：%w", err)
	}

	var result GetShopTagLinkResponse
	if err = json.Unmarshal(respBytes, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败：%w", err)
	}

	return &result, nil
}

// GetShopBasicInfoResponse 获取店铺基本信息响应
type GetShopBasicInfoResponse struct {
	ErrCode int              `json:"errcode"` // 错误码
	ErrMsg  string           `json:"errmsg"`  // 错误信息
	Info    ShopBasicInfoDTO `json:"info"`    // 店铺基础信息
}

// ShopBasicInfoDTO 店铺基础信息
type ShopBasicInfoDTO struct {
	Nickname    string `json:"nickname"`       // 店铺名称
	HeadImgURL  string `json:"headimg_url"`    // 店铺头像 URL
	SubjectType string `json:"subject_type"`   // 店铺主体类型
	Status      string `json:"status"`         // 店铺状态
	Username    string `json:"username"`       // 店铺原始 ID
	IsLocalLife int    `json:"is_local_life"`  // 是否本地生活小店
	OpenTime    int64  `json:"open_timestamp"` // 开店时间戳
}

// GetShopH5URLRequest 获取店铺 H5 链接请求参数
type GetShopH5URLRequest struct {
	WecomCorpID string `json:"wecom_corp_id,omitempty"` // 企业微信 ID
	WecomUserID string `json:"wecom_user_id,omitempty"` // 企业微信成员 ID
}

// GetShopH5URLResponse 获取店铺 H5 链接响应
type GetShopH5URLResponse struct {
	ErrCode   int    `json:"errcode"`    // 错误码
	ErrMsg    string `json:"errmsg"`     // 错误信息
	ShopH5URL string `json:"shop_h5url"` // 店铺 H5 链接
}

// GetShopQRCodeRequest 获取店铺二维码请求参数
type GetShopQRCodeRequest struct {
	WecomCorpID string `json:"wecom_corp_id,omitempty"` // 企业微信 ID
	WecomUserID string `json:"wecom_user_id,omitempty"` // 企业微信成员 ID
	QRCodeType  int    `json:"qrcode_type"`             // 二维码类型
}

// GetShopQRCodeResponse 获取店铺二维码响应
type GetShopQRCodeResponse struct {
	ErrCode    int    `json:"errcode"`     // 错误码
	ErrMsg     string `json:"errmsg"`      // 错误信息
	ShopQRCode string `json:"shop_qrcode"` // 店铺二维码链接
}

// GetShopTagLinkRequest 获取店铺口令请求参数
type GetShopTagLinkRequest struct {
	WecomCorpID string `json:"wecom_corp_id,omitempty"` // 企业微信 ID
	WecomUserID string `json:"wecom_user_id,omitempty"` // 企业微信成员 ID
}

// GetShopTagLinkResponse 获取店铺口令响应
type GetShopTagLinkResponse struct {
	ErrCode     int    `json:"errcode"`      // 错误码
	ErrMsg      string `json:"errmsg"`       // 错误信息
	ShopTagLink string `json:"shop_taglink"` // 店铺口令
}
