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
 @Time    : 2024/10/14 -- 18:47
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: singlecall.go
*/

package paypal

import (
	"context"
	"github.com/Bishoptylaor/paypay"
	"github.com/Bishoptylaor/paypay/paypal/consts"
	"github.com/Bishoptylaor/paypay/paypal/entity"
	"github.com/Bishoptylaor/paypay/pkg"
)

// CustomSingleCall
// single call with custom settings
// you can check out usage in examples
// ALERT: DO NOT support changes on ClientID or Secret
func (c *Client) CustomSingleCall(
	ctx context.Context,
	method Method,
	getKV func() (string, string),
	pl paypay.Payload,
	query paypay.Payload,
	patches []*entity.Patch,
	res interface{},
	response interface{},
	ops ...Settings) error {

	// generate a new client only in this func
	c2 := &Client{
		Config:   c.Config,
		Operates: c.Operates,
	}
	for _, op := range append(c.ops, ops...) {
		op(c2)
	}

	c.EmptyChecker = method.Checker
	k, v := getKV()
	if k != pkg.NULL && v == pkg.NULL {
		return pkg.ErrPaypalMissingQueryId
	}

	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		k: v,
		"params": func() string {
			if query != nil {
				return query.EncodeURLParams()
			}
			return ""
		}(), // it is useful when method has a placeholder
	}), pl, patches)
	if err != nil {
		return err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.AddTrackerForOrderRes{EmptyRes: emptyRes}
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, response); err != nil {
		return err
	}
	return nil
}
