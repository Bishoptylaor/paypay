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
 @Time    : 2024/10/16 -- 15:46
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: advanced.go
*/

package main

import (
	"context"
	"github.com/Bishoptylaor/paypay"
	"github.com/Bishoptylaor/paypay/paypal"
	"github.com/Bishoptylaor/paypay/paypal/entity"
	"github.com/Bishoptylaor/paypay/pkg/xlog"
	"github.com/Bishoptylaor/paypay/pkg/xnet/xhttp"
	"net/http"
	"time"
)

func initAdvanced() {
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

func customSettings() paypal.Settings {
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
