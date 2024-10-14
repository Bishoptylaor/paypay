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
 @Time    : 2024/10/14 -- 17:36
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: authorizeOrder.go
*/

package orders

import (
	"context"
	"github.com/Bishoptylaor/paypay/paypal"
	"github.com/Bishoptylaor/paypay/paypal/consts"
	"github.com/Bishoptylaor/paypay/pkg/xlog"
)

func AuthorizeOrder(ctx context.Context, client *paypal.Client) {
	res, err := client.AuthorizeOrder(ctx, "4HD07929GA974304A", nil)
	if err != nil {
		xlog.Error(err)
		return
	}

	if res.Code != consts.Success {
		xlog.Infof(ctx, "res.Code: %+v", res.Code)
		xlog.Infof(ctx, "res.Error: %+v", res.Error)
		xlog.Infof(ctx, "res.ErrorResponse: %+v", res.ErrorResponse)
		return
	}
	xlog.Infof(ctx, "res.Response: %+v", res.Response)
	for _, v := range res.Response.PurchaseUnits {
		xlog.Infof(ctx, "res.Response.PurchaseUnit.ReferenceId: %+v", v.ReferenceId)
		xlog.Infof(ctx, "res.Response.PurchaseUnit.Amount: %+v", v.Amount)
		if v.Shipping != nil && v.Shipping.Address != nil {
			xlog.Infof(ctx, "res.Response.PurchaseUnit.Shipping.Address: %+v", v.Shipping.Address)
		}
		xlog.Infof(ctx, "res.Response.PurchaseUnit.Description: %+v", v.Description)
		if v.Payments != nil && v.Payments.Authorizations != nil {
			xlog.Infof(ctx, "res.Response.PurchaseUnit.Payments.Authorizations: %+v", v.Payments.Authorizations)
		}
	}
	for _, v := range res.Response.Links {
		xlog.Infof(ctx, "res.Response.Links: %+v", v)
	}
}
