package payment

import (
	"fmt"
	"github.com/Bishoptylaor/paypay"
	"github.com/Bishoptylaor/paypay/alipay"
	"github.com/Bishoptylaor/paypay/alipay/consts"
	"github.com/Bishoptylaor/paypay/alipay/flow"
	"github.com/Bishoptylaor/paypay/alipay/service"
)

type PreLicensingCaller interface {
	flow.PreLicensingPay
	UseCustomPayloadCheckRuler(custom map[string][]paypay.Ruler)
}

type PreLicensingService service.Service

// NewPreLicensingCaller
//
// 初始化 预授权 支付相关接口功能
//
// 提供预设参数与参数校验能力
//
// 在保证参数正确的情况下，用户也可直接调用 client 中的相关接口实现
//
// 产品介绍 https://opendocs.alipay.com/open/06de96?pathHash=f6dfbf6f
func NewPreLicensingCaller(c *alipay.Client) PreLicensingCaller {
	// do some implantation
	caller := &PreLicensingService{Client: c}
	// 设置本产品相关接口默认参数校验规则
	caller.RulersMap = caller.setDefaultPayloadCheckRuler()
	alipay.Checker(caller.payloadChecker())(c)
	alipay.PayloadPreSetter(map[string][]paypay.PayloadPreSetter{
		"alipay.fund.auth.order.app.freeze": []paypay.PayloadPreSetter{
			paypay.PreSetter("product_code", "PREAUTH_PAY"),
		},
		"alipay.fund.auth.order.voucher.create": []paypay.PayloadPreSetter{
			paypay.PreSetter("product_code", "PREAUTH_PAY"),
		},
	})(c)
	return caller
}

// UseCustomPayloadCheckRuler 外部可调用，用于自定义参数校验规则
func (c PreLicensingService) UseCustomPayloadCheckRuler(custom map[string][]paypay.Ruler) {
	if c.RulersMap == nil {
		c.RulersMap = c.setDefaultPayloadCheckRuler()
	}
	for k, v := range custom {
		c.RulersMap[k] = v
	}
}

func (c PreLicensingService) payloadChecker() paypay.PayloadRuler {
	return func(method string) []paypay.Ruler {
		if rulers, ok := c.RulersMap[method]; ok {
			return rulers
		} else {
			return []paypay.Ruler{}
		}
	}
}

func (c PreLicensingService) setDefaultPayloadCheckRuler() map[string][]paypay.Ruler {
	return map[string][]paypay.Ruler{
		"alipay.fund.auth.order.app.freeze": []paypay.Ruler{
			paypay.NewRuler("商户授权资金订单号", "out_order_no != nil", fmt.Sprintf(consts.FmtEmptyAlert, "out_order_no")),
			paypay.NewRuler("商户本次资金操作的请求流水号，用于标示请求流水的唯一性。可与out_order_no相同，仅支持字母、数字、下划线。",
				"out_request_no != nil",
				fmt.Sprintf(consts.FmtEmptyAlert, "out_request_no"),
			),
			paypay.NewRuler("订单标题", "order_title != nil", fmt.Sprintf(consts.FmtEmptyAlert, "order_title")),
			paypay.NewRuler("需要冻结的金额", fmt.Sprintf(consts.FmtAmountDefaultRule, "amount"), "请检查价格金额区间是否在 0.01 ~ 100000000"),
			paypay.NewRuler("销售产品码，商家和支付宝签约的产品码", `product_code == "PREAUTH_PAY"`, "当前场景 product_code 传值必为 PREAUTH_PAY"),
			paypay.NewRuler("免押受理台模式",
				`deposit_product_mode == nil || deposit_product_mode in ["POSTPAY", "POSTPAY_UNCERTAIN", "DEPOSIT_ONLY"]`, // https://opendocs.alipay.com/b/08tf3t?pathHash=d67d7545
				"使用免押产品必传该字段。后付金额已知: POSTPAY\n后付金额未知: POSTPAY_UNCERTAIN\n纯免押: DEPOSIT_ONLY",
			),
			paypay.NewRuler("后付费项目",
				`deposit_product_mode == "POSTPAY" && post_payments?.name != nil && post_payments?.amount != nil || `+
					`deposit_product_mode == "POSTPAY_UNCERTAIN" && post_payments?.name != nil && post_payments?.description != nil && post_payments?.amount == nil`, // https://opendocs.alipay.com/b/08tf3t?pathHash=d67d7545
				"当受理台模式（deposit_product_mode）传入POSTPAY 时，后付费项目名称（name）、金额（amount）必传；当传入 POSTPAY_UNCERTAIN 时，后付费项目名称（name）、计费说明（description）必传，金额（amount）不传。",
			),
		},
		"alipay.fund.auth.operation.detail.query": []paypay.Ruler{
			// 详情查看 https://opendocs.alipay.com/open/064jhg?scene=common&pathHash=44be9c20
			paypay.NewRuler("auth_no 支付宝授权资金订单号 + out_order_no 商户的授权资金订单号",
				fmt.Sprintf(consts.Fmt2in1DefaultRule, "auth_no", "out_order_no"),
				"参数必传其一，同时传入优先级 auth_no > out_order_no",
			),
			paypay.NewRuler("operation_id 支付宝的授权资金操作流水号 + out_request_no 商户的授权资金操作流水号",
				fmt.Sprintf(consts.Fmt2in1DefaultRule, "operation_id", "out_request_no"),
				"参数必传其一，同时传入优先级 operation_id > out_request_no",
			),
		},
		"alipay.fund.auth.operation.cancel": []paypay.Ruler{
			paypay.NewRuler("商户对本次撤销操作的附言描述。", "remark != nil", fmt.Sprintf(consts.FmtEmptyAlert, "remark")),
			paypay.NewRuler("auth_no 支付宝授权资金订单号 + out_order_no 商户的授权资金订单号",
				fmt.Sprintf(consts.Fmt2in1DefaultRule, "auth_no", "out_order_no"),
				"参数必传其一，同时传入优先级 auth_no > out_order_no",
			),
			paypay.NewRuler("operation_id 支付宝的授权资金操作流水号 + out_request_no 商户的授权资金操作流水号",
				fmt.Sprintf(consts.Fmt2in1DefaultRule, "operation_id", "out_request_no"),
				"参数必传其一，同时传入优先级 operation_id > out_request_no",
			),
		},
		"alipay.fund.auth.order.unfreeze": []paypay.Ruler{
			paypay.NewRuler("支付宝资金授权订单号。", "auth_no != nil", fmt.Sprintf(consts.FmtEmptyAlert, "auth_no")),
			paypay.NewRuler("解冻请求流水号。", "out_request_no != nil", fmt.Sprintf(consts.FmtEmptyAlert, "out_request_no")),
			paypay.NewRuler("本次操作解冻的金额。", fmt.Sprintf(consts.FmtAmountDefaultRule, "amount"), "请检查价格金额区间是否在 0.01 ~ 100000000"),
			paypay.NewRuler("商户对本次撤销操作的附言描述。", "remark != nil", fmt.Sprintf(consts.FmtEmptyAlert, "remark")),
		},
		"alipay.fund.auth.order.voucher.create": []paypay.Ruler{
			paypay.NewRuler("商户授权资金订单号。", "out_order_no != nil", fmt.Sprintf(consts.FmtEmptyAlert, "out_order_no")),
			paypay.NewRuler("商户本次资金操作的请求流水号，用于标示请求流水的唯一性。", "out_request_no != nil", fmt.Sprintf(consts.FmtEmptyAlert, "out_request_no")),
			paypay.NewRuler("订单标题。", "order_title != nil", fmt.Sprintf(consts.FmtEmptyAlert, "order_title")),
			paypay.NewRuler("需要冻结的金额。", fmt.Sprintf(consts.FmtAmountDefaultRule, "amount"), "请检查价格金额区间是否在 0.01 ~ 100000000"),
			paypay.NewRuler("销售产品码，商家和支付宝签约的产品码", `product_code == "PREAUTH_PAY"`, "当前场景 product_code 传值必为 PREAUTH_PAY"),
			paypay.NewRuler("免押受理台模式",
				`deposit_product_mode == nil || deposit_product_mode in ["POSTPAY", "POSTPAY_UNCERTAIN", "DEPOSIT_ONLY"]`, // https://opendocs.alipay.com/b/08tf3t?pathHash=d67d7545
				"使用免押产品必传该字段。后付金额已知: POSTPAY\n后付金额未知: POSTPAY_UNCERTAIN\n纯免押: DEPOSIT_ONLY",
			),
			paypay.NewRuler("后付费项目",
				`deposit_product_mode == "POSTPAY" && post_payments?.name != nil && post_payments?.amount != nil || `+
					`deposit_product_mode == "POSTPAY_UNCERTAIN" && post_payments?.name != nil && post_payments?.description != nil && post_payments?.amount == nil`, // https://opendocs.alipay.com/b/08tf3t?pathHash=d67d7545
				"当受理台模式（deposit_product_mode）传入POSTPAY 时，后付费项目名称（name）、金额（amount）必传；当传入 POSTPAY_UNCERTAIN 时，后付费项目名称（name）、计费说明（description）必传，金额（amount）不传。",
			),
		},
		"alipay.fund.auth.order.freeze": []paypay.Ruler{
			paypay.NewRuler("用户付款码。", "auth_code != nil", fmt.Sprintf(consts.FmtEmptyAlert, "auth_code")),
			paypay.NewRuler("付款码类型。",
				`auth_code_type in ["bar_code", "security_code"]`,
				fmt.Sprintf(consts.FmtWithinAlert, "auth_code_type", `条码场景: bar_code; 刷脸场景: security_code`),
			),
			paypay.NewRuler("商户授权资金订单号。", "out_order_no != nil", fmt.Sprintf(consts.FmtEmptyAlert, "out_order_no")),
			paypay.NewRuler("商户本次资金操作的请求流水号，用于标示请求流水的唯一性。", "out_request_no != nil", fmt.Sprintf(consts.FmtEmptyAlert, "out_request_no")),
			paypay.NewRuler("订单标题。", "order_title != nil", fmt.Sprintf(consts.FmtEmptyAlert, "order_title")),
			paypay.NewRuler("需要冻结的金额。", fmt.Sprintf(consts.FmtAmountDefaultRule, "amount"), "请检查价格金额区间是否在 0.01 ~ 100000000"),
			paypay.NewRuler("销售产品码，商家和支付宝签约的产品码", `product_code == "PREAUTH_PAY"`, "当前场景 product_code 传值必为 PREAUTH_PAY"),
		},

		"alipay.trade.pay": []paypay.Ruler{
			paypay.NewRuler("商户订单号", "out_trade_no != nil && len(out_trade_no) <= 64", fmt.Sprintf(consts.FmtEmptyAlert, "out_trade_no")),
			consts.TotalAmountDefaultRuler,
			paypay.NewRuler("订单标题", "subject != nil", fmt.Sprintf(consts.FmtEmptyAlert, "subject")),
			paypay.NewRuler("销售产品码，商家和支付宝签约的产品码", `product_code == "PREAUTH_PAY"`, "当前场景 product_code 传值必为 PREAUTH_PAY"),
			paypay.NewRuler("资金预授权单号", "auth_no != nil", fmt.Sprintf(consts.FmtEmptyAlert, "auth_no")),
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
		"alipay.trade.refund": []paypay.Ruler{
			paypay.NewRuler("退款金额", "refund_amount != nil", fmt.Sprintf(consts.FmtEmptyAlert, "refund_amount")),
			consts.TradeNo2in1DefaultRuler,
			paypay.NewRuler("退款原因说明", "refund_reason != nil", fmt.Sprintf(consts.FmtEmptyAlert, "refund_reason")),
			// 一般来说这个字段只有在部分退款时必传，这里进行一般性校验，保证调用方业务一致性
			paypay.NewRuler("退款请求号", "out_request_no != nil", fmt.Sprintf(consts.FmtEmptyAlert, "out_request_no")),
		},
		"alipay.trade.orderinfo.sync": []paypay.Ruler{
			paypay.NewRuler("支付宝交易号", "trade_no != nil", fmt.Sprintf(consts.FmtEmptyAlert, "trade_no")),
			paypay.NewRuler("外部请求号", "out_request_no != nil", fmt.Sprintf(consts.FmtEmptyAlert, "out_request_no")),
			paypay.NewRuler("交易信息同步对应的业务类型",
				`biz_type in ["CREDIT_AUTH", "CREDIT_DEDUCT"]`, // https://opendocs.alipay.com/b/08tf3t?pathHash=d67d7545
				"信用授权场景下传CREDIT_AUTH 信用代扣场景下传CREDIT_DEDUCT",
			),
		},

		"alipay.data.dataservice.bill.downloadurl.query": consts.DataDownloadRuler,
	}
}
