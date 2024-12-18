package entity

type FundAccountQueryResponse struct {
	Response     *FundAccountQuery `json:"alipay_fund_account_query_response"`
	AlipayCertSn string            `json:"alipay_cert_sn,omitempty"`
	SignData     string            `json:"-"`
	Sign         string            `json:"sign"`
}

type FundAccountQuery struct {
	ErrorResponse
	AvailableAmount string       `json:"available_amount,omitempty"`
	FreezeAmount    string       `json:"freeze_amount,omitempty"`
	ExtCardInfo     *ExtCardInfo `json:"ext_card_info,omitempty"`
}

// =========================================================分割=========================================================

type FundQuotaQueryResponse struct {
	Response     *FundQuotaQuery `json:"alipay_fund_quota_query_response"`
	AlipayCertSn string          `json:"alipay_cert_sn,omitempty"`
	SignData     string          `json:"-"`
	Sign         string          `json:"sign"`
}

type FundQuotaQuery struct {
	ErrorResponse
	ToCorporateDailyAvailableAmount   string `json:"to_corporate_daily_available_amount"`   // 对公日可用额度，单位为元，精确到小数点后两位
	ToPrivateDailyAvailableAmount     string `json:"to_private_daily_available_amount"`     // 对私日可用额度，单位为元，精确到小数点后两位
	ToCorporateMonthlyAvailableAmount string `json:"to_corporate_monthly_available_amount"` // 对公月可用额度，单位为元，精确到小数点后两位
	ToPrivateMonthlyAvailableAmount   string `json:"to_private_monthly_available_amount"`   // 对私月可用额度，单位为元，精确到小数点后两位
}

// =========================================================分割=========================================================

type FundTransUniTransferResponse struct {
	Response     *TransUniTransfer `json:"alipay_fund_trans_uni_transfer_response"`
	AlipayCertSn string            `json:"alipay_cert_sn,omitempty"`
	SignData     string            `json:"-"`
	Sign         string            `json:"sign"`
}

type TransUniTransfer struct {
	ErrorResponse
	OutBizNo       string `json:"out_biz_no,omitempty"`
	OrderId        string `json:"order_id,omitempty"`
	PayFundOrderId string `json:"pay_fund_order_id,omitempty"`
	Status         string `json:"status,omitempty"`
	TransDate      string `json:"trans_date,omitempty"`
}

// =========================================================分割=========================================================

type FundTransCommonQueryResponse struct {
	Response     *FundTransCommonQuery `json:"alipay_fund_trans_common_query_response"`
	AlipayCertSn string                `json:"alipay_cert_sn,omitempty"`
	SignData     string                `json:"-"`
	Sign         string                `json:"sign"`
}

type FundTransCommonQuery struct {
	ErrorResponse
	OrderId        string `json:"order_id,omitempty"`
	PayFundOrderId string `json:"pay_fund_order_id,omitempty"`
	OutBizNo       string `json:"out_biz_no,omitempty"`
	TransAmount    string `json:"trans_amount,omitempty"`
	Status         string `json:"status,omitempty"`
	PayDate        string `json:"pay_date,omitempty"`
	ArrivalTimeEnd string `json:"arrival_time_end,omitempty"`
	OrderFee       string `json:"order_fee,omitempty"`
	ErrorCode      string `json:"error_code,omitempty"`
	FailReason     string `json:"fail_reason,omitempty"`
}

// =========================================================分割=========================================================

type FundTransOrderQueryResponse struct {
	Response     *FundTransOrderQuery `json:"alipay_fund_trans_order_query_response"`
	AlipayCertSn string               `json:"alipay_cert_sn,omitempty"`
	SignData     string               `json:"-"`
	Sign         string               `json:"sign"`
}

type FundTransOrderQuery struct {
	ErrorResponse
	OrderId        string `json:"order_id,omitempty"`
	Status         string `json:"status,omitempty"`
	PayDate        string `json:"pay_date,omitempty"`
	ArrivalTimeEnd string `json:"arrival_time_end,omitempty"`
	OrderFee       string `json:"order_fee,omitempty"`
	FailReason     string `json:"fail_reason,omitempty"`
	OutBizNo       string `json:"out_biz_no,omitempty"`
	ErrorCode      string `json:"error_code,omitempty"`
}

// =========================================================分割=========================================================

type FundTransRefundResponse struct {
	Response     *FundTransRefund `json:"alipay_fund_trans_refund_response"`
	AlipayCertSn string           `json:"alipay_cert_sn,omitempty"`
	SignData     string           `json:"-"`
	Sign         string           `json:"sign"`
}

type FundTransRefund struct {
	ErrorResponse
	RefundOrderId string `json:"refund_order_id"`
	OrderId       string `json:"order_id"`
	OutRequestNo  string `json:"out_request_no"`
	Status        string `json:"status"`
	RefundAmount  string `json:"refund_amount"`
	RefundDate    string `json:"refund_date"`
}

// =========================================================分割=========================================================

type FundAuthOrderFreezeResponse struct {
	Response     *FundAuthOrderFreeze `json:"alipay_fund_auth_order_freeze_response"`
	AlipayCertSn string               `json:"alipay_cert_sn,omitempty"`
	SignData     string               `json:"-"`
	Sign         string               `json:"sign"`
}

type FundAuthOrderFreeze struct {
	ErrorResponse
	AuthNo        string `json:"auth_no"`
	OutOrderNo    string `json:"out_order_no"`
	OperationId   string `json:"operation_id"`
	OutRequestNo  string `json:"out_request_no"`
	Amount        string `json:"amount"`
	Status        string `json:"status"`
	PayerUserId   string `json:"payer_user_id,omitempty"`
	PayerOpenId   string `json:"payer_open_id,omitempty"`
	PayerLogonId  string `json:"payer_logon_id,omitempty"`
	GmtTrans      string `json:"gmt_trans,omitempty"`      // 资金授权成功时间 格式：YYYY-MM-DD HH:MM:SS
	PreAuthType   string `json:"pre_auth_type,omitempty"`  // 预授权类型，目前支持 CREDIT_AUTH(信用预授权)
	TransCurrency string `json:"trans_currency,omitempty"` // 标价币种, amount 对应的币种单位。支持澳元：AUD, 新西兰元：NZD, 台币：TWD, 美元：USD, 欧元：EUR, 英镑：GBP
	CreditAmount  string `json:"credit_amount,omitempty"`
	FundAmount    string `json:"fund_amount,omitempty"`
}

// =========================================================分割=========================================================

type FundAuthOrdervoucherCreateResponse struct {
	Response     *FundAuthOrdervoucherCreate `json:"alipay_fund_auth_order_voucher_create_response"`
	AlipayCertSn string                      `json:"alipay_cert_sn,omitempty"`
	SignData     string                      `json:"-"`
	Sign         string                      `json:"sign"`
}

type FundAuthOrdervoucherCreate struct {
	ErrorResponse
	OutOrderNo   string `json:"out_order_no"`
	OutRequestNo string `json:"out_request_no"`
	CodeType     string `json:"code_type"`  // 码类型，分为 barCode：条形码 (一维码) 和 qrCode:二维码(qrCode) ； 目前发码只支持 qrCode
	CodeValue    string `json:"code_value"` // 当前发码请求生成的二维码码串，商户端可以利用二维码生成工具根据该码串值生成对应的二维码
	CodeUrl      string `json:"code_url"`   // 生成的带有支付宝logo的二维码地址
}

// =========================================================分割=========================================================

type FundAuthOrderUnfreezeResponse struct {
	Response     *FundAuthOrderUnfreeze `json:"alipay_fund_auth_order_unfreeze_response"`
	AlipayCertSn string                 `json:"alipay_cert_sn,omitempty"`
	SignData     string                 `json:"-"`
	Sign         string                 `json:"sign"`
}

type FundAuthOrderUnfreeze struct {
	ErrorResponse
	AuthNo       string `json:"auth_no"`
	OutOrderNo   string `json:"out_order_no,omitempty"`
	OperationId  string `json:"operation_id"`
	OutRequestNo string `json:"out_request_no"`
	Amount       string `json:"amount"`
	Status       string `json:"status"`
	GmtTrans     string `json:"gmt_trans,omitempty"`
	CreditAmount string `json:"credit_amount,omitempty"`
	FundAmount   string `json:"fund_amount,omitempty"`
}

// =========================================================分割=========================================================

type FundAuthOperationDetailQueryResponse struct {
	Response     *FundAuthOperationDetailQuery `json:"alipay_fund_auth_operation_detail_query_response"`
	AlipayCertSn string                        `json:"alipay_cert_sn,omitempty"`
	SignData     string                        `json:"-"`
	Sign         string                        `json:"sign"`
}

type FundAuthOperationDetailQuery struct {
	ErrorResponse
	AuthNo                  string `json:"auth_no"`             // 支付宝资金授权订单号
	OutOrderNo              string `json:"out_order_no"`        // 商户的授权资金订单号
	OrderStatus             string `json:"order_status"`        // 授权单状态：INIT（初始状态：已创建未授权）、AUTHORIZED（已授权状态：授权成功，可以进行转支付或解冻操作）、FINISH（完成状态：转支付完成且无剩余冻结资金）、CLOSED（关闭状态：授权未完成超时关闭或冻结资金全额解冻）
	TotalFreezeAmount       string `json:"total_freeze_amount"` // 订单累计的冻结金额，单位为：元（人民币）
	RestAmount              string `json:"rest_amount"`         // 订单当前剩余冻结金额，单位为：元（人民币）
	TotalPayAmount          string `json:"total_pay_amount"`    // 订单累计用于支付的金额，单位为：元（人民币）
	OrderTitle              string `json:"order_title"`         // 业务订单的简单描述，如商品名称等
	OperationId             string `json:"operation_id"`        // 支付宝资金操作流水号
	OutRequestNo            string `json:"out_request_no"`      // 商户资金操作的请求流水号
	Amount                  string `json:"amount"`              // 该笔资金操作流水operation_id对应的操作金额，单位为：元（人民币）
	PayerLogonId            string `json:"payer_logon_id,omitempty"`
	PayerUserId             string `json:"payer_user_id,omitempty"`
	PayerOpenId             string `json:"payer_open_id,omitempty"`
	ExtraParam              string `json:"extra_param,omitempty"`                // 商户请求创建预授权订单时传入的扩展参数，仅返回商户自定义的扩展信息（merchantExt）
	OperationType           string `json:"operation_type"`                       // 支付宝资金操作类型，表示当前查询到的这笔明细的操作类型。
	Status                  string `json:"status"`                               // 资金操作流水的状态， 目前支持： INIT：初始 SUCCESS：成功 CLOSED：关闭
	Remark                  string `json:"remark"`                               // 商户对本次操作的附言描述，长度不超过100个字母或50个汉字
	GmtCreate               string `json:"gmt_create"`                           // 资金授权单据操作流水创建时间， 格式：YYYY-MM-DD HH:MM:SS
	GmtTrans                string `json:"gmt_trans,omitempty"`                  // 支付宝账务处理成功时间， 格式：YYYY-MM-DD HH:MM:SS
	PreAuthType             string `json:"pre_auth_type,omitempty"`              // 预授权类型，信用预授权情况下值为 CREDIT_AUTH，表示该笔预授权为信用预授权，实际没有冻结用户资金；其它情况均不返回该字段。
	TransCurrency           string `json:"trans_currency,omitempty"`             // 标价币种, amount 对应的币种单位。支持澳元：AUD, 新西兰元：NZD, 台币：TWD, 美元：USD, 欧元：EUR, 英镑：GBP
	TotalFreezeCreditAmount string `json:"total_freeze_credit_amount,omitempty"` // 累计冻结信用金额，单位为：元（人民币），精确到小数点后两位
	TotalFreezeFundAmount   string `json:"total_freeze_fund_amount,omitempty"`   // 累计冻结自有资金金额，单位为：元（人民币），精确到小数点后两位
	TotalPayCreditAmount    string `json:"total_pay_credit_amount,omitempty"`    // 累计支付信用金额，单位为：元（人民币），精确到小数点后两位
	TotalPayFundAmount      string `json:"total_pay_fund_amount,omitempty"`      // 累计支付自有资金金额，单位为：元（人民币），精确到小数点后两位
	RestCreditAmount        string `json:"rest_credit_amount,omitempty"`         // 剩余冻结信用金额，单位为：元（人民币），精确到小数点后两位
	RestFundAmount          string `json:"rest_fund_amount,omitempty"`           // 剩余冻结自有资金金额，单位为：元（人民币），精确到小数点后两位
	CreditAmount            string `json:"credit_amount,omitempty"`              // 该笔资金操作流水operation_id对应的操作信用金额
	FundAmount              string `json:"fund_amount,omitempty"`                // 该笔资金操作流水operation_id对应的操作自有资金金额
	CreditMerchantExt       string `json:"credit_merchant_ext,omitempty"`        // 芝麻透出给商户的信息，具体内容由商户与芝麻约定后返回
}

// =========================================================分割=========================================================

type FundAuthOperationCancelResponse struct {
	Response     *FundAuthOperationCancel `json:"alipay_fund_auth_operation_cancel_response"`
	AlipayCertSn string                   `json:"alipay_cert_sn,omitempty"`
	SignData     string                   `json:"-"`
	Sign         string                   `json:"sign"`
}

type FundAuthOperationCancel struct {
	ErrorResponse
	AuthNo       string `json:"auth_no"`      // 商户的授权资金订单号。
	OutOrderNo   string `json:"out_order_no"` // 商户的冻结操作流水号 。
	OperationId  string `json:"operation_id,omitempty"`
	OutRequestNo string `json:"out_request_no,omitempty"`
	Action       string `json:"action,omitempty"` // 本次撤销触发的资金动作 close：关闭冻结明细，无资金解冻 unfreeze：产生了资金解冻
}

// =========================================================分割=========================================================

type FundBatchCreateResponse struct {
	Response     *FundBatchCreate `json:"alipay_fund_batch_create_response"`
	AlipayCertSn string           `json:"alipay_cert_sn,omitempty"`
	SignData     string           `json:"-"`
	Sign         string           `json:"sign"`
}

type FundBatchCreate struct {
	ErrorResponse
	OutBatchNo   string `json:"out_batch_no,omitempty"`
	BatchTransId string `json:"batch_trans_id,omitempty"`
	Status       string `json:"status,omitempty"`
}

// =========================================================分割=========================================================

type FundBatchCloseResponse struct {
	Response     *FundBatchClose `json:"alipay_fund_batch_close_response"`
	AlipayCertSn string          `json:"alipay_cert_sn,omitempty"`
	SignData     string          `json:"-"`
	Sign         string          `json:"sign"`
}

type FundBatchClose struct {
	ErrorResponse
	BatchTransId string `json:"batch_trans_id,omitempty"`
	Status       string `json:"status,omitempty"`
}

// =========================================================分割=========================================================

type FundBatchDetailQueryResponse struct {
	Response     *FundBatchDetailQuery `json:"alipay_fund_batch_detail_query_response"`
	AlipayCertSn string                `json:"alipay_cert_sn,omitempty"`
	SignData     string                `json:"-"`
	Sign         string                `json:"sign"`
}

type FundBatchDetailQuery struct {
	ErrorResponse
	BatchTransId    string       `json:"batch_trans_id,omitempty"`
	BatchNo         string       `json:"batch_no,omitempty"`
	BizCode         string       `json:"biz_code,omitempty"`
	BizScene        string       `json:"biz_scene,omitempty"`
	BatchStatus     string       `json:"batch_status,omitempty"`
	ApprovalStatus  string       `json:"approval_status,omitempty"`
	ErrorCode       string       `json:"error_code,omitempty"`
	FailReason      string       `json:"fail_reason,omitempty"`
	SignPrincipal   string       `json:"sign_principal,omitempty"`
	PaymentAmount   string       `json:"payment_amount,omitempty"`
	PaymentCurrency string       `json:"payment_currency,omitempty"`
	PageSize        int          `json:"page_size,omitempty"`
	PageNum         int          `json:"page_num,omitempty"`
	ProductCode     string       `json:"product_code,omitempty"`
	TotalPageCount  string       `json:"total_page_count,omitempty"`
	OutBatchNo      string       `json:"out_batch_no,omitempty"`
	GmtFinish       string       `json:"gmt_finish,omitempty"`
	TotalAmount     string       `json:"total_amount,omitempty"`
	GmtPayFinish    string       `json:"gmt_pay_finish,omitempty"`
	PayerId         string       `json:"payer_id,omitempty"`
	SuccessAmount   string       `json:"success_amount,omitempty"`
	FailAmount      string       `json:"fail_amount,omitempty"`
	FailCount       string       `json:"fail_count,omitempty"`
	SuccessCount    string       `json:"success_count,omitempty"`
	TotalItemCount  string       `json:"total_item_count,omitempty"`
	AccDetailList   []*AccDetail `json:"acc_detail_list,omitempty"`
}

// =========================================================分割=========================================================

type FundTransAppPayResponse struct {
	Response     *FundTransAppPay `json:"alipay_fund_trans_app_pay_response"`
	AlipayCertSn string           `json:"alipay_cert_sn,omitempty"`
	SignData     string           `json:"-"`
	Sign         string           `json:"sign"`
}

type FundTransAppPay struct {
	ErrorResponse
	OutBizNo string `json:"out_biz_no,omitempty"`
	OrderId  string `json:"order_id,omitempty"`
	Status   string `json:"status,omitempty"`
}

// =========================================================分割=========================================================

type FundTransPayeeBindQueryResponse struct {
	Response     *FundTransPayeeBindQuery `json:"alipay_fund_trans_payee_bind_query_response"`
	AlipayCertSn string                   `json:"alipay_cert_sn,omitempty"`
	SignData     string                   `json:"-"`
	Sign         string                   `json:"sign"`
}

type FundTransPayeeBindQuery struct {
	ErrorResponse
	Bind string `json:"bind"` // 是否绑定收款账号。true：已绑定；false：未绑定
}

// =========================================================分割=========================================================

type FundTransPagePayResponse struct {
	Response     *FundTransPagePay `json:"alipay_fund_trans_page_pay_response"`
	AlipayCertSn string            `json:"alipay_cert_sn,omitempty"`
	SignData     string            `json:"-"`
	Sign         string            `json:"sign"`
}

type FundTransPagePay struct {
	ErrorResponse
	OutBizNo string `json:"out_biz_no"`
	OrderID  string `json:"order_id,omitempty"`
	Status   string `json:"status"`
}

// =========================================================分割=========================================================

type FundAuthOrderAppFreezeResponse struct {
	Response     *FundAuthOrderAppFreeze `json:"alipay_fund_auth_order_app_freeze_response"`
	AlipayCertSn string                  `json:"alipay_cert_sn,omitempty"`
	SignData     string                  `json:"-"`
	Sign         string                  `json:"sign"`
}

type FundAuthOrderAppFreeze struct {
	ErrorResponse
	AuthNo        string `json:"auth_no,omitempty"`
	OutOrderNo    string `json:"out_order_no,omitempty"`
	OperationId   string `json:"operation_id,omitempty"`
	OutRequestNo  string `json:"out_request_no,omitempty"`
	Amount        string `json:"amount,omitempty"`
	Status        string `json:"status,omitempty"`
	PayerUserId   string `json:"payer_user_id,omitempty"`
	GmtTrans      string `json:"gmt_trans,omitempty"`
	PreAuthType   string `json:"pre_auth_type,omitempty"`
	CreditAmount  string `json:"credit_amount,omitempty"`
	FundAmount    string `json:"fund_amount,omitempty"`
	TransCurrency string `json:"trans_currency,omitempty"`
}

// =========================================================分割=========================================================

type ExtCardInfo struct {
	CardNo       string `json:"card_no,omitempty"`
	BankAccName  string `json:"bank_acc_name,omitempty"`
	CardBranch   string `json:"card_branch,omitempty"`
	CardBank     string `json:"card_bank,omitempty"`
	CardLocation string `json:"card_location,omitempty"`
	CardDeposit  string `json:"card_deposit,omitempty"`
	Status       string `json:"status,omitempty"`
}

type AccDetail struct {
	DetailNo           string `json:"detail_no,omitempty"`
	PaymentAmount      string `json:"payment_amount,omitempty"`
	PaymentCurrency    string `json:"payment_currency,omitempty"`
	TransAmount        string `json:"trans_amount,omitempty"`
	TransCurrency      string `json:"trans_currency,omitempty"`
	SettlementAmount   string `json:"settlement_amount,omitempty"`
	SettlementCurrency string `json:"settlement_currency,omitempty"`
	PayeeInfo          *struct {
		PayeeAccount string `json:"payee_account,omitempty"`
		PayeeType    string `json:"payee_type,omitempty"`
		PayeeName    string `json:"payee_name,omitempty"`
	} `json:"payee_info,omitempty"`
	CertInfo *struct {
		CertNo   string `json:"cert_no,omitempty"`
		CertType string `json:"cert_type,omitempty"`
	} `json:"cert_info,omitempty"`
	Remark       string `json:"remark,omitempty"`
	Status       string `json:"status,omitempty"`
	ExchangeRate *struct {
		Rate             string `json:"rate,omitempty"`
		BaseCurrency     string `json:"base_currency,omitempty"`
		ExchangeCurrency string `json:"exchange_currency,omitempty"`
	} `json:"exchange_rate,omitempty"`
	NeedRetry     string `json:"need_retry,omitempty"`
	AlipayOrderNo string `json:"alipay_order_no,omitempty"`
	OutBizNo      string `json:"out_biz_no,omitempty"`
	DetailId      string `json:"detail_id,omitempty"`
	ErrorCode     string `json:"error_code,omitempty"`
	ErrorMsg      string `json:"error_msg,omitempty"`
	GmtCreate     string `json:"gmt_create,omitempty"`
	GmtFinish     string `json:"gmt_finish,omitempty"`
	SubStatus     string `json:"sub_status,omitempty"`
}
