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
	"errors"
	"github.com/Bishoptylaor/paypay"
	"github.com/Bishoptylaor/paypay/paypal/entity"
	"github.com/Bishoptylaor/paypay/pkg"
)

// CustomCallOnce
// single call with custom settings
// @params: ctx = context of any function caller
//
// @params: method = Paypal request predefined info
//
// @params: getKv = placeholder and target value of uri eg:
//     /v2/checkout/orders/{{.order_id}}?{{.params}}
//     getKv can be : func() map[string]string {return map[string]string{OrderId.String(): "yourId"}}
//     you can skip params because its a build-in placeholder no need concerned from outside of the function.
//     will return an error if getKv is nil.
//
// @params: pl = payload you build to make a http request. should be nil if no use.
//
// @params: query = will transform to urlparams and attached to uri in method. should be nil if no use.
//
// @params: patches = payload of your changes in update* requests. should be nil if no use.
//
// @params: emptyRes = err msg in case there`s sth. wrong with the request.
//
// @params: response
//
// @params: ops = custom settings you want to use in this request.
//
// you can check out usage in examples
func (c *Client) CustomCallOnce(
	ctx context.Context,
	method Method,
	getKV func() map[string]string,
	pl paypay.Payload,
	query paypay.Payload,
	patches []*entity.Patch,
	emptyRes entity.EmptyRes,
	response interface{},
	ops ...Settings) error {

	if getKV == nil {
		return errors.New("getKV function is required")
	}

	// c2 is a copied client only in this func
	// use original client`s config as default
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
	if err = c2.handleResponse(ctx, method, httpRes, bs, &emptyRes, response); err != nil {
		return err
	}
	return nil
}
