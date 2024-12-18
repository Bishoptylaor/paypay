package flow

import (
	"context"
	"github.com/Bishoptylaor/paypay"
	"github.com/Bishoptylaor/paypay/alipay/entity"
)

// PaymentCoupon 支付券
type PaymentCoupon interface {
	// 支付券
	MarketingMaterialImageUpload(ctx context.Context, pl paypay.Payload) (aliRes *entity.MarketingMaterialImageUploadResponse, err error)       // 营销图片资源上传接口
	MarketingActivityVoucherCreate(ctx context.Context, pl paypay.Payload) (aliRes *entity.MarketingActivityVoucherCreateResponse, err error)   // 创建支付券
	MarketingActivityVoucherPublish(ctx context.Context, pl paypay.Payload) (aliRes *entity.MarketingActivityVoucherPublishResponse, err error) // 发布支付券
	MarketingActivityVoucherQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.MarketingActivityVoucherQueryResponse, err error)     // 查询支付券详情
	MarketingActivityVoucherModify(ctx context.Context, pl paypay.Payload) (aliRes *entity.MarketingActivityVoucherModifyResponse, err error)   // 修改支付券基本信息
	MarketingActivityVoucherAppend(ctx context.Context, pl paypay.Payload) (aliRes *entity.MarketingActivityVoucherAppendResponse, err error)   // 追加支付券预算
	MarketingActivityVoucherStop(ctx context.Context, pl paypay.Payload) (aliRes *entity.MarketingActivityVoucherStopResponse, err error)       // 停止支付券

	// 私域营销
	MarketingActivityBatchquery(ctx context.Context, pl paypay.Payload) (aliRes *entity.MarketingActivityBatchqueryResponse, err error)                       // 条件查询活动列表
	MarketingActivityConsult(ctx context.Context, pl paypay.Payload) (aliRes *entity.MarketingActivityConsultResponse, err error)                             // 活动领取咨询接口
	MarketingActivityQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.MarketingActivityQueryResponse, err error)                                 // 查询活动详情
	MarketingActivityMerchantBatchquery(ctx context.Context, pl paypay.Payload) (aliRes *entity.MarketingActivityMerchantBatchqueryResponse, err error)       // 查询活动可用商户
	MarketingActivityAppBatchquery(ctx context.Context, pl paypay.Payload) (aliRes *entity.MarketingActivityAppBatchqueryResponse, err error)                 // 查询活动可用小程序
	MarketingActivityShopBatchquery(ctx context.Context, pl paypay.Payload) (aliRes *entity.MarketingActivityShopBatchqueryResponse, err error)               // 查询活动可用门店
	MarketingActivityGoodsBatchquery(ctx context.Context, pl paypay.Payload) (aliRes *entity.MarketingActivityGoodsBatchqueryResponse, err error)             // 查询活动适用商品
	MarketingActivityUserBatchqueryvoucher(ctx context.Context, pl paypay.Payload) (aliRes *entity.MarketingActivityUserBatchqueryvoucherResponse, err error) // 条件查询用户券
	MarketingActivityUserQueryvoucher(ctx context.Context, pl paypay.Payload) (aliRes *entity.MarketingActivityUserQueryvoucherResponse, err error)           // 查询用户券详情
	MarketingCampaignOrderVoucherConsult(ctx context.Context, pl paypay.Payload) (aliRes *entity.MarketingCampaignOrderVoucherConsultResponse, err error)     // 订单优惠前置咨询

	// 蚂蚁门店管理
	AntMerchantExpandShopCreate(ctx context.Context, pl paypay.Payload) (aliRes *entity.AntMerchantExpandShopCreateResponse, err error)                   // 蚂蚁店铺创建
	AntMerchantExpandShopModify(ctx context.Context, pl paypay.Payload) (aliRes *entity.AntMerchantExpandShopModifyResponse, err error)                   // 修改蚂蚁店铺
	AntMerchantExpandOrderQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.AntMerchantExpandOrderQueryResponse, err error)                   // 商户申请单查询
	AntMerchantExpandShopPageQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.AntMerchantExpandShopPageQueryResponse, err error)             // 店铺分页查询接口
	AntMerchantExpandShopQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.AntMerchantExpandShopQueryResponse, err error)                     // 店铺查询接口
	AntMerchantExpandShopClose(ctx context.Context, pl paypay.Payload) (aliRes *entity.AntMerchantExpandShopCloseResponse, err error)                     // 蚂蚁店铺关闭
	AntMerchantExpandIndirectImageUpload(ctx context.Context, pl paypay.Payload) (aliRes *entity.AntMerchantExpandIndirectImageUploadResponse, err error) // 图片上传
	AntMerchantExpandMccQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.AntMerchantExpandMccQueryResponse, err error)                       // 商户mcc信息查询
}

// MerchantCoupon 商家券
type MerchantCoupon interface{}

// MerchantMemberCard 商家会员卡
type MerchantMemberCard interface{}

// MarketingRedEnv 营销活动红包
type MarketingRedEnv interface{}

// RedEnvelope 红包
type RedEnvelope interface {
}

// Chessboard 棋盘密云
type Chessboard interface{}
