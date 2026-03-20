package wxshop_go_sdk

const (
	MsgTypeOrderNew          = "channels_ec_order_new"           // 新订单
	MsgTypeOrderCancel       = "channels_ec_order_cancel"        // 取消订单
	MsgTypeOrderPay          = "channels_ec_order_pay"           // 订单支付
	MsgTypeOrderWaitShipping = "channels_ec_order_wait_shipping" // 等待发货
	MsgTypeOrderDeliver      = "channels_ec_order_deliver"       // 订单发货
	MsgTypeOrderConfirm      = "channels_ec_order_confirm"       // 订单确认
	MsgTypeOrderSettle       = "channels_ec_order_settle"        // 订单结算
	MsgTypeOrderWaitOrder    = "channels_ec_order_wait_order"    // 其他
)

type (
	// MsgOrder 订单消息 - 公共字段
	MsgOrder struct {
		ToUserName   string `json:"ToUserName"`
		FromUserName string `json:"FromUserName"`
		CreateTime   int    `json:"CreateTime"`
		MsgType      string `json:"MsgType"`
		Event        string `json:"Event"`
	}

	// MsgOrderNew 订单消息 - 新订单
	MsgOrderNew struct {
		MsgOrder
		OrderInfo struct {
			OrderId int64 `json:"order_id"`
		} `json:"order_info"`
	}

	// MsgOrderCancel 订单消息 - 取消订单
	MsgOrderCancel struct {
		MsgOrder
		OrderInfo struct {
			OrderId    int64 `json:"order_id"`
			CancelType int   `json:"cancel_type"`
		} `json:"order_info"`
	}

	// MsgOrderPay 订单消息 - 订单支付
	MsgOrderPay struct {
		MsgOrder
		OrderInfo struct {
			OrderId int64 `json:"order_id"`
			PayTime int   `json:"pay_time"`
		} `json:"order_info"`
	}

	// MsgOrderWaitShipping 订单消息 - 等待发货
	MsgOrderWaitShipping struct {
		MsgOrder
		OrderInfo struct {
			OrderId int64 `json:"order_id"`
		} `json:"order_info"`
	}

	// MsgOrderDeliver 订单消息 - 订单发货
	MsgOrderDeliver struct {
		MsgOrder
		OrderInfo struct {
			OrderId        int64 `json:"order_id"`
			FinishDelivery int64 `json:"finish_delivery"` // 0:尚未全部发货；1:全部商品发货完成
		} `json:"order_info"`
	}

	// MsgOrderConfirm 订单消息 - 订单确认
	MsgOrderConfirm struct {
		MsgOrder
		OrderInfo struct {
			OrderId     int64 `json:"order_id"`
			ConfirmType int64 `json:"confirm_type"` // 1:用户确认收货；2:超时自动确认收货
		} `json:"order_info"`
	}

	// MsgOrderSettle 订单消息 - 订单结算
	MsgOrderSettle struct {
		MsgOrder
		OrderInfo struct {
			OrderId    int64 `json:"order_id"`
			SettleTime int64 `json:"settle_time"`
		} `json:"order_info"`
	}

	// MsgOrderWaitOrder 订单消息 - 其他
	MsgOrderWaitOrder struct {
		MsgOrder
		OrderInfo struct {
			OrderId int64 `json:"order_id"`
			Type    int64 `json:"type"` //
		} `json:"order_info"`
	}
)

//	1：联盟佣金信息
//	2：商家主动地址修改或通过用户修改地址申请
//	3：商家备注修改
//	4：用户发起申请修改收货地址，特殊条件下需要商家审批
//	5：订单虚拟号码信息更新
//	6：分享员信息更新
//	7：用户催发货
//	8：商家修改礼物单备注
//	9：用户发起发货前更换sku请求
//	10：商家同意用户发货前更换sku请求
//	11：用户收礼时更换sku
//	12：商家修改订单价格
//	13：真实号审核结果
//	14：分配代发
//	15：取消代发
//	16：买家已同意修改，请按新地址发货
//	17：买家已拒绝修改，请按原地址发货
//	18：买家超时未确认，修改申请已失效
//	19：买家已主动修改地址，请按新地址发货
//	20：该订单已发货，修改申请已失效
