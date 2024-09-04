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
	"github.com/Bishoptylaor/paypay/paypal/consts"
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

func (c Config) GenUrl(path string) string {
	if c.Prod {
		return consts.BaseUrl + path
	}
	return consts.SandboxBaseUrl + path
}
