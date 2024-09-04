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
 @Description: auth.go
*/

package paypal

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/Bishoptylaor/paypay"
	"github.com/Bishoptylaor/paypay/paypal/consts"
	"github.com/Bishoptylaor/paypay/paypal/entity"
	"github.com/Bishoptylaor/paypay/pkg/xnet/xhttp"
	"github.com/Bishoptylaor/paypay/pkg/xutils"
	"net/http"
	"runtime"
	"time"
)

// GetAccessToken
// 获取AccessToken（Get an access token）
// 文档：https://developer.paypal.com/docs/api/reference/get-an-access-token
func (c *Client) GetAccessToken(ctx context.Context) (token *entity.AccessToken, err error) {
	// Authorization
	authHeader := consts.AuthorizationPrefixBasic + base64.StdEncoding.EncodeToString([]byte(c.ClientID+":"+c.ClientSecret))

	// Body
	pl := make(paypay.Payload)
	pl.Set("grant_type", "client_credentials")

	res, bs, err := c.HClient.CallOp(ctx, pl,
		xhttp.Req(xhttp.TypeFormData),
		xhttp.Post(c.GenUrl(consts.AccessTokenPath)),
		xhttp.Header(map[string]string{
			"Accept":                   "*/*",
			consts.HeaderAuthorization: authHeader,
		}),
		xhttp.Prefix(PPReqPrefix(c.debug, c.Logger)),
		xhttp.Suffix(PPResSuffix(c.debug, c.Logger)),
	)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}

	token = new(entity.AccessToken)
	if err = json.Unmarshal(bs, token); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	}
	c.Appid = token.Appid
	c.AccessToken = token.AccessToken
	c.ExpireIn = time.Duration(token.ExpiresIn)
	return token, nil
}

// autoRefreshToken 自动刷新 token
func (c *Client) autoRefreshToken(ctx context.Context) {
	defer func() {
		if r := recover(); r != nil {
			buf := make([]byte, 64<<10)
			buf = buf[:runtime.Stack(buf, false)]
			c.Logger.Errorf("paypal_goAuthRefreshToken: panic recovered: %s\n%s", r, buf)
		}
	}()
	for {
		time.Sleep(c.ExpireIn / 2 * time.Second)
		err := xutils.Retry(func() error {
			_, err := c.GetAccessToken(ctx)
			if err != nil {
				return err
			}
			return nil
		}, 3, time.Second)
		if err != nil {
			c.Logger.Errorf("PayPal GetAccessToken Error: %s", err.Error())
		}
	}
}
