package flow

import (
	"context"
	"github.com/Bishoptylaor/paypay"
	"github.com/Bishoptylaor/paypay/alipay/entity"
)

// AppletComplaints 小程序交易投诉处理
type AppletComplaints interface {
}

// AntShopManager 蚂蚁店铺管理
type AntShopManager interface {
}

// CCMPlugIn CCM插件化能力
type CCMPlugIn interface {
}

// SchoolMod 学校库
type SchoolMod interface {
}

// MerchantDataQueryDownload 商家账单数据查询及下载
type MerchantDataQueryDownload interface {
	DataBillSellQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.DataBillSellQueryResponse, err error)             // 支付宝商家账户卖出交易查询
	DataBillBuyQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.DataBillBuyQueryResponse, err error)               // 支付宝商家账户买入交易查询
	DataBillAccountlogQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.DataBillAccountlogQueryResponse, err error) // 支付宝商家账户账务明细查询
	DataBillTransferQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.DataBillTransferQueryResponse, err error)     // 支付宝商家账户充值，转账，提现查询
	DataBillBailQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.DataBillBailQueryResponse, err error)             // 支付宝商家账户保证金查询
	DataBillBalanceQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.DataBillBalanceQueryResponse, err error)       // 支付宝商家账户当前余额查询
	DataBillBalancehisQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.DataBillBalancehisQueryResponse, err error) // 支付宝商家账户历史余额查询
	DataBillDownloadFlow
}

// MerchantMonitor 商家自主监控
type MerchantMonitor interface {
}

// AlipayCards 支付宝卡包
type AlipayCards interface {
}

// EInvoice 电子发票
type EInvoice interface {
}

// DailyExpenses 生活缴费
type DailyExpenses interface {
}

// PreCollegeExpenses 中小学教育缴费
type PreCollegeExpenses interface {
}

// MetroTicketOnlinePurchase 地铁线上购票
type MetroTicketOnlinePurchase interface {
}

// ParkingExpenses 停车在线缴费
type ParkingExpenses interface {
}

// KouBeiStore 口碑开店
type KouBeiStore interface {
}

// FuelFilling 支付宝加油
type FuelFilling interface {
}

// PrivateMarketing 私域营销
type PrivateMarketing interface {
}

// MerchantFeeApply 商家费率申请
type MerchantFeeApply interface {
}
