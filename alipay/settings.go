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
 @Time    : 2024/8/26 -- 17:36
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: Settings.go
*/

package alipay

import (
	"encoding/base64"
	"fmt"
	"github.com/Bishoptylaor/paypay"
	"github.com/Bishoptylaor/paypay/alipay/cert"
	"github.com/Bishoptylaor/paypay/alipay/consts"
	"github.com/Bishoptylaor/paypay/pkg"
	"github.com/Bishoptylaor/paypay/pkg/xcrypto"
	"github.com/Bishoptylaor/paypay/pkg/xlog"
	"github.com/Bishoptylaor/paypay/pkg/xnet/xhttp"
	"time"
)

type Settings func(*Client)

func Checker(checker paypay.PayloadRuler) Settings {
	return func(client *Client) {
		client.EmptyChecker = checker
	}
}

func DefaultChecker() Settings {
	return Checker(func(uri string) []paypay.Ruler {
		return []paypay.Ruler{}
	})
}

func PayloadPreSetter(setter map[string][]paypay.PayloadPreSetter) Settings {
	return func(client *Client) {
		client.PayloadPreSetter = setter
	}
}

// Location 设置 时区，不设置或出错均为默认服务器时间
func Location(location *time.Location) Settings {
	return func(client *Client) {
		client.Location = location
	}
}

func DefaultLocation() Settings {
	return func(client *Client) {
		loc, err := time.LoadLocation("Asia/Shanghai")
		if err != nil {
			client.Logger.Error(err.Error())
			client.Logger.Infof("using UTC")
			return
		}
		Location(loc)(client)
	}
}

// DefaultCharset utf-8
func DefaultCharset() Settings {
	return Charset(consts.UTF8)
}

// Charset 设置编码格式，如utf-8,gbk,gb2312等，默认推荐使用 utf-8
func Charset(charset string) Settings {
	return func(client *Client) {
		if charset != pkg.NULL {
			client.Charset = charset
		}
	}
}

// DefaultVersion 1.0
func DefaultVersion() Settings {
	return func(client *Client) {
		client.Version = "1.0"
	}
}

// DefaultFormat JSON
func DefaultFormat() Settings {
	return func(client *Client) {
		client.Format = "JSON"
	}
}

// ReturnUrl 设置支付后的ReturnUrl
func ReturnUrl(url string) Settings {
	return func(client *Client) {
		client.ReturnUrl = url
	}
}

// NotifyUrl 设置支付宝服务器主动通知商户服务器里指定的页面http/https路径。
func NotifyUrl(url string) Settings {
	return func(client *Client) {
		client.NotifyUrl = url
	}
}

func Debug(d bool) Settings {
	return func(client *Client) {
		client.debug = d
	}
}

// DefaultSignType RSA2
func DefaultSignType() Settings {
	return SignType(consts.RSA2)
}

// SignType 设置签名算法类型，目前支持RSA2和RSA，默认推荐使用 RSA2
func SignType(signType string) Settings {
	return func(client *Client) {
		if signType != pkg.NULL {
			client.SignType = signType
		}
	}
}

func HClient(hc xhttp.HttpClientWrapper) Settings {
	return func(client *Client) {
		client.HClient = hc
	}
}

func DefaultHClient() Settings {
	return HClient(xhttp.GetDefaultClient())
}

func SetLogger(logger *xlog.Logger) Settings {
	return func(client *Client) {
		client.Logger = logger
	}
}

func DefaultLogger() Settings {
	return SetLogger(xlog.NewLogger())
}

// AppId set AppId
func AppId(appId string) Settings {
	return func(client *Client) {
		client.AppId = appId
	}
}

// PrivateKey set PrivateKey
func PrivateKey(pk string) Settings {
	return func(client *Client) {
		key := cert.FormatAlipayPrivateKey(pk)
		priKey, err := xcrypto.LoadPrivateKey(key)
		if err != nil {
			client.Logger.Error(err.Error())
			return
		}
		client.privateKey = priKey
	}
}

// PublicKey set PublicKey
func PublicKey(pk string) Settings {
	return func(client *Client) {
		key := cert.FormatAlipayPublicKey(pk)
		client.AutoVerifySign(key)
	}
}

func AutoVerify() Settings {
	return func(client *Client) {
		client.autoSign = true
	}
}

func SetEncryptKey(key string) Settings {
	return func(client *Client) {
		if key == "" {
			client.needEncrypt = false
			return
		}

		var data, err = base64.StdEncoding.DecodeString(key)
		if err != nil {
			client.needEncrypt = false
			return
		}

		client.needEncrypt = true
		client.encryptIV = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		client.encryptType = "AES"
		client.encryptKey = data
		client.encryptPadding = xcrypto.PKCS7
		return
	}
}

// AppAuthToken set AppAuthToken
func AppAuthToken(appAuthToken string) Settings {
	return func(client *Client) {
		client.AppAuthToken = appAuthToken
	}
}

func DefaultSign() Settings {
	return func(client *Client) {
		client.encoder = &Encoder{}
		client.signer = xcrypto.New(
			xcrypto.WithMethod(xcrypto.NewRSAMethod(GetHashBySignType(client), client.privateKey, nil)),
			xcrypto.WithEncoder(client.encoder),
		)
		client.verifiers = make(map[string]Verifier)
	}
}

func Prod(prod bool) Settings {
	return func(client *Client) {
		client.Prod = prod
	}
}

func CertSnContent(appCertContent, aliPayRootCertContent, aliPayPublicCertContent []byte) Settings {
	return func(client *Client) {
		err := client.LoadCertSnContent(appCertContent, aliPayRootCertContent, aliPayPublicCertContent)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func CertSnFile(appCertPath, aliPayRootCertPath, aliPayPublicCertPath string) Settings {
	return func(client *Client) {
		err := client.LoadCertSnFromFile(appCertPath, aliPayRootCertPath, aliPayPublicCertPath)
		if err != nil {
			fmt.Println(err)
		}
	}
}

// NewSettings 标准初始化配置
func NewSettings(ins ...Settings) []Settings {
	return append(
		append(
			[]Settings{},
			DefaultLogger(),   // 设置 logger
			DefaultHClient(),  // 设置 Http client
			DefaultSignType(), // 设置 RSA2
			DefaultFormat(),   // 设置 Json
			DefaultCharset(),  // 设置 UTF-8
			DefaultLocation(), // 设置 shanghai
			DefaultVersion(),  // 设置 1.0
			DefaultChecker(),  // 设置 checker 初始化
		), ins...,
	)
}

// DefaultSettings 默认沙盒配置，可自定义追加，按照顺序执行，自定义追加的部分会覆盖掉默认内容
func DefaultSettings(ins ...Settings) []Settings {
	return append(
		append(
			NewSettings(),
			AppId(consts.Appid),           // 设置 沙盒 appid
			PrivateKey(consts.PrivateKey), // 设置 沙盒 private key
			Prod(pkg.SandBox),             // 设置 沙盒环境
			Debug(pkg.DebugOn),            // debug on
			DefaultSign(),                 // 设置 签名机 需要先设置 private key, signType
		), ins...,
	)
}

func PackSettings(i1 []Settings, i2 ...Settings) []Settings {
	return append(i1, i2...)
}
