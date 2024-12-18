package private_domain

import (
	"fmt"
	"github.com/Bishoptylaor/paypay"
	"github.com/Bishoptylaor/paypay/alipay"
	"github.com/Bishoptylaor/paypay/alipay/consts"
	"github.com/Bishoptylaor/paypay/alipay/flow"
	"github.com/Bishoptylaor/paypay/alipay/service"
)

type PayWithGiftCaller interface {
	flow.PayWithGift
	UseCustomPayloadCheckRuler(custom map[string][]paypay.Ruler)
}

type PayWithGiftService service.Service

// NewPayWithGiftCaller
//
// 初始化 支付有礼 相关接口功能
//
// 提供预设参数与参数校验能力
//
// 在保证参数正确的情况下，用户也可直接调用 client 中的相关接口实现
//
// 产品介绍 https://opendocs.alipay.com/open/03o2f7?pathHash=e2a381af
func NewPayWithGiftCaller(c *alipay.Client) PayWithGiftCaller {
	// do some implantation
	caller := &PayWithGiftService{Client: c}
	// 设置本产品相关接口默认参数校验规则
	caller.RulersMap = caller.setDefaultPayloadCheckRuler()
	alipay.Checker(caller.payloadChecker())(c)
	alipay.PayloadPreSetter(map[string][]paypay.PayloadPreSetter{
		"alipay.trade.app.pay": []paypay.PayloadPreSetter{
			paypay.PreSetter("product_code", "QUICK_MSECURITY_PAY"),
		},
	})(c)
	return caller
}

// UseCustomPayloadCheckRuler 外部可调用，用于自定义参数校验规则
func (c PayWithGiftService) UseCustomPayloadCheckRuler(custom map[string][]paypay.Ruler) {
	if c.RulersMap == nil {
		c.RulersMap = c.setDefaultPayloadCheckRuler()
	}
	for k, v := range custom {
		c.RulersMap[k] = v
	}
}

func (c PayWithGiftService) payloadChecker() paypay.PayloadRuler {
	return func(method string) []paypay.Ruler {
		if rulers, ok := c.RulersMap[method]; ok {
			return rulers
		} else {
			return []paypay.Ruler{}
		}
	}
}

func (c PayWithGiftService) setDefaultPayloadCheckRuler() map[string][]paypay.Ruler {
	return map[string][]paypay.Ruler{
		"alipay.marketing.activity.delivery.stop": []paypay.Ruler{
			paypay.NewRuler("投放计划id", "delivery_id != nil && len(delivery_id) <= 128 && len(delivery_id) >= 1", fmt.Sprintf(consts.FmtEmptyAlert, "delivery_id")),
			paypay.NewRuler("外部业务单号", "out_biz_no != nil && len(delivery_id) <= 64 && len(delivery_id) >= 1", fmt.Sprintf(consts.FmtEmptyAlert, "out_biz_no")),
			paypay.NewRuler("商户接入模式",
				`merchant_access_mode in ["SELF_MODE", "AGENCY_MODE"]`,
				"merchant_access_mode 枚举值为 SELF_MODE, AGENCY_MODE",
			),
		},
		"alipay.marketing.activity.delivery.query": []paypay.Ruler{
			paypay.NewRuler("投放计划id", "delivery_id != nil && len(delivery_id) <= 128 && len(delivery_id) >= 1", fmt.Sprintf(consts.FmtEmptyAlert, "delivery_id")),
			paypay.NewRuler("商户接入模式", `merchant_access_mode in ["SELF_MODE", "AGENCY_MODE"]`, fmt.Sprintf(consts.FmtEmptyAlert, "out_biz_no")),
		},
		"alipay.marketing.activity.delivery.create": []paypay.Ruler{
			paypay.NewRuler("展位编码", `delivery_booth_code == "PAYMENT_RESULT"`, "PAYMENT_RESULT：支付有礼"),
			paypay.NewRuler("外部业务单号", "out_biz_no != nil && len(delivery_id) <= 64 && len(delivery_id) >= 1", fmt.Sprintf(consts.FmtEmptyAlert, "out_biz_no")),
			paypay.NewRuler("商户接入模式",
				`merchant_access_mode in ["SELF_MODE", "AGENCY_MODE"]`,
				"merchant_access_mode 枚举值为 SELF_MODE, AGENCY_MODE",
			),
			paypay.NewRuler("投放计划基础信息",
				"delivery_base_info != nil && "+
					"delivery_base_info.delivery_name != nil && len(delivery_base_info.delivery_name) <= 20 && len(delivery_base_info.delivery_name) >= 1"+
					"delivery_base_info.delivery_begin_time != nil && "+
					"delivery_base_info.delivery_end_time != nil",
				"时间格式为：yyyy-MM-dd HH:mm:ss",
			),
			paypay.NewRuler("投放计划玩法配置", "delivery_play_config != nil", fmt.Sprintf(consts.FmtEmptyAlert, "delivery_play_config")),
			paypay.NewRuler("投放计划玩法配置详细校验",
				`delivery_play_config.delivery_single_send_config?.delivery_content_info != nil &&
delivery_play_config.delivery_single_send_config?.delivery_content_info.delivery_content_type in ["ACTIVITY", "MINI_APP", "MINI_APP_SERVICE"] &&
(
	(delivery_play_config.delivery_single_send_config?.delivery_content_info.delivery_content_type == "ACTIVITY" && delivery_play_config.delivery_single_send_config?.delivery_content_info.delivery_activity_content != nil && delivery_play_config.delivery_single_send_config?.delivery_content_info.delivery_activity_content.activity_id != nil)
	||
	(delivery_play_config.delivery_single_send_config?.delivery_content_info.delivery_content_type in ["MINI_APP", "MINI_APP_SERVICE"] && delivery_play_config.delivery_single_send_config?.delivery_content_info.delivery_display_info != nil)
) &&
float(delivery_play_config.delivery_full_send_config?.delivery_floor_amount) >= 0.01 && float(delivery_play_config.delivery_full_send_config?.delivery_floor_amount) <= 99999 &&
delivery_play_config.delivery_full_send_config?.delivery_content_info != nil &&
delivery_play_config.delivery_full_send_config?.delivery_content_info.delivery_content_type in ["ACTIVITY", "MINI_APP", "MINI_APP_SERVICE"] &&
(
	(delivery_play_config.delivery_full_send_config?.delivery_content_info.delivery_content_type == "ACTIVITY" && delivery_play_config.delivery_full_send_config?.delivery_content_info.delivery_activity_content != nil && delivery_play_config.delivery_full_send_config?.delivery_content_info.delivery_activity_content.activity_id != nil)
	||
	(delivery_play_config.delivery_full_send_config?.delivery_content_info.delivery_content_type in ["MINI_APP", "MINI_APP_SERVICE"] && delivery_play_config.delivery_full_send_config?.delivery_content_info.delivery_display_info != nil)
)
`,
				"本接口参数较复杂，请结合原始文档判断 https://opendocs.alipay.com/open/47498bf2_alipay.marketing.activity.delivery.create?scene=8cb6764a5b944aa09266ed8109f74f62&pathHash=e94a8478",
			),
		},
	}
}
