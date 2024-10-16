
## Quick Start

We have integrated basic settings in NewClient function, you can create a Paypal client easily.

```go
package main

import (
	"context"
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
		paypal.ClientID("Your ClientId"),
		paypal.Secret("Your ClientSecret"),
	)
	if err != nil {
		xlog.Error(err)
		return
	}
}

```

## Settings

You can also use settings for client like ↓ and use your custom setting options.

```go
package main

import (
	"context"
	"github.com/Bishoptylaor/paypay/paypal"
	"github.com/Bishoptylaor/paypay/pkg/xlog"
	"github.com/Bishoptylaor/paypay/pkg/xnet/xhttp"
)

var (
	ctx    context.Context
	client *paypal.Client
)

func init() {
	var err error
	ctx = context.Background()
	client, err = paypal.NewClient(ctx)
	if err != nil {
		xlog.Error(err)
		return
	}

	// set client id and secret
	// 设置账户 ID 和 密钥
	client.Use(paypal.ClientID(""))
	client.Use(paypal.Secret(""))
	// debug flag
	client.Use(paypal.Debug(true))

	// chaining 
	client.
		// use custom headers
		// 自定义 Headers
		Use(paypal.Headers(map[string]string{})).
		// use custom logger
		// 自定义 Logger
		// built-in xlog.Logger
		Use(paypal.SetLogger(xlog.NewLogger())).
		// change prod flag
		// 修改客户端访问环境
		Use(paypal.Prod(false))

	// use multiple settings
	client.Use(
		// use your http client configs
		// 使用你自己的 http client 配置
		paypal.HClient(
			xhttp.NewHttpClientWrapper(&http.Client{
				Timeout: time.Minute,
			}),
		),
		// set params checker before make requests to Paypal
		// 在调用接口前增加参数校验，一般不在 client 初始化阶段使用
		// 规则语法参考 
		paypal.Checker(
			paypay.InjectRuler(map[string][]paypay.Ruler{
				"uri path to your function": []paypay.Ruler{
					paypay.NewRuler("field", `field != nil`, "field in payload cannot be nil"),
				},
			}),
		),
	)
}
```

## Dynamic request to Paypal

### Usage

custom request

```go
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
		paypal.ListProducts,
		func() map[string]string {
			return map[string]string{}
		},
		nil,
		query,
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
```

another example

```go
// any field in your config or db
func customSettings(field1, field2 any) paypal.Settings {
	return func(c *paypal.Client) {
		// set you own info on any field exported
		c.ClientID = "your client id"
		c.ClientSecret = "your client secret"
		c.Prod = false
		// etc.
	}
}

func initCustom() {
	var err error
	client, err = paypal.NewClient(ctx, customSettings())
	if err != nil {
		xlog.Error(err)
	}

	emptyRes := entity.EmptyRes{}

	err = client.CustomCallOnce(
		context.Background(),
		paypal.EmptyMethod,
		func() map[string]string { return map[string]string{} },
		nil,
		nil,
		nil,
		emptyRes,
		nil,
		customSettings(),
	)
}
```

## Hooks

We provide hooks which run before request started, and after response received, so you can do something like inject context, tracing's span etc. it is just similar to web middleware.
Here is a default usage, and we`ve implanted it to Client, help you have a concept. 

```go

// PPReqPrefix 闭包注入 logger 和 debug 信息
func PPReqPrefix(debug bool, log xlog.XLogger) xhttp.ReqPrefixFunc {
	return func(ctx context.Context, req *http.Request) context.Context {
		if debug == pkg.DebugOn {
			log.Debugf("PayPal_Url: %s", req.URL)
			log.Debugf("PayPal_Req_Body: %s", req.Body)
			log.Debugf("PayPal_Req_Headers: %#v", req.Header)
		} else {
			body, err := io.ReadAll(req.Body)
			if err != nil {
				log.Errorf("[Read Req body] error: %s", err.Error())
				return ctx
			}
			enEscapeUrl, err := url.QueryUnescape(string(body))
			if err == nil {
				log.Infof("[Req] %s", enEscapeUrl)
			}
			req.Body = io.NopCloser(bytes.NewBuffer(body))
		}
		return ctx
	}
}

// PPResSuffix 闭包注入 logger 和 debug 信息
func PPResSuffix(debug bool, log xlog.XLogger) xhttp.ResSuffixFunc {
	return func(ctx context.Context, res *http.Response) context.Context {
		if debug == pkg.DebugOn {
			log.Debugf("PayPal_Response: %d > %s", res.StatusCode, res.Body)
			log.Debugf("PayPal_Rsp_Headers: %#v", res.Header)
		} else {
			body, err := io.ReadAll(res.Body)
			if err != nil {
				log.Errorf("[Read Res body] error: %s", err.Error())
				return ctx
			}
			res.Body = io.NopCloser(bytes.NewBuffer(body))
			log.Infof("[Res] %s", string(body))
		}
		return ctx
	}
}
```

If former function not suitable for your project, you can use PrefixFunc and SuffixFunc from settings.go to build your own.

You can pass suppress field true to overwrite default functions

Here`s an example of usage.
```go
func DefaultSuffixFunc() Settings {
    return func(client *Client) {
        SuffixFunc(true, PPResSuffix(client.debug, client.Logger))
    }
}

func DefaultPrefixFunc() Settings {
	return func(client *Client) {
		PrefixFunc(true, PPReqPrefix(client.debug, client.Logger))
	}
}
```

## Following functions of Paypal are Supported

[Catalog Products](https://developer.paypal.com/docs/api/catalog-products/v1/)

[Disputes](https://developer.paypal.com/docs/api/customer-disputes/v1/)

[Invoicing](https://developer.paypal.com/docs/api/invoicing/v2/)

[Orders](https://developer.paypal.com/docs/api/orders/v2/)

[Payments](https://developer.paypal.com/docs/api/payments/v2/)

[Payouts](https://developer.paypal.com/docs/api/payments/v2/)

[Subscriptions](https://developer.paypal.com/docs/api/subscriptions/v1/)

[Transaction Search](https://developer.paypal.com/docs/api/transaction-search/v1/)
