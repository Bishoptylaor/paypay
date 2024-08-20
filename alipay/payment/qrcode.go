package payment

import (
	"github.com/Bishoptylaor/paypay/alipay"
	"github.com/Bishoptylaor/paypay/alipay/aliClient"
	"github.com/Bishoptylaor/paypay/alipay/consts"
	"github.com/Bishoptylaor/paypay/alipay/flow"
)

type QrcodeCaller struct {
}

func NewQrcodeCaller(a *aliClient.Client) flow.QrcodePay {
	// do some implantation
	md := QrcodeCaller{}
	aliClient.SetChecker(md.payloadChecker())(a)
	aliClient.SetPayloadPreSetter(map[string][]alipay.PayloadPreSetter{
		"alipay.trade.precreate": []alipay.PayloadPreSetter{
			preSetter("product_code", "QR_CODE_OFFLINE"),
		},
	})(a)
	return a
}

func (QrcodeCaller) payloadChecker() alipay.PayloadRuler {
	return func(caller string) []string {
		switch caller {
		case "alipay.trade.precreate":
			return []string{
				"out_trade_no != nil", consts.TotalAmountDefaultRule, "subject != nil", `product_code == "QR_CODE_OFFLINE"`}
		case "alipay.trade.close":
			return []string{
				consts.TradeNo2in1DefaultRule,
			}
		case "alipay.trade.fastpay.refund.query":
			return []string{
				consts.TradeNo2in1DefaultRule,
				"out_request_no != nil",
			}
		case "alipay.trade.cancel":
			return []string{
				consts.TradeNo2in1DefaultRule,
			}
		case "alipay.trade.query":
			return []string{
				consts.TradeNo2in1DefaultRule,
			}
		case "alipay.trade.refund":
			return []string{
				"refund_amount != nil",
				consts.TradeNo2in1DefaultRule,
				"refund_reason != nil",
			}
		case "alipay.trade.refund.depositback.completed":
			return []string{
				"trade_no != nil",
				"out_trade_no != nil",
				"out_request_no != nil",
				"dback_status != nil",
				"dback_amount != nil",
			}

		case "alipay.data.dataservice.bill.downloadurl.query":
			return []string{
				`bill_type in ["trade", "signcustomer", "merchant_act", "trade_zft_merchant", "zft_acc", "settlementMerge"]`,
				"bill_date != nil",
			}
		default:
			return []string{}
		}
	}
}
