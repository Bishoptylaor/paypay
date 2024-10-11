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
 @Time    : 2024/8/16 -- 11:18
 @Author  : 亓官竹
 @Copyright 2024 亓官竹
 @Description: client.go
*/

package alipay

import (
	"context"
	"crypto"
	"crypto/rsa"
	"fmt"
	"github.com/Bishoptylaor/paypay"
	"github.com/Bishoptylaor/paypay/alipay/config"
	"github.com/Bishoptylaor/paypay/operate"
	"github.com/Bishoptylaor/paypay/pkg"
	"github.com/Bishoptylaor/paypay/pkg/xcrypto"
	"sync"
)

type Client struct {
	// 配置组
	config.Config

	// 操作组
	operate.Operates
	debug bool
	mu    sync.Mutex

	// 加密组
	privateKey      *rsa.PrivateKey // 本地私钥，用于发起请求时对全部参数进行签名，支持PKCS1和PKCS8
	aliPayPublicKey *rsa.PublicKey  // alipay公钥证书，由本地公钥在商家后台换取
	autoSign        bool
	needEncrypt     bool
	encryptIV       []byte
	encryptType     string
	encryptKey      []byte
	encryptPadding  xcrypto.Pad

	// 签名组
	encoder   xcrypto.Encoder
	signer    Signer
	verifiers map[string]Verifier
}

// NewClient 初始化支付宝客户端
//
// 通过 settings 类型自定义参数
func NewClient(ctx context.Context, ops ...Settings) (client *Client, err error) {
	client = &Client{}
	// 预设关键参数
	// 默认关闭 debug
	client.Use(Debug(pkg.DebugOff))
	client.Use(NewSettings()...)
	// NewSettings 为下方预设参数集合
	// client.Use(DefaultLogger())
	// client.Use(DefaultHClient())
	// client.Use(DefaultSign())
	// client.Use(DefaultSignType())
	// client.Use(DefaultFormat())
	// client.Use(DefaultCharset())

	if len(ops) == 0 {
		client.Use(DefaultSettings()...)
	}
	client.Use(ops...)

	if err = client.setupCheck(); err != nil {
		return nil, err
	}

	return client, nil
}

func (c *Client) Use(ops ...Settings) {
	for _, op := range ops {
		op(c)
	}
}

func (c *Client) setupCheck() error {
	if c.AppId == pkg.NULL || c.privateKey == nil {
		return pkg.ErrMissingInitParams
	}
	return nil
}

func (c *Client) baseCall(ctx context.Context, pl paypay.Payload, caller string) (bs []byte, err error) {
	if bs, err = c.callAli(ctx, pl, caller); err != nil {
		return nil, err
	}
	return bs, nil
}

// AutoVerifySign 开启请求完自动验签功能（默认不开启，推荐开启，只支持证书模式）
// 注意：只支持证书模式
// alipayPublicKeyContent：支付宝公钥证书文件内容[]byte
func (c *Client) AutoVerifySign(alipayPublicKeyContent string) {
	pubKey, err := xcrypto.LoadPublicKey(alipayPublicKeyContent)
	if err != nil {
		c.Logger.Errorf("AutoVerifySign(%s),err:%+v", alipayPublicKeyContent, err)
	}
	if pubKey != nil {
		c.aliPayPublicKey = pubKey
		c.autoSign = true
	}
	var verifier = xcrypto.New(
		xcrypto.WithMethod(xcrypto.NewRSAMethod(crypto.SHA256, nil, pubKey)), xcrypto.WithEncoder(c.encoder))
	if c.verifiers == nil {
		c.verifiers = make(map[string]Verifier)
	}
	c.verifiers["alipayPublicKeyContent"] = verifier
}

// RequestParam 获取支付宝完整请求参数包含签名
// 注意：biz_content 需要自行通过pl.SetPayload()设置，不设置则没有此参数
func (c *Client) RequestParam(ctx context.Context, pl paypay.Payload, method string) (string, error) {
	var (
		err        error
		sign       string
		bizContent paypay.Payload
		ok         bool
	)
	// check if there is biz_content
	bz := pl.GetAny("biz_content")
	if bizContent, ok = bz.(paypay.Payload); ok {
	}

	err = paypay.ExecuteQueue(
		SetMethod(method),
		SetPublicParam(c),
		SetOptionalParam(c),
		SetBizContent(bizContent),
	)(pl)
	if err != nil {
		return "", err
	}

	// 公共参数校验
	err = paypay.ExecuteQueue(
		SetByKey("app_id", pl),
		SetByKey("app_cert_sn", pl),
		SetByKey("notify_url", pl),
		SetByKey("alipay_root_cert_sn", pl),
		SetByKey("return_url", pl),
		SetByKey("notify_url", pl),
		SetByKey("app_auth_token", pl),
	)(pl)
	if err != nil {
		return "", err
	}

	// check sign
	if pl.GetString("sign") == "" {
		sign, err = c.sign(ctx, pl.EncodeAliPaySignParams())
		if err != nil {
			return "", fmt.Errorf("GetRsaSign Error: %w", err)
		}
		pl.Set("sign", sign)
	}

	if c.debug {
		c.Logger.Debugf("Alipay_Request: %s", pl.JsonBody())
	}
	return pl.EncodeURLParams(), nil
}

// 公共参数处理
func (c *Client) publicParamsHandler(ctx context.Context, pl paypay.Payload, method, bizContent string, authToken ...string) (pubBody paypay.Payload, param string, err error) {
	if len(authToken) > 0 {
		c.AuthToken = authToken[0]
	}
	pubBody = make(paypay.Payload)
	err = paypay.ExecuteQueue(
		SetMethod(method),
		SetPublicParam(c),
		SetOptionalParam(c),
		SetBizContentString(ctx, c, bizContent),
		SetByKey("version", pl),
		SetByKey("return_url", pl),
		SetByKey("notify_url", pl),
		SetByKey("app_auth_token", pl),
	)(pubBody)
	if err != nil {
		return nil, "", err
	}

	// sign
	sign, err := c.sign(ctx, pubBody.EncodeAliPaySignParams())
	if err != nil {
		return nil, "", fmt.Errorf("GetRsaSign Error: %w", err)
	}
	pubBody.Set("sign", sign)
	if c.debug == pkg.DebugOn {
		c.Logger.Debugf("Alipay_Request: %s", pubBody.JsonBody())
	}
	param = pubBody.EncodeURLParams()
	return
}
