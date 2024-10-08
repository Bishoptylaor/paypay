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
 @Time    : 2024/9/11 -- 18:08
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: payment.go
*/

package entity

type ShowAuthorizedPaymentDetailsRes struct {
	EmptyRes
	Response *PaymentAuthorizationDetail `json:"response,omitempty"`
}

type CaptureAuthorizedPaymentRes struct {
	EmptyRes
	Response *PaymentCapture `json:"response,omitempty"`
}

type ReauthorizePaymentRes struct {
	EmptyRes
	Response *PaymentAuthorizationDetail `json:"response,omitempty"`
}

type VoidAuthorizePaymentRes struct {
	EmptyRes
}

type ShowCapturedPaymentRes struct {
	EmptyRes
	Response *PaymentCapture `json:"response,omitempty"`
}

type RefundCapturedPaymentRes struct {
	EmptyRes
	Response *PaymentRefund `json:"response,omitempty"`
}

type ShowRefundDetailsRes struct {
	EmptyRes
	Response *PaymentRefund `json:"response,omitempty"`
}
