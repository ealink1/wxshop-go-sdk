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
//type GetProductResponse struct {
//	ErrCode       int             `json:"errcode"`         // 错误码
//	ErrMsg        string          `json:"errmsg"`          // 错误信息
//	Product       json.RawMessage `json:"product"`         // 商品线上数据
//	EditProduct   json.RawMessage `json:"edit_product"`    // 商品草稿数据
//	SaleLimitInfo json.RawMessage `json:"sale_limit_info"` // 售卖上限提醒
//	InfoScore     json.RawMessage `json:"info_score"`      // 商品信息质量
//	CmpPriceInfo  json.RawMessage `json:"cmp_price_info"`  // 商品高价预警
//	AuditInfo     json.RawMessage `json:"audit_info"`      // 审核信息
//}

type GetProductResponse struct {
	Errcode int    `json:"errcode"` // 错误码
	Errmsg  string `json:"errmsg"`  // 错误信息
	Product struct {
		ProductId    string   `json:"product_id"`     // 小店内部商品 ID
		OutProductId string   `json:"out_product_id"` // 外部平台自定义商品 ID
		Title        string   `json:"title"`          // 商品标题
		SubTitle     string   `json:"sub_title"`      // 商品副标题（已废弃）
		HeadImgs     []string `json:"head_imgs"`      // 商品主图列表
		DescInfo     struct {
			Imgs []string `json:"imgs"` // 详情图片列表
			Desc string   `json:"desc"` // 详情文本
		} `json:"desc_info"` // 商品详情
		Cats []struct {
			CatId string `json:"cat_id"` // 旧类目 ID
		} `json:"cats"` // 旧类目树
		Attrs []struct {
			AttrKey   string `json:"attr_key"`   // 属性键
			AttrValue string `json:"attr_value"` // 属性值
		} `json:"attrs"` // 商品属性
		ExpressInfo struct {
			TemplateId string `json:"template_id"` // 运费模板 ID
			Weight     int    `json:"weight"`      // 重量（克）
		} `json:"express_info"` // 运费信息
		Status int `json:"status"` // 商品线上状态
		Skus   []struct {
			SkuId     string `json:"sku_id"`     // SKU ID
			OutSkuId  string `json:"out_sku_id"` // 外部平台 SKU ID
			ThumbImg  string `json:"thumb_img"`  // SKU 图片
			SalePrice int    `json:"sale_price"` // 售价（分）
			StockNum  int    `json:"stock_num"`  // 库存
			SkuCode   string `json:"sku_code"`   // 商家自定义 SKU 编码
			SkuAttrs  []struct {
				AttrKey   string `json:"attr_key"`   // 销售属性键
				AttrValue string `json:"attr_value"` // 销售属性值
			} `json:"sku_attrs"` // SKU 销售属性
			Status         int `json:"status"` // SKU 状态
			SkuDeliverInfo struct {
				StockType                      int `json:"stock_type"`                         // 库存类型
				FullPaymentPresaleDeliveryType int `json:"full_payment_presale_delivery_type"` // 全款预售发货类型
				PresaleBeginTime               int `json:"presale_begin_time"`                 // 预售开始时间
				PresaleEndTime                 int `json:"presale_end_time"`                   // 预售结束时间
				FullPaymentPresaleDeliveryTime int `json:"full_payment_presale_delivery_time"` // 全款预售发货时间
				SpotAfterPresaleEnd            int `json:"spot_after_presale_end"`             // 预售结束后是否转现货
			} `json:"sku_deliver_info"` // SKU 发货配置
			BarCode string `json:"bar_code"` // 商品条形码
		} `json:"skus"` // SKU 列表
		MinPrice      int    `json:"min_price"`      // SKU 最低价（分）
		SpuCode       string `json:"spu_code"`       // 商家自定义 SPU 编码
		DeliverMethod int    `json:"deliver_method"` // 发货方式
		AftersaleDesc string `json:"aftersale_desc"` // 售后说明
		LimitedInfo   struct {
			PeriodType    int `json:"period_type"`     // 限购周期类型
			LimitedBuyNum int `json:"limited_buy_num"` // 限购数量
		} `json:"limited_info"` // 限购信息
		BrandId        string        `json:"brand_id"`       // 品牌 ID，无品牌为 2100000000
		Qualifications []interface{} `json:"qualifications"` // 商品资质（旧字段）
		ExtraService   struct {
			SevenDayReturn   int `json:"seven_day_return"`   // 7 天无理由
			PayAfterUse      int `json:"pay_after_use"`      // 先用后付
			FreightInsurance int `json:"freight_insurance"`  // 运费险
			DamageGuarantee  int `json:"damage_guarantee"`   // 破损包退换
			FakeOnePayThree  int `json:"fake_one_pay_three"` // 假一赔三
			ExchangeSupport  int `json:"exchange_support"`   // 换货保障
		} `json:"extra_service"` // 额外服务
		ProductType   int `json:"product_type"` // 商品类型
		EditTime      int `json:"edit_time"`    // 草稿最近修改时间（秒级时间戳）
		AfterSaleInfo struct {
			AfterSaleAddressId string `json:"after_sale_address_id"` // 售后地址 ID
		} `json:"after_sale_info"` // 售后信息
		HideInWindow    int `json:"hide_in_window"` // 是否在店铺首页隐藏
		ProductQuaInfos []struct {
			QuaId  string   `json:"qua_id"`  // 资质 ID
			QuaUrl []string `json:"qua_url"` // 资质文件 URL 列表
		} `json:"product_qua_infos"` // 商品资质列表
		SizeChart struct {
			Enable            bool          `json:"enable"`             // 是否启用尺码表
			SpecificationList []interface{} `json:"specification_list"` // 尺码规格信息
		} `json:"size_chart"` // 尺码信息
		DeliverAcctType []interface{} `json:"deliver_acct_type"` // 无需快递时的发货账号类型
		DomainType      int           `json:"domain_type"`       // 领域类型
		CustomConfig    struct {
			CustomType           []interface{} `json:"custom_type"`             // 定制类型
			CustomTextMaxLength  int           `json:"custom_text_max_length"`  // 定制文案最大长度
			CustomTextInputType  int           `json:"custom_text_input_type"`  // 定制输入类型
			CustomTextDirection  int           `json:"custom_text_direction"`   // 文案方向
			CustomTextFontSize   int           `json:"custom_text_font_size"`   // 文案字号
			CustomTextColor      string        `json:"custom_text_color"`       // 文案颜色
			CustomTextInputTypes []interface{} `json:"custom_text_input_types"` // 定制输入类型集合
			OpenCustom           bool          `json:"open_custom"`             // 是否开启定制
			DescImgList          []interface{} `json:"desc_img_list"`           // 定制说明图
			DeliveryTime         int           `json:"delivery_time"`           // 交付时间
			PreviewType          int           `json:"preview_type"`            // 预览类型
		} `json:"custom_config"` // 定制化配置
		ShortTitle     string        `json:"short_title"`    // 短标题
		TotalSoldNum   int           `json:"total_sold_num"` // 销量
		ReleaseMode    int           `json:"release_mode"`   // 发布模式
		HeadVideos     []interface{} `json:"head_videos"`    // 主图视频列表
		SpuDeliverInfo struct {
			SkuDeliverInfo struct {
				StockType                      int `json:"stock_type"`                         // 库存类型
				FullPaymentPresaleDeliveryType int `json:"full_payment_presale_delivery_type"` // 全款预售发货类型
				PresaleBeginTime               int `json:"presale_begin_time"`                 // 预售开始时间
				PresaleEndTime                 int `json:"presale_end_time"`                   // 预售结束时间
				FullPaymentPresaleDeliveryTime int `json:"full_payment_presale_delivery_time"` // 全款预售发货时间
				SpotAfterPresaleEnd            int `json:"spot_after_presale_end"`             // 预售结束后是否转现货
			} `json:"sku_deliver_info"` // SPU 维度发货配置
			IsSpuRange int `json:"is_spu_range"` // 是否按 SPU 维度生效
		} `json:"spu_deliver_info"` // SPU 维度预售配置
		CatsV2 []struct {
			CatId string `json:"cat_id"` // 新类目 ID
		} `json:"cats_v2"` // 新类目树
		TimingOnsaleInfo struct {
			Status      int `json:"status"`        // 待开售状态
			OnsaleTime  int `json:"onsale_time"`   // 开售时间（秒级时间戳）
			IsHidePrice int `json:"is_hide_price"` // 开售前是否隐藏价格
			TaskId      int `json:"task_id"`       // 定时任务 ID
		} `json:"timing_onsale_info"` // 待开售信息
	} `json:"product"` // 商品线上数据
	EditProduct struct {
		ProductId    string   `json:"product_id"`     // 小店内部商品 ID
		OutProductId string   `json:"out_product_id"` // 外部平台自定义商品 ID
		Title        string   `json:"title"`          // 商品标题
		SubTitle     string   `json:"sub_title"`      // 商品副标题（已废弃）
		HeadImgs     []string `json:"head_imgs"`      // 商品主图列表
		DescInfo     struct {
			Imgs []string `json:"imgs"` // 详情图片列表
			Desc string   `json:"desc"` // 详情文本
		} `json:"desc_info"` // 商品详情
		Cats []struct {
			CatId string `json:"cat_id"` // 旧类目 ID
		} `json:"cats"` // 旧类目树
		Attrs []struct {
			AttrKey   string `json:"attr_key"`   // 属性键
			AttrValue string `json:"attr_value"` // 属性值
		} `json:"attrs"` // 商品属性
		ExpressInfo struct {
			TemplateId string `json:"template_id"` // 运费模板 ID
			Weight     int    `json:"weight"`      // 重量（克）
		} `json:"express_info"` // 运费信息
		Status     int `json:"status"`      // 商品线上状态
		EditStatus int `json:"edit_status"` // 商品草稿状态
		Skus       []struct {
			SkuId     string `json:"sku_id"`     // SKU ID
			OutSkuId  string `json:"out_sku_id"` // 外部平台 SKU ID
			ThumbImg  string `json:"thumb_img"`  // SKU 图片
			SalePrice int    `json:"sale_price"` // 售价（分）
			StockNum  int    `json:"stock_num"`  // 库存
			SkuCode   string `json:"sku_code"`   // 商家自定义 SKU 编码
			SkuAttrs  []struct {
				AttrKey   string `json:"attr_key"`   // 销售属性键
				AttrValue string `json:"attr_value"` // 销售属性值
			} `json:"sku_attrs"` // SKU 销售属性
			Status         int `json:"status"` // SKU 状态
			SkuDeliverInfo struct {
				StockType                      int `json:"stock_type"`                         // 库存类型
				FullPaymentPresaleDeliveryType int `json:"full_payment_presale_delivery_type"` // 全款预售发货类型
				PresaleBeginTime               int `json:"presale_begin_time"`                 // 预售开始时间
				PresaleEndTime                 int `json:"presale_end_time"`                   // 预售结束时间
				FullPaymentPresaleDeliveryTime int `json:"full_payment_presale_delivery_time"` // 全款预售发货时间
				SpotAfterPresaleEnd            int `json:"spot_after_presale_end"`             // 预售结束后是否转现货
			} `json:"sku_deliver_info"` // SKU 发货配置
			BarCode string `json:"bar_code"` // 商品条形码
		} `json:"skus"` // SKU 列表
		MinPrice      int    `json:"min_price"`      // SKU 最低价（分）
		SpuCode       string `json:"spu_code"`       // 商家自定义 SPU 编码
		DeliverMethod int    `json:"deliver_method"` // 发货方式
		AftersaleDesc string `json:"aftersale_desc"` // 售后说明
		LimitedInfo   struct {
			PeriodType    int `json:"period_type"`     // 限购周期类型
			LimitedBuyNum int `json:"limited_buy_num"` // 限购数量
		} `json:"limited_info"` // 限购信息
		BrandId        string        `json:"brand_id"`       // 品牌 ID，无品牌为 2100000000
		Qualifications []interface{} `json:"qualifications"` // 商品资质（旧字段）
		ExtraService   struct {
			SevenDayReturn   int `json:"seven_day_return"`   // 7 天无理由
			PayAfterUse      int `json:"pay_after_use"`      // 先用后付
			FreightInsurance int `json:"freight_insurance"`  // 运费险
			DamageGuarantee  int `json:"damage_guarantee"`   // 破损包退换
			FakeOnePayThree  int `json:"fake_one_pay_three"` // 假一赔三
			ExchangeSupport  int `json:"exchange_support"`   // 换货保障
		} `json:"extra_service"` // 额外服务
		ProductType   int `json:"product_type"` // 商品类型
		EditTime      int `json:"edit_time"`    // 草稿最近修改时间（秒级时间戳）
		AfterSaleInfo struct {
			AfterSaleAddressId string `json:"after_sale_address_id"` // 售后地址 ID
		} `json:"after_sale_info"` // 售后信息
		HideInWindow    int `json:"hide_in_window"` // 是否在店铺首页隐藏
		ProductQuaInfos []struct {
			QuaId  string   `json:"qua_id"`  // 资质 ID
			QuaUrl []string `json:"qua_url"` // 资质文件 URL 列表
		} `json:"product_qua_infos"` // 商品资质列表
		SizeChart struct {
			Enable            bool          `json:"enable"`             // 是否启用尺码表
			SpecificationList []interface{} `json:"specification_list"` // 尺码规格信息
		} `json:"size_chart"` // 尺码信息
		DeliverAcctType []interface{} `json:"deliver_acct_type"` // 无需快递时的发货账号类型
		DomainType      int           `json:"domain_type"`       // 领域类型
		CustomConfig    struct {
			CustomType           []interface{} `json:"custom_type"`             // 定制类型
			CustomTextMaxLength  int           `json:"custom_text_max_length"`  // 定制文案最大长度
			CustomTextInputType  int           `json:"custom_text_input_type"`  // 定制输入类型
			CustomTextDirection  int           `json:"custom_text_direction"`   // 文案方向
			CustomTextFontSize   int           `json:"custom_text_font_size"`   // 文案字号
			CustomTextColor      string        `json:"custom_text_color"`       // 文案颜色
			CustomTextInputTypes []interface{} `json:"custom_text_input_types"` // 定制输入类型集合
			OpenCustom           bool          `json:"open_custom"`             // 是否开启定制
			DescImgList          []interface{} `json:"desc_img_list"`           // 定制说明图
			DeliveryTime         int           `json:"delivery_time"`           // 交付时间
			PreviewType          int           `json:"preview_type"`            // 预览类型
		} `json:"custom_config"` // 定制化配置
		ShortTitle     string        `json:"short_title"`  // 短标题
		ReleaseMode    int           `json:"release_mode"` // 发布模式
		HeadVideos     []interface{} `json:"head_videos"`  // 主图视频列表
		SpuDeliverInfo struct {
			SkuDeliverInfo struct {
				StockType                      int `json:"stock_type"`                         // 库存类型
				FullPaymentPresaleDeliveryType int `json:"full_payment_presale_delivery_type"` // 全款预售发货类型
				PresaleBeginTime               int `json:"presale_begin_time"`                 // 预售开始时间
				PresaleEndTime                 int `json:"presale_end_time"`                   // 预售结束时间
				FullPaymentPresaleDeliveryTime int `json:"full_payment_presale_delivery_time"` // 全款预售发货时间
				SpotAfterPresaleEnd            int `json:"spot_after_presale_end"`             // 预售结束后是否转现货
			} `json:"sku_deliver_info"` // SPU 维度发货配置
			IsSpuRange int `json:"is_spu_range"` // 是否按 SPU 维度生效
		} `json:"spu_deliver_info"` // SPU 维度预售配置
		CatsV2 []struct {
			CatId string `json:"cat_id"` // 新类目 ID
		} `json:"cats_v2"` // 新类目树
		TimingOnsaleInfo struct {
			Status      int `json:"status"`        // 待开售状态
			OnsaleTime  int `json:"onsale_time"`   // 开售时间（秒级时间戳）
			IsHidePrice int `json:"is_hide_price"` // 开售前是否隐藏价格
			TaskId      int `json:"task_id"`       // 定时任务 ID
		} `json:"timing_onsale_info"` // 待开售信息
	} `json:"edit_product"` // 商品草稿数据
	InfoScore struct {
		SubScoreList []struct {
			AuditRemark      string `json:"audit_remark"`        // 审核建议
			BusiDataFieldApi string `json:"busi_data_field_api"` // 业务字段 API 标识
			FieldName        string `json:"field_name"`          // 字段名称
		} `json:"sub_score_list"` // 子维度评分
		ScoreLevel int `json:"score_level"` // 信息质量评分等级
	} `json:"info_score"` // 商品信息质量
	CmpPriceInfo struct {
		SkuResultList []interface{} `json:"sku_result_list"` // SKU 高价结果
		Status        int           `json:"status"`          // 高价预警状态
	} `json:"cmp_price_info"` // 商品高价预警
	AuditInfo struct {
		UserStrategyFlagList []interface{} `json:"user_strategy_flag_list"` // 用户策略标记列表
	} `json:"audit_info"` // 审核信息
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
