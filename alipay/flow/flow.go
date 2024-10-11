package flow

import (
	"context"
	"github.com/Bishoptylaor/paypay"
	"github.com/Bishoptylaor/paypay/alipay/entity"
)

// 本包主要定义各种支付宝产品的具体工作流内容
// 具体文档参考：https://opendocs.alipay.com/open/065yhr?pathHash=d43962c0

type ExecuteFlow interface {
	PageExecute(ctx context.Context, pl paypay.Payload, method string, authToken ...string) (url string, err error)
}

type DataBillDownloadFlow interface {
	DataBillDownloadUrlQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.DataBillDownloadUrlQueryResponse, err error) // 对账单下载地址接口
}

type NotifyFlow interface {
	TradeRefundDepositbackCompleted(ctx context.Context, vals map[string][]string) (*entity.TradeRefundDepositbackCompletedReq, error)   // 收单退款冲退完成通知
	MarketingActivityDeliveryChanged(ctx context.Context, vals map[string][]string) (*entity.MarketingActivityDeliveryChangedReq, error) // 推广计划状态变更消息
}
