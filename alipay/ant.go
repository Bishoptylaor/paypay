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
 @Time    : 2024/8/26 -- 11:17
 @Author  : 亓官竹
 @Copyright 2024 亓官竹
 @Description: ant.x 相关接口
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

// AntMerchantExpandShopCreate
// ant.merchant.expand.shop.create(蚂蚁店铺创建)
// 文档地址：https://opendocs.alipay.com/apis/api_1/ant.merchant.expand.shop.create
func (c *Client) AntMerchantExpandShopCreate(ctx context.Context, pl paypay.Payload) (aliRes *entity.AntMerchantExpandShopCreateResponse, err error) {
	var bs []byte
	var method = "ant.merchant.expand.shop.create"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.AntMerchantExpandShopCreateResponse)
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

// AntMerchantExpandShopModify
// ant.merchant.expand.shop.modify(修改蚂蚁店铺)
// 文档地址：https://opendocs.alipay.com/apis/014tmb
func (c *Client) AntMerchantExpandShopModify(ctx context.Context, pl paypay.Payload) (aliRes *entity.AntMerchantExpandShopModifyResponse, err error) {
	var bs []byte
	var method = "ant.merchant.expand.shop.modify"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.AntMerchantExpandShopModifyResponse)
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

// AntMerchantExpandOrderQuery
// ant.merchant.expand.order.query(商户申请单查询)
// 文档地址：https://opendocs.alipay.com/apis/api_1/ant.merchant.expand.order.query
func (c *Client) AntMerchantExpandOrderQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.AntMerchantExpandOrderQueryResponse, err error) {
	var bs []byte
	var method = "ant.merchant.expand.order.query"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.AntMerchantExpandOrderQueryResponse)
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

// AntMerchantExpandShopPageQuery
// ant.merchant.expand.shop.page.query(店铺分页查询接口)
// 文档地址：https://opendocs.alipay.com/open/04fgwq
func (c *Client) AntMerchantExpandShopPageQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.AntMerchantExpandShopPageQueryResponse, err error) {
	var bs []byte
	var method = "ant.merchant.expand.shop.page.query"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.AntMerchantExpandShopPageQueryResponse)
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

// AntMerchantExpandShopQuery
// ant.merchant.expand.shop.query(店铺查询接口)
// 文档地址：https://opendocs.alipay.com/apis/api_1/ant.merchant.expand.shop.query
func (c *Client) AntMerchantExpandShopQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.AntMerchantExpandShopQueryResponse, err error) {
	var bs []byte
	var method = "ant.merchant.expand.shop.query"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.AntMerchantExpandShopQueryResponse)
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

// AntMerchantExpandShopClose
// ant.merchant.expand.shop.close(蚂蚁店铺关闭)
// 文档地址：https://opendocs.alipay.com/apis/api_1/ant.merchant.expand.shop.close
func (c *Client) AntMerchantExpandShopClose(ctx context.Context, pl paypay.Payload) (aliRes *entity.AntMerchantExpandShopCloseResponse, err error) {
	var bs []byte
	var method = "ant.merchant.expand.shop.close"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.AntMerchantExpandShopCloseResponse)
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

// AntMerchantExpandIndirectImageUpload
// ant.merchant.expand.indirect.image.upload(图片上传)
// pl参数中 image_content 可不传，file为必传参数
// 文档地址：https://opendocs.alipay.com/open/04fgwt
func (c *Client) AntMerchantExpandIndirectImageUpload(ctx context.Context, pl paypay.Payload) (aliRes *entity.AntMerchantExpandIndirectImageUploadResponse, err error) {
	var bs []byte
	var method = "ant.merchant.expand.indirect.image.upload"
	if bs, err = c.FileUploadRequest(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.AntMerchantExpandIndirectImageUploadResponse)
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

// AntMerchantExpandMccQuery
// ant.merchant.expand.mcc.query(商户mcc信息查询)
// 文档地址：https://opendocs.alipay.com/open/04fgwu
func (c *Client) AntMerchantExpandMccQuery(ctx context.Context, pl paypay.Payload) (aliRes *entity.AntMerchantExpandMccQueryResponse, err error) {
	var bs []byte
	var method = "ant.merchant.expand.mcc.query"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.AntMerchantExpandMccQueryResponse)
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

// AntMerchantExpandShopConsult
// ant.merchant.expand.shop.consult(蚂蚁店铺创建咨询)
// 文档地址：https://opendocs.alipay.com/apis/014yig
func (c *Client) AntMerchantExpandShopConsult(ctx context.Context, pl paypay.Payload) (aliRes *entity.AntMerchantExpandShopConsultResponse, err error) {
	var bs []byte
	var method = "ant.merchant.expand.shop.consult"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.AntMerchantExpandShopConsultResponse)
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

// AntMerchantExpandShopReceiptaccountSave
// ant.merchant.expand.shop.receiptaccount.save(店铺增加收单账号)
// 文档地址：https://opendocs.alipay.com/open/54b69b89_ant.merchant.expand.shop.receiptaccount.save
func (c *Client) AntMerchantExpandShopReceiptaccountSave(ctx context.Context, pl paypay.Payload) (aliRes *entity.AntMerchantExpandShopReceiptaccountSaveResponse, err error) {
	var bs []byte
	var method = "ant.merchant.expand.shop.receiptaccount.save"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.AntMerchantExpandShopReceiptaccountSaveResponse)
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
