package main

import (
	"context"
	"github.com/Bishoptylaor/paypay"
	"github.com/Bishoptylaor/paypay/alipay/service/payment"
	"github.com/Bishoptylaor/paypay/pkg/xlog"
)

type pc struct{}

func (pc) TradePagePay(ctx context.Context, caller payment.PCPageCaller) {
	payload := make(paypay.Payload)
	payload.Set("out_trade_no", getTradeNo())             // 商户订单号
	payload.Set("total_amount", "12")                     // 订单总金额，首次支付的金额，不算在周期扣总金额里。
	payload.Set("subject", "电脑支付")                        // 订单标题
	payload.Set("product_code", "FAST_INSTANT_TRADE_PAY") // 产品码,固定值
	// payload.Set("time_expire", "2024-12-31 10:05:00")     //超时时间

	res, err := caller.TradePagePay(ctx, payload)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Info(res)
	return
}
