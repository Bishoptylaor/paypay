package aliClient

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Bishoptylaor/paypay"
	"github.com/Bishoptylaor/paypay/alipay/entity"
	"github.com/Bishoptylaor/paypay/alipay/utils"
	"github.com/Bishoptylaor/paypay/pkg"
	"net/url"
	"strings"
)

// UserAgreementPageSign
// alipay.user.agreement.page.sign(支付宝个人协议页面签约接口)
// 文档地址：https://opendocs.alipay.com/open/8bccfa0b_alipay.user.agreement.page.sign
func (c *Client) UserAgreementPageSign(ctx context.Context, pl paypay.Payload) (ret string, err error) {
	var bs []byte
	if bs, err = c.callAli(ctx, pl, "alipay.user.agreement.page.sign"); err != nil {
		return "", err
	}
	return string(bs), nil
}

// UserAgreementPageSignInApp
// alipay.user.agreement.page.sign(APP 支付宝个人协议页面签约接口)
// 文档地址：https://opendocs.alipay.com/open/00a05b  通过 App 唤起支付宝的签约页面
func (c *Client) UserAgreementPageSignInApp(ctx context.Context, pl paypay.Payload) (ret string, err error) {
	var s string
	// 参考官方示例
	// PageExecute get方式，生成url
	if s, err = c.PageExecute(ctx, pl, "alipay.user.agreement.page.sign"); err != nil {
		return "", err
	}

	// / 生成的url地址去除 http://openapi.alipay.com/gateway.do
	replaceUrl := c.Url() + "?"
	signParams := strings.Replace(s, replaceUrl, "", 1)

	// 该链接里面的 APPID 为固定值，不可修改；生成唤起客户端。把signParams使用 UTF-8 字符集整体做一次 encode
	link := "alipays://platformapi/startapp?appId=60000157&appClearTop=false&startMultApp=YES&sign_params=" + url.QueryEscape(signParams)
	return link, nil
}

// UserAgreementPageUnSign
// alipay.user.agreement.unsign(支付宝个人代扣协议解约接口)
// 文档地址：https://opendocs.alipay.com/open/b841da1f_alipay.user.agreement.unsign
func (c *Client) UserAgreementPageUnSign(ctx context.Context, pl paypay.Payload) (aliRes *entity.UserAgreementPageUnSignResponse, err error) {
	var bs []byte
	var method = "alipay.user.agreement.page.unsign"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.UserAgreementPageUnSignResponse)
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

// UserAgreementQuery
// alipay.user.agreement.query(支付宝个人代扣协议查询接口)
// 文档地址：https://opendocs.alipay.com/open/3dab71bc_alipay.user.agreement.query
func (c *Client) UserAgreementQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.UserAgreementQueryResponse, err error) {
	var bs []byte
	var method = "alipay.user.agreement.query"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.UserAgreementQueryResponse)
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

// UserAgreementExecutionplanModify
// alipay.user.agreement.executionplan.modify(周期性扣款协议执行计划修改接口)
// 文档地址：https://opendocs.alipay.com/open/ed428330_alipay.user.agreement.executionplan.modify
func (c *Client) UserAgreementExecutionplanModify(ctx context.Context, pl paypay.Payload) (aliRes *entity.UserAgreementExecutionplanModifyResponse, err error) {
	var bs []byte
	var method = "alipay.user.agreement.executionplan.modify"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.UserAgreementExecutionplanModifyResponse)
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

// UserAgreementTransfer
// alipay.user.agreement.transfer(协议由普通通用代扣协议产品转移到周期扣协议产品) 由商户调用，将商户之前通用代扣产品转移到周期扣的协议产品
// 文档地址：https://opendocs.alipay.com/open/02fkar
func (c *Client) UserAgreementTransfer(ctx context.Context, pl paypay.Payload) (aliRes *entity.UserAgreementTransferResponse, err error) {
	var bs []byte
	var method = "alipay.user.agreement.transfer"
	bs, err = c.callAli(ctx, pl, method)
	if err != nil {
		return nil, err
	}
	aliRes = new(entity.UserAgreementTransferResponse)
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
