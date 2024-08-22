package payment

import (
	"fmt"
	"github.com/Bishoptylaor/paypay"
	"github.com/Bishoptylaor/paypay/alipay/aliClient"
	"github.com/Bishoptylaor/paypay/alipay/consts"
	"github.com/Bishoptylaor/paypay/alipay/flow"
)

type AppCaller struct{}

// NewAppCaller
//
// 初始化 app 支付相关接口功能
//
// 提供预设参数与参数校验能力
//
// 在保证参数正确的情况下，用户也可直接调用 client 中的相关接口实现
//
// 产品介绍 https://opendocs.alipay.com/open/204/105051?pathHash=b91b9616
func NewAppCaller(a *aliClient.Client) flow.AppPay {
	// do some implantation
	md := AppCaller{}
	aliClient.SetChecker(md.payloadChecker())(a)
	aliClient.SetPayloadPreSetter(map[string][]paypay.PayloadPreSetter{
		"alipay.trade.app.pay": []paypay.PayloadPreSetter{
			paypay.PreSetter("product_code", "QUICK_MSECURITY_PAY"),
		},
	})(a)
	return a
}

func (AppCaller) payloadChecker() paypay.PayloadRuler {
	return func(caller string) []paypay.Ruler {
		switch caller {
		case "alipay.trade.app.pay":
			return []paypay.Ruler{
				paypay.NewRuler("商户订单号", "out_trade_no != nil && len(out_trade_no) <= 64", fmt.Sprintf(consts.FmtEmptyAlert, "out_trade_no")),
				consts.TotalAmountDefaultRuler,
				paypay.NewRuler("订单标题", "subject != nil", fmt.Sprintf(consts.FmtEmptyAlert, "subject")),
				paypay.NewRuler("销售产品码，商家和支付宝签约的产品码", `product_code == "QUICK_MSECURITY_PAY"`, "当前场景 product_code 传值必为 QUICK_MSECURITY_PAY"),
			}
		case "alipay.trade.refund":
			return []paypay.Ruler{
				paypay.NewRuler("退款金额", "refund_amount != nil", fmt.Sprintf(consts.FmtEmptyAlert, "refund_amount")),
				consts.TradeNo2in1DefaultRuler,
				paypay.NewRuler("退款原因说明", "refund_reason != nil", fmt.Sprintf(consts.FmtEmptyAlert, "refund_reason")),
				// 一般来说这个字段只有在部分退款时必传，这里进行一般性校验，保证调用方业务一致性
				paypay.NewRuler("退款请求号", "out_request_no != nil", fmt.Sprintf(consts.FmtEmptyAlert, "out_request_no")),
			}
		case "alipay.trade.close":
			return []paypay.Ruler{
				consts.TradeNo2in1DefaultRuler,
			}
		case "alipay.trade.fastpay.refund.query":
			return []paypay.Ruler{
				consts.TradeNo2in1DefaultRuler,
				paypay.NewRuler("退款请求号", "out_request_no != nil", fmt.Sprintf(consts.FmtEmptyAlert, "out_request_no")),
			}
		case "alipay.trade.query":
			return []paypay.Ruler{
				consts.TradeNo2in1DefaultRuler,
			}

		case "alipay.data.dataservice.bill.downloadurl.query":
			return consts.DataDownloadRuler
		default:
			return []paypay.Ruler{}
		}
	}
}
