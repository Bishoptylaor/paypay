package alipay

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Bishoptylaor/paypay"
	"github.com/Bishoptylaor/paypay/alipay/entity"
	"github.com/Bishoptylaor/paypay/alipay/utils"
	"github.com/Bishoptylaor/paypay/pkg"
)

// MarketingActivityDeliveryStop
// alipay.marketing.activity.delivery.stop(停止推广计划)
// 文档地址：https://opendocs.alipay.com/open/39c69f03_alipay.marketing.activity.delivery.stop
func (c *Client) MarketingActivityDeliveryStop(ctx context.Context, pl paypay.Payload) (aliRes *entity.MarketingActivityDeliveryStopResponse, err error) {
	var bs []byte
	var method = "alipay.marketing.activity.delivery.stop"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.MarketingActivityDeliveryStopResponse)
	if err = json.Unmarshal(bs, aliRes); err != nil || aliRes.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", pkg.ErrUnmarshal, string(bs))
	}
	if err = utils.ExtractBussErr(aliRes.Response.ErrorResponse); err != nil {
		return aliRes, err
	}
	signData, signDataErr := c.autoVerifySignByCert(ctx, bs, method, aliRes.AlipayCertSn)
	aliRes.SignData = signData
	return aliRes, signDataErr
}

// MarketingActivityDeliveryQuery
// alipay.marketing.activity.delivery.query(查询推广计划)
// 文档地址：https://opendocs.alipay.com/open/69c6d6c2_alipay.marketing.activity.delivery.query
func (c *Client) MarketingActivityDeliveryQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.MarketingActivityDeliveryQueryResponse, err error) {
	var bs []byte
	var method = "alipay.marketing.activity.delivery.query"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.MarketingActivityDeliveryQueryResponse)
	if err = json.Unmarshal(bs, aliRes); err != nil || aliRes.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", pkg.ErrUnmarshal, string(bs))
	}
	if err = utils.ExtractBussErr(aliRes.Response.ErrorResponse); err != nil {
		return aliRes, err
	}
	signData, signDataErr := c.autoVerifySignByCert(ctx, bs, method, aliRes.AlipayCertSn)
	aliRes.SignData = signData
	return aliRes, signDataErr
}

// MarketingActivityDeliveryCreate
// alipay.marketing.activity.delivery.create(创建推广计划)
// 文档地址：https://opendocs.alipay.com/open/47498bf2_alipay.marketing.activity.delivery.create
func (c *Client) MarketingActivityDeliveryCreate(ctx context.Context, pl paypay.Payload) (aliRes *entity.MarketingActivityDeliveryCreateResponse, err error) {
	var bs []byte
	var method = "alipay.marketing.activity.delivery.create"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.MarketingActivityDeliveryCreateResponse)
	if err = json.Unmarshal(bs, aliRes); err != nil || aliRes.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", pkg.ErrUnmarshal, string(bs))
	}
	if err = utils.ExtractBussErr(aliRes.Response.ErrorResponse); err != nil {
		return aliRes, err
	}
	signData, signDataErr := c.autoVerifySignByCert(ctx, bs, method, aliRes.AlipayCertSn)
	aliRes.SignData = signData
	return aliRes, signDataErr
}

// MarketingActivityVoucherCreate
// alipay.marketing.activity.voucher.create(创建支付券)
// 文档地址：https://opendocs.alipay.com/open/049d65
func (c *Client) MarketingActivityVoucherCreate(ctx context.Context, pl paypay.Payload) (aliRes *entity.MarketingActivityVoucherCreateResponse, err error) {
	var bs []byte
	var method = "alipay.marketing.activity.voucher.create"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.MarketingActivityVoucherCreateResponse)
	if err = json.Unmarshal(bs, aliRes); err != nil || aliRes.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", pkg.ErrUnmarshal, string(bs))
	}
	if err = utils.ExtractBussErr(aliRes.Response.ErrorResponse); err != nil {
		return aliRes, err
	}
	signData, signDataErr := c.autoVerifySignByCert(ctx, bs, method, aliRes.AlipayCertSn)
	aliRes.SignData = signData
	return aliRes, signDataErr
}

// MarketingActivityVoucherPublish
// alipay.marketing.activity.voucher.publish(激活支付券)
// 文档地址：https://opendocs.alipay.com/open/049d66
func (c *Client) MarketingActivityVoucherPublish(ctx context.Context, pl paypay.Payload) (aliRes *entity.MarketingActivityVoucherPublishResponse, err error) {
	var bs []byte
	var method = "alipay.marketing.activity.voucher.publish"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.MarketingActivityVoucherPublishResponse)
	if err = json.Unmarshal(bs, aliRes); err != nil || aliRes.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", pkg.ErrUnmarshal, string(bs))
	}
	if err = utils.ExtractBussErr(aliRes.Response.ErrorResponse); err != nil {
		return aliRes, err
	}
	signData, signDataErr := c.autoVerifySignByCert(ctx, bs, method, aliRes.AlipayCertSn)
	aliRes.SignData = signData
	return aliRes, signDataErr
}

// MarketingActivityVoucherQuery
// alipay.marketing.activity.voucher.query(查询支付券详情)
// 文档地址：https://opendocs.alipay.com/open/049d6g
func (c *Client) MarketingActivityVoucherQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.MarketingActivityVoucherQueryResponse, err error) {
	var bs []byte
	var method = "alipay.marketing.activity.voucher.query"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.MarketingActivityVoucherQueryResponse)
	if err = json.Unmarshal(bs, aliRes); err != nil || aliRes.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", pkg.ErrUnmarshal, string(bs))
	}
	if err = utils.ExtractBussErr(aliRes.Response.ErrorResponse); err != nil {
		return aliRes, err
	}
	signData, signDataErr := c.autoVerifySignByCert(ctx, bs, method, aliRes.AlipayCertSn)
	aliRes.SignData = signData
	return aliRes, signDataErr
}

// MarketingActivityVoucherModify
// alipay.marketing.activity.voucher.modify(修改支付券基本信息)
// 文档地址：https://opendocs.alipay.com/open/049d67
func (c *Client) MarketingActivityVoucherModify(ctx context.Context, pl paypay.Payload) (aliRes *entity.MarketingActivityVoucherModifyResponse, err error) {
	var bs []byte
	var method = "alipay.marketing.activity.voucher.modify"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.MarketingActivityVoucherModifyResponse)
	if err = json.Unmarshal(bs, aliRes); err != nil || aliRes.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", pkg.ErrUnmarshal, string(bs))
	}
	if err = utils.ExtractBussErr(aliRes.Response.ErrorResponse); err != nil {
		return aliRes, err
	}
	signData, signDataErr := c.autoVerifySignByCert(ctx, bs, method, aliRes.AlipayCertSn)
	aliRes.SignData = signData
	return aliRes, signDataErr
}

// MarketingActivityVoucherAppend
// alipay.marketing.activity.voucher.append(追加支付券预算)
// 文档地址：https://opendocs.alipay.com/open/049d68
func (c *Client) MarketingActivityVoucherAppend(ctx context.Context, pl paypay.Payload) (aliRes *entity.MarketingActivityVoucherAppendResponse, err error) {
	var bs []byte
	var method = "alipay.marketing.activity.voucher.append"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.MarketingActivityVoucherAppendResponse)
	if err = json.Unmarshal(bs, aliRes); err != nil || aliRes.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", pkg.ErrUnmarshal, string(bs))
	}
	if err = utils.ExtractBussErr(aliRes.Response.ErrorResponse); err != nil {
		return aliRes, err
	}
	signData, signDataErr := c.autoVerifySignByCert(ctx, bs, method, aliRes.AlipayCertSn)
	aliRes.SignData = signData
	return aliRes, signDataErr
}

// MarketingActivityVoucherStop
// alipay.marketing.activity.voucher.stop(停止支付券)
// 文档地址：https://opendocs.alipay.com/open/049d69
func (c *Client) MarketingActivityVoucherStop(ctx context.Context, pl paypay.Payload) (aliRes *entity.MarketingActivityVoucherStopResponse, err error) {
	var bs []byte
	var method = "alipay.marketing.activity.voucher.stop"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.MarketingActivityVoucherStopResponse)
	if err = json.Unmarshal(bs, aliRes); err != nil || aliRes.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", pkg.ErrUnmarshal, string(bs))
	}
	if err = utils.ExtractBussErr(aliRes.Response.ErrorResponse); err != nil {
		return aliRes, err
	}
	signData, signDataErr := c.autoVerifySignByCert(ctx, bs, method, aliRes.AlipayCertSn)
	aliRes.SignData = signData
	return aliRes, signDataErr
}

// MarketingActivityOrdervoucherCreate
// alipay.marketing.activity.ordervoucher.create(创建商家券活动)
// 文档地址：https://opendocs.alipay.com/open/7ad3a7bf_alipay.marketing.activity.ordervoucher.create
func (c *Client) MarketingActivityOrdervoucherCreate(ctx context.Context, pl paypay.Payload) (aliRes *entity.MarketingActivityOrdervoucherCreateResponse, err error) {
	var bs []byte
	var method = "alipay.marketing.activity.ordervoucher.create"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.MarketingActivityOrdervoucherCreateResponse)
	if err = json.Unmarshal(bs, aliRes); err != nil || aliRes.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", pkg.ErrUnmarshal, string(bs))
	}
	if err = utils.ExtractBussErr(aliRes.Response.ErrorResponse); err != nil {
		return aliRes, err
	}
	signData, signDataErr := c.autoVerifySignByCert(ctx, bs, method, aliRes.AlipayCertSn)
	aliRes.SignData = signData
	return aliRes, signDataErr
}

// MarketingActivityOrdervoucherCodedeposit
// alipay.marketing.activity.ordervoucher.codedeposit(同步商家券券码)
// 文档地址：https://opendocs.alipay.com/open/7ed0450d_alipay.marketing.activity.ordervoucher.codedeposit
func (c *Client) MarketingActivityOrdervoucherCodedeposit(ctx context.Context, pl paypay.Payload) (aliRes *entity.MarketingActivityOrdervoucherCodedepositResponse, err error) {
	var bs []byte
	var method = "alipay.marketing.activity.ordervoucher.codedeposit"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.MarketingActivityOrdervoucherCodedepositResponse)
	if err = json.Unmarshal(bs, aliRes); err != nil || aliRes.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", pkg.ErrUnmarshal, string(bs))
	}
	if err = utils.ExtractBussErr(aliRes.Response.ErrorResponse); err != nil {
		return aliRes, err
	}
	signData, signDataErr := c.autoVerifySignByCert(ctx, bs, method, aliRes.AlipayCertSn)
	aliRes.SignData = signData
	return aliRes, signDataErr
}

// MarketingActivityOrdervoucherModify
// alipay.marketing.activity.ordervoucher.modify(修改商家券活动基本信息)
// 文档地址：https://opendocs.alipay.com/open/528f83f6_alipay.marketing.activity.ordervoucher.modify
func (c *Client) MarketingActivityOrdervoucherModify(ctx context.Context, pl paypay.Payload) (aliRes *entity.MarketingActivityOrdervoucherModifyResponse, err error) {
	var bs []byte
	var method = "alipay.marketing.activity.ordervoucher.modify"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.MarketingActivityOrdervoucherModifyResponse)
	if err = json.Unmarshal(bs, aliRes); err != nil || aliRes.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", pkg.ErrUnmarshal, string(bs))
	}
	if err = utils.ExtractBussErr(aliRes.Response.ErrorResponse); err != nil {
		return aliRes, err
	}
	signData, signDataErr := c.autoVerifySignByCert(ctx, bs, method, aliRes.AlipayCertSn)
	aliRes.SignData = signData
	return aliRes, signDataErr
}

// MarketingActivityOrdervoucherStop
// alipay.marketing.activity.ordervoucher.stop(停止商家券活动)
// 文档地址：https://opendocs.alipay.com/open/16803efe_alipay.marketing.activity.ordervoucher.stop
func (c *Client) MarketingActivityOrdervoucherStop(ctx context.Context, pl paypay.Payload) (aliRes *entity.MarketingActivityOrdervoucherStopResponse, err error) {
	var bs []byte
	var method = "alipay.marketing.activity.ordervoucher.stop"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.MarketingActivityOrdervoucherStopResponse)
	if err = json.Unmarshal(bs, aliRes); err != nil || aliRes.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", pkg.ErrUnmarshal, string(bs))
	}
	if err = utils.ExtractBussErr(aliRes.Response.ErrorResponse); err != nil {
		return aliRes, err
	}
	signData, signDataErr := c.autoVerifySignByCert(ctx, bs, method, aliRes.AlipayCertSn)
	aliRes.SignData = signData
	return aliRes, signDataErr
}

// MarketingActivityOrdervoucherAppend
// alipay.marketing.activity.ordervoucher.append(修改商家券活动发券数量上限)
// 文档地址：https://opendocs.alipay.com/open/4e2acff5_alipay.marketing.activity.ordervoucher.append
func (c *Client) MarketingActivityOrdervoucherAppend(ctx context.Context, pl paypay.Payload) (aliRes *entity.MarketingActivityOrdervoucherAppendResponse, err error) {
	var bs []byte
	var method = "alipay.marketing.activity.ordervoucher.append"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.MarketingActivityOrdervoucherAppendResponse)
	if err = json.Unmarshal(bs, aliRes); err != nil || aliRes.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", pkg.ErrUnmarshal, string(bs))
	}
	if err = utils.ExtractBussErr(aliRes.Response.ErrorResponse); err != nil {
		return aliRes, err
	}
	signData, signDataErr := c.autoVerifySignByCert(ctx, bs, method, aliRes.AlipayCertSn)
	aliRes.SignData = signData
	return aliRes, signDataErr
}

// MarketingActivityOrdervoucherUse
// alipay.marketing.activity.ordervoucher.use(同步券核销状态)
// 文档地址：https://opendocs.alipay.com/open/3ffce87f_alipay.marketing.activity.ordervoucher.use
func (c *Client) MarketingActivityOrdervoucherUse(ctx context.Context, pl paypay.Payload) (aliRes *entity.MarketingActivityOrdervoucherUseResponse, err error) {
	var bs []byte
	var method = "alipay.marketing.activity.ordervoucher.use"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.MarketingActivityOrdervoucherUseResponse)
	if err = json.Unmarshal(bs, aliRes); err != nil || aliRes.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", pkg.ErrUnmarshal, string(bs))
	}
	if err = utils.ExtractBussErr(aliRes.Response.ErrorResponse); err != nil {
		return aliRes, err
	}
	signData, signDataErr := c.autoVerifySignByCert(ctx, bs, method, aliRes.AlipayCertSn)
	aliRes.SignData = signData
	return aliRes, signDataErr
}

// MarketingActivityOrdervoucherRefund
// alipay.marketing.activity.ordervoucher.refund(取消券核销状态)
// 文档地址：https://opendocs.alipay.com/open/4682759b_alipay.marketing.activity.ordervoucher.refund?scene=common
func (c *Client) MarketingActivityOrdervoucherRefund(ctx context.Context, pl paypay.Payload) (aliRes *entity.MarketingActivityOrdervoucherRefundResponse, err error) {
	var bs []byte
	var method = "alipay.marketing.activity.ordervoucher.refund"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.MarketingActivityOrdervoucherRefundResponse)
	if err = json.Unmarshal(bs, aliRes); err != nil || aliRes.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", pkg.ErrUnmarshal, string(bs))
	}
	if err = utils.ExtractBussErr(aliRes.Response.ErrorResponse); err != nil {
		return aliRes, err
	}
	signData, signDataErr := c.autoVerifySignByCert(ctx, bs, method, aliRes.AlipayCertSn)
	aliRes.SignData = signData
	return aliRes, signDataErr
}

// MarketingActivityOrdervoucherQuery
// alipay.marketing.activity.ordervoucher.query(查询商家券活动)
// 文档地址：https://opendocs.alipay.com/open/51f5946e_alipay.marketing.activity.ordervoucher.query
func (c *Client) MarketingActivityOrdervoucherQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.MarketingActivityOrdervoucherQueryResponse, err error) {
	var bs []byte
	var method = "alipay.marketing.activity.ordervoucher.query"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.MarketingActivityOrdervoucherQueryResponse)
	if err = json.Unmarshal(bs, aliRes); err != nil || aliRes.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", pkg.ErrUnmarshal, string(bs))
	}
	if err = utils.ExtractBussErr(aliRes.Response.ErrorResponse); err != nil {
		return aliRes, err
	}
	signData, signDataErr := c.autoVerifySignByCert(ctx, bs, method, aliRes.AlipayCertSn)
	aliRes.SignData = signData
	return aliRes, signDataErr
}

// MarketingActivityOrdervoucherCodecount
// alipay.marketing.activity.ordervoucher.codecount(统计商家券券码数量)
// 文档地址：https://opendocs.alipay.com/open/f6e49e82_alipay.marketing.activity.ordervoucher.codecount
func (c *Client) MarketingActivityOrdervoucherCodecount(ctx context.Context, pl paypay.Payload) (aliRes *entity.MarketingActivityOrdervoucherCodecountResponse, err error) {
	var bs []byte
	var method = "alipay.marketing.activity.ordervoucher.codecount"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.MarketingActivityOrdervoucherCodecountResponse)
	if err = json.Unmarshal(bs, aliRes); err != nil || aliRes.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", pkg.ErrUnmarshal, string(bs))
	}
	if err = utils.ExtractBussErr(aliRes.Response.ErrorResponse); err != nil {
		return aliRes, err
	}
	signData, signDataErr := c.autoVerifySignByCert(ctx, bs, method, aliRes.AlipayCertSn)
	aliRes.SignData = signData
	return aliRes, signDataErr
}

// MarketingActivityBatchquery
// alipay.marketing.activity.batchquery(条件查询活动列表)
// 文档地址：https://opendocs.alipay.com/open/04fgw9
func (c *Client) MarketingActivityBatchquery(ctx context.Context, pl paypay.Payload) (aliRes *entity.MarketingActivityBatchqueryResponse, err error) {
	var bs []byte
	var method = "alipay.marketing.activity.batchquery"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.MarketingActivityBatchqueryResponse)
	if err = json.Unmarshal(bs, aliRes); err != nil || aliRes.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", pkg.ErrUnmarshal, string(bs))
	}
	if err = utils.ExtractBussErr(aliRes.Response.ErrorResponse); err != nil {
		return aliRes, err
	}
	signData, signDataErr := c.autoVerifySignByCert(ctx, bs, method, aliRes.AlipayCertSn)
	aliRes.SignData = signData
	return aliRes, signDataErr
}

// MarketingActivityConsult
// alipay.marketing.activity.consult(活动领取咨询接口)
// 文档地址：https://opendocs.alipay.com/open/04fgwa
func (c *Client) MarketingActivityConsult(ctx context.Context, pl paypay.Payload) (aliRes *entity.MarketingActivityConsultResponse, err error) {
	var bs []byte
	var method = "alipay.marketing.activity.consult"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.MarketingActivityConsultResponse)
	if err = json.Unmarshal(bs, aliRes); err != nil || aliRes.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", pkg.ErrUnmarshal, string(bs))
	}
	if err = utils.ExtractBussErr(aliRes.Response.ErrorResponse); err != nil {
		return aliRes, err
	}
	signData, signDataErr := c.autoVerifySignByCert(ctx, bs, method, aliRes.AlipayCertSn)
	aliRes.SignData = signData
	return aliRes, signDataErr
}

// MarketingActivityQuery
// alipay.marketing.activity.query(查询活动详情)
// 文档地址：https://opendocs.alipay.com/open/04fgwb
func (c *Client) MarketingActivityQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.MarketingActivityQueryResponse, err error) {
	var bs []byte
	var method = "alipay.marketing.activity.query"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.MarketingActivityQueryResponse)
	if err = json.Unmarshal(bs, aliRes); err != nil || aliRes.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", pkg.ErrUnmarshal, string(bs))
	}
	if err = utils.ExtractBussErr(aliRes.Response.ErrorResponse); err != nil {
		return aliRes, err
	}
	signData, signDataErr := c.autoVerifySignByCert(ctx, bs, method, aliRes.AlipayCertSn)
	aliRes.SignData = signData
	return aliRes, signDataErr
}

// MarketingActivityMerchantBatchquery
// alipay.marketing.activity.merchant.batchquery(查询活动可用商户)
// 文档地址：https://opendocs.alipay.com/open/04fgwc
func (c *Client) MarketingActivityMerchantBatchquery(ctx context.Context, pl paypay.Payload) (aliRes *entity.MarketingActivityMerchantBatchqueryResponse, err error) {
	var bs []byte
	var method = "alipay.marketing.activity.merchant.batchquery"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.MarketingActivityMerchantBatchqueryResponse)
	if err = json.Unmarshal(bs, aliRes); err != nil || aliRes.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", pkg.ErrUnmarshal, string(bs))
	}
	if err = utils.ExtractBussErr(aliRes.Response.ErrorResponse); err != nil {
		return aliRes, err
	}
	signData, signDataErr := c.autoVerifySignByCert(ctx, bs, method, aliRes.AlipayCertSn)
	aliRes.SignData = signData
	return aliRes, signDataErr
}

// MarketingActivityAppBatchquery
// alipay.marketing.activity.app.batchquery(查询活动可用小程序)
// 文档地址：https://opendocs.alipay.com/open/04fgwd
func (c *Client) MarketingActivityAppBatchquery(ctx context.Context, pl paypay.Payload) (aliRes *entity.MarketingActivityAppBatchqueryResponse, err error) {
	var bs []byte
	var method = "alipay.marketing.activity.app.batchquery"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.MarketingActivityAppBatchqueryResponse)
	if err = json.Unmarshal(bs, aliRes); err != nil || aliRes.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", pkg.ErrUnmarshal, string(bs))
	}
	if err = utils.ExtractBussErr(aliRes.Response.ErrorResponse); err != nil {
		return aliRes, err
	}
	signData, signDataErr := c.autoVerifySignByCert(ctx, bs, method, aliRes.AlipayCertSn)
	aliRes.SignData = signData
	return aliRes, signDataErr
}

// MarketingActivityShopBatchquery
// alipay.marketing.activity.shop.batchquery(查询活动可用门店)
// 文档地址：https://opendocs.alipay.com/open/04fgwe
func (c *Client) MarketingActivityShopBatchquery(ctx context.Context, pl paypay.Payload) (aliRes *entity.MarketingActivityShopBatchqueryResponse, err error) {
	var bs []byte
	var method = "alipay.marketing.activity.shop.batchquery"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.MarketingActivityShopBatchqueryResponse)
	if err = json.Unmarshal(bs, aliRes); err != nil || aliRes.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", pkg.ErrUnmarshal, string(bs))
	}
	if err = utils.ExtractBussErr(aliRes.Response.ErrorResponse); err != nil {
		return aliRes, err
	}
	signData, signDataErr := c.autoVerifySignByCert(ctx, bs, method, aliRes.AlipayCertSn)
	aliRes.SignData = signData
	return aliRes, signDataErr
}

// MarketingActivityGoodsBatchquery
// alipay.marketing.activity.goods.batchquery(查询活动适用商品)
// 文档地址：https://opendocs.alipay.com/open/04fgwf
func (c *Client) MarketingActivityGoodsBatchquery(ctx context.Context, pl paypay.Payload) (aliRes *entity.MarketingActivityGoodsBatchqueryResponse, err error) {
	var bs []byte
	var method = "alipay.marketing.activity.goods.batchquery"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.MarketingActivityGoodsBatchqueryResponse)
	if err = json.Unmarshal(bs, aliRes); err != nil || aliRes.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", pkg.ErrUnmarshal, string(bs))
	}
	if err = utils.ExtractBussErr(aliRes.Response.ErrorResponse); err != nil {
		return aliRes, err
	}
	signData, signDataErr := c.autoVerifySignByCert(ctx, bs, method, aliRes.AlipayCertSn)
	aliRes.SignData = signData
	return aliRes, signDataErr
}

// MarketingActivityUserBatchqueryvoucher
// alipay.marketing.activity.user.batchqueryvoucher(条件查询用户券)
// 文档地址：https://opendocs.alipay.com/open/04fgwg
func (c *Client) MarketingActivityUserBatchqueryvoucher(ctx context.Context, pl paypay.Payload) (aliRes *entity.MarketingActivityUserBatchqueryvoucherResponse, err error) {
	var bs []byte
	var method = "alipay.marketing.activity.user.batchqueryvoucher"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.MarketingActivityUserBatchqueryvoucherResponse)
	if err = json.Unmarshal(bs, aliRes); err != nil || aliRes.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", pkg.ErrUnmarshal, string(bs))
	}
	if err = utils.ExtractBussErr(aliRes.Response.ErrorResponse); err != nil {
		return aliRes, err
	}
	signData, signDataErr := c.autoVerifySignByCert(ctx, bs, method, aliRes.AlipayCertSn)
	aliRes.SignData = signData
	return aliRes, signDataErr
}

// MarketingActivityUserQueryvoucher
// alipay.marketing.activity.user.queryvoucher(查询用户券详情)
// 文档地址：https://opendocs.alipay.com/open/04fgwh
func (c *Client) MarketingActivityUserQueryvoucher(ctx context.Context, pl paypay.Payload) (aliRes *entity.MarketingActivityUserQueryvoucherResponse, err error) {
	var bs []byte
	var method = "alipay.marketing.activity.user.queryvoucher"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.MarketingActivityUserQueryvoucherResponse)
	if err = json.Unmarshal(bs, aliRes); err != nil || aliRes.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", pkg.ErrUnmarshal, string(bs))
	}
	if err = utils.ExtractBussErr(aliRes.Response.ErrorResponse); err != nil {
		return aliRes, err
	}
	signData, signDataErr := c.autoVerifySignByCert(ctx, bs, method, aliRes.AlipayCertSn)
	aliRes.SignData = signData
	return aliRes, signDataErr
}
