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

// TradePay
// alipay.trade.pay(统一收单交易支付接口)
// 文档地址：https://opendocs.alipay.com/open/02cdx8
func (c *Client) TradePay(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradePayResponse, err error) {
	var bs []byte
	var method = "alipay.trade.pay"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.TradePayResponse)
	if err = json.Unmarshal(bs, aliRes); err != nil || aliRes.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", pkg.ErrUnmarshal, string(bs))
	}
	if err = utils.ExtractBussErrFunc(aliRes.Response.ErrorResponse, func() []string {
		return []string{"10000", "10003"}
	}); err != nil {
		return aliRes, err
	}
	signData, signDataErr := c.autoVerifySignByCert(ctx, bs, method, aliRes.AlipayCertSn)
	aliRes.SignData = signData
	return aliRes, signDataErr
}

// TradePrecreate
// alipay.trade.precreate(统一收单线下交易预创建)
// 文档地址：https://opendocs.alipay.com/open/02ekfg
func (c *Client) TradePrecreate(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradePrecreateResponse, err error) {
	var bs []byte
	var method = "alipay.trade.precreate"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.TradePrecreateResponse)
	if err = json.Unmarshal(bs, aliRes); err != nil || aliRes.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", pkg.ErrUnmarshal, string(bs))
	}
	if err = utils.ExtractBussErr(aliRes.Response.ErrorResponse); err != nil {
		return aliRes, err
	}
	if aliRes.NullResponse != nil {
		info := aliRes.NullResponse
		return aliRes, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	signData, signDataErr := c.autoVerifySignByCert(ctx, bs, method, aliRes.AlipayCertSn)
	aliRes.SignData = signData
	return aliRes, signDataErr
}

// TradeAppPay
// alipay.trade.app.pay(app支付接口2.0)
// 文档地址：https://opendocs.alipay.com/open/02e7gq
func (c *Client) TradeAppPay(ctx context.Context, pl paypay.Payload) (payParam string, err error) {
	var bs []byte
	var method = "alipay.trade.app.pay"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return pkg.NULL, err
	}
	payParam = string(bs)
	return payParam, nil
}

// TradeWapPay
// alipay.trade.wap.pay(手机网站支付接口2.0)
// 文档地址：https://opendocs.alipay.com/open/02ivbs
func (c *Client) TradeWapPay(ctx context.Context, pl paypay.Payload) (payUrl string, err error) {
	var bs []byte
	if bs, err = c.callAli(ctx, pl, "alipay.trade.wap.pay"); err != nil {
		return pkg.NULL, err
	}
	payUrl = string(bs)
	return payUrl, nil
}

// TradePagePay
// alipay.trade.page.pay(统一收单下单并支付页面接口)
// 文档地址：https://opendocs.alipay.com/open/028r8t
func (c *Client) TradePagePay(ctx context.Context, pl paypay.Payload) (payUrl string, err error) {
	var bs []byte
	var method = "alipay.trade.page.pay"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return pkg.NULL, err
	}
	payUrl = string(bs)
	return payUrl, nil
}

// TradeCreate
// alipay.trade.create(统一收单交易创建接口)
// 文档地址：https://opendocs.alipay.com/open/02ekfj
func (c *Client) TradeCreate(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeCreateResponse, err error) {
	var bs []byte
	var method = "alipay.trade.create"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.TradeCreateResponse)
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

// TradeOrderPay
// alipay.trade.order.pay(统一收单交易订单支付接口)
// 文档地址：https://opendocs.alipay.com/open/03vtew
func (c *Client) TradeOrderPay(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeOrderPayResponse, err error) {
	var bs []byte
	var method = "alipay.trade.order.pay"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.TradeOrderPayResponse)
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

// TradeWapMergePay
// alipay.trade.wap.merge.pay(无线Wap合并支付接口2.0)
// 文档地址：https://opendocs.alipay.com/solution/fc0cb136_alipay.trade.wap.merge.pay
func (c *Client) TradeWapMergePay(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeWapMergePayResponse, err error) {
	var bs []byte
	var method = "alipay.trade.wap.merge.pay"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.TradeWapMergePayResponse)
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

// TradeQuery
// alipay.trade.query(统一收单线下交易查询)
// 文档地址：https://opendocs.alipay.com/open/02e7gm
func (c *Client) TradeQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeQueryResponse, err error) {
	var bs []byte
	var method = "alipay.trade.query"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.TradeQueryResponse)
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

// TradeCancel
// alipay.trade.cancel(统一收单交易撤销接口)
// 文档地址：https://opendocs.alipay.com/open/02ekfi
func (c *Client) TradeCancel(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeCancelResponse, err error) {
	var bs []byte
	var method = "alipay.trade.cancel"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.TradeCancelResponse)
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

// TradeClose
// alipay.trade.close(统一收单交易关闭接口)
// 文档地址：https://opendocs.alipay.com/open/02e7gn
func (c *Client) TradeClose(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeCloseResponse, err error) {
	var bs []byte
	var method = "alipay.trade.close"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.TradeCloseResponse)
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

// TradeRefund
// alipay.trade.refund(统一收单交易退款接口)
// 文档地址：https://opendocs.alipay.com/open/02e7go
func (c *Client) TradeRefund(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeRefundResponse, err error) {
	var bs []byte
	var method = "alipay.trade.refund"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.TradeRefundResponse)
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

// // TradePageRefund
// // alipay.trade.page.refund(统一收单退款页面接口)
// // 文档地址：https://opendocs.alipay.com/apis/api_1/alipay.trade.page.refund
// func (c *Client) TradePageRefund(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradePageRefundResponse, err error) {
//	if pl.GetString("out_trade_no") == paypay.NULL && pl.GetString("trade_no") == paypay.NULL {
//		return nil, errors.New("out_trade_no and trade_no are not allowed to be null at the same time")
//	}
//	err = pl.CheckEmptyError("out_request_no", "refund_amount")
//	if err != nil {
//		return nil, err
//	}
//	var bs []byte
//	var method = "alipay.trade.page.refund"
//	if bs, err = c.callAli(ctx, pl, method); err != nil {
//		return nil, err
//	}
//	aliRes = new(entity.TradePageRefundResponse)
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

// TradeFastPayRefundQuery
// alipay.trade.fastpay.refund.query(统一收单交易退款查询)
// 文档地址：https://opendocs.alipay.com/open/02e7gp
func (c *Client) TradeFastPayRefundQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeFastpayRefundQueryResponse, err error) {
	var bs []byte
	var method = "alipay.trade.fastpay.refund.query"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.TradeFastpayRefundQueryResponse)
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

// TradeOrderInfoSync
// alipay.trade.orderinfo.sync(支付宝订单信息同步接口)
// 文档地址：https://opendocs.alipay.com/open/02cnou
func (c *Client) TradeOrderInfoSync(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeOrderInfoSyncResponse, err error) {
	var bs []byte
	var method = "alipay.trade.orderinfo.sync"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.TradeOrderInfoSyncResponse)
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

// TradeAdvanceConsult
// alipay.trade.advance.consult(订单咨询服务)
// 文档地址：https://opendocs.alipay.com/apis/api_1/alipay.trade.advance.consult
func (c *Client) TradeAdvanceConsult(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeAdvanceConsultResponse, err error) {
	var bs []byte
	var method = "alipay.trade.advance.consult"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.TradeAdvanceConsultResponse)
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

// TradeRepaybillQuery
// alipay.trade.repaybill.query(还款账单查询)
// 文档地址：https://opendocs.alipay.com/apis/api_1/alipay.trade.repaybill.query
func (c *Client) TradeRepaybillQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeRepaybillQueryResponse, err error) {
	var bs []byte
	var method = "alipay.trade.repaybill.query"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.TradeRepaybillQueryResponse)
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
