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
 @Time    : 2024/9/3 -- 18:31
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: orders.go
*/

package paypal

import (
	"context"
	"github.com/Bishoptylaor/paypay"
	"github.com/Bishoptylaor/paypay/paypal/consts"
	"github.com/Bishoptylaor/paypay/paypal/entity"
	"github.com/Bishoptylaor/paypay/pkg"
)

// CreateOrder
// 创建订单（Create order）
// 文档：https://developer.paypal.com/docs/api/orders/v2/#orders_create
func (c *Client) CreateOrder(ctx context.Context, pl paypay.Payload) (res *entity.CreateOrderRes, err error) {
	method := CreateOrder
	c.EmptyChecker = method.Checker

	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, nil), pl, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.CreateOrderRes{EmptyRes: emptyRes}
	res.Response = new(entity.OrderDetail)
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, res.Response); err != nil {
		return res, err
	}
	return res, nil
}

// ShowOrderDetails
// 查看订单详情（Show order details）
// 文档：https://developer.paypal.com/docs/api/orders/v2/#orders_get
func (c *Client) ShowOrderDetails(ctx context.Context, orderId string, pl paypay.Payload) (res *entity.ShowOrderDetailsRes, err error) {
	method := ShowOrderDetails
	if orderId == pkg.NULL {
		return nil, pkg.ErrPaypalMissingOrderId
	}

	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"id":     orderId,
		"params": pl.EncodeURLParams(),
	}), nil, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.ShowOrderDetailsRes{EmptyRes: emptyRes}
	res.Response = new(entity.OrderDetail)
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, res.Response); err != nil {
		return res, err
	}
	return res, nil
}

// UpdateOrder
// 更新订单（Update order）
// 文档：https://developer.paypal.com/docs/api/orders/v2/#orders_patch
func (c *Client) UpdateOrder(ctx context.Context, orderId string, patches []*entity.Patch) (res *entity.UpdateOrderRes, err error) {
	method := UpdateOrder
	if orderId == pkg.NULL {
		return nil, pkg.ErrPaypalMissingOrderId
	}
	if len(patches) == 0 {
		return nil, pkg.ErrPaypalNothingToChange
	}

	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"id": orderId,
	}), nil, patches)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.UpdateOrderRes{EmptyRes: emptyRes}
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, new(struct{})); err != nil {
		return res, err
	}
	return res, nil
}

// ConfirmOrder
// 订单支付确认（Confirm the Order）
// 文档：https://developer.paypal.com/docs/api/orders/v2/#orders_confirm
func (c *Client) ConfirmOrder(ctx context.Context, orderId string, pl paypay.Payload) (res *entity.ConfirmOrderRes, err error) {
	method := ConfirmOrder
	c.EmptyChecker = method.Checker
	if orderId == pkg.NULL {
		return nil, pkg.ErrPaypalMissingOrderId
	}

	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"id": orderId,
	}), pl, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.ConfirmOrderRes{EmptyRes: emptyRes}
	res.Response = new(entity.OrderDetail)
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, res.Response); err != nil {
		return res, err
	}
	return res, nil
}

// AuthorizeOrder
// 订单支付确认（Authorize payment for order）
// 文档：https://developer.paypal.com/docs/api/orders/v2/#orders_authorize
func (c *Client) AuthorizeOrder(ctx context.Context, orderId string, pl paypay.Payload) (res *entity.AuthorizeOrderRes, err error) {
	method := AuthorizeOrder
	c.EmptyChecker = method.Checker
	if orderId == pkg.NULL {
		return nil, pkg.ErrPaypalMissingOrderId
	}

	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"id": orderId,
	}), pl, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.AuthorizeOrderRes{EmptyRes: emptyRes}
	res.Response = new(entity.OrderDetail)
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, res.Response); err != nil {
		return res, err
	}
	return res, nil
}

// CaptureOrder
// 订单支付确认（Capture payment for order）
// 文档：https://developer.paypal.com/docs/api/orders/v2/#orders_capture
func (c *Client) CaptureOrder(ctx context.Context, orderId string, pl paypay.Payload) (res *entity.CaptureOrderRes, err error) {
	method := CaptureOrder
	c.EmptyChecker = method.Checker
	if orderId == pkg.NULL {
		return nil, pkg.ErrPaypalMissingOrderId
	}

	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"id": orderId,
	}), pl, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.CaptureOrderRes{EmptyRes: emptyRes}
	res.Response = new(entity.OrderDetail)
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, res.Response); err != nil {
		return res, err
	}
	return res, nil
}

// AddTrackerForOrder
// 给订单添加物流信息（Add tracking information for an Order）
// 文档：https://developer.paypal.com/docs/api/orders/v2/#orders_track_create
func (c *Client) AddTrackerForOrder(ctx context.Context, orderId string, pl paypay.Payload) (res *entity.AddTrackerForOrderRes, err error) {
	method := AddTracking4Order
	c.EmptyChecker = method.Checker
	if orderId == pkg.NULL {
		return nil, pkg.ErrPaypalMissingOrderId
	}

	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"id": orderId,
	}), pl, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.AddTrackerForOrderRes{EmptyRes: emptyRes}
	res.Response = new(entity.OrderDetail)
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, res.Response); err != nil {
		return res, err
	}
	return res, nil
}

// TrackersOfOrder
// 更新或取消物流信息（Update or cancel tracking information for a PayPal order）
// 文档：https://developer.paypal.com/docs/api/orders/v2/#orders_trackers_patch
func (c *Client) TrackersOfOrder(ctx context.Context, orderId, trackerId string, patches []*entity.Patch) (res *entity.TrackersOfOrderRes, err error) {
	method := AddTracking4Order
	if orderId == pkg.NULL {
		return nil, pkg.ErrPaypalMissingOrderId
	}
	if len(patches) == 0 {
		return nil, pkg.ErrPaypalNothingToChange
	}

	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"id":         orderId,
		"tracker_id": trackerId,
	}), nil, patches)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.TrackersOfOrderRes{EmptyRes: emptyRes}
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, new(struct{})); err != nil {
		return res, err
	}
	return res, nil
}
