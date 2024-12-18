package flow

import (
	"context"
	"github.com/Bishoptylaor/paypay"
	"github.com/Bishoptylaor/paypay/alipay/entity"
)

// MerchantLedger 商家分账；见于预授权支付和商家扣款
type MerchantLedger interface {
	// 分账关系维护
	TradeRoyaltyRelationBind(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeRoyaltyRelationBindResponse, err error)             // 分账关系绑定
	TradeRoyaltyRelationUnbind(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeRoyaltyRelationUnbindResponse, err error)         // 分账关系解绑
	TradeRoyaltyRelationBatchquery(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeRoyaltyRelationBatchqueryResponse, err error) // 分账关系查询

	// 分账请求
	TradeOrderSettle(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeOrderSettleResponse, err error) // 统一收单交易结算接口

	// 分账查询
	TradeRoyaltyRateQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeRoyaltyRateQueryResponse, err error)     // 分账比例查询
	TradeOrderOnsettleQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeOrderOnsettleQueryResponse, err error) // 分账剩余金额查询
	TradeOrderSettleQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeOrderSettleQueryResponse, err error)     // 交易分账查询接口
}

// HuaBeiInstallment 花呗分期
type HuaBeiInstallment interface {
	TradeWapPay(ctx context.Context, pl paypay.Payload) (payUrl string, err error)                            // 手机网站支付接口2.0
	TradeAppPay(ctx context.Context, pl paypay.Payload) (payParam string, err error)                          // app支付接口2.0
	TradePay(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradePayResponse, err error)             // 统一收单交易支付接口
	TradePrecreate(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradePrecreateResponse, err error) // 统一收单线下交易预创建
}

// PureTransfer 转账到支付宝账户
type PureTransfer interface {
	FundAccountQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.FundAccountQueryResponse, err error)         // 支付宝资金账户资产查询接口
	FundQuotaQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.FundQuotaQueryResponse, err error)             // 转账额度查询接口
	FundTransUniTransfer(ctx context.Context, pl paypay.Payload) (aliRes *entity.FundTransUniTransferResponse, err error) // 单笔转账接口
	FundTransCommonQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.FundTransCommonQueryResponse, err error) // 转账业务单据查询接口

	DataBillEreceiptApply(ctx context.Context, pl paypay.Payload) (aliRes *entity.DataBillEreceiptApplyResponse, err error) // 申请电子回单(incubating)
	DataBillEreceiptQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.DataBillEreceiptQueryResponse, err error) // 查询电子回单状态(incubating)
	DataBillDownloadFlow
}
