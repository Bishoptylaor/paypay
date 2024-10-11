package entity

var (
	// AppAuthTokenInBizContent 不需要处理AppAuthToken的方法
	AppAuthTokenInBizContent = map[string]bool{
		"alipay.open.auth.token.app.query": true,
	}
)

type UserPhone struct {
	ErrorResponse
	Mobile string `json:"mobile,omitempty"`
}

type ErrorResponse struct {
	Code    string `json:"code"`
	Msg     string `json:"msg"`
	SubCode string `json:"sub_code,omitempty"`
	SubMsg  string `json:"sub_msg,omitempty"`
}

// =========================================================分割=========================================================

type OpenAuthTokenAppResponse struct {
	Response     *AuthTokenApp `json:"alipay_open_auth_token_app_response"`
	AlipayCertSn string        `json:"alipay_cert_sn,omitempty"`
	SignData     string        `json:"-"`
	Sign         string        `json:"sign"`
}

type AuthTokenApp struct {
	ErrorResponse
	UserId          string   `json:"user_id,omitempty"`
	AuthAppId       string   `json:"auth_app_id,omitempty"`
	AppAuthToken    string   `json:"app_auth_token,omitempty"`
	ExpiresIn       int      `json:"expires_in,omitempty"`
	AppRefreshToken string   `json:"app_refresh_token,omitempty"`
	ReExpiresIn     int      `json:"re_expires_in,omitempty"`
	Tokens          []*Token `json:"tokens,omitempty"`
}

// =========================================================分割=========================================================

type OpenAuthTokenAppInviteCreateResponse struct {
	Response     *OpenAuthTokenAppInviteCreate `json:"alipay_open_auth_appauth_invite_create_response"`
	AlipayCertSn string                        `json:"alipay_cert_sn,omitempty"`
	SignData     string                        `json:"-"`
	Sign         string                        `json:"sign"`
}

type OpenAuthTokenAppInviteCreate struct {
	ErrorResponse
	TaskPageUrl string `json:"task_page_url,omitempty"`
}

// =========================================================分割=========================================================

type OpenAuthTokenAppQueryResponse struct {
	Response     *AuthTokenAppQuery `json:"alipay_open_auth_token_app_query_response"`
	AlipayCertSn string             `json:"alipay_cert_sn,omitempty"`
	SignData     string             `json:"-"`
	Sign         string             `json:"sign"`
}

type AuthTokenAppQuery struct {
	ErrorResponse
	UserId      string   `json:"user_id"`      //授权商户的user_id
	AuthAppId   string   `json:"auth_app_id"`  //授权商户的appid
	ExpiresIn   int      `json:"expires_in"`   //应用授权令牌失效时间，单位到秒
	AuthMethods []string `json:"auth_methods"` //当前app_auth_token的授权接口列表
	AuthStart   string   `json:"auth_start"`   //授权生效时间
	AuthEnd     string   `json:"auth_end"`     //授权失效时间
	Status      string   `json:"status"`       //valid：有效状态；invalid：无效状态
}

// =========================================================分割=========================================================

type UserInfoAuthResponse struct {
	Response     *ErrorResponse `json:"alipay_user_info_auth_response"`
	AlipayCertSn string         `json:"alipay_cert_sn,omitempty"`
	SignData     string         `json:"-"`
	Sign         string         `json:"sign"`
}

// =========================================================分割=========================================================

type MonitorHeartbeatSynResponse struct {
	Response     *MonitorHeartbeatSynRes `json:"monitor_heartbeat_syn_response"`
	AlipayCertSn string                  `json:"alipay_cert_sn,omitempty"`
	SignData     string                  `json:"-"`
	Sign         string                  `json:"sign"`
}

type MonitorHeartbeatSynRes struct {
	ErrorResponse
	Pid string `json:"pid"`
}

// =========================================================分割=========================================================

type OpenAppQrcodeCreateRes struct {
	Response     *OpenAppQrcodeCreate `json:"alipay_open_app_qrcode_create_response"`
	AlipayCertSn string               `json:"alipay_cert_sn,omitempty"`
	SignData     string               `json:"-"`
	Sign         string               `json:"sign"`
}

type OpenAppQrcodeCreate struct {
	ErrorResponse
	QrCodeUrl string `json:"qr_code_url"`
}

// =========================================================分割=========================================================

type MerchantItemFileUploadRes struct {
	Response     *MerchantItemFileUpload `json:"alipay_merchant_item_file_upload_response"`
	AlipayCertSn string                  `json:"alipay_cert_sn,omitempty"`
	SignData     string                  `json:"-"`
	Sign         string                  `json:"sign"`
}

type MerchantItemFileUpload struct {
	ErrorResponse
	MaterialId  string `json:"material_id"`  // 文件在商品中心的素材标识（素材ID长期有效）
	MaterialKey string `json:"material_key"` // 文件在商品中心的素材标示，创建/更新商品时使用
}

// =========================================================分割=========================================================

type DataDataserviceAdDataQueryRes struct {
	Response     *DataDataserviceAdDataQuery `json:"alipay_data_dataservice_ad_data_query_response"`
	AlipayCertSn string                      `json:"alipay_cert_sn,omitempty"`
	SignData     string                      `json:"-"`
	Sign         string                      `json:"sign"`
}

type DataDataserviceAdDataQuery struct {
	ErrorResponse
	DataList []*DataDetail `json:"data_list,omitempty"`
}

// =========================================================分割=========================================================

type Token struct {
	AuthAppId       string `json:"auth_app_id,omitempty"`
	AppAuthToken    string `json:"app_auth_token,omitempty"`
	ExpiresIn       int    `json:"expires_in,omitempty"`
	AppRefreshToken string `json:"app_refresh_token,omitempty"`
	ReExpiresIn     int    `json:"re_expires_in,omitempty"`
	UserId          string `json:"user_id,omitempty"`
}

type DataDetail struct {
	OuterId            string                  `json:"outer_id,omitempty"`
	Impression         int64                   `json:"impression,omitempty"`
	Click              int64                   `json:"click,omitempty"`
	Cost               int64                   `json:"cost,omitempty"`
	ConversionDataList []*ConversionDataDetail `json:"conversion_data_list,omitempty"`
	BizDate            string                  `json:"biz_date,omitempty"`
}

type ConversionDataDetail struct {
	ConversionId     string `json:"conversion_id,omitempty"`
	ConversionResult string `json:"conversion_result,omitempty"`
}
