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

// ZolozAuthenticationSmilepayInitialize
// zoloz.authentication.smilepay.initialize(刷脸支付初始化)
// 文档地址：https://opendocs.alipay.com/open/2f7c1d5f_zoloz.authentication.smilepay.initialize
func (c *Client) ZolozAuthenticationSmilepayInitialize(ctx context.Context, pl paypay.Payload) (aliRes *entity.ZolozAuthenticationSmilepayInitializeResponse, err error) {
	var bs []byte
	var method = "zoloz.authentication.smilepay.initialize"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.ZolozAuthenticationSmilepayInitializeResponse)
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

// ZolozAuthenticationCustomerFtokenQuery
// zoloz.authentication.customer.ftoken.query(查询刷脸结果信息接口)
// 文档地址：https://opendocs.alipay.com/open/c8e4d285_zoloz.authentication.customer.ftoken.query
func (c *Client) ZolozAuthenticationCustomerFtokenQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.ZolozAuthenticationCustomerFtokenQueryResponse, err error) {
	var bs []byte
	var method = "zoloz.authentication.customer.ftoken.query"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.ZolozAuthenticationCustomerFtokenQueryResponse)
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
