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

package singlecall

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
		Set("total_required", "false")

	res := new(entity.ListProductsRes)
	res.Response = new(entity.ProductList)
	err := client.CustomSingleCall(ctx,
		paypal.ListProducts,
		func() map[string]string {
			return map[string]string{}
		},
		nil,
		query,
		nil,
		res,
		res.Response,
		paypal.Headers(map[string]string{
			"Prefer": "return=representation",
		}),
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
