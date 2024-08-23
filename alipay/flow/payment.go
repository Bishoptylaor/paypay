package flow

import (
	"context"
	"github.com/Bishoptylaor/paypay"
	"github.com/Bishoptylaor/paypay/alipay/entity"
)

// Face2FacePay 当面付
type Face2FacePay interface {
	// 付款码支付
	TradeAppPay(ctx context.Context, pl paypay.Payload) (payParam string, err error) // 统一收单交易支付接口

	// 扫码支付
	TradePrecreate(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradePrecreateResponse, err error) // 统一收单线下交易预创建
	TradeCreate(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeCreateResponse, err error)       // 统一收单交易创建接口

	TradeQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeQueryResponse, err error)                           // 统一收单交易查询接口
	TradeRefund(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeRefundResponse, err error)                         // 统一收单交易退款接口
	TradeFastPayRefundQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeFastpayRefundQueryResponse, err error) // 统一收单交易退款查询
	TradeCancel(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeCancelResponse, err error)                         // 统一收单交易撤销接口
	TradeClose(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeCloseResponse, err error)                           // 统一收单交易关闭接口
	DataBillFlow
	ExecuteFlow
}

// AppPay app支付
type AppPay interface {
	TradeAppPay(ctx context.Context, pl paypay.Payload) (payParam string, err error)                                            // 统一收单交易支付接口
	TradeQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeQueryResponse, err error)                           // 统一收单交易查询接口
	TradeRefund(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeRefundResponse, err error)                         // 统一收单交易退款接口
	TradeClose(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeCloseResponse, err error)                           // 统一收单交易关闭接口
	TradeFastPayRefundQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeFastpayRefundQueryResponse, err error) // 统一收单交易退款查询

	DataBillFlow
	ExecuteFlow
}

// PhoneWebPay 手机网页支付
type PhoneWebPay interface {
	TradeWapPay(ctx context.Context, pl paypay.Payload) (payUrl string, err error)                                              // 手机网页支付接口
	TradeQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeQueryResponse, err error)                           // 统一收单交易查询接口
	TradeRefund(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeRefundResponse, err error)                         // 统一收单交易退款接口
	TradeClose(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeCloseResponse, err error)                           // 统一收单交易关闭接口
	TradeFastPayRefundQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeFastpayRefundQueryResponse, err error) // 统一收单交易退款查询

	DataBillFlow
	ExecuteFlow
}

// PcPagePay 电脑网页支付
type PcPagePay interface {
	TradePagePay(ctx context.Context, pl paypay.Payload) (payUrl string, err error)                                             // PC网页支付接口
	TradeQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeQueryResponse, err error)                           // 统一收单交易查询接口
	TradeRefund(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeRefundResponse, err error)                         // 统一收单交易退款接口
	TradeClose(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeCloseResponse, err error)                           // 统一收单交易关闭接口
	TradeFastPayRefundQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeFastpayRefundQueryResponse, err error) // 统一收单交易退款查询

	DataBillFlow
	ExecuteFlow
}

// MerchantDeduction 商家扣款 | 周期扣款
type MerchantDeduction interface {
	UserAgreementPageSign(ctx context.Context, pl paypay.Payload) (ret string, err error)                                                         // 个人代扣协议签约接口
	UserAgreementPageSignInApp(ctx context.Context, pl paypay.Payload) (ret string, err error)                                                    // 个人代扣协议签约接口 不跳转支付宝
	UserAgreementPageUnSign(ctx context.Context, pl paypay.Payload) (aliRes *entity.UserAgreementPageUnSignResponse, err error)                   // 个人代扣协议解约
	UserAgreementQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.UserAgreementQueryResponse, err error)                             // 个人代扣协议查询
	UserAgreementExecutionplanModify(ctx context.Context, pl paypay.Payload) (aliRes *entity.UserAgreementExecutionplanModifyResponse, err error) // 个人代扣协议执行计划修改

	TradeAppPay(ctx context.Context, pl paypay.Payload) (payParam string, err error)                    // 支付并签约接口
	TradePay(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradePayResponse, err error)       // 周期扣款接口
	TradeQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeQueryResponse, err error)   // 统一收单交易查询接口
	TradeRefund(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeRefundResponse, err error) // 统一收单交易退款接口
	TradeClose(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeCloseResponse, err error)   // 统一收单交易关闭接口
	TradeCancel(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeCancelResponse, err error) // 统一收单交易撤销接口

	DataBillFlow
	ExecuteFlow
}

// SmilePay 刷脸付
type SmilePay interface {
	ZolozAuthenticationSmilepayInitialize(ctx context.Context, pl paypay.Payload) (aliRes *entity.ZolozAuthenticationSmilepayInitializeResponse, err error)   // 刷脸支付初始化
	ZolozAuthenticationCustomerFtokenQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.ZolozAuthenticationCustomerFtokenQueryResponse, err error) // 查询刷脸结果信息接口
}

// PreLicensingPay 预授权支付
type PreLicensingPay interface {
	// 预授权
	FundAuthOrderAppFreeze(ctx context.Context, pl paypay.Payload) (payParam string, err error)                                           // 线上资金授权冻结接口
	FundAuthOperationDetailQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.FundAuthOperationDetailQueryResponse, err error) // 资金授权操作查询接口
	FundAuthOperationCancel(ctx context.Context, pl paypay.Payload) (aliRes *entity.FundAuthOperationCancelResponse, err error)           // 资金授权撤销接口
	FundAuthOrderUnfreeze(ctx context.Context, pl paypay.Payload) (aliRes *entity.FundAuthOrderUnfreezeResponse, err error)               // 资金授权解冻接口
	FundAuthOrdervoucherCreate(ctx context.Context, pl paypay.Payload) (aliRes *entity.FundAuthOrdervoucherCreateResponse, err error)     // 资金授权发码接口
	FundAuthOrderFreeze(ctx context.Context, pl paypay.Payload) (aliRes *entity.FundAuthOrderFreezeResponse, err error)                   // 资金授权冻结接口

	// 交易
	TradePay(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradePayResponse, err error)                               // 统一收单交易支付接口
	TradeClose(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeCloseResponse, err error)                           // 统一收单交易关闭接口
	TradeFastPayRefundQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeFastpayRefundQueryResponse, err error) // 统一收单交易退款查询
	TradeQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeQueryResponse, err error)                           // 统一收单交易查询接口
	TradeRefund(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeRefundResponse, err error)                         // 统一收单交易退款接口
	TradeOrderInfoSync(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeOrderInfoSyncResponse, err error)           // 支付宝订单信息同步接口

	// 账单
	DataBillFlow
}

// QrcodePay 订单码支付
type QrcodePay interface {
	TradePrecreate(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradePrecreateResponse, err error)                   // 统一收单线下交易预创建
	TradeClose(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeCloseResponse, err error)                           // 统一收单交易关闭接口
	TradeFastPayRefundQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeFastpayRefundQueryResponse, err error) // 统一收单交易退款查询
	TradeCancel(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeCancelResponse, err error)                         // 统一收单交易撤销接口
	TradeQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeQueryResponse, err error)                           // 统一收单交易查询接口
	TradeRefund(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeRefundResponse, err error)                         // 统一收单交易退款接口

	DataBillFlow
	ExecuteFlow
}

// JSAPIPay JSAPI支付
type JSAPIPay interface{}
