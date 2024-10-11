package alipay

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Bishoptylaor/paypay"
	"github.com/Bishoptylaor/paypay/pkg"
	"github.com/Bishoptylaor/paypay/pkg/xlog"
	"github.com/Bishoptylaor/paypay/pkg/xnet/xhttp"
	"io"
	"net/http"
	"net/url"
)

// PostAliPayAPISelfV2 支付宝接口自行实现方法
// 注意：biz_content 需要自行通过pl.SetPayload()设置，不设置则没有此参数
// 示例：请参考 client_test.go 的 TestClient_PostAliPayAPISelfV2() 方法
func (c *Client) PostAliPayAPISelfV2(ctx context.Context, Payload paypay.Payload, method string, aliRes any) (err error) {
	var (
		bs, bodyBs []byte
	)
	// check if there is biz_content
	bz := Payload.GetAny("biz_content")
	if bzBody, ok := bz.(paypay.Payload); ok {
		if bodyBs, err = json.Marshal(bzBody); err != nil {
			return fmt.Errorf("json.Marshal(%v)：%w", bzBody, err)
		}
		Payload.Set("biz_content", string(bodyBs))
	}

	if bs, err = c.callAliDirect(ctx, Payload, method); err != nil {
		return err
	}
	if err = json.Unmarshal(bs, aliRes); err != nil {
		return err
	}
	return nil
}

// PostFileAliPayAPISelfV2 用于支付宝带有文件上传的接口自行实现方法
// 注意：最新版本的支付宝接口，对于文件的上传已统一改为通过formData上传
// 请求form格式如下： {file: "fileData", "data": Payload{"key": "value"}}
// 其中file为file请求字段名称，data为其他请求参数（key为文件名，value为文件内容）
func (c *Client) PostFileAliPayAPISelfV2(ctx context.Context, pl paypay.Payload, method string, aliRes any) (err error) {
	var (
		sign string
	)
	err = paypay.ExecuteQueue(
		SetMethod(method),
		SetPublicParam(c),
		SetOptionalParam(c),
	)(pl)
	if err != nil {
		return err
	}

	// check sign, 需要先移除文件字段
	fm := make(paypay.Payload)
	for k, v := range pl {
		if _, ok := v.(*paypay.File); ok {
			fm.Set(k, v)
			pl.Remove(k)
			continue
		}
	}
	if pl.GetString("sign") == "" {
		sign, err = c.sign(ctx, pl.EncodeAliPaySignParams())
		if err != nil {
			return fmt.Errorf("GetRsaSign Error: %w", err)
		}
		pl.Set("sign", sign)
	}
	// 增加文件字段
	for k, v := range fm {
		pl.Set(k, v)
	}

	res, err := c.HClient.CallOpWrite(ctx, pl, aliRes,
		xhttp.Req(xhttp.TypeMultipartFormData),
		xhttp.Post(c.Utf8Url()),
		xhttp.Prefix(AliReqPrefix(c.debug, c.Logger)),
		xhttp.Suffix(AliResSuffix(c.debug, c.Logger)),
	)
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		return fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}
	return nil
}

// 向支付宝发送自定义请求
func (c *Client) callAliDirect(ctx context.Context, pl paypay.Payload, method string) (bs []byte, err error) {
	var sign string
	err = paypay.ExecuteQueue(
		SetMethod(method),
		SetPublicParam(c),
		SetOptionalParam(c),
	)(pl)
	if err != nil {
		return bs, err
	}
	// check sign
	if pl.GetString("sign") == "" {
		sign, err = c.sign(ctx, pl.EncodeAliPaySignParams())
		if err != nil {
			return nil, fmt.Errorf("GetRsaSign Error: %w", err)
		}
		pl.Set("sign", sign)
	}
	if c.debug == pkg.DebugOn {
		c.Logger.Debugf("Alipay_Request: %s", pl.JsonBody())
	}

	res, bs, err := c.HClient.CallOp(ctx, pl,
		xhttp.Req(xhttp.TypeFormData),
		xhttp.Post(c.Utf8Url()),
		xhttp.Prefix(AliReqPrefix(c.debug, c.Logger)),
		xhttp.Suffix(AliResSuffix(c.debug, c.Logger)),
	)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}
	return bs, nil
}

// 向支付宝发送请求
func (c *Client) callAli(ctx context.Context, pl paypay.Payload, method string, authToken ...string) (bs []byte, err error) {
	var bizContent string
	err = paypay.ExecuteQueue(
		// 注入该场景下固定参数
		InjectFixedPayload(c, method),
		// 校验 biz_content 参数规则
		IntegrityCheck(ctx, c, method),
		FixAppAuthToken(ctx, c, method, &bizContent),
	)(pl)
	if err != nil {
		return []byte{}, err
	}

	// 处理公共参数
	data, param, err := c.publicParamsHandler(ctx, pl, method, bizContent, authToken...)
	if err != nil {
		return nil, err
	}
	switch method {
	case "alipay.trade.app.pay", "alipay.fund.auth.order.app.freeze", "zhima.credit.pe.zmgo.sign.apply", "zhima.credit.payafteruse.creditagreement.sign":
		return []byte(param), nil
	case "alipay.trade.wap.pay", "alipay.trade.page.pay", "alipay.user.certify.open.certify":
		return []byte(c.Url() + "?" + param), nil
	default:
		res, bs, err := c.HClient.CallOp(ctx, data,
			xhttp.Req(xhttp.TypeFormData),
			xhttp.Post(c.Utf8Url()),
			xhttp.Prefix(AliReqPrefix(c.debug, c.Logger)),
			xhttp.Suffix(AliResSuffix(c.debug, c.Logger)),
		)
		if err != nil {
			return nil, err
		}
		if res.StatusCode != 200 {
			return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
		}
		return bs, nil
	}
}

// CallAli 外部调用：向支付宝发送请求
func (c *Client) CallAli(ctx context.Context, pl paypay.Payload, method string, authToken ...string) (bs []byte, err error) {
	return c.callAli(ctx, pl, method, authToken...)
}

func (c *Client) PageExecute(ctx context.Context, pl paypay.Payload, method string, authToken ...string) (url string, err error) {
	var (
		bizContent string
	)
	err = paypay.ExecuteQueue(
		// 注入该场景下固定参数
		InjectFixedPayload(c, method),
		// 校验 biz_content 参数规则
		IntegrityCheck(ctx, c, method),
		FixAppAuthToken(ctx, c, method, &bizContent),
	)(pl)
	if err != nil {
		return pkg.NULL, err
	}

	// 处理公共参数
	_, param, err := c.publicParamsHandler(ctx, pl, method, bizContent, authToken...)
	if err != nil {
		return "", err
	}

	return c.Url() + "?" + param, nil
}

// FileUploadRequest 文件上传
func (c *Client) FileUploadRequest(ctx context.Context, pl paypay.Payload, method string) (bs []byte, err error) {
	var (
		aat string
	)
	if pl != nil {
		aat = pl.GetString("app_auth_token")
		pl.Remove("app_auth_token")
	}
	pubBody := make(paypay.Payload)
	err = paypay.ExecuteQueue(
		SetPublicParam(c),
		SetMethod(method),
	)(pubBody)
	if err != nil {
		return nil, err
	}

	// if user set app_auth_token in body_map, use this
	if aat != pkg.NULL {
		pubBody.Set("app_auth_token", aat)
	}
	// 文件上传除文件外其他参数也需要签名
	for k, v := range pl {
		if _, ok := v.(*paypay.File); !ok {
			pubBody.Set(k, v)
		}
	}
	// todo check sign
	sign, err := c.sign(ctx, pubBody.EncodeAliPaySignParams())
	if err != nil {
		return nil, fmt.Errorf("GetRsaSign Error: %w", err)
	}
	// 文件签名完移除query params
	for k := range pl {
		pubBody.Remove(k)
	}
	pubBody.Set("sign", sign)

	res, bs, err := c.HClient.CallOp(ctx, pubBody,
		xhttp.Req(xhttp.TypeMultipartFormData),
		xhttp.Post(c.Utf8Url()+"&"+pubBody.EncodeURLParams()),
		xhttp.Prefix(AliReqPrefix(c.debug, c.Logger)),
		xhttp.Suffix(AliResSuffix(c.debug, c.Logger)),
	)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}
	return bs, nil
}

// AliReqPrefix 闭包注入 logger 和 debug 信息
func AliReqPrefix(debug bool, log xlog.XLogger) xhttp.ReqPrefixFunc {
	return func(ctx context.Context, req *http.Request) context.Context {
		if debug == pkg.DebugOn {
			log.Debugf("Alipay_Request: %s", req.Body)
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

// AliResSuffix 闭包注入 logger 和 debug 信息
func AliResSuffix(debug bool, log xlog.XLogger) xhttp.ResSuffixFunc {
	return func(ctx context.Context, res *http.Response) context.Context {
		if debug == pkg.DebugOn {
			log.Debugf("Alipay_Response: %d, %s", res.StatusCode, res.Body)
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
