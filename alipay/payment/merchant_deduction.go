package payment

import (
	"fmt"
	"github.com/Bishoptylaor/paypay"
	"github.com/Bishoptylaor/paypay/alipay/aliClient"
	"github.com/Bishoptylaor/paypay/alipay/consts"
	"github.com/Bishoptylaor/paypay/alipay/flow"
)

type MerchantDeductionProxy interface {
	flow.MerchantDeduction
	InjectCustomPayloadCheckRuler(custom map[string][]paypay.Ruler)
}

type MerchantDeductionCaller struct {
	*aliClient.Client
	rulersMap map[string][]paypay.Ruler
}

// NewMerchantDeductionCaller
//
// 初始化 商家扣款 支付相关接口功能
//
// 提供预设参数与参数校验能力
//
// 在保证参数正确的情况下，用户也可直接调用 client 中的相关接口实现
//
// 产品介绍 https://opendocs.alipay.com/open/06de8c?pathHash=654eb816
func NewMerchantDeductionCaller(c *aliClient.Client) MerchantDeductionProxy {
	// do some implantation
	caller := MerchantDeductionCaller{Client: c}
	// 设置本产品相关接口默认参数校验规则
	caller.rulersMap = caller.setDefaultPayloadCheckRuler()
	aliClient.SetChecker(caller.payloadChecker())(c)
	aliClient.SetPayloadPreSetter(map[string][]paypay.PayloadPreSetter{
		"alipay.trade.app.pay": []paypay.PayloadPreSetter{
			paypay.PreSetter("product_code", "GENERAL_WITHHOLDING"),
		},
		"alipay.user.agreement.query": []paypay.PayloadPreSetter{
			paypay.PreSetter("personal_product_code", "CYCLE_PAY_AUTH_P"),
		},
	})(c)
	return caller
}

// InjectCustomPayloadCheckRuler 外部可调用，用于自定义参数校验规则
func (c MerchantDeductionCaller) InjectCustomPayloadCheckRuler(custom map[string][]paypay.Ruler) {
	if c.rulersMap == nil {
		c.rulersMap = c.setDefaultPayloadCheckRuler()
	}
	for k, v := range custom {
		c.rulersMap[k] = v
	}
}

func (c MerchantDeductionCaller) payloadChecker() paypay.PayloadRuler {
	return func(method string) []paypay.Ruler {
		if rulers, ok := c.rulersMap[method]; ok {
			return rulers
		} else {
			return []paypay.Ruler{}
		}
	}
}

func (c MerchantDeductionCaller) setDefaultPayloadCheckRuler() map[string][]paypay.Ruler {
	return map[string][]paypay.Ruler{
		"alipay.trade.pay": []paypay.Ruler{
			paypay.NewRuler("商户订单号", "out_trade_no != nil && len(out_trade_no) <= 64", fmt.Sprintf(consts.FmtEmptyAlert, "out_trade_no")),
			consts.TotalAmountDefaultRuler,
			paypay.NewRuler("订单标题", "subject != nil", fmt.Sprintf(consts.FmtEmptyAlert, "subject")),
			paypay.NewRuler("销售产品码，商家和支付宝签约的产品码", `product_code == "GENERAL_WITHHOLDING"`, "当前场景 product_code 传值必为 GENERAL_WITHHOLDING"),
			paypay.NewRuler("代扣信息",
				"agreement_params != nil && agreement_params.agreement_no != nil",
				"代扣信息不能为空; 代扣协议号不能为空",
			),
		},
		"alipay.trade.app.pay": []paypay.Ruler{
			paypay.NewRuler("商户订单号", "out_trade_no != nil && len(out_trade_no) <= 64", fmt.Sprintf(consts.FmtEmptyAlert, "out_trade_no")),
			consts.TotalAmountDefaultRuler,
			paypay.NewRuler("订单标题", "subject != nil", fmt.Sprintf(consts.FmtEmptyAlert, "subject")),
			paypay.NewRuler("销售产品码",
				`product_code in ["GENERAL_WITHHOLDING", "QUICK_MSECURITY_PAY"]`,
				fmt.Sprintf(consts.FmtWithinAlert, "product_code", `GENERAL_WITHHOLDING: 商家扣款场景; QUICK_MSECURITY_PAY: 无线快捷支付`),
			),
			paypay.NewRuler("签约参数",
				`agreement_sign_params?.product_code == "GENERAL_WITHHOLDING" &&
agreement_sign_params?.personal_product_code == "CYCLE_PAY_AUTH_P" &&
agreement_sign_params?.sign_scene != nil &&
agreement_sign_params?.access_params != nil &&
agreement_sign_params?.access_params.channel in ["ALIPAYAPP", "QRCODE", "QRCODEORSMS"] &&
agreement_sign_params?.period_rule_params != nil &&
agreement_sign_params?.period_rule_params.period_type in ["DAY", "MONTH"] &&
agreement_sign_params?.period_rule_params.period != nil &&
agreement_sign_params?.period_rule_params.execute_time != nil &&
agreement_sign_params?.period_rule_params.single_amount != nil`,
				"签约产品码: 固定值 GENERAL_WITHHOLDING; \n个人签约产品码: 固定值 CYCLE_PAY_AUTH_P; \n sign_scene 不为空; \n"+
					"sign_scene 详见 https://b.alipay.com/page/product-mall/all-product 商家扣款 > 功能管理 > 修改 > 设置模版 商家自行配置",
			),
			paypay.NewRuler("商户签约号 + 协议签约场景",
				`agreement_sign_params?.external_agreement_no == nil || (agreement_sign_params?.external_agreement_no != nil && len(agreement_sign_params?.external_agreement_no) <= 32 && agreement_sign_params?.sign_scene != nil && agreement_sign_params?.sign_scene != "DEFAULT|DEFAULT")`,
				"当传入商户签约号 external_agreement_no 时，本参数必填，不能为默认值 DEFAULT|DEFAULT。scene 参数根据常见场景值选择，也可在商户商家扣款产品自定义", // https://opendocs.alipay.com/open/20190319114403226822/signscene
			),
			paypay.NewRuler("签约参数关联校验",
				`agreement_sign_params?.period_rule_params.period_type != "MONTH" || 
(agreement_sign_params?.period_rule_params.period_type == "MONTH" && date(agreement_sign_params?.period_rule_params.execute_time, "2006-01-02").Day() <= 28)`,
				"period_type = MONTH 的时候，execute_time 不能大于 28 日",
			),
			paypay.NewRuler("签约参数关联校验",
				`agreement_sign_params?.period_rule_params.period_type != "DAY" || 
(agreement_sign_params?.period_rule_params.period_type == "DAY" && int(agreement_sign_params?.period_rule_params.period) >= 7)`,
				"period_type = DAY 的时候，period 周期小于 7 天，最小周期不能小于 7 天",
			),
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
		"alipay.trade.close": []paypay.Ruler{
			consts.TradeNo2in1DefaultRuler,
		},
		"alipay.trade.cancel": []paypay.Ruler{
			consts.TradeNo2in1DefaultRuler,
		},

		"alipay.user.agreement.page.sign": []paypay.Ruler{
			paypay.NewRuler("个人签约产品码", `personal_product_code == "CYCLE_PAY_AUTH_P"`, "当前场景 personal_product_code 传值必为 CYCLE_PAY_AUTH_P"),
			paypay.NewRuler("销售产品码，商家和支付宝签约的产品码", `product_code == "GENERAL_WITHHOLDING"`, "当前场景 product_code 传值必为 GENERAL_WITHHOLDING"),
			paypay.NewRuler("",
				`access_params != nil && access_params.channel in ["ALIPAYAPP", "QRCODE", "QRCODEORSMS"]`,
				"1. ALIPAYAPP （钱包h5页面签约） 2. QRCODE(扫码签约) 3. QRCODEORSMS(扫码签约或者短信签约)",
			),
			paypay.NewRuler("周期管控规则参数",
				`period_rule_params?.period_type in ["DAY", "MONTH"] &&`+
					`period_rule_params?.period != nil &&`+
					`period_rule_params?.execute_time != nil &&`+
					`period_rule_params?.single_amount != nil`,
				"周期类型 period_type 是周期扣款产品必填，枚举值为 DAY 和 MONTH。周期数 period。首次执行时间 execute_time 格式为 yyyy-MM-dd 日期。single_amount 单次扣款最大金额",
			),
			paypay.NewRuler("商户签约号 + 协议签约场景",
				`external_agreement_no == nil || (external_agreement_no != nil && sign_scene != nil && sign_scene != "DEFAULT|DEFAULT")`,
				"当传入商户签约号 external_agreement_no 时，本参数必填，不能为默认值 DEFAULT|DEFAULT。scene 参数根据常见场景值选择，也可自定义", // https://opendocs.alipay.com/open/20190319114403226822/signscene
			),
			paypay.NewRuler("签约参数关联校验",
				`period_rule_params.period_type != "MONTH" || (period_rule_params.period_type == "MONTH" && date(period_rule_params.execute_time, "2006-01-02").Day() <= 28)`,
				"period_type = MONTH 的时候，execute_time 不能大于 28 日",
			),
			paypay.NewRuler("签约参数关联校验",
				`period_rule_params.period_type != "DAY" || (period_rule_params.period_type == "DAY" && int(period_rule_params.period) >= 7)`,
				"period_type = DAY 的时候，period 周期小于 7 天，最小周期不能小于 7 天",
			),
		},
		"alipay.user.agreement.unsign": []paypay.Ruler{
			paypay.NewRuler("alipay_open_id 支付宝用户号 + alipay_logon_id 支付宝登录账号",
				fmt.Sprintf(consts.Fmt2in1DefaultRule, "alipay_open_id", "alipay_logon_id"),
				"参数必传其一，同时传入优先级 alipay_open_id > alipay_logon_id。推荐使用 alipay_open_id",
			),
			paypay.NewRuler("agreement_no 支付宝系统协议号 + external_agreement_no 商户的协议号",
				fmt.Sprintf(consts.Fmt2in1DefaultRule, "external_agreement_no", "agreement_no"),
				"参数必传其一，同时传入优先级 agreement_no > external_agreement_no",
			),
		},
		"alipay.user.agreement.query": []paypay.Ruler{
			paypay.NewRuler("外部/商户的签约ID", "external_agreement_no != nil", fmt.Sprintf(consts.FmtEmptyAlert, "external_agreement_no")),
			paypay.NewRuler("签约场景值", "sign_scene != nil", fmt.Sprintf(consts.FmtEmptyAlert, "sign_scene")),
			paypay.NewRuler("协议产品码", "personal_product_code != nil", fmt.Sprintf(consts.FmtEmptyAlert, "personal_product_code")),
		},
		"alipay.user.agreement.executionplan.modify": []paypay.Ruler{
			paypay.NewRuler("", "agreement_no != nil", fmt.Sprintf(consts.FmtEmptyAlert, "agreement_no")),
			paypay.NewRuler("", "deduct_time != nil", fmt.Sprintf(consts.FmtEmptyAlert, "deduct_time")),
		},

		"alipay.data.dataservice.bill.downloadurl.query": consts.DataDownloadRuler,
	}
}
