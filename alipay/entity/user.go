package entity

type UserAgreementPageUnSignResponse struct {
	Response     *UserAgreementPageUnSign `json:"alipay_user_agreement_unsign_response"`
	AlipayCertSn string                   `json:"alipay_cert_sn,omitempty"`
	SignData     string                   `json:"-"`
	Sign         string                   `json:"sign"`
}

type UserAgreementPageUnSign struct {
	ErrorResponse
}

// =========================================================分割=========================================================

type UserAgreementExecutionplanModifyResponse struct {
	Response     *UserAgreementExecutionplanModify `json:"alipay_user_agreement_executionplan_modify_response"`
	AlipayCertSn string                            `json:"alipay_cert_sn,omitempty"`
	SignData     string                            `json:"-"`
	Sign         string                            `json:"sign"`
}

type UserAgreementExecutionplanModify struct {
	ErrorResponse
	AgreementNo string `json:"agreement_no"` // 周期性扣款产品，授权免密支付协议号
	DeductTime  string `json:"deduct_time"`  // 商户下一次扣款时间，格式 "yyyy-MM-dd"。 例如：用户在1月1日开通了连续包月，使用了10天又另行购买了“季度包”，如果此时商户希望“季度包”立即优先生效，在季度包结束后能继续使用连续包月，那么原定的周期就被延后了。此时可以通过本接口将预计扣款时间推后“季度包”的时长。
}

// =========================================================分割=========================================================

type UserAgreementQueryResponse struct {
	Response     *UserAgreementQuery `json:"alipay_user_agreement_query_response"`
	AlipayCertSn string              `json:"alipay_cert_sn,omitempty"`
	SignData     string              `json:"-"`
	Sign         string              `json:"sign"`
}

type UserAgreementQuery struct {
	ErrorResponse
	ValidTime           string `json:"valid_time"`                 // 协议生效时间，格式为 yyyyMM-dd HH:mm:ss
	AlipayLogonId       string `json:"alipay_logon_id"`            // 返回脱敏的支付宝账号
	InvalidTime         string `json:"invalid_time"`               // 协议失效时间，格式为 yyyyMM-dd HH:mm:ss
	PricipalType        string `json:"pricipal_type"`              // 签约主体类型。 CARD:支付宝账号 CUSTOMER:支付宝用户
	SignScene           string `json:"sign_scene"`                 // 签约协议的场景
	AgreementNo         string `json:"agreement_no"`               // 用户签约成功后的协议号
	ThirdPartyType      string `json:"third_party_type"`           // 签约第三方主体类型。对于三方协议，表示当前用户和哪一类的第三方主体进行签约。 1.PARTNER（平台商户）;2.MERCHANT（集团商户），集团下子商户可共享用户签约内容;默认为PARTNER
	Status              string `json:"status"`                     // 协议当前状态 1.TEMP：暂存，协议未生效过；2.NORMAL：正常；3.STOP：暂停
	SignTime            string `json:"sign_time"`                  // 协议签约时间，格式为 yyyyMM-dd HH:mm:ss
	PersonalProductCode string `json:"personal_product_code"`      // 协议产品码，商户和支付宝签约时确定，不同业务场景对应不同的签约产品码
	PrincipalOpenId     string `json:"principal_open_id"`          // 签约主体标识。当principal_type为CARD时，该字段为支付宝用户号;当principal_type为CUSTOMER时，该字段为支付宝用户标识。
	DeviceId            string `json:"device_id"`                  // 设备Id
	ExternalAgreementNo string `json:"external_agreement_no"`      // 代扣协议中标示用户的唯一签约号(确保在商户系统中唯一)
	ZmOpenId            string `json:"zm_open_id"`                 // 用户的芝麻信用 openId，供商户查询用户芝麻信用使用。
	ExternalLogonId     string `json:"external_logon_id"`          // 外部登录Id
	CreditAuthMode      string `json:"credit_auth_mode"`           // 授信模式，取值：DEDUCT_HUAZHI-花芝GO。目前只在花芝代扣（即花芝go）协议时才会返回
	SingleQuota         string `json:"single_quota"`               // 单笔代扣额度
	LastDeductTime      string `json:"last_deduct_time,omitempty"` // 周期扣协议，上次扣款成功时间
	NextDeductTime      string `json:"next_deduct_time,omitempty"` // 周期扣协议，预计下次扣款时间
	ExecutionPlans      []struct {
		SingleAmount      string `json:"single_amount"`                 // 单笔金额
		PeriodId          string `json:"period_id,omitempty"`           // 期数
		ExecuteTime       string `json:"execute_time,omitempty"`        // 预期执行时间
		LatestExecuteTime string `json:"latest_execute_time,omitempty"` // 最晚执行时间
	} `json:"execution_plans,omitempty"`
}

// =========================================================分割=========================================================

type UserAgreementTransferResponse struct {
	Response     *UserAgreementTransfer `json:"alipay_user_agreement_transfer_response"`
	AlipayCertSn string                 `json:"alipay_cert_sn,omitempty"`
	SignData     string                 `json:"-"`
	Sign         string                 `json:"sign"`
}

type UserAgreementTransfer struct {
	ErrorResponse
	ExecuteTime   string `json:"execute_time,omitempty"`
	PeriodType    string `json:"period_type,omitempty"`
	Amount        string `json:"amount,omitempty"`
	TotalAmount   string `json:"total_amount,omitempty"`
	TotalPayments string `json:"total_payments,omitempty"`
	Period        string `json:"period,omitempty"`
}
