package payment

import (
	"fmt"
	"github.com/Bishoptylaor/paypay"
	"github.com/Bishoptylaor/paypay/alipay/aliClient"
	"github.com/Bishoptylaor/paypay/alipay/consts"
	"github.com/Bishoptylaor/paypay/alipay/flow"
)

type SmileCaller struct{}

// NewSmileCaller
//
// 初始化 刷脸 支付相关接口功能
//
// 提供预设参数与参数校验能力
//
// 在保证参数正确的情况下，用户也可直接调用 client 中的相关接口实现
//
// 产品介绍 https://opendocs.alipay.com/open/20180402104715814204/intro?pathHash=5b9e1a85
func NewSmileCaller(a *aliClient.Client) flow.SmilePay {
	// do some implantation
	md := SmileCaller{}
	aliClient.SetChecker(md.payloadChecker())(a)
	return a
}

func (SmileCaller) payloadChecker() paypay.PayloadRuler {
	return func(caller string) []paypay.Ruler {
		switch caller {
		case "zoloz.authentication.smilepay.initialize":
			return []paypay.Ruler{}
		case "zoloz.authentication.customer.ftoken.query":
			return []paypay.Ruler{
				paypay.NewRuler("人脸token", "ftoken != nil", fmt.Sprintf(consts.FmtEmptyAlert, "ftoken")),
				paypay.NewRuler("识别能力",
					`biz_type in ["1", "2", "3", "4"]`,
					fmt.Sprintf(consts.FmtWithinAlert, "biz_type", `1、1：1人脸验证能力 2、1：n人脸搜索能力（支付宝uid入库） 3、1：n人脸搜索能力（支付宝手机号入库） 4、手机号和人脸识别综合能力`),
				),
				consts.TradeNo2in1DefaultRuler,
			}
		default:
			return []paypay.Ruler{}
		}
	}
}
