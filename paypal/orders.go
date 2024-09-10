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
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/orders/v2/#orders_create
func (c *Client) CreateOrder(ctx context.Context, pl paypay.Payload) (res *entity.CreateOrderRes, err error) {
	method := consts.CreateOrder
	c.EmptyChecker = func(uri string) []paypay.Ruler {
		_map := map[string][]paypay.Ruler{
			method.Uri: []paypay.Ruler{
				paypay.NewRuler("purchase_units",
					`purchase_units != nil && len(purchase_units) <= 10 &&
all(purchase_units, {.Amount != nil}) `,
					"purchase_units 最多一次性传入10个",
				),
				paypay.NewRuler("intent", `intent in ["CAPTURE", "AUTHORIZE"]`, ""),
			},
		}
		if rulers, ok := _map[uri]; ok {
			return rulers
		} else {
			return []paypay.Ruler{}
		}
	}

	httpRes, bs, err := c.doPayPalPost(ctx, pl, consts.CreateOrder.Uri, c.GenUrl(ctx, nil))
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
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/orders/v2/#orders_get
func (c *Client) ShowOrderDetails(ctx context.Context, orderId string, pl paypay.Payload) (res *entity.ShowOrderDetailsRes, err error) {
	method := consts.ShowOrderDetails
	if orderId == pkg.NULL {
		return nil, pkg.ErrMissingInitParams
	}

	httpRes, bs, err := c.doPayPalGet(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"id":     orderId,
		"params": pl.EncodeURLParams(),
	}))
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
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/orders/v2/#orders_patch
func (c *Client) UpdateOrder(ctx context.Context, orderId string, patches []*entity.Patch) (res *entity.UpdateOrderRes, err error) {
	method := consts.UpdateOrder
	if orderId == pkg.NULL {
		return nil, pkg.ErrMissingInitParams
	}

	httpRes, bs, err := c.doPayPalPatch(ctx, patches, method.Uri, c.GenUrl(ctx, map[string]string{
		"id": orderId,
	}))
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
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/orders/v2/#orders_confirm
func (c *Client) ConfirmOrder(ctx context.Context, orderId string, pl paypay.Payload) (res *entity.ConfirmOrderRes, err error) {
	method := consts.ConfirmOrder
	c.EmptyChecker = func(uri string) []paypay.Ruler {
		_map := map[string][]paypay.Ruler{
			method.Uri: []paypay.Ruler{
				paypay.NewRuler("payment_source", `payment_source != nil`, "payment_source 不为空"),
			},
		}
		if rulers, ok := _map[uri]; ok {
			return rulers
		} else {
			return []paypay.Ruler{}
		}
	}
	if orderId == pkg.NULL {
		return nil, pkg.ErrMissingInitParams
	}

	httpRes, bs, err := c.doPayPalPost(ctx, pl, method.Uri, c.GenUrl(ctx, map[string]string{
		"id": orderId,
	}))
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
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/orders/v2/#orders_authorize
func (c *Client) AuthorizeOrder(ctx context.Context, orderId string, pl paypay.Payload) (res *entity.AuthorizeOrderRes, err error) {
	method := consts.AuthorizeOrder
	c.EmptyChecker = func(uri string) []paypay.Ruler {
		_map := map[string][]paypay.Ruler{
			method.Uri: []paypay.Ruler{
				paypay.NewRuler("payment_source", `payment_source != nil`, "payment_source 不为空"),
			},
		}
		if rulers, ok := _map[uri]; ok {
			return rulers
		} else {
			return []paypay.Ruler{}
		}
	}
	if orderId == pkg.NULL {
		return nil, pkg.ErrMissingInitParams
	}

	httpRes, bs, err := c.doPayPalPost(ctx, pl, method.Uri, c.GenUrl(ctx, map[string]string{
		"id": orderId,
	}))
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
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/orders/v2/#orders_capture
func (c *Client) CaptureOrder(ctx context.Context, orderId string, pl paypay.Payload) (res *entity.CaptureOrderRes, err error) {
	method := consts.CaptureOrder
	c.EmptyChecker = func(uri string) []paypay.Ruler {
		_map := map[string][]paypay.Ruler{
			method.Uri: []paypay.Ruler{
				paypay.NewRuler("payment_source", `payment_source != nil`, "payment_source 不为空"),
			},
		}
		if rulers, ok := _map[uri]; ok {
			return rulers
		} else {
			return []paypay.Ruler{}
		}
	}
	if orderId == pkg.NULL {
		return nil, pkg.ErrMissingInitParams
	}

	httpRes, bs, err := c.doPayPalPost(ctx, pl, method.Uri, c.GenUrl(ctx, map[string]string{
		"id": orderId,
	}))
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
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/orders/v2/#orders_track_create
func (c *Client) AddTrackerForOrder(ctx context.Context, orderId string, pl paypay.Payload) (res *entity.AddTrackerForOrderRes, err error) {
	method := consts.AddTracking4Order
	c.EmptyChecker = func(uri string) []paypay.Ruler {
		_map := map[string][]paypay.Ruler{
			method.Uri: []paypay.Ruler{
				paypay.NewRuler("tracking_number", `tracking_number != nil`, "运单号不为空"),
				paypay.NewRuler("carrier", `carrier != nil`, "承运机构不为空"),
				paypay.NewRuler("capture_id", `capture_id != nil`, "capture_id 不为空"),
			},
		}
		if rulers, ok := _map[uri]; ok {
			return rulers
		} else {
			return []paypay.Ruler{}
		}
	}
	if orderId == pkg.NULL {
		return nil, pkg.ErrMissingInitParams
	}

	httpRes, bs, err := c.doPayPalPost(ctx, pl, method.Uri, c.GenUrl(ctx, map[string]string{
		"id": orderId,
	}))
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
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/orders/v2/#orders_trackers_patch
func (c *Client) TrackersOfOrder(ctx context.Context, orderId, trackerId string, patches []*entity.Patch) (res *entity.TrackersOfOrderRes, err error) {
	method := consts.AddTracking4Order
	if orderId == pkg.NULL {
		return nil, pkg.ErrMissingInitParams
	}

	httpRes, bs, err := c.doPayPalPatch(ctx, patches, method.Uri, c.GenUrl(ctx, map[string]string{
		"id":         orderId,
		"tracker_id": trackerId,
	}))
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
