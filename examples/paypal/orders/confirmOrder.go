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
 @Description: confirmOrder.go
*/

package orders

import (
	"context"
	"github.com/Bishoptylaor/paypay"
	"github.com/Bishoptylaor/paypay/paypal"
	"github.com/Bishoptylaor/paypay/paypal/consts"
	"github.com/Bishoptylaor/paypay/pkg/xlog"
)

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
