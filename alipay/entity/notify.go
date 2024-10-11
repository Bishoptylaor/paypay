/*
 *  ┏┓      ┏┓
 *┏━┛┻━━━━━━┛┻┓
 *┃　　　━　　  ┃
 *┃   ┳┛ ┗┳   ┃
 *┃           ┃
 *┃     ┻     ┃
 *┗━━━┓     ┏━┛
 *　　 ┃　　　┃神兽保佑
 *　　 ┃　　　┃代码无BUG！
 *　　 ┃　　　┗━━━┓
 *　　 ┃         ┣┓
 *　　 ┃         ┏┛
 *　　 ┗━┓┓┏━━┳┓┏┛
 *　　   ┃┫┫  ┃┫┫
 *      ┗┻┛　 ┗┻┛
 @Time    : 2024/8/28 -- 18:28
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: notify.go
*/

package entity

type StdIn struct {
	NotifyId     string `json:"notify_id"`
	UtcTimestamp int64  `json:"utc_timestamp"`
	MsgMethod    string `json:"msg_method"`
	Appid        string `json:"appid"`
	Version      string `json:"version"`
	Sign         string `json:"sign"`
	SignType     string `json:"sign_type"`
	Charset      string `json:"charset"`
}

type TradeRefundDepositbackCompletedReq struct {
	StdIn
	BizContent TradeRefundDepositbackCompleted `json:"biz_content"`
}

type TradeRefundDepositbackCompleted struct {
	TradeNo            string `json:"trade_no"`                        // 支付宝交易号
	OutTradeNo         string `json:"out_trade_no"`                    // 商户订单号
	OutRequestNo       string `json:"out_request_no"`                  // 退款请求号
	DbackStatus        string `json:"dback_status"`                    // 银行卡冲退状态。S-成功，F-失败。银行卡冲退失败，资金自动转入用户支付宝余额。
	DbackAmount        string `json:"dback_amount"`                    // 银行卡冲退金额，仅当dback_status=S时，才会返回。单位：元。
	BankAckTime        string `json:"bank_ack_time,omitempty"`         // 银行响应时间，格式为yyyy-MM-dd HH:mm:ss
	EstBankReceiptTime string `json:"est_bank_receipt_time,omitempty"` // 预估银行入账时间，格式为yyyy-MM-dd HH:mm:ss
}

// =========================================================分割=========================================================

type MarketingActivityDeliveryChangedReq struct {
	StdIn
	BizContent MarketingActivityDeliveryChanged `json:"biz_content"`
}

type MarketingActivityDeliveryChanged struct {
	EventTime         string `json:"event_time"`
	DeliveryId        string `json:"delivery_id"`
	DeliveryStatus    string `json:"delivery_status"`
	DeliveryErrorMsg  string `json:"delivery_error_msg"`
	DeliveryBoothCode string `json:"delivery_booth_code"`
}

// =========================================================分割=========================================================

type MarketingActivityDeliverycreateNotifyReq struct {
	StdIn
	BizContent MarketingActivityDeliverycreateNotify `json:"biz_content"`
}

type MarketingActivityDeliverycreateNotify struct {
	SubjectId              string `json:"subject_id"`                // 推广内容ID
	SubjectType            string `json:"subject_type"`              // 推广内容类型
	EventTime              string `json:"event_time"`                // 事件创建时间
	OutBizNo               string `json:"out_biz_no"`                // 外部业务唯一单号
	DeliveryCreateStatus   string `json:"delivery_create_status"`    // 推广创建结果
	DeliveryId             string `json:"delivery_id"`               // 推广计划id
	DeliveryBoothCode      string `json:"delivery_booth_code"`       // 展位编码
	DeliveryCreateErrorMsg string `json:"delivery_create_error_msg"` // 推广创建失败原因
}
