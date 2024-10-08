package payment

import (
	"fmt"
	"github.com/Bishoptylaor/paypay"
	"github.com/Bishoptylaor/paypay/alipay/aliClient"
	"github.com/Bishoptylaor/paypay/alipay/consts"
	"github.com/Bishoptylaor/paypay/alipay/flow"
)

type SmileProxy interface {
	flow.SmilePay
	InjectCustomPayloadCheckRuler(custom map[string][]paypay.Ruler)
}

type SmileCaller struct {
	*aliClient.Client
	rulersMap map[string][]paypay.Ruler
}

// NewSmileCaller
//
// 初始化 刷脸 支付相关接口功能
//
// 提供预设参数与参数校验能力
//
// 在保证参数正确的情况下，用户也可直接调用 client 中的相关接口实现
//
// 产品介绍 https://opendocs.alipay.com/open/20180402104715814204/intro?pathHash=5b9e1a85
func NewSmileCaller(c *aliClient.Client) SmileProxy {
	// do some implantation
	caller := &SmileCaller{Client: c}
	// 设置本产品相关接口默认参数校验规则
	caller.rulersMap = caller.setDefaultPayloadCheckRuler()
	aliClient.SetChecker(caller.payloadChecker())(c)
	return caller
}

// InjectCustomPayloadCheckRuler 外部可调用，用于自定义参数校验规则
func (c SmileCaller) InjectCustomPayloadCheckRuler(custom map[string][]paypay.Ruler) {
	if c.rulersMap == nil {
		c.rulersMap = c.setDefaultPayloadCheckRuler()
	}
	for k, v := range custom {
		c.rulersMap[k] = v
	}
}

func (c SmileCaller) payloadChecker() paypay.PayloadRuler {
	return func(method string) []paypay.Ruler {
		if rulers, ok := c.rulersMap[method]; ok {
			return rulers
		} else {
			return []paypay.Ruler{}
		}
	}
}

func (c SmileCaller) setDefaultPayloadCheckRuler() map[string][]paypay.Ruler {
	return map[string][]paypay.Ruler{
		"zoloz.authentication.smilepay.initialize": []paypay.Ruler{},
		"zoloz.authentication.customer.ftoken.query": []paypay.Ruler{
			paypay.NewRuler("人脸token", "ftoken != nil", fmt.Sprintf(consts.FmtEmptyAlert, "ftoken")),
			paypay.NewRuler("识别能力",
				`biz_type in ["1", "2", "3", "4"]`,
				fmt.Sprintf(consts.FmtWithinAlert, "biz_type", `1、1：1人脸验证能力 2、1：n人脸搜索能力（支付宝uid入库） 3、1：n人脸搜索能力（支付宝手机号入库） 4、手机号和人脸识别综合能力`),
			),
			consts.TradeNo2in1DefaultRuler,
		},
	}
}
