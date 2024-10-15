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
	getKV func() map[string]string,
	pl paypay.Payload,
	query paypay.Payload,
	patches []*entity.Patch,
	res interface{},
	response interface{},
	ops ...Settings) error {

	// c2 is a copied client only in this func
	c2 := &Client{
		Config:   c.Config,
		Operates: c.Operates,
	}
	for _, op := range append(c.ops, ops...) {
		op(c2)
	}

	c2.EmptyChecker = method.Checker
	_kv := getKV()
	for _, v := range _kv {
		if v == pkg.NULL {
			return pkg.ErrPaypalMissingQueryId
		}
	}
	// it is useful when method has a placeholder
	_kv["params"] = func() string {
		if query != nil {
			return query.EncodeURLParams()
		}
		return ""
	}()

	httpRes, bs, err := method.Do(c2)(ctx, method.Uri, c2.GenUrl(ctx, _kv), pl, patches)
	if err != nil {
		return err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.AddTrackerForOrderRes{EmptyRes: emptyRes}
	if err = c2.handleResponse(ctx, method, httpRes, bs, &emptyRes, response); err != nil {
		return err
	}
	return nil
}
