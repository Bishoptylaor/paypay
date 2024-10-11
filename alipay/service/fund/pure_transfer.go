/*
 *  ┏┓      ┏┓
 *┏━┛┻━━━━━━┛┻┓
 *┃　　　━　　  ┃
 *┃   ┳┛ ┗┳   ┃
 *┃           ┃
 *┃     ┻     ┃
 *┗━━━┓     ┏━┛
 *　　 ┃　　　┃神兽保佑
 *　　 ┃　　　┃代码无BUG！
 *　　 ┃　　　┗━━━┓
 *　　 ┃         ┣┓
 *　　 ┃         ┏┛
 *　　 ┗━┓┓┏━━┳┓┏┛
 *　　   ┃┫┫  ┃┫┫
 *      ┗┻┛　 ┗┻┛
 @Time    : 2024/9/3 -- 10:29
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: pure_transfer.go
*/

package fund

import (
	"fmt"
	"github.com/Bishoptylaor/paypay"
	"github.com/Bishoptylaor/paypay/alipay"
	"github.com/Bishoptylaor/paypay/alipay/consts"
	"github.com/Bishoptylaor/paypay/alipay/flow"
	"github.com/Bishoptylaor/paypay/alipay/service"
)

type PureTransferCaller interface {
	flow.PureTransfer
	UseCustomPayloadCheckRuler(custom map[string][]paypay.Ruler)
}

type PureTransferService service.Service

// NewPureTransferCaller
//
// 初始化 转账到支付宝账户 相关接口功能
//
// 提供预设参数与参数校验能力
//
// 在保证参数正确的情况下，用户也可直接调用 client 中的相关接口实现
//
// 产品介绍 https://opendocs.alipay.com/open/06de8c?pathHash=654eb816
func NewPureTransferCaller(c *alipay.Client) PureTransferCaller {
	// do some implantation
	caller := &PureTransferService{Client: c}
	// 设置本产品相关接口默认参数校验规则
	caller.RulersMap = caller.setDefaultPayloadCheckRuler()
	alipay.Checker(caller.payloadChecker())(c)
	alipay.PayloadPreSetter(map[string][]paypay.PayloadPreSetter{
		"alipay.trade.app.pay": []paypay.PayloadPreSetter{
			paypay.PreSetter("product_code", "GENERAL_WITHHOLDING"),
		},
		"alipay.user.agreement.query": []paypay.PayloadPreSetter{
			paypay.PreSetter("personal_product_code", "CYCLE_PAY_AUTH_P"),
		},
	})(c)
	return caller
}

// UseCustomPayloadCheckRuler 外部可调用，用于自定义参数校验规则
func (c PureTransferService) UseCustomPayloadCheckRuler(custom map[string][]paypay.Ruler) {
	if c.RulersMap == nil {
		c.RulersMap = c.setDefaultPayloadCheckRuler()
	}
	for k, v := range custom {
		c.RulersMap[k] = v
	}
}

func (c PureTransferService) payloadChecker() paypay.PayloadRuler {
	return func(method string) []paypay.Ruler {
		if rulers, ok := c.RulersMap[method]; ok {
			return rulers
		} else {
			return []paypay.Ruler{}
		}
	}
}

func (c PureTransferService) setDefaultPayloadCheckRuler() map[string][]paypay.Ruler {
	bindUnbind := []paypay.Ruler{
		paypay.NewRuler("分账接收方列表",
			`receiver_list != nil && len(receiver_list) < 20 && `+
				`all(receiver_list, {.type in ["userId", "loginName", "openId"]}) && `+
				`none(receiver_list, {.type == "loginName" && .name == nil})`,
			"receiver_list 不为空 & 元素个数小于 20 &\nreceiver_list 所有元素中 type 字段取值范围为：\"userId\", \"loginName\", \"openId\"\nreceiver_list 所有元素中 如果 type == loginName 则 name 字段不能为空",
		),
		paypay.NewRuler("外部请求号，由商家自定义", "out_request_no != nil && len(out_request_no) <= 32", fmt.Sprintf(consts.FmtEmptyAlert, "out_request_no")),
	}
	return map[string][]paypay.Ruler{
		"alipay.fund.account.query": []paypay.Ruler{
			paypay.NewRuler("查询的账号类型，查询余额账户值为ACCTRANS_ACCOUNT。必填。",
				`account_type in [ACCTRANS_ACCOUNT, TRUSTEESHIP_ACCOUNT]`,
				"余额户查询: ACCTRANS_ACCOUNT \n托管账户查询: TRUSTEESHIP_ACCOUNT ",
			),
		},
		"alipay.fund.quota.query": []paypay.Ruler{
			paypay.NewRuler("产品编码",
				``,
				"业务产品码， 单笔转账到支付宝账户固定为: TRANS_ACCOUNT_NO_PWD； 收发现金红包固定为: STD_RED_PACKET；单笔付款到卡固定为：TRANS_BANKCARD_NO_PWD；单笔付款到卡固定为：TRANS_BANKCARD_NO_PWD；使用alipay.fund.trans.toaccount.transfer接口固定为DEFAULT",
			),
			paypay.NewRuler("业务场景",
				`biz_scene in [DIRECT_TRANSFER, DEFAULT]`,
				"DIRECT_TRANSFER：单笔无密转账到支付宝，单笔无密转账到银行卡，现金红包; DEFAULT：使用alipay.fund.trans.toaccount.transfer转账到户场景",
			),
		},

		"alipay.data.dataservice.bill.downloadurl.query": consts.DataDownloadRuler,
	}
}
