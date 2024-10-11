package alipay

import (
	"context"
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"github.com/Bishoptylaor/paypay"
	"github.com/Bishoptylaor/paypay/alipay/cert"
	"github.com/Bishoptylaor/paypay/alipay/utils"
	"github.com/Bishoptylaor/paypay/pkg"
	"github.com/Bishoptylaor/paypay/pkg/xcrypto"
	"strings"
)

type Signer interface {
	SignBytes(ctx context.Context, data []byte, opts ...xcrypto.SignOption) ([]byte, error)
}

type Verifier interface {
	VerifyBytes(ctx context.Context, data []byte, signature []byte, opts ...xcrypto.SignOption) error
}

// LoadCertSnFromFile
// 通过应用公钥证书路径设置 appCertSN、aliPayPublicCertSN、aliPayRootCertSN
// appCertPath：应用公钥证书路径
// aliPayRootCertPath：支付宝根证书文件路径
// aliPayPublicCertPath：支付宝公钥证书文件路径
func (c *Client) LoadCertSnFromFile(appCertPath, aliPayRootCertPath, aliPayPublicCertPath string) (err error) {
	appCertSn, _, err := cert.GetCertSN(appCertPath)
	if err != nil {
		return fmt.Errorf("get app_cert_sn return err, but alse return alipay client. err: %w", err)
	}
	rootCertSn, err := cert.GetRootCertSN(aliPayRootCertPath)
	if err != nil {
		return fmt.Errorf("get alipay_root_cert_sn return err, but alse return alipay client. err: %w", err)
	}
	publicCertSn, publicPub, err := cert.GetCertSN(aliPayPublicCertPath)
	if err != nil {
		return fmt.Errorf("get alipay_cert_sn return err, but alse return alipay client. err: %w", err)
	}
	c.AppCertSN = appCertSn
	c.AliPayRootCertSN = rootCertSn
	c.AliPayPublicCertSN = publicCertSn
	var verifier = xcrypto.New(xcrypto.WithMethod(xcrypto.NewRSAMethod(crypto.SHA256, nil, publicPub)), xcrypto.WithEncoder(c.encoder))
	c.verifiers[c.AliPayPublicCertSN] = verifier
	return nil
}

// LoadCertSnContent
// 通过应用公钥证书内容设置 appCertSN、aliPayPublicCertSN、aliPayRootCertSN
// appCertContent：应用公钥证书文件内容
// aliPayRootCertContent：支付宝根证书文件内容
// aliPayPublicCertContent：支付宝公钥证书文件内容
func (c *Client) LoadCertSnContent(appCertContent, aliPayRootCertContent, aliPayPublicCertContent []byte) (err error) {
	appCertSn, _, err := cert.GetCertSN(appCertContent)
	if err != nil {
		return fmt.Errorf("get app_cert_sn return err, but alse return alipay client. err: %w", err)
	}
	rootCertSn, err := cert.GetRootCertSN(aliPayRootCertContent)
	if err != nil {
		return fmt.Errorf("get alipay_root_cert_sn return err, but alse return alipay client. err: %w", err)
	}
	publicCertSn, publicPub, err := cert.GetCertSN(aliPayPublicCertContent)
	if err != nil {
		return fmt.Errorf("get alipay_cert_sn return err, but alse return alipay client. err: %w", err)
	}
	c.AppCertSN = appCertSn
	c.AliPayRootCertSN = rootCertSn
	c.AliPayPublicCertSN = publicCertSn
	var verifier = xcrypto.New(xcrypto.WithMethod(xcrypto.NewRSAMethod(crypto.SHA256, nil, publicPub)), xcrypto.WithEncoder(c.encoder))
	c.verifiers[c.AliPayPublicCertSN] = verifier
	return nil
}

// LoadPublicSnContent
// 通过应用公钥证书内容设置 appCertSN、aliPayPublicCertSN、aliPayRootCertSN
// appCertContent：应用公钥证书文件内容
// aliPayRootCertContent：支付宝根证书文件内容
// aliPayPublicCertContent：支付宝公钥证书文件内容
func (c *Client) LoadPublicSnContent(appCertContent, aliPayRootCertContent, aliPayPublicCertContent []byte) (err error) {
	appCertSn, _, err := cert.GetCertSN(appCertContent)
	if err != nil {
		return fmt.Errorf("get app_cert_sn return err, but alse return alipay client. err: %w", err)
	}
	rootCertSn, err := cert.GetRootCertSN(aliPayRootCertContent)
	if err != nil {
		return fmt.Errorf("get alipay_root_cert_sn return err, but alse return alipay client. err: %w", err)
	}
	publicCertSn, publicPub, err := cert.GetCertSN(aliPayPublicCertContent)
	if err != nil {
		return fmt.Errorf("get alipay_cert_sn return err, but alse return alipay client. err: %w", err)
	}
	c.AppCertSN = appCertSn
	c.AliPayRootCertSN = rootCertSn
	c.AliPayPublicCertSN = publicCertSn
	var verifier = xcrypto.New(xcrypto.WithMethod(xcrypto.NewRSAMethod(crypto.SHA256, nil, publicPub)), xcrypto.WithEncoder(c.encoder))
	c.verifiers[c.AliPayPublicCertSN] = verifier
	return nil
}

// =============================== 获取SignData ===============================

// autoVerifySignByCert 同步验签
// 需注意的是，公钥签名模式和公钥证书签名模式的不同之处
// 验签文档：https://opendocs.alipay.com/open/200/106120
func (c *Client) autoVerifySignByCert(ctx context.Context, bs []byte, method string, alipayCertSN string) (signData string, err error) {
	var raw = make(map[string]json.RawMessage)
	if err = json.Unmarshal(bs, &raw); err != nil {
		return
	}

	var bizFieldName = strings.Replace(method, ".", "_", -1) + "_response"
	var signBytes = raw["sign"]
	var certBytes = raw["alipay_cert_sn"]
	var bizBytes = raw[bizFieldName]
	var errBytes = raw["error_response"]

	if len(certBytes) > 1 {
		certBytes = certBytes[1 : len(certBytes)-1]
	}
	if len(signBytes) > 1 {
		signBytes = signBytes[1 : len(signBytes)-1]
	}

	if len(bizBytes) == 0 {
		if len(errBytes) > 0 {
			var rErr *utils.BussError
			if err = json.Unmarshal(errBytes, &rErr); err != nil {
				return pkg.NULL, err
			}
			return pkg.NULL, rErr
		}
		return pkg.NULL, pkg.ErrBadResponse
	}

	if alipayCertSN != "" {
		// 公钥证书模式
		if alipayCertSN != c.AliPayPublicCertSN {
			return pkg.NULL, fmt.Errorf("[%w]: 当前使用的支付宝公钥证书SN[%s]与网关响应报文中的SN[%s]不匹配", pkg.ErrCertNotMatch, c.AliPayPublicCertSN, alipayCertSN)
		}
	} else {
		certBytes = []byte("alipayPublicKeyContent")
	}

	// 解密
	var origData []byte
	if origData, err = c.decrypt(bizBytes); err != nil {
		return pkg.NULL, err
	}

	// 验签 仅在证书模式下生效
	if c.autoSign && c.aliPayPublicKey != nil {
		if c.debug == pkg.DebugOn {
			c.Logger.Debugf("[Alipay_SyncSignData]: %s, Sign=[%s]", signData, signBytes)
		}
		c.Logger.Warnf("[Alipay_SyncSignData]: certBytes=[%s], bizBytes=[%s], Sign=[%s]", certBytes, bizBytes, signBytes)

		if len(signBytes) == 0 {
			// 没有签名数据，返回的内容一般为错误信息
			var rErr *utils.BussError
			if err = json.Unmarshal(origData, &rErr); err != nil {
				return pkg.NULL, err
			}
			return pkg.NULL, rErr
		}

		// 验证签名
		if err = c.verify(ctx, string(certBytes), bizBytes, signBytes); err != nil {
			return pkg.NULL, err
		}
	}
	return string(signBytes), nil
}

func (c *Client) decrypt(data []byte) ([]byte, error) {
	var plaintext = data
	if len(data) > 1 && data[0] == '"' {
		var ciphertext, err = xcrypto.Base64.Decode(string(data[1 : len(data)-1]))
		if err != nil {
			return nil, err
		}
		plaintext, err = xcrypto.AESCBCDecrypt(ciphertext, c.encryptKey, c.encryptIV, c.encryptPadding)
		if err != nil {
			return nil, err
		}
	}
	return plaintext, nil
}

func (c *Client) encrypt(originData string) (string, error) {
	encryptData, err := xcrypto.AESCBCEncryptWithBase64([]byte(originData), c.encryptKey, c.encryptIV, xcrypto.PKCS7)
	if err != nil {
		return "", err
	}
	return encryptData, nil
}

func (c *Client) sign(ctx context.Context, originData string) (signature string, err error) {
	sBytes, err := c.signer.SignBytes(ctx, []byte(originData))
	if err != nil {
		return "", err
	}
	signature = xcrypto.Base64.Encode(sBytes)
	return signature, nil
}

func (c *Client) verify(ctx context.Context, certSN string, data, signature []byte) (err error) {
	var verifier Verifier
	if verifier, err = c.getVerifier(certSN); err != nil {
		return err
	}

	if signature, err = xcrypto.Base64.Decode(string(signature)); err != nil {
		return err
	}

	if err = verifier.VerifyBytes(ctx, data, signature); err != nil {
		return err
	}
	return nil
}

func (c *Client) getVerifier(certSN string) (verifier Verifier, err error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if certSN == "" {
		certSN = c.AliPayRootCertSN
	}

	verifier = c.verifiers[certSN]

	if verifier == nil {
		if !c.Prod {
			return nil, pkg.ErrAliPublicKeyNotFound
		}

		certs, err := c.downloadAliPayCert(certSN)
		if err != nil {
			return nil, err
		}

		pub, ok := certs.PublicKey.(*rsa.PublicKey)
		if !ok {
			return nil, pkg.ErrAliPublicKeyNotFound
		}
		verifier = xcrypto.New(xcrypto.WithMethod(xcrypto.NewRSAMethod(crypto.SHA256, nil, pub)), xcrypto.WithEncoder(c.encoder))
	}
	return verifier, nil
}

func (c *Client) downloadAliPayCert(certSN string) (cert *x509.Certificate, err error) {
	res, err := c.PublicCertDownload(context.Background(), make(paypay.Payload).Set("alipay_cert_sn", certSN))
	if err != nil {
		return nil, err
	}
	certBytes, err := base64.StdEncoding.DecodeString(res.Response.AlipayCertContent)
	if err != nil {
		return nil, err
	}

	if block, _ := pem.Decode(certBytes); block != nil {
		cert, err = x509.ParseCertificate(block.Bytes)
		if err != nil {
			return nil, err
		}
	}

	return cert, nil
}
