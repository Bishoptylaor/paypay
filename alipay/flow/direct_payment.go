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
 @Time    : 2024/8/26 -- 12:19
 @Author  : 亓官竹
 @Copyright 2024 亓官竹
 @Description: 标准解决方案：互联网平台直付通
*/

package flow

import (
	"context"
	"github.com/Bishoptylaor/paypay"
	"github.com/Bishoptylaor/paypay/alipay/entity"
)

type StandardSolutionDirectPayment interface {
	// 进件
	//AntMerchantExpandIndirectZftConsult(ctx context.Context, pl paypay.Payload) (aliRes *entity.AntMerchantExpandIndirectZftConsultResponse, err error)
	//AntMerchantExpandIndirectZftSettlementmodify(ctx context.Context, pl paypay.Payload) (aliRes *entity.AntMerchantExpandIndirectZftSettlementmodifyResponse, err error)
	//AntMerchantExpandIndirectZftUpgrade(ctx context.Context, pl paypay.Payload) (aliRes *entity.AntMerchantExpandIndirectZftUpgradeResponse, err error)
	//AntMerchantExpandIndirectZftSimplecreate(ctx context.Context, pl paypay.Payload) (aliRes *entity.AntMerchantExpandIndirectZftSimplecreateResponse, err error)
	//AntMerchantExpandIndirectZftCreate(ctx context.Context, pl paypay.Payload) (aliRes *entity.AntMerchantExpandIndirectZftCreateResponse, err error)
	//AntMerchantExpandIndirectZftModify(ctx context.Context, pl paypay.Payload) (aliRes *entity.AntMerchantExpandIndirectZftModifyResponse, err error)
	//AntMerchantExpandIndirectZftorderQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.AntMerchantExpandIndirectZftorderQueryResponse, err error)
	//AntMerchantExpandIndirectImageUpload(ctx context.Context, pl paypay.Payload) (aliRes *entity.AntMerchantExpandIndirectImageUploadResponse, err error) // 图片上传
	//AntMerchantExpandIndirectZftDelete(ctx context.Context, pl paypay.Payload) (aliRes *entity.AntMerchantExpandIndirectZftDeleteResponse, err error)

	// 分账关系维护
	TradeRoyaltyRelationBind(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeRoyaltyRelationBindResponse, err error)             // 分账关系绑定
	TradeRoyaltyRelationUnbind(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeRoyaltyRelationUnbindResponse, err error)         // 分账关系解绑
	TradeRoyaltyRelationBatchquery(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeRoyaltyRelationBatchqueryResponse, err error) // 分账关系查询

	// 交易收款
	TradeAppPay(ctx context.Context, pl paypay.Payload) (payParam string, err error) // 统一收单交易支付接口
	TradeWapPay(ctx context.Context, pl paypay.Payload) (payUrl string, err error)   // 手机网页支付接口
	//TradeAppMergePay(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeAppMergePayResponse, err error)
	TradeWapMergePay(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeWapMergePayResponse, err error)
	TradeCreate(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeCreateResponse, err error)
	//TradeMergePrecreate(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeMergePrecreateResponse, err error)
	//TradeMergeCreate(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeMergeCreateResponse, err error)
	TradePagePay(ctx context.Context, pl paypay.Payload) (payUrl string, err error)                   // 统一收单下单并支付页面接口
	TradeQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeQueryResponse, err error) // 统一收单交易查询接口
	DataBillDownloadFlow

	// 资金结算
	TradeSettleConfirm(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeSettleConfirmResponse, err error)

	// 分账/补差
	TradeOrderSettle(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeOrderSettleResponse, err error)
	TradeOrderSettleQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeOrderSettleQueryResponse, err error)
	TradeOrderOnsettleQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeOrderOnsettleQueryResponse, err error)

	// 退款/退分账/退补差
	TradeRefund(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeRefundResponse, err error)                         // 统一收单交易退款接口
	TradeFastPayRefundQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeFastpayRefundQueryResponse, err error) // 统一收单交易退款查询
}
