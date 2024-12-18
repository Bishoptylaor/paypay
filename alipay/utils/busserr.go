package utils

import (
	"fmt"
	"github.com/Bishoptylaor/paypay/alipay/entity"
	"slices"
)

type BussError struct {
	Code     string `json:"code"`
	Msg      string `json:"msg"`
	SubCode  string `json:"sub_code"`
	SubMsg   string `json:"sub_msg"`
	Solution string `json:"solution"`
}

// ExtractBussErr 检查业务码是否为10000 否则返回一个BizErr
func ExtractBussErr(errRes entity.ErrorResponse) error {
	if errRes.Code != "10000" {
		suggesstion, ok := BussErrorCodeMap[errRes.SubCode]
		return &BussError{
			Code:    errRes.Code,
			Msg:     errRes.Msg,
			SubCode: errRes.SubCode,
			SubMsg:  errRes.SubMsg,
			Solution: func() string {
				if ok {
					return suggesstion.Solution
				}
				return ""
			}(),
		}
	}
	return nil
}

// ExtractBussErrFunc 如果有其他 code 算作业务成功，可以使用这个接口个性化出来
func ExtractBussErrFunc(errRes entity.ErrorResponse, f func() []string) error {
	if !slices.Contains(f(), errRes.Code) {
		return &BussError{
			Code:    errRes.Code,
			Msg:     errRes.Msg,
			SubCode: errRes.SubCode,
			SubMsg:  errRes.SubMsg,
		}
	}
	return nil
}

func (e *BussError) Error() string {
	return fmt.Sprintf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s", "suggesstion": "%s"}`, e.Code, e.Msg, e.SubCode, e.SubMsg, e.Solution)
}

func IsBussError(err error) (*BussError, bool) {
	if bussErr, ok := err.(*BussError); ok {
		return bussErr, true
	}
	return nil, false
}

var BussErrorCodeMap = map[string]BussError{
	"ACQ.TRADE_NOT_EXIST": {
		SubCode:  "ACQ.TRADE_NOT_EXIST",
		SubMsg:   "交易不存在",
		Solution: "请检查 trade_no",
	},
	"ACQ.ACCESS_FORBIDDEN": {
		SubCode:  "ACQ.ACCESS_FORBIDDEN",
		SubMsg:   "无权限使用接口",
		Solution: "未签约对应的产品合约； 1、请校验传入的product_code参数是否正确； 2、确认请求商户是否签约了对应的产品合约；",
	},
	"ACQ.AGREEMENT_ERROR": {
		SubCode:  "ACQ.AGREEMENT_ERROR",
		SubMsg:   "协议信息异常",
		Solution: "请检查传入的协议信息是否正确",
	},
	"ACQ.AGREEMENT_INVALID": {
		SubCode:  "ACQ.AGREEMENT_INVALID",
		SubMsg:   "用户协议失效",
		Solution: "代扣业务传入的协议号对应的用户协议已经失效，需要用户重新签约",
	},
	"ACQ.AGREEMENT_NOT_EXIST": {
		SubCode:  "ACQ.AGREEMENT_NOT_EXIST",
		SubMsg:   "用户协议不存在或已解约",
		Solution: "协议号传递错误或用户已经解约，请重新引导用户进行签约，使用新生成的协议号发起代扣请求。",
	},
	"ACQ.AGREEMENT_STATUS_NOT_NORMAL": {
		SubCode:  "ACQ.AGREEMENT_STATUS_NOT_NORMAL",
		SubMsg:   "用户协议状态非NORMAL",
		Solution: "代扣业务用户协议状态非正常状态，需要用户解约后重新签约",
	},
	"ACQ.AMOUNT_OR_CURRENCY_ERROR": {
		SubCode:  "ACQ.AMOUNT_OR_CURRENCY_ERROR",
		SubMsg:   "订单金额或币种信息错误",
		Solution: "检查订单传入的金额信息是否有误，或者是不是当前币种未签约",
	},
	"ACQ.AUTH_AMOUNT_NOT_ENOUGH": {
		SubCode:  "ACQ.AUTH_AMOUNT_NOT_ENOUGH",
		SubMsg:   "授权金额不足",
		Solution: "订单金额大于授权剩余金额，请检查授权单剩余金额信息",
	},
	"ACQ.AUTH_NO_ERROR": {
		SubCode:  "ACQ.AUTH_NO_ERROR",
		SubMsg:   "预授权号错误或状态不对",
		Solution: "1、确认预授权单号（auth_no）是否正确；2、确认预授权订单的参与方与支付单的参与方是否一致；3、确认预授权订单的状态是否为已授权状态；",
	},
	"ACQ.AUTH_ORDER_HAS_CLOSED": {
		SubCode:  "ACQ.AUTH_ORDER_HAS_CLOSED",
		SubMsg:   "预授权订单已关闭",
		Solution: "预授权订单已撤销、解冻或未支付超时关闭，授权单状态已关闭，不能再发起转交易扣款",
	},
	"ACQ.AUTH_ORDER_HAS_FINISHED": {
		SubCode:  "ACQ.AUTH_ORDER_HAS_FINISHED",
		SubMsg:   "预授权订单已经完结",
		Solution: "预授权订单已经转交易支付成功，授权单状态已完结，不能再发起转交易扣款",
	},
	"ACQ.AUTH_ORDER_NOT_PAID": {
		SubCode:  "ACQ.AUTH_ORDER_NOT_PAID",
		SubMsg:   "预授权订单未支付",
		Solution: "先引导用户完成预授权订单的支付，再重新请求转交易支付接口",
	},
	"ACQ.AUTH_TOKEN_IS_NOT_EXIST": {
		SubCode:  "ACQ.AUTH_TOKEN_IS_NOT_EXIST",
		SubMsg:   "支付授权码为空",
		Solution: "请检查请求参数是否正确，支付授权码、协议信息或预授权号是否正确传入",
	},
	"ACQ.BEYOND_PAY_RESTRICTION": {
		SubCode:  "ACQ.BEYOND_PAY_RESTRICTION",
		SubMsg:   "商户收款额度超限",
		Solution: "联系支付宝小二提高限额（联系电话：4007585858）",
	},
	"ACQ.BEYOND_PER_RECEIPT_DAY_RESTRICTION": {
		SubCode:  "ACQ.BEYOND_PER_RECEIPT_DAY_RESTRICTION",
		SubMsg:   "订单金额超过当日累计限额",
		Solution: "联系支付宝小二提高限额（联系电话：4007585858）",
	},
	"ACQ.BEYOND_PER_RECEIPT_RESTRICTION": {
		SubCode:  "ACQ.BEYOND_PER_RECEIPT_RESTRICTION",
		SubMsg:   "商户收款金额超过月限额",
		Solution: "联系支付宝小二提高限额（联系电话：4007585858）",
	},
	"ACQ.BEYOND_PER_RECEIPT_SINGLE_RESTRICTION": {
		SubCode:  "ACQ.BEYOND_PER_RECEIPT_SINGLE_RESTRICTION",
		SubMsg:   "订单金额超过单笔限额",
		Solution: "联系支付宝小二提高限额（联系电话：4007585858）",
	},
	"ACQ.BUYER_BALANCE_NOT_ENOUGH": {
		SubCode:  "ACQ.BUYER_BALANCE_NOT_ENOUGH",
		SubMsg:   "买家余额不足",
		Solution: "买家绑定新的银行卡或者支付宝余额有钱后再发起支付",
	},
	"ACQ.BUYER_BANKCARD_BALANCE_NOT_ENOUGH": {
		SubCode:  "ACQ.BUYER_BANKCARD_BALANCE_NOT_ENOUGH",
		SubMsg:   "用户银行卡余额不足",
		Solution: "建议买家更换支付宝进行支付或者更换其它付款方式",
	},
	"ACQ.BUYER_ENABLE_STATUS_FORBID": {
		SubCode:  "ACQ.BUYER_ENABLE_STATUS_FORBID",
		SubMsg:   "买家状态非法",
		Solution: "用户联系支付宝小二（联系支付宝文档右边的客服头像或到支持中心咨询），确认买家状态为什么非法",
	},
	"ACQ.BUYER_NOT_EXIST": {
		SubCode:  "ACQ.BUYER_NOT_EXIST",
		SubMsg:   "买家不存在",
		Solution: "联系支付宝小二，确认买家是否已经注销账号",
	},
	"ACQ.BUYER_NOT_MAINLAND_CERT": {
		SubCode:  "ACQ.BUYER_NOT_MAINLAND_CERT",
		SubMsg:   "买家证件类型非大陆身份证",
		Solution: "该服务仅支持中国大陆身份证实名制用户,建议买家完善实名信息或者更换其它付款方式",
	},
	"ACQ.BUYER_PAYMENT_AMOUNT_DAY_LIMIT_ERROR": {
		SubCode:  "ACQ.BUYER_PAYMENT_AMOUNT_DAY_LIMIT_ERROR",
		SubMsg:   "买家付款日限额超限",
		Solution: "更换买家进行支付",
	},
	"ACQ.BUYER_PAYMENT_AMOUNT_MONTH_LIMIT_ERROR": {
		SubCode:  "ACQ.BUYER_PAYMENT_AMOUNT_MONTH_LIMIT_ERROR",
		SubMsg:   "买家付款月额度超限",
		Solution: "让买家更换账号后，重新付款或者更换其它付款方式",
	},
	"ACQ.BUYER_SELLER_EQUAL": {
		SubCode:  "ACQ.BUYER_SELLER_EQUAL",
		SubMsg:   "买卖家不能相同",
		Solution: "交易的买家和卖家不能相同，请更换买家后重新发起支付请求。",
	},
	"ACQ.BUYER_UNSUPPORT_ADVANCE": {
		SubCode:  "ACQ.BUYER_UNSUPPORT_ADVANCE",
		SubMsg:   "先享后付2.0准入失败,买家不满足垫资条件",
		Solution: "先享后付2.0准入失败,买家不满足垫资条件",
	},
	"ACQ.CARD_TYPE_ERROR": {
		SubCode:  "ACQ.CARD_TYPE_ERROR",
		SubMsg:   "卡类型错误",
		Solution: "检查传入的卡类型",
	},
	"ACQ.CARD_USER_NOT_MATCH": {
		SubCode:  "ACQ.CARD_USER_NOT_MATCH",
		SubMsg:   "脱机记录用户信息不匹配",
		Solution: "请检查传入的进展出站记录是否正确",
	},
	"ACQ.CERT_EXPIRED": {
		SubCode:  "ACQ.CERT_EXPIRED",
		SubMsg:   "凭证过期",
		Solution: "凭证过期",
	},
	"ACQ.CONTEXT_INCONSISTENT": {
		SubCode:  "ACQ.CONTEXT_INCONSISTENT",
		SubMsg:   "订单信息不一致",
		Solution: "商户订单号已经创建交易成功，且本次请求的交易关键信息（如订单金额、订单标题等）与已存交易的信息不一致，请检查传入的订单参数是否正确。 如原有交易未支付成功，请更换订单号重新发起支付请求。",
	},
	"ACQ.CURRENCY_NOT_SUPPORT": {
		SubCode:  "ACQ.CURRENCY_NOT_SUPPORT",
		SubMsg:   "订单币种不支持",
		Solution: "请检查是否签约对应的币种",
	},
	"ACQ.CYCLE_PAY_DATE_NOT_MATCH": {
		SubCode:  "ACQ.CYCLE_PAY_DATE_NOT_MATCH",
		SubMsg:   "扣款日期不在签约时的允许范围之内",
		Solution: "对于商家扣款产品，签约时会约定扣款的周期。如果发起扣款的日期不符合约定的周期，则不允许扣款。请重新检查扣款日期，在符合约定的日期发起扣款。",
	},
	"ACQ.CYCLE_PAY_SINGLE_FEE_EXCEED": {
		SubCode:  "ACQ.CYCLE_PAY_SINGLE_FEE_EXCEED",
		SubMsg:   "商家扣款的单笔金额超过签约时限制",
		Solution: "对于商家扣款产品，签约时会约定单笔扣款的最大金额。如果发起扣款的金额大于约定上限，则不允许扣款。请在允许的金额范围内扣款。",
	},
	"ACQ.CYCLE_PAY_TOTAL_FEE_EXCEED": {
		SubCode:  "ACQ.CYCLE_PAY_TOTAL_FEE_EXCEED",
		SubMsg:   "商家扣款的累计金额超过签约时限制",
		Solution: "对于商家扣款产品，签约时可以约定多次扣款的累计金额限制。如果发起扣款的累计金额大于约定上限，则不允许扣款。请在允许的金额范围内扣款。",
	},
	"ACQ.CYCLE_PAY_TOTAL_TIMES_EXCEED": {
		SubCode:  "ACQ.CYCLE_PAY_TOTAL_TIMES_EXCEED",
		SubMsg:   "商家扣款的总次数超过签约时限制",
		Solution: "对于商家扣款产品，签约时可以约定多次扣款的总次数限制。如果发起扣款的总次数大于约定上限，则不允许扣款。请在允许的次数范围内扣款",
	},
	"ACQ.ERROR_BALANCE_PAYMENT_DISABLE": {
		SubCode:  "ACQ.ERROR_BALANCE_PAYMENT_DISABLE",
		SubMsg:   "余额支付功能关闭",
		Solution: "用户打开余额支付开关后，再重新进行支付",
	},
	"ACQ.ERROR_BUYER_CERTIFY_LEVEL_LIMIT": {
		SubCode:  "ACQ.ERROR_BUYER_CERTIFY_LEVEL_LIMIT",
		SubMsg:   "买家未通过人行认证",
		Solution: "让用户联系支付宝小二并更换其它付款方式（联系电话：4007585858）",
	},
	"ACQ.EXIST_FORBIDDEN_WORD": {
		SubCode:  "ACQ.EXIST_FORBIDDEN_WORD",
		SubMsg:   "订单信息中包含违禁词",
		Solution: "请检查订单标题和订单描述是否包含敏感词，修改订单信息后，重新发起请求",
	},
	"ACQ.INVALID_PARAMETER": {
		SubCode:  "ACQ.INVALID_PARAMETER",
		SubMsg:   "参数无效",
		Solution: "请根据接口返回的错误信息，检查请求参数，修改后重新发起请求",
	},
	"ACQ.INVALID_RECEIVE_ACCOUNT": {
		SubCode:  "ACQ.INVALID_RECEIVE_ACCOUNT",
		SubMsg:   "收款账户不支持",
		Solution: "seller_id不在请求商户设置的收款账户限制集中，请确认是否需要收款到指定的支付宝账户，如需要联系支付宝小二进行配置处理。 如果不需要，则seller_id不需要传递，资金默认收款到签约商户账户中。",
	},
	"ACQ.INVALID_STORE_ID": {
		SubCode:  "ACQ.INVALID_STORE_ID",
		SubMsg:   "商户门店编号无效",
		Solution: "检查传入的门店编号是否有效",
	},
	"ACQ.MERCHANT_AGREEMENT_INVALID": {
		SubCode:  "ACQ.MERCHANT_AGREEMENT_INVALID",
		SubMsg:   "商户协议已失效",
		Solution: "商户与支付宝合同已失效，需要重新签约",
	},
	"ACQ.MERCHANT_AGREEMENT_NOT_EXIST": {
		SubCode:  "ACQ.MERCHANT_AGREEMENT_NOT_EXIST",
		SubMsg:   "商户协议不存在",
		Solution: "确认商户与支付宝是否已签约",
	},
	"ACQ.MERCHANT_PERM_RECEIPT_DAY_LIMIT": {
		SubCode:  "ACQ.MERCHANT_PERM_RECEIPT_DAY_LIMIT",
		SubMsg:   "超过单日累计收款额度",
		Solution: "联系支付宝小二处理（联系电话：4007585858）",
	},
	"ACQ.MERCHANT_PERM_RECEIPT_SINGLE_LIMIT": {
		SubCode:  "ACQ.MERCHANT_PERM_RECEIPT_SINGLE_LIMIT",
		SubMsg:   "超过单笔收款限额",
		Solution: "联系支付宝小二处理（联系电话：4007585858）",
	},
	"ACQ.MERCHANT_PERM_RECEIPT_SUSPEND_LIMIT": {
		SubCode:  "ACQ.MERCHANT_PERM_RECEIPT_SUSPEND_LIMIT",
		SubMsg:   "商户暂停收款",
		Solution: "联系支付宝小二处理（联系电话：4007585858）",
	},
	"ACQ.MERCHANT_STATUS_NOT_NORMAL": {
		SubCode:  "ACQ.MERCHANT_STATUS_NOT_NORMAL",
		SubMsg:   "商户状态异常",
		Solution: "因商户超过三个月未产生交易，需重新激活后可正常收单。1、进入支付宝商家中心，重新确认激活商家信息 或2、联系支付宝小二处理（联系电话：4007585858）",
	},
	"ACQ.MERCHANT_UNSUPPORT_ADVANCE": {
		SubCode:  "ACQ.MERCHANT_UNSUPPORT_ADVANCE",
		SubMsg:   "先享后付2.0准入失败,商户不支持垫资支付产品",
		Solution: "先享后付2.0准入失败,商户不支持垫资支付产品",
	},
	"ACQ.MOBILE_PAYMENT_SWITCH_OFF": {
		SubCode:  "ACQ.MOBILE_PAYMENT_SWITCH_OFF",
		SubMsg:   "用户的无线支付开关关闭",
		Solution: "用户在PC上打开无线支付开关后，再重新发起支付",
	},
	"ACQ.NOT_CERTIFIED_USER": {
		SubCode:  "ACQ.NOT_CERTIFIED_USER",
		SubMsg:   "买家非实名认证用户",
		Solution: "建议买家完善实名信息后再重试或者更换其它付款方式",
	},
	"ACQ.NOT_SUPPORT_PAYMENT_INST": {
		SubCode:  "ACQ.NOT_SUPPORT_PAYMENT_INST",
		SubMsg:   "不支持的钱包版本",
		Solution: "业务不支持使用该客户端支付，建议买家更换客户端进行支付或者更换其它付款方式",
	},
	"ACQ.NOW_TIME_AFTER_EXPIRE_TIME_ERROR": {
		SubCode:  "ACQ.NOW_TIME_AFTER_EXPIRE_TIME_ERROR",
		SubMsg:   "当前时间已超过允许支付的时间",
		Solution: "请检查传入的支付超时时间是否正确",
	},
	"ACQ.NO_PAYMENT_INSTRUMENTS_AVAILABLE": {
		SubCode:  "ACQ.NO_PAYMENT_INSTRUMENTS_AVAILABLE",
		SubMsg:   "没有可用的支付工具",
		Solution: "更换其它付款方式",
	},
	"ACQ.ORDER_UNSUPPORT_ADVANCE": {
		SubCode:  "ACQ.ORDER_UNSUPPORT_ADVANCE",
		SubMsg:   "订单不支持先享后付垫资",
		Solution: "订单不支持先享后付垫资",
	},
	"ACQ.PARTNER_ERROR": {
		SubCode:  "ACQ.PARTNER_ERROR",
		SubMsg:   "应用APP_ID填写错误",
		Solution: "联系支付宝小二（联系支付宝文档右边的客服头像或到支持中心咨询），确认APP_ID的状态",
	},
	"ACQ.PAYER_UNMATCHED": {
		SubCode:  "ACQ.PAYER_UNMATCHED",
		SubMsg:   "付款人不匹配",
		Solution: "建议用户更换为指定的支付宝账号进行支付",
	},
	"ACQ.PAYMENT_AUTH_CODE_INVALID": {
		SubCode:  "ACQ.PAYMENT_AUTH_CODE_INVALID",
		SubMsg:   "付款码无效",
		Solution: "1、请确认auth_code参数传递正确； 2、请用户刷新付款码后，重新扫码发起请求； 3、请用户确认付款码是否正确； 4、请用户确认付款码是否过期； 5、请用户确认付款码是否被使用； 6、请用户确认付款码是否被篡改； 7、请用户确认付款码是否被泄露； 8、请用户确认付款码是否被他人获取； 9、请用户确认付款码是否被他人篡改； ",
	},
	"ACQ.PAYMENT_FAIL": {
		SubCode:  "ACQ.PAYMENT_FAIL",
		SubMsg:   "支付失败",
		Solution: "用户刷新条码后，重新发起请求，如果重试一次后仍未成功，更换其它方式付款",
	},
	"ACQ.PAYMENT_REQUEST_HAS_RISK": {
		SubCode:  "ACQ.PAYMENT_REQUEST_HAS_RISK",
		SubMsg:   "支付有风险",
		Solution: "更换其它付款方式",
	},
	"ACQ.PLATFORM_BUSINESS_ACQUIRE_MODE_MUST_MERCHANT_ID": {
		SubCode:  "ACQ.PLATFORM_BUSINESS_ACQUIRE_MODE_MUST_MERCHANT_ID",
		SubMsg:   "二级商户编码为空",
		Solution: "二级商户编号(sub_merchant.merchant_id)不能为空。 直付通模式下，二级商户信息为必填项，如果不使用直付通模式接入，请联系支付宝小二改签。",
	},
	"ACQ.PRE_AUTH_PROD_CODE_INCONSISTENT": {
		SubCode:  "ACQ.PRE_AUTH_PROD_CODE_INCONSISTENT",
		SubMsg:   "预授权产品码不一致",
		Solution: "请检查预授权订单和转交易订单传入的产品码是否一致",
	},
	"ACQ.PRODUCT_AMOUNT_LIMIT_ERROR": {
		SubCode:  "ACQ.PRODUCT_AMOUNT_LIMIT_ERROR",
		SubMsg:   "产品额度超限",
		Solution: "联系支付宝小二提高限额（联系电话：4007585858）",
	},
	"ACQ.PULL_MOBILE_CASHIER_FAIL": {
		SubCode:  "ACQ.PULL_MOBILE_CASHIER_FAIL",
		SubMsg:   "唤起移动收银台失败",
		Solution: "用户刷新条码后，重新扫码发起请求",
	},
	"ACQ.REQUEST_AMOUNT_EXCEED": {
		SubCode:  "ACQ.REQUEST_AMOUNT_EXCEED",
		SubMsg:   "请求金额超限",
		Solution: "请检查传入的订单金额是否正确，预授权订单场景下请检查订单金额是否大于冻结金额",
	},
	"ACQ.RESTRICTED_MERCHANT_INDUSTRY": {
		SubCode:  "ACQ.RESTRICTED_MERCHANT_INDUSTRY",
		SubMsg:   "行业信息交易受限",
		Solution: "订单金额超过所属行业支持的最大金额",
	},
	"ACQ.RISK_MERCHANT_IP_NOT_EXIST": {
		SubCode:  "ACQ.RISK_MERCHANT_IP_NOT_EXIST",
		SubMsg:   "当前交易未传入IP信息，创单失败，请传入IP后再发起支付",
		Solution: "检查请求参数是否已经传入用户IP信息",
	},
	"ACQ.ROYALTY_ACCOUNT_NOT_EXIST": {
		SubCode:  "ACQ.ROYALTY_ACCOUNT_NOT_EXIST",
		SubMsg:   "分账收款方账号不存在",
		Solution: "请确认分账收款方账号是否正确",
	},
	"ACQ.SECONDARY_MERCHANT_ALIPAY_ACCOUNT_INVALID": {
		SubCode:  "ACQ.SECONDARY_MERCHANT_ALIPAY_ACCOUNT_INVALID",
		SubMsg:   "二级商户账户异常",
		Solution: "确认传入的二级商户结算账户是否与进件时设置的结算账户一致，如果一致可联系支付宝小二确认是否商户的账号信息有变更",
	},
	"ACQ.SECONDARY_MERCHANT_ID_BLANK": {
		SubCode:  "ACQ.SECONDARY_MERCHANT_ID_BLANK",
		SubMsg:   "二级商户编号错误",
		Solution: "请检查是否正确传入二级商户编号",
	},
	"ACQ.SECONDARY_MERCHANT_ID_INVALID": {
		SubCode:  "ACQ.SECONDARY_MERCHANT_ID_INVALID",
		SubMsg:   "二级商户不存在",
		Solution: "请检查传入的二级商户编号是否正确",
	},
	"ACQ.SECONDARY_MERCHANT_ISV_PUNISH_INDIRECT": {
		SubCode:  "ACQ.SECONDARY_MERCHANT_ISV_PUNISH_INDIRECT",
		SubMsg:   "商户状态异常",
		Solution: "请联系对应的服务商咨询",
	},
	"ACQ.SECONDARY_MERCHANT_NOT_MATCH": {
		SubCode:  "ACQ.SECONDARY_MERCHANT_NOT_MATCH",
		SubMsg:   "二级商户信息不匹配",
		Solution: "1、请检查发起支付请求的商户账号是否与请求创建二级商户接口所使用的商户账号一致； 2、如果接入的是直付通模式，请联系BD确认是否签约了直付通收单模式；",
	},
	"ACQ.SECONDARY_MERCHANT_STATUS_ERROR": {
		SubCode:  "ACQ.SECONDARY_MERCHANT_STATUS_ERROR",
		SubMsg:   "商户状态异常",
		Solution: "请联系对应的服务商咨询",
	},
	"ACQ.SELLER_BEEN_BLOCKED": {
		SubCode:  "ACQ.SELLER_BEEN_BLOCKED",
		SubMsg:   "商家账号被冻结",
		Solution: "联系支付宝小二，解冻账号（联系电话：4007585858）",
	},
	"ACQ.SELLER_NOT_EXIST": {
		SubCode:  "ACQ.SELLER_NOT_EXIST",
		SubMsg:   "卖家不存在",
		Solution: "确认卖家信息是否传递正确",
	},
	"ACQ.SMILE_PAY_MERCHANT_NOT_MATCH": {
		SubCode:  "ACQ.SMILE_PAY_MERCHANT_NOT_MATCH",
		SubMsg:   "请求支付和刷脸服务的商户身份不一致",
		Solution: "请检查请求支付和刷脸服务使用的pid是否一致",
	},
	"ACQ.SUB_GOODS_SIZE_MAX_COUNT": {
		SubCode:  "ACQ.SUB_GOODS_SIZE_MAX_COUNT",
		SubMsg:   "子商品明细超长",
		Solution: "请检查子商品明细是否超过了150条",
	},
	"ACQ.SUB_MERCHANT_CREATE_FAIL": {
		SubCode:  "ACQ.SUB_MERCHANT_CREATE_FAIL",
		SubMsg:   "二级商户创建失败",
		Solution: "检查上送的二级商户信息是否有效",
	},
	"ACQ.SUB_MERCHANT_TYPE_INVALID": {
		SubCode:  "ACQ.SUB_MERCHANT_TYPE_INVALID",
		SubMsg:   "二级商户类型非法",
		Solution: "检查上传的二级商户类型是否有效",
	},
	"ACQ.SYSTEM_ERROR": {
		SubCode:  "ACQ.SYSTEM_ERROR",
		SubMsg:   "系统异常",
		Solution: "系统异常错误下该笔交易可能成功也可能失败，请调用查询订单API，查询当前订单的状态，并根据订单状态决定下一步的操作，如果多次调用依然报此错误码，请联系支付宝客服",
	},
	"ACQ.TOTAL_FEE_EXCEED": {
		SubCode:  "ACQ.TOTAL_FEE_EXCEED",
		SubMsg:   "订单总金额超过限额",
		Solution: "订单金额不能小于等于0，且不能大于100000000元，请修改订单金额后重新发起支付请求。",
	},
	"ACQ.TRADE_BUYER_NOT_MATCH": {
		SubCode:  "ACQ.TRADE_BUYER_NOT_MATCH",
		SubMsg:   "交易买家不匹配",
		Solution: "请确认该订单号是否重复支付，如果是新订单，请更换商户订单号后重新提交支付。",
	},
	"ACQ.TRADE_HAS_CLOSE": {
		SubCode:  "ACQ.TRADE_HAS_CLOSE",
		SubMsg:   "交易已经关闭",
		Solution: "商户订单号对应的交易已经关闭（超时未支付或者已全额退款），请更换商户订单号后重新发起支付请求。",
	},
	"ACQ.TRADE_HAS_SUCCESS": {
		SubCode:  "ACQ.TRADE_HAS_SUCCESS",
		SubMsg:   "交易已被支付",
		Solution: "请确认该订单号是否重复支付，如果是新订单，请更换商户订单号后重新提交支付。",
	},
	"ACQ.TRADE_SETTLE_ERROR": {
		SubCode:  "ACQ.TRADE_SETTLE_ERROR",
		SubMsg:   "交易结算异常",
		Solution: "请检查传入的结算项信息是否正确，如果正确请联系支付宝小二",
	},
	"ACQ.TRADE_STATUS_ERROR": {
		SubCode:  "ACQ.TRADE_STATUS_ERROR",
		SubMsg:   "交易状态异常",
		Solution: "请检查订单状态是否已经支付成功",
	},
	"ACQ.UN_SUPPORT_TRADE_SCENE": {
		SubCode:  "ACQ.UN_SUPPORT_TRADE_SCENE",
		SubMsg:   "不支持该交易场景",
		Solution: "改签合约，勾选支持对应的交易场景",
	},
	"ACQ.USER_FACE_PAYMENT_SWITCH_OFF": {
		SubCode:  "ACQ.USER_FACE_PAYMENT_SWITCH_OFF",
		SubMsg:   "用户当面付付款开关关闭",
		Solution: "让用户在手机上打开当面付付款开关",
	},
	"ACQ.USER_LOGONID_DUP": {
		SubCode:  "ACQ.USER_LOGONID_DUP",
		SubMsg:   "用户账号重复",
		Solution: "用户手机账户名与他人重复，无法进行收付款。为了保障资金安全，建议您通知对方修改账户名，并与对方核对后更新对方账户名",
	},
	"ACQ.ZM_AUTH_AMOUNT_EXCEED": {
		SubCode:  "ACQ.ZM_AUTH_AMOUNT_EXCEED",
		SubMsg:   "先用后付场景下超过约定的免密支付金额",
		Solution: "超过约定的免密支付金额，需要商户调用支付宝SDK唤起收银台，用户确认后付款",
	},
	"ACQ.ZM_AUTH_RULE_LIMIT": {
		SubCode:  "ACQ.ZM_AUTH_RULE_LIMIT",
		SubMsg:   "调用芝麻做先用后付鉴权芝麻授权失败，命中先用后付场景限制规则",
		Solution: "请更换调用 alipay.trade.app.pay服务，继续信用下单流程",
	},
	"ACQ.ZM_CREDIT_AUTH_FAIL": {
		SubCode:  "ACQ.ZM_CREDIT_AUTH_FAIL",
		SubMsg:   "综合评估不通过",
		Solution: "建议用户按时履约，提升芝麻信用等级",
	},
	"ACQ.CANCEL_NOT_ALLOWED": {
		SubCode:  "ACQ.CANCEL_NOT_ALLOWED",
		SubMsg:   "交易不允许撤销",
		Solution: "该笔交易不支持撤销",
	},
	"ACQ.REASON_TRADE_BEEN_FREEZEN": {
		SubCode:  "ACQ.REASON_TRADE_BEEN_FREEZEN",
		SubMsg:   "当前交易被冻结，不允许进行撤销",
		Solution: "联系支付宝小二，确认该笔交易的具体情况",
	},
	"ACQ.REASON_TRADE_REFUND_FEE_ERR": {
		SubCode:  "ACQ.REASON_TRADE_REFUND_FEE_ERR",
		SubMsg:   "退款金额无效",
		Solution: "请确认交易是否已经发起过退款，已经退款场景下不能再发起撤销",
	},
	"ACQ.SELLER_BALANCE_NOT_ENOUGH": {
		SubCode:  "ACQ.SELLER_BALANCE_NOT_ENOUGH",
		SubMsg:   "商户的支付宝账户中无足够的资金进行撤销",
		Solution: "商户支付宝账户充值后重新发起撤销即可",
	},
	"ACQ.TRADE_CANCEL_TIME_OUT": {
		SubCode:  "ACQ.TRADE_CANCEL_TIME_OUT",
		SubMsg:   "超过撤销时间范围",
		Solution: "交易超过了撤销的时间范围，可使用退款接口发起交易退款",
	},
	"ACQ.TRADE_HAS_FINISHED": {
		SubCode:  "ACQ.TRADE_HAS_FINISHED",
		SubMsg:   "交易已经完结",
		Solution: "请使用相同的参数再次调用",
	},
	"BILL_NOT_EXIST": {
		SubCode:  "BILL_NOT_EXIST",
		SubMsg:   "账单不存在",
		Solution: "确认参数后重新查询",
	},
	"ACQ.ENTERPRISE_PAY_BIZ_ERROR": {
		SubCode:  "ACQ.ENTERPRISE_PAY_BIZ_ERROR",
		SubMsg:   "因公付业务异常",
		Solution: "重新发起查询请求，如果多次重试后仍返回同样的错误，请联系支付宝小二处理",
	},
	"ACQ.ALLOC_AMOUNT_VALIDATE_ERROR": {
		SubCode:  "ACQ.ALLOC_AMOUNT_VALIDATE_ERROR",
		SubMsg:   "退分账金额超限",
		Solution: "请调整退分账金额后重试",
	},
	"ACQ.BUYER_ERROR": {
		SubCode:  "ACQ.BUYER_ERROR",
		SubMsg:   "买家状态异常",
		Solution: "联系支付宝小二确认买家状态异常原因，或者可联系买家进行线下退款处理",
	},
	"ACQ.CUSTOMER_VALIDATE_ERROR": {
		SubCode:  "ACQ.CUSTOMER_VALIDATE_ERROR",
		SubMsg:   "账户已注销或者被冻结",
		Solution: "请查询账户状态：1. 如果账户已注销，请线下处理；2. 如果账户已冻结，请联系支付宝小二确认冻结原因。",
	},
	"ACQ.DISCORDANT_REPEAT_REQUEST": {
		SubCode:  "ACQ.DISCORDANT_REPEAT_REQUEST",
		SubMsg:   "请求信息不一致",
		Solution: "退款请求号对应的退款已经执行成功，且本次请求的退款金额与之前请求的金额不一致，请检查传入的退款金额是否正确。 或者通过退款查询接口获取退款执行结果。",
	},
	"ACQ.NOT_ALLOW_PARTIAL_REFUND": {
		SubCode:  "ACQ.NOT_ALLOW_PARTIAL_REFUND",
		SubMsg:   "不支持部分退款",
		Solution: "由于交易使用了特定的优惠券等场景，该笔交易不支持部分退款，请对交易进行全额退款或者联系买家进行线下退款处理",
	},
	"ACQ.ONLINE_TRADE_VOUCHER_NOT_ALLOW_REFUND": {
		SubCode:  "ACQ.ONLINE_TRADE_VOUCHER_NOT_ALLOW_REFUND",
		SubMsg:   "交易不允许退款",
		Solution: "此交易中核销了购买的代金券，不允许进行退款，可联系买家进行线下退款处理",
	},
	"ACQ.OVERDRAFT_AGREEMENT_NOT_MATCH": {
		SubCode:  "ACQ.OVERDRAFT_AGREEMENT_NOT_MATCH",
		SubMsg:   "垫资退款接口传入模式和签约配置不一致",
		Solution: "请检查垫资退款合约中的出资方式，修改合约或接口传参后重试",
	},
	"ACQ.OVERDRAFT_ASSIGN_ACCOUNT_INVALID": {
		SubCode:  "ACQ.OVERDRAFT_ASSIGN_ACCOUNT_INVALID",
		SubMsg:   "垫资退款出资账号和商户信息不一致",
		Solution: "垫资退款出资账号必须为商户名下支付宝账号，请更换出资账号后重试",
	},
	"ACQ.REASON_TRADE_STATUS_INVALID": {
		SubCode:  "ACQ.REASON_TRADE_STATUS_INVALID",
		SubMsg:   "交易状态异常",
		Solution: "查询交易，确认交易是否是支付成功状态，是的话可联系支付宝小二确认交易状态",
	},
	"ACQ.REFUND_ACCOUNT_NOT_EXIST": {
		SubCode:  "ACQ.REFUND_ACCOUNT_NOT_EXIST",
		SubMsg:   "退款出资账号不存在或账号异常",
		Solution: "检查退款出资账号状态，账号正常后重试",
	},
	"ACQ.REFUND_AMT_NOT_EQUAL_TOTAL": {
		SubCode:  "ACQ.REFUND_AMT_NOT_EQUAL_TOTAL",
		SubMsg:   "退款金额超限",
		Solution: "1、请检查退款金额是否正确，请求的退款金额不能大于交易总金额； 2、如果不是全额退款，退款请求号必填，请检查是否传入了退款请求号；",
	},
	"ACQ.REFUND_CHARGE_ERROR": {
		SubCode:  "ACQ.REFUND_CHARGE_ERROR",
		SubMsg:   "退收费异常",
		Solution: "请过一段时间后再重试发起退款",
	},
	"ACQ.REFUND_FEE_ERROR": {
		SubCode:  "ACQ.REFUND_FEE_ERROR",
		SubMsg:   "交易退款金额有误",
		Solution: "请检查传入的退款金额是否正确",
	},
	"ACQ.REFUND_ROYALTY_PAYEE_ACCOUNT_NOT_EXIST": {
		SubCode:  "ACQ.REFUND_ROYALTY_PAYEE_ACCOUNT_NOT_EXIST",
		SubMsg:   "退分账收入方账户不存在",
		Solution: "退分账收入方账户不存在，请确认收入方账号是否正确，更换账号后重新发起",
	},
	"ACQ.TRADE_NOT_ALLOW_REFUND": {
		SubCode:  "ACQ.TRADE_NOT_ALLOW_REFUND",
		SubMsg:   "当前交易不允许退款",
		Solution: "检查当前交易的状态是否为交易成功状态以及签约的退款属性是否允许退款，确认后，重新发起请求",
	},
	"AGREEMENT_HAS_UNSIGNED": {
		SubCode:  "AGREEMENT_HAS_UNSIGNED",
		SubMsg:   "用户协议已解约",
		Solution: "用户于支付宝侧协议已解约，商户需解约对应协议",
	},
	"AUTHOREE_IS_NOT_MATCH": {
		SubCode:  "AUTHOREE_IS_NOT_MATCH",
		SubMsg:   "被授权方不匹配",
		Solution: "确认商户app_id对应的被授权方与用户协议中的被授权方是否一致",
	},
	"MERCHANT_STATUS_IS_NOT_NORMAL": {
		SubCode:  "MERCHANT_STATUS_IS_NOT_NORMAL",
		SubMsg:   "商户协议状态不正常",
		Solution: "检查商户协议是否正确",
	},
	"PARENT_MERCHANT_QUERY_FAIL": {
		SubCode:  "PARENT_MERCHANT_QUERY_FAIL",
		SubMsg:   "平台商户查询失败",
		Solution: "检查父商户是否存在，或重试处理",
	},
	"PRODUCT_CODE_NOT_SUPPORTED_ERROR": {
		SubCode:  "PRODUCT_CODE_NOT_SUPPORTED_ERROR",
		SubMsg:   "无效的个人产品码",
		Solution: "商户确认个人产品码填写是否正确",
	},
	"USER_AGREEMENT_NOT_EXIST": {
		SubCode:  "USER_AGREEMENT_NOT_EXIST",
		SubMsg:   "用户协议不存在",
		Solution: "检查用户协议是否正确",
	},
	"USER_NOT_EXIST_ERROR": {
		SubCode:  "USER_NOT_EXIST_ERROR",
		SubMsg:   "支付宝用户信息不存在",
		Solution: "检查用户信息是否正确",
	},
	"USER_NOT_EXSIT_ERROR": {
		SubCode:  "USER_NOT_EXSIT_ERROR",
		SubMsg:   "用户信息不存在",
		Solution: "检查用户信息是否正确",
	},
	"temp": {
		SubCode:  "",
		SubMsg:   "",
		Solution: "",
	},
}
