package main

import (
	"context"
	"github.com/Bishoptylaor/paypay"
	"github.com/Bishoptylaor/paypay/alipay/service/payment"
	"github.com/Bishoptylaor/paypay/pkg/xlog"
)

type app struct{}

func (app) TradeAppPay(ctx context.Context, caller payment.AppCaller) {
	payload := make(paypay.Payload)
	payload.Set("out_trade_no", getTradeNo())          // 商户订单号
	payload.Set("total_amount", "0.01")                // 订单总金额，首次支付的金额，不算在周期扣总金额里。
	payload.Set("subject", "测试支付签约")                   // 订单标题
	payload.Set("product_code", "QUICK_MSECURITY_PAY") // 产品码,固定值
	// payload.Set("time_expire", "2024-12-31 10:05:00")  //超时时间

	res, err := caller.TradeAppPay(ctx, payload)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Info(res)
	res, err = caller.PageExecute(ctx, payload, "alipay.trade.app.pay")
	xlog.Info(res)
	return
}
