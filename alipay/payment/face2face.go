package payment

import (
	"fmt"
	"github.com/Bishoptylaor/paypay"
	"github.com/Bishoptylaor/paypay/alipay/aliClient"
	"github.com/Bishoptylaor/paypay/alipay/consts"
	"github.com/Bishoptylaor/paypay/alipay/flow"
)

type Face2FacePayProxy interface {
	flow.Face2FacePay
	InjectCustomPayloadCheckRuler(custom map[string][]paypay.Ruler)
}

type Face2FaceCaller struct {
	*aliClient.Client
	rulersMap map[string][]paypay.Ruler
}

// NewFace2FaceCaller
//
// 初始化 当面付 支付相关接口功能
//
// 提供预设参数与参数校验能力
//
// 在保证参数正确的情况下，用户也可直接调用 client 中的相关接口实现
//
// 产品介绍 https://opendocs.alipay.com/open/194/105072?pathHash=45357796
func NewFace2FaceCaller(c *aliClient.Client) Face2FacePayProxy {
	// do some implantation
	caller := Face2FaceCaller{Client: c}
	// 设置本产品相关接口默认参数校验规则
	caller.rulersMap = caller.setDefaultPayloadCheckRuler()
	aliClient.SetChecker(caller.payloadChecker())(c)
	aliClient.SetPayloadPreSetter(map[string][]paypay.PayloadPreSetter{
		"alipay.trade.create": []paypay.PayloadPreSetter{
			paypay.PreSetter("product_code", "FACE_TO_FACE_PAYMENT"),
		},
	})(c)
	return caller
}

// InjectCustomPayloadCheckRuler 外部可调用，用于自定义参数校验规则
func (c Face2FaceCaller) InjectCustomPayloadCheckRuler(custom map[string][]paypay.Ruler) {
	if c.rulersMap == nil {
		c.rulersMap = c.setDefaultPayloadCheckRuler()
	}
	for k, v := range custom {
		c.rulersMap[k] = v
	}
}

func (c Face2FaceCaller) payloadChecker() paypay.PayloadRuler {
	return func(method string) []paypay.Ruler {
		if rulers, ok := c.rulersMap[method]; ok {
			return rulers
		} else {
			return []paypay.Ruler{}
		}
	}
}

func (c Face2FaceCaller) setDefaultPayloadCheckRuler() map[string][]paypay.Ruler {
	return map[string][]paypay.Ruler{
		"alipay.trade.pay": []paypay.Ruler{
			paypay.NewRuler("商户订单号", "out_trade_no != nil && len(out_trade_no) <= 64", fmt.Sprintf(consts.FmtEmptyAlert, "out_trade_no")),
			consts.TotalAmountDefaultRuler,
			paypay.NewRuler("订单标题", "subject != nil", fmt.Sprintf(consts.FmtEmptyAlert, "subject")),
			paypay.NewRuler("支付授权码", "auth_code != nil", fmt.Sprintf(consts.FmtEmptyAlert, "auth_code")),
			paypay.NewRuler("支付场景",
				`scene in ["bar_code", "security_code"]`,
				fmt.Sprintf(consts.FmtWithinAlert, "scene", `bar_code: 当面付条码支付场景/默认; security_code: 当面付刷脸支付场景`),
			),
			paypay.NewRuler("支付场景 + 授权码联合校验", `scene == "security_code" && hasPrefix(auth_code, "fp")`, "当面付刷脸支付场景，对应的 auth_code 为 fp 开头的刷脸标识串"),
			paypay.NewRuler("销售产品码",
				`product_code == nil || product_code in ["OFFLINE_PAYMENT", "FACE_TO_FACE_PAYMENT"]`,
				fmt.Sprintf(consts.FmtWithinAlert, "product_code", `OFFLINE_PAYMENT: 当面付快捷版; FACE_TO_FACE_PAYMENT: 其它支付宝当面付产品/默认`),
			),
		},
		"alipay.trade.precreate": []paypay.Ruler{
			paypay.NewRuler("商户订单号", "out_trade_no != nil && len(out_trade_no) <= 64", fmt.Sprintf(consts.FmtEmptyAlert, "out_trade_no")),
			consts.TotalAmountDefaultRuler,
			paypay.NewRuler("订单标题", "subject != nil", fmt.Sprintf(consts.FmtEmptyAlert, "subject")),
			paypay.NewRuler("销售产品码",
				`product_code in ["OFFLINE_PAYMENT", "FACE_TO_FACE_PAYMENT"]`,
				fmt.Sprintf(consts.FmtWithinAlert, "product_code", `OFFLINE_PAYMENT: 当面付快捷版; FACE_TO_FACE_PAYMENT: 其它支付宝当面付产品/默认`),
			),
		},
		"alipay.trade.create": []paypay.Ruler{
			paypay.NewRuler("商户订单号", "out_trade_no != nil && len(out_trade_no) <= 64", fmt.Sprintf(consts.FmtEmptyAlert, "out_trade_no")),
			consts.TotalAmountDefaultRuler,
			paypay.NewRuler("订单标题", "subject != nil", fmt.Sprintf(consts.FmtEmptyAlert, "subject")),
			paypay.NewRuler("销售产品码", `product_code == "FACE_TO_FACE_PAYMENT"`,
				fmt.Sprintf(consts.FmtWithinAlert, "product_code", `["FACE_TO_FACE_PAYMENT"]`)),
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
		"alipay.trade.fastpay.refund.query": []paypay.Ruler{
			consts.TradeNo2in1DefaultRuler,
			paypay.NewRuler("退款请求号", "out_request_no != nil", fmt.Sprintf(consts.FmtEmptyAlert, "out_request_no")),
		},
		"alipay.trade.cancel": []paypay.Ruler{
			consts.TradeNo2in1DefaultRuler,
		},
		"alipay.trade.close": []paypay.Ruler{
			consts.TradeNo2in1DefaultRuler,
		},
		"alipay.data.dataservice.bill.downloadurl.query": consts.DataDownloadRuler,
	}
}
