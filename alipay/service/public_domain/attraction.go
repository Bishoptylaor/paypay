package private_domain

import (
	"fmt"
	"github.com/Bishoptylaor/paypay"
	"github.com/Bishoptylaor/paypay/alipay"
	"github.com/Bishoptylaor/paypay/alipay/consts"
	"github.com/Bishoptylaor/paypay/alipay/flow"
	"github.com/Bishoptylaor/paypay/alipay/service"
)

type AttractingTrafficPromotionCaller interface {
	flow.AttractingTrafficPromotion
	UseCustomPayloadCheckRuler(custom map[string][]paypay.Ruler)
}

type AttractingTrafficPromotionService service.Service

// NewAttractingTrafficPromotionCaller
//
// 初始化 支付有礼 相关接口功能
//
// 提供预设参数与参数校验能力
//
// 在保证参数正确的情况下，用户也可直接调用 client 中的相关接口实现
//
// 产品介绍 https://opendocs.alipay.com/open/03o2f7?pathHash=e2a381af
func NewAttractingTrafficPromotionCaller(c *alipay.Client) AttractingTrafficPromotionCaller {
	// do some implantation
	caller := &AttractingTrafficPromotionService{Client: c}
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
func (c AttractingTrafficPromotionService) UseCustomPayloadCheckRuler(custom map[string][]paypay.Ruler) {
	if c.RulersMap == nil {
		c.RulersMap = c.setDefaultPayloadCheckRuler()
	}
	for k, v := range custom {
		c.RulersMap[k] = v
	}
}

func (c AttractingTrafficPromotionService) payloadChecker() paypay.PayloadRuler {
	return func(method string) []paypay.Ruler {
		if rulers, ok := c.RulersMap[method]; ok {
			return rulers
		} else {
			return []paypay.Ruler{}
		}
	}
}

func (c AttractingTrafficPromotionService) setDefaultPayloadCheckRuler() map[string][]paypay.Ruler {
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
			paypay.NewRuler("展位编码", `delivery_booth_code == "PUBLIC_UNION"`, "PUBLIC_UNION：日常推广"),
			paypay.NewRuler("外部业务单号", "out_biz_no != nil && len(delivery_id) <= 64 && len(delivery_id) >= 1", fmt.Sprintf(consts.FmtEmptyAlert, "out_biz_no")),
			paypay.NewRuler("商户接入模式",
				`merchant_access_mode in ["SELF_MODE", "AGENCY_MODE"]`,
				"merchant_access_mode 枚举值为 SELF_MODE, AGENCY_MODE",
			),
			paypay.NewRuler("投放计划基础信息",
				`delivery_base_info != nil && 
delivery_base_info.delivery_name != nil && len(delivery_base_info.delivery_name) <= 20 && len(delivery_base_info.delivery_name) >= 1 &&
delivery_base_info.delivery_begin_time != nil &&
delivery_base_info.delivery_end_time != nil && 
delivery_base_info.delivery_material != nil && delivery_base_info.delivery_material.delivery_single_material != nil && delivery_base_info.delivery_material.delivery_single_material.delivery_image != nil &&
`,
				"时间格式为：yyyy-MM-dd HH:mm:ss",
			),
			paypay.NewRuler("投放计划玩法配置", "delivery_play_config != nil", fmt.Sprintf(consts.FmtEmptyAlert, "delivery_play_config")),
			paypay.NewRuler("投放计划玩法配置详细校验",
				`delivery_play_config.delivery_single_send_config != nil &&
delivery_play_config.delivery_single_send_config.delivery_content_info != nil &&
delivery_play_config.delivery_single_send_config.delivery_content_info.delivery_content_type in ["ACTIVITY", "MINI_APP"] &&
delivery_play_config.delivery_single_send_config.delivery_content_info.delivery_activity_content != nil
`,
				"本接口参数较复杂，请结合原始文档判断 https://opendocs.alipay.com/open/f1382294_alipay.marketing.activity.delivery.create?scene=bbdb22eb41db4052bd7e5908e82be9a3&pathHash=8108c8f4",
			),
		},
		"alipay.marketing.material.image.upload": []paypay.Ruler{
			paypay.NewRuler("图片内容", "file_content != nil", fmt.Sprintf(consts.FmtEmptyAlert, "file_content")),
			paypay.NewRuler("商户接入模式", `merchant_access_mode in ["SELF_MODE", "AGENCY_MODE"]`, "商户自接入模式: SELF_MODE\n服务商代接入模式: AGENCY_MODE"),
			paypay.NewRuler("文件业务标识",
				`file_key in ["PUBLIC_UNION_CHANNEL_PIC", "DELIVERY_CHANNEL_PIC", "PROMO_BRAND_LOGO", "PROMO_VOUCHER_IMAGE"]`,
				"枚举值 alipay.marketing.activity.delivery.create接口中 delivery_base_info.delivery_material.delivery_single_material.delivery_image 当delivery_booth_code=PUBLIC_UNION，上传图片接口需指定file_key=PUBLIC_UNION_CHANNEL_PIC。上传图片尺寸600*600，支持格式：png、jpg、jpeg、bmp，大小不超过200kb； 当delivery_booth_code=PAYMENT_RESULT，上传图片接口需指定file_key=DELIVERY_CHANNEL_PIC。上传图片尺寸600*600，支持格式：png、jpg、jpeg、bmp，大小不超过200kb。 上传图片更多要求参考文档： https://render.alipay.com/p/c/18tpirlg12e8?operateFrom=BALIPAY alipay.marketing.activity.ordervoucher.create接口中 voucher_display_info.brand_logo字段,file_key=PROMO_BRAND_LOGO，上传图片尺寸600*600，支持格式：png、jpg、jpeg、bmp，大小不超过2MB voucher_display_info.voucher_image字段,file_key=PROMO_VOUCHER_IMAGE,上传图片尺寸600*600，支持格式：png、jpg、jpeg、bmp，大小不超过2MB"),
		},
	}
}
