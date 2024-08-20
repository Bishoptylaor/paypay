package payment

import (
	"github.com/Bishoptylaor/paypay/alipay"
	"github.com/Bishoptylaor/paypay/alipay/aliClient"
	"github.com/Bishoptylaor/paypay/alipay/consts"
	"github.com/Bishoptylaor/paypay/alipay/flow"
)

type SmileCaller struct {
}

func NewSmileCaller(a *aliClient.Client) flow.SmilePay {
	// do some implantation
	md := SmileCaller{}
	aliClient.SetChecker(md.payloadChecker())(a)
	return a
}

func (SmileCaller) payloadChecker() alipay.PayloadRuler {
	return func(caller string) []string {
		switch caller {
		case "zoloz.authentication.smilepay.initialize":
			return []string{}
		case "zoloz.authentication.customer.ftoken.query":
			return []string{
				"ftoken != nil",
				consts.TradeNo2in1DefaultRule,
				`biz_type in ["1", "2", "3", "4"]`, // 1、1：1人脸验证能力 2、1：n人脸搜索能力（支付宝uid入库） 3、1：n人脸搜索能力（支付宝手机号入库） 4、手机号和人脸识别综合能力
			}
		default:
			return []string{}
		}
	}
}
