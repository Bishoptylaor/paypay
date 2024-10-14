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
 @Time    : 2024/9/11 -- 17:03
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: orders.go
*/

package main

import (
	"context"
	"github.com/Bishoptylaor/paypay"
	"github.com/Bishoptylaor/paypay/paypal"
	"github.com/Bishoptylaor/paypay/paypal/consts"
	"github.com/Bishoptylaor/paypay/paypal/entity"
	"github.com/Bishoptylaor/paypay/pkg/xlog"
	"github.com/Bishoptylaor/paypay/pkg/xutils"
)

func RunOrderExamples() {
	// CreateOrder(ctx, client)
	ShowOrderDetails(ctx, client)
	// UpdateOrder(ctx, client)
	// ConfirmOrder(ctx, client)
	// 登录支付
	// AuthorizeOrder(ctx, client)
}

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

func ShowOrderDetails(ctx context.Context, client *paypal.Client) {
	res, err := client.ShowOrderDetails(ctx, "4HD07929GA974304A", nil)
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
	/*
		{"id":"4HD07929GA974304A","intent":"CAPTURE","status":"CREATED","purchase_units":[{"reference_id":"HqSGj0ZW89jkzeUJ","amount":{"currency_code":"USD","value":"8.00"},"payee":{"email_address":"sb-fylie32523003@business.example.com","merchant_id":"M8MGJQBRHW74J","display_data":{"brand_name":"paypay"}}}],"create_time":"2024-09-11T09:13:37Z","links":[{"href":"https://api.sandbox.paypal.com/v2/checkout/orders/4HD07929GA974304A","rel":"self","method":"GET"},{"href":"https://www.sandbox.paypal.com/checkoutnow?token=4HD07929GA974304A","rel":"approve","method":"GET"},{"href":"https://api.sandbox.paypal.com/v2/checkout/orders/4HD07929GA974304A","rel":"update","method":"PATCH"},{"href":"https://api.sandbox.paypal.com/v2/checkout/orders/4HD07929GA974304A/capture","rel":"capture","method":"POST"}]}
	*/
}

func UpdateOrder(ctx context.Context, client *paypal.Client) {
	ps := []*entity.Patch{
		&entity.Patch{
			Op:   "replace",
			Path: "/purchase_units/@reference_id=='HqSGj0ZW89jkzeUJ'/shipping/address", // reference_id is yourself set when create order
			Value: &entity.Address{
				AddressLine1: "321 Townsend St",
				AddressLine2: "Floor 7",
				AdminArea1:   "San Francisco",
				AdminArea2:   "CA",
				PostalCode:   "94107",
				CountryCode:  "US",
			},
		},
		// &entity.Patch{
		// 	Op:    "add",
		// 	Path:  "/purchase_units/@reference_id=='HqSGj0ZW89jkzeUJ'/description",
		// 	Value: "I am patch info",
		// },
	}

	res, err := client.UpdateOrder(ctx, "4HD07929GA974304A", ps)
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
	xlog.Infof(ctx, "res: %+v", res)
}

func ConfirmOrder(ctx context.Context, client *paypal.Client) {
	pl := make(paypay.Payload)
	pl.Set("payment_source", func(l1 paypay.Payload) {
		l1.Set("paypal", func(l2 paypay.Payload) {
			l2.Set("name", func(l3 paypay.Payload) {
				l3.Set("given_name", "John")
				l3.Set("surname", "Doe")
			})
			l2.Set("email_address", "sb-xepyh32505727@personal.example.com")
			l2.Set("experience_context", func(l3 paypay.Payload) {
				l3.Set("payment_method_preference", "IMMEDIATE_PAYMENT_REQUIRED")
				l3.Set("brand_name", "EXAMPLE INC")
				l3.Set("locale", "en-US")
				l3.Set("landing_page", "LOGIN")
				l3.Set("shipping_preference", "SET_PROVIDED_ADDRESS")
				l3.Set("user_action", "PAY_NOW")
				l3.Set("return_url", "https://example.com/returnUrl")
				l3.Set("cancel_url", "https://example.com/cancelUrl")
			})
		})
	})
	res, err := client.ConfirmOrder(ctx, "4HD07929GA974304A", pl)
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
}

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
