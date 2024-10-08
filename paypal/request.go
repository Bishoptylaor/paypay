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

type DoPayPalRequest func(ctx context.Context, uri string, urlGenerator func(string) string, pl paypay.Payload, patches []*entity.Patch) (res *http.Response, bs []byte, err error)

func GetPayPal(c *Client) DoPayPalRequest {
	return func(ctx context.Context, uri string, urlGenerator func(string) string, _ paypay.Payload, _ []*entity.Patch) (res *http.Response, bs []byte, err error) {
		return c.doPayPal(ctx, xhttp.Get(urlGenerator(uri)), "", nil)
	}
}

func PostPayPal(c *Client) DoPayPalRequest {
	return func(ctx context.Context, uri string, urlGenerator func(string) string, pl paypay.Payload, _ []*entity.Patch) (res *http.Response, bs []byte, err error) {
		err = paypay.ExecuteQueue(
			// 校验 biz_content 参数规则
			IntegrityCheck(ctx, c, uri),
		)(pl)
		if err != nil {
			return nil, nil, err
		}
		return c.doPayPal(ctx, xhttp.Post(urlGenerator(uri)), pl, nil)
	}
}

func PutPayPal(c *Client) DoPayPalRequest {
	return func(ctx context.Context, uri string, urlGenerator func(string) string, pl paypay.Payload, _ []*entity.Patch) (res *http.Response, bs []byte, err error) {
		err = paypay.ExecuteQueue(
			// 校验参数规则
			IntegrityCheck(ctx, c, uri),
		)(pl)
		if err != nil {
			return nil, nil, err
		}
		return c.doPayPal(ctx, xhttp.Put(urlGenerator(uri)), pl, nil)
	}
}

func PatchPayPal(c *Client) DoPayPalRequest {
	return func(ctx context.Context, uri string, urlGenerator func(string) string, _ paypay.Payload, patches []*entity.Patch) (res *http.Response, bs []byte, err error) {
		// data, _ := json.Marshal(patches)
		return c.doPayPal(ctx, xhttp.Patch(urlGenerator(uri)), patches, nil)
	}
}

func DeletePayPal(c *Client) DoPayPalRequest {
	return func(ctx context.Context, uri string, urlGenerator func(string) string, _ paypay.Payload, _ []*entity.Patch) (res *http.Response, bs []byte, err error) {
		return c.doPayPal(ctx, xhttp.Delete(urlGenerator(uri)), "", nil)
	}
}

func (c *Client) doPayPal(ctx context.Context, op xhttp.CfgOp, data any, headers map[string]string) (res *http.Response, bs []byte, err error) {
	res, bs, err = c.HClient.CallOp(ctx, data,
		xhttp.Req(xhttp.TypeJSON), // default json
		op,
		xhttp.Header(map[string]string{
			"Accept":                   "*/*",
			consts.HeaderAuthorization: consts.AuthorizationPrefixBearer + c.AccessToken,
		}),
		xhttp.Header(headers),
		xhttp.Prefix(PPReqPrefix(c.debug, c.Logger)),
		xhttp.Suffix(PPResSuffix(c.debug, c.Logger)),
	)
	if err != nil {
		return nil, nil, err
	}
	return res, bs, nil
}

// PPReqPrefix 闭包注入 logger 和 debug 信息
func PPReqPrefix(debug bool, log xlog.XLogger) xhttp.ReqPrefixFunc {
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
func PPResSuffix(debug bool, log xlog.XLogger) xhttp.ResSuffixFunc {
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
