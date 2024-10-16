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
 @Time    : 2024/10/15 -- 09:57
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: call.go
*/

package cco

import (
	"context"
	"github.com/Bishoptylaor/paypay"
	"github.com/Bishoptylaor/paypay/paypal"
	"github.com/Bishoptylaor/paypay/paypal/consts"
	"github.com/Bishoptylaor/paypay/paypal/entity"
	"github.com/Bishoptylaor/paypay/pkg/xlog"
)

func ListProducts(ctx context.Context, client *paypal.Client) {
	query := make(paypay.Payload)
	query.Set("page_size", "10").
		Set("page", "1"). // starts from 1
		Set("total_required", "true")

	emptyRes := entity.EmptyRes{Code: consts.Success}
	res := &entity.ListProductsRes{EmptyRes: emptyRes}
	res.Response = new(entity.ProductList)
	err := client.CustomCallOnce(ctx,
		// the method you want to use
		paypal.ListProducts,
		// params like order_id || dispute_id.
		func() map[string]string {
			return map[string]string{}
		},
		// payload
		nil,
		// query params
		query,
		// update patches
		nil,
		emptyRes,
		res.Response,
		paypal.Headers(map[string]string{
			"Prefer": "return=representation",
		}),
		// if you need a call with new Paypal Account , you can use like this ↓
		// paypal.NewToken(ctx, "new client id", "new client secret"),
	)
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
	for _, product := range res.Response.Products {
		xlog.Infof(ctx, "product: %+v", product)
	}

	for _, v := range res.Response.Links {
		xlog.Infof(ctx, "res.Response.Links: %+v", v)
	}

}

func ShowProductDetails(ctx context.Context, client *paypal.Client) {
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res := &entity.ShowProductDetailsRes{EmptyRes: emptyRes}
	res.Response = new(entity.ProductDetail)
	err := client.CustomCallOnce(ctx,
		// the method you want to use
		paypal.ShowProductDetails,
		// params like order_id || dispute_id.
		func() map[string]string {
			return map[string]string{
				paypal.ProductId.String(): "PROD-97K41316KK798123J",
			}
		},
		// payload
		nil,
		// query params
		nil,
		// update patches
		nil,
		emptyRes,
		res.Response,
		paypal.Headers(map[string]string{
			"Prefer": "return=representation",
		}),
		// if you need a call with new Paypal Account , you can use like this ↓
		paypal.NewToken(ctx, "new client id", "new client secret"),
	)
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
