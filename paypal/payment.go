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

package paypal

import (
	"context"
	"github.com/Bishoptylaor/paypay"
	"github.com/Bishoptylaor/paypay/paypal/consts"
	"github.com/Bishoptylaor/paypay/paypal/entity"
	"github.com/Bishoptylaor/paypay/pkg"
)

// ShowAuthorizedPaymentDetails
// 支付授权详情（Show details for authorized payment）
// 文档：https://developer.paypal.com/docs/api/payments/v2/#authorizations_get
func (c *Client) ShowAuthorizedPaymentDetails(ctx context.Context, authorizationId string) (res *entity.ShowAuthorizedPaymentDetailsRes, err error) {
	method := ShowAuthorizedPaymentDetails
	if authorizationId == pkg.NULL {
		return nil, pkg.ErrPaypalMissingAuthorizeId
	}

	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"authorization_id": authorizationId,
	}), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.ShowAuthorizedPaymentDetailsRes{EmptyRes: emptyRes}
	res.Response = new(entity.PaymentAuthorizationDetail)
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, res.Response); err != nil {
		return res, err
	}
	return res, nil
}

// CaptureAuthorizedPayment
// 支付授权捕获（Capture authorized payment）
// 文档：https://developer.paypal.com/docs/api/payments/v2/#authorizations_capture
func (c *Client) CaptureAuthorizedPayment(ctx context.Context, authorizationId string, pl paypay.Payload) (res *entity.CaptureAuthorizedPaymentRes, err error) {
	method := CaptureAuthorizedPayment
	if authorizationId == pkg.NULL {
		return nil, pkg.ErrPaypalMissingAuthorizeId
	}

	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"authorization_id": authorizationId,
	}), pl, nil, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.CaptureAuthorizedPaymentRes{EmptyRes: emptyRes}
	res.Response = new(entity.PaymentCapture)
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, res.Response); err != nil {
		return res, err
	}
	return res, nil
}

// ReauthorizePayment
// 重新授权支付授权（Reauthorize authorized payment）
// 文档：https://developer.paypal.com/docs/api/payments/v2/#authorizations_reauthorize
func (c *Client) ReauthorizePayment(ctx context.Context, authorizationId string, pl paypay.Payload) (res *entity.ReauthorizePaymentRes, err error) {
	method := ReauthorizePayment
	if authorizationId == pkg.NULL {
		return nil, pkg.ErrPaypalMissingAuthorizeId
	}

	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"authorization_id": authorizationId,
	}), pl, nil, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.ReauthorizePaymentRes{EmptyRes: emptyRes}
	res.Response = new(entity.PaymentAuthorizationDetail)
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, res.Response); err != nil {
		return res, err
	}
	return res, nil
}

// VoidAuthorizePayment
// 作废支付授权（Void authorized payment）
// 文档：https://developer.paypal.com/docs/api/payments/v2/#authorizations_void
func (c *Client) VoidAuthorizePayment(ctx context.Context, authorizationId string, pl paypay.Payload) (res *entity.VoidAuthorizePaymentRes, err error) {
	method := VoidAuthorizePayment
	if authorizationId == pkg.NULL {
		return nil, pkg.ErrPaypalMissingAuthorizeId
	}

	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"authorization_id": authorizationId,
	}), pl, nil, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.VoidAuthorizePaymentRes{EmptyRes: emptyRes}
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, new(struct{})); err != nil {
		return res, err
	}
	return res, nil
}

// ShowCapturedPayment
// 支付捕获详情（Show captured payment details）
// 文档：https://developer.paypal.com/docs/api/payments/v2/#captures_get
func (c *Client) ShowCapturedPayment(ctx context.Context, captureId string) (res *entity.ShowCapturedPaymentRes, err error) {
	method := ShowCapturedPayment
	if captureId == pkg.NULL {
		return nil, pkg.ErrPaypalMissingCaptureId
	}

	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"capture_id": captureId,
	}), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.ShowCapturedPaymentRes{EmptyRes: emptyRes}
	res.Response = new(entity.PaymentCapture)
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, res.Response); err != nil {
		return res, err
	}
	return res, nil
}

// RefundCapturedPayment
// 支付捕获退款（Refund captured payment）
// 文档：https://developer.paypal.com/docs/api/payments/v2/#captures_refund
func (c *Client) RefundCapturedPayment(ctx context.Context, captureId string, pl paypay.Payload) (res *entity.RefundCapturedPaymentRes, err error) {
	method := RefundCapturedPayment
	if captureId == pkg.NULL {
		return nil, pkg.ErrPaypalMissingCaptureId
	}

	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"capture_id": captureId,
	}), pl, nil, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.RefundCapturedPaymentRes{EmptyRes: emptyRes}
	res.Response = new(entity.PaymentRefund)
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, res.Response); err != nil {
		return res, err
	}
	return res, nil
}

// ShowRefundDetails
// 支付捕获退款（Refund captured payment）
// 文档：https://developer.paypal.com/docs/api/payments/v2/#captures_refund
func (c *Client) ShowRefundDetails(ctx context.Context, refundId string) (res *entity.ShowRefundDetailsRes, err error) {
	method := ShowRefundDetails
	if refundId == pkg.NULL {
		return nil, pkg.ErrPaypalMissingRefundId
	}

	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"refund_id": refundId,
	}), nil, nil, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.ShowRefundDetailsRes{EmptyRes: emptyRes}
	res.Response = new(entity.PaymentRefund)
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, res.Response); err != nil {
		return res, err
	}
	return res, nil
}
