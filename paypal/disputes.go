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

package paypal

import (
	"context"
	"github.com/Bishoptylaor/paypay"
	"github.com/Bishoptylaor/paypay/paypal/consts"
	"github.com/Bishoptylaor/paypay/paypal/entity"
	"github.com/Bishoptylaor/paypay/pkg"
	"github.com/Bishoptylaor/paypay/pkg/xnet/xhttp"
	"github.com/pkg/errors"
)

// EscalateDisputeToClaim
// 将投诉升级为索赔（Escalate dispute to claim）
// 文档：https://developer.paypal.com/docs/api/customer-disputes/v1/#disputes_escalate
// make sure dispute lifecycle must be INQUIRY.
func (c *Client) EscalateDisputeToClaim(ctx context.Context, disputeId string, pl paypay.Payload) (res *entity.EscalateDisputeToClaimRes, err error) {
	method := EscalateDisputeToClaim
	c.EmptyChecker = method.Checker
	if disputeId == pkg.NULL {
		return nil, pkg.ErrPaypalMissingDisputeId
	}

	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"dispute_id": disputeId,
	}), pl, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.EscalateDisputeToClaimRes{EmptyRes: emptyRes}
	res.Response = new(entity.V1Links)
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, res.Response); err != nil {
		return res, err
	}
	return res, nil
}

// AcceptOffer2Resolve
// 客户接受商家的解决方案，结束争议（Accept offer to resolve dispute）
// 文档：https://developer.paypal.com/docs/api/customer-disputes/v1/#disputes_accept-offer
func (c *Client) AcceptOffer2Resolve(ctx context.Context, disputeId string, pl paypay.Payload) (res *entity.AcceptOffer2ResolveRes, err error) {
	method := AcceptOffer2Resolve
	if disputeId == pkg.NULL {
		return nil, pkg.ErrPaypalMissingDisputeId
	}

	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"dispute_id": disputeId,
	}), pl, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.AcceptOffer2ResolveRes{EmptyRes: emptyRes}
	res.Response = new(entity.V1Links)
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, res.Response); err != nil {
		return res, err
	}
	return res, nil
}

// ListDisputes
// 获取争议列表（List disputes）
// 文档：https://developer.paypal.com/docs/api/customer-disputes/v1/#disputes_accept-offer
func (c *Client) ListDisputes(ctx context.Context, query paypay.Payload) (res *entity.ListDisputesRes, err error) {
	method := ListDisputes

	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"params": query.EncodeURLParams(),
	}), nil, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.ListDisputesRes{EmptyRes: emptyRes}
	res.Response = new(entity.DisputeLists)
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, res.Response); err != nil {
		return res, err
	}
	return res, nil
}

// ProvideInfo4Dispute
// 提供有效信息（Provide supporting information for dispute）
// 文档：https://developer.paypal.com/docs/api/customer-disputes/v1/#disputes_provide-supporting-info
// fixme : testify
func (c *Client) ProvideInfo4Dispute(ctx context.Context, disputeId string, getFiles func() map[string]paypay.File) (res *entity.ProvideInfo4DisputeRes, err error) {
	method := ProvideInfo4Dispute
	if disputeId == pkg.NULL {
		return nil, pkg.ErrPaypalMissingDisputeId
	}

	httpRes, bs, err := UploadFilePaypal(c,
		getFiles(),
		xhttp.Req(xhttp.TypeMultipartFormData),
	)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"dispute_id": disputeId,
	}),
		nil,
		nil,
	)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.ProvideInfo4DisputeRes{EmptyRes: emptyRes}
	res.Response = new(entity.V1Links)
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, res.Response); err != nil {
		return res, err
	}
	return res, nil
}

// ShowDisputeDetails
// 获取争议详情c（Show dispute details）
// 文档：https://developer.paypal.com/docs/api/customer-disputes/v1/#disputes_get
func (c *Client) ShowDisputeDetails(ctx context.Context, disputeId string) (res *entity.ShowDisputeDetailsRes, err error) {
	method := ShowDisputeDetails
	if disputeId == pkg.NULL {
		return nil, pkg.ErrPaypalMissingDisputeId
	}

	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"dispute_id": disputeId,
	}), nil, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.ShowDisputeDetailsRes{EmptyRes: emptyRes}
	res.Response = new(entity.DisputeDetail)
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, res.Response); err != nil {
		return res, err
	}
	return res, nil
}

// PartiallyUpdateDispute
// 更新争议（Update dispute）
// 文档：https://developer.paypal.com/docs/api/customer-disputes/v1/#disputes_patch
func (c *Client) PartiallyUpdateDispute(ctx context.Context, disputeId string, patches []*entity.Patch) (res *entity.PartiallyUpdateDisputeRes, err error) {
	method := PartiallyUpdateDispute
	if disputeId == pkg.NULL {
		return nil, pkg.ErrPaypalMissingDisputeId
	}

	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"dispute_id": disputeId,
	}), nil, patches)
	if err != nil {
		return nil, errors.Wrap(err, "[PartiallyUpdateDispute] do method")
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.PartiallyUpdateDisputeRes{EmptyRes: emptyRes}
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, nil); err != nil {
		return res, errors.Wrap(err, "[PartiallyUpdateDispute]")
	}
	return res, nil
}

// DenyOffer2Resolve
// 拒绝方案（Deny offer to resolve dispute）
// 文档：https://developer.paypal.com/docs/api/customer-disputes/v1/#disputes_deny-offer
func (c *Client) DenyOffer2Resolve(ctx context.Context, disputeId string, pl paypay.Payload) (res *entity.DenyOffer2ResolveRes, err error) {
	method := DenyOffer2Resolve
	if disputeId == pkg.NULL {
		return nil, pkg.ErrPaypalMissingDisputeId
	}

	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"dispute_id": disputeId,
	}), pl, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.DenyOffer2ResolveRes{EmptyRes: emptyRes}
	res.Response = new(entity.V1Links)
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, res.Response); err != nil {
		return res, err
	}
	return res, nil
}

// MakeOffer2Resolve
// 发起方案（Make offer to resolve dispute）
// 文档：https://developer.paypal.com/docs/api/customer-disputes/v1/#disputes_make-offer
func (c *Client) MakeOffer2Resolve(ctx context.Context, disputeId string, pl paypay.Payload) (res *entity.MakeOffer2ResolveRes, err error) {
	method := MakeOffer2Resolve
	if disputeId == pkg.NULL {
		return nil, pkg.ErrPaypalMissingDisputeId
	}

	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"dispute_id": disputeId,
	}), pl, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.MakeOffer2ResolveRes{EmptyRes: emptyRes}
	res.Response = new(entity.V1Links)
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, res.Response); err != nil {
		return res, err
	}
	return res, nil
}

// AppealDispute
// 上诉（Appeal dispute）
// 文档：https://developer.paypal.com/docs/api/customer-disputes/v1/#disputes_appeal
// fixme : testify
func (c *Client) AppealDispute(ctx context.Context, disputeId string, getFiles func() map[string]paypay.File) (res *entity.AppealDisputeRes, err error) {
	method := AppealDispute
	if disputeId == pkg.NULL {
		return nil, pkg.ErrPaypalMissingDisputeId
	}

	httpRes, bs, err := UploadFilePaypal(c,
		getFiles(),
		xhttp.Req(xhttp.TypeMultipartFormData),
	)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"dispute_id": disputeId,
	}),
		nil,
		nil,
	)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.AppealDisputeRes{EmptyRes: emptyRes}
	res.Response = new(entity.V1Links)
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, res.Response); err != nil {
		return res, err
	}
	return res, nil
}

// ProvideEvidence
// 上诉（Provide evidence）
// 文档：https://developer.paypal.com/docs/api/customer-disputes/v1/#disputes_provide-evidence
// fixme : testify
func (c *Client) ProvideEvidence(ctx context.Context, disputeId string, getFiles func() map[string]paypay.File) (res *entity.ProvideEvidenceRes, err error) {
	method := ProvideEvidence
	if disputeId == pkg.NULL {
		return nil, pkg.ErrPaypalMissingDisputeId
	}

	httpRes, bs, err := UploadFilePaypal(c,
		getFiles(),
		xhttp.Req(xhttp.TypeMultipartFormData),
	)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"dispute_id": disputeId,
	}),
		nil,
		nil,
	)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.ProvideEvidenceRes{EmptyRes: emptyRes}
	res.Response = new(entity.V1Links)
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, res.Response); err != nil {
		return res, err
	}
	return res, nil
}

// AckReturnedItem
// ack 退回的商品（Acknowledge returned item）
// 文档：https://developer.paypal.com/docs/api/customer-disputes/v1/#disputes_acknowledge-return-item
// fixme : testify
func (c *Client) AckReturnedItem(ctx context.Context, disputeId string, pl paypay.Payload) (res *entity.AckReturnedItemRes, err error) {
	method := AckReturnedItem
	if disputeId == pkg.NULL {
		return nil, pkg.ErrPaypalMissingDisputeId
	}

	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"dispute_id": disputeId,
	}), pl, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.AckReturnedItemRes{EmptyRes: emptyRes}
	res.Response = new(entity.V1Links)
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, res.Response); err != nil {
		return res, err
	}
	return res, nil
}

// NotifyDispute2ThirdParty
// 向第三方发送相关内容（Send message about dispute to other party）
// 文档：https://developer.paypal.com/docs/api/customer-disputes/v1/#disputes_send-message
// fixme : testify
func (c *Client) NotifyDispute2ThirdParty(ctx context.Context, disputeId string, getFiles func() map[string]paypay.File) (res *entity.NotifyDispute2ThirdPartyRes, err error) {
	method := NotifyDispute2ThirdParty
	if disputeId == pkg.NULL {
		return nil, pkg.ErrPaypalMissingDisputeId
	}

	httpRes, bs, err := UploadFilePaypal(c,
		func() map[string]paypay.File {
			_map := getFiles()
			if _, ok := _map["message_document"]; ok {
				return _map
			}
			return make(map[string]paypay.File)
		}(),
		xhttp.Req(xhttp.TypeMultipartFormData),
	)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"dispute_id": disputeId,
	}),
		nil,
		nil,
	)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.NotifyDispute2ThirdPartyRes{EmptyRes: emptyRes}
	res.Response = new(entity.V1Links)
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, res.Response); err != nil {
		return res, err
	}
	return res, nil
}

// AcceptClaim
// 接受索赔（Accept claim）
// 文档：https://developer.paypal.com/docs/api/customer-disputes/v1/#disputes_accept-claim
// Bug：原文有错误，文档中显示上传文件，例子中是正常的 payload
// fixme : testify
func (c *Client) AcceptClaim(ctx context.Context, disputeId string, pl paypay.Payload) (res *entity.AcceptClaimRes, err error) {
	method := AcceptClaim
	if disputeId == pkg.NULL {
		return nil, pkg.ErrPaypalMissingDisputeId
	}

	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"dispute_id": disputeId,
	}), pl, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.AcceptClaimRes{EmptyRes: emptyRes}
	res.Response = new(entity.V1Links)
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, res.Response); err != nil {
		return res, err
	}
	return res, nil
}
