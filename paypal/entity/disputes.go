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
 @Time    : 2024/10/15 -- 12:06
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: disputes.go
*/

package entity

type EscalateDisputeToClaimRes struct {
	EmptyRes
	Response *V1Links `json:"response,omitempty"`
}

type AcceptOffer2ResolveRes struct {
	EmptyRes
	Response *V1Links `json:"response,omitempty"`
}

type ListDisputesRes struct {
	EmptyRes
	Response *DisputeLists `json:"response,omitempty"`
}

type ProvideInfo4DisputeRes struct {
	EmptyRes
	Response *V1Links `json:"response,omitempty"`
}

type ShowDisputeDetailsRes struct {
	EmptyRes
	Response *DisputeDetail `json:"response,omitempty"`
}

type PartiallyUpdateDisputeRes struct {
	EmptyRes
}

type DenyOffer2ResolveRes struct {
	EmptyRes
	Response *V1Links `json:"response,omitempty"`
}

type MakeOffer2ResolveRes struct {
	EmptyRes
	Response *V1Links `json:"response,omitempty"`
}

type AppealDisputeRes struct {
	EmptyRes
	Response *V1Links `json:"response,omitempty"`
}

type ProvideEvidenceRes struct {
	EmptyRes
	Response *V1Links `json:"response,omitempty"`
}

type AckReturnedItemRes struct {
	EmptyRes
	Response *V1Links `json:"response,omitempty"`
}

type NotifyDispute2ThirdPartyRes struct {
	EmptyRes
	Response *V1Links `json:"response,omitempty"`
}

type AcceptClaimRes struct {
	EmptyRes
	Response *V1Links `json:"response,omitempty"`
}

// =========================================================分割=========================================================

type DisputeDetail struct {
	DisputeId              string                  `json:"dispute_id,omitempty"`
	DisputeTransactions    []*DisputeTransaction   `json:"dispute_transactions,omitempty"`
	ExternalReasonCode     string                  `json:"external_reason_code,omitempty"`
	Adjudications          []*Adjudication         `json:"adjudications,omitempty"`
	MoneyMovements         []*MoneyMovement        `json:"money_movements,omitempty"`
	Messages               []*Message              `json:"messages,omitempty"`
	Evidences              []*EvidenceInfo         `json:"evidences,omitempty"`
	SupportingInfo         []*SupportingInfo       `json:"supporting_info,omitempty"`
	Links                  []*Link                 `json:"links,omitempty"`
	CreateTime             string                  `json:"create_time,omitempty"`
	UpdateTime             string                  `json:"update_time,omitempty"`
	Reason                 string                  `json:"reason,omitempty"`
	Status                 string                  `json:"status,omitempty"`
	DisputeStatus          string                  `json:"dispute_status,omitempty"`
	DisputeAmount          *Amount                 `json:"dispute_amount,omitempty"`
	DisputeAsset           *CryptoBalance          `json:"dispute_asset,omitempty"`
	DisputeLifeCycleStage  string                  `json:"dispute_life_cycle_stage,omitempty"`
	DisputeChannel         string                  `json:"dispute_channel,omitempty"`
	BuyerResponseDueDate   string                  `json:"buyer_response_due_date,omitempty"`
	SellerResponseDueDate  string                  `json:"seller_response_due_date,omitempty"`
	FeePolicy              *FeePolicy              `json:"fee_policy,omitempty"`
	DisputeOutcome         *DisputeOutcome         `json:"dispute_outcome,omitempty"`
	Extensions             *Extensions             `json:"extensions,omitempty"`
	Offer                  *Offer                  `json:"offer,omitempty"`
	RefundDetails          *RefundDetail           `json:"refund_details,omitempty"`
	CommunicationDetails   []CommunicationDetail   `json:"communication_details,omitempty"`
	AllowedResponseOptions *AllowedResponseOptions `json:"allowed_response_options,omitempty"`
}

type Adjudication struct {
	Type                  string `json:"type"`
	AdjudicationTime      string `json:"adjudication_time"`
	Reason                string `json:"reason,omitempty"`
	DisputeLifeCycleStage string `json:"dispute_life_cycle_stage,omitempty"`
}

type MoneyMovement struct {
	AffectedParty string         `json:"affected_party,omitempty"`
	Type          string         `json:"type,omitempty"`
	Amount        *Amount        `json:"amount,omitempty"`
	Asset         *CryptoBalance `json:"asset,omitempty"`
	InitiatedTime string         `json:"initiated_time,omitempty"`
	Reason        string         `json:"reason,omitempty"`
}

type Message struct {
	PostedBy   string      `json:"posted_by,omitempty"`
	Content    string      `json:"content,omitempty"`
	Documents  []*Document `json:"documents,omitempty"`
	TimePosted string      `json:"time_posted,omitempty"`
}

type EvidenceInfo struct {
	TrackingInfo []*Tracking `json:"tracking_info,omitempty"`
	RefundIds    []string    `json:"refund_ids,omitempty"`
}

type Tracking struct {
	CarrierName      string `json:"carrier_name,omitempty"`
	CarrierNameOther string `json:"carrier_name_other,omitempty"`
	TrackingUrl      string `json:"tracking_url,omitempty"`
	TrackingNumber   string `json:"tracking_number,omitempty"`
}

type SupportingInfo struct {
	Notes                 string      `json:"notes,omitempty"`
	Documents             []*Document `json:"documents,omitempty"`
	Source                string      `json:"source,omitempty"`
	ProvidedTime          string      `json:"provided_time,omitempty"`
	DisputeLifeCycleStage string      `json:"dispute_life_cycle_stage,omitempty"`
}

type Document struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}

type FeePolicy struct {
}

type DisputeOutcome struct {
	OutcomeCode    string         `json:"outcome_code,omitempty"`
	AmountRefunded *Amount        `json:"amount_refunded,omitempty"`
	AssetRefunded  *CryptoBalance `json:"asset_refunded,omitempty"`
}

type Extensions struct {
	MerchantContacted            bool                    `json:"merchant_contacted,omitempty"`
	BuyerContactedChannel        string                  `json:"buyer_contacted_channel,omitempty"`
	MerchantContactedOutcome     string                  `json:"merchant_contacted_outcome,omitempty"`
	MerchantContactedTime        string                  `json:"merchant_contacted_time,omitempty"`
	MerchantContactedMode        string                  `json:"merchant_contacted_mode,omitempty"`
	BuyerContactedTime           string                  `json:"buyer_contacted_time,omitempty"`
	BillingDisputeProperties     *BillingIssueDetail     `json:"billing_dispute_properties,omitempty"`
	MerchandizeDisputeProperties *MerchandiseIssueDetail `json:"merchandize_dispute_properties,omitempty"`
}

type BillingIssueDetail struct {
	DuplicateTransaction bool `json:"duplicate_transaction,omitempty"`
}

type MerchandiseIssueDetail struct {
	IssueType             string               `json:"issue_type,omitempty"`
	ProductDetail         *ProductDetails      `json:"product_detail,omitempty"`
	ServiceDetail         *ServiceDetails      `json:"service_detail,omitempty"`
	CancellationDetails   *CancellationDetails `json:"cancellation_details,omitempty"`
	ReturnShippingAddress *Address             `json:"return_shipping_address,omitempty"`
}

type CancellationDetails struct {
	CancellationNumber string `json:"cancellation_number,omitempty"`
	Cancelled          bool   `json:"cancelled,omitempty"`
	CancellationMode   string `json:"cancellation_mode,omitempty"`
	CancellationDate   string `json:"cancellation_date,omitempty"`
}

type ProductDetails struct {
	Description          string         `json:"description,omitempty"`
	ProductReceived      string         `json:"product_received,omitempty"`
	SubReasons           []string       `json:"sub_reasons,omitempty"`
	PurchaseUrl          string         `json:"purchase_url,omitempty"`
	ProductReceivedTime  string         `json:"product_received_time,omitempty"`
	ExpectedDeliveryDate string         `json:"expected_delivery_date,omitempty"`
	ReturnDetails        *ReturnDetails `json:"return_details,omitempty"`
}

type ReturnDetails struct {
	Mode                     string `json:"mode,omitempty"`
	Receipt                  bool   `json:"receipt,omitempty"`
	ReturnConfirmationNumber string `json:"return_confirmation_number,omitempty"`
	Returned                 bool   `json:"returned,omitempty"`
	ReturnTime               string `json:"return_time,omitempty"`
}

type ServiceDetails struct {
	Description    string   `json:"description,omitempty"`
	ServiceStarted string   `json:"service_started,omitempty"`
	Note           string   `json:"note,omitempty"`
	SubReasons     []string `json:"sub_reasons,omitempty"`
	PurchaseUrl    string   `json:"purchase_url,omitempty"`
}

type Offer struct {
	History              []*History `json:"history,omitempty"`
	BuyerRequestedAmount *Amount    `json:"buyer_requested_amount,omitempty"`
	SellerOfferedAmount  *Amount    `json:"seller_offered_amount,omitempty"`
	OfferType            string     `json:"offer_type,omitempty"`
}

type History struct {
	Actor                 string  `json:"actor,omitempty"`
	EventType             string  `json:"event_type,omitempty"`
	Notes                 string  `json:"notes,omitempty"`
	OfferTime             string  `json:"offer_time,omitempty"`
	OfferType             string  `json:"offer_type,omitempty"`
	OfferAmount           *Amount `json:"offer_amount,omitempty"`
	DisputeLifeCycleStage string  `json:"dispute_life_cycle_stage,omitempty"`
}

type RefundDetail struct {
	AllowedRefundAmount *Amount `json:"allowed_refund_amount,omitempty"`
}

type CommunicationDetail struct {
	Note       string `json:"note,omitempty"`
	Email      string `json:"email,omitempty"`
	TimePosted string `json:"time_posted,omitempty"`
}

type AllowedResponseOptions struct {
	MakeOffer             *OfferType           `json:"make_offer,omitempty"`
	AcceptClaim           *AcceptClaimType     `json:"accept_claim,omitempty"`
	AcknowledgeReturnItem *AcknowledgementType `json:"acknowledge_return_item,omitempty"`
}

type AcknowledgementType struct {
	AcknowledgementTypes []string `json:"acknowledgement_types,omitempty"`
}

type AcceptClaimType struct {
	AcceptClaimTypes []string `json:"accept_claim_types,omitempty"`
}

type OfferType struct {
	OfferTypes []string `json:"offer_types,omitempty"`
}

type DisputeTransaction struct {
	BuyerTransactionId  string           `json:"buyer_transaction_id,omitempty"`
	SellerTransactionId string           `json:"seller_transaction_id,omitempty"`
	ReferenceId         string           `json:"reference_id,omitempty"`
	TransactionStatus   string           `json:"transaction_status,omitempty"`
	InvoiceNumber       string           `json:"invoice_number,omitempty"`
	Custom              string           `json:"custom,omitempty"`
	Items               []*PurchasedItem `json:"items,omitempty"`
	CreateTime          string           `json:"create_time,omitempty"`
	GrossAmount         *Amount          `json:"gross_amount,omitempty"`
	GrossAsset          *CryptoBalance   `json:"gross_asset,omitempty"`
	Buyer               *Buyer           `json:"buyer,omitempty"`
	Seller              *Seller          `json:"seller,omitempty"`
}

type PurchasedItem struct {
	ItemId               string  `json:"item_id,omitempty"`
	ItemName             string  `json:"item_name,omitempty"`
	ItemDescription      string  `json:"item_description,omitempty"`
	ItemQuantity         string  `json:"item_quantity,omitempty"`
	PartnerTransactionId string  `json:"partner_transaction_id,omitempty"`
	Reason               string  `json:"reason,omitempty"`
	Notes                string  `json:"notes,omitempty"`
	ItemType             string  `json:"item_type,omitempty"`
	DisputeAmount        *Amount `json:"dispute_amount,omitempty"`
}

type Buyer struct {
	Name string `json:"name,omitempty"`
}

type Seller struct {
	MerchantId string `json:"merchant_id,omitempty"`
	Name       string `json:"name,omitempty"`
	Email      string `json:"email,omitempty"`
}

type DisputeLists struct {
	Items []*DisputeDetail `json:"items,omitempty"`
	Links []*Link          `json:"links,omitempty"`
}

type V1Links struct {
	Links []*Link `json:"links,omitempty"`
}
