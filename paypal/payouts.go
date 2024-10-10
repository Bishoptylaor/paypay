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

package paypal

import (
	"context"
	"github.com/Bishoptylaor/paypay"
	"github.com/Bishoptylaor/paypay/paypal/consts"
	"github.com/Bishoptylaor/paypay/paypal/entity"
	"github.com/Bishoptylaor/paypay/pkg"
)

// CreateBatchPayout
// 创建批量支付（Create batch payout）
// 文档：https://developer.paypal.com/docs/api/payments.payouts-batch/v1/#payouts_post
func (c *Client) CreateBatchPayout(ctx context.Context, pl paypay.Payload) (res *entity.CreateBatchPayoutRes, err error) {
	method := CreateBatchPayout
	c.EmptyChecker = method.Checker

	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, nil), pl, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.CreateBatchPayoutRes{EmptyRes: emptyRes}
	res.Response = new(entity.BatchPayout)
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, res.Response); err != nil {
		return res, err
	}
	return res, nil
}

// ShowPayoutBatchDetail
// 查看批量支付详情（Show payout batch details）
// 文档：https://developer.paypal.com/docs/api/payments.payouts-batch/v1/#payouts_get
func (c *Client) ShowPayoutBatchDetail(ctx context.Context, batchId string, pl paypay.Payload) (res *entity.ShowPayoutBatchDetailRes, err error) {
	method := ShowPayoutBatchDetail
	if batchId == pkg.NULL {
		return nil, pkg.ErrPaypalMissingPayoutBatchId
	}

	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"payout_batch_id": batchId,
		"params":          pl.EncodeURLParams(),
	}), nil, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.ShowPayoutBatchDetailRes{EmptyRes: emptyRes}
	res.Response = new(entity.PayoutBatchDetail)
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, res.Response); err != nil {
		return res, err
	}
	return res, nil
}

// ShowPayoutItemDetail
// 查看支付详情（Show payout detail）
// 文档：https://developer.paypal.com/docs/api/payments.payouts-batch/v1/#payouts-item_get
func (c *Client) ShowPayoutItemDetail(ctx context.Context, payoutItemId string) (res *entity.ShowPayoutItemDetailRes, err error) {
	method := ShowPayoutItemDetail
	if payoutItemId == pkg.NULL {
		return nil, pkg.ErrPaypalMissingPayoutItemId
	}

	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"payout_item_id": payoutItemId,
	}), nil, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.ShowPayoutItemDetailRes{EmptyRes: emptyRes}
	res.Response = new(entity.PayoutItemDetail)
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, res.Response); err != nil {
		return res, err
	}
	return res, nil
}

// CancelUnclaimedPayoutItem
// 取消批量支付中收款人无PayPal账号的项目（Cancel Unclaimed Payout Item）主动取消，超过30天 Paypal 系统自动取消
// 文档：https://developer.paypal.com/docs/api/payments.payouts-batch/v1/#payouts-item_cancel
func (c *Client) CancelUnclaimedPayoutItem(ctx context.Context, payoutItemId string) (res *entity.CancelUnclaimedPayoutItemRes, err error) {
	method := CancelUnclaimedPayoutItem
	if payoutItemId == pkg.NULL {
		return nil, pkg.ErrPaypalMissingPayoutItemId
	}

	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"payout_item_id": payoutItemId,
	}), nil, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.CancelUnclaimedPayoutItemRes{EmptyRes: emptyRes}
	res.Response = new(entity.PayoutItemDetail)
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, res.Response); err != nil {
		return res, err
	}
	return res, nil
}
