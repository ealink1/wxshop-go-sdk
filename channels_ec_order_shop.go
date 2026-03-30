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

// ============================================================================
// 订单状态枚举
// ============================================================================
const (
	OrderStatusUnpaid         = 10  // 待付款
	OrderStatusGiftPending    = 12  // 礼物待收下
	OrderStatusGrouping       = 13  // 一起买待成团/凑单买凑团中
	OrderStatusPendingShip    = 20  // 待发货（包括部分发货）
	OrderStatusPartialShipped = 21  // 部分发货
	OrderStatusPendingReceive = 30  // 待收货（包括部分发货）
	OrderStatusCompleted      = 100 // 完成
	OrderStatusCancelled      = 250 // 订单取消（包括未付款取消，售后取消等）
)

// ============================================================================
// 获取订单列表接口
// ============================================================================

// GetOrderList 获取订单列表
func (c *Client) GetOrderList(ctx context.Context, reqData *GetOrderListRequest) (*GetOrderListResponse, error) {
	if c.AccessToken == "" {
		return nil, fmt.Errorf("access_token 不能为空")
	}
	if reqData == nil {
		return nil, fmt.Errorf("请求参数不能为空")
	}
	if reqData.CreateTimeRange == nil && reqData.UpdateTimeRange == nil {
		return nil, fmt.Errorf("create_time_range 和 update_time_range 至少填一个")
	}

	query := url.Values{}
	query.Set("access_token", c.AccessToken)

	bodyBytes, err := json.Marshal(reqData)
	if err != nil {
		return nil, fmt.Errorf("序列化请求参数失败：%w", err)
	}

	reqURL := c.Env + OrderListGetApi + "?" + query.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("构建获取订单列表请求失败：%w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求获取订单列表失败：%w", err)
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败：%w", err)
	}

	var result GetOrderListResponse
	if err = json.Unmarshal(respBytes, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败：%w", err)
	}

	return &result, nil
}

// ============================================================================
// 获取订单详情接口
// ============================================================================

// GetOrder 获取订单详情
func (c *Client) GetOrder(ctx context.Context, reqData *GetOrderRequest) (*GetOrderResponse, error) {
	if c.AccessToken == "" {
		return nil, fmt.Errorf("access_token 不能为空")
	}
	if reqData == nil {
		return nil, fmt.Errorf("请求参数不能为空")
	}
	if reqData.OrderID == "" {
		return nil, fmt.Errorf("order_id 不能为空")
	}

	query := url.Values{}
	query.Set("access_token", c.AccessToken)

	bodyBytes, err := json.Marshal(reqData)
	if err != nil {
		return nil, fmt.Errorf("序列化请求参数失败：%w", err)
	}

	reqURL := c.Env + OrderGetApi + "?" + query.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("构建获取订单详情请求失败：%w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求获取订单详情失败：%w", err)
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败：%w", err)
	}

	var result GetOrderResponse
	if err = json.Unmarshal(respBytes, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败：%w", err)
	}

	return &result, nil
}

// ============================================================================
// 搜索订单接口
// ============================================================================

// SearchOrder 搜索订单
func (c *Client) SearchOrder(ctx context.Context, reqData *SearchOrderRequest) (*SearchOrderResponse, error) {
	if c.AccessToken == "" {
		return nil, fmt.Errorf("access_token 不能为空")
	}
	if reqData == nil {
		return nil, fmt.Errorf("请求参数不能为空")
	}
	if reqData.SearchCondition == nil {
		return nil, fmt.Errorf("search_condition 不能为空")
	}

	query := url.Values{}
	query.Set("access_token", c.AccessToken)

	bodyBytes, err := json.Marshal(reqData)
	if err != nil {
		return nil, fmt.Errorf("序列化请求参数失败：%w", err)
	}

	reqURL := c.Env + OrderSearchApi + "?" + query.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("构建搜索订单请求失败：%w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求搜索订单失败：%w", err)
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败：%w", err)
	}

	var result SearchOrderResponse
	if err = json.Unmarshal(respBytes, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败：%w", err)
	}

	return &result, nil
}

// ============================================================================
// 请求参数结构体
// ============================================================================

// GetOrderListRequest 获取订单列表请求参数
type GetOrderListRequest struct {
	CreateTimeRange *OrderTimeRange `json:"create_time_range,omitempty"` // 订单创建时间范围
	UpdateTimeRange *OrderTimeRange `json:"update_time_range,omitempty"` // 订单更新时间范围
	Status          int             `json:"status,omitempty"`            // 订单状态
	OpenID          string          `json:"openid,omitempty"`            // 买家身份标识
	PageSize        int             `json:"page_size,omitempty"`         // 每页数量 (不超过 100)
	NextKey         string          `json:"next_key,omitempty"`          // 分页参数
}

// OrderTimeRange 时间范围
type OrderTimeRange struct {
	StartTime int64 `json:"start_time"` // 秒级时间戳（距离 end_time 不可超过 7 天）
	EndTime   int64 `json:"end_time"`   // 秒级时间戳（距离 start_time 不可超过 7 天）
}

// GetOrderRequest 获取订单详情请求参数
type GetOrderRequest struct {
	OrderID string `json:"order_id"` // 订单 ID
}

// SearchOrderRequest 搜索订单请求参数
type SearchOrderRequest struct {
	SearchCondition       *SearchCondition `json:"search_condition"`                   // 搜索条件
	OnAftersaleOrderExist int              `json:"on_aftersale_order_exist,omitempty"` // 是否存在售后单
	Status                int              `json:"status,omitempty"`                   // 订单状态
	PageSize              int              `json:"page_size"`                          // 每页数量 (不超过 100)
	NextKey               string           `json:"next_key"`                           // 分页参数
}

// SearchCondition 搜索条件
type SearchCondition struct {
	Title              string `json:"title,omitempty"`                // 商品标题关键词
	SKUCode            string `json:"sku_code,omitempty"`             // 商品编码
	UserName           string `json:"user_name,omitempty"`            // 收件人
	TelNumber          string `json:"tel_number,omitempty"`           // 收件人电话（已废弃）
	OrderID            string `json:"order_id,omitempty"`             // 订单 ID
	MerchantNotes      string `json:"merchant_notes,omitempty"`       // 商家备注
	CustomerNotes      string `json:"customer_notes,omitempty"`       // 买家备注
	TelNumberLast4     string `json:"tel_number_last4,omitempty"`     // 收件人电话后四位
	AddressUnderReview bool   `json:"address_under_review,omitempty"` // 申请修改地址审核中
	PresentOrderID     string `json:"present_order_id,omitempty"`     // 礼物单号
}

// ============================================================================
// 响应参数结构体
// ============================================================================

// GetOrderListResponse 获取订单列表响应
type GetOrderListResponse struct {
	ErrCode     int      `json:"errcode"`       // 错误码
	ErrMsg      string   `json:"errmsg"`        // 错误信息
	OrderIDList []string `json:"order_id_list"` // 订单号列表
	NextKey     string   `json:"next_key"`      // 分页参数
	HasMore     bool     `json:"has_more"`      // 是否还有下一页
}

// GetOrderResponse 获取订单详情响应
type GetOrderResponse struct {
	ErrCode int         `json:"errcode"` // 错误码
	ErrMsg  string      `json:"errmsg"`  // 错误信息
	Order   OrderDetail `json:"order"`   // 订单详情
}

// SearchOrderResponse 搜索订单响应
type SearchOrderResponse struct {
	ErrCode     int      `json:"errcode"`       // 错误码
	ErrMsg      string   `json:"errmsg"`        // 错误信息
	OrderIDList []string `json:"order_id_list"` // 订单号列表
	NextKey     string   `json:"next_key"`      // 分页参数
	HasMore     bool     `json:"has_more"`      // 是否还有下一页
}

// OrderDetail 订单详情
type OrderDetail struct {
	OrderID             int64           `json:"order_id"`              // 订单 ID
	CreateTime          int64           `json:"create_time"`           // 创建时间，秒级时间戳
	UpdateTime          int64           `json:"update_time"`           // 更新时间，秒级时间戳
	Status              int             `json:"status"`                // 订单状态
	OrderDetail         OrderDetailInfo `json:"order_detail"`          // 订单详细数据信息
	AftersaleDetail     AftersaleDetail `json:"aftersale_detail"`      // 售后信息
	OpenID              string          `json:"openid"`                // 订单归属人身份标识
	UnionID             string          `json:"unionid"`               // 订单归属人在开放平台的唯一标识符
	IsPresent           bool            `json:"is_present"`            // 是否礼物订单
	PresentOrderID      int64           `json:"present_order_id"`      // 礼物 id【废弃】
	PresentNote         string          `json:"present_note"`          // 礼物订单留言
	PresentGiverOpenID  string          `json:"present_giver_openid"`  // 礼物订单赠送者 openid
	PresentGiverUnionID string          `json:"present_giver_unionid"` // 礼物订单赠送者在开放平台的唯一标识符
	PresentOrderIDStr   string          `json:"present_order_id_str"`  // 礼物订单 ID
	PresentSendType     int             `json:"present_send_type"`     // 礼物单类型
	OrderPresentInfo    interface{}     `json:"order_present_info"`    // 订单对应礼物单信息
	IsFlashSaleOrder    bool            `json:"is_flash_sale_order"`   // 是否闪购订单
	IntraCityOrderInfo  interface{}     `json:"intra_city_order_info"` // 同城单信息
}

// OrderDetailInfo 订单详细数据信息
type OrderDetailInfo struct {
	ProductInfos     []ProductInfo   `json:"product_infos"`      // 商品列表
	PayInfo          PayInfo         `json:"pay_info"`           // 支付信息
	PriceInfo        PriceInfo       `json:"price_info"`         // 价格信息
	DeliveryInfo     DeliveryInfo    `json:"delivery_info"`      // 配送信息
	ExtInfo          ExtInfo         `json:"ext_info"`           // 额外信息
	CouponInfo       CouponInfo      `json:"coupon_info"`        // 优惠券信息
	CommissionInfos  []interface{}   `json:"commission_infos"`   // 分佣信息
	SharerInfo       SharerInfo      `json:"sharer_info"`        // 分享员信息【已经下线】
	SettleInfo       SettleInfo      `json:"settle_info"`        // 结算信息
	SkuSharerInfos   []SkuSharerInfo `json:"sku_sharer_infos"`   // 分享员信息【已经下线】
	AgentInfo        interface{}     `json:"agent_info"`         // 授权账号信息
	SourceInfos      []interface{}   `json:"source_infos"`       // 订单来源信息
	RefundInfo       interface{}     `json:"refund_info"`        // 订单退款信息
	GreetingCardInfo interface{}     `json:"greeting_card_info"` // 需代写的商品贺卡信息
	CustomInfo       interface{}     `json:"custom_info"`        // 商品定制信息
}

// ProductInfo 商品信息
type ProductInfo struct {
	ProductID                               int64                    `json:"product_id"`                                  // 商品 id
	SKUID                                   int64                    `json:"sku_id"`                                      // 商品 skuid
	ThumbImg                                string                   `json:"thumb_img"`                                   // sku 小图
	SalePrice                               int                      `json:"sale_price"`                                  // 售卖单价，单位为分
	SKUCnt                                  int                      `json:"sku_cnt"`                                     // sku 数量
	Title                                   string                   `json:"title"`                                       // 商品标题
	OnAftersaleSKUCnt                       int                      `json:"on_aftersale_sku_cnt"`                        // 正在售后/退款流程中的 sku 数量
	FinishAftersaleSKUCnt                   int                      `json:"finish_aftersale_sku_cnt"`                    // 完成售后/退款的 sku 数量
	SKUCode                                 string                   `json:"sku_code"`                                    // sku 编码（商家自定义编码）
	MarketPrice                             int                      `json:"market_price"`                                // 市场单价，单位为分
	SKUAttrs                                []SKUAttr                `json:"sku_attrs"`                                   // sku 属性
	RealPrice                               int                      `json:"real_price"`                                  // sku 实付总价
	OutProductID                            string                   `json:"out_product_id"`                              // 商品外部 spuid
	OutSKUID                                string                   `json:"out_sku_id"`                                  // 商品外部 skuid
	IsDiscounted                            bool                     `json:"is_discounted"`                               // 是否有商家优惠金额
	EstimatePrice                           int                      `json:"estimate_price"`                              // 使用所有优惠后 sku 总价
	IsChangePrice                           bool                     `json:"is_change_price"`                             // 是否修改过价格
	ChangePrice                             int                      `json:"change_price"`                                // 改价后 sku 总价
	OutWarehouseID                          string                   `json:"out_warehouse_id"`                            // 区域库存 id
	SKUDeliverInfo                          SKUDeliverInfo           `json:"sku_deliver_info"`                            // 商品发货信息
	ExtraService                            ExtraService             `json:"extra_service"`                               // 商品额外服务信息
	UseDeduction                            bool                     `json:"use_deduction"`                               // 是否使用了会员积分抵扣
	DeductionPrice                          int                      `json:"deduction_price"`                             // 会员积分抵扣金额，单位为分
	OrderProductCouponInfoList              []OrderProductCouponInfo `json:"order_product_coupon_info_list"`              // 商品优惠券信息
	DeliveryDeadline                        int                      `json:"delivery_deadline"`                           // 商品发货时效
	MerchantDiscountedPrice                 int                      `json:"merchant_discounted_price"`                   // 商家优惠金额，单位为分
	FinderDiscountedPrice                   int                      `json:"finder_discounted_price"`                     // 达人优惠金额，单位为分
	IsFreeGift                              int                      `json:"is_free_gift"`                                // 是否赠品，1:是赠品
	VIPDiscountedPrice                      int                      `json:"vip_discounted_price"`                        // 订单内商品维度会员权益优惠金额
	ProductUniqueID                         string                   `json:"product_unique_id"`                           // 商品常量编号
	ChangeSKUInfo                           interface{}              `json:"change_sku_info"`                             // 更换 sku 信息
	FreeGiftInfo                            interface{}              `json:"free_gift_info"`                              // 赠品信息
	BulkbuyDiscountedPrice                  int                      `json:"bulkbuy_discounted_price"`                    // 订单内商品维度一起买优惠金额
	NationalSubsidyDiscountedPrice          int                      `json:"national_subsidy_discounted_price"`           // 订单内商品维度国补优惠金额
	DropshipInfo                            interface{}              `json:"dropship_info"`                               // 代发相关信息
	IsFlashSale                             bool                     `json:"is_flash_sale"`                               // 是否闪购商品
	NationalSubsidyMerchantDiscountedPrice  int                      `json:"national_subsidy_merchant_discounted_price"`  // 地方补贴优惠金额 (商家出资)
	PlatformActivityMerchantDiscountedPrice int                      `json:"platform_activity_merchant_discounted_price"` // 平台活动商家补贴
	CashCouponDiscountedPrice               int                      `json:"cash_coupon_discounted_price"`                // 平台券优惠金额
	LimitedDiscountDiscountedPrice          int                      `json:"limited_discount_discounted_price"`           // 限时抢购优惠金额
}

// SKUAttr SKU 属性
type SKUAttr struct {
	AttrKey   string `json:"attr_key"`   // 属性键
	AttrValue string `json:"attr_value"` // 属性值
}

// SKUDeliverInfo 商品发货信息
type SKUDeliverInfo struct {
	StockType                      int `json:"stock_type"`                         // 商品发货类型：0：现货，1：全款预售
	PredictDeliveryTime            int `json:"predict_delivery_time"`              // 预计发货时间 (stock_type=1 时返回该字段)
	FullPaymentPresaleDeliveryType int `json:"full_payment_presale_delivery_type"` // 预售类型 0:付款后 n 天发货，1:预售结束后 n 天发货
}

// ExtraService 商品额外服务信息
type ExtraService struct {
	SevenDayReturn   int `json:"seven_day_return"`  // 7 天无理由：0：不支持，1：支持
	FreightInsurance int `json:"freight_insurance"` // 商家运费险：0：不支持，1：支持
}

// OrderProductCouponInfo 商品优惠券信息
type OrderProductCouponInfo struct {
	UserCouponID    string `json:"user_coupon_id"`   // 用户优惠券 id
	CouponType      int    `json:"coupon_type"`      // 优惠券类型
	DiscountedPrice int    `json:"discounted_price"` // 优惠金额，单位为分
	CouponID        string `json:"coupon_id"`        // 优惠券 id
}

// ============================================================================
// 订单详情子结构体
// ============================================================================

// AftersaleDetail 售后信息
type AftersaleDetail struct {
	AftersaleOrderList  []AftersaleOrder `json:"aftersale_order_list"`   // 售后单列表
	OnAftersaleOrderCnt int              `json:"on_aftersale_order_cnt"` // 正在售后中的售后单数量
}

// AftersaleOrder 售后单
type AftersaleOrder struct {
	AftersaleOrderID string `json:"aftersale_order_id"` // 售后单 ID
	Status           int    `json:"status"`             // 售后单状态
}

// PayInfo 支付信息
type PayInfo struct {
	PrepayID      string `json:"prepay_id"`      // 预支付 ID
	TransactionID string `json:"transaction_id"` // 交易 ID
	PrepayTime    int64  `json:"prepay_time"`    // 预支付时间，秒级时间戳
	PayTime       int64  `json:"pay_time"`       // 支付时间，秒级时间戳
	PaymentMethod int    `json:"payment_method"` // 支付方式
}

// PriceInfo 价格信息
type PriceInfo struct {
	ProductPrice    int  `json:"product_price"`    // 商品总价，单位为分
	OrderPrice      int  `json:"order_price"`      // 订单总价，单位为分
	Freight         int  `json:"freight"`          // 运费，单位为分
	DiscountedPrice int  `json:"discounted_price"` // 优惠金额，单位为分
	IsDiscounted    bool `json:"is_discounted"`    // 是否有优惠
}

// DeliveryInfo 配送信息
type DeliveryInfo struct {
	AddressInfo         AddressInfo           `json:"address_info"`          // 地址信息
	DeliveryProductInfo []DeliveryProductInfo `json:"delivery_product_info"` // 配送商品信息
	ShipDoneTime        int64                 `json:"ship_done_time"`        // 发货完成时间，秒级时间戳
	DeliverMethod       int                   `json:"deliver_method"`        // 配送方式
}

// AddressInfo 地址信息
type AddressInfo struct {
	UserName     string `json:"user_name"`     // 收件人姓名
	PostalCode   string `json:"postal_code"`   // 邮政编码
	ProvinceName string `json:"province_name"` // 省份
	CityName     string `json:"city_name"`     // 城市
	CountyName   string `json:"county_name"`   // 区县
	DetailInfo   string `json:"detail_info"`   // 详细地址
	TelNumber    string `json:"tel_number"`    // 联系电话
}

// DeliveryProductInfo 配送商品信息
type DeliveryProductInfo struct {
	WaybillID    string                    `json:"waybill_id"`    // 运单 ID
	DeliveryID   string                    `json:"delivery_id"`   // 物流公司 ID
	DeliveryTime int64                     `json:"delivery_time"` // 发货时间，秒级时间戳
	DeliverType  int                       `json:"deliver_type"`  // 发货类型
	ProductInfos []DeliveryProductItemInfo `json:"product_infos"` // 商品信息
}

// DeliveryProductItemInfo 配送商品信息项
type DeliveryProductItemInfo struct {
	ProductID  string `json:"product_id"`  // 商品 ID
	SKUID      string `json:"sku_id"`      // SKU ID
	ProductCnt int    `json:"product_cnt"` // 商品数量
}

// ExtInfo 额外信息
type ExtInfo struct {
	CustomerNotes string `json:"customer_notes"` // 买家备注
	MerchantNotes string `json:"merchant_notes"` // 商家备注
	FinderID      string `json:"finder_id"`      // 视频号 ID
	LiveID        string `json:"live_id"`        // 直播间 ID
	OrderScene    int    `json:"order_scene"`    // 订单场景
}

// CouponInfo 优惠券信息
type CouponInfo struct {
	UserCouponID string `json:"user_coupon_id"` // 用户优惠券 ID
}

// SharerInfo 分享员信息
type SharerInfo struct {
	SharerOpenid     string `json:"sharer_openid"`     // 分享者 openid
	SharerUnionid    string `json:"sharer_unionid"`    // 分享者 unionid
	SharerType       int    `json:"sharer_type"`       // 分享者类型
	ShareScene       int    `json:"share_scene"`       // 分享场景
	HandlingProgress int    `json:"handling_progress"` // 处理进度
}

// SettleInfo 结算信息
type SettleInfo struct {
	CommissionFee        int `json:"commission_fee"`         // 佣金，单位为分
	PredictCommissionFee int `json:"predict_commission_fee"` // 预估佣金，单位为分
}

// SkuSharerInfo SKU 分享员信息
type SkuSharerInfo struct {
	SharerOpenid  string `json:"sharer_openid"`  // 分享者 openid
	SharerUnionid string `json:"sharer_unionid"` // 分享者 unionid
	SharerType    int    `json:"sharer_type"`    // 分享者类型
	ShareScene    int    `json:"share_scene"`    // 分享场景
	SKUID         string `json:"sku_id"`         // SKU ID
}
