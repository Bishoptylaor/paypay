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
 @Time    : 2024/8/20 -- 10:55
 @Author  : 亓官竹
 @Description: data.go
*/

package aliClient

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Bishoptylaor/paypay"
	"github.com/Bishoptylaor/paypay/alipay/entity"
	"github.com/Bishoptylaor/paypay/alipay/utils"
	"github.com/Bishoptylaor/paypay/pkg"
)

// DataBillEreceiptApply
// alipay.data.bill.ereceipt.apply(申请电子回单(incubating))
// 文档地址：https://opendocs.alipay.com/open/1aad1956_alipay.data.bill.ereceipt.apply
func (c *Client) DataBillEreceiptApply(ctx context.Context, pl paypay.Payload) (aliRes *entity.DataBillEreceiptApplyResponse, err error) {
	var bs []byte
	var method = "alipay.data.bill.ereceipt.apply"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.DataBillEreceiptApplyResponse)
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

// DataBillEreceiptQuery
// alipay.data.bill.ereceipt.query(查询电子回单状态(incubating))
// 文档地址：https://opendocs.alipay.com/open/30b94a2f_alipay.data.bill.ereceipt.query
func (c *Client) DataBillEreceiptQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.DataBillEreceiptQueryResponse, err error) {
	var bs []byte
	var method = "alipay.data.bill.ereceipt.query"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.DataBillEreceiptQueryResponse)
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

// DataBillDownloadUrlQuery
// alipay.data.dataservice.bill.downloadurl.query(查询对账单下载地址)
// 文档地址：https://opendocs.alipay.com/open/02e7gr
func (c *Client) DataBillDownloadUrlQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.DataBillDownloadUrlQueryResponse, err error) {
	var bs []byte
	var method = "alipay.data.dataservice.bill.downloadurl.query"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.DataBillDownloadUrlQueryResponse)
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

// DataBillBalanceQuery
// alipay.data.bill.balance.query(支付宝商家账户当前余额查询)
// 文档地址：https://opendocs.alipay.com/open/2acb3c34_alipay.data.bill.balance.query
func (c *Client) DataBillBalanceQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.DataBillBalanceQueryResponse, err error) {
	var bs []byte
	var method = "alipay.data.bill.balance.query"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.DataBillBalanceQueryResponse)
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

// DataBillBalancehisQuery
// alipay.data.bill.balancehis.query(支付宝商家账户历史余额查询)
// 文档地址：https://opendocs.alipay.com/open/2cb36cd5_alipay.data.bill.balancehis.query
func (c *Client) DataBillBalancehisQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.DataBillBalancehisQueryResponse, err error) {
	var bs []byte
	var method = "alipay.data.bill.balancehis.query"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.DataBillBalancehisQueryResponse)
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

// DataBillAccountlogQuery
// alipay.data.bill.accountlog.query(支付宝商家账户账务明细查询)
// 文档地址：https://opendocs.alipay.com/open/dae3ac99_alipay.data.bill.accountlog.query
func (c *Client) DataBillAccountlogQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.DataBillAccountlogQueryResponse, err error) {
	var bs []byte
	var method = "alipay.data.bill.accountlog.query"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.DataBillAccountlogQueryResponse)
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

// DataBillSellQuery
// alipay.data.bill.sell.query(支付宝商家账户卖出交易查询)
// 文档地址：https://opendocs.alipay.com/open/8a737327_alipay.data.bill.sell.query
func (c *Client) DataBillSellQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.DataBillSellQueryResponse, err error) {
	var bs []byte
	var method = "alipay.data.bill.sell.query"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.DataBillSellQueryResponse)
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

// DataBillBuyQuery
// alipay.data.bill.buy.query(支付宝商家账户买入交易查询)
// 文档地址：https://opendocs.alipay.com/open/4e92e4e7_alipay.data.bill.buy.query
func (c *Client) DataBillBuyQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.DataBillBuyQueryResponse, err error) {
	var bs []byte
	var method = "alipay.data.bill.buy.query"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.DataBillBuyQueryResponse)
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

// DataBillTransferQuery
// alipay.data.bill.transfer.query(支付宝商家账户充值，转账，提现查询)
// 文档地址：https://opendocs.alipay.com/open/0d2f1256_alipay.data.bill.transfer.query
func (c *Client) DataBillTransferQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.DataBillTransferQueryResponse, err error) {
	var bs []byte
	var method = "alipay.data.bill.transfer.query"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.DataBillTransferQueryResponse)
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

// DataBillBailQuery
// alipay.data.bill.bail.query(支付宝商家账户保证金查询)
// 文档地址：https://opendocs.alipay.com/open/10e1feb7_alipay.data.bill.bail.query
func (c *Client) DataBillBailQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.DataBillBailQueryResponse, err error) {
	var bs []byte
	var method = "alipay.data.bill.bail.query"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.DataBillBailQueryResponse)
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
