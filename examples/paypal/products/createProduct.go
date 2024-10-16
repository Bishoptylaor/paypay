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
 @Time    : 2024/10/14 -- 15:37
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: products.go
*/

package products

import (
	"context"
	"encoding/json"
	"github.com/Bishoptylaor/paypay"
	"github.com/Bishoptylaor/paypay/examples"
	"github.com/Bishoptylaor/paypay/paypal"
	"github.com/Bishoptylaor/paypay/paypal/consts"
	"github.com/Bishoptylaor/paypay/pkg/xlog"
)

func CreateProduct(ctx context.Context, client *paypal.Client) {
	pl := make(paypay.Payload)
	pl.Set("name", "Video Streaming Service").
		Set("description", "Video streaming service").
		Set("type", "SERVICE").
		Set("category", "SOFTWARE").
		Set("image_url", "https://example.com/streaming.jpg").
		Set("home_url", "https://example.com/home")

	res, err := client.CreateProduct(ctx, pl)
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
		// some information will contain abs time info || genetic ids, will delete them to make compare
		delete(_json, "links")
		delete(_json, "id")
		delete(_json, "create_time")
		delete(_json, "update_time")
		return _json
	}, func(f func(bs []byte) map[string]interface{}) map[string]interface{} {
		b, err := json.Marshal(res.Response)
		if err != nil {
			xlog.Error(err)
		}
		return f(b)
	}, func(f func(bs []byte) map[string]interface{}) map[string]interface{} {
		return f([]byte(createProductExample))
	})

}

var createProductExample = `
	{
	  "id": "PROD-1YN92718J3553841U",
	  "name": "Video Streaming Service",
	  "description": "Video streaming service",
	  "type": "SERVICE",
	  "category": "SOFTWARE",
	  "image_url": "https://example.com/streaming.jpg",
	  "home_url": "https://example.com/home",
	  "create_time": "2024-10-14T08:20:01Z",
	  "update_time": "2024-10-14T08:20:01Z",
	  "links": [
	    {
	      "href": "https://api-m.paypal.com/v1/catalogs/products/PROD-1YN92718J3553841U ",
	      "rel": "self",
	      "method": "GET"
	    },
	    {
	      "href": "https://api-m.paypal.com/v1/catalogs/products/PROD-1YN92718J3553841U ",
	      "rel": "edit",
	      "method": "PATCH"
	    }
	  ]
	}
`
