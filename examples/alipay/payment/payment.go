package main

import (
	"context"
	"fmt"
	"github.com/Bishoptylaor/paypay/alipay"
	payment2 "github.com/Bishoptylaor/paypay/alipay/service/payment"
	"github.com/Bishoptylaor/paypay/pkg/xlog"
	"github.com/Bishoptylaor/paypay/pkg/xutils"
	qrcode2 "github.com/skip2/go-qrcode"
	"time"
)

func NewMerchantDeductionCaller(ctx context.Context, client *alipay.Client) {
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

func NewQrcodeCaller(ctx context.Context, client *alipay.Client) {
	// not supported in sandbox
	caller := payment2.NewQrcodeCaller(client)
	s := qrcode{}
	s.TradePreCreate(ctx, caller)
}

func NewAppCaller(ctx context.Context, client *alipay.Client) {
	caller := payment2.NewAppCaller(client)
	s := app{}
	s.TradeAppPay(ctx, caller)
}

func NewPCPageCaller(ctx context.Context, client *alipay.Client) {
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
