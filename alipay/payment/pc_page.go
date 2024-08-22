package payment

import (
	"fmt"
	"github.com/Bishoptylaor/paypay"
	"github.com/Bishoptylaor/paypay/alipay/aliClient"
	"github.com/Bishoptylaor/paypay/alipay/consts"
	"github.com/Bishoptylaor/paypay/alipay/flow"
)

type PCPageProxy interface {
	flow.PcPagePay
	InjectCustomPayloadCheckRuler(custom map[string][]paypay.Ruler)
}

type PCPageCaller struct {
	*aliClient.Client
	rulersMap map[string][]paypay.Ruler
}

// NewPCPageCaller
//
// 初始化 电脑网站 支付相关接口功能
//
// 提供预设参数与参数校验能力
//
// 在保证参数正确的情况下，用户也可直接调用 client 中的相关接口实现
//
// 产品介绍 https://opendocs.alipay.com/open/270/105898?pathHash=b3b2b667
func NewPCPageCaller(c *aliClient.Client) PCPageProxy {
	// do some implantation
	caller := &PCPageCaller{Client: c}
	// 设置本产品相关接口默认参数校验规则
	caller.rulersMap = caller.setDefaultPayloadCheckRuler()
	aliClient.SetChecker(caller.payloadChecker())(c)
	aliClient.SetPayloadPreSetter(map[string][]paypay.PayloadPreSetter{
		"alipay.trade.page.pay": []paypay.PayloadPreSetter{
			paypay.PreSetter("product_code", "FAST_INSTANT_TRADE_PAY"),
		},
	})(c)
	return caller
}

// InjectCustomPayloadCheckRuler 外部可调用，用于自定义参数校验规则
func (c PCPageCaller) InjectCustomPayloadCheckRuler(custom map[string][]paypay.Ruler) {
	if c.rulersMap == nil {
		c.rulersMap = c.setDefaultPayloadCheckRuler()
	}
	for k, v := range custom {
		c.rulersMap[k] = v
	}
}

func (c PCPageCaller) payloadChecker() paypay.PayloadRuler {
	return func(method string) []paypay.Ruler {
		if rulers, ok := c.rulersMap[method]; ok {
			return rulers
		} else {
			return []paypay.Ruler{}
		}
	}
}

func (c PCPageCaller) setDefaultPayloadCheckRuler() map[string][]paypay.Ruler {
	return map[string][]paypay.Ruler{
		"alipay.trade.page.pay": []paypay.Ruler{
			paypay.NewRuler("商户订单号", "out_trade_no != nil && len(out_trade_no) <= 64", fmt.Sprintf(consts.FmtEmptyAlert, "out_trade_no")),
			consts.TotalAmountDefaultRuler,
			paypay.NewRuler("订单标题", "subject != nil", fmt.Sprintf(consts.FmtEmptyAlert, "subject")),
			paypay.NewRuler("销售产品码，商家和支付宝签约的产品码", `product_code == "FAST_INSTANT_TRADE_PAY"`, "当前场景 product_code 传值必为 FAST_INSTANT_TRADE_PAY"),
		},
		"alipay.trade.refund": []paypay.Ruler{
			paypay.NewRuler("退款金额", "refund_amount != nil", fmt.Sprintf(consts.FmtEmptyAlert, "refund_amount")),
			consts.TradeNo2in1DefaultRuler,
			paypay.NewRuler("退款原因说明", "refund_reason != nil", fmt.Sprintf(consts.FmtEmptyAlert, "refund_reason")),
			// 一般来说这个字段只有在部分退款时必传，这里进行一般性校验，保证调用方业务一致性
			paypay.NewRuler("退款请求号", "out_request_no != nil", fmt.Sprintf(consts.FmtEmptyAlert, "out_request_no")),
		},
		"alipay.trade.close": []paypay.Ruler{
			consts.TradeNo2in1DefaultRuler,
		},
		"alipay.trade.fastpay.refund.query": []paypay.Ruler{
			consts.TradeNo2in1DefaultRuler,
			paypay.NewRuler("退款请求号", "out_request_no != nil", fmt.Sprintf(consts.FmtEmptyAlert, "out_request_no")),
		},
		"alipay.trade.query": []paypay.Ruler{
			consts.TradeNo2in1DefaultRuler,
		},

		"alipay.data.dataservice.bill.downloadurl.query": consts.DataDownloadRuler,
	}
}
