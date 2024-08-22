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

type DataBillFlow interface {
	DataBillDownloadUrlQuery(ctx context.Context, pl paypay.Payload) (aliRsp *entity.DataBillDownloadUrlQueryResponse, err error) // 对账单下载地址接口
}
