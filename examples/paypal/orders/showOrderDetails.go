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
 @Description: showOrderDetails.go
*/

package orders

import (
	"context"
	"github.com/Bishoptylaor/paypay/paypal"
	"github.com/Bishoptylaor/paypay/paypal/consts"
	"github.com/Bishoptylaor/paypay/pkg/xlog"
)

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
}

var showOrderDetailsExample = `
{
    "id": "4HD07929GA974304A",
    "intent": "CAPTURE",
    "status": "CREATED",
    "purchase_units": [{
        "reference_id": "HqSGj0ZW89jkzeUJ",
        "amount": {
            "currency_code": "USD",
            "value": "8.00"
        },
        "payee": {
            "email_address": "sb-fylie32523003@business.example.com",
            "merchant_id": "M8MGJQBRHW74J",
            "display_data": {
                "brand_name": "paypay"
            }
        }
    }],
    "create_time": "2024-09-11T09:13:37Z",
    "links": [{
        "href": "https://api.sandbox.paypal.com/v2/checkout/orders/4HD07929GA974304A",
        "rel": "self",
        "method": "GET"
    }, {
        "href": "https://www.sandbox.paypal.com/checkoutnow?token=4HD07929GA974304A",
        "rel": "approve",
        "method": "GET"
    }, {
        "href": "https://api.sandbox.paypal.com/v2/checkout/orders/4HD07929GA974304A",
        "rel": "update",
        "method": "PATCH"
    }, {
        "href": "https://api.sandbox.paypal.com/v2/checkout/orders/4HD07929GA974304A/capture",
        "rel": "capture",
        "method": "POST"
    }]
}
`
