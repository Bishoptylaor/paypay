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
 @Time    : 2024/10/8 -- 14:40
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: payouts.go
*/

package entity

type CreateBatchPayoutRes struct {
	EmptyRes
	Response *BatchPayout `json:"response,omitempty"`
}

type ShowPayoutBatchDetailRes struct {
	EmptyRes
	Response *PayoutBatchDetail `json:"response,omitempty"`
}

type ShowPayoutItemDetailRes struct {
	EmptyRes
	Response *PayoutItemDetail `json:"response,omitempty"`
}

type CancelUnclaimedPayoutItemRes struct {
	EmptyRes
	Response *PayoutItemDetail `json:"response,omitempty"`
}

// =========================================================分割=========================================================

type BatchPayout struct {
	Links       []Link       `json:"links,omitempty"`
	BatchHeader *BatchHeader `json:"batch_header,omitempty"`
}

type BatchHeader struct {
	PayoutBatchId     string             `json:"payout_batch_id"`
	TimeCreated       string             `json:"time_created,omitempty"`
	BatchStatus       string             `json:"batch_status"`
	SenderBatchHeader *SenderBatchHeader `json:"sender_batch_header,omitempty"`
}

type SenderBatchHeader struct {
	SenderBatchId string `json:"sender_batch_id,omitempty"`
	EmailSubject  string `json:"email_subject,omitempty"`
	EmailMessage  string `json:"email_message,omitempty"`
	RecipientType string `json:"recipient_type,omitempty"`
}

type PayoutBatchDetail struct {
	TotalItems  int                 `json:"total_items"`
	TotalPages  int                 `json:"total_pages"`
	Items       []*PayoutItemDetail `json:"items,omitempty"`
	Links       []Link              `json:"links,omitempty"`
	BatchHeader *BatchHeader        `json:"batch_header,omitempty"`
}

type PayoutItemDetail struct {
	PayoutItemId       string              `json:"payout_item_id"`
	TransactionId      string              `json:"transaction_id"`
	ActivityId         string              `json:"activity_id,omitempty"`
	PayoutBatchId      string              `json:"payout_batch_id"`
	TimeProcessed      string              `json:"time_processed"`
	Links              []*Link             `json:"links,omitempty"`
	TransactionStatus  string              `json:"transaction_status"` // SUCCESS、FAILED、PENDING、UNCLAIMED、RETURNED、ONHOLD、BLOCKED、REFUNDED、REVERSED
	PayoutItemFee      *V1Amount           `json:"payout_item_fee"`
	PayoutItem         *PayoutItem         `json:"payout_item"`
	CurrencyConversion *CurrencyConversion `json:"currency_conversion,omitempty"`
	Errors             *Errors             `json:"errors,omitempty"`
	SenderBatchId      string              `json:"sender_batch_id,omitempty"`
}

type V1Amount struct {
	Currency string `json:"currency"` // The three-character ISO-4217 currency code.
	Value    string `json:"value"`
}

type PayoutItem struct {
	RecipientType string    `json:"recipient_type"`
	Amount        *V1Amount `json:"amount"`
	Note          string    `json:"note"`
	Receiver      string    `json:"receiver"`
	SenderItemId  string    `json:"sender_item_id"`
}

type CurrencyConversion struct {
	ExchangeRate string    `json:"exchange_rate"`
	FromAmount   *V1Amount `json:"from_amount"`
	ToAmount     *V1Amount `json:"to_amount"`
}

type Errors struct {
	Name            string `json:"name"`
	Message         string `json:"message"`
	InformationLink string `json:"information_link"`
}
