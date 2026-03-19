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
	// 资金相关 API 接口路径
	FundsGetBalanceApi         = "/channels/ec/funds/getbalance"         // 获取账户余额
	FundsGetBankAcctApi        = "/channels/ec/funds/getbankacct"        // 获取结算账户
	FundsGetFundsFlowDetailApi = "/channels/ec/funds/getfundsflowdetail" // 获取资金流水详情
	FundsGetFundsFlowListApi   = "/channels/ec/funds/getfundsflowlist"   // 获取资金流水列表
	FundsGetWithdrawListApi    = "/channels/ec/funds/getwithdrawlist"    // 获取提现记录列表
	FundsListOrderFlowApi      = "/channels/ec/funds/listorderflow"      // 查询订单流水列表
)

// GetFundsBalance 获取账户余额
func (c *Client) GetFundsBalance(ctx context.Context) (*GetFundsBalanceResponse, error) {
	if c.AccessToken == "" {
		return nil, fmt.Errorf("access_token 不能为空")
	}

	query := url.Values{}
	query.Set("access_token", c.AccessToken)

	bodyBytes, err := json.Marshal(map[string]interface{}{})
	if err != nil {
		return nil, fmt.Errorf("序列化请求参数失败：%w", err)
	}

	reqURL := c.Env + FundsGetBalanceApi + "?" + query.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("构建获取账户余额请求失败：%w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求获取账户余额失败：%w", err)
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败：%w", err)
	}

	var result GetFundsBalanceResponse
	if err = json.Unmarshal(respBytes, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败：%w", err)
	}

	return &result, nil
}

// GetFundsBankAcct 获取结算账户
func (c *Client) GetFundsBankAcct(ctx context.Context) (*GetFundsBankAcctResponse, error) {
	if c.AccessToken == "" {
		return nil, fmt.Errorf("access_token 不能为空")
	}

	query := url.Values{}
	query.Set("access_token", c.AccessToken)

	bodyBytes, err := json.Marshal(map[string]interface{}{})
	if err != nil {
		return nil, fmt.Errorf("序列化请求参数失败：%w", err)
	}

	reqURL := c.Env + FundsGetBankAcctApi + "?" + query.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("构建获取结算账户请求失败：%w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求获取结算账户失败：%w", err)
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败：%w", err)
	}

	var result GetFundsBankAcctResponse
	if err = json.Unmarshal(respBytes, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败：%w", err)
	}

	return &result, nil
}

// GetFundsFlowDetail 获取资金流水详情
func (c *Client) GetFundsFlowDetail(ctx context.Context, reqData *GetFundsFlowDetailRequest) (*GetFundsFlowDetailResponse, error) {
	if c.AccessToken == "" {
		return nil, fmt.Errorf("access_token 不能为空")
	}
	if reqData == nil || reqData.FlowID == "" {
		return nil, fmt.Errorf("flow_id 不能为空")
	}

	query := url.Values{}
	query.Set("access_token", c.AccessToken)

	bodyBytes, err := json.Marshal(reqData)
	if err != nil {
		return nil, fmt.Errorf("序列化请求参数失败：%w", err)
	}

	reqURL := c.Env + FundsGetFundsFlowDetailApi + "?" + query.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("构建获取资金流水详情请求失败：%w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求获取资金流水详情失败：%w", err)
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败：%w", err)
	}

	var result GetFundsFlowDetailResponse
	if err = json.Unmarshal(respBytes, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败：%w", err)
	}

	return &result, nil
}

// GetFundsFlowList 获取资金流水列表
func (c *Client) GetFundsFlowList(ctx context.Context, reqData *GetFundsFlowListRequest) (*GetFundsFlowListResponse, error) {
	if c.AccessToken == "" {
		return nil, fmt.Errorf("access_token 不能为空")
	}
	if reqData == nil {
		reqData = &GetFundsFlowListRequest{}
	}

	query := url.Values{}
	query.Set("access_token", c.AccessToken)

	bodyBytes, err := json.Marshal(reqData)
	if err != nil {
		return nil, fmt.Errorf("序列化请求参数失败：%w", err)
	}

	reqURL := c.Env + FundsGetFundsFlowListApi + "?" + query.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("构建获取资金流水列表请求失败：%w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求获取资金流水列表失败：%w", err)
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败：%w", err)
	}

	var result GetFundsFlowListResponse
	if err = json.Unmarshal(respBytes, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败：%w", err)
	}

	return &result, nil
}

// GetFundsWithdrawList 获取提现记录列表
func (c *Client) GetFundsWithdrawList(ctx context.Context, reqData *GetFundsWithdrawListRequest) (*GetFundsWithdrawListResponse, error) {
	if c.AccessToken == "" {
		return nil, fmt.Errorf("access_token 不能为空")
	}
	if reqData == nil {
		return nil, fmt.Errorf("请求参数不能为空")
	}
	if reqData.PageNum <= 0 {
		return nil, fmt.Errorf("page_num 必须大于 0")
	}
	if reqData.PageSize <= 0 {
		return nil, fmt.Errorf("page_size 必须大于 0")
	}

	query := url.Values{}
	query.Set("access_token", c.AccessToken)

	bodyBytes, err := json.Marshal(reqData)
	if err != nil {
		return nil, fmt.Errorf("序列化请求参数失败：%w", err)
	}

	reqURL := c.Env + FundsGetWithdrawListApi + "?" + query.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("构建获取提现记录列表请求失败：%w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求获取提现记录列表失败：%w", err)
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败：%w", err)
	}

	var result GetFundsWithdrawListResponse
	if err = json.Unmarshal(respBytes, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败：%w", err)
	}

	return &result, nil
}

// ListOrderFlow 查询订单流水列表
func (c *Client) ListOrderFlow(ctx context.Context, reqData *ListOrderFlowRequest) (*ListOrderFlowResponse, error) {
	if c.AccessToken == "" {
		return nil, fmt.Errorf("access_token 不能为空")
	}
	if reqData == nil {
		return nil, fmt.Errorf("请求参数不能为空")
	}
	if reqData.OrderSettleState == 0 && reqData.OrderID == "" {
		return nil, fmt.Errorf("order_settle_state 或 order_id 至少传一个")
	}

	query := url.Values{}
	query.Set("access_token", c.AccessToken)

	bodyBytes, err := json.Marshal(reqData)
	if err != nil {
		return nil, fmt.Errorf("序列化请求参数失败：%w", err)
	}

	reqURL := c.Env + FundsListOrderFlowApi + "?" + query.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("构建查询订单流水列表请求失败：%w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求查询订单流水列表失败：%w", err)
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败：%w", err)
	}

	var result ListOrderFlowResponse
	if err = json.Unmarshal(respBytes, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败：%w", err)
	}

	return &result, nil
}

// ==================== 请求结构体 ====================

// GetFundsFlowDetailRequest 获取资金流水详情请求参数
type GetFundsFlowDetailRequest struct {
	FlowID string `json:"flow_id"` // 流水 id，可通过获取资金流水列表获取
}

// GetFundsFlowListRequest 获取资金流水列表请求参数
type GetFundsFlowListRequest struct {
	Page          int    `json:"page,omitempty"`           // 页码，从 1 开始
	PageSize      int    `json:"page_size,omitempty"`      // 页数，不填默认为 10
	NextKey       string `json:"next_key,omitempty"`       // 分页参数，翻页时写入上一页返回的 next_key
	StartTime     int64  `json:"start_time,omitempty"`     // 流水产生的开始时间，unix 时间戳
	EndTime       int64  `json:"end_time,omitempty"`       // 流水产生的结束时间，unix 时间戳
	TransactionID string `json:"transaction_id,omitempty"` // 支付单号
}

// GetFundsWithdrawListRequest 获取提现记录列表请求参数
type GetFundsWithdrawListRequest struct {
	PageNum   int   `json:"page_num"`             // 页码
	PageSize  int   `json:"page_size"`            // 每页大小
	StartTime int64 `json:"start_time,omitempty"` // 开始时间
	EndTime   int64 `json:"end_time,omitempty"`   // 结束时间
}

// ListOrderFlowRequest 查询订单流水列表请求参数
type ListOrderFlowRequest struct {
	OrderSettleState int                      `json:"order_settle_state"`          // 订单结算状态
	OrderState       int                      `json:"order_state,omitempty"`       // 订单状态
	OrderPayMethod   int                      `json:"order_pay_method,omitempty"`  // 订单支付方式
	OrderID          string                   `json:"order_id,omitempty"`          // 指定订单 id 查询
	PaginationInfo   OrderFlowPaginationInfo  `json:"pagination_info,omitempty"`   // 分页信息
	CreateTimeRange  OrderFlowCreateTimeRange `json:"create_time_range,omitempty"` // 订单创建时间范围
}

// OrderFlowPaginationInfo 分页信息
type OrderFlowPaginationInfo struct {
	Limit      int    `json:"limit"`                  // 页大小
	Offset     int    `json:"offset,omitempty"`       // 偏移量
	UsePageCtx bool   `json:"use_page_ctx,omitempty"` // 是否使用分页上下文
	PageCtx    string `json:"page_ctx,omitempty"`     // 分页上下文
}

// OrderFlowCreateTimeRange 订单创建时间范围
type OrderFlowCreateTimeRange struct {
	Begin int64 `json:"begin,omitempty"` // 订单创建时间开始（闭区间）
	End   int64 `json:"end,omitempty"`   // 订单创建时间结束（开区间）
}

// ==================== 响应结构体 ====================

// GetFundsBalanceResponse 获取账户余额响应
type GetFundsBalanceResponse struct {
	ErrCode         int    `json:"errcode"`          // 错误码
	ErrMsg          string `json:"errmsg"`           // 错误信息
	AvailableAmount int    `json:"available_amount"` // 可提现余额
	PendingAmount   int    `json:"pending_amount"`   // 待结算余额
	SubMchid        string `json:"sub_mchid"`        // 二级商户号
}

// GetFundsBankAcctResponse 获取结算账户响应
type GetFundsBankAcctResponse struct {
	ErrCode     int             `json:"errcode"`      // 错误码
	ErrMsg      string          `json:"errmsg"`       // 错误信息
	AccountInfo BankAccountInfo `json:"account_info"` // 账户信息
}

// BankAccountInfo 银行账户信息
type BankAccountInfo struct {
	BankAccountType string `json:"bank_account_type"` // 账户类型
	AccountBank     string `json:"account_bank"`      // 开户银行
	BankAddressCode string `json:"bank_address_code"` // 开户银行省市编码
	BankBranchID    string `json:"bank_branch_id"`    // 开户银行联行号
	BankName        string `json:"bank_name"`         // 开户银行全称
	AccountNumber   string `json:"account_number"`    // 银行账号
	AccountName     string `json:"account_name"`      // 账户名称
}

// GetFundsFlowDetailResponse 获取资金流水详情响应
type GetFundsFlowDetailResponse struct {
	ErrCode   int                `json:"errcode"`    // 错误码
	ErrMsg    string             `json:"errmsg"`     // 错误信息
	FundsFlow FundsFlowDetailDTO `json:"funds_flow"` // 流水信息
}

// FundsFlowDetailDTO 资金流水详情
type FundsFlowDetailDTO struct {
	FlowID          string            `json:"flow_id"`           // 流水 id
	FundsType       int               `json:"funds_type"`        // 资金类型
	FlowType        int               `json:"flow_type"`         // 流水类型，1 收入，2 支出
	Amount          int               `json:"amount"`            // 流水金额
	Balance         int               `json:"balance"`           // 余额
	RelatedInfoList []FundRelatedInfo `json:"related_info_list"` // 流水关联信息
	BookkeepingTime string            `json:"bookkeeping_time"`  // 记账时间
	Remark          string            `json:"remark"`            // 备注
	FundsTypeDesc   string            `json:"funds_type_desc"`   // 资金类型描述
}

// FundRelatedInfo 流水关联信息
type FundRelatedInfo struct {
	RelatedType                int    `json:"related_type"`                              // 关联类型
	OrderID                    string `json:"order_id,omitempty"`                        // 关联订单号
	AftersaleID                string `json:"aftersale_id,omitempty"`                    // 关联售后单号
	WithdrawID                 string `json:"withdraw_id,omitempty"`                     // 关联提现单号
	BookkeepingTime            string `json:"bookkeeping_time,omitempty"`                // 记账时间
	InsuranceID                string `json:"insurance_id,omitempty"`                    // 关联运费险单号
	TransactionID              string `json:"transaction_id,omitempty"`                  // 关联支付单号
	GuaranteeID                int    `json:"guarantee_id,omitempty"`                    // 关联保障单号
	PresentID                  int    `json:"present_id,omitempty"`                      // 关联礼物单号
	GroupPresentSubOrderIDList string `json:"group_present_sub_order_id_list,omitempty"` // 群送礼关联订单号列表
	IntraCityShopID            int    `json:"intra_city_shop_id,omitempty"`              // 同城配送门店 id
}

// GetFundsFlowListResponse 获取资金流水列表响应
type GetFundsFlowListResponse struct {
	ErrCode int      `json:"errcode"`  // 错误码
	ErrMsg  string   `json:"errmsg"`   // 错误信息
	FlowIDs []string `json:"flow_ids"` // 流水单号列表
	HasMore bool     `json:"has_more"` // 是否还有下一页
	NextKey string   `json:"next_key"` // 分页参数，深翻页时使用
}

// GetFundsWithdrawListResponse 获取提现记录列表响应
type GetFundsWithdrawListResponse struct {
	ErrCode     int      `json:"errcode"`      // 错误码
	ErrMsg      string   `json:"errmsg"`       // 错误信息
	WithdrawIDs []string `json:"withdraw_ids"` // 提现单号列表
	TotalNum    int      `json:"total_num"`    // 提现单号总数
}

// ListOrderFlowResponse 查询订单流水列表响应
type ListOrderFlowResponse struct {
	ErrCode    int                     `json:"errcode"`     // 错误码
	ErrMsg     string                  `json:"errmsg"`      // 错误信息
	TotalCount int                     `json:"total_count"` // 满足条件的总数量
	DataList   []OrderFlowDataListItem `json:"data_list"`   // 订单流水列表
	PageCtx    string                  `json:"page_ctx"`    // 分页上下文
}

// OrderFlowDataListItem 订单流水列表项
type OrderFlowDataListItem struct {
	OrderID                           string                 `json:"order_id"`                                // 订单 id
	OrderState                        int                    `json:"order_state"`                             // 订单状态
	OrderSettleState                  int                    `json:"order_settle_state"`                      // 订单结算状态
	OrderCreateTime                   int64                  `json:"order_create_time"`                       // 订单创建时间
	OrderPaidTime                     int64                  `json:"order_paid_time"`                         // 订单支付时间
	OrderPayMethod                    int                    `json:"order_pay_method"`                        // 订单支付方式
	OrderType                         int                    `json:"order_type"`                              // 订单类型
	MchReceivedAmount                 int                    `json:"mch_received_amount"`                     // 商户实收金额
	ExpenseAmount                     int                    `json:"expense_amount"`                          // 支出金额
	MchSettleAmount                   int                    `json:"mch_settle_amount"`                       // (预计) 结算金额
	MchSettleTime                     int64                  `json:"mch_settle_time"`                         // (预计) 商家货款结算时间
	ProductList                       []OrderFlowProductItem `json:"product_list"`                            // 商品列表
	ProductTotalAmount                int                    `json:"product_total_amount"`                    // 商品总金额
	FreightAmount                     int                    `json:"freight_amount"`                          // 运费金额
	ChangeDownPrice                   int                    `json:"change_down_price"`                       // 改价金额
	MchDiscountAmount                 int                    `json:"mch_discount_amount"`                     // 商户优惠金额
	ScoreDiscountAmount               int                    `json:"score_discount_amount"`                   // 积分抵扣金额
	BuyerPaidAmount                   int                    `json:"buyer_paid_amount"`                       // 用户实付金额
	PromoterDiscountAmount            int                    `json:"promoter_discount_amount"`                // 达人优惠金额
	PlatformDiscountAmount            int                    `json:"platform_discount_amount"`                // 平台优惠金额
	NationalSubsidyDiscountAmount     int                    `json:"national_subsidy_discount_amount"`        // 国家补贴金额
	FreightMakeUpAmount               int                    `json:"freight_make_up_amount"`                  // 补交运费
	CrossShopDiscountAmount           int                    `json:"cross_shop_discount_amount"`              // 跨店优惠金额
	BuyerRefundAmount                 int                    `json:"buyer_refund_amount"`                     // 用户退款金额
	PlatformDiscountRefundAmount      int                    `json:"platform_discount_refund_amount"`         // (预计) 平台优惠退款金额
	PromoterDiscountRefundAmount      int                    `json:"promoter_discount_refund_amount"`         // 达人优惠退款金额
	OriginalPlatformCommissionAmount  int                    `json:"original_platform_commission_amount"`     // 原技术服务费
	PlatformCommissionAmount          int                    `json:"platform_commission_amount"`              // (预计) 技术服务费
	FreightInsuranceSubsidyAmount     int                    `json:"freight_insurance_subsidy_amount"`        // 运费险补贴减免技术服务费
	SupplierCommissionAmount          int                    `json:"supplier_commission_amount"`              // (预计) 机构服务费
	SupplierCommissionSettleState     int                    `json:"supplier_commission_settle_state"`        // 机构服务费结算状态
	PromoterCommissionAmount          int                    `json:"promoter_commission_amount"`              // (预计) 达人服务费
	PromoterCommissionSettleState     int                    `json:"promoter_commission_settle_state"`        // 达人服务费结算状态
	FreightInsuranceAmount            int                    `json:"freight_insurance_amount"`                // (预计) 运费险金额
	FreightInsuranceSettleState       int                    `json:"freight_insurance_settle_state"`          // 运费险结算状态
	FreightInsuranceMakeUpAmount      int                    `json:"freight_insurance_make_up_amount"`        // 运费险补缴他单金额
	FreightInsuranceMakeUpOrderIDList []string               `json:"freight_insurance_make_up_order_id_list"` // 运费险补缴他单订单 id 列表
	PlatformCommissionSettleTime      int64                  `json:"platform_commission_settle_time"`         // (预计) 技术服务费结算时间
	PromoterCommissionSettleTime      int64                  `json:"promoter_commission_settle_time"`         // (预计) 达人服务费结算时间
	SupplierCommissionSettleTime      int64                  `json:"supplier_commission_settle_time"`         // (预计) 机构服务费结算时间
	FreightInsuranceSettleTime        int64                  `json:"freight_insurance_settle_time"`           // (预计) 运费险结算时间
	FreightInsuranceMakeUpSettleTime  int64                  `json:"freight_insurance_make_up_settle_time"`   // (预计) 运费险补缴他单结算时间
	PreFreightRefundAmount            int                    `json:"pre_freight_refund_amount"`               // 预付运费退回金额
	PostSettlementExpense             PostSettlementExpense  `json:"post_settlement_expense"`                 // 结算后支出
	RefundBeforeSettlement            int                    `json:"refund_before_settlement"`                // 结算前退款
	OtherExpenseAmount                int                    `json:"other_expense_amount"`                    // 其他支出
	PlatformCommissionSettleState     int                    `json:"platform_commission_settle_state"`        // 平台服务费结算状态
	FreightInsuranceMakeUpSettleState int                    `json:"freight_insurance_make_up_settle_state"`  // 运费险补缴他单结算状态
	IntraCityShopID                   int                    `json:"intra_city_shop_id"`                      // 同城配送门店 id
}

// OrderFlowProductItem 订单流水商品项
type OrderFlowProductItem struct {
	ProductID   int           `json:"product_id"`   // 商品 id
	ParamList   []ProductSpec `json:"param_list"`   // 商品规格列表
	SalePrice   int           `json:"sale_price"`   // 商品销售价格
	Count       int           `json:"count"`        // 商品数量
	ProductName string        `json:"product_name"` // 商品名称
	IsGift      bool          `json:"is_gift"`      // 是否赠品
}

// ProductSpec 商品规格
type ProductSpec struct {
	Key   string `json:"key"`   // 规格名称
	Value string `json:"value"` // 规格值
}

// PostSettlementExpense 结算后支出
type PostSettlementExpense struct {
	BuyerRefundAmount                 int `json:"buyer_refund_amount"`                    // 买家退款
	PlatformDiscountRefundAmount      int `json:"platform_discount_refund_amount"`        // 平台优惠退款金额
	PromoterRefundAmount              int `json:"promoter_refund_amount"`                 // 达人退款金额
	FreightInsuranceMakeUpAmount      int `json:"freight_insurance_make_up_amount"`       // 运费险保费
	FreightInsuranceMakeUpSettleState int `json:"freight_insurance_make_up_settle_state"` // 运费险补缴本单结算状态
	FreightInsuranceMakeUpOrderID     int `json:"freight_insurance_make_up_order_id"`     // 运费险补缴本单的订单 id
}
