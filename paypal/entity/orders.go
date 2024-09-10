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
 @Time    : 2024/9/5 -- 14:49
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: order.go
*/

package entity

type CreateOrderRes struct {
	EmptyRes
	Response *OrderDetail `json:"response,omitempty"`
}

type ShowOrderDetailsRes struct {
	EmptyRes
	Response *OrderDetail `json:"response,omitempty"`
}

type UpdateOrderRes struct {
	EmptyRes
	// Response *OrderDetail `json:"response,omitempty"`
}

type ConfirmOrderRes struct {
	EmptyRes
	Response *OrderDetail `json:"response,omitempty"`
}

type AuthorizeOrderRes struct {
	EmptyRes
	Response *OrderDetail `json:"response,omitempty"`
}

type CaptureOrderRes struct {
	EmptyRes
	Response *OrderDetail `json:"response,omitempty"`
}

type AddTrackerForOrderRes struct {
	EmptyRes
	Response *OrderDetail `json:"response,omitempty"`
}

type TrackersOfOrderRes struct {
	EmptyRes
	// Response *OrderDetail `json:"response,omitempty"`
}

// =========================================================分割=========================================================

type OrderDetail struct {
	CreateTime            string           `json:"create_time,omitempty"`
	UpdateTime            string           `json:"update_time,omitempty"`
	Id                    string           `json:"id,omitempty"`
	ProcessingInstruction string           `json:"processing_instruction,omitempty"`
	PurchaseUnits         []*PurchaseUnit  `json:"purchase_units,omitempty"`
	Links                 []*Link          `json:"links,omitempty"`
	PaymentSource         []*PaymentSource `json:"payment_source,omitempty"`
	Intent                string           `json:"intent,omitempty"`
	Payer                 *Payer           `json:"payer,omitempty"`
	Status                string           `json:"status,omitempty"`
}

type PurchaseUnit struct {
	ReferenceId        string              `json:"reference_id,omitempty"`
	Description        string              `json:"description,omitempty"`
	CustomId           string              `json:"custom_id,omitempty"`
	InvoiceId          string              `json:"invoice_id,omitempty"`
	Id                 string              `json:"id,omitempty"`
	SoftDescriptor     string              `json:"soft_descriptor,omitempty"`
	Items              []*Item             `json:"items,omitempty"`
	Amount             *Amount             `json:"amount,omitempty"`
	Payee              *Payee              `json:"payee,omitempty"`
	PaymentInstruction *PaymentInstruction `json:"payment_instruction,omitempty"`
	Shipping           *Shipping           `json:"shipping,omitempty"`
	SupplementaryData  *SupplementaryData  `json:"supplementary_data,omitempty"`
	Payments           *Payments           `json:"payments,omitempty"`
}

type Item struct {
	Name        string  `json:"name"`
	Quantity    string  `json:"quantity"`
	Description string  `json:"description,omitempty"`
	Sku         string  `json:"sku,omitempty"`
	Url         string  `json:"url,omitempty"`
	Category    string  `json:"category,omitempty"`
	ImageUrl    string  `json:"image_url,omitempty"`
	UnitAmount  *Amount `json:"unit_amount,omitempty"`
	Tax         *Amount `json:"tax,omitempty"`
	Upc         *Upc    `json:"upc,omitempty"`
}

type PaymentInstruction struct {
	PlatformFees            []*PlatformFee `json:"platform_fees,omitempty"`
	PayeePricingTierId      string         `json:"payee_pricing_tier_id,omitempty"`
	PayeeReceivableFxRateId string         `json:"payee_receivable_fx_rate_id,omitempty"`
	DisbursementMode        string         `json:"disbursement_mode,omitempty"`
}

type PlatformFee struct {
	Amount *Amount `json:"amount"`
	Payee  *Payee  `json:"payee,omitempty"`
}

type Shipping struct {
	Type        string       `json:"type,omitempty"`
	Options     *Option      `json:"options,omitempty"`
	Name        *Name        `json:"name,omitempty"`
	PhoneNumber *PhoneNumber `json:"phone_number,omitempty"`
	Address     *Address     `json:"address,omitempty"`
	Trackers    []*Tracker   `json:"trackers,omitempty"`
}

type Option struct {
	Id       string  `json:"id"`
	Label    string  `json:"label"`
	Selected bool    `json:"selected"`
	Type     string  `json:"type,omitempty"`
	Amount   *Amount `json:"amount,omitempty"`
}

type Tracker struct {
	Id         string         `json:"id,omitempty"`
	Status     string         `json:"status,omitempty"`
	Items      []*TrackerItem `json:"items,omitempty"`
	Links      []*Link        `json:"links,omitempty"`
	CreateTime string         `json:"create_time,omitempty"`
	UpdateTime string         `json:"update_time,omitempty"`
}

type TrackerItem struct {
	Name     string `json:"name,omitempty"`
	Quantity string `json:"quantity,omitempty"`
	Sku      string `json:"sku,omitempty"`
	Url      string `json:"url,omitempty"`
	ImageUrl string `json:"image_url,omitempty"`
	Upc      *Upc   `json:"upc,omitempty"`
}

type Upc struct {
	Type string `json:"type"`
	Code string `json:"code"`
}

type SupplementaryData struct {
	Card *Card `json:"card,omitempty"`
}

type Payments struct {
	Authorizations []*Authorizations `json:"authorizations,omitempty"`
	Captures       []*Captures       `json:"captures,omitempty"`
	Refunds        []*Refunds        `json:"refunds,omitempty"`
}

type Authorizations struct {
	Status                      string                       `json:"status,omitempty"`
	StatusDetails               *StatusDetails               `json:"status_details,omitempty"`
	Id                          string                       `json:"id,omitempty"`
	InvoiceId                   string                       `json:"invoice_id,omitempty"`
	CustomId                    string                       `json:"custom_id,omitempty"`
	Links                       []*Link                      `json:"links,omitempty"`
	Amount                      *Amount                      `json:"amount,omitempty"`
	NetworkTransactionReference *NetworkTransactionReference `json:"network_transaction_reference,omitempty"`
	SellerProtection            *SellerProtection            `json:"seller_protection,omitempty"`
	ExpirationTime              string                       `json:"expiration_time,omitempty"`
	CreateTime                  string                       `json:"create_time,omitempty"`
	UpdateTime                  string                       `json:"update_time,omitempty"`
}

type StatusDetails struct {
	Reason string `json:"reason,omitempty"`
}

type Captures struct {
	Status                      string                       `json:"status,omitempty"`
	StatusDetails               *StatusDetails               `json:"status_details"`
	Id                          string                       `json:"id,omitempty"`
	InvoiceId                   string                       `json:"invoice_id,omitempty"`
	CustomId                    string                       `json:"custom_id,omitempty"`
	FinalCapture                bool                         `json:"final_capture,omitempty"`
	DisbursementMode            string                       `json:"disbursement_mode,omitempty"`
	Links                       []*Link                      `json:"links,omitempty"`
	Amount                      *Amount                      `json:"amount,omitempty"`
	NetworkTransactionReference *NetworkTransactionReference `json:"network_transaction_reference,omitempty"`
	SellerProtection            *SellerProtection            `json:"seller_protection,omitempty"`
	SellerPayableBreakdown      *SellerPayableBreakdown      `json:"seller_payable_breakdown,omitempty"`
	ProcessorResponse           *ProcessorResponse           `json:"processor_response,omitempty"`
	CreateTime                  string                       `json:"create_time,omitempty"`
	UpdateTime                  string                       `json:"update_time,omitempty"`
}

type ProcessorResponse struct {
	AvsCode           string `json:"avs_code,omitempty"`
	CvvCode           string `json:"cvv_code,omitempty"`
	ResponseCode      string `json:"response_code,omitempty"`
	PaymentAdviceCode string `json:"payment_advice_code,omitempty"`
}

type SellerProtection struct {
	Status            string   `json:"status,omitempty"`
	DisputeCategories []string `json:"dispute_categories,omitempty"`
}

type NetworkTransactionReference struct {
	Id                      string `json:"id"`
	Date                    string `json:"date"`
	AcquirerReferenceNumber string `json:"acquirer_reference_number,omitempty"`
	Network                 string `json:"network,omitempty"`
}

type Refunds struct {
	Status        string `json:"status,omitempty"`
	StatusDetails struct {
		Reason string `json:"reason,omitempty"`
	} `json:"status_details"`
	Id                      string                  `json:"id,omitempty"`
	InvoiceId               string                  `json:"invoice_id,omitempty"`
	CustomId                string                  `json:"custom_id,omitempty"`
	AcquirerReferenceNumber string                  `json:"acquirer_reference_number,omitempty"`
	NoteToPayer             string                  `json:"note_to_payer,omitempty"`
	SellerPayableBreakdown  *SellerPayableBreakdown `json:"seller_payable_breakdown,omitempty"`
	Links                   []*Link                 `json:"links,omitempty"`
	Amount                  *Amount                 `json:"amount,omitempty"`
	Payer                   *Payer                  `json:"payer,omitempty"`
	CreateTime              string                  `json:"create_time,omitempty"`
	UpdateTime              string                  `json:"update_time,omitempty"`
}

type SellerPayableBreakdown struct {
	PlatformFees []struct {
		Amount *Amount `json:"amount"`
		Payee  *Payee  `json:"payee,omitempty"`
	} `json:"platform_fees,omitempty"`
	NetAmountBreakdown []struct {
		PayableAmount   *Amount `json:"payable_amount"`
		ConvertedAmount *Amount `json:"converted_amount"`
		ExchangeRate    *Rate   `json:"exchange_rate"`
	} `json:"net_amount_breakdown,omitempty"`
	GrossAmount                   *Amount `json:"gross_amount,omitempty"`
	PaypalFee                     *Amount `json:"paypal_fee,omitempty"`
	PaypalFeeInReceivableCurrency *Amount `json:"paypal_fee_in_receivable_currency,omitempty"`
	NetAmount                     *Amount `json:"net_amount,omitempty"`
	NetFeeInReceivableCurrency    *Amount `json:"net_fee_in_receivable_currency,omitempty"`
	TotalRefundedAmount           *Amount `json:"total_refunded_amount,omitempty"`
}

type Rate struct {
	Value          string `json:"value,omitempty"`
	SourceCurrency string `json:"source_currency,omitempty"`
	TargetCurrency string `json:"target_currency,omitempty"`
}

type Amount struct {
	CurrencyCode string `json:"currency_code"`
	Value        string `json:"value"`
}

type Payee struct {
	EmailAddress string `json:"email_address,omitempty"`
	MerchantId   string `json:"merchant_id,omitempty"`
}

type PaymentSource struct {
	Card       *Card              `json:"card,omitempty"`
	BanContact *Bancontact        `json:"bancontact,omitempty"`
	Blik       *Blik              `json:"blik,omitempty"`
	Eps        *BasePaymentSource `json:"eps,omitempty"`
	Giropay    *BasePaymentSource `json:"giropay,omitempty"`
	Ideal      *BasePaymentSource `json:"ideal,omitempty"`
	Mybank     *BasePaymentSource `json:"mybank,omitempty"`
	P24        *P24               `json:"p24,omitempty"`
	Sofort     *BasePaymentSource `json:"sofort,omitempty"`
	Trustly    *BasePaymentSource `json:"trustly,omitempty"`
	Venmo      *Venmo             `json:"venmo,omitempty"`
	Paypal     *Paypal            `json:"paypal,omitempty"`
	ApplePay   *ApplePay          `json:"apple_pay,omitempty"`
	GooglePay  *GooglePay         `json:"google_pay,omitempty"`
}

type Card struct {
	Name                 string                `json:"name,omitempty"`
	LastDigits           string                `json:"last_digits,omitempty"`
	AvailableNetworks    []string              `json:"available_networks,omitempty"`
	FromRequest          *FromRequest          `json:"from_request,omitempty"`
	Brand                string                `json:"brand,omitempty"`
	Type                 string                `json:"type,omitempty"`
	AuthenticationResult *AuthenticationResult `json:"authentication_result,omitempty"`
	Attributes           *Attributes           `json:"attributes,omitempty"`
	Expiry               string                `json:"expiry,omitempty"`
	BinDetails           []*BinDetails         `json:"bin_details,omitempty"`
	BillingAddress       *Address              `json:"billing_address,omitempty"`
	CountryCode          string                `json:"country_code,omitempty"`
}

type FromRequest struct {
	LastDigits string `json:"last_digits,omitempty"`
	Expiry     string `json:"expiry,omitempty"`
}

type AuthenticationResult struct {
	LiabilityShift string        `json:"liability_shift"`
	ThreeDSecure   *ThreeDSecure `json:"three_d_secure,omitempty"`
}

type ThreeDSecure struct {
	AuthenticationStatus string `json:"authentication_status"`
	EnrollmentStatus     string `json:"enrollment_status"`
}

type Attributes struct {
	Vault []struct {
		Id       string    `json:"id,omitempty"`
		Status   string    `json:"status,omitempty"`
		Links    []string  `json:"links,omitempty"`
		Customer *Customer `json:"customer,omitempty"`
	}
}

type Customer struct {
	Id                 string `json:"id,omitempty"`
	EmailAddress       string `json:"email_address,omitempty"`
	Phone              *Phone `json:"phone,omitempty"`
	MerchantCustomerId string `json:"merchant_customer_id,omitempty"`
}

type Bancontact struct {
	CardLastDigits string `json:"card_last_digits,omitempty"`
	Name           string `json:"name"`
	CountryCode    string `json:"country_code,omitempty"`
	Bic            string `json:"bic,omitempty"`
	IbanLastChars  string `json:"iban_last_chars,omitempty"`
}

type Blik struct {
	Name        string `json:"name,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
	Email       string `json:"email,omitempty"`
	OneClick    struct {
		ConsumerReference string `json:"consumer_reference,omitempty"`
	} `json:"one_click"`
}

type P24 struct {
	PaymentDescriptor string `json:"payment_descriptor,omitempty"`
	MethodId          string `json:"method_id,omitempty"`
	MethodDescription string `json:"method_description,omitempty"`
	Name              string `json:"name,omitempty"`
	Email             string `json:"email,omitempty"`
	CountryCode       string `json:"country_code,omitempty"`
}

type BasePaymentSource struct {
	Name          string `json:"name,omitempty"`
	CountryCode   string `json:"country_code,omitempty"`
	Bic           string `json:"bic,omitempty"`
	IbanLastChars string `json:"iban_last_chars,omitempty"`
}

type Venmo struct {
	UserName     string       `json:"user_name,omitempty"`
	Attributes   *Attributes  `json:"attributes,omitempty"`
	EmailAddress string       `json:"email_address,omitempty"`
	AccountId    string       `json:"account_id,omitempty"`
	Name         *Name        `json:"name,omitempty"`
	PhoneNumber  *PhoneNumber `json:"phone_number,omitempty"`
	Address      *Address     `json:"address,omitempty"`
}

type Paypal struct {
	AccountStatus string       `json:"account_status,omitempty"`
	PhoneType     string       `json:"phone_type,omitempty"`
	BusinessName  string       `json:"business_name,omitempty"`
	Attributes    *Attributes  `json:"attributes,omitempty"`
	EmailAddress  string       `json:"email_address,omitempty"`
	AccountId     string       `json:"account_id,omitempty"`
	Name          *Name        `json:"name,omitempty"`
	PhoneNumber   *PhoneNumber `json:"phone_number,omitempty"`
	BirthDate     string       `json:"birth_date,omitempty"`
	TaxInfo       *TaxInfo     `json:"tax_info,omitempty"`
	Address       *Address     `json:"address,omitempty"`
}

type ApplePay struct {
	Id           string       `json:"id,omitempty"`
	Token        string       `json:"token,omitempty"`
	Name         string       `json:"name,omitempty"`
	EmailAddress string       `json:"email_address,omitempty"`
	PhoneNumber  *PhoneNumber `json:"phone_number,omitempty"`
	Card         *Card        `json:"card,omitempty"`
	Attributes   *Attributes  `json:"attributes,omitempty"`
}

type GooglePay struct {
	Name         string       `json:"name,omitempty"`
	EmailAddress string       `json:"email_address,omitempty"`
	PhoneNumber  *PhoneNumber `json:"phone_number,omitempty"`
	Card         *Card        `json:"card,omitempty"`
}

type PhoneNumber struct {
	CountryCode    string `json:"country_code,omitempty"`
	NationalNumber string `json:"national_number,omitempty"`
}

type BinDetails struct {
	Bin            string   `json:"bin,omitempty"`
	IssuingBank    string   `json:"issuing_bank,omitempty"`
	Products       []string `json:"products,omitempty"`
	BinCountryCode string   `json:"bin_country_code,omitempty"`
}

type Payer struct {
	EmailAddress string   `json:"email_address,omitempty"`
	PayerId      string   `json:"payer_id,omitempty"`
	Name         *Name    `json:"name,omitempty"`
	Phone        *Phone   `json:"phone,omitempty"`
	BirthDate    string   `json:"birth_date,omitempty"`
	TaxInfo      *TaxInfo `json:"tax_info,omitempty"`
	Address      *Address `json:"address,omitempty"`
}

type Name struct {
	GivenName string `json:"given_name,omitempty"`
	Surname   string `json:"surname,omitempty"`
}

type Phone struct {
	PhoneType   string `json:"phone_type,omitempty"`
	PhoneNumber string `json:"phone_number"`
}

type TaxInfo struct {
	TaxId     string `json:"tax_id,omitempty"`
	TaxIdType string `json:"tax_id_type,omitempty"`
}

type Address struct {
	AddressLine1 string `json:"address_line_1,omitempty"`
	AddressLine2 string `json:"address_line_2,omitempty"`
	AdminArea1   string `json:"admin_area_1,omitempty"`
	AdminArea2   string `json:"admin_area_2,omitempty"`
	PostalCode   string `json:"postal_code,omitempty"`
	CountryCode  string `json:"country_code,omitempty"`
}
