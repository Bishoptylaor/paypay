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
 @Time    : 2024/8/26 -- 15:05
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: operate.go
*/

package operate

import (
	"bytes"
	"context"
	"github.com/Bishoptylaor/paypay"
	"github.com/Bishoptylaor/paypay/pkg/zlog"
	"github.com/Bishoptylaor/paypay/pkg/znet/zhttp"
	"io"
	"net/http"
	"net/url"
)

type Operator func(*Operates)

type Operates struct {
	Prefix []ReqPrefixFunc
	Suffix []ResSuffixFunc

	// 其他工具组
	EmptyChecker     paypay.PayloadRuler
	PayloadPreSetter map[string][]paypay.PayloadPreSetter
	HClient          *zhttp.HttpClientWrapper
	Logger           *zlog.Logger
}

// ReqPrefixFunc a hook Before the request started
type ReqPrefixFunc func(context context.Context, req *http.Request, log *zlog.Logger) context.Context

// ResSuffixFunc a hook After the request finished
type ResSuffixFunc func(context context.Context, res *http.Response, log *zlog.Logger) context.Context

// ConvertFunc convert res body into a struct or map todo
type ConvertFunc func(context.Context, interface{}, interface{}) error

func DefaultReqPrefix(ctx context.Context, req *http.Request, log *zlog.Logger) context.Context {
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
	return ctx
}

func DefaultResSuffix(ctx context.Context, res *http.Response, log *zlog.Logger) context.Context {
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Errorf("[Read Res body] error: %s", err.Error())
		return ctx
	}
	res.Body = io.NopCloser(bytes.NewBuffer(body))
	log.Infof("[Res] %s", string(body))
	return ctx
}
