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
 @Time    : 2024/8/19 -- 11:28
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

// // alipay.open.app.qrcode.create(小程序生成推广二维码接口)
// // 文档地址：https://opendocs.alipay.com/apis/009zva
// func (c *Client) OpenAppQrcodeCreate(ctx context.Context, pl paypay.Payload) (aliRes *OpenAppQrcodeCreateRsp, err error) {
//	err = pl.CheckEmptyError("url_param", "query_param", "describe")
//	if err != nil {
//		return nil, err
//	}
//	var bs []byte
//	if bs, err = c.callAli(ctx, pl, "alipay.open.app.qrcode.create"); err != nil {
//		return nil, err
//	}
//	aliRes = new(OpenAppQrcodeCreateRsp)
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
//
// // alipay.marketing.campaign.cash.create(创建现金活动接口)
// // 文档地址：https://opendocs.alipay.com/open/029yy9
// func (c *Client) MarketingCampaignCashCreate(ctx context.Context, pl paypay.Payload) (aliRes *MarketingCampaignCashCreateRsp, err error) {
//	err = pl.CheckEmptyError("coupon_name", "prize_type", "total_money", "total_num", "prize_msg", "start_time", "end_time", "merchant_link")
//	if err != nil {
//		return nil, err
//	}
//	var bs []byte
//	if bs, err = c.callAli(ctx, pl, "alipay.marketing.campaign.cash.create"); err != nil {
//		return nil, err
//	}
//	aliRes = new(MarketingCampaignCashCreateRsp)
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
//
// // alipay.marketing.campaign.cash.trigger(触发现金红包活动)
// // 文档地址：https://opendocs.alipay.com/open/029yya
// func (c *Client) MarketingCampaignCashTrigger(ctx context.Context, pl paypay.Payload) (aliRes *MarketingCampaignCashTriggerRsp, err error) {
//	err = pl.CheckEmptyError("user_id", "crowd_no", "login_id", "order_price", "out_biz_no")
//	if err != nil {
//		return nil, err
//	}
//	var bs []byte
//	if bs, err = c.callAli(ctx, pl, "alipay.marketing.campaign.cash.trigger"); err != nil {
//		return nil, err
//	}
//	aliRes = new(MarketingCampaignCashTriggerRsp)
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
//
// // alipay.marketing.campaign.cash.status.modify(更改现金活动状态接口)
// // 文档地址：https://opendocs.alipay.com/open/029yyb
// func (c *Client) MarketingCampaignCashStatusModify(ctx context.Context, pl paypay.Payload) (aliRes *MarketingCampaignCashStatusModifyRsp, err error) {
//	err = pl.CheckEmptyError("crowd_no", "camp_status")
//	if err != nil {
//		return nil, err
//	}
//	var bs []byte
//	if bs, err = c.callAli(ctx, pl, "alipay.marketing.campaign.cash.status.modify"); err != nil {
//		return nil, err
//	}
//	aliRes = new(MarketingCampaignCashStatusModifyRsp)
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

// MarketingCampaignCashListQuery
// alipay.marketing.campaign.cash.list.query(现金活动列表查询接口)
// 文档地址：https://opendocs.alipay.com/open/02a1f9
func (c *Client) MarketingCampaignCashListQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.MarketingCampaignCashListQueryResponse, err error) {
	var bs []byte
	var method = "alipay.marketing.campaign.cash.list.query"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.MarketingCampaignCashListQueryResponse)
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

// MarketingCampaignCashDetailQuery
// alipay.marketing.campaign.cash.detail.query(现金活动详情查询)
// 文档地址：https://opendocs.alipay.com/open/02a1fa
func (c *Client) MarketingCampaignCashDetailQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.MarketingCampaignCashDetailQueryResponse, err error) {
	var bs []byte
	var method = "alipay.marketing.campaign.cash.detail.query"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.MarketingCampaignCashDetailQueryResponse)
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

// MarketingCampaignOrderVoucherConsult
// alipay.marketing.campaign.order.voucher.consult(订单优惠前置咨询)
// 文档地址：https://opendocs.alipay.com/open/04fgwi
func (c *Client) MarketingCampaignOrderVoucherConsult(ctx context.Context, pl paypay.Payload) (aliRes *entity.MarketingCampaignOrderVoucherConsultResponse, err error) {
	var bs []byte
	var method = "alipay.marketing.campaign.order.voucher.consult"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.MarketingCampaignOrderVoucherConsultResponse)
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

// MarketingMaterialImageUpload
// alipay.marketing.material.image.upload(营销图片资源上传接口)
// 文档地址：https://opendocs.alipay.com/open/389b24b6_alipay.marketing.material.image.upload
func (c *Client) MarketingMaterialImageUpload(ctx context.Context, pl paypay.Payload) (aliRes *entity.MarketingMaterialImageUploadResponse, err error) {
	var bs []byte
	var method = "alipay.marketing.material.image.upload"
	if bs, err = c.FileUploadRequest(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.MarketingMaterialImageUploadResponse)
	if err = json.Unmarshal(bs, aliRes); err != nil {
		return nil, err
	}
	if aliRes.Response != nil && aliRes.Response.Code != "10000" {
		info := aliRes.Response
		return aliRes, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	signData, signDataErr := c.autoVerifySignByCert(ctx, bs, method, aliRes.AlipayCertSn)
	aliRes.SignData = signData
	return aliRes, signDataErr
}
