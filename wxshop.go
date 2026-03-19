package wxshop_go_sdk

// Client 是与微信小店 API 交互的主客户端
type Client struct {
	// 客户端配置字段
	AppID       string // 应用 ID
	AppSecret   string // 应用密钥
	AccessToken string // access_token 管理器
	Env         string
}

// NewClient 创建一个新的微信小店客户端
func NewClient(appID, appSecret string) *Client {
	return &Client{
		AppID:     appID,
		AppSecret: appSecret,
		Env:       OnlineEnv,
	}
}

// SetAccessToken 设置接口调用凭证
func (c *Client) SetAccessToken(accToken string) {
	c.AccessToken = accToken
}

// SetAppID 设置应用 ID
func (c *Client) SetAppID(appid string) {
	c.AppID = appid
}

// SetAppSecret 设置应用密钥
func (c *Client) SetAppSecret(secret string) {
	c.AppSecret = secret
}

// SetEnv 设置接口请求环境地址
func (c *Client) SetEnv(env string) {
	c.Env = env
}

const (
	OnlineEnv = "https://api.weixin.qq.com"

	AccessTokenApi             = "/cgi-bin/token"                         // 获取 access_token
	AccessStableTokenApi       = "/cgi-bin/stable_token"                  // 获取稳定版接口调用凭据
	OpenApiQuotaGetApi         = "/cgi-bin/openapi/quota/get"             // 查询 API 调用额度
	OpenApiQuotaClearApi       = "/cgi-bin/openapi/quota/clear"           // 重置指定 API 调用次数
	ClearQuotaApi              = "/cgi-bin/clear_quota"                   // 重置每日 API 调用次数
	ClearQuotaBySecretApi      = "/cgi-bin/clear_quota/v2"                // 使用 AppSecret 重置每日 API 调用次数
	CallbackCheckApi           = "/cgi-bin/callback/check"                // 网络通信检测
	GetRidInfoApi              = "/cgi-bin/openapi/rid/get"               // 查询 rid 信息
	BasicsInfoGetApi           = "/channels/ec/basics/info/get"           // 获取店铺基本信息
	ShopH5URLGetApi            = "/channels/ec/basics/shop/h5url/get"     // 获取店铺 H5 链接
	ShopQRCodeGetApi           = "/channels/ec/basics/shop/qrcode/get"    // 获取店铺二维码
	ShopTagLinkGetApi          = "/channels/ec/basics/shop/taglink/get"   // 获取店铺口令
	ProductGetApi              = "/channels/ec/product/get"               // 获取商品详情
	ProductListGetApi          = "/channels/ec/product/list/get"          // 获取商品列表
	ProductH5URLGetApi         = "/channels/ec/product/h5url/get"         // 获取商品 H5 短链
	ProductTagLinkGetApi       = "/channels/ec/product/taglink/get"       // 获取商品口令
	ProductQRCodeGetApi        = "/channels/ec/product/qrcode/get"        // 获取商品二维码
	ProductSchemeGetApi        = "/channels/ec/product/scheme/get"        // 获取商品移动应用跳转 scheme 码
	ProductAuditStrategyGetApi = "/channels/ec/product/auditstrategy/get" // 获取商品上架策略
	OrderListGetApi            = "/channels/ec/order/list/get"            // 获取订单列表
	OrderGetApi                = "/channels/ec/order/get"                 // 获取订单详情
	OrderSearchApi             = "/channels/ec/order/search"              // 搜索订单
)
