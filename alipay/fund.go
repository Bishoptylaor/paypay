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
 @Time    : 2024/8/22 -- 12:19
 @Author  : 亓官竹
 @Copyright 2024 亓官竹
 @Description: fund.go
*/

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

// FundAccountQuery
// alipay.fund.account.query(支付宝资金账户资产查询接口)
// 文档地址：https://opendocs.alipay.com/open/02byuq
func (c *Client) FundAccountQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.FundAccountQueryResponse, err error) {
	var bs []byte
	var method = "alipay.fund.account.query"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.FundAccountQueryResponse)
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

// FundQuotaQuery
// alipay.fund.account.query(支付宝资金账户资产查询接口)
// 文档地址：https://opendocs.alipay.com/open/02byuq
func (c *Client) FundQuotaQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.FundQuotaQueryResponse, err error) {
	var bs []byte
	var method = "alipay.fund.quota.query"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.FundQuotaQueryResponse)
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

// FundTransUniTransfer
// alipay.fund.trans.uni.transfer(单笔转账接口)
// 文档地址：https://opendocs.alipay.com/open/02byuo
func (c *Client) FundTransUniTransfer(ctx context.Context, pl paypay.Payload) (aliRes *entity.FundTransUniTransferResponse, err error) {
	var bs []byte
	var method = "alipay.fund.trans.uni.transfer"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.FundTransUniTransferResponse)
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

// FundTransCommonQuery
// alipay.fund.trans.common.query(转账业务单据查询接口)
// 文档地址：https://opendocs.alipay.com/open/02byup
func (c *Client) FundTransCommonQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.FundTransCommonQueryResponse, err error) {
	var bs []byte
	var method = "alipay.fund.trans.common.query"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.FundTransCommonQueryResponse)
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

// FundTransOrderQuery
// alipay.fund.trans.order.query(查询转账订单接口)
// 文档地址: https://opendocs.alipay.com/apis/api_28/alipay.fund.trans.order.query
func (c *Client) FundTransOrderQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.FundTransOrderQueryResponse, err error) {
	var bs []byte
	var method = "alipay.fund.trans.order.query"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}

	aliRes = new(entity.FundTransOrderQueryResponse)
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

// FundTransRefund
// alipay.fund.trans.refund(资金退回接口)
// 文档地址: https://opendocs.alipay.com/open/02byvd
func (c *Client) FundTransRefund(ctx context.Context, pl paypay.Payload) (aliRes *entity.FundTransRefundResponse, err error) {
	var bs []byte
	var method = "alipay.fund.trans.refund"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.FundTransRefundResponse)
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

// FundAuthOrderFreeze
// alipay.fund.auth.order.freeze(资金授权冻结接口)
// 文档地址: https://opendocs.alipay.com/open/02fkb9
func (c *Client) FundAuthOrderFreeze(ctx context.Context, pl paypay.Payload) (aliRes *entity.FundAuthOrderFreezeResponse, err error) {
	var bs []byte
	var method = "alipay.fund.auth.order.freeze"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.FundAuthOrderFreezeResponse)
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

// FundAuthOrdervoucherCreate
// alipay.fund.auth.order.voucher.create(资金授权发码接口)
// 文档地址: https://opendocs.alipay.com/open/02fit5
func (c *Client) FundAuthOrdervoucherCreate(ctx context.Context, pl paypay.Payload) (aliRes *entity.FundAuthOrdervoucherCreateResponse, err error) {
	var bs []byte
	var method = "alipay.fund.auth.order.voucher.create"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.FundAuthOrdervoucherCreateResponse)
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

// FundAuthOrderAppFreeze
// alipay.fund.auth.order.app.freeze(线上资金授权冻结接口)
// 文档地址: https://opendocs.alipay.com/open/02f912
func (c *Client) FundAuthOrderAppFreeze(ctx context.Context, pl paypay.Payload) (payParam string, err error) {
	var bs []byte
	if bs, err = c.callAli(ctx, pl, "alipay.fund.auth.order.app.freeze"); err != nil {
		return "", err
	}
	payParam = string(bs)
	return payParam, nil
}

// FundAuthOrderUnfreeze
// alipay.fund.auth.order.unfreeze(资金授权解冻接口)
// 文档地址: https://opendocs.alipay.com/open/02fkbc
func (c *Client) FundAuthOrderUnfreeze(ctx context.Context, pl paypay.Payload) (aliRes *entity.FundAuthOrderUnfreezeResponse, err error) {
	var bs []byte
	var method = "alipay.fund.auth.order.unfreeze"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.FundAuthOrderUnfreezeResponse)
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

// FundAuthOperationDetailQuery
// alipay.fund.auth.operation.detail.query(资金授权操作查询接口)
// 文档地址: https://opendocs.alipay.com/open/02fkbd
func (c *Client) FundAuthOperationDetailQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.FundAuthOperationDetailQueryResponse, err error) {
	var bs []byte
	var method = "alipay.fund.auth.operation.detail.query"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.FundAuthOperationDetailQueryResponse)
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

// FundAuthOperationCancel
// alipay.fund.auth.operation.cancel(资金授权撤销接口)
// 文档地址: https://opendocs.alipay.com/open/02fkbb
func (c *Client) FundAuthOperationCancel(ctx context.Context, pl paypay.Payload) (aliRes *entity.FundAuthOperationCancelResponse, err error) {
	var bs []byte
	var method = "alipay.fund.auth.operation.cancel"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.FundAuthOperationCancelResponse)
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

// // alipay.fund.batch.create(批次下单接口)
// // 文档地址: https://opendocs.alipay.com/apis/api_28/alipay.fund.batch.create
// func (c *Client) FundBatchCreate(ctx context.Context, pl paypay.Payload) (aliRes *entity.FundBatchCreateResponse, err error) {
//	err = pl.CheckEmptyError("out_batch_no", "product_code", "biz_scene", "order_title", "total_trans_amount", "total_count", "trans_order_list")
//	if err != nil {
//		return nil, err
//	}
//	var bs []byte
//	if bs, err = c.callAli(ctx, pl, "alipay.fund.batch.create"); err != nil {
//		return nil, err
//	}
//	aliRes = new(entity.FundBatchCreateResponse)
//	if err = json.Unmarshal(bs, aliRes); err != nil || aliRes.Response == nil {
//		return nil, fmt.Errorf("[%w], bytes: %s", pkg.ErrUnmarshal, string(bs))
//	}
//	if err = utils.ExtractBussErr(aliRes.Response.ErrorResponse); err != nil {
//		return aliRes, err
//	}
//	signData, signDataErr := c.autoVerifySignByCert(ctx, bs, method, aliRes.AlipayCertSn)
//	aliRes.SignData = signData
//	return aliRes, signDataErr
// }

// // alipay.fund.batch.close(批量转账关单接口)
// // 文档地址: https://opendocs.alipay.com/apis/api_28/alipay.fund.batch.close
// func (c *Client) FundBatchClose(ctx context.Context, pl paypay.Payload) (aliRes *entity.FundBatchCloseResponse, err error) {
//	err = pl.CheckEmptyError("biz_scene")
//	if err != nil {
//		return nil, err
//	}
//	var bs []byte
//	if bs, err = c.callAli(ctx, pl, "alipay.fund.batch.close"); err != nil {
//		return nil, err
//	}
//	aliRes = new(entity.FundBatchCloseResponse)
//	if err = json.Unmarshal(bs, aliRes); err != nil || aliRes.Response == nil {
//		return nil, fmt.Errorf("[%w], bytes: %s", pkg.ErrUnmarshal, string(bs))
//	}
//	if err = utils.ExtractBussErr(aliRes.Response.ErrorResponse); err != nil {
//		return aliRes, err
//	}
//	signData, signDataErr := c.autoVerifySignByCert(ctx, bs, method, aliRes.AlipayCertSn)
//	aliRes.SignData = signData
//	return aliRes, signDataErr
// }

// // alipay.fund.batch.detail.query(批量转账明细查询接口)
// // 文档地址: https://opendocs.alipay.com/apis/api_28/alipay.fund.batch.detail.query
// func (c *Client) FundBatchDetailQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.FundBatchDetailQueryResponse, err error) {
//	err = pl.CheckEmptyError("biz_scene")
//	if err != nil {
//		return nil, err
//	}
//	var bs []byte
//	if bs, err = c.callAli(ctx, pl, "alipay.fund.batch.detail.query"); err != nil {
//		return nil, err
//	}
//	aliRes = new(entity.FundBatchDetailQueryResponse)
//	if err = json.Unmarshal(bs, aliRes); err != nil || aliRes.Response == nil {
//		return nil, fmt.Errorf("[%w], bytes: %s", pkg.ErrUnmarshal, string(bs))
//	}
//	if err = utils.ExtractBussErr(aliRes.Response.ErrorResponse); err != nil {
//		return aliRes, err
//	}
//	signData, signDataErr := c.autoVerifySignByCert(ctx, bs, method, aliRes.AlipayCertSn)
//	aliRes.SignData = signData
//	return aliRes, signDataErr
// }

// // alipay.fund.trans.app.pay(现金红包无线支付接口)
// // 文档地址: https://opendocs.alipay.com/open/03rbyf
// func (c *Client) FundTransAppPay(ctx context.Context, pl paypay.Payload) (aliRes *entity.FundTransAppPayResponse, err error) {
//	err = pl.CheckEmptyError("out_biz_no", "trans_amount", "product_code", "biz_scene")
//	if err != nil {
//		return nil, err
//	}
//	var bs []byte
//	if bs, err = c.callAli(ctx, pl, "alipay.fund.trans.app.pay"); err != nil {
//		return nil, err
//	}
//	aliRes = new(entity.FundTransAppPayResponse)
//	if err = json.Unmarshal(bs, aliRes); err != nil || aliRes.Response == nil {
//		return nil, fmt.Errorf("[%w], bytes: %s", pkg.ErrUnmarshal, string(bs))
//	}
//	if err = utils.ExtractBussErr(aliRes.Response.ErrorResponse); err != nil {
//		return aliRes, err
//	}
//	signData, signDataErr := c.autoVerifySignByCert(ctx, bs, method, aliRes.AlipayCertSn)
//	aliRes.SignData = signData
//	return aliRes, signDataErr
// }

// // alipay.fund.trans.payee.bind.query(资金收款账号绑定关系查询)
// // 文档地址: https://opendocs.alipay.com/apis/020tl1
// func (c *Client) FundTransPayeeBindQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.FundTransPayeeBindQueryResponse, err error) {
//	err = pl.CheckEmptyError("identity", "identity_type")
//	if err != nil {
//		return nil, err
//	}
//	var bs []byte
//	if bs, err = c.callAli(ctx, pl, "alipay.fund.trans.payee.bind.query"); err != nil {
//		return nil, err
//	}
//	aliRes = new(entity.FundTransPayeeBindQueryResponse)
//	if err = json.Unmarshal(bs, aliRes); err != nil || aliRes.Response == nil {
//		return nil, fmt.Errorf("[%w], bytes: %s", pkg.ErrUnmarshal, string(bs))
//	}
//	if err = utils.ExtractBussErr(aliRes.Response.ErrorResponse); err != nil {
//		return aliRes, err
//	}
//	signData, signDataErr := c.autoVerifySignByCert(ctx, bs, method, aliRes.AlipayCertSn)
//	aliRes.SignData = signData
//	return aliRes, signDataErr
// }

// // alipay.fund.trans.page.pay(资金转账页面支付接口)
// // 文档地址: https://opendocs.alipay.com/open/03rbye
// func (c *Client) FundTransPagePay(ctx context.Context, pl paypay.Payload) (aliRes *entity.FundTransPagePayResponse, err error) {
//	err = pl.CheckEmptyError("out_biz_no", "trans_amount", "product_code", "biz_scene")
//	if err != nil {
//		return nil, err
//	}
//	var bs []byte
//	if bs, err = c.callAli(ctx, pl, "alipay.fund.trans.page.pay"); err != nil {
//		return nil, err
//	}
//	aliRes = new(entity.FundTransPagePayResponse)
//	if err = json.Unmarshal(bs, aliRes); err != nil || aliRes.Response == nil {
//		return nil, fmt.Errorf("[%w], bytes: %s", pkg.ErrUnmarshal, string(bs))
//	}
//	if err = utils.ExtractBussErr(aliRes.Response.ErrorResponse); err != nil {
//		return aliRes, err
//	}
//	signData, signDataErr := c.autoVerifySignByCert(ctx, bs, method, aliRes.AlipayCertSn)
//	aliRes.SignData = signData
//	return aliRes, signDataErr
// }
