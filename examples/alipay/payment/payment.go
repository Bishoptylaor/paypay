package main

import (
	"context"
	"fmt"
	"github.com/Bishoptylaor/paypay/alipay"
	"github.com/Bishoptylaor/paypay/alipay/consts"
	payment2 "github.com/Bishoptylaor/paypay/alipay/service/payment"
	"github.com/Bishoptylaor/paypay/pkg"
	"github.com/Bishoptylaor/paypay/pkg/xlog"
	"github.com/Bishoptylaor/paypay/pkg/xutils"
	qrcode2 "github.com/skip2/go-qrcode"
	"time"
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

	NewMerchantDeductionCaller(client)
	// NewAppCaller(client)
	// NewPCPageCaller(client)
	// NewQrcodeCaller(client)
}

func NewMerchantDeductionCaller(client *alipay.Client) {
	caller := payment2.NewMerchantDeductionCaller(client)
	s := md{}
	// s.TradeAppPay(ctx, caller)
	// s.TradeQuery(ctx, caller)
	// s.UserAgreementPageSignInApp(ctx, caller)
	// s.DataBillDownloadUrlQuery(ctx, caller)
	// s.UserAgreementPageUnSign(ctx, caller)
	// s.UserAgreementQuery(ctx, caller)
	s.TradePay(ctx, caller)
	// s.TradeRefund(ctx, caller)
}

func NewQrcodeCaller(client *alipay.Client) {
	// not supported in sandbox
	caller := payment2.NewQrcodeCaller(client)
	s := qrcode{}
	s.TradePreCreate(ctx, caller)
}

func NewAppCaller(client *alipay.Client) {
	caller := payment2.NewAppCaller(client)
	s := app{}
	s.TradeAppPay(ctx, caller)
}

func NewPCPageCaller(client *alipay.Client) {
	caller := payment2.NewPCPageCaller(client)
	s := pc{}
	s.TradePagePay(ctx, caller)
}

func getTradeNo() string {
	return fmt.Sprintf("%s%d", xutils.RandomString(6), time.Now().Unix())
}

func Link2QRCode(link, filename string) {
	// 1. 生成一张 png 图片
	// var png []byte
	filePath := filename + ".png"
	// 生成二维码
	err := qrcode2.WriteFile(link, qrcode2.Medium, 256, filePath)
	if err != nil {
		xlog.Error("err:", err)
	}
	fmt.Println("QR code generated successfully.")
}
