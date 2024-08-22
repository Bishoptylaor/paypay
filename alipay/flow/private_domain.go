package flow

import (
	"context"
	"github.com/Bishoptylaor/paypay"
	"github.com/Bishoptylaor/paypay/alipay/entity"
)

// PayWithGift 支付有礼 https://opendocs.alipay.com/open/03o2f7?pathHash=e2a381af
type PayWithGift interface {
	MarketingActivityDeliveryStop(ctx context.Context, pl paypay.Payload) (aliRsp *entity.MarketingActivityDeliveryStopResponse, err error)     // 停止推广计划
	MarketingActivityDeliveryQuery(ctx context.Context, pl paypay.Payload) (aliRsp *entity.MarketingActivityDeliveryQueryResponse, err error)   // 查询推广计划
	MarketingActivityDeliveryCreate(ctx context.Context, pl paypay.Payload) (aliRsp *entity.MarketingActivityDeliveryCreateResponse, err error) // 创建推广计划
}
