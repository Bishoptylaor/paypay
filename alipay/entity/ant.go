package entity

type AntMerchantExpandShopCreateResponse struct {
	Response     *AntMerchantExpandShopCreate `json:"ant_merchant_expand_shop_create_response"`
	AlipayCertSn string                       `json:"alipay_cert_sn,omitempty"`
	SignData     string                       `json:"-"`
	Sign         string                       `json:"sign"`
}
type AntMerchantExpandShopModifyResponse struct {
	Response     *AntMerchantExpandShopModify `json:"ant_merchant_expand_shop_modify_response"`
	AlipayCertSn string                       `json:"alipay_cert_sn,omitempty"`
	SignData     string                       `json:"-"`
	Sign         string                       `json:"sign"`
}

type AntMerchantExpandOrderQueryResponse struct {
	Response     *AntMerchantExpandOrderQuery `json:"ant_merchant_expand_order_query_response"`
	AlipayCertSn string                       `json:"alipay_cert_sn,omitempty"`
	SignData     string                       `json:"-"`
	Sign         string                       `json:"sign"`
}

type AntMerchantExpandShopPageQueryResponse struct {
	Response     *AntMerchantExpandShopPageQuery `json:"ant_merchant_expand_shop_page_query_response"`
	AlipayCertSn string                          `json:"alipay_cert_sn,omitempty"`
	SignData     string                          `json:"-"`
	Sign         string                          `json:"sign"`
}

type AntMerchantExpandShopQueryResponse struct {
	Response     *AntMerchantExpandShopQuery `json:"ant_merchant_expand_shop_query_response"`
	AlipayCertSn string                      `json:"alipay_cert_sn,omitempty"`
	SignData     string                      `json:"-"`
	Sign         string                      `json:"sign"`
}

type AntMerchantExpandShopCloseResponse struct {
	Response     *AntMerchantExpandShopClose `json:"ant_merchant_expand_shop_close_response"`
	AlipayCertSn string                      `json:"alipay_cert_sn,omitempty"`
	SignData     string                      `json:"-"`
	Sign         string                      `json:"sign"`
}

type AntMerchantExpandIndirectImageUploadResponse struct {
	Response     *AntMerchantExpandIndirectImageUpload `json:"ant_merchant_expand_indirect_image_upload_response"`
	AlipayCertSn string                                `json:"alipay_cert_sn,omitempty"`
	SignData     string                                `json:"-"`
	Sign         string                                `json:"sign"`
}

type AntMerchantExpandMccQueryResponse struct {
	Response     *AntMerchantExpandMccQuery `json:"ant_merchant_expand_mcc_query_response"`
	AlipayCertSn string                     `json:"alipay_cert_sn,omitempty"`
	SignData     string                     `json:"-"`
	Sign         string                     `json:"sign"`
}

type AntMerchantExpandShopConsultResponse struct {
	Response     *AntMerchantExpandShopConsult `json:"ant_merchant_expand_shop_consult_response"`
	AlipayCertSn string                        `json:"alipay_cert_sn,omitempty"`
	SignData     string                        `json:"-"`
	Sign         string                        `json:"sign"`
}

type AntMerchantExpandShopReceiptaccountSaveResponse struct {
	Response     *AntMerchantExpandShopReceiptaccountSave `json:"ant_merchant_expand_shop_receiptaccount_save_response"`
	AlipayCertSn string                                   `json:"alipay_cert_sn,omitempty"`
	SignData     string                                   `json:"-"`
	Sign         string                                   `json:"sign"`
}

// =========================================================分割=========================================================

type AntMerchantExpandShopCreate struct {
	ErrorResponse
	OrderId string `json:"order_id"`
}
type AntMerchantExpandShopModify struct {
	ErrorResponse
	OrderId string `json:"order_id"`
}

type AntMerchantExpandOrderQuery struct {
	ErrorResponse
	IPRoleID     []string `json:"ip_role_id,omitempty"`
	MerchantName string   `json:"merchant_name"`
	Status       string   `json:"status"`
	ApplyTime    string   `json:"apply_time"`
	ExtInfo      string   `json:"ext_info"`
}

type AntMerchantExpandShopPageQuery struct {
	ErrorResponse
	ShopInfos   []*AntShopInfo `json:"shop_infos"`
	TotalPages  int            `json:"total_pages"`
	AlipayPoiid string         `json:"alipay_poiid"`
}

type AntShopInfo struct {
	ShopId          string           `json:"shop_id"`
	BusinessAddress *BusinessAddress `json:"business_address"`
	ShopCategory    string           `json:"shop_category"`
	NewShopCategory string           `json:"new_shop_category"`
	StoreId         string           `json:"store_id"`
	ShopType        string           `json:"shop_type"`
	ShopName        string           `json:"shop_name"`
	ContactPhone    string           `json:"contact_phone"`
	ContactMobile   string           `json:"contact_mobile"`
	BusinessTime    []*BusinessTime  `json:"business_time"`
	ShopStatus      string           `json:"shop_status"`
	ShopInfoStatus  string           `json:"shop_info_status"`
}

type BusinessAddress struct {
	CityCode     string `json:"city_code"`
	DistrictCode string `json:"district_code"`
	Address      string `json:"address"`
	ProvinceCode string `json:"province_code"`
	Poiid        string `json:"poiid"`
	Longitude    string `json:"longitude"`
	Latitude     string `json:"latitude"`
	Type         string `json:"type"`
}

type BusinessTime struct {
	WeekDay   int    `json:"week_day"`
	OpenTime  string `json:"open_time"`
	CloseTime string `json:"close_time"`
}

type AntMerchantExpandShopQuery struct {
	ErrorResponse
	ShopID          string           `json:"shop_id"`
	BusinessAddress *BusinessAddress `json:"business_address"`
	ShopCategory    string           `json:"shop_category"`
	StoreID         string           `json:"store_id"`
	ShopType        string           `json:"shop_type"`
	IPRoleID        string           `json:"ip_role_id"`
	ShopName        string           `json:"shop_name"`
	ContactPhone    string           `json:"contact_phone"`
	ContactMobile   string           `json:"contact_mobile"`
	CertNo          string           `json:"cert_no"`
	OutDoorImages   []string         `json:"out_door_images,omitempty"`
	Qualifications  []struct {
		IndustryQualificationType  string `json:"industry_qualification_type,omitempty"`
		IndustryQualificationImage string `json:"industry_qualification_image,omitempty"`
	} `json:"qualifications,omitempty"`
	CertType               string `json:"cert_type,omitempty"`
	CertName               string `json:"cert_name,omitempty"`
	CertImage              string `json:"cert_image,omitempty"`
	LegalName              string `json:"legal_name,omitempty"`
	LegalCertNo            string `json:"legal_cert_no,omitempty"`
	LicenseAuthLetterImage string `json:"license_auth_letter_image,omitempty"`
	SettleAlipayLogonID    string `json:"settle_alipay_logon_id,omitempty"`
	ExtInfos               []struct {
		KeyName string `json:"key_name"`
		Value   string `json:"value"`
	} `json:"ext_infos,omitempty"`
	BusinessTime []*BusinessTime `json:"business_time,omitempty"`
	ContactInfos []struct {
		Name     string   `json:"name"`
		Phone    string   `json:"phone,omitempty"`
		Mobile   string   `json:"mobile,omitempty"`
		Email    string   `json:"email,omitempty"`
		Tag      []string `json:"tag"`
		Type     string   `json:"type"`
		IDCardNo string   `json:"id_card_no,omitempty"`
	} `json:"contact_infos,omitempty"`
	Memo            string `json:"memo,omitempty"`
	BrandID         string `json:"brand_id,omitempty"`
	Scene           string `json:"scene,omitempty"`
	NewShopCategory string `json:"new_shop_category"`
}

type AntMerchantExpandShopClose struct {
	ErrorResponse
}

type AntMerchantExpandIndirectImageUpload struct {
	ErrorResponse
	ImageId string `json:"image_id"`
}

type AntMerchantExpandMccQuery struct {
	ErrorResponse
	MccInfoList []*MccInfo `json:"mcc_info_list"`
}

type MccInfo struct {
	MccLevel1           string `json:"mcc_level_1"`
	MccLevel1Name       string `json:"mcc_level_1_name"`
	MccLevel2           string `json:"mcc_level_2"`
	MccLevel2Name       string `json:"mcc_level_2_name"`
	IsSpecial           bool   `json:"is_special"`
	SpecialQualRequired bool   `json:"special_qual_required"`
	MccRequirements     string `json:"mcc_requirements"`
}

type AntMerchantExpandShopConsult struct {
	ErrorResponse
	AccountAudit bool   `json:"account_audit"`
	RiskAudit    bool   `json:"risk_audit"`
	OrderId      string `json:"order_id"`
}

type AntMerchantExpandShopReceiptaccountSave struct {
	ErrorResponse
}
