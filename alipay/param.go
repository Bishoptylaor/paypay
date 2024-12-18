package alipay

import (
	"context"
	"crypto"
	"encoding/json"
	"fmt"
	"github.com/Bishoptylaor/paypay"
	"github.com/Bishoptylaor/paypay/alipay/consts"
	"github.com/Bishoptylaor/paypay/alipay/entity"
	"github.com/Bishoptylaor/paypay/pkg"
	"github.com/Bishoptylaor/paypay/pkg/xutils"
	"time"
)

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

func InjectFixedPayload(c *Client, method string) paypay.ExecuteElem {
	return func(pl paypay.Payload) error {
		if setters, ok := c.PayloadPreSetter[method]; ok {
			for _, setter := range setters {
				setter(pl)
			}
		}
		return nil
	}
}

// AppId   string `json:"app_id"`   //支付宝分配给开发者的应用ID
// Method  string `json:"method"`   //接口名称
// Format  string `json:"format"`   //仅支持 JSON
// ReturnUrl  string `json:"return_url"`  //HTTP/HTTPS开头字符串
// Charset string `json:"charset"`  //请求使用的编码格式，如utf-8,gbk,gb2312等，推荐使用 utf-8
// SignType   string `json:"sign_type"`   //商户生成签名字符串所使用的签名算法类型，目前支持RSA2和RSA，推荐使用 RSA2
// Sign    string `json:"sign"`  //商户请求参数的签名串
// Timestamp  string `json:"timestamp"`   //发送请求的时间，格式"yyyy-MM-dd HH:mm:ss"
// Version string `json:"version"`  //调用的接口版本，固定为：1.0
// NotifyUrl  string `json:"notify_url"`  //支付宝服务器主动通知商户服务器里指定的页面http/https路径。
// BizContent string `json:"biz_content"` //业务请求参数的集合，最大长度不限，除公共参数外所有请求参数都必须放在这个参数中传递，具体参照各产品快速接入文档

func FixAppAuthToken(ctx context.Context, c *Client, method string, bizContent *string) paypay.ExecuteElem {
	return func(pl paypay.Payload) error {
		if pl != nil {
			var (
				bodyBs []byte
				err    error
			)
			_, has := entity.AppAuthTokenInBizContent[method]
			if !has {
				aat := pl.GetString("app_auth_token")
				pl.Remove("app_auth_token")
				if bodyBs, err = json.Marshal(pl); err != nil {
					return fmt.Errorf("json.Marshal：%w", err)
				}
				*bizContent = string(bodyBs)
				pl.Set("app_auth_token", aat)
			} else {
				if bodyBs, err = json.Marshal(pl); err != nil {
					return fmt.Errorf("json.Marshal：%w", err)
				}
				*bizContent = string(bodyBs)
				pl.Remove("app_auth_token")
			}
		}
		return nil
	}
}

func SetPublicParam(c *Client) paypay.ExecuteElem {
	return func(pl paypay.Payload) error {
		pl.
			Set("app_id", c.AppId).
			Set("format", c.Format).
			Set("charset", c.Charset).
			Set("sign_type", c.SignType).
			Set("version", c.Version).
			Set("timestamp", time.Now().In(c.Location).Format(pkg.TimeLayout))
		return nil
	}
}

func SetOptionalParam(c *Client) paypay.ExecuteElem {
	return func(pl paypay.Payload) error {
		opt := paypay.PayloadOptions{}
		opt.SetOptional(true)
		// add optional public params
		pl.
			Set("app_cert_sn", c.AppCertSN, opt).
			Set("alipay_root_cert_sn", c.AliPayRootCertSN, opt).
			Set("return_url", c.ReturnUrl, opt).
			Set("notify_url", c.NotifyUrl, opt).
			Set("app_auth_token", c.AppAuthToken, opt).
			Set("auth_token", c.AuthToken, opt)
		return nil
	}
}

// SetMethod set method
func SetMethod(method string) paypay.ExecuteElem {
	return func(pl paypay.Payload) error {
		pl.Set("method", method)
		return nil
	}
}

// SetMethod set method
func SetAppCertSn(acs string) paypay.ExecuteElem {
	return func(pl paypay.Payload) error {
		pl.Set("app_cert_sn", acs)
		return nil
	}
}

// SetByKey 如果有预设或者单次传入的字段值，支持使用最新传入的内容
func SetByKey(_k string, orig paypay.Payload) paypay.ExecuteElem {
	return func(pl paypay.Payload) error {
		if _v := pl.GetString(_k); _v != pkg.NULL {
			pl.Set(_k, _v)
		}

		return nil
	}
}

// SetBizContent set bizContent
func SetBizContent(biz paypay.Payload) paypay.ExecuteElem {
	return func(pl paypay.Payload) error {
		bytes, err := json.Marshal(biz)
		if err != nil {
			return fmt.Errorf("json.Marshal：%w", err)
		}
		pl["biz_content"] = string(bytes)
		return nil
	}
}

// SetBizContentString set bizContent
func SetBizContentString(ctx context.Context, c *Client, biz string) paypay.ExecuteElem {
	return func(pl paypay.Payload) error {
		if !c.needEncrypt {
			pl.Set("biz_content", biz)
		} else {
			// AES Encrypt biz_content
			encryptBizContent, err := c.encrypt(biz)
			if err != nil {
				return fmt.Errorf("EncryptBizContent Error: %w", err)
			}
			if c.debug == pkg.DebugOn {
				c.Logger.Debugf("Alipay_Origin_BizContent: %s", biz)
				c.Logger.Debugf("Alipay_Encrypt_BizContent: %s", encryptBizContent)
			}
			pl.Set("biz_content", encryptBizContent)
		}
		return nil
	}
}

func GetHashBySignType(c *Client) crypto.Hash {
	switch c.SignType {
	case consts.RSA:
		return crypto.SHA1
	case consts.RSA2:
		return crypto.SHA256
	default:
		return crypto.SHA256
	}
}
