package flow

import (
	"context"
	"github.com/Bishoptylaor/paypay"
	"github.com/Bishoptylaor/paypay/alipay/entity"
)

// AttractingTrafficPromotion 经营推广-引流转化
type AttractingTrafficPromotion interface {
	MarketingActivityDeliveryStop(ctx context.Context, pl paypay.Payload) (aliRsp *entity.MarketingActivityDeliveryStopResponse, err error)     // 停止推广计划
	MarketingActivityDeliveryQuery(ctx context.Context, pl paypay.Payload) (aliRsp *entity.MarketingActivityDeliveryQueryResponse, err error)   // 查询推广计划
	MarketingActivityDeliveryCreate(ctx context.Context, pl paypay.Payload) (aliRsp *entity.MarketingActivityDeliveryCreateResponse, err error) // 创建推广计划

	MarketingMaterialImageUpload(ctx context.Context, pl paypay.Payload) (aliRsp *entity.MarketingMaterialImageUploadResponse, err error) // 营销图片资源上传接口
}

// GoodsPromotion 经营推广-商品卖货
type GoodsPromotion interface {
	MarketingActivityDeliveryStop(ctx context.Context, pl paypay.Payload) (aliRsp *entity.MarketingActivityDeliveryStopResponse, err error)     // 停止推广计划
	MarketingActivityDeliveryQuery(ctx context.Context, pl paypay.Payload) (aliRsp *entity.MarketingActivityDeliveryQueryResponse, err error)   // 查询推广计划
	MarketingActivityDeliveryCreate(ctx context.Context, pl paypay.Payload) (aliRsp *entity.MarketingActivityDeliveryCreateResponse, err error) // 创建推广计划
}
