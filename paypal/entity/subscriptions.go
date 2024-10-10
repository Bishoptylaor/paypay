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
 @Time    : 2024/10/8 -- 16:41
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: subscriptions.go
*/

package entity

type CreatePlanRes struct {
	EmptyRes
	Response *BillingDetail `json:"response,omitempty"`
}

type ListPlansRes struct {
	EmptyRes
	Response *PlanDetails `json:"response,omitempty"`
}

type ShowPlanDetailsRes struct {
	EmptyRes
	Response *BillingDetail `json:"response,omitempty"`
}

type UpdatePlanRes struct {
	EmptyRes
}

type ActivePlanRes struct {
	EmptyRes
}

type DeactivePlanRes struct {
	EmptyRes
}

type UpdatePricingRes struct {
	EmptyRes
}

type CreateSubscriptionRes struct {
	EmptyRes
	Response *SubscriptionDetail `json:"response,omitempty"`
}

type ShowSubscriptionDetailsRes struct {
	EmptyRes
	Response *SubscriptionDetail `json:"response,omitempty"`
}

type UpdateSubscriptionRes struct {
	EmptyRes
}

type RevisePlanOrQuantityOfSubsriptionRes struct {
	EmptyRes
	Response *SubscriptionDetail `json:"response,omitempty"`
}

type SuspendSubscriptionRes struct {
	EmptyRes
}

type CancelSubscriptionRes struct {
	EmptyRes
}

type ActivateSubscriptionRes struct {
	EmptyRes
}

type CaptureAuthoriedPaymentOnSubscriptionRes struct {
	EmptyRes
}

type ListTransactions4SubscriptionRes struct {
	EmptyRes
	Response *SubscriptionDetail `json:"response,omitempty"`
}

// =========================================================分割=========================================================

type SubscriptionDetail struct {
	Status           string               `json:"status,omitempty"`
	StatusChangeNote string               `json:"status_change_note,omitempty"`
	StatusUpdateTime string               `json:"status_update_time,omitempty"`
	Id               string               `json:"id,omitempty"`
	PlanId           string               `json:"plan_id,omitempty"`
	Quantity         string               `json:"quantity,omitempty"`
	CustomId         string               `json:"custom_id,omitempty"`
	PlanOverridden   bool                 `json:"plan_overridden,omitempty"`
	Links            []*Link              `json:"links,omitempty"`
	StartTime        string               `json:"start_time,omitempty"`
	ShippingAmount   *V1Amount            `json:"shipping_amount,omitempty"`
	Subscriber       *Subscriber          `json:"subscriber,omitempty"`
	BillingInfo      *SubsctiptionBilling `json:"billing_info,omitempty"`
	CreateTime       string               `json:"create_time,omitempty"`
	UpdateTime       string               `json:"update_time,omitempty"`
	Plan             *BillingDetail       `json:"plan,omitempty"`
}

type Subscriber struct {
	EmailAddress    string         `json:"email_address,omitempty"`
	PayerId         string         `json:"payer_id,omitempty"`
	Name            *Name          `json:"name,omitempty"`
	Phone           *Phone         `json:"phone,omitempty"`
	ShippingAddress *Shipping      `json:"shipping_address,omitempty"`
	PaymentSource   *PaymentSource `json:"payment_source,omitempty"`
}

type SubsctiptionBilling struct {
	CycleExecutions     []CycleExecution   `json:"cycle_executions,omitempty"`
	FailedPaymentsCount int                `json:"failed_payments_count,omitempty"`
	OutstandingBalance  *V1Amount          `json:"outstanding_balance,omitempty"`
	LastPayment         *LastPayment       `json:"last_payment,omitempty"`
	NextBillingTime     string             `json:"next_billing_time,omitempty"`
	FinalBillingTime    string             `json:"final_billing_time,omitempty"`
	LastFailedPayment   *LastFailedPayment `json:"last_failed_payment,omitempty"`
}

type LastFailedPayment struct {
	ReasonCode           string    `json:"reason_code,omitempty"`
	Amount               *V1Amount `json:"amount,omitempty"`
	Time                 string    `json:"time,omitempty"`
	NextPaymentRetryTime string    `json:"next_payment_retry_time,omitempty"`
}

type LastPayment struct {
	Status string    `json:"status,omitempty"`
	Amount *V1Amount `json:"amount,omitempty"`
	Time   string    `json:"time,omitempty"`
}

type CycleExecution struct {
	TenureType                  string `json:"tenure_type,omitempty"`
	Sqeuence                    string `json:"sqeuence,omitempty"`
	CyclesCompleted             int    `json:"cycles_completed,omitempty"`
	CyclesRemaining             int    `json:"cycles_remaining,omitempty"`
	CurrentPricingSchemeVersion string `json:"current_pricing_scheme_version,omitempty"`
	TotalCycles                 int    `json:"total_cycles,omitempty"`
}

type PlanDetails struct {
	Plans      []*BillingDetail `json:"plans,omitempty"`
	TotalItems int64            `json:"total_items,omitempty"`
	TotalPages int64            `json:"total_pages,omitempty"`
	Links      []*Link          `json:"links,omitempty"`
}

type BillingDetail struct {
	ID                 string               `json:"id,omitempty"`
	Name               string               `json:"name,omitempty"`
	ProductId          string               `json:"product_id,omitempty"`
	Description        string               `json:"description,omitempty"`
	Status             string               `json:"status,omitempty"`
	BillingCycles      []*BillingCycle      `json:"billing_cycles,omitempty"`
	QuantitySupported  bool                 `json:"quantity_supported,omitempty"`
	Links              []*Link              `json:"links,omitempty"`
	PaymentPreferences []*PaymentPreference `json:"payment_preferences,omitempty"`
	Taxes              []*V1Tax             `json:"taxes,omitempty"`
	CreateTime         string               `json:"create_time,omitempty"`
	UpdateTime         string               `json:"update_time,omitempty"`
}

type BillingCycle struct {
	Frequency     *Frequency     `json:"frequency,omitempty"`
	TrnureType    string         `json:"trnure_type,omitempty"`
	Sqeuence      int            `json:"sqeuence,omitempty"`
	TotalCycles   int            `json:"total_cycles,omitempty"`
	PricingScheme *PricingScheme `json:"pricing_scheme,omitempty"`
}

type Frequency struct {
	IntervalUnit  string `json:"interval_unit"`
	IntervalCount int    `json:"interval_count"`
}

type PricingScheme struct {
	FixedPrice *FixedPrice `json:"fixed_price"`
	Version    int         `json:"version,omitempty"`
	CreateTime string      `json:"create_time,omitempty"`
	UpdateTime string      `json:"update_time,omitempty"`
}

type FixedPrice struct {
	Value        string `json:"value"`
	CurrencyCode string `json:"currency_code"`
}

type PaymentPreference struct {
	AutoBillOutstanding     bool      `json:"auto_bill_outstanding"`
	SetupFee                *V1Amount `json:"setup_fee,omitempty"`
	SetupFeeFailureAction   string    `json:"setup_fee_failure_action"`
	PaymentFailureThreshold int       `json:"payment_failure_threshold"`
}

type V1Tax struct {
	Percentage string `json:"percentage"`
	Inclusive  bool   `json:"inclusive"`
}
