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
 @Time    : 2024/9/10 -- 15:35
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: methods.go
*/

package paypal

import (
	"github.com/Bishoptylaor/paypay"
	"net/http"
)

type Method struct {
	Uri             string
	ValidStatusCode int
	Do              func(c *Client) DoPayPalRequest
	Checker         paypay.PayloadRuler
}

type QueryPlaceholder string

var (
	OrderId         QueryPlaceholder = "order_id"
	TrackerId       QueryPlaceholder = "tracker_id"
	AuthorizationId QueryPlaceholder = "authorization_id"
	CaptureId       QueryPlaceholder = "capture_id"
	RefundId        QueryPlaceholder = "refund_id"
	PayoutItemId    QueryPlaceholder = "payout_item_id"
	InvoiceId       QueryPlaceholder = "invoice_id"
	TemplateId      QueryPlaceholder = "template_id"
	SubscriptionId  QueryPlaceholder = "subscription_id"
	PlanId          QueryPlaceholder = "plan_id"
	ProductId       QueryPlaceholder = "product_id"
	DisputeId       QueryPlaceholder = "dispute_id"
)

func (qp QueryPlaceholder) String() string {
	return string(qp)
}

// 订单相关
var (
	// CreateOrder 创建订单 POST
	CreateOrder Method = Method{
		Uri:             "/v2/checkout/orders",
		ValidStatusCode: http.StatusCreated,
		Do:              PostPayPal,
		Checker: paypay.InjectRuler(map[string][]paypay.Ruler{
			"/v2/checkout/orders": []paypay.Ruler{
				paypay.NewRuler("purchase_units",
					`purchase_units != nil && len(purchase_units) <= 10 &&
all(purchase_units, {.Amount != nil}) `,
					"purchase_units 最多一次性传入10个",
				),
				paypay.NewRuler("intent", `intent in ["CAPTURE", "AUTHORIZE"]`, ""),
			},
		}),
	}
	// ShowOrderDetails order_id 查看订单详情 GET
	ShowOrderDetails Method = Method{
		Uri:             "/v2/checkout/orders/{{.order_id}}?{{.params}}",
		ValidStatusCode: http.StatusOK,
		Do:              GetPayPal,
	}
	// UpdateOrder order_id 更新订单 PATCH
	UpdateOrder Method = Method{
		Uri:             "/v2/checkout/orders/{{.order_id}}",
		ValidStatusCode: http.StatusNoContent,
		Do:              PatchPayPal,
	}
	// ConfirmOrder order_id 订单支付确认 POST
	ConfirmOrder Method = Method{
		Uri:             "/v2/checkout/orders/{{.order_id}}/confirm-payment-source",
		ValidStatusCode: http.StatusOK,
		Do:              PostPayPal,
		Checker: paypay.InjectRuler(map[string][]paypay.Ruler{
			"/v2/checkout/orders/{{.order_id}}/confirm-payment-source": []paypay.Ruler{
				paypay.NewRuler("payment_source", `payment_source != nil`, "payment_source 不为空"),
			},
		}),
	}
	// AuthorizeOrder order_id 订单支付授权 POST
	AuthorizeOrder Method = Method{
		Uri:             "/v2/checkout/orders/{{.order_id}}/authorize",
		ValidStatusCode: http.StatusCreated,
		Do:              PostPayPal,
	}
	// CaptureOrder order_id 订单支付捕获 POST
	CaptureOrder Method = Method{
		Uri:             "/v2/checkout/orders/{{.order_id}}/capture",
		ValidStatusCode: http.StatusCreated,
		Do:              PostPayPal,
		Checker: paypay.InjectRuler(map[string][]paypay.Ruler{
			"/v2/checkout/orders/{{.order_id}}/capture": []paypay.Ruler{
				paypay.NewRuler("payment_source", `payment_source != nil`, "payment_source 不为空"),
			},
		}),
	}
	// AddTracking4Order order_id 订单追踪 POST
	AddTracking4Order Method = Method{
		Uri:             "/v2/checkout/orders/{{.order_id}}/track",
		ValidStatusCode: http.StatusCreated,
		Do:              PostPayPal,
		Checker: paypay.InjectRuler(map[string][]paypay.Ruler{
			"/v2/checkout/orders/{{.order_id}}/track": []paypay.Ruler{
				paypay.NewRuler("tracking_number", `tracking_number != nil`, "运单号不为空"),
				paypay.NewRuler("carrier", `carrier != nil`, "承运机构不为空"),
				paypay.NewRuler("capture_id", `capture_id != nil`, "capture_id 不为空"),
			},
		}),
	}
	// UpOrCancelTracking4Order order_id, tracker_id 更新或取消订单追踪 POST
	UpOrCancelTracking4Order Method = Method{
		Uri:             "/v2/checkout/orders/{{.order_id}}/trackers/{{.tracker_id}}",
		ValidStatusCode: http.StatusNoContent,
		Do:              PostPayPal,
	}
)

// 支付相关
var (
	// ShowAuthorizedPaymentDetails authorization_id 支付授权详情 GET
	ShowAuthorizedPaymentDetails Method = Method{
		Uri:             "/v2/payments/authorizations/{{.authorization_id}}",
		ValidStatusCode: http.StatusCreated,
		Do:              GetPayPal,
	}
	// CaptureAuthorizedPayment authorization_id 支付授权捕获 POST
	CaptureAuthorizedPayment Method = Method{
		Uri:             "/v2/payments/authorizations/{{.authorization_id}}/capture",
		ValidStatusCode: http.StatusCreated,
		Do:              PostPayPal,
	}
	// ReauthorizePayment authorization_id 重新授权支付授权 POST
	ReauthorizePayment Method = Method{
		Uri:             "/v2/payments/authorizations/{{.authorization_id}}/reauthorize",
		ValidStatusCode: http.StatusCreated,
		Do:              PostPayPal,
	}
	// VoidAuthorizePayment authorization_id 作废支付授权 POST
	VoidAuthorizePayment Method = Method{
		Uri:             "/v2/payments/authorizations/{{.authorization_id}}/void",
		ValidStatusCode: http.StatusNoContent,
		Do:              PostPayPal,
	}
	// ShowCapturedPayment capture_id 支付捕获详情 GET
	ShowCapturedPayment Method = Method{
		Uri:             "/v2/payments/captures/{{.capture_id}}",
		ValidStatusCode: http.StatusOK,
		Do:              GetPayPal,
	}
	// RefundCapturedPayment capture_id 支付捕获退款 POST
	RefundCapturedPayment Method = Method{
		Uri:             "/v2/payments/captures/{{.capture_id}}/refund",
		ValidStatusCode: http.StatusCreated,
		Do:              PostPayPal,
	}
	// ShowRefundDetails refund_id 支付退款详情 GET
	ShowRefundDetails Method = Method{
		Uri:             "/v2/payments/refunds/{{.refund_id}}",
		ValidStatusCode: http.StatusOK,
		Do:              GetPayPal,
	}
)

// 支付 Payouts
var (
	// CreateBatchPayout 创建批量支付 POST
	CreateBatchPayout Method = Method{
		Uri:             "/v1/payments/payouts",
		ValidStatusCode: http.StatusCreated,
		Do:              PostPayPal,
		Checker: paypay.InjectRuler(map[string][]paypay.Ruler{
			"/v1/payments/payouts": []paypay.Ruler{
				paypay.NewRuler("items", `items != nil`, "支付详情列表不能为空"),
				paypay.NewRuler("sender_batch_header", `sender_batch_header != nil`, "批次号不能为空"),
			},
		}),
	}
	// ShowPayoutBatchDetail payout_batch_id 获取批量支付详情 GET
	ShowPayoutBatchDetail Method = Method{
		Uri:             "/v1/payments/payouts/{{.payout_batch_id}}?{{.params}}",
		ValidStatusCode: http.StatusOK,
		Do:              GetPayPal,
	}
	// ShowPayoutItemDetail payout_item_id 获取支付项目详情 GET
	ShowPayoutItemDetail Method = Method{
		Uri:             "/v1/payments/payouts-item/{{.payout_item_id}}",
		ValidStatusCode: http.StatusOK,
		Do:              GetPayPal,
	}
	// CancelUnclaimedPayoutItem payout_item_id 取消支付项目 POST
	CancelUnclaimedPayoutItem Method = Method{
		Uri:             "/v1/payments/payouts-item/{{.payout_item_id}}/cancel",
		ValidStatusCode: http.StatusOK,
		Do:              PostPayPal,
	}
)

// 发票 Invoices
var (
	// CreateInvoices 创建发票 POST
	CreateInvoices Method = Method{
		Uri:             "/v2/invoicing/invoices",
		ValidStatusCode: http.StatusOK,
		Do:              PostPayPal,
	}
	// ListInvoices 查看发票列表 GET
	ListInvoices Method = Method{
		Uri:             "/v2/invoicing/invoices?{{.params}}",
		ValidStatusCode: http.StatusOK,
		Do:              GetPayPal,
	}
	// SendInvoice invoice_id 发送发票 POST
	SendInvoice Method = Method{
		Uri:             "/v2/invoicing/invoices/{{.invoice_id}}/send",
		ValidStatusCode: http.StatusAccepted,
		Do:              PostPayPal,
	}
	// SendInvoiceReminder invoice_id 发送发票提醒 POST
	SendInvoiceReminder Method = Method{
		Uri:             "/v2/invoicing/invoices/{{.invoice_id}}/remind",
		ValidStatusCode: http.StatusNoContent,
		Do:              PostPayPal,
	}
	// CancelSentInvoice invoice_id 取消已发送发票 POST
	CancelSentInvoice Method = Method{
		Uri:             "/v2/invoicing/invoices/{{.invoice_id}}/cancel",
		ValidStatusCode: http.StatusNoContent,
		Do:              PostPayPal,
	}
	// RecordPaymentForInvoice invoice_id 记录发票付款 POST
	RecordPaymentForInvoice Method = Method{
		Uri:             "/v2/invoicing/invoices/{{.invoice_id}}/payments",
		ValidStatusCode: http.StatusOK,
		Do:              PostPayPal,
		Checker: paypay.InjectRuler(map[string][]paypay.Ruler{
			"/v2/invoicing/invoices/{{.invoice_id}}/payments": []paypay.Ruler{
				paypay.NewRuler(
					"method",
					`method != nil && method in [BANK_TRANSFER, CASH, CHECK, CREDIT_CARD, DEBIT_CARD, PAYPAL, WIRE_TRANSFER, OTHER]`,
					"支付方案不能为空且要在以下内容中选取 [BANK_TRANSFER, CASH, CHECK, CREDIT_CARD, DEBIT_CARD, PAYPAL, WIRE_TRANSFER, OTHER]"),
			},
		}),
	}
	// DeleteExternalPayment invoice_id,transaction_id 删除额外支付 DELETE
	DeleteExternalPayment Method = Method{
		Uri:             "/v2/invoicing/invoices/{{.invoice_id}}/payments/{{.transaction_id}}",
		ValidStatusCode: http.StatusNoContent,
		Do:              DeletePayPal,
	}
	// RecordRefundForInvoice invoice_id 记录发票退款 POST
	RecordRefundForInvoice Method = Method{
		Uri:             "/v2/invoicing/invoices/{{.invoice_id}}/refunds",
		ValidStatusCode: http.StatusOK,
		Do:              PostPayPal,
		Checker: paypay.InjectRuler(map[string][]paypay.Ruler{
			"/v2/invoicing/invoices/{{.invoice_id}}/payments": []paypay.Ruler{
				paypay.NewRuler(
					"method",
					`method != nil && method in [BANK_TRANSFER, CASH, CHECK, CREDIT_CARD, DEBIT_CARD, PAYPAL, WIRE_TRANSFER, OTHER]`,
					"支付方案不能为空且要在以下内容中选取 [BANK_TRANSFER, CASH, CHECK, CREDIT_CARD, DEBIT_CARD, PAYPAL, WIRE_TRANSFER, OTHER]"),
			},
		}),
	}
	// DeleteExternalRefund invoice_id,transaction_id 删除额外退款 DELETE
	DeleteExternalRefund Method = Method{
		Uri:             "/v2/invoicing/invoices/{{.invoice_id}}/refunds/{{.transaction_id}}",
		ValidStatusCode: http.StatusNoContent,
		Do:              DeletePayPal,
	}
	// GenerateInvoiceQRCode invoice_id 生成发票二维码 POST
	GenerateInvoiceQRCode Method = Method{
		Uri:             "/v2/invoicing/invoices/{{.invoice_id}}/generate-qr-code",
		ValidStatusCode: http.StatusOK,
		Do:              PostPayPal,
	}
	// GenerateInvoiceNumber 生成下个发票号码 POST
	GenerateInvoiceNumber Method = Method{
		Uri:             "/v2/invoicing/generate-next-invoice-number",
		ValidStatusCode: http.StatusOK,
		Do:              PostPayPal,
	}
	// ShowInvoiceDetail invoice_id 获取发票详情 GET
	ShowInvoiceDetail Method = Method{
		Uri:             "/v2/invoicing/invoices/{{.invoice_id}}",
		ValidStatusCode: http.StatusOK,
		Do:              GetPayPal,
	}
	// FullyUpdateInvoice invoice_id 全量更新发票 PUT
	FullyUpdateInvoice Method = Method{
		Uri:             "/v2/invoicing/invoices/{{.invoice_id}}?{{.params}}",
		ValidStatusCode: http.StatusOK,
		Do:              PutPayPal,
		Checker: paypay.InjectRuler(map[string][]paypay.Ruler{
			"/v2/invoicing/invoices/{{.invoice_id}}?{{.params}}": []paypay.Ruler{
				paypay.NewRuler("detail", `detail != nil`, "detail 不为空"),
			},
		}),
	}
	// DeleteInvoice // invoice_id 删除发票 DELETE
	DeleteInvoice Method = Method{
		Uri:             "/v2/invoicing/invoices/{{.invoice_id}}",
		ValidStatusCode: http.StatusNoContent,
		Do:              DeletePayPal,
	}
	// SearchInvoice 搜索发票 POST
	SearchInvoice Method = Method{
		Uri:             "/v2/invoicing/search-invoices?{{.params}}",
		ValidStatusCode: http.StatusOK,
		Do:              PostPayPal,
		Checker:         nil,
	}
	// ListInvoiceTemplate 获取发票模板列表 GET
	ListInvoiceTemplate Method = Method{
		Uri:             "/v2/invoicing/templates?{{.params}}",
		ValidStatusCode: http.StatusOK,
		Do:              GetPayPal,
	}
	// CreateInvoiceTemplate 创建发票模板 POST
	CreateInvoiceTemplate Method = Method{
		Uri:             "/v2/invoicing/templates",
		ValidStatusCode: http.StatusOK,
		Do:              PostPayPal,
	}
	// ShowTemplateDetails template_id 获取模版详情 GET
	ShowTemplateDetails Method = Method{
		Uri:             "/v2/invoicing/templates/{{.template_id}}",
		ValidStatusCode: http.StatusOK,
		Do:              GetPayPal,
	}
	// FullyUpdateInvoiceTemplate template_id 全量更新发票模板 PUT
	FullyUpdateInvoiceTemplate Method = Method{
		Uri:             "/v2/invoicing/templates/{{.template_id}}",
		ValidStatusCode: http.StatusOK,
		Do:              PutPayPal,
	}
	// DeleteInvoiceTemplate template_id 删除发票模板 DELETE
	DeleteInvoiceTemplate Method = Method{
		Uri:             "/v2/invoicing/templates/{{.template_id}}",
		ValidStatusCode: http.StatusNoContent,
		Do:              DeletePayPal,
	}
)

// 订阅 类似支付宝周期扣款
var (
	// CreatePlan 创建订阅 POST
	CreatePlan Method = Method{
		Uri:             "/v1/billing/plans",
		ValidStatusCode: http.StatusCreated,
		Do:              PostPayPal,
		Checker: paypay.InjectRuler(map[string][]paypay.Ruler{
			"/v1/billing/plans": []paypay.Ruler{
				paypay.NewRuler("product_id", `product_id != nil`, "通过目录产品 API 创建的产品的 ID 不为空"),
				paypay.NewRuler("name", `name != nil`, "计划名 不为空"),
				paypay.NewRuler("billing_cycles", `billing_cycles != nil`, "billing_cycles 不为空"),
				paypay.NewRuler("payment_preferences", `payment_preferences != nil`, "payment_preferences 不为空"),
				paypay.NewRuler("status", `(status != nil && status in ["CREATED", "INACTIVE", "ACTIVE"]) || status == nil`,
					"status 默认 ACTIVE 或者 枚举值有误：ACTIVE，INACTIVE，CREATED"),
			},
		}),
	}
	// ListPlans 列表展示计划 GET
	ListPlans Method = Method{
		Uri:             "/v1/billing/plans?{{.params}}",
		ValidStatusCode: http.StatusOK,
		Do:              GetPayPal,
	}
	// ShowPlanDetails plan_id 展示计划详情 GET
	ShowPlanDetails Method = Method{
		Uri:             "/v1/billing/plans/{{.plan_id}}",
		ValidStatusCode: http.StatusOK,
		Do:              GetPayPal,
	}
	// UpdatePlan plan_id 更新计划 patch
	UpdatePlan Method = Method{
		Uri:             "/v1/billing/plans/{{.plan_id}}",
		ValidStatusCode: http.StatusNoContent,
		Do:              PatchPayPal,
	}
	// ActivePlan plan_id 激活计划 POST
	ActivePlan Method = Method{
		Uri:             "/v1/billing/plans/{{.plan_id}}/activate",
		ValidStatusCode: http.StatusNoContent,
		Do:              PostPayPal,
	}
	DeactivePlan Method = Method{
		Uri:             "/v1/billing/plans/{{.plan_id}}/deactivate",
		ValidStatusCode: http.StatusNoContent,
		Do:              PostPayPal,
	}
	// UpdatePricing plan_id 更新计划价格方案 POST
	UpdatePricing Method = Method{
		Uri:             "/v1/billing/plans/{{.plan_id}}/update-pricing-schemes",
		ValidStatusCode: http.StatusNoContent,
		Do:              PostPayPal,
	}

	// CreateSubscription 创建订阅 POST
	CreateSubscription Method = Method{
		Uri:             "/v1/billing/subscriptions",
		ValidStatusCode: http.StatusCreated,
		Do:              PostPayPal,
		Checker: paypay.InjectRuler(map[string][]paypay.Ruler{
			"/v1/billing/subscriptions": []paypay.Ruler{
				paypay.NewRuler("plan_id", `plan_id != nil`, "plan_id 不为空"),
			},
		}),
	}
	// ShowSubscriptionDetails subscription_id 展示订阅详情 GET
	ShowSubscriptionDetails Method = Method{
		Uri:             "/v1/billing/subscriptions/{{.subscription_id}}",
		ValidStatusCode: http.StatusOK,
		Do:              GetPayPal,
	}
	// UpdateSubscription subscription_id 更新订阅 PATCH
	UpdateSubscription Method = Method{
		Uri:             "/v1/billing/subscriptions/{{.subscription_id}}",
		ValidStatusCode: http.StatusNoContent,
		Do:              PatchPayPal,
	}
	// RevisePlanOrQuantityOfSubsription subscription_id 更新计划或者数量 POST
	RevisePlanOrQuantityOfSubsription Method = Method{
		Uri:             "/v1/billing/subscriptions/{{.subscription_id}}/revise",
		ValidStatusCode: http.StatusOK,
		Do:              PostPayPal,
	}
	// SuspendSubscription subscription_id 暂定订阅计划 POST
	SuspendSubscription Method = Method{
		Uri:             "/v1/billing/subscriptions/{{.subscription_id}}/suspend",
		ValidStatusCode: http.StatusNoContent,
		Do:              PostPayPal,
		Checker: paypay.InjectRuler(map[string][]paypay.Ruler{
			"/v1/billing/subscriptions/{{.subscription_id}}/suspend": []paypay.Ruler{
				paypay.NewRuler("reason", `reason != nil`, "暂定原因 reason 不为空"),
			},
		}),
	}
	// CancelSubscription subscription_id 取消订阅 POST
	CancelSubscription Method = Method{
		Uri:             "/v1/billing/subscriptions/{{.subscription_id}}/cancel",
		ValidStatusCode: http.StatusNoContent,
		Do:              PostPayPal,
		Checker: paypay.InjectRuler(map[string][]paypay.Ruler{
			"/v1/billing/subscriptions/{{.subscription_id}}/cancel": []paypay.Ruler{
				paypay.NewRuler("reason", `reason != nil`, "取消原因 reason 不为空"),
			},
		}),
	}
	// ActivateSubscription subscription_id 激活订阅 POST
	ActivateSubscription Method = Method{
		Uri:             "/v1/billing/subscriptions/{{.subscription_id}}/activate",
		ValidStatusCode: http.StatusNoContent,
		Do:              PostPayPal,
	}
	// CaptureAuthoriedPaymentOnSubscription subscription_id 捕获订阅的授权支付信息 POST
	CaptureAuthoriedPaymentOnSubscription Method = Method{
		Uri:             "/v1/billing/subscriptions/{{.subscription_id}}/capture",
		ValidStatusCode: http.StatusAccepted,
		Do:              PostPayPal,
		Checker: paypay.InjectRuler(map[string][]paypay.Ruler{
			"/v1/billing/subscriptions/{{.subscription_id}}/capture": []paypay.Ruler{
				paypay.NewRuler("note", `note != nil`, "note 不为空"),
				paypay.NewRuler("capture_type", `capture_type == "OUTSTANDING_BALANCE"`, "capture_type = OUTSTANDING_BALANCE The outstanding balance that the subscriber must clear"),
				paypay.NewRuler("amount", `amount != nil && amount.currency_code != nil && amount.value != nil`, "amount 及其字段不为空"),
			},
		}),
	}
	// ListTransactions4Subscription subscription_id 列出一个订阅的所有交易记录
	ListTransactions4Subscription Method = Method{
		Uri:             "/v1/billing/subscriptions/{{.subscription_id}}/transactions?{{.params}}",
		ValidStatusCode: http.StatusOK,
		Do:              GetPayPal,
		Checker: paypay.InjectRuler(map[string][]paypay.Ruler{
			"/v1/billing/subscriptions/{{.subscription_id}}/transactions": []paypay.Ruler{
				paypay.NewRuler("start_time", `start_time != nil`, "query 参数 start_time 不为空"),
				paypay.NewRuler("end_time", `end_time != nil`, "query 参数 end_time 不为空"),
			},
		}),
	}
)

// Transaction 交易
var (
	// ListTransactions 交易列表
	ListTransactions Method = Method{
		Uri:             "/v1/reporting/transactions?{{.params}}",
		ValidStatusCode: http.StatusOK,
		Do:              GetPayPal,
		Checker: paypay.InjectRuler(map[string][]paypay.Ruler{
			"/v1/reporting/transactions": []paypay.Ruler{
				paypay.NewRuler("start_date", `start_date != nil`, "query 参数 start_date 不为空"),
				paypay.NewRuler("end_date", `end_date != nil`, "query 参数 end_date 不为空"),
			},
		}),
	}
	// ListAllBalances 获取所有余额
	ListAllBalances Method = Method{
		Uri:             "/v1/reporting/balances?{{.params}}",
		ValidStatusCode: http.StatusOK,
		Do:              GetPayPal,
	}
)

// Catalog Products 目录商品
var (
	// CreateProduct 创建商品 POST
	CreateProduct Method = Method{
		Uri:             "/v1/catalogs/products",
		ValidStatusCode: http.StatusCreated,
		Do:              PostPayPal,
		Checker: paypay.InjectRuler(map[string][]paypay.Ruler{
			"/v1/catalogs/products": []paypay.Ruler{
				paypay.NewRuler("name", `name != nil`, "产品名不为空"),
				// paypay.NewRuler("type",
				// 	`type in ["PHYSICAL", "DIGITAL", "SERVICE"]`,
				// 	"产品类型不为空，取值范围：[PHYSICAL, DIGITAL, SERVICE]"),
			},
		}),
	}
	// ListProducts 商品列表 GET
	ListProducts Method = Method{
		Uri:             "/v1/catalogs/products",
		ValidStatusCode: http.StatusOK,
		Do:              GetPayPal,
	}
	// ShowProductDetails product_id 获取商品详情 GET
	ShowProductDetails Method = Method{
		Uri:             "/v1/catalogs/products/{{.product_id}}",
		ValidStatusCode: http.StatusOK,
		Do:              GetPayPal,
	}
	// UpdateProduct product_id 更新商品 PATCH
	UpdateProduct Method = Method{
		Uri:             "/v1/catalogs/products/{{.product_id}}",
		ValidStatusCode: http.StatusNoContent,
		Do:              PatchPayPal,
	}
)

// Disputes 投诉 || 争议
var (
	// EscalateDisputeToClaim dispute_id 将投诉升级为索赔 POST
	EscalateDisputeToClaim Method = Method{
		Uri:             "/v1/customer/disputes/{{.dispute_id}}/escalate",
		ValidStatusCode: http.StatusOK,
		Do:              PostPayPal,
		Checker: paypay.InjectRuler(map[string][]paypay.Ruler{
			"/v1/customer/disputes/{{.dispute_id}}/escalate": []paypay.Ruler{
				paypay.NewRuler("note", `note != nil`, "note 不为空"),
			},
		}),
	}
	// AcceptOffer2Resolve dispute_id 客户接受商家的解决方案，结束争议 POST
	AcceptOffer2Resolve Method = Method{
		Uri:             "/v1/customer/disputes/{{.dispute_id}}/accept-offer",
		ValidStatusCode: http.StatusAccepted,
		Do:              PostPayPal,
	}
	// ListDisputes 获取争议列表 GET
	ListDisputes Method = Method{
		Uri:             "/v1/customer/disputes?{{.params}}",
		ValidStatusCode: http.StatusOK,
		Do:              GetPayPal,
	}
	// ProvideInfo4Dispute dispute_id 提供有效信息 POST
	ProvideInfo4Dispute Method = Method{
		Uri:             "/v1/customer/disputes/{{.dispute_id}}/provide-supporting-info",
		ValidStatusCode: http.StatusOK,
		Do:              EmptyPaypal,
	}
	// ShowDisputeDetails dispute_id 获取争议详情 GET
	ShowDisputeDetails Method = Method{
		Uri:             "/v1/customer/disputes/{{.dispute_id}}",
		ValidStatusCode: http.StatusOK,
		Do:              GetPayPal,
	}
	// PartiallyUpdateDispute dispute_id 部分更新 POST
	PartiallyUpdateDispute Method = Method{
		Uri:             "/v1/customer/disputes/{{.dispute_id}}",
		ValidStatusCode: http.StatusNoContent,
		Do:              PatchPayPal,
	}
	// DenyOffer2Resolve dispute_id 拒绝方案 POST
	DenyOffer2Resolve Method = Method{
		Uri:             "/v1/customer/disputes/{{.dispute_id}}/deny-offer",
		ValidStatusCode: http.StatusOK,
		Do:              PostPayPal,
	}
	// MakeOffer2Resolve dispute_id 发起一个解决方案 POST
	MakeOffer2Resolve Method = Method{
		Uri:             "/v1/customer/disputes/{{.dispute_id}}/make-offer",
		ValidStatusCode: http.StatusOK,
		Do:              PostPayPal,
	}
	// AppealDispute dispute_id 上诉 POST
	AppealDispute Method = Method{
		Uri:             "/v1/customer/disputes/{{.dispute_id}}/appeal",
		ValidStatusCode: http.StatusOK,
		Do:              EmptyPaypal,
	}
	ProvideEvidence Method = Method{
		Uri:             "/v1/customer/disputes/{{.dispute_id}}/provide-evidence",
		ValidStatusCode: http.StatusOK,
		Do:              EmptyPaypal,
	}
	AckReturnedItem Method = Method{
		Uri:             "/v1/customer/disputes/{{.dispute_id}}/acknowledge-return-item",
		ValidStatusCode: http.StatusOK,
		Do:              EmptyPaypal,
	}
	NotifyDispute2ThirdParty Method = Method{
		Uri:             "/v1/customer/disputes/{{.dispute_id}}/send-message",
		ValidStatusCode: http.StatusOK,
		Do:              EmptyPaypal,
	}
	AcceptClaim Method = Method{
		Uri:             "/v1/customer/disputes/{{.dispute_id}}/accept-claim",
		ValidStatusCode: http.StatusOK,
		Do:              PostPayPal,
		Checker: paypay.InjectRuler(map[string][]paypay.Ruler{
			"/v1/customer/disputes/{{.dispute_id}}/accept-claim": []paypay.Ruler{
				paypay.NewRuler("note", `note != nil`, "note 不为空"),
			},
		}),
	}
)

var EmptyMethod Method = Method{
	Uri:             "",
	ValidStatusCode: http.StatusOK,
	Do:              EmptyPaypal,
}
