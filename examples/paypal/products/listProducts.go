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
 @Time    : 2024/10/14 -- 16:52
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: listProducts.go
*/

package products

import (
	"context"
	"github.com/Bishoptylaor/paypay"
	"github.com/Bishoptylaor/paypay/paypal"
	"github.com/Bishoptylaor/paypay/paypal/consts"
	"github.com/Bishoptylaor/paypay/pkg/xlog"
)

func ListProducts(ctx context.Context, client *paypal.Client) {
	pl := make(paypay.Payload)
	pl.Set("page_size", "10").
		Set("page", "1"). // starts from 1
		Set("total_required", "false")

	res, err := client.ListProducts(ctx, pl)
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

var listProductExamples = `
{
  "total_items": 20,
  "total_pages": 10,
  "products": [
    {
      "id": "72255d4849af8ed6e0df1173",
      "name": "Video Streaming Service",
      "description": "Video streaming service",
      "create_time": "2018-12-10T21:20:49Z",
      "links": [
        {
          "href": "https://api-m.paypal.com/v1/catalogs/products/72255d4849af8ed6e0df1173",
          "rel": "self",
          "method": "GET"
        }
      ]
    },
    {
      "id": "PROD-XYAB12ABSB7868434",
      "name": "Video Streaming Service",
      "description": "Audio streaming service",
      "create_time": "2018-12-10T21:20:49Z",
      "links": [
        {
          "href": "https://api-m.paypal.com/v1/catalogs/products/125d4849af8ed6e0df18",
          "rel": "self",
          "method": "GET"
        }
      ]
    }
  ],
  "links": [
    {
      "href": "https://api-m.paypal.com/v1/catalogs/products?page_size=2&page=1",
      "rel": "self",
      "method": "GET"
    },
    {
      "href": "https://api-m.paypal.com/v1/catalogs/products?page_size=2&page=2",
      "rel": "next",
      "method": "GET"
    },
    {
      "href": "https://api-m.paypal.com/v1/catalogs/products?page_size=2&page=10",
      "rel": "last",
      "method": "GET"
    }
  ]
}
`
