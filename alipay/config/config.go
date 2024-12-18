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
 @Time    : 2024/8/26 -- 15:01
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: config.go
*/

package config

import (
	"github.com/Bishoptylaor/paypay/alipay/consts"
	"time"
)

type Config struct {
	// 可访问参数 & 请求通用参数
	AppId        string // 应用ID
	ReturnUrl    string
	NotifyUrl    string
	Charset      string
	Version      string
	Format       string
	SignType     string
	AppAuthToken string
	AuthToken    string
	Prod         bool // 是否是正式环境，沙箱环境请选择新版沙箱应用信息。
	Location     *time.Location

	// 公钥证书组
	AppCertSN          string // alipay 应用公钥证书
	AliPayPublicCertSN string // alipay 支付宝公钥证书
	AliPayRootCertSN   string // alipay 支付宝根证书
}

// Url get alipay gateway endpoint
func (c Config) Url() string {
	if c.Prod {
		return consts.BaseUrl
	}
	return consts.SandboxBaseUrl
}

// Utf8Url get alipay gateway endpoint
func (c Config) Utf8Url() string {
	if c.Prod {
		return consts.BaseUrlUtf8
	}
	return consts.SandboxBaseUrlUtf8
}
