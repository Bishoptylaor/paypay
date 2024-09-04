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
	"fmt"
	"github.com/Bishoptylaor/paypay"
	"github.com/Bishoptylaor/paypay/operate"
	"github.com/Bishoptylaor/paypay/paypal/config"
	"github.com/Bishoptylaor/paypay/pkg"
	"github.com/Bishoptylaor/paypay/pkg/xutils"
)

type Client struct {
	// 配置组
	config.Config

	// 操作组
	operate.Operates
	debug bool
}

// NewClient 初始化 Paypal 客户端
func NewClient(ctx context.Context, ops ...Settings) (client *Client, err error) {
	client = &Client{}
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

	go client.autoRefreshToken(ctx)

	return client, nil
}

func (c *Client) Use(ops ...Settings) {
	for _, op := range ops {
		op(c)
	}
}

func (c *Client) setupCheck() error {
	if c.ClientID == pkg.NULL || c.ClientSecret == pkg.NULL {
		return pkg.ErrMissingInitParams
	}
	return nil
}

func IntegrityCheck(ctx context.Context, c *Client, method string) paypay.ExecuteElem {
	return func(pl paypay.Payload) error {
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
