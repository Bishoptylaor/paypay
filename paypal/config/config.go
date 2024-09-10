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
 @Time    : 2024/9/3 -- 17:34
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: config.go
*/

package config

import (
	"bytes"
	"context"
	"github.com/Bishoptylaor/paypay/paypal/consts"
	"github.com/Bishoptylaor/paypay/pkg/xlog"
	"text/template"
	"time"
)

type Config struct {
	ClientID     string
	ClientSecret string
	Appid        string
	AccessToken  string
	ExpireIn     time.Duration
	Prod         bool   // 是否是正式环境，沙箱环境请选择新版沙箱应用信息。
	ProxyProd    string // 代理 URL
	ProxySandbox string // 代理 沙盒 URL
}

func (c Config) GenUrl(ctx context.Context, params interface{}) func(format string) string {
	return func(format string) string {
		fun := "BuildUri -->"
		b := &bytes.Buffer{}
		defer func() {
			if r := recover(); r != nil {
				xlog.Errorf(ctx, "%s fail err: template.Panic. format:%s, params:%+v", fun, format, params)
			}
		}()
		_ = template.Must(template.New("").Parse(format)).Execute(b, params)

		if c.Prod {
			return consts.BaseUrl + b.String()
		}
		return consts.SandboxBaseUrl + b.String()
	}
}
