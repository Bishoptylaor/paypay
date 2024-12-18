package alipay

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Bishoptylaor/paypay"
	"github.com/Bishoptylaor/paypay/alipay/consts"
	"github.com/Bishoptylaor/paypay/alipay/entity"
	"github.com/Bishoptylaor/paypay/alipay/utils"
	"github.com/Bishoptylaor/paypay/pkg"
)

// TradeRoyaltyRelationBind
// alipay.trade.royalty.relation.bind(分账关系绑定)
// 文档地址：https://opendocs.alipay.com/open/c21931d6_alipay.trade.royalty.relation.bind
func (c *Client) TradeRoyaltyRelationBind(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeRoyaltyRelationBindResponse, err error) {
	var bs []byte
	var method = "alipay.trade.royalty.relation.bind"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.TradeRoyaltyRelationBindResponse)
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

// TradeRoyaltyRelationUnbind
// alipay.trade.royalty.relation.unbind(分账关系解绑)
// 文档地址：https://opendocs.alipay.com/open/3613f4e1_alipay.trade.royalty.relation.unbind
func (c *Client) TradeRoyaltyRelationUnbind(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeRoyaltyRelationUnbindResponse, err error) {
	var bs []byte
	var method = "alipay.trade.royalty.relation.unbind"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.TradeRoyaltyRelationUnbindResponse)
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

// TradeRoyaltyRelationBatchquery
// alipay.trade.royalty.relation.batchquery(分账关系查询)
// 文档地址：https://opendocs.alipay.com/open/1860be54_alipay.trade.royalty.relation.batchquery
func (c *Client) TradeRoyaltyRelationBatchquery(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeRoyaltyRelationBatchqueryResponse, err error) {
	var bs []byte
	var method = "alipay.trade.royalty.relation.batchquery"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.TradeRoyaltyRelationBatchqueryResponse)
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

// TradeSettleConfirm
// alipay.trade.settle.confirm(统一收单确认结算接口)
// 文档地址：https://opendocs.alipay.com/open/028xqy
func (c *Client) TradeSettleConfirm(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeSettleConfirmResponse, err error) {
	c.EmptyChecker = paypay.InjectRuler(
		map[string][]paypay.Ruler{
			"alipay.trade.settle.confirm": []paypay.Ruler{
				paypay.NewRuler("确认结算请求流水号", "out_request_no != nil", fmt.Sprintf(consts.FmtEmptyAlert, "out_request_no")),
				paypay.NewRuler("支付宝交易号", "trade_no != nil", fmt.Sprintf(consts.FmtEmptyAlert, "trade_no")),
				paypay.NewRuler("描述结算信息",
					`settle_info != nil && len(settle_info?.settle_detail_infos) == 1 && `+
						`all(settle_info?.settle_detail_infos, {.trans_in_type in ["userId", "cardAliasNo", "loginName", "defaultSettle"]}) &&`+
						`none(settle_info?.settle_detail_infos, {.trans_in_type == "userId" && !(hasPrefix(.trans_in, "2088" && len(.trans_in) == 16))}) &&`+
						`none(settle_info?.settle_detail_infos, {.trans_in_type == "defaultSettle" && .trans_in != nil}) &&`+
						`all(settle_info?.settle_detail_infos, {.amount != nil})`,
					"",
				),
			},
		},
	)
	var bs []byte
	var method = "alipay.trade.settle.confirm"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.TradeSettleConfirmResponse)
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

// TradeOrderSettle
// alipay.trade.order.settle(统一收单交易结算接口)
// 文档地址：https://opendocs.alipay.com/open/c3b24498_alipay.trade.order.settle
func (c *Client) TradeOrderSettle(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeOrderSettleResponse, err error) {
	var bs []byte
	var method = "alipay.trade.order.settle"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.TradeOrderSettleResponse)
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

// TradeOrderSettleQuery
// alipay.trade.order.settle.query(交易分账查询接口)
// 文档地址：https://opendocs.alipay.com/open/02pj6l
func (c *Client) TradeOrderSettleQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeOrderSettleQueryResponse, err error) {
	var bs []byte
	var method = "alipay.trade.order.settle.query"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.TradeOrderSettleQueryResponse)
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

// TradeOrderOnsettleQuery
// alipay.trade.order.onsettle.query(分账剩余金额查询)
// 文档地址：https://opendocs.alipay.com/open/d87dc009_alipay.trade.order.onsettle.query
func (c *Client) TradeOrderOnsettleQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeOrderOnsettleQueryResponse, err error) {
	var bs []byte
	var method = "alipay.trade.order.onsettle.query"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.TradeOrderOnsettleQueryResponse)
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

// TradeRoyaltyRateQuery
// alipay.trade.royalty.rate.query(分账比例查询)
// 文档地址：https://opendocs.alipay.com/open/6f314ee9_alipay.trade.royalty.rate.query
func (c *Client) TradeRoyaltyRateQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.TradeRoyaltyRateQueryResponse, err error) {
	var bs []byte
	var method = "alipay.trade.royalty.rate.query"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.TradeRoyaltyRateQueryResponse)
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

// // alipay.pay.app.marketing.consult(商户前置内容咨询接口)
// // 文档地址：https://opendocs.alipay.com/pre-open/296d225f_alipay.pay.app.marketing.consult
// func (c *Client) PayAppMarketingConsult(ctx context.Context, pl paypay.Payload) (aliRes *entity.PayAppMarketingConsultResponse, err error) {
//	err = pl.CheckEmptyError("biz_scene", "total_amount", "product_code")
//	if err != nil {
//		return nil, err
//	}
//	var bs []byte
//	var method = "alipay.pay.app.marketing.consult"
//	if bs, err = c.callAli(ctx, pl, method); err != nil {
//		return nil, err
//	}
//	aliRes = new(entity.PayAppMarketingConsultResponse)
//	if err = json.Unmarshal(bs, aliRes); err != nil || aliRes.Response == nil {
//		return nil, fmt.Errorf("[%w], bytes: %s", pkg.ErrUnmarshal, string(bs))
//	}
//
//	if err = utils.ExtractBussErr(aliRes.Response.ErrorResponse); err != nil {
//		return aliRes, err
//	}
//	signData, signDataErr := c.autoVerifySignByCert(ctx, bs, method, aliRes.AlipayCertSn)
//	aliRes.SignData = signData
//	return aliRes, signDataErr
// }
