# wxshop-go-sdk

微信小店 Go 语言 SDK。

## 功能

提供微信小店 API 的 Go 语言封装，方便开发者快速接入。

## 快速开始

### 安装

```bash
go get github.com/ealink1/wxshop_go_sdk
go get -u github.com/ealink1/wxshop_go_sdk@latest
```

> 注意：请根据实际仓库地址调整上述 import 路径。

### 使用示例

```go
package main

import (
	"context"
	"fmt"
	"log"

	wxshop "github.com/ealink1/wxshop_go_sdk"
)

func main() {
	appID := "your_app_id"
	appSecret := "your_app_secret"
	client := wxshop.NewClient(appID, appSecret)
	ctx := context.Background()

	tokenResult, err := client.GetStableAccessTokenDirect(ctx, true)
	if err != nil {
		log.Fatalf("获取 stable access_token 失败：%v", err)
	}

	client.SetAccessToken(tokenResult.AccessToken)
	quotaResult, err := client.GetApiQuota(ctx, &wxshop.GetApiQuotaRequest{
		CGIPath: wxshop.AccessStableTokenApi,
	})
	if err != nil {
		log.Fatalf("查询 API 调用额度失败：%v", err)
	}

	fmt.Printf("access_token=%s\n", tokenResult.AccessToken)
	fmt.Printf("daily_limit=%d used=%d remain=%d\n", quotaResult.Quota.DailyLimit, quotaResult.Quota.Used, quotaResult.Quota.Remain)
}
```

## 目录结构

- `wxshop.go`: 客户端结构定义及环境配置
- `cgi_bin_common.go`: 通用接口实现（token、quota、rid、callback 等）
- `channels_ec_basics_shop.go`: 店铺管理接口实现
- `channels_ec_product_shop.go`: 商品管理查询类接口实现
- `channels_ec_order_shop.go`: 订单管理查询类接口实现
- `channels_ec_funds_shop.go`: 资金管理查询类接口实现

## 接口列表

| 方法名 | 接口路径 | 说明 |
| --- | --- | --- |
| `GetAccessTokenDirect` | `/cgi-bin/token` | 直接获取 access_token |
| `GetStableAccessTokenDirect` | `/cgi-bin/stable_token` | 获取稳定版接口调用凭据 |
| `GetApiQuota` | `/cgi-bin/openapi/quota/get` | 查询指定 API 的调用额度与频率限制 |
| `ClearApiQuota` | `/cgi-bin/openapi/quota/clear` | 重置指定 API 的每日调用次数 |
| `ClearQuota` | `/cgi-bin/clear_quota` | 使用 access_token 重置每日 API 调用次数 |
| `ClearQuotaByAppSecret` | `/cgi-bin/clear_quota/v2` | 使用 AppSecret 重置每日 API 调用次数 |
| `CallbackCheck` | `/cgi-bin/callback/check` | 对回调地址执行域名解析和 ping 检测 |
| `GetRidInfo` | `/cgi-bin/openapi/rid/get` | 查询接口报错返回 rid 的详细信息 |
| `GetShopBasicInfo` | `/channels/ec/basics/info/get` | 获取店铺基本信息 |
| `GetShopH5URL` | `/channels/ec/basics/shop/h5url/get` | 获取店铺 H5 链接 |
| `GetShopQRCode` | `/channels/ec/basics/shop/qrcode/get` | 获取店铺二维码 |
| `GetShopTagLink` | `/channels/ec/basics/shop/taglink/get` | 获取店铺微信口令 |
| `GetProduct` | `/channels/ec/product/get` | 获取商品详情（线上/草稿） |
| `GetProductList` | `/channels/ec/product/list/get` | 获取商品列表 |
| `GetProductH5URL` | `/channels/ec/product/h5url/get` | 获取商品 H5 短链 |
| `GetProductTagLink` | `/channels/ec/product/taglink/get` | 获取商品微信口令 |
| `GetProductQRCode` | `/channels/ec/product/qrcode/get` | 获取商品二维码 |
| `GetProductScheme` | `/channels/ec/product/scheme/get` | 获取商品移动应用跳转 scheme 码 |
| `GetProductAuditStrategy` | `/channels/ec/product/auditstrategy/get` | 获取商品上架策略 |
| `GetOrderList` | `/channels/ec/order/list/get` | 获取订单列表 |
| `GetOrder` | `/channels/ec/order/get` | 获取订单详情 |
| `SearchOrder` | `/channels/ec/order/search` | 搜索订单 |
| `GetFundsBalance` | `/channels/ec/funds/getbalance` | 获取账户余额 |
| `GetFundsBankAcct` | `/channels/ec/funds/getbankacct` | 获取结算账户 |
| `GetFundsFlowDetail` | `/channels/ec/funds/getfundsflowdetail` | 获取资金流水详情 |
| `GetFundsFlowList` | `/channels/ec/funds/getfundsflowlist` | 获取资金流水列表 |
| `GetFundsWithdrawList` | `/channels/ec/funds/getwithdrawlist` | 获取提现记录列表 |
| `ListOrderFlow` | `/channels/ec/funds/listorderflow` | 查询订单流水列表 |

## License

MIT
