package entity

type DataBillBalanceQueryResponse struct {
	Response     *DataBillBalanceQuery `json:"alipay_data_bill_balance_query_response"`
	AlipayCertSn string                `json:"alipay_cert_sn,omitempty"`
	SignData     string                `json:"-"`
	Sign         string                `json:"sign"`
}
type DataBillBalancehisQueryResponse struct {
	Response     *DataBillBalancehisQuery `json:"alipay_data_bill_balancehis_query_response"`
	AlipayCertSn string                   `json:"alipay_cert_sn,omitempty"`
	SignData     string                   `json:"-"`
	Sign         string                   `json:"sign"`
}

type DataBillAccountLogQueryResponse struct {
	Response     *DataBillAccountLogQuery `json:"alipay_data_bill_accountlog_query_response"`
	AlipayCertSn string                   `json:"alipay_cert_sn,omitempty"`
	SignData     string                   `json:"-"`
	Sign         string                   `json:"sign"`
}

type DataBillEreceiptApplyResponse struct {
	Response     *DataBillEreceiptApply `json:"alipay_data_bill_ereceipt_apply_response"`
	AlipayCertSn string                 `json:"alipay_cert_sn,omitempty"`
	SignData     string                 `json:"-"`
	Sign         string                 `json:"sign"`
}

type DataBillEreceiptQueryResponse struct {
	Response     *DataBillEreceiptQuery `json:"alipay_data_bill_ereceipt_query_response"`
	AlipayCertSn string                 `json:"alipay_cert_sn,omitempty"`
	SignData     string                 `json:"-"`
	Sign         string                 `json:"sign"`
}

type DataBillDownloadUrlQueryResponse struct {
	Response     *DataBillDownloadUrlQuery `json:"alipay_data_dataservice_bill_downloadurl_query_response"`
	AlipayCertSn string                    `json:"alipay_cert_sn,omitempty"`
	SignData     string                    `json:"-"`
	Sign         string                    `json:"sign"`
}

type DataBillSellQueryResponse struct {
	Response     *DataBillSellQuery `json:"alipay_data_bill_sell_query_response"`
	AlipayCertSn string             `json:"alipay_cert_sn,omitempty"`
	SignData     string             `json:"-"`
	Sign         string             `json:"sign"`
}

type DataBillBuyQueryResponse struct {
	Response     *DataBillBuyQuery `json:"alipay_data_bill_buy_query_response"`
	AlipayCertSn string            `json:"alipay_cert_sn,omitempty"`
	SignData     string            `json:"-"`
	Sign         string            `json:"sign"`
}

type DataBillTransferQueryResponse struct {
	Response     *DataBillTransferQuery `json:"alipay_data_bill_transfer_query_response"`
	AlipayCertSn string                 `json:"alipay_cert_sn,omitempty"`
	SignData     string                 `json:"-"`
	Sign         string                 `json:"sign"`
}

type DataBillBailQueryResponse struct {
	Response     *DataBillBailQuery `json:"alipay_data_bill_bail_query_response"`
	AlipayCertSn string             `json:"alipay_cert_sn,omitempty"`
	SignData     string             `json:"-"`
	Sign         string             `json:"sign"`
}

// =========================================================分割=========================================================

type DataBillBalanceQuery struct {
	ErrorResponse
	TotalAmount     string `json:"total_amount"`            // 支付宝账户余额
	AvailableAmount string `json:"available_amount"`        // 账户可用余额
	FreezeAmount    string `json:"freeze_amount"`           // 冻结金额。单位（元）
	SettleAmount    string `json:"settle_amount,omitempty"` // 当前账户的待结算金额，单位（元）
}

type DataBillBalancehisQuery struct {
	ErrorResponse
	BeginningBalance string `json:"beginning_balance"` // 期初余额
	EndingBalance    string `json:"ending_balance"`    // 期末余额
}

type DataBillAccountLogQuery struct {
	ErrorResponse
	PageNo     string                  `json:"page_no"`    // 分页号，从1开始
	PageSize   string                  `json:"page_size"`  // 分页大小1000-2000
	TotalSize  string                  `json:"total_size"` // 账务明细总数。返回满足查询条件的明细的数量
	DetailList []*AccountLogItemResult `json:"detail_list,omitempty"`
}

type AccountLogItemResult struct {
	TransDt             string `json:"trans_dt"`                         // 入账时间
	AccountLogId        string `json:"account_log_id"`                   // 支付宝账务流水号。对账使用，不脱敏
	AlipayOrderNo       string `json:"alipay_order_no"`                  // 支付宝订单号。对账使用，不脱敏
	MerchantOrderNo     string `json:"merchant_order_no"`                // 商户订单号，创建支付宝交易时传入的信息。对账使用，不脱敏
	TransAmount         string `json:"trans_amount"`                     // 金额
	Balance             string `json:"balance"`                          // 余额 当前不是最终结果，具有不确定性。最终一致
	Type                string `json:"type"`                             // 账务记录的类型，仅供参考
	OtherAccount        string `json:"other_account"`                    // 对方账户
	TransMemo           string `json:"trans_memo"`                       // 收入/支出。表示收支。amount是正数，返回“收入”。amount是负数，返回“支出”
	Direction           string `json:"direction,omitempty"`              // 账务备注。由上游业务决定，不可依赖此字段进行对账
	BillSource          string `json:"bill_source,omitempty"`            // 业务账单来源，资金收支对应的上游业务订单数据来源，确认业务订单出处。此字段供商户对账使用，不脱敏。
	BizNos              string `json:"biz_nos,omitempty"`                // 业务订单号，资金收支相关的业务场景订单号明细，字母大写；M：平台交易主单号，S：平台交易子单号，O：业务系统单据号（如退款订单号）。此字段供商户对账使用，不脱敏。
	BizOrigNo           string `json:"biz_orig_no,omitempty"`            // 业务基础订单号，资金收支对应的原始业务订单唯一识别编号。此字段供商户对账使用，不脱敏。
	BizDesc             string `json:"biz_desc,omitempty"`               // 业务描述，资金收支对应的详细业务场景信息。此字段供商户对账使用，不脱敏。
	MerchantOutRefundNo string `json:"merchant_out_refund_no,omitempty"` // 支付宝交易商户退款请求号。对应商户在调用收单退款接口openApi请求传入的outRequestNo参数值
	ComplementInfo      string `json:"complement_info,omitempty"`        // 账单的补全信息，用于特殊场景，普通商户不需要传值，对账时可忽略。
	StoreName           string `json:"store_name,omitempty"`             // 门店信息
}

type DataBillDownloadUrlQuery struct {
	ErrorResponse
	BillDownloadUrl string `json:"bill_download_url"` // 账单下载地址链接，获取连接后30秒后未下载，链接地址失效。
	BillFileCode    string `json:"bill_file_code"`    // 描述本次申请的账单文件状态。 EMPTY_DATA_WITH_BILL_FILE：当天无账单业务数据&&可以获取到空数据账单文件
}

type DataBillEreceiptApply struct {
	ErrorResponse
	FileId string `json:"file_id"` // 文件申请号file_id信息。使用file_id可以查询处理状态，有效期：2天
}

type DataBillEreceiptQuery struct {
	ErrorResponse
	Status       string `json:"status"`                  // 处理状态。枚举值如下： INIT：初始化。 PROCESS：处理中。 SUCCESS：成功。 FAIL：失败。
	DownloadUrl  string `json:"download_url,omitempty"`  // 下载链接，status 为 SUCCESS 时返回。用户可以使用此http链接下载文件内容。有效时间为 30s。 生成文件为PDF格式，下载即可获取电子回单 PDF 内容。
	ErrorMessage string `json:"error_message,omitempty"` // 如果生成失败，则会返回失败原因
}

type DataBillSellQuery struct {
	ErrorResponse
	PageNo     string             `json:"page_no"`     // 分页号，从1开始
	PageSize   string             `json:"page_size"`   // 分页大小1000-2000
	TotalSize  string             `json:"total_size"`  // 账务明细总数。返回满足查询条件的明细的数量
	DetailList []*TradeItemResult `json:"detail_list"` // 交易流水详情
}

type TradeItemResult struct {
	GmtCreate       string `json:"gmt_create"`              // 交易创建时间
	GmtPay          string `json:"gmt_pay"`                 // 交易支付时间
	AlipayOrderNo   string `json:"alipay_order_no"`         // 支付宝订单号。对账使用，不脱敏
	MerchantOrderNo string `json:"merchant_order_no"`       // 商户订单号，创建支付宝交易时传入的信息。对账使用，不脱敏
	OtherAccount    string `json:"other_account"`           // 对方账户
	GoodsTitle      string `json:"goods_title"`             // 商品名称
	TotalAmount     string `json:"total_amount"`            // 订单金额
	TradeStatus     string `json:"trade_status"`            // 订单状态(待付款,成功,关闭,待发货,待确认收货,已预付,进行中)
	TradeType       string `json:"trade_type"`              // 业务类型，帮助商户作为对账参考
	GmtRefund       string `json:"gmt_refund,omitempty"`    // 交易退款时间
	NetMdiscount    string `json:"net_mdiscount,omitempty"` // 商家优惠金额
	RefundAmount    string `json:"refund_amount,omitempty"` // 订单退款金额
	ServiceFee      string `json:"service_fee,omitempty"`   // 服务费金额
	StoreNo         string `json:"store_no,omitempty"`      // 门店名称
	StoreName       string `json:"store_name,omitempty"`    // 门店名称
	GoodsMemo       string `json:"goods_memo,omitempty"`    // 商品备注信息
}

type DataBillBuyQuery struct {
	ErrorResponse
	PageNo     string             `json:"page_no"`     // 分页号，从1开始
	PageSize   string             `json:"page_size"`   // 分页大小1000-2000
	TotalSize  string             `json:"total_size"`  // 账务明细总数。返回满足查询条件的明细的数量
	DetailList []*TradeItemResult `json:"detail_list"` // 交易流水详情
}

type DataBillTransferQuery struct {
	ErrorResponse
	PageNo     string                  `json:"page_no"`     // 分页号，从1开始
	PageSize   string                  `json:"page_size"`   // 分页大小1000-2000
	TotalSize  string                  `json:"total_size"`  // 账务明细总数。返回满足查询条件的明细的数量
	DetailList []*TransferDetailResult `json:"detail_list"` // 交易流水详情
}

type TransferDetailResult struct {
	TransDt       string `json:"trans_dt"`                 // 业务发生时间
	OrderNo       string `json:"order_no"`                 // 业务订单号。该笔业务单据的唯一识别编号
	TypeDesc      string `json:"type_desc"`                // 查询类型描述：充值、转账、提现
	FundDesc      string `json:"fund_desc"`                // 资金来源/去向类型。在充值记录中，表示资金来源类型，在转账和提现类型中，表示去向类型
	Account       string `json:"account"`                  // 付款/收款账户。充值记录中是付款账户。提现、转账记录中是收款账户。支付宝名称及账号脱敏；银行账户的户名脱敏，银行账户显示银行名称+银行卡号后四位
	Amount        string `json:"amount"`                   // 金额
	Status        string `json:"status"`                   // 资金状态
	SubTypeDesc   string `json:"sub_type_desc,omitempty"`  // 子类型。“充值类型”，普通充值、大额充值。“转账类型”，暂无实现。转账至支付宝账户、转账至银行卡、批量转账支付宝账户、批量转账至银行卡、批量付款。“提现类型”，暂无实现。普通提现、批量委托提现。对账使用，无需脱敏
	ServiceFee    string `json:"service_fee,omitempty"`    // 服务费金额
	InstructionId string `json:"instruction_id,omitempty"` // 银行单据号。对账使用，无需脱敏
	Memo          string `json:"memo,omitempty"`           // 备注信息
}

type DataBillBailQuery struct {
	ErrorResponse
	DetailList []*BailDetailResult `json:"detail_list"` // 保证金明细列表，最多返回5000条
}

type BailDetailResult struct {
	TransDt    string `json:"trans_dt"`              // 业务发生时间
	TransLogId string `json:"trans_log_id"`          // 保证金业务流水号
	BailType   string `json:"bail_type"`             // 保证金类型描述，仅供参考
	Amount     string `json:"amount"`                // 保证金收支金额
	Balance    string `json:"balance"`               // 本次操作后的保证金余额。字段数据展示为"--"，表明数据暂未更新，请稍等1分钟后重试。
	Memo       string `json:"memo,omitempty"`        // 保证金说明
	BizDesc    string `json:"biz_desc,omitempty"`    // 业务描述，资金收支对应的详细业务场景信息
	BizOrigNo  string `json:"biz_orig_no,omitempty"` // 业务基础订单号，资金收支对应的原始业务订单唯一识别编号
}
