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
 @Time    : 2024/8/22 -- 12:19
 @Author  : 亓官竹
 @Copyright 2024 亓官竹
 @Description: consts.go
*/

package consts

import (
	"fmt"
	"github.com/Bishoptylaor/paypay"
	"regexp"
)

var (
	jsonCheck = regexp.MustCompile("(?i:(?:application|text)/json)")
	xmlCheck  = regexp.MustCompile("(?i:(?:application|text)/xml)")
)

type PKCSType uint8

const (
	BaseUrl        = "https://openapi.alipay.com/gateway.do"               // 正式环境请求地址
	SandboxBaseUrl = "https://openapi-sandbox.dl.alipaydev.com/gateway.do" // 沙箱环境请求地址

	BaseUrlUtf8        = "https://openapi.alipay.com/gateway.do?charset=utf-8"               // 正式环境请求地址 utf-8
	SandboxBaseUrlUtf8 = "https://openapi-sandbox.dl.alipaydev.com/gateway.do?charset=utf-8" // 沙箱环境请求地址 utf-8

	PublicAppAuthUrlProduction  = "https://openauth.alipay.com/oauth2/publicAppAuthorize.htm"    // 正式环境授权登录地址
	PublicAppAuthUrlDevelopment = "https://openauth.alipaydev.com/oauth2/publicAppAuthorize.htm" // 沙箱环境授权登录地址

	AppToAppAuthUrlProduction  = "https://openauth.alipay.com/oauth2/appToAppAuth.htm"    // 正式环境第三方授权登录地址
	AppToAppAuthUrlDevelopment = "https://openauth.alipaydev.com/oauth2/appToAppAuth.htm" // 沙箱环境第三方授权登录地址

	BaseUrlV3        = "https://openapi.alipay.com/v3/"
	SandboxBaseUrlV3 = "https://openapi.alipay.com/v3/"

	LocationShanghai = "Asia/Shanghai"
	RSA              = "RSA"
	RSA2             = "RSA2"
	UTF8             = "utf-8"
	PKCS1            = "PKCS1"
	PKCS8            = "PKCS8"
)

var (
	TotalAmountDefaultRuler = paypay.NewRuler("订单总金额", `float(total_amount) >= 0.01 && float(total_amount) <= 100000000`, "请检查价格金额区间是否在 0.01 ~ 100000000")
	TradeNo2in1DefaultRuler = paypay.NewRuler("商户订单号 + 支付宝交易号", "out_trade_no != nil || trade_no != nil", "请检查 out_trade_no 和 trade_no 不能同时为空")

	FmtAmountDefaultRule = `float(%s) >= 0.01 && float(%s) <= 100000000`
	Fmt2in1DefaultRule   = `%s != nil || %s != nil`

	// FmtEmptyAlert 1 params key
	FmtEmptyAlert = `请检查 %s 是否传值`
	// FmtWithinAlert 2 params key; needed values
	FmtWithinAlert = `请检查 %s 取值是否在 %+v 内`
	// Fmt2in1Alert 2 params key1; key2
	Fmt2in1Alert = "请检查 %s 和 %s 不能同时为空"
)

var DataDownloadRuler = []paypay.Ruler{
	paypay.NewRuler("", `bill_type in ["trade", "signcustomer", "merchant_act", "trade_zft_merchant", "zft_acc", "settlementMerge"]`,
		fmt.Sprintf(FmtWithinAlert, "bill_type", `["trade", "signcustomer", "merchant_act", "trade_zft_merchant", "zft_acc", "settlementMerge"]`)),
	paypay.NewRuler("", "bill_date != nil", fmt.Sprintf(FmtEmptyAlert, "bill_date")),
}
