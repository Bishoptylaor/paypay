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
 @Time    : 2024/10/14 -- 17:34
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: createOrder.go
*/

package orders

import (
	"context"
	"github.com/Bishoptylaor/paypay"
	"github.com/Bishoptylaor/paypay/paypal"
	"github.com/Bishoptylaor/paypay/paypal/consts"
	"github.com/Bishoptylaor/paypay/paypal/entity"
	"github.com/Bishoptylaor/paypay/pkg/xlog"
	"github.com/Bishoptylaor/paypay/pkg/xutils"
)

func CreateOrder(ctx context.Context, client *paypal.Client) {
	var pus []*entity.PurchaseUnit
	var item = &entity.PurchaseUnit{
		ReferenceId: xutils.RandomString(16),
		Amount: &entity.Amount{
			CurrencyCode: "USD",
			Value:        "8",
		},
	}
	pus = append(pus, item)

	bm := make(paypay.Payload)
	bm.Set("intent", "CAPTURE").
		Set("purchase_units", pus).
		Set("application_context", func(b paypay.Payload) {
			b.Set("brand_name", "paypay").
				Set("locale", "en-PT").
				Set("return_url", "https://example.com/returnUrl").
				Set("cancel_url", "https://example.com/cancelUrl")
		})

	res, err := client.CreateOrder(ctx, bm)
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

	for _, v := range res.Response.Links {
		xlog.Infof(ctx, "res.Response.Links: %+v", v)
	}
}
