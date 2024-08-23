package entity

type ZolozAuthenticationSmilepayInitializeResponse struct {
	Response     *ZolozAuthenticationSmilepayInitialize `json:"zoloz_authentication_smilepay_initialize_response"`
	AlipayCertSn string                                 `json:"alipay_cert_sn,omitempty"`
	SignData     string                                 `json:"-"`
	Sign         string                                 `json:"sign"`
}

type ZolozAuthenticationCustomerFtokenQueryResponse struct {
	Response     *ZolozAuthenticationCustomerFtokenQuery `json:"zoloz_authentication_customer_ftoken_query_response"`
	AlipayCertSn string                                  `json:"alipay_cert_sn,omitempty"`
	SignData     string                                  `json:"-"`
	Sign         string                                  `json:"sign"`
}

// =========================================================分割=========================================================

type ZolozAuthenticationSmilepayInitialize struct {
	ErrorResponse
	RetCodeSub        string `json:"ret_code_sub"`         // 返回详细码
	RetMessageSub     string `json:"ret_message_sub"`      // 返回详细信息
	ZimId             string `json:"zim_id"`               // ZIM上下文ID
	ZimInitClientData string `json:"zim_init_client_data"` // 客户端协议
}

type ZolozAuthenticationCustomerFtokenQuery struct {
	ErrorResponse
	Uid            string        `json:"uid,omitempty"`              // 支付宝uid，逐步替换为 open_id
	OpenId         string        `json:"open_id"`                    // 支付宝用户open_id
	UidTelPairList []*UidTelPair `json:"uid_tel_pair_list"`          // 用户名信息返回的列表
	AgeCheckResult string        `json:"age_check_result,omitempty"` // 年龄是否在指定范围内，未指定范围则返回空，true/false
	CertNo         string        `json:"cert_no,omitempty"`          // 身份证号码
	CertName       string        `json:"cert_name,omitempty"`        // 证件姓名
	FaceId         string        `json:"face_id,omitempty"`          // 由ISV定义的对自然人唯一编码，举例可以是身份证号码和姓名的MD5值，或者是其他编码方式，要求脱敏、随机且在ISV可以唯一说明一个自然人
}

type UidTelPair struct {
	UserId string `json:"user_id,omitempty"`
	OpenId string `json:"open_id"`
}
