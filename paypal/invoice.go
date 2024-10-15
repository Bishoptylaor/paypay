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

package paypal

import (
	"context"
	"github.com/Bishoptylaor/paypay"
	"github.com/Bishoptylaor/paypay/paypal/consts"
	"github.com/Bishoptylaor/paypay/paypal/entity"
	"github.com/Bishoptylaor/paypay/pkg"
)

// CreateInvoice
// 创建虚拟发票（Create draft invoice）
// 文档：https://developer.paypal.com/docs/api/invoicing/v2/#invoices_create
func (c *Client) CreateInvoice(ctx context.Context, pl paypay.Payload) (res *entity.CreateInvoiceRes, err error) {
	method := CreateInvoices
	c.EmptyChecker = method.Checker
	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, nil), pl, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.CreateInvoiceRes{EmptyRes: emptyRes}
	res.Response = new(entity.Invoice)
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, res.Response); err != nil {
		return res, err
	}
	return res, nil
}

// ListInvoice
// 发票列表（List invoices）
// 文档：https://developer.paypal.com/docs/api/invoicing/v2/#invoices_list
func (c *Client) ListInvoice(ctx context.Context, query paypay.Payload) (res *entity.ListInvoiceRes, err error) {
	method := ListInvoices
	c.EmptyChecker = method.Checker
	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"params": query.EncodeURLParams(),
	}), query, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.ListInvoiceRes{EmptyRes: emptyRes}
	res.Response = new(entity.InvoiceList)
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, res.Response); err != nil {
		return res, err
	}
	return res, nil
}

// SendInvoice
// 发票列表（List invoices）
// 文档：https://developer.paypal.com/docs/api/invoicing/v2/#invoices_send
func (c *Client) SendInvoice(ctx context.Context, invoiceId string, pl paypay.Payload) (res *entity.SendInvoiceRes, err error) {
	method := SendInvoice
	c.EmptyChecker = method.Checker
	if invoiceId == pkg.NULL {
		return nil, pkg.ErrPaypalMissingInvoiceId
	}
	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"invoice_id": invoiceId,
	}), pl, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.SendInvoiceRes{EmptyRes: emptyRes}
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, nil); err != nil {
		return res, err
	}
	return res, nil
}

// SendInvoiceReminder
// 发票列表（List invoices）
// 文档：https://developer.paypal.com/docs/api/invoicing/v2/#invoices_send
func (c *Client) SendInvoiceReminder(ctx context.Context, invoiceId string, pl paypay.Payload) (res *entity.SendInvoiceReminderRes, err error) {
	method := SendInvoiceReminder
	c.EmptyChecker = method.Checker
	if invoiceId == pkg.NULL {
		return nil, pkg.ErrPaypalMissingInvoiceId
	}
	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"invoice_id": invoiceId,
	}), pl, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.SendInvoiceReminderRes{EmptyRes: emptyRes}
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, nil); err != nil {
		return res, err
	}
	return res, nil
}

// CancelSentInvoice
// 取消已发送发票（Cancel sent invoice）
// 文档：https://developer.paypal.com/docs/api/invoicing/v2/#invoices_cancel
func (c *Client) CancelSentInvoice(ctx context.Context, invoiceId string, pl paypay.Payload) (res *entity.CancelSentInvoiceRes, err error) {
	method := CancelSentInvoice
	c.EmptyChecker = method.Checker
	if invoiceId == pkg.NULL {
		return nil, pkg.ErrPaypalMissingInvoiceId
	}
	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"invoice_id": invoiceId,
	}), pl, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.CancelSentInvoiceRes{EmptyRes: emptyRes}
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, nil); err != nil {
		return res, err
	}
	return res, nil
}

// RecordPaymentForInvoice
// 记录发票付款（Record payment for invoice）
// 文档：https://developer.paypal.com/docs/api/invoicing/v2/#invoices_payments
func (c *Client) RecordPaymentForInvoice(ctx context.Context, invoiceId string, pl paypay.Payload) (res *entity.RecordPaymentForInvoiceRes, err error) {
	method := RecordPaymentForInvoice
	c.EmptyChecker = method.Checker
	if invoiceId == pkg.NULL {
		return nil, pkg.ErrPaypalMissingInvoiceId
	}

	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"invoice_id": invoiceId,
	}), pl, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.RecordPaymentForInvoiceRes{EmptyRes: emptyRes}
	res.Response = new(entity.RecordPaymentForInvoice)
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, res.Response); err != nil {
		return res, err
	}
	return res, nil
}

// DeleteExternalPayment
// 删除额外支付（Delete external payment）
// 文档：https://developer.paypal.com/docs/api/invoicing/v2/#invoices_payments-delete
func (c *Client) DeleteExternalPayment(ctx context.Context, invoiceId, transactionId string) (res *entity.DeleteExternalPaymentRes, err error) {
	method := DeleteExternalPayment
	c.EmptyChecker = method.Checker
	if invoiceId == pkg.NULL || transactionId == pkg.NULL {
		return nil, pkg.ErrPaypalMissingInvoiceTransactionId
	}
	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"invoice_id":     invoiceId,
		"transaction_id": transactionId,
	}), nil, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.DeleteExternalPaymentRes{EmptyRes: emptyRes}
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, nil); err != nil {
		return res, err
	}
	return res, nil
}

// RecordRefundForInvoice
// 记录发票退款（Record refund for invoice）
// 文档：https://developer.paypal.com/docs/api/invoicing/v2/#invoices_refunds
func (c *Client) RecordRefundForInvoice(ctx context.Context, invoiceId string, pl paypay.Payload) (res *entity.RecordRefundForInvoiceRes, err error) {
	method := RecordRefundForInvoice
	c.EmptyChecker = method.Checker
	if invoiceId == pkg.NULL {
		return nil, pkg.ErrPaypalMissingInvoiceId
	}

	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"invoice_id": invoiceId,
	}), pl, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.RecordRefundForInvoiceRes{EmptyRes: emptyRes}
	res.Response = new(entity.RecordRefundForInvoice)
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, res.Response); err != nil {
		return res, err
	}
	return res, nil
}

// DeleteExternalRefund
// 删除额外支付（Delete external refund）
// 文档：https://developer.paypal.com/docs/api/invoicing/v2/#invoices_payments-delete
func (c *Client) DeleteExternalRefund(ctx context.Context, invoiceId, transactionId string) (res *entity.DeleteExternalRefundRes, err error) {
	method := DeleteExternalRefund
	c.EmptyChecker = method.Checker
	if invoiceId == pkg.NULL || transactionId == pkg.NULL {
		return nil, pkg.ErrPaypalMissingInvoiceTransactionId
	}

	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"invoice_id":     invoiceId,
		"transaction_id": transactionId,
	}), nil, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.DeleteExternalRefundRes{EmptyRes: emptyRes}
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, nil); err != nil {
		return res, err
	}
	return res, nil
}

// GenerateInvoiceQRCode
// 生成发票二维码（Generate QR code）
// 文档：https://developer.paypal.com/docs/api/invoicing/v2/#invoices_generate-qr-code
func (c *Client) GenerateInvoiceQRCode(ctx context.Context, invoiceId string, pl paypay.Payload) (res *entity.GenerateInvoiceQRCodeRes, err error) {
	method := GenerateInvoiceQRCode
	c.EmptyChecker = method.Checker
	if invoiceId == pkg.NULL {
		return nil, pkg.ErrPaypalMissingInvoiceId
	}
	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"invoice_id": invoiceId,
	}), pl, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.GenerateInvoiceQRCodeRes{EmptyRes: emptyRes}
	res.Response = new(entity.Qrcode)
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, res.Response); err != nil {
		return res, err
	}
	res.Response = &entity.Qrcode{Base64Image: string(bs)}
	return res, nil
}

// GenerateInvoiceNumber
// 生成发票码（Generate invoice number）
// 文档：https://developer.paypal.com/docs/api/invoicing/v2/#invoicing_generate-next-invoice-number
func (c *Client) GenerateInvoiceNumber(ctx context.Context, pl paypay.Payload) (res *entity.GenerateInvoiceNumberRes, err error) {
	method := GenerateInvoiceNumber
	c.EmptyChecker = method.Checker
	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{}), nil, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.GenerateInvoiceNumberRes{EmptyRes: emptyRes}
	res.Response = new(entity.InvoiceNumber)
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, res.Response); err != nil {
		return res, err
	}
	return res, nil
}

// ShowInvoiceDetail
// 查看发票详情（Show invoice details）
// 文档：https://developer.paypal.com/docs/api/invoicing/v2/#invoices_get
func (c *Client) ShowInvoiceDetail(ctx context.Context, invoiceId string, pl paypay.Payload) (res *entity.ShowInvoiceDetailRes, err error) {
	method := ShowInvoiceDetail
	c.EmptyChecker = method.Checker
	if invoiceId == pkg.NULL {
		return nil, pkg.ErrPaypalMissingInvoiceId
	}
	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"invoice_id": invoiceId,
	}), nil, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.ShowInvoiceDetailRes{EmptyRes: emptyRes}
	res.Response = new(entity.Invoice)
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, res.Response); err != nil {
		return res, err
	}
	return res, nil
}

// FullyUpdateInvoice
// 更新发票（Fully update invoice）
// 文档：https://developer.paypal.com/docs/api/invoicing/v2/#invoices_update
func (c *Client) FullyUpdateInvoice(ctx context.Context, invoiceId string, query, pl paypay.Payload) (res *entity.FullyUpdateInvoiceRes, err error) {
	method := FullyUpdateInvoice
	c.EmptyChecker = method.Checker
	if invoiceId == pkg.NULL {
		return nil, pkg.ErrPaypalMissingInvoiceId
	}
	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"invoice_id": invoiceId,
		"params":     query.EncodeURLParams(),
	}), pl, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.FullyUpdateInvoiceRes{EmptyRes: emptyRes}
	res.Response = new(entity.Invoice)
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, res.Response); err != nil {
		return res, err
	}
	return res, nil
}

// DeleteInvoice
// 删除发票（Delete invoice）
// 文档：https://developer.paypal.com/docs/api/invoicing/v2/#invoices_delete
func (c *Client) DeleteInvoice(ctx context.Context, invoiceId string) (res *entity.DeleteInvoiceRes, err error) {
	method := DeleteInvoice
	c.EmptyChecker = method.Checker
	if invoiceId == pkg.NULL {
		return nil, pkg.ErrPaypalMissingInvoiceId
	}

	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"invoice_id": invoiceId,
	}), nil, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.DeleteInvoiceRes{EmptyRes: emptyRes}
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, nil); err != nil {
		return res, err
	}
	return res, nil
}

// SearchInvoice
// 查询发票（Search for invoices）
// 文档：https://developer.paypal.com/docs/api/invoicing/v2/#invoices_search-invoices
func (c *Client) SearchInvoice(ctx context.Context, query, pl paypay.Payload) (res *entity.SearchInvoiceRes, err error) {
	method := SearchInvoice
	c.EmptyChecker = method.Checker
	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"params": query.EncodeURLParams(),
	}), pl, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.SearchInvoiceRes{EmptyRes: emptyRes}
	res.Response = new(entity.SearchInvoice)
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, res.Response); err != nil {
		return res, err
	}
	return res, nil
}

// ListInvoiceTemplate
// 发票模板列表（List templates）
// 文档：https://developer.paypal.com/docs/api/invoicing/v2/#templates_list
func (c *Client) ListInvoiceTemplate(ctx context.Context, query paypay.Payload) (res *entity.ListInvoiceTemplateRes, err error) {
	method := ListInvoiceTemplate
	c.EmptyChecker = method.Checker
	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"params": query.EncodeURLParams(),
	}), query, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.ListInvoiceTemplateRes{EmptyRes: emptyRes}
	res.Response = new(entity.InvoiceTemplate)
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, res.Response); err != nil {
		return res, err
	}
	return res, nil
}

// CreateInvoiceTemplate
// 创建发票模板（Create template）
// 文档：https://developer.paypal.com/docs/api/invoicing/v2/#templates_create
func (c *Client) CreateInvoiceTemplate(ctx context.Context, pl paypay.Payload) (res *entity.CreateInvoiceTemplateRes, err error) {
	method := CreateInvoiceTemplate
	c.EmptyChecker = method.Checker
	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, nil), pl, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.CreateInvoiceTemplateRes{EmptyRes: emptyRes}
	res.Response = new(entity.Template)
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, res.Response); err != nil {
		return res, err
	}
	return res, nil
}

// ShowTemplateDetails
// 查看模板详情（Show template details）
// 文档：https://developer.paypal.com/docs/api/invoicing/v2/#templates_get
func (c *Client) ShowTemplateDetails(ctx context.Context, templateId string, _ paypay.Payload) (res *entity.ShowTemplateDetailsRes, err error) {
	method := ShowTemplateDetails
	c.EmptyChecker = method.Checker
	if templateId == pkg.NULL {
		return nil, pkg.ErrPaypalMissingTemplateId
	}
	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"template_id": templateId,
	}), nil, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.ShowTemplateDetailsRes{EmptyRes: emptyRes}
	res.Response = new(entity.Template)
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, res.Response); err != nil {
		return res, err
	}
	return res, nil
}

// FullyUpdateInvoiceTemplate
// 更新发票模板（Fully update template）
// 文档：https://developer.paypal.com/docs/api/invoicing/v2/#templates_update
func (c *Client) FullyUpdateInvoiceTemplate(ctx context.Context, templateId string, pl paypay.Payload) (res *entity.FullyUpdateInvoiceTemplateRes, err error) {
	method := FullyUpdateInvoiceTemplate
	c.EmptyChecker = method.Checker
	if templateId == pkg.NULL {
		return nil, pkg.ErrPaypalMissingTemplateId
	}
	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"template_id": templateId,
	}), pl, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.FullyUpdateInvoiceTemplateRes{EmptyRes: emptyRes}
	res.Response = new(entity.Template)
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, res.Response); err != nil {
		return res, err
	}
	return res, nil
}

// DeleteInvoiceTemplate
// 删除发票模板（Delete template）
// 文档：https://developer.paypal.com/docs/api/invoicing/v2/#invoices_delete
func (c *Client) DeleteInvoiceTemplate(ctx context.Context, templateId string) (res *entity.DeleteInvoiceTemplateRes, err error) {
	method := DeleteInvoiceTemplate
	c.EmptyChecker = method.Checker
	if templateId == pkg.NULL {
		return nil, pkg.ErrPaypalMissingTemplateId
	}

	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"template_id": templateId,
	}), nil, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.DeleteInvoiceTemplateRes{EmptyRes: emptyRes}
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, nil); err != nil {
		return res, err
	}
	return res, nil
}
