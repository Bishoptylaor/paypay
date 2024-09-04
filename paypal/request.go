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
 @Time    : 2024/9/3 -- 17:50
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: request.go
*/

package paypal

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/Bishoptylaor/paypay"
	"github.com/Bishoptylaor/paypay/paypal/consts"
	"github.com/Bishoptylaor/paypay/paypal/entity"
	"github.com/Bishoptylaor/paypay/pkg"
	"github.com/Bishoptylaor/paypay/pkg/xlog"
	"github.com/Bishoptylaor/paypay/pkg/xnet/xhttp"
	"io"
	"net/http"
	"net/url"
)

func (c *Client) doPayPalGet(ctx context.Context, path string) (res *http.Response, bs []byte, err error) {
	return c.doPayPal(ctx, xhttp.Get(c.GenUrl(path)), "")
}

func (c *Client) doPayPalPost(ctx context.Context, pl paypay.Payload, path string) (res *http.Response, bs []byte, err error) {
	err = paypay.ExecuteQueue(
		// 校验 biz_content 参数规则
		IntegrityCheck(ctx, c, path),
	)(pl)
	if err != nil {
		return nil, nil, err
	}
	return c.doPayPal(ctx, xhttp.Post(c.GenUrl(path)), pl)
}

func (c *Client) doPayPalPut(ctx context.Context, pl paypay.Payload, path string) (res *http.Response, bs []byte, err error) {
	err = paypay.ExecuteQueue(
		// 校验 biz_content 参数规则
		IntegrityCheck(ctx, c, path),
	)(pl)
	if err != nil {
		return nil, nil, err
	}
	return c.doPayPal(ctx, xhttp.Put(c.GenUrl(path)), pl)
}

func (c *Client) doPayPalPatch(ctx context.Context, patchs []*entity.Patch, path string) (res *http.Response, bs []byte, err error) {
	body, _ := json.Marshal(patchs)
	return c.doPayPal(ctx, xhttp.Patch(c.GenUrl(path)), body)
}

func (c *Client) doPayPalDelete(ctx context.Context, path string) (res *http.Response, bs []byte, err error) {
	return c.doPayPal(ctx, xhttp.Delete(c.GenUrl(path)), "")
}

func (c *Client) doPayPal(ctx context.Context, method xhttp.CfgOp, data any) (res *http.Response, bs []byte, err error) {
	res, bs, err = c.HClient.CallOp(ctx, data,
		xhttp.Req(xhttp.TypeJSON), // default json
		method,
		xhttp.Header(map[string]string{
			"Accept":                   "*/*",
			consts.HeaderAuthorization: consts.AuthorizationPrefixBearer + c.AccessToken,
		}),
		xhttp.Prefix(PPReqPrefix(c.debug, c.Logger)),
		xhttp.Suffix(PPResSuffix(c.debug, c.Logger)),
	)
	if err != nil {
		return nil, nil, err
	}
	return res, bs, nil
}

// PPReqPrefix 闭包注入 logger 和 debug 信息
func PPReqPrefix(debug bool, log xlog.ZLogger) xhttp.ReqPrefixFunc {
	return func(ctx context.Context, req *http.Request) context.Context {
		if debug == pkg.DebugOn {
			log.Debugf("PayPal_Url: %s", req.URL)
			log.Debugf("PayPal_Req_Body: %s", req.Body)
			log.Debugf("PayPal_Req_Headers: %#v", req.Header)
		} else {
			body, err := io.ReadAll(req.Body)
			if err != nil {
				log.Errorf("[Read Req body] error: %s", err.Error())
				return ctx
			}
			enEscapeUrl, err := url.QueryUnescape(string(body))
			if err == nil {
				log.Infof("[Req] %s", enEscapeUrl)
			}
			req.Body = io.NopCloser(bytes.NewBuffer(body))
		}
		return ctx
	}
}

// PPResSuffix 闭包注入 logger 和 debug 信息
func PPResSuffix(debug bool, log xlog.ZLogger) xhttp.ResSuffixFunc {
	return func(ctx context.Context, res *http.Response) context.Context {
		if debug == pkg.DebugOn {
			log.Debugf("PayPal_Response: %d > %s", res.StatusCode, res.Body)
			log.Debugf("PayPal_Rsp_Headers: %#v", res.Header)
		} else {
			body, err := io.ReadAll(res.Body)
			if err != nil {
				log.Errorf("[Read Res body] error: %s", err.Error())
				return ctx
			}
			res.Body = io.NopCloser(bytes.NewBuffer(body))
			log.Infof("[Res] %s", string(body))
		}
		return ctx
	}
}
