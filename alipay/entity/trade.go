package entity

type TradePayResponse struct {
	Response     *TradePay `json:"alipay_trade_pay_response"`
	AlipayCertSn string    `json:"alipay_cert_sn,omitempty"`
	SignData     string    `json:"-"`
	Sign         string    `json:"sign"`
}

type TradePay struct {
	ErrorResponse
	TradeNo             string           `json:"trade_no"`       // 支付宝交易号
	OutTradeNo          string           `json:"out_trade_no"`   // 创建交易传入的商户订单号
	BuyerLogonId        string           `json:"buyer_logon_id"` // 买家支付宝账号
	TotalAmount         string           `json:"total_amount"`   // 交易金额
	ReceiptAmount       string           `json:"receipt_amount"` // 实收金额，单位为元，两位小数
	GmtPayment          string           `json:"gmt_payment"`    // 交易支付时间
	FundBillList        []*TradeFundBill `json:"fund_bill_list"` // 交易支付使用的资金渠道。
	BuyerUserId         string           `json:"buyer_user_id,omitempty"`
	BuyerOpenId         string           `json:"buyer_open_id"`
	BuyerPayAmount      string           `json:"buyer_pay_amount,omitempty"`      // 买家付款的金额
	PointAmount         string           `json:"point_amount,omitempty"`          // 使用集分宝付款的金额
	InvoiceAmount       string           `json:"invoice_amount,omitempty"`        // 交易中可给用户开具发票的金额
	StoreName           string           `json:"store_name,omitempty"`            // 发生支付交易的商户门店名称
	DiscountGoodsDetail string           `json:"discount_goods_detail,omitempty"` // 本次交易支付所使用的单品券优惠的商品优惠信息。只有在query_options中指定时才返回该字段信息
	AsyncPaymentMode    string           `json:"async_payment_mode,omitempty"`    // 异步支付模式，目前有五种值：ASYNC_DELAY_PAY(异步延时付款);ASYNC_REALTIME_PAY(异步准实时付款);SYNC_DIRECT_PAY(同步直接扣款);NORMAL_ASYNC_PAY(纯异步付款);QUOTA_OCCUPYIED_ASYNC_PAY(异步支付并且预占了先享后付额度);
	VoucherDetailList   []*VoucherDetail `json:"voucher_detail_list,omitempty"`   // 本交易支付时使用的所有优惠券信息。只有在query_options中指定时才返回该字段信息。
	AdvanceAmount       string           `json:"advance_amount,omitempty"`        // 先享后付2.0垫资金额,不返回表示没有走垫资，非空表示垫资支付的金额
	ChargeFlags         string           `json:"charge_flags,omitempty"`          // 费率活动标识。当交易享受特殊行业或活动费率时，返回该场景的标识。具体场景如下：trade_special_00：订单优惠费率；industry_special_on_00：线上行业特殊费率0；industry_special_on_01：线上行业特殊费率1；industry_special_00：线下行业特殊费率0；industry_special_01：线下行业特殊费率1；bluesea_1：蓝海活动优惠费率标签；注：只在机构间联模式下返回，其它场景下不返回该字段；
	AuthTradePayMode    string           `json:"auth_trade_pay_mode,omitempty"`
	MdiscountAmount     string           `json:"mdiscount_amount,omitempty"` // 商家优惠金额
	DiscountAmount      string           `json:"discount_amount,omitempty"`  // 平台优惠金额
	CreditPayMode       string           `json:"credit_pay_mode"`
	CreditBizOrderId    string           `json:"credit_biz_order_id"`
}

// =========================================================分割=========================================================

type TradePrecreateResponse struct {
	Response     *TradePrecreate `json:"alipay_trade_precreate_response"`
	NullResponse *ErrorResponse  `json:"null_response,omitempty"`
	AlipayCertSn string          `json:"alipay_cert_sn,omitempty"`
	SignData     string          `json:"-"`
	Sign         string          `json:"sign"`
}

type TradePrecreate struct {
	ErrorResponse
	OutTradeNo string `json:"out_trade_no"` // 商户的订单号
	QrCode     string `json:"qr_code"`      // 当前预下单请求生成的二维码码串，有效时间2小时，可以用二维码生成工具根据该码串值生成对应的二维码
}

// =========================================================分割=========================================================

type TradeCreateResponse struct {
	Response     *TradeCreate `json:"alipay_trade_create_response"`
	AlipayCertSn string       `json:"alipay_cert_sn,omitempty"`
	SignData     string       `json:"-"`
	Sign         string       `json:"sign"`
}

type TradeCreate struct {
	ErrorResponse
	TradeNo    string `json:"trade_no"`
	OutTradeNo string `json:"out_trade_no"`
}

// =========================================================分割=========================================================

type TradeOrderPayResponse struct {
	Response     *TradeOrderPay `json:"alipay_trade_order_pay_response"`
	AlipayCertSn string         `json:"alipay_cert_sn,omitempty"`
	SignData     string         `json:"-"`
	Sign         string         `json:"sign"`
}

type TradeOrderPay struct {
	ErrorResponse
	TradeNo           string `json:"trade_no"`
	OutTradeNo        string `json:"out_trade_no"`
	FulfillmentAmount string `json:"fulfillment_amount,omitempty"` // 实际履约金额，单位（元）。仅履约场景才会返回
	OutRequestNo      string `json:"out_request_no,omitempty"`
	TotalAmount       string `json:"total_amount,omitempty"`
	GmtPayment        string `json:"gmt_payment,omitempty"`
	AsyncPaymentMode  string `json:"async_payment_mode,omitempty"`
}

// =========================================================分割=========================================================

type TradeWapMergePayResponse struct {
	Response     *TradeWapMergePay `json:"alipay_trade_wap_merge_pay_response"`
	AlipayCertSn string            `json:"alipay_cert_sn,omitempty"`
	SignData     string            `json:"-"`
	Sign         string            `json:"sign"`
}

type TradeWapMergePay struct {
	ErrorResponse
	OutMergeNo         string               `json:"out_merge_no"`         // 如果和支付宝约定子订单必须同时支付成功或者同时失败并且请求时传递了`out_merge_no`时才存在
	MergePayStatus     string               `json:"merge_pay_status"`     // 1. FINISHED：全部订单付款成功 2. CLOSED：全部订单付款失败
	OrderDetailResults []*OrderDetailResult `json:"order_detail_results"` // 合并子订单中所有订单的支付结果信息
}

// =========================================================分割=========================================================

type TradeQueryResponse struct {
	Response     *TradeQuery `json:"alipay_trade_query_response"`
	AlipayCertSn string      `json:"alipay_cert_sn,omitempty"`
	SignData     string      `json:"-"`
	Sign         string      `json:"sign"`
}

type TradeQuery struct {
	ErrorResponse
	TradeNo               string               `json:"trade_no"`                    // 支付宝交易号
	OutTradeNo            string               `json:"out_trade_no"`                // 商家订单号
	BuyerLogonId          string               `json:"buyer_logon_id"`              // 买家支付宝账号
	TradeStatus           string               `json:"trade_status"`                // 交易状态：WAIT_BUYER_PAY（交易创建，等待买家付款）、TRADE_CLOSED（未付款交易超时关闭，或支付完成后全额退款）、TRADE_SUCCESS（交易支付成功）、TRADE_FINISHED（交易结束，不可退款）
	TotalAmount           string               `json:"total_amount"`                // 交易的订单金额，单位为元，两位小数。该参数的值为支付时传入的total_amount
	ReqGoodsDetail        []*GoodsDetail       `json:"req_goods_detail"`            // 支付请求的商品明细列表
	PeriodScene           string               `json:"period_scene,omitempty"`      // 该字段用于描述当前账期交易的场景。
	TransCurrency         string               `json:"trans_currency,omitempty"`    // 支持英镑：GBP、港币：HKD、美元：USD、新加坡元：SGD、日元：JPY、加拿大元：CAD、澳元：AUD、欧元：EUR、新西兰元：NZD、韩元：KRW、泰铢：THB、瑞士法郎：CHF、瑞典克朗：SEK、丹麦克朗：DKK、挪威克朗：NOK、马来西亚林吉特：MYR、印尼卢比：IDR、菲律宾比索：PHP、毛里求斯卢比：MUR、以色列新谢克尔：ILS、斯里兰卡卢比：LKR、俄罗斯卢布：RUB、阿联酋迪拉姆：AED、捷克克朗：CZK、南非兰特：ZAR、人民币：CNY、新台币：TWD。当trans_currency 和 settle_currency 不一致时，trans_currency支持人民币：CNY、新台币：TWD
	SettleCurrency        string               `json:"settle_currency,omitempty"`   // 订单结算币种，对应支付接口传入的settle_currency
	SettleAmount          string               `json:"settle_amount,omitempty"`     // 结算币种订单金额
	PayCurrency           string               `json:"pay_currency,omitempty"`      // 订单支付币种
	PayAmount             string               `json:"pay_amount,omitempty"`        // 支付币种订单金额
	SettleTransRate       string               `json:"settle_trans_rate,omitempty"` // 结算币种兑换标价币种汇率
	TransPayRate          string               `json:"trans_pay_rate,omitempty"`    // 标价币种兑换支付币种汇率
	BuyerPayAmount        string               `json:"buyer_pay_amount,omitempty"`  // 买家实付金额，单位为元，两位小数。该金额代表该笔交易买家实际支付的金额，不包含商户折扣等金额
	PointAmount           string               `json:"point_amount,omitempty"`      // 积分支付的金额，单位为元，两位小数。该金额代表该笔交易中用户使用积分支付的金额，比如集分宝或者支付宝实时优惠等
	InvoiceAmount         string               `json:"invoice_amount,omitempty"`    // 交易中用户支付的可开具发票的金额，单位为元，两位小数。该金额代表该笔交易中可以给用户开具发票的金额
	SendPayDate           string               `json:"send_pay_date,omitempty"`     // 本次交易打款给卖家的时间
	ReceiptAmount         string               `json:"receipt_amount,omitempty"`    // 实收金额，单位为元，两位小数。该金额为本笔交易，商户账户能够实际收到的金额
	StoreId               string               `json:"store_id,omitempty"`          // 商户门店编号
	TerminalId            string               `json:"terminal_id,omitempty"`       // 商户机具终端编号
	StoreName             string               `json:"store_name,omitempty"`        // 请求交易支付中的商户店铺的名称
	BuyerUserId           string               `json:"buyer_user_id,omitempty"`     // 买家在支付宝的用户id
	BuyerOpenId           string               `json:"buyer_open_id,omitempty"`     // 买家支付宝用户唯一标识
	DiscountGoodsDetail   string               `json:"discount_goods_detail,omitempty"`
	IndustrySepcDetail    string               `json:"industry_sepc_detail,omitempty"`
	IndustrySepcDetailGov string               `json:"industry_sepc_detail_gov,omitempty"` // 行业特殊信息-统筹相关
	IndustrySepcDetailAcc string               `json:"industry_sepc_detail_acc,omitempty"` // 行业特殊信息-个账相关
	ChargeAmount          string               `json:"charge_amount,omitempty"`            // 该笔交易针对收款方的收费金额；单位：元。 只在银行间联交易场景下返回该信息
	ChargeFlags           string               `json:"charge_flags,omitempty"`             // 费率活动标识。
	SettlementId          string               `json:"settlement_id,omitempty"`            // 支付清算编号，用于清算对账使用； 只在银行间联交易场景下返回该信息
	TradeSettleInfo       *TradeSettleInfo     `json:"trade_settle_info,omitempty"`        // 返回的交易结算信息，包含分账、补差等信息。 只有在query_options中指定时才返回该字段信息
	AuthTradePayMode      string               `json:"auth_trade_pay_mode,omitempty"`      // 预授权支付模式，该参数仅在信用预授权支付场景下返回。信用预授权支付：CREDIT_PREAUTH_PAY
	BuyerUserType         string               `json:"buyer_user_type,omitempty"`          // 买家用户类型。CORPORATE:企业用户；PRIVATE:个人用户。
	MdiscountAmount       string               `json:"mdiscount_amount,omitempty"`         // 商家优惠金额。单位：元
	DiscountAmount        string               `json:"discount_amount,omitempty"`          // 平台优惠金额。单位：元
	Subject               string               `json:"subject,omitempty"`                  // 订单标题； 只在银行间联交易场景下返回该信息
	Body                  string               `json:"body,omitempty"`                     // 订单描述； 只在银行间联交易场景下返回该信息
	AlipaySubMerchantId   string               `json:"alipay_sub_merchant_id,omitempty"`   // 间连商户在支付宝端的商户编号； 只在银行间联交易场景下返回该信息
	ExtInfos              string               `json:"ext_infos,omitempty"`                // 交易额外信息，特殊场景下与支付宝约定返回
	HbFqPayInfo           *HbFqPayInfo         `json:"hb_fq_pay_info,omitempty"`           // 若用户使用花呗分期支付，且商家开通返回此通知参数，则会返回花呗分期信息
	FulfillmentDetailList []*FulfillmentDetail `json:"fulfillment_detail_list,omitempty"`  // 履约详情列表。 只有入参的query_options中指定fulfillment_detail_list并且所查询的交易存在履约明细时才返回该字段信息
	AdditionalStatus      string               `json:"additional_status,omitempty"`        // 交易附加状态： SELLER_NOT_RECEIVED（买家已付款，卖家未收款）
	FundBillList          []*TradeFundBill     `json:"fund_bill_list"`                     // 交易支付使用的资金渠道
	PassbackParams        string               `json:"passback_params,omitempty"`          // 公用回传参数
	CreditPayMode         string               `json:"credit_pay_mode,omitempty"`          // 信用支付模式
	CreditBizOrderId      string               `json:"credit_biz_order_id,omitempty"`      // 信用业务单号。信用支付场景才有值，先用后付产品里是芝麻订单号。
	HybAmount             string               `json:"hyb_amount,omitempty"`               // 惠营宝回票金额。单位：元
	BkagentRespInfo       *BkAgentRespInfo     `json:"bkagent_resp_info,omitempty"`        // 间联交易下，返回给机构的信息
	ChargeInfoList        []*ChargeInfo        `json:"charge_info_list,omitempty"`         // 计费信息列表
	BizSettleMode         string               `json:"biz_settle_mode,omitempty"`          // 账期结算标识，指已完成支付的订单会进行账期管控，不会实时结算。该参数目前会在使用小程序交易组件场景下返回
	AsyncPayApplyStatus   string               `json:"async_pay_apply_status,omitempty"`   // 异步支付受理状态，仅异步支付模式且query_options指定async_pay_info时返回。S：受理成功，支付宝内部会在一定期限内捞起任务推进支付，直到支付成功或超出可重试期限；其它：受理结果未知，可重试查询。
}

// =========================================================分割=========================================================

type TradeCancelResponse struct {
	Response     *TradeCancel `json:"alipay_trade_cancel_response"`
	AlipayCertSn string       `json:"alipay_cert_sn,omitempty"`
	SignData     string       `json:"-"`
	Sign         string       `json:"sign"`
}

type TradeCancel struct {
	ErrorResponse
	TradeNo            string `json:"trade_no,omitempty"` // 支付宝交易号; 当发生交易关闭或交易退款时返回
	OutTradeNo         string `json:"out_trade_no"`       // 商户订单号
	RetryFlag          string `json:"retry_flag"`         // 是否需要重试
	Action             string `json:"action,omitempty"`   // 本次撤销触发的交易动作,接口调用成功且交易存在时返回。可能的返回值： close：交易未支付，触发关闭交易动作，无退款； refund：交易已支付，触发交易退款动作； 未返回：未查询到交易，或接口调用失败
	GmtRefundPay       string `json:"gmt_refund_pay,omitempty"`
	RefundSettlementId string `json:"refund_settlement_id,omitempty"`
}

// =========================================================分割=========================================================

type TradeCloseResponse struct {
	Response     *TradeClose `json:"alipay_trade_close_response"`
	AlipayCertSn string      `json:"alipay_cert_sn,omitempty"`
	SignData     string      `json:"-"`
	Sign         string      `json:"sign"`
}

type TradeClose struct {
	ErrorResponse
	TradeNo    string `json:"trade_no,omitempty"`     // 支付宝交易号
	OutTradeNo string `json:"out_trade_no,omitempty"` // 创建交易传入的商户订单号
}

// =========================================================分割=========================================================

type TradeRefundResponse struct {
	Response     *TradeRefund `json:"alipay_trade_refund_response"`
	AlipayCertSn string       `json:"alipay_cert_sn,omitempty"`
	SignData     string       `json:"-"`
	Sign         string       `json:"sign"`
}

type TradeRefund struct {
	ErrorResponse
	TradeNo                 string              `json:"trade_no"`                             // 支付宝交易号
	OutTradeNo              string              `json:"out_trade_no"`                         // 商户订单号
	BuyerLogonId            string              `json:"buyer_logon_id"`                       // 用户的登录id
	RefundFee               string              `json:"refund_fee"`                           // 退款总金额。单位：元。 指该笔交易累计已经退款成功的金额。
	RefundDetailItemList    []*TradeFundBill    `json:"refund_detail_item_list,omitempty"`    // 退款使用的资金渠道
	StoreName               string              `json:"store_name,omitempty"`                 // 交易在支付时候的门店名称
	BuyerUserId             string              `json:"buyer_user_id,omitempty"`              // 买家在支付宝的用户id
	BuyerOpenId             string              `json:"buyer_open_id"`                        // 买家支付宝用户唯一标识
	SendBackFee             string              `json:"send_back_fee,omitempty"`              // 本次商户实际退回金额。单位：元。 说明：如需获取该值，需在入参query_options中传入 refund_detail_item_list
	FundChange              string              `json:"fund_change,omitempty"`                // 本次退款是否发生了资金变化
	RefundHybAmount         string              `json:"refund_hyb_amount,omitempty"`          // 本次请求退惠营宝金额。单位：元
	RefundChargeInfoList    []*RefundChargeInfo `json:"refund_charge_info_list,omitempty"`    // 退费信息
	RefundVoucherDetailList []*VoucherDetail    `json:"refund_voucher_detail_list,omitempty"` // 本交易支付时使用的所有优惠券信息。 只有在query_options中指定了refund_voucher_detail_list时才返回该字段信息

	OpenId                       string                 `json:"open_id,omitempty"`
	RefundCurrency               string                 `json:"refund_currency,omitempty"`
	GmtRefundPay                 string                 `json:"gmt_refund_pay,omitempty"`
	RefundPresetPaytoolList      []*RefundPresetPaytool `json:"refund_preset_paytool_list,omitempty"`
	RefundChargeAmount           string                 `json:"refund_charge_amount,omitempty"`
	RefundSettlementId           string                 `json:"refund_settlement_id,omitempty"`
	PresentRefundBuyerAmount     string                 `json:"present_refund_buyer_amount,omitempty"`
	PresentRefundDiscountAmount  string                 `json:"present_refund_discount_amount,omitempty"`
	PresentRefundMdiscountAmount string                 `json:"present_refund_mdiscount_amount,omitempty"`
	HasDepositBack               string                 `json:"has_deposit_back,omitempty"`
}

// =========================================================分割=========================================================

type TradePageRefundResponse struct {
	Response     *TradePageRefund `json:"alipay_trade_page_refund_response"`
	AlipayCertSn string           `json:"alipay_cert_sn,omitempty"`
	SignData     string           `json:"-"`
	Sign         string           `json:"sign"`
}

type TradePageRefund struct {
	ErrorResponse
	TradeNo      string `json:"trade_no,omitempty"`
	OutTradeNo   string `json:"out_trade_no,omitempty"`
	OutRequestNo string `json:"out_request_no,omitempty"`
	RefundAmount string `json:"refund_amount,omitempty"`
}

// =========================================================分割=========================================================

type TradeFastpayRefundQueryResponse struct {
	Response     *TradeRefundQuery `json:"alipay_trade_fastpay_refund_query_response"`
	AlipayCertSn string            `json:"alipay_cert_sn,omitempty"`
	SignData     string            `json:"-"`
	Sign         string            `json:"sign"`
}

type TradeRefundQuery struct {
	ErrorResponse
	TradeNo                 string              `json:"trade_no,omitempty"`                   // 支付宝交易号
	OutTradeNo              string              `json:"out_trade_no,omitempty"`               // 创建交易传入的商户订单号
	OutRequestNo            string              `json:"out_request_no,omitempty"`             // 本笔退款对应的退款请求号
	RefundReason            string              `json:"refund_reason,omitempty"`              //
	TotalAmount             string              `json:"total_amount,omitempty"`               // 该笔退款所对应的交易的订单金额。单位：元。
	RefundAmount            string              `json:"refund_amount,omitempty"`              // 本次退款请求，对应的退款金额。单位：元。
	RefundStatus            string              `json:"refund_status,omitempty"`              // 退款状态。枚举值： REFUND_SUCCESS 退款处理成功； 未返回该字段表示退款请求未收到或者退款失败； 注：如果退款查询发起时间早于退款时间，或者间隔退款发起时间太短，可能出现退款查询时还没处理成功，后面又处理成功的情况，建议商户在退款发起后间隔10秒以上再发起退款查询请求。
	RefundRoyaltys          []*RefundRoyalty    `json:"refund_royaltys,omitempty"`            // 退分账明细信息，当前仅在直付通产品中返回。
	GmtRefundPay            string              `json:"gmt_refund_pay,omitempty"`             // 退款时间。默认不返回该信息，需要在入参的query_options中指定"gmt_refund_pay"值时才返回该字段信息。
	RefundDetailItemList    []*TradeFundBill    `json:"refund_detail_item_list,omitempty"`    // 本次退款使用的资金渠道； 默认不返回该信息，需要在入参的query_options中指定"refund_detail_item_list"值时才返回该字段信息。
	SendBackFee             string              `json:"send_back_fee,omitempty"`              // 本次商户实际退回金额；单位：元。
	DepositBackInfo         *DepositBackInfo    `json:"deposit_back_info,omitempty"`          // 银行卡冲退信息； 默认不返回该信息，需要在入参的query_options中指定"deposit_back_info"值时才返回该字段信息。
	RefundVoucherDetailList []VoucherDetail     `json:"refund_voucher_detail_list,omitempty"` // 本交易支付时使用的所有优惠券信息。
	RefundHybAmount         string              `json:"refund_hyb_amount,omitempty"`          // 本次退款金额中退惠营宝的金额。单位：元。
	RefundChargeInfoList    []*RefundChargeInfo `json:"refund_charge_info_list,omitempty"`    // 退费信息
	DepositBackInfoList     []*DepositBackInfo  `json:"deposit_back_info_list,omitempty"`     // 银行卡冲退信息列表。 默认不返回该信息，需要在入参的query_options中指定"deposit_back_info_list"值时才返回该字段信息。
}

// =========================================================分割=========================================================

type TradeOrderInfoSyncResponse struct {
	Response     *TradeOrderInfoSync `json:"alipay_trade_orderinfo_sync_response"`
	AlipayCertSn string              `json:"alipay_cert_sn,omitempty"`
	SignData     string              `json:"-"`
	Sign         string              `json:"sign"`
}

type TradeOrderInfoSync struct {
	ErrorResponse
	TradeNo     string `json:"trade_no"`
	OutTradeNo  string `json:"out_trade_no"`
	BuyerUserId string `json:"buyer_user_id"`
	BuyerOpenId string `json:"buyer_open_id,omitempty"`
}

// =========================================================分割=========================================================

type TradeAdvanceConsultResponse struct {
	Response     *TradeAdvanceConsult `json:"alipay_trade_advance_consult_response"`
	AlipayCertSn string               `json:"alipay_cert_sn,omitempty"`
	SignData     string               `json:"-"`
	Sign         string               `json:"sign"`
}

type TradeAdvanceConsult struct {
	ErrorResponse
	ReferResult             bool                      `json:"refer_result"`                         // true 代表当前时间点，用户允许垫资 false 代表当前时间，用户不允许垫资
	WaitRepaymentOrderInfos []*WaitRepaymentOrderInfo `json:"wait_repayment_order_infos,omitempty"` // 待还订单列表，无论用户当前状态是否允许垫资，都会返回当前用户在商户下的待还订单信息
	WaitRepaymentAmount     string                    `json:"wait_repayment_amount,omitempty"`      // 用户剩余的总待还金额，无论当前用户是否允许垫资，都会返回该属性。
	WaitRepaymentOrderCount string                    `json:"wait_repayment_order_count,omitempty"` // 用户总的未还的垫资笔数，无论用户是否允许垫资，都会返回该属性
	RiskLevel               string                    `json:"risk_level,omitempty"`                 // 订单风险评估等级，在单笔订单风险预评估时返回。当基础风险校验通过时，可通过该值获取业务风险评估等级。取值：2-高风险；1-低风险。
	ResultMessage           string                    `json:"result_message"`                       // 返回用户不准入原因
	ResultCode              string                    `json:"result_code"`                          // 用户被注销
	UserRiskPrediction      *UserRiskPrediction       `json:"user_risk_prediction,omitempty"`       // 用户风险预测结果，包括用户拒付风险等级、用户绑定手机号被二次放号风险等级。
}

// =========================================================分割=========================================================

type PcreditHuabeiAuthSettleApplyResponse struct {
	Response     *PcreditHuabeiAuthSettleApply `json:"alipay_pcredit_huabei_auth_settle_apply_response"`
	AlipayCertSn string                        `json:"alipay_cert_sn,omitempty"`
	SignData     string                        `json:"-"`
	Sign         string                        `json:"sign"`
}

type PcreditHuabeiAuthSettleApply struct {
	ErrorResponse
	OutRequestNo string `json:"out_request_no"`
	FailReason   string `json:"fail_reason,omitempty"`
}

// =========================================================分割=========================================================

type PaymentTradeOrderCreateResponse struct {
	Response     *PaymentTradeOrderCreate `json:"mybank_payment_trade_order_create_response"`
	AlipayCertSn string                   `json:"alipay_cert_sn,omitempty"`
	SignData     string                   `json:"-"`
	Sign         string                   `json:"sign"`
}

type PaymentTradeOrderCreate struct {
	ErrorResponse
}

// =========================================================分割=========================================================

type TradeRepaybillQueryResponse struct {
	Response     *TradeRepaybillQuery `json:"alipay_trade_repaybill_query_response"`
	AlipayCertSn string               `json:"alipay_cert_sn,omitempty"`
	SignData     string               `json:"-"`
	Sign         string               `json:"sign"`
}

type TradeRepaybillQuery struct {
	ErrorResponse
	BillNo                string `json:"bill_no"`
	BillAmount            string `json:"bill_amount"`
	BillOverdueAmount     string `json:"bill_overdue_amount"`
	BillPaidAmount        string `json:"bill_paid_amount"`
	BillPaidRevokedAmount string `json:"bill_paid_revoked_amount"`
	BillRevokedAmount     string `json:"bill_revoked_amount"`
	BillStatus            string `json:"bill_status"`
}

// =========================================================分割=========================================================

type TradeFundBill struct {
	FundChannel string `json:"fund_channel"` // 交易使用的资金渠道，详见 https://opendocs.alipay.com/open/common/103259
	Amount      string `json:"amount"`       // 该支付工具类型所使用的金额
	RealAmount  string `json:"real_amount"`  // 渠道实际付款金额
	FundType    string `json:"fund_type,omitempty"`
}

type VoucherDetail struct {
	Id                         string `json:"id"`                                     // 券id
	Name                       string `json:"name"`                                   // 券名称
	Type                       string `json:"type"`                                   // 全场代金券: ALIPAY_FIX_VOUCHER 折扣券: ALIPAY_DISCOUNT_VOUCHER单品优惠券: ALIPAY_ITEM_VOUCHER 现金抵价券: ALIPAY_CASH_VOUCHER 商家全场券: ALIPAY_BIZ_VOUCHER
	Amount                     string `json:"amount"`                                 // 优惠券面额，它应该会等于商家出资加上其他出资方出资
	MerchantContribute         string `json:"merchant_contribute,omitempty"`          // 商家出资（特指发起交易的商家出资金额）
	OtherContribute            string `json:"other_contribute,omitempty"`             // 其他出资方出资金额，可能是支付宝，可能是品牌商，或者其他方，也可能是他们的一起出资
	Memo                       string `json:"memo,omitempty"`                         // 优惠券备注信息
	TemplateId                 string `json:"template_id,omitempty"`                  // 券模板id
	PurchaseBuyerContribute    string `json:"purchase_buyer_contribute,omitempty"`    // 如果使用的这张券是用户购买的，则该字段代表用户在购买这张券时用户实际付款的金额
	PurchaseMerchantContribute string `json:"purchase_merchant_contribute,omitempty"` // 如果使用的这张券是用户购买的，则该字段代表用户在购买这张券时商户优惠的金额
	PurchaseAntContribute      string `json:"purchase_ant_contribute,omitempty"`      // 如果使用的这张券是用户购买的，则该字段代表用户在购买这张券时平台优惠的金额
}

type OrderDetailResult struct {
	Appid          string `json:"appid"`        // 应用唯一标识
	OutTradeNo     string `json:"out_trade_no"` // 商户订单号
	TradeNo        string `json:"trade_no"`     // 支付宝交易号
	TradeStatus    string `json:"trade_status"` // 1. TRADE_SUCCESS：付款成功 2. TRADE_FINISHED：交易完成 3. WAIT_BUYER_PAY：等待支付 4. TRADE_CLOSED：交易关闭
	Subject        string `json:"subject"`
	TotalAmount    string `json:"total_amount"`
	SellerId       string `json:"seller_id"`       // 卖家支付宝用户ID。
	PassbackParams string `json:"passback_params"` // 公用回传参数，如果请求时传递了该参数，则返回给商户时会回传该参数。
}

type GoodsDetail struct {
	GoodsId        string `json:"goods_id"`                  // 商品的编号，该参数传入支付券上绑定商品goods_id, 倘若无支付券需要消费，该字段传入商品最小粒度的商品ID（如：若商品有sku粒度，则传商户sku粒度的ID）
	GoodsName      string `json:"goods_name"`                // 商品名称
	Quantity       string `json:"quantity"`                  // 商品数量
	Price          string `json:"price"`                     // 商品单价，单位为元
	AlipayGoodsId  string `json:"alipay_goods_id,omitempty"` // 支付宝定义的统一商品编号
	GoodsCategory  string `json:"goods_category,omitempty"`  // 商品类目
	CategoriesTree string `json:"categories_tree,omitempty"` // 商品类目树，从商品类目根节点到叶子节点的类目id组成，类目id值使用|分割
	Body           string `json:"body,omitempty"`            // 商品描述信息
	ShowUrl        string `json:"show_url,omitempty"`        // 商品的展示地址
	OutItemId      string `json:"out_item_id,omitempty"`     // 商家侧小程序商品ID，指商家提报给小程序商品库的商品。
	OutSkuId       string `json:"out_sku_id,omitempty"`      // 商家侧小程序商品ID，指商家提报给小程序商品库的商品。
}

type FulfillmentDetail struct {
	FulfillmentAmount string `json:"fulfillment_amount"` // 履约金额
	OutRequestNo      string `json:"out_request_no"`     // 商户发起履约请求时，传入的out_request_no，标识一次请求的唯一id
	GmtPayment        string `json:"gmt_payment"`        // 履约支付时间
}

type BkAgentRespInfo struct {
	BindtrxId        string `json:"bindtrx_id,omitempty"`        // 原快捷交易流水号
	BindclrissrId    string `json:"bindclrissr_id,omitempty"`    // 枚举值，01 银联；02 网联；03 连通等
	BindpyeracctbkId string `json:"bindpyeracctbk_id,omitempty"` // 付款机构在清算组织登记或分配的机构代码
	BkpyeruserCode   string `json:"bkpyeruser_code,omitempty"`   // 用户在银行付款账号的标记化处理编号
	EstterLocation   string `json:"estter_location,omitempty"`   // 设备推测位置 +37.28/-121.268
}

type ChargeInfo struct {
	ChargeFee               string    `json:"charge_fee,omitempty"`                  // 实收费用。单位：元。
	OriginalChargeFee       string    `json:"original_charge_fee,omitempty"`         // 原始费用。单位：元。
	SwitchFeeRate           string    `json:"switch_fee_rate,omitempty"`             // 签约费率
	IsRatingOnTradeReceiver string    `json:"is_rating_on_trade_receiver,omitempty"` // 是否收款账号出资，值为"Y"或"N"
	IsRatingOnSwitch        string    `json:"is_rating_on_switch,omitempty"`         // 是否合约指定收费账号出资，值为"Y"或"N"
	ChargeType              string    `json:"charge_type,omitempty"`                 // 收单手续费trade，花呗分期手续hbfq，其他手续费charge
	SubFeeDetailList        []*SubFee `json:"sub_fee_detail_list,omitempty"`         // 组合支付收费明细
}

type SubFee struct {
	ChargeFee         string `json:"charge_fee,omitempty"`          // 实收费用。单位：元。
	OriginalChargeFee string `json:"original_charge_fee,omitempty"` // 原始费用。单位：元。
	SwitchFeeRate     string `json:"switch_fee_rate,omitempty"`     // 签约费率
}

type TradeSettleInfo struct {
	//trade_unsettled_amount｜剩余待结算金额
	TradeUnsettledAmount  string               `json:"trade_unsettled_amount,omitempty"`
	TradeSettleDetailList []*TradeSettleDetail `json:"trade_settle_detail_list,omitempty"`
}

type TradeSettleDetail struct {
	OperationType     string `json:"operation_type,omitempty"`
	OperationSerialNo string `json:"operation_serial_no,omitempty"`
	OperationDt       string `json:"operation_dt,omitempty"`
	TransOut          string `json:"trans_out,omitempty"`
	TransIn           string `json:"trans_in,omitempty"`
	Amount            string `json:"amount,omitempty"`
	OriTransOut       string `json:"ori_trans_out,omitempty"`
	OriTransIn        string `json:"ori_trans_in,omitempty"`
}

type HbFqPayInfo struct {
	UserInstallNum string `json:"user_install_num,omitempty"`
}

type RefundChargeInfo struct {
	RefundChargeFee        string          `json:"refund_charge_fee,omitempty"`          // 实退费用。单位：元。
	SwitchFeeRate          string          `json:"switch_fee_rate,omitempty"`            // 签约费率
	ChargeType             string          `json:"charge_type,omitempty"`                // 收单手续费trade，花呗分期手续hbfq，其他手续费charge
	RefundSubFeeDetailList []*RefundSubFee `json:"refund_sub_fee_detail_list,omitempty"` // 组合支付退费明细
}

type RefundSubFee struct {
	RefundChargeFee string `json:"refund_charge_fee,omitempty"`
	SwitchFeeRate   string `json:"switch_fee_rate,omitempty"`
}

type RefundPresetPaytool struct {
	Amount         []string `json:"amount,omitempty"`
	AssertTypeCode string   `json:"assert_type_code,omitempty"`
}

type RefundRoyalty struct {
	RefundAmount  string `json:"refund_amount"`             // 退分账金额。单位：元。
	ResultCode    string `json:"result_code"`               // 退分账结果码
	RoyaltyType   string `json:"royalty_type,omitempty"`    // 分账类型. 字段为空默认为普通分账类型transfer
	TransOut      string `json:"trans_out,omitempty"`       // 转出人支付宝账号对应用户ID
	TransOutEmail string `json:"trans_out_email,omitempty"` // 转出人支付宝账号
	TransIn       string `json:"trans_in,omitempty"`        // 转入人支付宝账号对应用户ID
	TransInEmail  string `json:"trans_in_email,omitempty"`  // 转入人支付宝账号
	OriTransOut   string `json:"ori_trans_out,omitempty"`   // 商户请求的转出账号
	OriTransIn    string `json:"ori_trans_in,omitempty"`    // 商户请求的转入账号
}

type DepositBackInfo struct {
	HasDepositBack     string `json:"has_deposit_back,omitempty"`      // 是否存在银行卡冲退信息。
	DbackStatus        string `json:"dback_status,omitempty"`          // 银行卡冲退状态。S-成功，F-失败，P-处理中。银行卡冲退失败，资金自动转入用户支付宝余额。
	DbackAmount        string `json:"dback_amount,omitempty"`          // 银行卡冲退金额
	BankAckTime        string `json:"bank_ack_time,omitempty"`         // 银行响应时间，格式为yyyy-MM-dd HH:mm:ss
	EstBankReceiptTime string `json:"est_bank_receipt_time,omitempty"` // 预估银行到账时间，格式为yyyy-MM-dd HH:mm:ss
}

type WaitRepaymentOrderInfo struct {
	AdvanceOrderId      string `json:"advance_order_id"` // 垫资单id
	AlipayUserId        string `json:"alipay_user_id,omitempty"`
	OpenId              string `json:"open_id"`
	OrigBizOrderId      string `json:"orig_biz_order_id"`     // 原始的业务单号，通常为支付宝交易号
	BizProduct          string `json:"biz_product"`           // 通常为商户签约的收单产品码
	WaitRepaymentAmount string `json:"wait_repayment_amount"` // 垫资金额
}

type UserRiskPrediction struct {
	RefusedPaymentRiskLevel string `json:"refused_payment_risk_level,omitempty"` // 用户拒付风险等级。
	PhoneRecycleRiskLevel   string `json:"phone_recycle_risk_leve,omitempty"`    // 用户绑定手机号被二次放号风险等级。
}
