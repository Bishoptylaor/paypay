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
 @Time    : 2024/10/10 -- 17:46
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: transcation.go
*/

package entity

type ListTranscationsRes struct {
	EmptyRes
	Response *TranscationInfos `json:"response,omitempty"`
}

type ListAllBalancesRes struct {
	EmptyRes
	Response *AllBalances `json:"response,omitempty"`
}

// =========================================================分割=========================================================

type AllBalances struct {
	Balances        []*Balance       `json:"balances,omitempty"`
	CryptoBalances  []*CryptoBalance `json:"crypto_balances,omitempty"`
	AccountId       string           `json:"account_id,omitempty"`
	AsOfTime        string           `json:"as_of_time,omitempty"`
	LastRefreshTime string           `json:"last_refresh_time,omitempty"`
}

type CryptoBalance struct {
	AssetSymbol string `json:"asset_symbol,omitempty"`
	Quantity    string `json:"quantity,omitempty"`
}

type Balance struct {
	Primary          bool      `json:"primary,omitempty"`
	Currency         string    `json:"currency"`
	TotalBalance     *V1Amount `json:"total_balance"`
	AvailableBalance *V1Amount `json:"available_balance,omitempty"`
	WithheldBalance  *V1Amount `json:"withheld_balance,omitempty"`
}

type TranscationInfos struct {
	TranscationDetails    []*TranscationDetail `json:"transcation_details,omitempty"`
	AccountNumber         string               `json:"account_number,omitempty"`
	Page                  int                  `json:"page,omitempty"`
	TotalItems            int                  `json:"total_items,omitempty"`
	TotalPages            int                  `json:"total_pages,omitempty"`
	Links                 []Link               `json:"links,omitempty"`
	StartDate             string               `json:"start_date,omitempty"`
	EndDate               string               `json:"end_date,omitempty"`
	LastRefreshedDatetime string               `json:"last_refreshed_datetime,omitempty"`
}

type TranscationDetail struct {
	TransactionInfo *Transcation    `json:"transaction_info,omitempty"`
	PayerInfo       *PayerInfo      `json:"payer_info,omitempty"`
	ShippingInfo    *V1ShippingInfo `json:"shipping_info,omitempty"`
	CartInfo        *CartInfo       `json:"cart_info,omitempty"`
	StoreInfo       *StoreInfo      `json:"store_info,omitempty"`
	AuctionInfo     *AuctionInfo    `json:"auction_info,omitempty"`
	IncentiveInfo   *IncentiveInfo  `json:"incentive_info,omitempty"`
}

type IncentiveInfo struct {
	IncentiveDetails []*IncentiveDetail `json:"incentive_details,omitempty"`
}

type IncentiveDetail struct {
	IncentiveType        string    `json:"incentive_type,omitempty"`
	IncentiveCode        string    `json:"incentive_code,omitempty"`
	IncentiveProgramCode string    `json:"incentive_program_code,omitempty"`
	IncentiveAmount      *V1Amount `json:"incentive_amount,omitempty"`
}

type AuctionInfo struct {
	AuctionSite        string `json:"auction_site,omitempty"`
	AuctionItemSite    string `json:"auction_item_site,omitempty"`
	AuctionBuyerId     string `json:"auction_buyer_id,omitempty"`
	AuctionClosingDate string `json:"auction_closing_date,omitempty"`
}

type StoreInfo struct {
	StoreId    string `json:"store_id,omitempty"`
	TerminalId string `json:"terminal_id,omitempty"`
}

type CartInfo struct {
	ItemDetails     []*CartDetail `json:"item_details,omitempty"`
	TaxInclusive    bool          `json:"tax_inclusive,omitempty"`
	PaypalInvoiceId string        `json:"paypal_invoice_id,omitempty"`
}

type CartDetail struct {
	ItemCode            string            `json:"item_code,omitempty"`
	ItemName            string            `json:"item_name,omitempty"`
	ItemDescription     string            `json:"item_description,omitempty"`
	ItemOptions         string            `json:"item_options,omitempty"`
	ItemQuantity        string            `json:"item_quantity,omitempty"`
	TaxAmounts          []*V1Amount       `json:"tax_amounts,omitempty"`
	InvoiceNumber       string            `json:"invoice_number,omitempty"`
	CheckoutOptions     []*CheckoutOption `json:"checkout_options,omitempty"`
	ItemUnitPrice       *V1Amount         `json:"item_unit_price,omitempty"`
	ItemAmount          *V1Amount         `json:"item_amount,omitempty"`
	DiscountAmount      *V1Amount         `json:"discount_amount,omitempty"`
	AdjustmentAmount    *V1Amount         `json:"adjustment_amount,omitempty"`
	GiftWrapAmount      *V1Amount         `json:"gift_wrap_amount,omitempty"`
	TaxPercentage       string            `json:"tax_percentage,omitempty"`
	BasicShippingAmount *V1Amount         `json:"basic_shipping_amount,omitempty"`
	ExtraShippingAmount *V1Amount         `json:"extra_shipping_amount,omitempty"`
	HandlingAmount      *V1Amount         `json:"handling_amount,omitempty"`
	InsuranceAmount     *V1Amount         `json:"insurance_amount,omitempty"`
	TotalItemAmount     *V1Amount         `json:"total_item_amount,omitempty"`
}

type CheckoutOption struct {
	CheckoutOptionName  string `json:"checkout_option_name,omitempty"`
	CheckoutOptionValue string `json:"checkout_option_value,omitempty"`
}

type Transcation struct {
	PaypalAccountId           string    `json:"paypal_account_id,omitempty"`
	TranscationId             string    `json:"transcation_id,omitempty"`
	PaypalReferenceId         string    `json:"paypal_reference_id,omitempty"`
	PaypalReferenceIdType     string    `json:"paypal_reference_id_type,omitempty"`
	TranscationEventCode      string    `json:"transcation_event_code,omitempty"`
	TranscationStatus         string    `json:"transcation_status,omitempty"`
	TranscationSubject        string    `json:"transcation_subject,omitempty"`
	TranscationNote           string    `json:"transcation_note,omitempty"`
	PaymentTrackingId         string    `json:"payment_tracking_id,omitempty"`
	BankReferenceId           string    `json:"bank_reference_id,omitempty"`
	InvoiceId                 string    `json:"invoice_id,omitempty"`
	CustomeField              string    `json:"custome_field,omitempty"`
	ProtectionEligibility     string    `json:"protection_eligibility,omitempty"`
	CreditTerm                string    `json:"credit_term,omitempty"`
	PaymentMethodType         string    `json:"payment_method_type,omitempty"`
	InstrumentType            string    `json:"instrument_type,omitempty"`
	InstrumentSubType         string    `json:"instrument_sub_type,omitempty"`
	TranscationInitiationDate string    `json:"transaction_initiation_date,omitempty"`
	TranscationUpdatedDate    string    `json:"transcation_updated_date,omitempty"`
	TranscationAmount         *V1Amount `json:"transcation_amount,omitempty"`
	FeeAmount                 *V1Amount `json:"fee_amount,omitempty"`
	DiscountAmount            *V1Amount `json:"discount_amount,omitempty"`
	InsuranceAmount           *V1Amount `json:"insurance_amount,omitempty"`
	SalesTaxAmount            *V1Amount `json:"sales_tax_amount,omitempty"`
	ShippingAmount            *V1Amount `json:"shipping_amount,omitempty"`
	ShippingDiscountAmount    *V1Amount `json:"shipping_discount_amount,omitempty"`
	ShippingTaxAmount         *V1Amount `json:"shipping_tax_amount,omitempty"`
	OtherAmount               *V1Amount `json:"other_amount,omitempty"`
	TipAmount                 *V1Amount `json:"tip_amount,omitempty"`
	EndingAmount              *V1Amount `json:"ending_amount,omitempty"`
	AvailableBalance          *V1Amount `json:"available_balance,omitempty"`
	CredisTranscationalFee    *V1Amount `json:"credit_transcational_fee,omitempty"`
	CreditPromotionalFee      *V1Amount `json:"credit_promotional_fee,omitempty"`
	AnnualPercentageRate      string    `json:"annual_percentage_rate,omitempty"`
}

type PayerInfo struct {
	AccountId     string       `json:"account_id,omitempty"`
	AddressStatus string       `json:"address_status,omitempty"`
	PayerStatus   string       `json:"payer_status,omitempty"`
	EmailAddress  string       `json:"email_address,omitempty"`
	PhoneNumber   *PhoneNumber `json:"phone_number,omitempty"`
	PayerName     *Name        `json:"payer_name,omitempty"`
	CountryCode   string       `json:"country_code,omitempty"`
	Adderess      *V1Address   `json:"adderess,omitempty"`
}

type V1Address struct {
	Line1       string `json:"line1"`
	Line2       string `json:"line2,omitempty"`
	City        string `json:"city"`
	State       string `json:"state,omitempty"`
	CountryCode string `json:"country_code"`
	PostalCode  string `json:"postal_code,omitempty"`
}

type V1ShippingInfo struct {
	Name                     string     `json:"name,omitempty"`
	Method                   string     `json:"method,omitempty"`
	Address                  *V1Address `json:"address,omitempty"`
	SecondaryShippingAddress *V1Address `json:"secondary_shipping_address,omitempty"`
}
