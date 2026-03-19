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
	ShopQRCodeTypeNormal = 1
	ShopQRCodeTypeAsset  = 2
	ShopQRCodeTypeGift   = 3
)

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

type GetShopBasicInfoResponse struct {
	ErrCode int              `json:"errcode"`
	ErrMsg  string           `json:"errmsg"`
	Info    ShopBasicInfoDTO `json:"info"`
}

type ShopBasicInfoDTO struct {
	Nickname    string `json:"nickname"`
	HeadImgURL  string `json:"headimg_url"`
	SubjectType string `json:"subject_type"`
	Status      string `json:"status"`
	Username    string `json:"username"`
	IsLocalLife int    `json:"is_local_life"`
	OpenTime    int64  `json:"open_timestamp"`
}

type GetShopH5URLRequest struct {
	WecomCorpID string `json:"wecom_corp_id,omitempty"`
	WecomUserID string `json:"wecom_user_id,omitempty"`
}

type GetShopH5URLResponse struct {
	ErrCode   int    `json:"errcode"`
	ErrMsg    string `json:"errmsg"`
	ShopH5URL string `json:"shop_h5url"`
}

type GetShopQRCodeRequest struct {
	WecomCorpID string `json:"wecom_corp_id,omitempty"`
	WecomUserID string `json:"wecom_user_id,omitempty"`
	QRCodeType  int    `json:"qrcode_type"`
}

type GetShopQRCodeResponse struct {
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
	ShopQRCode string `json:"shop_qrcode"`
}

type GetShopTagLinkRequest struct {
	WecomCorpID string `json:"wecom_corp_id,omitempty"`
	WecomUserID string `json:"wecom_user_id,omitempty"`
}

type GetShopTagLinkResponse struct {
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
	ShopTagLink string `json:"shop_taglink"`
}
