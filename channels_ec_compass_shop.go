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

// GetCompassShopFinderAuthorizationList 获取授权视频号列表
func (c *Client) GetCompassShopFinderAuthorizationList(ctx context.Context) (*GetCompassShopFinderAuthorizationListResponse, error) {
	if c.AccessToken == "" {
		return nil, fmt.Errorf("access_token 不能为空")
	}

	query := url.Values{}
	query.Set("access_token", c.AccessToken)

	bodyBytes, err := json.Marshal(map[string]interface{}{})
	if err != nil {
		return nil, fmt.Errorf("序列化请求参数失败：%w", err)
	}

	reqURL := c.Env + CompassShopFinderAuthorizationListGetApi + "?" + query.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("构建获取授权视频号列表请求失败：%w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求获取授权视频号列表失败：%w", err)
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败：%w", err)
	}

	var result GetCompassShopFinderAuthorizationListResponse
	if err = json.Unmarshal(respBytes, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败：%w", err)
	}

	return &result, nil
}

// GetCompassShopFinderOverall 获取带货数据概览
func (c *Client) GetCompassShopFinderOverall(ctx context.Context, reqData *GetCompassShopFinderOverallRequest) (*GetCompassShopFinderOverallResponse, error) {
	if c.AccessToken == "" {
		return nil, fmt.Errorf("access_token 不能为空")
	}
	if reqData == nil || reqData.Ds == "" {
		return nil, fmt.Errorf("ds（日期）不能为空")
	}

	query := url.Values{}
	query.Set("access_token", c.AccessToken)

	bodyBytes, err := json.Marshal(reqData)
	if err != nil {
		return nil, fmt.Errorf("序列化请求参数失败：%w", err)
	}

	reqURL := c.Env + CompassShopFinderOverallGetApi + "?" + query.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("构建获取带货数据概览请求失败：%w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求获取带货数据概览失败：%w", err)
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败：%w", err)
	}

	var result GetCompassShopFinderOverallResponse
	if err = json.Unmarshal(respBytes, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败：%w", err)
	}

	return &result, nil
}

// GetCompassShopFinderProductList 获取带货达人商品列表
func (c *Client) GetCompassShopFinderProductList(ctx context.Context, reqData *GetCompassShopFinderProductListRequest) (*GetCompassShopFinderProductListResponse, error) {
	if c.AccessToken == "" {
		return nil, fmt.Errorf("access_token 不能为空")
	}
	if reqData == nil || reqData.Ds == "" {
		return nil, fmt.Errorf("ds（日期）不能为空")
	}

	query := url.Values{}
	query.Set("access_token", c.AccessToken)

	bodyBytes, err := json.Marshal(reqData)
	if err != nil {
		return nil, fmt.Errorf("序列化请求参数失败：%w", err)
	}

	reqURL := c.Env + CompassShopFinderProductListGetApi + "?" + query.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("构建获取带货达人商品列表请求失败：%w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求获取带货达人商品列表失败：%w", err)
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败：%w", err)
	}

	var result GetCompassShopFinderProductListResponse
	if err = json.Unmarshal(respBytes, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败：%w", err)
	}

	return &result, nil
}

// GetCompassShopFinderProductOverall 获取带货达人详情
func (c *Client) GetCompassShopFinderProductOverall(ctx context.Context, reqData *GetCompassShopFinderProductOverallRequest) (*GetCompassShopFinderProductOverallResponse, error) {
	if c.AccessToken == "" {
		return nil, fmt.Errorf("access_token 不能为空")
	}
	if reqData == nil || reqData.Ds == "" {
		return nil, fmt.Errorf("ds（日期）不能为空")
	}

	query := url.Values{}
	query.Set("access_token", c.AccessToken)

	bodyBytes, err := json.Marshal(reqData)
	if err != nil {
		return nil, fmt.Errorf("序列化请求参数失败：%w", err)
	}

	reqURL := c.Env + CompassShopFinderProductOverallGetApi + "?" + query.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("构建获取带货达人详情请求失败：%w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求获取带货达人详情失败：%w", err)
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败：%w", err)
	}

	var result GetCompassShopFinderProductOverallResponse
	if err = json.Unmarshal(respBytes, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败：%w", err)
	}

	return &result, nil
}

// GetCompassShopLiveList 获取店铺开播列表
func (c *Client) GetCompassShopLiveList(ctx context.Context, reqData *GetCompassShopLiveListRequest) (*GetCompassShopLiveListResponse, error) {
	if c.AccessToken == "" {
		return nil, fmt.Errorf("access_token 不能为空")
	}
	if reqData == nil || reqData.Ds == "" {
		return nil, fmt.Errorf("ds（日期）不能为空")
	}
	if reqData.FinderID == "" {
		return nil, fmt.Errorf("finder_id（视频号 ID）不能为空")
	}

	query := url.Values{}
	query.Set("access_token", c.AccessToken)

	bodyBytes, err := json.Marshal(reqData)
	if err != nil {
		return nil, fmt.Errorf("序列化请求参数失败：%w", err)
	}

	reqURL := c.Env + CompassShopLiveListGetApi + "?" + query.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("构建获取店铺开播列表请求失败：%w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求获取店铺开播列表失败：%w", err)
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败：%w", err)
	}

	var result GetCompassShopLiveListResponse
	if err = json.Unmarshal(respBytes, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败：%w", err)
	}

	return &result, nil
}

// GetCompassShopOverall 获取电商数据概览
func (c *Client) GetCompassShopOverall(ctx context.Context, reqData *GetCompassShopOverallRequest) (*GetCompassShopOverallResponse, error) {
	if c.AccessToken == "" {
		return nil, fmt.Errorf("access_token 不能为空")
	}
	if reqData == nil || reqData.Ds == "" {
		return nil, fmt.Errorf("ds（日期）不能为空")
	}

	query := url.Values{}
	query.Set("access_token", c.AccessToken)

	bodyBytes, err := json.Marshal(reqData)
	if err != nil {
		return nil, fmt.Errorf("序列化请求参数失败：%w", err)
	}

	reqURL := c.Env + CompassShopOverallGetApi + "?" + query.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("构建获取电商数据概览请求失败：%w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求获取电商数据概览失败：%w", err)
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败：%w", err)
	}

	var result GetCompassShopOverallResponse
	if err = json.Unmarshal(respBytes, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败：%w", err)
	}

	return &result, nil
}

// GetCompassShopProductData 获取商品详细信息
func (c *Client) GetCompassShopProductData(ctx context.Context, reqData *GetCompassShopProductDataRequest) (*GetCompassShopProductDataResponse, error) {
	if c.AccessToken == "" {
		return nil, fmt.Errorf("access_token 不能为空")
	}
	if reqData == nil || reqData.Ds == "" {
		return nil, fmt.Errorf("ds（日期）不能为空")
	}
	if reqData.ProductID <= 0 {
		return nil, fmt.Errorf("product_id（商品 id）不能为空")
	}

	query := url.Values{}
	query.Set("access_token", c.AccessToken)

	bodyBytes, err := json.Marshal(reqData)
	if err != nil {
		return nil, fmt.Errorf("序列化请求参数失败：%w", err)
	}

	reqURL := c.Env + CompassShopProductDataGetApi + "?" + query.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("构建获取商品详细信息请求失败：%w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求获取商品详细信息失败：%w", err)
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败：%w", err)
	}

	var result GetCompassShopProductDataResponse
	if err = json.Unmarshal(respBytes, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败：%w", err)
	}

	return &result, nil
}

// GetCompassShopProductList 获取商品列表
func (c *Client) GetCompassShopProductList(ctx context.Context, reqData *GetCompassShopProductListRequest) (*GetCompassShopProductListResponse, error) {
	if c.AccessToken == "" {
		return nil, fmt.Errorf("access_token 不能为空")
	}
	if reqData == nil || reqData.Ds == "" {
		return nil, fmt.Errorf("ds（日期）不能为空")
	}

	query := url.Values{}
	query.Set("access_token", c.AccessToken)

	bodyBytes, err := json.Marshal(reqData)
	if err != nil {
		return nil, fmt.Errorf("序列化请求参数失败：%w", err)
	}

	reqURL := c.Env + CompassShopProductListGetApi + "?" + query.Encode()
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

	var result GetCompassShopProductListResponse
	if err = json.Unmarshal(respBytes, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败：%w", err)
	}

	return &result, nil
}

// GetCompassShopSaleProfileData 获取店铺人群数据
func (c *Client) GetCompassShopSaleProfileData(ctx context.Context, reqData *GetCompassShopSaleProfileDataRequest) (*GetCompassShopSaleProfileDataResponse, error) {
	if c.AccessToken == "" {
		return nil, fmt.Errorf("access_token 不能为空")
	}
	if reqData == nil || reqData.Ds == "" {
		return nil, fmt.Errorf("ds（日期）不能为空")
	}
	if reqData.Type <= 0 || reqData.Type > 5 {
		return nil, fmt.Errorf("type（用户类型）必须在 1-5 之间")
	}

	query := url.Values{}
	query.Set("access_token", c.AccessToken)

	bodyBytes, err := json.Marshal(reqData)
	if err != nil {
		return nil, fmt.Errorf("序列化请求参数失败：%w", err)
	}

	reqURL := c.Env + CompassShopSaleProfileDataGetApi + "?" + query.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("构建获取店铺人群数据请求失败：%w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求获取店铺人群数据失败：%w", err)
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败：%w", err)
	}

	var result GetCompassShopSaleProfileDataResponse
	if err = json.Unmarshal(respBytes, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败：%w", err)
	}

	return &result, nil
}

// ==================== 请求结构体 ====================

// GetCompassShopFinderOverallRequest 获取带货数据概览请求参数
type GetCompassShopFinderOverallRequest struct {
	Ds string `json:"ds"` // 日期，格式 YYYYMMDD
}

// GetCompassShopFinderProductListRequest 获取带货达人商品列表请求参数
type GetCompassShopFinderProductListRequest struct {
	Ds       string `json:"ds"`                  // 日期，格式 YYYYMMDD
	FinderID string `json:"finder_id,omitempty"` // 视频号 ID（finder_id、talent_id、mp_id 三者必填其一）
	TalentID string `json:"talent_id,omitempty"` // 达人号 ID（finder_id、talent_id、mp_id 三者必填其一）
	MpID     string `json:"mp_id,omitempty"`     // 公众号 ID（finder_id、talent_id、mp_id 三者必填其一）
}

// GetCompassShopFinderProductOverallRequest 获取带货达人详情请求参数
type GetCompassShopFinderProductOverallRequest struct {
	Ds       string `json:"ds"`                  // 日期，格式 YYYYMMDD
	FinderID string `json:"finder_id,omitempty"` // 视频号 id（finder_id、talent_id、mp_id 三者必填其一）
	TalentID string `json:"talent_id,omitempty"` // 达人号 ID（finder_id、talent_id、mp_id 三者必填其一）
	MpID     string `json:"mp_id,omitempty"`     // 公众号 ID（finder_id、talent_id、mp_id 三者必填其一）
}

// GetCompassShopLiveListRequest 获取店铺开播列表请求参数
type GetCompassShopLiveListRequest struct {
	Ds       string `json:"ds"`        // 日期，格式 YYYYMMDD
	FinderID string `json:"finder_id"` // 视频号 ID
}

// GetCompassShopOverallRequest 获取电商数据概览请求参数
type GetCompassShopOverallRequest struct {
	Ds string `json:"ds"` // 日期，格式 YYYYMMDD
}

// GetCompassShopProductDataRequest 获取商品详细信息请求参数
type GetCompassShopProductDataRequest struct {
	Ds        string `json:"ds"`         // 日期，格式 YYYYMMDD
	ProductID int64  `json:"product_id"` // 商品 id
}

// GetCompassShopProductListRequest 获取商品列表请求参数
type GetCompassShopProductListRequest struct {
	Ds     string `json:"ds"`               // 日期，格式 YYYYMMDD
	Limit  int    `json:"limit,omitempty"`  // 分页参数，如果不填，默认只会返回 10 条商品数据
	Offset int    `json:"offset,omitempty"` // 分页参数，如果不填，默认只会返回 10 条商品数据
}

// GetCompassShopSaleProfileDataRequest 获取店铺人群数据请求参数
type GetCompassShopSaleProfileDataRequest struct {
	Ds   string `json:"ds"`   // 日期，格式 YYYYMMDD
	Type int    `json:"type"` // 用户类型：1-商品曝光用户，2-商品点击用户，3-购买用户，4-首购用户，5-复购用户
}

// ==================== 响应结构体 ====================

// GetCompassShopFinderAuthorizationListResponse 获取授权视频号列表响应
type GetCompassShopFinderAuthorizationListResponse struct {
	ErrCode                int      `json:"errcode"`                   // 错误码
	ErrMsg                 string   `json:"errmsg"`                    // 错误信息
	AuthorizedFinderIDList []string `json:"authorized_finder_id_list"` // 授权视频号 id 列表
}

// GetCompassShopFinderOverallResponse 获取带货数据概览响应
type GetCompassShopFinderOverallResponse struct {
	ErrCode int                  `json:"errcode"` // 错误码
	ErrMsg  string               `json:"errmsg"`  // 错误信息
	Data    FinderOverallDataDTO `json:"data"`    // 带货数据
}

// FinderOverallDataDTO 带货数据概览
type FinderOverallDataDTO struct {
	PayGmv            string  `json:"pay_gmv"`               // 成交金额，单位分
	PaySalesFinderCnt string  `json:"pay_sales_finder_cnt"`  // 动销达人数
	PayProductIDCnt   string  `json:"pay_product_id_cnt"`    // 动销商品数
	ClickToPayUvRatio float64 `json:"click_to_pay_uv_ratio"` // 点击 - 成交转化率
}

// GetCompassShopFinderProductListResponse 获取带货达人商品列表响应
type GetCompassShopFinderProductListResponse struct {
	ErrCode     int                        `json:"errcode"`      // 错误码
	ErrMsg      string                     `json:"errmsg"`       // 错误信息
	ProductList []FinderProductListItemDTO `json:"product_list"` // 商品信息
}

// FinderProductListItemDTO 带货达人商品信息
type FinderProductListItemDTO struct {
	ProductID        int64                      `json:"product_id"`         // 商品 id
	HeadImgURL       string                     `json:"head_img_url"`       // 商品头图
	Title            string                     `json:"title"`              // 商品标题
	Price            int64                      `json:"price"`              // 商品价格
	FirstCategoryID  int64                      `json:"first_category_id"`  // 商品 1 级类目
	SecondCategoryID int64                      `json:"second_category_id"` // 商品 2 级类目
	ThirdCategoryID  int64                      `json:"third_category_id"`  // 商品 3 级类目
	Data             FinderProductCommissionDTO `json:"data"`               // 佣金数据
}

// FinderProductCommissionDTO 佣金数据
type FinderProductCommissionDTO struct {
	CommissionRatio float64 `json:"commission_ratio"` // 佣金率
	PayGmv          string  `json:"pay_gmv"`          // 成交金额
}

// GetCompassShopFinderProductOverallResponse 获取带货达人详情响应
type GetCompassShopFinderProductOverallResponse struct {
	ErrCode int                         `json:"errcode"` // 错误码
	ErrMsg  string                      `json:"errmsg"`  // 错误信息
	Data    FinderProductOverallDataDTO `json:"data"`    // 带货达人详情数据
}

// FinderProductOverallDataDTO 带货达人详情数据
type FinderProductOverallDataDTO struct {
	PayGmv          string `json:"pay_gmv"`            // 成交金额，单位分
	PayProductIDCnt string `json:"pay_product_id_cnt"` // 动销商品数
	PayUv           string `json:"pay_uv"`             // 成交人数
	RefundGmv       string `json:"refund_gmv"`         // 退款金额，单位分
	PayRefundGmv    string `json:"pay_refund_gmv"`     // 成交退款金额，单位分
}

// GetCompassShopLiveListResponse 获取店铺开播列表响应
type GetCompassShopLiveListResponse struct {
	ErrCode  int               `json:"errcode"`   // 错误码
	ErrMsg   string            `json:"errmsg"`    // 错误信息
	LiveList []LiveListItemDTO `json:"live_list"` // 开播列表
}

// LiveListItemDTO 开播信息
type LiveListItemDTO struct {
	LiveID          string `json:"live_id"`            // 直播 id
	LiveTitle       string `json:"live_title"`         // 直播标题
	LiveTime        string `json:"live_time"`          // 开播时间，unix 时间戳
	LiveDuration    string `json:"live_duration"`      // 直播时长
	LiveCoverImgURL string `json:"live_cover_img_url"` // 直播封面
}

// GetCompassShopOverallResponse 获取电商数据概览响应
type GetCompassShopOverallResponse struct {
	ErrCode int                `json:"errcode"` // 错误码
	ErrMsg  string             `json:"errmsg"`  // 错误信息
	Data    ShopOverallDataDTO `json:"data"`    // 电商数据
}

// ShopOverallDataDTO 电商数据概览
type ShopOverallDataDTO struct {
	PayGmv         string `json:"pay_gmv"`          // 成交金额，单位分
	PayUv          string `json:"pay_uv"`           // 成交人数
	PayOrderCnt    string `json:"pay_order_cnt"`    // 成交订单数
	PayRefundGmv   string `json:"pay_refund_gmv"`   // 成交退款金额，单位分
	LivePayGmv     string `json:"live_pay_gmv"`     // 直播成交金额，单位分
	FeedPayGmv     string `json:"feed_pay_gmv"`     // 短视频成交金额，单位分
	ProductClickUv string `json:"product_click_uv"` // 商品点击人数
}

// GetCompassShopProductDataResponse 获取商品详细信息响应
type GetCompassShopProductDataResponse struct {
	ErrCode     int                  `json:"errcode"`      // 错误码
	ErrMsg      string               `json:"errmsg"`       // 错误信息
	ProductInfo ProductDetailInfoDTO `json:"product_info"` // 商品详细信息
}

// ProductDetailInfoDTO 商品详细信息
type ProductDetailInfoDTO struct {
	ProductID        string               `json:"product_id"`         // 商品 id
	HeadImgURL       string               `json:"head_img_url"`       // 商品图
	Title            string               `json:"title"`              // 商品标题
	Price            string               `json:"price"`              // 商品价格，单位分
	FirstCategoryID  string               `json:"first_category_id"`  // 商品一级类目
	SecondCategoryID string               `json:"second_category_id"` // 商品二级类目
	ThirdCategoryID  string               `json:"third_category_id"`  // 商品三级类目
	Data             ProductDataDetailDTO `json:"data"`               // 详细数据
}

// ProductDataDetailDTO 商品详细数据
type ProductDataDetailDTO struct {
	PayGmv                     string  `json:"pay_gmv"`                      // 成交金额，单位分
	CreateGmv                  string  `json:"create_gmv"`                   // 下单金额，单位分
	CreateCnt                  string  `json:"create_cnt"`                   // 下单订单数
	CreateUv                   string  `json:"create_uv"`                    // 下单人数
	CreateProductCnt           string  `json:"create_product_cnt"`           // 下单件数
	PayCnt                     string  `json:"pay_cnt"`                      // 成交订单数
	PayUv                      string  `json:"pay_uv"`                       // 成交人数
	PayProductCnt              string  `json:"pay_product_cnt"`              // 成交件数
	PurePayGmv                 string  `json:"pure_pay_gmv"`                 // 成交金额（剔除退款）
	PayGmvPerUv                string  `json:"pay_gmv_per_uv"`               // 成交客单价（剔除退款）
	SellerActualSettleAmount   string  `json:"seller_actual_settle_amount"`  // 实际结算金额，单位分
	PlatformActualCommission   string  `json:"platform_actual_commission"`   // 实际服务费金额，单位分
	FinderuinActualCommission  string  `json:"finderuin_actual_commission"`  // 实际达人佣金支出，单位分
	CaptainActualCommission    string  `json:"captain_actual_commission"`    // 实际团长佣金支出，单位分
	SellerPredictSettleAmount  string  `json:"seller_predict_settle_amount"` // 预估结算金额，单位分
	PlatformPredictCommission  string  `json:"platform_predict_commission"`  // 预估服务费金额，单位分
	FinderuinPredictCommission string  `json:"finderuin_predict_commission"` // 预估达人佣金支出，单位分
	CaptainPredictCommission   string  `json:"captain_predict_commission"`   // 预估团长佣金支出，单位分
	ProductClickUv             string  `json:"product_click_uv"`             // 商品点击人数
	ProductClickCnt            string  `json:"product_click_cnt"`            // 商品点击次数
	PayRefundGmv               string  `json:"pay_refund_gmv"`               // 成交退款金额，单位分
	PayRefundUv                string  `json:"pay_refund_uv"`                // 成交退款人数
	PayRefundRatio             float64 `json:"pay_refund_ratio"`             // 成交退款率
	PayRefundAfterSendRatio    float64 `json:"pay_refund_after_send_ratio"`  // 发货后成交退款率
	PayRefundCnt               string  `json:"pay_refund_cnt"`               // 成交退款订单数
	PayRefundProductCnt        string  `json:"pay_refund_product_cnt"`       // 成交退款件数
	PayRefundBeforeSendRatio   float64 `json:"pay_refund_before_send_ratio"` // 发货前成交退款率
	RefundGmv                  string  `json:"refund_gmv"`                   // 退款金额，单位分
	RefundProductCnt           string  `json:"refund_product_cnt"`           // 退款件数
	RefundCnt                  string  `json:"refund_cnt"`                   // 退款订单数
	RefundUv                   string  `json:"refund_uv"`                    // 退款人数
}

// GetCompassShopProductListResponse 获取商品列表响应
type GetCompassShopProductListResponse struct {
	ErrCode     int                  `json:"errcode"`      // 错误码
	ErrMsg      string               `json:"errmsg"`       // 错误信息
	ProductList []ProductListItemDTO `json:"product_list"` // 商品列表
	TotalCount  int                  `json:"total_count"`  // 一共有多少条商品信息
}

// ProductListItemDTO 商品列表项
type ProductListItemDTO struct {
	ProductID        string               `json:"product_id"`         // 商品 id
	HeadImgURL       string               `json:"head_img_url"`       // 商品头图
	Title            string               `json:"title"`              // 商品标题
	Price            string               `json:"price"`              // 商品价格，单位分
	FirstCategoryID  string               `json:"first_category_id"`  // 商品一级类目
	SecondCategoryID string               `json:"second_category_id"` // 商品二级类目
	ThirdCategoryID  string               `json:"third_category_id"`  // 商品三级类目
	Data             ProductDataDetailDTO `json:"data"`               // 详细数据
}

// GetCompassShopSaleProfileDataResponse 获取店铺人群数据响应
type GetCompassShopSaleProfileDataResponse struct {
	ErrCode int                `json:"errcode"` // 错误码
	ErrMsg  string             `json:"errmsg"`  // 错误信息
	Data    SaleProfileDataDTO `json:"data"`    // 店铺人群数据
}

// SaleProfileDataDTO 店铺人群数据
type SaleProfileDataDTO struct {
	FieldList []FieldListItemDTO `json:"field_list"` // 维度列表
}

// FieldListItemDTO 维度列表项
type FieldListItemDTO struct {
	FieldName string             `json:"field_name"` // 维度类别名
	DataList  []DimensionDataDTO `json:"data_list"`  // 指标
}

// DimensionDataDTO 维度指标数据
type DimensionDataDTO struct {
	DimKey   string `json:"dim_key"`   // 维度指标名
	DimValue string `json:"dim_value"` // 维度指标值
}
