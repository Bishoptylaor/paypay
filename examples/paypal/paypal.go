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
 @Time    : 2024/9/9 -- 16:44
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: paypel.go
*/

package main

import (
	"context"
	"github.com/Bishoptylaor/paypay/examples/paypal/orders"
	"github.com/Bishoptylaor/paypay/examples/paypal/products"
	"github.com/Bishoptylaor/paypay/examples/paypal/singlecall"
	"github.com/Bishoptylaor/paypay/paypal"
	"github.com/Bishoptylaor/paypay/pkg/xlog"
)

var (
	ctx    context.Context
	client *paypal.Client
)

func init() {
	var err error
	ctx = context.Background()
	client, err = paypal.NewClient(ctx,
		paypal.ClientID("AbSkAmeRl40PdAk7LTD-dKpu-I1kdTV6VYgywViwv7RKYjFmIeFDFeMBMoK2_uXadlrLLgsWjmTJV-xH"),
		paypal.Secret("EJ8qu25jspPSLe65hwslbCWLUAoEm0wE9lOiHhSMikpfBj1-lRIAOHzC7OxuLb0lgLj5XuvOWTwBVOHp"),
	)
	if err != nil {
		xlog.Error(err)
		return
	}
}

func main() {
	// RunOrderExamples()
	// RunInvoiceExamples()
	// RunProductsExamples()
	SingleCallUsage()
}

func RunProductsExamples() {
	// products.CreateProduct(ctx, client)
	products.ListProducts(ctx, client)
	// products.ShowProductDetails(ctx, client)

	// products.UpdateProduct(ctx, client)
	// products.ListProducts(ctx, client)
}

func RunOrderExamples() {
	// CreateOrder(ctx, client)
	orders.ShowOrderDetails(ctx, client)
	// UpdateOrder(ctx, client)
	// ConfirmOrder(ctx, client)
	// 登录支付
	// AuthorizeOrder(ctx, client)
}

func SingleCallUsage() {
	singlecall.ListProducts(ctx, client)
}
