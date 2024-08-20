package payment

import (
	"github.com/Bishoptylaor/paypay/alipay"
	"github.com/Bishoptylaor/paypay/alipay/aliClient"
	"github.com/Bishoptylaor/paypay/alipay/consts"
	"github.com/Bishoptylaor/paypay/alipay/flow"
)

type MerchantDeductionCaller struct {
}

func NewMerchantDeductionCaller(a *aliClient.Client) flow.MerchantDeduction {
	// do some implantation
	md := MerchantDeductionCaller{}
	aliClient.SetChecker(md.payloadChecker())(a)
	aliClient.SetPayloadPreSetter(map[string][]alipay.PayloadPreSetter{
		"alipay.trade.app.pay": []alipay.PayloadPreSetter{
			preSetter("product_code", "GENERAL_WITHHOLDING"),
		},
		"alipay.user.agreement.query": []alipay.PayloadPreSetter{
			preSetter("personal_product_code", "CYCLE_PAY_AUTH_P"),
		},
	})(a)
	return a
}

func (MerchantDeductionCaller) payloadChecker() alipay.PayloadRuler {
	return func(caller string) []string {
		switch caller {
		case "alipay.trade.pay":
			return []string{
				"out_trade_no != nil",
				consts.TotalAmountDefaultRule,
				"subject != nil",
				`product_code == "GENERAL_WITHHOLDING"`,
				"agreement_params != nil && agreement_params.agreement_no != nil",
			}
		case "alipay.trade.app.pay":
			return []string{
				"out_trade_no != nil",
				consts.TotalAmountDefaultRule,
				"subject != nil",
				`product_code in ["GENERAL_WITHHOLDING", "QUICK_MSECURITY_PAY"]`,
				`agreement_sign_params != nil && ` +
					`agreement_sign_params.product_code == "GENERAL_WITHHOLDING" &&` +
					`agreement_sign_params.personal_product_code == "CYCLE_PAY_AUTH_P &&"` +
					`agreement_sign_params.sign_scene != nil &&` +
					`agreement_sign_params.access_params != nil &&` +
					`agreement_sign_params.access_params.channel in [` + consts.ALIPAYAPP + consts.QRCODE + consts.QRCODEORSMS + `] &&` +
					`agreement_sign_params.period_rule_params != nil &&` +
					`agreement_sign_params.period_rule_params.period_type in ["DAY", "MONTH"] &&` +
					`agreement_sign_params.period_rule_params.period != nil &&` +
					`agreement_sign_params.period_rule_params.execute_time != nil &&` +
					`agreement_sign_params.period_rule_params.single_amount != nil &&`,
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
		case "alipay.trade.close":
			return []string{
				consts.TradeNo2in1DefaultRule,
			}
		case "alipay.trade.cancel":
			return []string{
				consts.TradeNo2in1DefaultRule,
				"retry_flag != nil",
			}

		case "alipay.user.agreement.page.sign":
			return []string{
				`personal_product_code == "GENERAL_WITHHOLDING_P"`,
				`product_code == "GENERAL_WITHHOLDING"`,
				`access_params != nil && access_params.channel in [` + consts.ALIPAYAPP + consts.QRCODE + consts.QRCODEORSMS + `]`,
				`period_rule_params != nil &&` +
					`period_rule_params.period_type in ["DAY", "MONTH"] &&` +
					`period_rule_params.period != nil &&` +
					`period_rule_params.execute_time != nil &&` +
					`period_rule_params.single_amount != nil`,
				"sign_scene != nil",
				"external_agreement_no != nil",
			}
		case "alipay.user.agreement.unsign":
			return []string{"agreement_no"}
		case "alipay.user.agreement.query":
			return []string{
				"external_agreement_no != nil",
				"sign_scene != nil",
				"personal_product_code != nil",
			}
		case "alipay.user.agreement.executionplan.modify":
			return []string{
				"agreement_no != nil",
				"deduct_time != nil",
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
