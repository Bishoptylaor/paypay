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
 @Time    : 2024/8/26 -- 18:42
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: security.go
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

// SecurityRiskCustomerriskSend
// alipay.security.risk.customerrisk.send(商户数据同步)
// 文档地址：https://opendocs.alipay.com/open/02qth4
func (c *Client) SecurityRiskCustomerriskSend(ctx context.Context, pl paypay.Payload) (aliRes *entity.SecurityRiskCustomerriskSendResponse, err error) {
	// err = pl.CheckEmptyError("process_code", "trade_no")
	// if err != nil {
	//	return nil, err
	// }
	var bs []byte
	var method = "alipay.security.risk.customerrisk.send"
	if bs, err = c.callAli(ctx, pl, method); err != nil {
		return nil, err
	}
	aliRes = new(entity.SecurityRiskCustomerriskSendResponse)
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
