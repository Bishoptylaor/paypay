package payment

import (
	"fmt"
	"github.com/Bishoptylaor/paypay"
	"github.com/Bishoptylaor/paypay/alipay/aliClient"
	"github.com/Bishoptylaor/paypay/alipay/consts"
	"github.com/Bishoptylaor/paypay/alipay/flow"
)

type MerchantLedgerCaller struct {
}

// NewMerchantLedgerCaller
//
// 初始化 商家分账 相关接口功能
//
// 提供预设参数与参数校验能力
//
// 在保证参数正确的情况下，用户也可直接调用 client 中的相关接口实现
//
// 产品介绍 https://opendocs.alipay.com/open/06de8c?pathHash=654eb816
func NewMerchantLedgerCaller(a *aliClient.Client) flow.MerchantLedger {
	// do some implantation
	md := MerchantLedgerCaller{}
	aliClient.SetChecker(md.payloadChecker())(a)
	aliClient.SetPayloadPreSetter(map[string][]paypay.PayloadPreSetter{
		"alipay.trade.app.pay": []paypay.PayloadPreSetter{
			paypay.PreSetter("product_code", "GENERAL_WITHHOLDING"),
		},
		"alipay.user.agreement.query": []paypay.PayloadPreSetter{
			paypay.PreSetter("personal_product_code", "CYCLE_PAY_AUTH_P"),
		},
	})(a)
	return a
}

func (MerchantLedgerCaller) payloadChecker() paypay.PayloadRuler {
	return func(caller string) []paypay.Ruler {
		switch caller {
		case "alipay.trade.royalty.relation.bind", "alipay.trade.royalty.relation.unbind":
			return []paypay.Ruler{
				paypay.NewRuler("分账接收方列表",
					`receiver_list != nil && len(receiver_list) < 20 && `+
						`all(receiver_list, {.type in ["userId", "loginName", "openId"]}) && `+
						`none(receiver_list, {.type == "loginName" && .name == nil})`,
					"receiver_list 不为空 & 元素个数小于 20 &\nreceiver_list 所有元素中 type 字段取值范围为：\"userId\", \"loginName\", \"openId\"\nreceiver_list 所有元素中 如果 type == loginName 则 name 字段不能为空",
				),
				paypay.NewRuler("外部请求号，由商家自定义", "out_request_no != nil && len(out_request_no) <= 32", fmt.Sprintf(consts.FmtEmptyAlert, "out_request_no")),
			}
		case "alipay.trade.royalty.relation.batchquery":
			return []paypay.Ruler{
				paypay.NewRuler("外部请求号，由商家自定义", "out_request_no != nil && len(out_request_no) <= 32", fmt.Sprintf(consts.FmtEmptyAlert, "out_request_no")),
			}
		case "alipay.trade.order.settle":
			return []paypay.Ruler{
				paypay.NewRuler("外部请求号，由商家自定义", "out_request_no != nil && len(out_request_no) <= 32", fmt.Sprintf(consts.FmtEmptyAlert, "out_request_no")),
				paypay.NewRuler("外部请求号，由商家自定义", "trade_no != nil && len(trade_no) <= 64", fmt.Sprintf(consts.FmtEmptyAlert, "trade_no")),
				paypay.NewRuler("分账明细信息",
					`royalty_parameters == nil || `+
						`(`+
						`none(royalty_parameters, {.trans_out_type == "userId" && (!hasPrefix(.trans_out, "2088") || !hasPrefix(.trans_in, "2088"))})`+
						`)`,
					"字段不传 或者 trans_out_type == userId 时 trans_out 和 trans_in 必须 2088 开头",
				),
			}
		case "alipay.trade.royalty.rate.query":
			return []paypay.Ruler{
				paypay.NewRuler("外部请求号，由商家自定义", "out_request_no != nil && len(out_request_no) <= 32", fmt.Sprintf(consts.FmtEmptyAlert, "out_request_no")),
			}
		case "alipay.trade.order.onsettle.query":
			return []paypay.Ruler{
				paypay.NewRuler("外部请求号，由商家自定义", "trade_no != nil && len(trade_no) <= 64", fmt.Sprintf(consts.FmtEmptyAlert, "trade_no")),
			}
		case "alipay.trade.order.settle.query":
			return []paypay.Ruler{
				paypay.NewRuler("",
					"(settle_no != nil && len(settle_no) <= 64) || (settle_no == nil && (out_request_no != nil && trade_no != nil))",
					"传入 settle_no, 无需再传 out_request_no 和 trade_no; 不传入 settle_no, out_request_no 和 trade_no 需要一起传入",
				),
			}
		default:
			return []paypay.Ruler{}
		}
	}
}
