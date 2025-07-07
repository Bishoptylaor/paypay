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
 @Time    : 2024/12/18 -- 18:11
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: alipay.go
*/

package main

import (
	"context"
	"fmt"
	"github.com/Bishoptylaor/paypay/alipay"
	"github.com/Bishoptylaor/paypay/alipay/cert"
	"github.com/Bishoptylaor/paypay/alipay/consts"
	"github.com/Bishoptylaor/paypay/pkg"
)

var ctx = context.Background()

func init() {
	Prod = false
	Cert = true
}

var (
	Prod bool
	Cert bool

	SandBoxSettingss = []alipay.Settings{
		alipay.AppId(consts.Appid),           // 设置 沙盒 appid
		alipay.PrivateKey(consts.PrivateKey), // 设置 沙盒 private key
		alipay.Prod(pkg.SandBox),             // 设置 沙盒环境
		alipay.Debug(pkg.DebugOff),           // debug on
		alipay.DefaultSign(),                 // 设置 签名机 需要先设置 private key, signType
	}
	SandBoxCertSettings = []alipay.Settings{
		alipay.CertSnContent(
			consts.AppPublicContent,
			consts.AlipayRootContent,
			consts.AlipayPublicContentRSA2,
		),
		// alipay.CertSnFile(
		// 	"assets/appPublicCert.crt",
		// 	"assets/alipayRootCert.crt",
		// 	"assets/alipayPublicCert.crt",
		// ),
		alipay.AutoVerify(),
	}

	ProdSettings = []alipay.Settings{
		alipay.AppId(consts.OnlineAppid),
		alipay.PrivateKey(consts.OnlinePrivateKey), // 设置 沙盒 private key
		alipay.Prod(pkg.Online),                    // 设置 沙盒环境
		alipay.Debug(pkg.DebugOff),                 // debug on
		alipay.DefaultSign(),                       // 设置 签名机 需要先设置 private key, signType
	}
	OnlineCertSettings = []alipay.Settings{}
)

func main() {
	client, err := alipay.NewClient(ctx,
		alipay.PackSettings(
			SandBoxSettingss,
			// alipay.PublicKey(consts.PublicKey),
			// alipay.SetEncryptKey("Dvz9Wh3RVrj6APII5hzDrw=="),
		)...,
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	if Cert {
		client.Use(SandBoxCertSettings...)
	}

	if Prod {
		client, err = alipay.NewClient(ctx,
			alipay.PackSettings(
				ProdSettings,
				alipay.NotifyUrl(""),
				alipay.ReturnUrl(""),
			)...,
		)
		if err != nil {
			fmt.Println(err)
			return
		}
		if Cert {
			client.Use(OnlineCertSettings...)
		}
	}

	fmt.Printf("Client: %+v\n", client)

	// payment.NewMerchantDeductionCaller(client)
	// NewAppCaller(client)
	// NewPCPageCaller(client)
	// NewQrcodeCaller(client)
	PrivateKey()
}

func PrivateKey() {
	key := cert.FormatAlipayPrivateKey("MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCjBNGgkG7p4gnZzagNJNiM1FisU+55amEwaAZKSuyxGgpQUynz+Enp+t+8gNFf5zFMPBdpa2pBW0DXr9T8BnXKyKI7+ZZX3PCKgXSe7pj2UNRvSOCqUh2ipVsgg0gagujrQ5LIY1K+i3YB+jhpPmfaD4WjwF7uP2HMCVaOB5W9JIL3TijKpwf6qg7xSPsdNrdCo+qXSSNITfqyfTIlRDMGqk6nNqWo4shsxScBYfUbyYXDHjoZ8Zke4uEci58A1N9+SVfY1aSkkx2kxnKnECE69lq94bWjG6+OhPF9X8+4YXb1471oy4JhZc8yo5k7mPkGMAwq7CKoT8+fWCTTH+4rAgMBAAECggEBAJJQL+nB2Jq4kwjWib+KU4ZKzG+XQVHeRe7RXtvtVXU7HwirK9tZ/NUxowAoSkOeWGPJB2ZJz2trgM5VdX7iOXTRELNttx+TIJ3fxREkq0tgjtp2+ZZE725cpWaQDfYlivIliyep5ck5mKf7sgfbyBAyoEqxzUJexuL/ep0td5rzLBan3PBU4GgtPuqdAPB90DhFDbEhteH5fVppyiGkG7ZxwX8d4ufErZwvOPXkgH0Jq5SrfH07wOMIX7JEIhjp2hHOdz5bmLj/bfs8OUMtGcHDs5SF4Z3c2e5wvVnjbgtuKmDTQskA+tNUW675uY65BqvY5iY7vk5Na0wFZ8kUXykCgYEA10SX/zTckGOaFTIdndTu4Nly5SO+WUsnNYAysbtNrPpvxKPIt8jAGKAC3fArGYTVG4Pvm6fzXASdotF14Otp3Y+bMwn8XJHgFRQYgtSX8wEtYL2QXmmWDXanPeBxd2swK4A4A1JJcNJHZZGuNyk4PlOKwllj2hux9Cg6hmaJOA0CgYEAwd1R14ludW1D+ILnrl2kaOnzI0EaV2xbXpFk9wN0csGoAFCqf51HbR3A/OLUvYtOvw+NLZ7S7faf+7cRWpsDURdmvdhAC4b1at6VK7Kiz38hfWbIm4HcrYVTMJePtzJ9sIWWnJ48hXlPJxdPZIh+YQfADMqidZ7OPOnHm39VORcCgYA3BYpQKcPCiCLHDiMxx04zXIEYflVV33arHxRWB0joUK6mtTJFf2NSJ/vznEJ36FLMEKH6yCfUuj1Mc3tMvP83KVCC8Cd3xbVfzd4h+pLIAuIFWw5+g0BeOpiuoSJ/Yn9SbxhiCPNKigTDukOsR7XfgFz0JgL2qaOHg4zJ84VZNQKBgCmMo6PMQq5MMssqUzldvIT3zzJ5G4otIRANnVQLm4OU2crnfwNG1C4guGH6w5m0J7ZvQNTMO2zg9LzgKYn/uyBAIbcCiAdwb1zuQ45QBPDv0F+uDcS2enN+Jaw+b4DWA3Y3y2XZzCLXMw2vw0ZKsrtdC94x98u5s33+pNdA3SUVAoGBAKRHNy5erb5KqAIdaj/qj+oZeJwT7eFXqhl7Oh9f013wNZP7KkqykZHHUz4ZEsaGEZTLVYKqx1bWJ8UEaDWCaCikPL5IVkIfaSW7uKSqkovzVbzvRECXkab7Oz+FXFqPU3Dpl46bJWfzgm//Tks3gkpcfzRrBgz+L26w2GY3up8V")

	fmt.Println(key)
}
