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
 @Time    : 2024/9/3 -- 17:45
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: settings.go
*/

package paypal

import (
	"context"
	"github.com/Bishoptylaor/paypay"
	"github.com/Bishoptylaor/paypay/alipay/consts"
	"github.com/Bishoptylaor/paypay/pkg"
	"github.com/Bishoptylaor/paypay/pkg/xlog"
	"github.com/Bishoptylaor/paypay/pkg/xnet/xhttp"
)

type Settings func(*Client)

func Checker(checker paypay.PayloadRuler) Settings {
	return func(client *Client) {
		client.EmptyChecker = checker
	}
}

func DefaultChecker() Settings {
	return Checker(func(uri string) []paypay.Ruler {
		return []paypay.Ruler{}
	})
}

func PayloadPreSetter(setter map[string][]paypay.PayloadPreSetter) Settings {
	return func(client *Client) {
		client.PayloadPreSetter = setter
	}
}

func ClientID(cid string) Settings {
	return func(client *Client) {
		client.ClientID = cid
	}
}

func Secret(secret string) Settings {
	return func(client *Client) {
		client.ClientSecret = secret
	}
}

func Debug(d bool) Settings {
	return func(client *Client) {
		client.debug = d
	}
}

func HClient(hc xhttp.HttpClientWrapper) Settings {
	return func(client *Client) {
		client.HClient = hc
	}
}

func DefaultHClient() Settings {
	return HClient(xhttp.GetDefaultClient())
}

func SetLogger(logger xlog.XLogger) Settings {
	return func(client *Client) {
		client.Logger = logger
	}
}

func DefaultLogger() Settings {
	return SetLogger(xlog.NewLogger())
}

func Prod(prod bool) Settings {
	return func(client *Client) {
		client.Prod = prod
	}
}

func Proxy(prod, sandbox string) Settings {
	return func(client *Client) {
		client.ProxyProd = prod
		client.ProxySandbox = sandbox
	}
}

func Headers(headers map[string]string) Settings {
	return func(client *Client) {
		if client.Headers == nil {
			client.Headers = headers
		} else {
			for k, v := range headers {
				client.Headers[k] = v
			}
		}
	}
}

func DefaultHeaders() Settings {
	return Headers(map[string]string{
		// representation || minimal ; minimal by default
		"Prefer": "minimal",
		// used for identifies a merchant. you can change it with the client.Use function
		// "PayPal-Auth-Assertion": "",
	})
}

func PrefixFunc(suppress bool, pres ...xhttp.ReqPrefixFunc) Settings {
	return func(client *Client) {
		if suppress {
			client.PrefixFunc = pres
			return
		}
		if client.PrefixFunc == nil {
			client.PrefixFunc = []xhttp.ReqPrefixFunc{}
		}
		client.PrefixFunc = append(client.PrefixFunc, pres...)
	}
}

func DefaultPrefixFunc() Settings {
	return func(client *Client) {
		PrefixFunc(true, PPReqPrefix(client.debug, client.Logger))
	}
}

func SuffixFunc(suppress bool, sufs ...xhttp.ResSuffixFunc) Settings {
	return func(client *Client) {
		if suppress {
			client.SuffixFunc = sufs
			return
		}
		if client.SuffixFunc == nil {
			client.SuffixFunc = []xhttp.ResSuffixFunc{}
		}
		client.SuffixFunc = append(client.SuffixFunc, sufs...)
	}
}

func DefaultSuffixFunc() Settings {
	return func(client *Client) {
		SuffixFunc(true, PPResSuffix(client.debug, client.Logger))
	}
}

// NewSettings 标准初始化配置
func NewSettings(ins ...Settings) []Settings {
	return append(
		append(
			[]Settings{},
			DefaultLogger(),     // 设置 logger
			DefaultHClient(),    // 设置 Http client
			DefaultChecker(),    // 设置 checker 初始化
			DefaultHeaders(),    // 设置 header 自定义部分
			DefaultPrefixFunc(), // 设置 http prefix hooks, will pass on to client.HClient and run before requests are made.
			DefaultSuffixFunc(), // 设置 http suffix hooks, will pass on to client.HClient and run after requests are made.
		), ins...,
	)
}

// DefaultSettings 默认沙盒配置，可自定义追加，按照顺序执行，自定义追加的部分会覆盖掉默认内容
func DefaultSettings(ins ...Settings) []Settings {
	return append(
		append(
			NewSettings(),
			ClientID(consts.Appid),    // 设置 沙盒 clientId
			Secret(consts.PrivateKey), // 设置 沙盒 secret
			Prod(pkg.SandBox),         // 设置 沙盒环境
			Debug(pkg.DebugOn),        // debug on
		), ins...,
	)
}

func PackSettings(i1 []Settings, i2 ...Settings) []Settings {
	return append(i1, i2...)
}

func NewToken(ctx context.Context, clientId, clientSecret string) Settings {
	return func(client *Client) {
		client.ClientID = clientId
		client.ClientSecret = clientSecret
		_, err := client.getAccessToken(ctx)
		if err != nil {
			client.Logger.Errorf("[NewToken]Error getting access token: %v", err)
		}
	}
}
