package main

import (
	"context"
	"encoding/json"
	"github.com/Bishoptylaor/paypay"
	"github.com/Bishoptylaor/paypay/alipay/service/payment"
	"github.com/Bishoptylaor/paypay/pkg/xlog"
	"time"
)

type md struct{}

func (md) TradeRefund(ctx context.Context, caller payment.MerchantDeductionCaller) {
	// 请求参数
	pl := make(paypay.Payload)
	pl.Set("refund_amount", "0.02")                    // 退款金额
	pl.Set("trade_no", "2024082022001422760503807863") // 支付宝订单号

	// 查询订单
	aliRes, err := caller.TradeRefund(ctx, pl)
	if err != nil {
		xlog.Error("err:", err)
		return
	}
	xlog.Info("aliRes:", *aliRes)
}

func (md) TradePay(ctx context.Context, caller payment.MerchantDeductionCaller) {
	// 请求参数
	pl := make(paypay.Payload)
	pl.Set("out_trade_no", getTradeNo())     // 商户订单号
	pl.Set("total_amount", "50")             // 本次支付金额。
	pl.Set("subject", "测试代扣")                // 订单标题
	pl.Set("product_code", "CYCLE_PAY_AUTH") // 产品码,固定值
	pl.Set("agreement_params", func(pl paypay.Payload) {
		pl.Set("agreement_no", "20245622025279568776")
	})
	pl.Set("product_code", "GENERAL_WITHHOLDING")

	// 查询订单
	aliRes, err := caller.TradePay(ctx, pl)
	if err != nil {
		xlog.Error("err:", err)
		return
	}
	xlog.Infof(ctx, "aliRes: %+v", *aliRes)
	xlog.Infof(ctx, "response: %+v", aliRes.Response)
}

func (md) UserAgreementQuery(ctx context.Context, caller payment.MerchantDeductionCaller) {
	// 请求参数
	pl := make(paypay.Payload)
	// pl.Set("out_trade_no", "GZ201909081743431443")
	// pl.Set("alipay_open_id", "")
	// pl.Set("agreement_no", "20245622095470574330")
	pl.Set("personal_product_code", "CYCLE_PAY_AUTH_P")
	pl.Set("sign_scene", "INDUSTRY|DEFAULT_SCENE")
	// pl.Set("external_agreement_no", "2024082210000123")
	pl.Set("external_agreement_no", "test1445544789")
	pl.Set("third_party_type", "")

	// 查询订单
	aliRes, err := caller.UserAgreementQuery(ctx, pl)
	if err != nil {
		xlog.Error("err:", err)
		return
	}
	/*
		aliRes:&{ErrorResponse:{Code:10000 Msg:Success SubCode: SubMsg:} PrincipalId: PrincipalOpenId:068sgou5i97DlNS0nPpqF6NEj4DajfCnK624rZX37whshYb ValidTime:2024-08-22 11:13:39 AlipayLogonId:aya***@sina.com InvalidTime:2115-02-01 00:00:00 PricipalType:CARD DeviceId: SignScene:INDUSTRY|DEFAULT_SCENE AgreementNo:20245622072469698668 ThirdPartyType:PARTNER Status:NORMAL SignTime:2024-08-22 11:13:39 PersonalProductCode:CYCLE_PAY_AUTH_P ExternalAgreementNo:2024082210000123 ZmOpenId: ExternalLogonId: CreditAuthMode: SingleQuota: LastDeductTime: NextDeductTime:2024-08-22}
	*/
	xlog.Infof(ctx, "aliRes:%+v", aliRes.Response)
}

func (md) TradeQuery(ctx context.Context, caller payment.MerchantDeductionCaller) {
	// 请求参数
	pl := make(paypay.Payload)
	pl.Set("out_trade_no", "080514075435733")

	// 查询订单
	aliRes, err := caller.TradeQuery(ctx, pl)
	if err != nil {
		xlog.Error("err:", err)
		return
	}
	xlog.Debug("aliRes:", *aliRes)
}

func (md) UserAgreementPageSignInApp(ctx context.Context, caller payment.MerchantDeductionCaller) {
	// 请求参数
	pl := make(paypay.Payload)
	pl.Set("personal_product_code", "CYCLE_PAY_AUTH_P")
	pl.Set("access_params", func(pl paypay.Payload) {
		pl.Set("channel", "ALIPAYAPP")
	})
	pl.Set("period_rule_params", func(pl paypay.Payload) {
		pl.Set("period_type", "DAY")
		pl.Set("period", "7")
		pl.Set("execute_time", time.Now().Format("2006-01-02"))
		pl.Set("single_amount", "99")
		// pl.Set("total_amount", "0.04")
		// pl.Set("total_payments", "2")
	})
	pl.Set("product_code", "GENERAL_WITHHOLDING")
	pl.Set("external_agreement_no", "test1445544789")
	pl.Set("sign_scene", "INDUSTRY|DEFAULT_SCENE")

	// 查询订单
	aliRes, err := caller.UserAgreementPageSignInApp(ctx, pl)
	if err != nil {
		xlog.Error("err:", err)
		return
	}
	xlog.Info(aliRes)
	xlog.Debug("aliRes:", aliRes)

	aliRes, err = caller.PageExecute(ctx, pl, "alipay.user.agreement.page.sign")
	if err != nil {
		xlog.Error("err:", err)
		return
	}
	xlog.Info(aliRes)
	xlog.Debug("aliRes:", aliRes)
}

func (md) DataBillDownloadUrlQuery(ctx context.Context, caller payment.MerchantDeductionCaller) {
	// 请求参数
	pl := make(paypay.Payload)
	pl.Set("biz_content", func(pl paypay.Payload) {
		pl.Set("bill_type", "trade")
		pl.Set("bill_date", "2024-08")
		// pl.Set("bill_date", "2024-08-05")
	})

	// 查询订单
	aliRes, err := caller.DataBillDownloadUrlQuery(ctx, pl)
	if err != nil {
		xlog.Error("err:", err)
		return
	}
	xlog.Debug("aliRes:", *aliRes)
}

func (md) TradeAppPay(ctx context.Context, caller payment.MerchantDeductionCaller) {
	payload := make(paypay.Payload)
	payload.Set("out_trade_no", getTradeNo()) // 商户订单号
	payload.Set("total_amount", "0.01")       // 订单总金额，首次支付的金额，不算在周期扣总金额里。
	payload.Set("subject", "测试支付签约")          // 订单标题
	// payload.Set("product_code", "GENERAL_WITHHOLDING") //产品码,固定值
	// payload.Set("product_code", "CYCLE_PAY_AUTH") //产品码,固定值 app支付
	payload.Set("product_code", "QUICK_MSECURITY_PAY")                            // 产品码,固定值 无线快捷支付
	payload.Set("time_expire", time.Now().AddDate(0, 0, 2).Format(time.DateTime)) // 超时时间
	payload.Set("agreement_sign_params", func(payload paypay.Payload) {
		payload.Set("product_code", "GENERAL_WITHHOLDING")
		payload.Set("personal_product_code", "CYCLE_PAY_AUTH_P") // 个人签约产品码固定为CYCLE_PAY_AUTH_P
		payload.Set("sign_scene", "INDUSTRY|DEFAULT_SCENE")      // 协议签约场景，参见下文sign_scene参数说明
		// payload.Set("external_agreement_no", "2024082210000123") // 线上 测试 已解约
		payload.Set("external_agreement_no", "2024082210000456") // 应用 订单ID //商户签约号，代扣协议中用户的唯一签约号
		// pl.Set("sign_notify_url", "https://testtt.cabc") //签约成功异步通知地址
		payload.Set("access_params", func(payload paypay.Payload) {
			payload.Set("channel", "ALIPAYAPP")
		}) // 签约接入的方式
		payload.Set("period_rule_params", func(payload paypay.Payload) { // 周期扣款协议信息
			payload.Set("period_type", "DAY")                             // 周期类型枚举值为 DAY 和 MONTH
			payload.Set("period", "7")                                    // 周期数，与 period_type 组合使用确定扣款周期
			payload.Set("execute_time", time.Now().Format(time.DateOnly)) // 用户签约后，下一次使用代扣支付扣款的时间
			payload.Set("single_amount", "99")                            // 单次扣款最大金额
			// pl.Set("total_amount", "0.04")       //周期内扣允许扣款的总金额，单位为元
			// pl.Set("total_payments", "2")
			// payload.Set("sign_notify_url", "url") // 签约成功后商户用于接收异步通知的地址 签约单独通知
		})
	})

	v, _ := json.Marshal(payload)
	xlog.Infof(ctx, "%+v", string(v))

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
