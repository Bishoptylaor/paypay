package main

import (
	"context"
	"github.com/Bishoptylaor/paypay"
	"github.com/Bishoptylaor/paypay/alipay/service/payment"
	"github.com/Bishoptylaor/paypay/pkg/xlog"
)

// 沙盒环境未开启相关功能
type qrcode struct{}

func (qrcode) TradePreCreate(ctx context.Context, caller payment.QrcodeCaller) {
	pl := make(paypay.Payload)
	tradeNo := getTradeNo()
	pl.Set("subject", "预创建商品订单").
		Set("out_trade_no", tradeNo).
		Set("total_amount", "2.00").
		Set("product_code", "QR_CODE_OFFLINE")

	// 创建订单
	aliRes, err := caller.TradePrecreate(ctx, pl)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Info("aliRes:%+v", aliRes.Response)
	xlog.Info("aliRes.QrCode:", aliRes.Response.QrCode)
	xlog.Info("aliRes.OutTradeNo:", aliRes.Response.OutTradeNo)

	Link2QRCode(aliRes.Response.QrCode, tradeNo)
}
