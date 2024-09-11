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
		Uri:             "/v2/checkout/orders/{{.id}}?{{.params}}",
		ValidStatusCode: http.StatusOK,
		Do:              GetPayPal,
	}
	// UpdateOrder order_id 更新订单 PATCH
	UpdateOrder Method = Method{
		Uri:             "/v2/checkout/orders/{{.id}}",
		ValidStatusCode: http.StatusNoContent,
		Do:              PatchPayPal,
	}
	// ConfirmOrder order_id 订单支付确认 POST
	ConfirmOrder Method = Method{
		Uri:             "/v2/checkout/orders/{{.id}}/confirm-payment-source",
		ValidStatusCode: http.StatusOK,
		Do:              PostPayPal,
		Checker: paypay.InjectRuler(map[string][]paypay.Ruler{
			"/v2/checkout/orders/{{.id}}/confirm-payment-source": []paypay.Ruler{
				paypay.NewRuler("payment_source", `payment_source != nil`, "payment_source 不为空"),
			},
		}),
	}
	// AuthorizeOrder order_id 订单支付授权 POST
	AuthorizeOrder Method = Method{
		Uri:             "/v2/checkout/orders/{{.id}}/authorize",
		ValidStatusCode: http.StatusCreated,
		Do:              PostPayPal,
		Checker: paypay.InjectRuler(map[string][]paypay.Ruler{
			"/v2/checkout/orders/{{.id}}/authorize": []paypay.Ruler{
				paypay.NewRuler("payment_source", `payment_source != nil`, "payment_source 不为空"),
			},
		}),
	}
	// CaptureOrder order_id 订单支付捕获 POST
	CaptureOrder Method = Method{
		Uri:             "/v2/checkout/orders/{{.id}}/capture",
		ValidStatusCode: http.StatusCreated,
		Do:              PostPayPal,
		Checker: paypay.InjectRuler(map[string][]paypay.Ruler{
			"/v2/checkout/orders/{{.id}}/capture": []paypay.Ruler{
				paypay.NewRuler("payment_source", `payment_source != nil`, "payment_source 不为空"),
			},
		}),
	}
	// AddTracking4Order order_id 订单追踪 POST
	AddTracking4Order Method = Method{
		Uri:             "/v2/checkout/orders/{{.id}}/track",
		ValidStatusCode: http.StatusCreated,
		Do:              PostPayPal,
		Checker: paypay.InjectRuler(map[string][]paypay.Ruler{
			"/v2/checkout/orders/{{.id}}/track": []paypay.Ruler{
				paypay.NewRuler("tracking_number", `tracking_number != nil`, "运单号不为空"),
				paypay.NewRuler("carrier", `carrier != nil`, "承运机构不为空"),
				paypay.NewRuler("capture_id", `capture_id != nil`, "capture_id 不为空"),
			},
		}),
	}
	// UpOrCancelTracking4Order order_id, tracker_id 更新或取消订单追踪 POST
	UpOrCancelTracking4Order Method = Method{
		Uri:             "/v2/checkout/orders/{{.id}}/trackers/{{.tracker_id}}",
		ValidStatusCode: http.StatusNoContent,
		Do:              PostPayPal,
	}
)

// 支付相关
var (
	paymentAuthorizeDetail  = "/v2/payments/authorizations/%s"             // authorization_id 支付授权详情 GET
	paymentAuthorizeCapture = "/v2/payments/authorizations/%s/capture"     // authorization_id 支付授权捕获 POST
	paymentReauthorize      = "/v2/payments/authorizations/%s/reauthorize" // authorization_id 重新授权支付授权 POST
	paymentAuthorizeVoid    = "/v2/payments/authorizations/%s/void"        // authorization_id 作废支付授权 POST
	paymentCaptureDetail    = "/v2/payments/captures/%s"                   // capture_id 支付捕获详情 GET
	paymentCaptureRefund    = "/v2/payments/captures/%s/refund"            // capture_id 支付捕获退款 POST
	paymentRefundDetail     = "/v2/payments/refunds/%s"                    // refund_id 支付退款详情 GET
)

var (
	// 支出
	createBatchPayout         = "/v1/payments/payouts"                // 创建批量支出 POST
	showPayoutBatchDetail     = "/v1/payments/payouts/%s"             // payout_batch_id 获取批量支出详情 GET
	showPayoutItemDetail      = "/v1/payments/payouts-item/%s"        // payout_item_id 获取支出项目详情 GET
	cancelUnclaimedPayoutItem = "/v1/payments/payouts-item/%s/cancel" // payout_item_id 取消支出项目 POST

	// 订阅
	subscriptionCreate = "/v1/billing/plans" // 创建订阅 POST
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
