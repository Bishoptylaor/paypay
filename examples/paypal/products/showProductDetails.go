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
 @Time    : 2024/10/14 -- 17:01
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: showProductDetails.go
*/

package products

import (
	"context"
	"encoding/json"
	"github.com/Bishoptylaor/paypay/examples"
	"github.com/Bishoptylaor/paypay/paypal"
	"github.com/Bishoptylaor/paypay/paypal/consts"
	"github.com/Bishoptylaor/paypay/pkg/xlog"
)

func ShowProductDetails(ctx context.Context, client *paypal.Client) {
	res, err := client.ShowProductDetails(ctx, "PROD-97K41316KK798123J")
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

	examples.Equal(func(bs []byte) map[string]interface{} {
		var _json map[string]interface{}
		_ = json.Unmarshal(bs, &_json)
		// different url domain
		// api.sandbox.paypal.com vs api-m.paypal.com
		delete(_json, "links")
		return _json
	}, func(f func([]byte) map[string]interface{}) map[string]interface{} {
		b, err := json.Marshal(res.Response)
		if err != nil {
			xlog.Error(err)
		}
		return f(b)
	}, func(f func([]byte) map[string]interface{}) map[string]interface{} {
		return f([]byte(showProductDetailsExample))
	})
}

var showProductDetailsExample = `{
  "id": "PROD-97K41316KK798123J",
  "name": "Video Streaming Service",
  "description": "Video streaming service",
  "type": "SERVICE",
  "category": "SOFTWARE",
  "image_url": "https://example.com/streaming.jpg",
  "home_url": "https://example.com/home",
  "create_time": "2024-10-14T08:34:47Z",
  "update_time": "2024-10-14T08:34:47Z",
  "links": [
    {
      "href": "https://api-m.paypal.com/v1/catalogs/products/PROD-97K41316KK798123J ",
      "rel": "self",
      "method": "GET"
    },
    {
      "href": "https://api-m.paypal.com/v1/catalogs/products/PROD-97K41316KK798123J ",
      "rel": "edit",
      "method": "PATCH"
    }
  ]
}
`
