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
 @Time    : 2024/9/10 -- 13:52
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: invoice.go
*/

package entity

type CreateInvoiceRes struct {
	EmptyRes
	Response *Invoice `json:"response,omitempty"`
}

type ListInvoiceRes struct {
	EmptyRes
	Response *InvoiceList `json:"response,omitempty"`
}

type SendInvoiceRes struct {
	EmptyRes
}

type SendInvoiceReminderRes struct {
	EmptyRes
}

type CancelSentInvoiceRes struct {
	EmptyRes
}

type RecordPaymentForInvoiceRes struct {
	EmptyRes
	Response *RecordPaymentForInvoice `json:"response,omitempty"`
}

type DeleteExternalPaymentRes struct {
	EmptyRes
}

type RecordRefundForInvoiceRes struct {
	EmptyRes
	Response *RecordRefundForInvoice `json:"response,omitempty"`
}

type DeleteExternalRefundRes struct {
	EmptyRes
}

type GenerateInvoiceQRCodeRes struct {
	EmptyRes
	Response *Qrcode `json:"response,omitempty"`
}

type GenerateInvoiceNumberRes struct {
	EmptyRes
	Response *InvoiceNumber `json:"response,omitempty"`
}

type ShowInvoiceDetailRes struct {
	EmptyRes
	Response *Invoice `json:"response,omitempty"`
}

type FullyUpdateInvoiceRes struct {
	EmptyRes
	Response *Invoice `json:"response,omitempty"`
}

type DeleteInvoiceRes struct {
	EmptyRes
}

type SearchInvoiceRes struct {
	EmptyRes
	Response *SearchInvoice `json:"response,omitempty"`
}

type ListInvoiceTemplateRes struct {
	EmptyRes
	Response *InvoiceTemplate `json:"response,omitempty"`
}

type CreateInvoiceTemplateRes struct {
	EmptyRes
	Response *Template `json:"response,omitempty"`
}

type ShowTemplateDetailsRes struct {
	EmptyRes
	Response *Template `json:"response,omitempty"`
}

type FullyUpdateInvoiceTemplateRes struct {
	EmptyRes
	Response *Template `json:"response,omitempty"`
}

type DeleteInvoiceTemplateRes struct {
	EmptyRes
}

// =========================================================分割=========================================================

type InvoiceTemplate struct {
	Addresses []*Address     `json:"addresses"`
	Phones    []*PhoneDetail `json:"phones"`
	Templates []*Template    `json:"templates"`
	Links     []*Link        `json:"links"`
	Emails    string         `json:"emails"`
}

type PhoneDetail struct {
	CountryCode     string `json:"country_code"`
	NationalNumber  string `json:"national_number"`
	ExtensionNumber string `json:"extension_number,omitempty"`
	PhoneType       string `json:"phone_type"`
}

type Template struct {
	Id               string        `json:"id,omitempty"`
	Name             string        `json:"name,omitempty"`
	DefaultTemplate  bool          `json:"default_template,omitempty"`
	StandardTemplate bool          `json:"standard_template,omitempty"`
	Links            []*Link       `json:"links,omitempty"`
	TemplateInfo     *TemplateInfo `json:"template_info,omitempty"`
	Settings         *Settings     `json:"settings,omitempty"`
	UnitOfMeasure    string        `json:"unit_of_measure,omitempty"`
}

type TemplateInfo struct {
	PrimaryRecipients    []*RecipientInfo `json:"primary_recipients,omitempty"`
	AdditionalRecipients []string         `json:"additional_recipients,omitempty"`
	Items                []*InvoiceItem   `json:"items,omitempty"`
	Detail               []*InvoiceDetail `json:"detail,omitempty"`
	Invoicer             *Invoicer        `json:"invoicer,omitempty"`
	Configuration        *Configuration   `json:"configuration,omitempty"`
	Amount               *Amount          `json:"amount,omitempty"`
	DueAmount            *Amount          `json:"due_amount,omitempty"`
}

type Settings struct {
	TemplateItemSettings     []*TemplateSetting `json:"template_item_settings,omitempty"`
	TemplateSubtotalSettings []*TemplateSetting `json:"template_subtotal_settings,omitempty"`
}

type TemplateSetting struct {
	FieldName         string             `json:"field_name,omitempty"`
	DisplayPreference *DisplayPreference `json:"display_preference,omitempty"`
}

type DisplayPreference struct {
	Hidden bool `json:"hidden,omitempty"`
}

type SearchInvoice struct {
	Items      []*Invoice `json:"items"`
	Links      []*Link    `json:"links"`
	TotalItems int        `json:"total_items"`
	TotalPages int        `json:"total_pages"`
}

type InvoiceNumber struct {
	InvoiceNumber string `json:"invoice_number"`
}

type Qrcode struct {
	Base64Image string
}

type RecordRefundForInvoice struct {
	RefundId string `json:"refundId,omitempty"`
}

type RecordPaymentForInvoice struct {
	PaymentId string `json:"payment_id,omitempty"`
}

type InvoiceList struct {
	TotalItems int        `json:"total_items"`
	TotalPages int        `json:"total_pages"`
	Items      []*Invoice `json:"items"`
	Links      []*Link    `json:"links,omitempty"`
}

type Invoice struct {
	Id                   string            `json:"id"`
	ParentId             string            `json:"parent_id,omitempty"`
	PrimaryRecipients    []*RecipientInfo  `json:"primary_recipients,omitempty"`
	AdditionalRecipients []string          `json:"additional_recipients,omitempty"`
	Items                []*InvoiceItem    `json:"items,omitempty"`
	Links                []*Link           `json:"links,omitempty"`
	Status               string            `json:"status"`
	Detail               *InvoiceDetail    `json:"detail"`
	Invoicer             *Invoicer         `json:"invoicer"`
	Configuration        *Configuration    `json:"configuration,omitempty"`
	Amount               *Amount           `json:"amount"`
	DueAmount            *Amount           `json:"due_amount"`
	Gratuity             *Amount           `json:"gratuity,omitempty"`
	Payments             []*InvoicePayment `json:"payments,omitempty"`
	Refunds              []*InvoiceRefund  `json:"refunds,omitempty"`
}

type RecipientInfo struct {
	BillingInfo  *BillingInfo  `json:"billing_info,omitempty"`
	ShippingInfo *ShippingInfo `json:"shipping_info,omitempty"`
}

type BillingInfo struct {
	BusinessName   string   `json:"business_name,omitempty"`
	Name           *Name    `json:"name,omitempty"`
	Address        *Address `json:"address,omitempty"`
	Phones         []*Phone `json:"phones,omitempty"`
	AdditionalInfo string   `json:"additional_info,omitempty"`
	EmailAddress   string   `json:"email_address,omitempty"`
	Language       string   `json:"language,omitempty"`
}

type InvoiceItem struct {
	Id            string    `json:"id,omitempty"`
	Name          string    `json:"name,omitempty"`
	Description   string    `json:"description,omitempty"`
	Quantity      string    `json:"quantity,omitempty"`
	UnitAmount    *Amount   `json:"unit_amount,omitempty"`
	Tax           *Tax      `json:"tax,omitempty"`
	ItemDate      string    `json:"item_date,omitempty"`
	Discount      *Discount `json:"discount,omitempty"`
	UnitOfMeasure string    `json:"unit_of_measure,omitempty"`
}

type Tax struct {
	Name    string  `json:"name"`
	TaxNote string  `json:"tax_note,omitempty"`
	Percent string  `json:"percent"`
	Amount  *Amount `json:"amount,omitempty"`
}

type Discount struct {
	Percent string  `json:"percent,omitempty"`
	Amount  *Amount `json:"amount,omitempty"`
}

type InvoiceDetail struct {
	Reference          string        `json:"reference,omitempty"`
	Note               string        `json:"note,omitempty"`
	TermsAndConditions string        `json:"terms_and_conditions,omitempty"`
	Memo               string        `json:"memo,omitempty"`
	Attachments        []*Attachment `json:"attachments,omitempty"`
	CurrencyCode       string        `json:"currency_code"`
	InvoiceNumber      string        `json:"invoice_number,omitempty"`
	InvoiceDate        string        `json:"invoice_date,omitempty"`
	PaymentTerm        *PaymentTerm  `json:"payment_term,omitempty"`
	MetaData           *MetaData     `json:"meta_data,omitempty"`
}

type Attachment struct {
	Id           string `json:"id,omitempty"`
	ReferenceUrl string `json:"reference_url,omitempty"`
	ContentType  string `json:"content_type,omitempty"`
	Size         string `json:"size,omitempty"`
	CreateTime   string `json:"create_time,omitempty"`
}

type PaymentTerm struct {
	TermType string `json:"term_type,omitempty"`
	DueDate  string `json:"due_date,omitempty"`
}

type MetaData struct {
	CreateBy         string `json:"create_by,omitempty"`
	LastUpdateBy     string `json:"last_update_by,omitempty"`
	CreateTime       string `json:"create_time,omitempty"`
	LastUpdateTime   string `json:"last_update_time,omitempty"`
	CancelledBy      string `json:"cancelled_by,omitempty"`
	LastSentBy       string `json:"last_sent_by,omitempty"`
	RecipientViewUrl string `json:"recipient_view_url,omitempty"`
	InvoicerViewUrl  string `json:"invoicer_view_url,omitempty"`
	CancelTime       string `json:"cancel_time,omitempty"`
	FirstSentTime    string `json:"first_sent_time,omitempty"`
	LastSentTime     string `json:"last_sent_time,omitempty"`
	CreatedByFlow    string `json:"created_by_flow,omitempty"`
}

type Invoicer struct {
	BusinessName    string   `json:"business_name,omitempty"`
	Name            *Name    `json:"name,omitempty"`
	Address         *Address `json:"address,omitempty"`
	Phones          []*Phone `json:"phones,omitempty"`
	Website         string   `json:"website,omitempty"`
	TaxId           string   `json:"tax_id,omitempty"`
	AdditionalNotes string   `json:"additional_notes,omitempty"`
	LogoUrl         string   `json:"logo_url,omitempty"`
	EmailAddress    string   `json:"email_address,omitempty"`
}

type Configuration struct {
	TaxCalculatedAfterDiscount bool            `json:"tax_calculated_after_discount,omitempty"`
	TaxInclusive               bool            `json:"tax_inclusive,omitempty"`
	AllowTip                   bool            `json:"allow_tip,omitempty"`
	PartialPayment             *PartialPayment `json:"partial_payment,omitempty"`
	TemplateId                 string          `json:"template_id,omitempty"`
}

type PartialPayment struct {
	AllowPartialPayment bool    `json:"allow_partial_payment,omitempty"`
	MinimumAmountDue    *Amount `json:"minimum_amount_due,omitempty"`
}

type InvoicePayment struct {
	Transactions []PaymentTransaction `json:"transactions,omitempty"`
	PaidAmount   *Amount              `json:"paid_amount,omitempty"`
}

type PaymentTransaction struct {
	PaymentId    string        `json:"payment_id,omitempty"`
	Note         string        `json:"note,omitempty"`
	Type         string        `json:"type,omitempty"`
	PaymentDate  string        `json:"payment_date,omitempty"`
	Method       string        `json:"method"`
	Amount       *Amount       `json:"amount,omitempty"`
	ShippingInfo *ShippingInfo `json:"shipping_info,omitempty"`
}

type ShippingInfo struct {
	BusinessName string   `json:"business_name,omitempty"`
	Name         *Name    `json:"name,omitempty"`
	Address      *Address `json:"address,omitempty"`
}

type InvoiceRefund struct {
	Transactions []*RefundTransaction `json:"transactions,omitempty"`
	RefundAmount *Amount              `json:"refund_amount,omitempty"`
}

type RefundTransaction struct {
	RefundId   string  `json:"refund_id,omitempty"`
	Type       string  `json:"type,omitempty"`
	RefundDate string  `json:"refund_date,omitempty"`
	Amount     *Amount `json:"amount,omitempty"`
	Method     string  `json:"method"`
}
