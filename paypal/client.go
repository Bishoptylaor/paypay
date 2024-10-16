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
 @Time    : 2024/9/3 -- 16:45
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: client.go
*/

package paypal

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Bishoptylaor/paypay"
	"github.com/Bishoptylaor/paypay/operate"
	"github.com/Bishoptylaor/paypay/paypal/config"
	"github.com/Bishoptylaor/paypay/paypal/entity"
	"github.com/Bishoptylaor/paypay/pkg"
	"github.com/Bishoptylaor/paypay/pkg/xutils"
	"github.com/pkg/errors"
	"net/http"
)

type Client struct {
	// 配置组
	config.Config

	// 操作组
	operate.Operates

	// to remember every setting from init. used when we need a temporary new client with tiny different settings
	ops   []Settings
	debug bool
}

// NewClient 初始化 Paypal 客户端
func NewClient(ctx context.Context, ops ...Settings) (client *Client, err error) {
	client = &Client{
		ops: make([]Settings, 0),
	}
	// 预设关键参数
	// 默认关闭 debug
	client.Use(Debug(pkg.DebugOff))
	client.Use(NewSettings()...)
	// NewSettings 为下方预设参数集合
	// client.Use(DefaultLogger())
	// client.Use(DefaultHClient())

	if len(ops) == 0 {
		client.Use(DefaultSettings()...)
	}
	client.Use(ops...)

	if err = client.setupCheck(); err != nil {
		return nil, err
	}

	_, err = client.getAccessToken(ctx)
	if err != nil {
		return nil, err
	}
	go client.autoRefreshToken(ctx)

	return client, nil
}

func (c *Client) Use(ops ...Settings) *Client {
	for _, op := range ops {
		op(c)
		c.ops = append(c.ops, op)
	}
	return c
}

func (c *Client) setupCheck() error {
	if c.ClientID == pkg.NULL || c.ClientSecret == pkg.NULL {
		return pkg.ErrMissingInitParams
	}
	return nil
}

// handleResponse 处理 HTTP 响应，并填充结果结构体
func (c *Client) handleResponse(ctx context.Context, method Method, httpRes *http.Response, bs []byte, emptyRes *entity.EmptyRes, response interface{}) error {
	if httpRes.StatusCode != method.ValidStatusCode {
		emptyRes.Code = httpRes.StatusCode
		emptyRes.Error = string(bs)
		emptyRes.ErrorResponse = new(entity.ErrorResponse)
		if err := json.Unmarshal(bs, emptyRes.ErrorResponse); err != nil {
			return errors.Wrap(err, "[handleResponse] wrong status code unmarshal err")
		}
		return nil
	}
	if response == nil || method.ValidStatusCode == http.StatusNoContent {
		return nil
	}

	if err := json.Unmarshal(bs, response); err != nil {
		return errors.Wrap(err, "[handleResponse] response unmarshal err")
	}
	return nil
}

func (c *Client) Print() {

}

func IntegrityCheck(ctx context.Context, c *Client, method string) paypay.ExecuteElem {
	return func(pl paypay.Payload) error {
		if c.EmptyChecker == nil {
			return nil
		}
		var ok bool
		var err error
		for _, ruler := range c.EmptyChecker(method) {
			ok, err = xutils.Expr(ctx, ruler.Rule, pl)
			if !ok || err != nil {
				return fmt.Errorf("[IntegrityCheck]: rule:[%s], err[%s], [%s]", ruler.Des, err, ruler.Alert)
			}
		}
		return nil
	}
}
