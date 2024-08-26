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
 @Time    : 2024/8/26 -- 17:09
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: cert.go
*/

package cert

import (
	"crypto/md5"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"github.com/Bishoptylaor/paypay/pkg"
	"os"
	"strings"
)

// 允许进行 sn 提取的证书签名算法
var allowSignatureAlgorithm = map[string]bool{
	"MD2-RSA":       true,
	"MD5-RSA":       true,
	"SHA1-RSA":      true,
	"SHA256-RSA":    true,
	"SHA384-RSA":    true,
	"SHA512-RSA":    true,
	"SHA256-RSAPSS": true,
	"SHA384-RSAPSS": true,
	"SHA512-RSAPSS": true,
}

/*
Q：使用公钥证书签名方式下，为什么开放平台网关的响应报文需要携带支付宝公钥证书SN（alipay_cert_sn）？
**
A：开发者上传自己的应用公钥证书后，开放平台会为开发者应用自动签发支付宝公钥证书供开发者下载，用来对开放平台网关响应报文做验签。

但是支付宝公钥证书可能因证书到期或者变更CA签发机构等原因，可能会重新签发证书。在重新签发前，开放平台会在门户上提前提醒开发者支付宝应用公钥证书变更时间。

但为避免开发者因未能及时感知支付宝公钥证书变更而导致验签失败，开放平台提供了一种支付宝公钥证书无感知升级机制，具体流程如下：
1）开放平台网关在响应报文中会多返回支付宝公钥证书SN
2）开放平台网关提供根据SN下载对应支付宝公钥证书的API接口
3）开发者在验签过程中，先比较本地使用的支付宝公钥证书SN与开放平台网关响应中SN是否一致。若不一致，可调用支付宝公钥证书下载接口下载对应SN的支付宝公钥证书。
4）对下载的支付宝公钥证书执行证书链校验，若校验通过，则用该证书验签。

基于该机制可实现支付宝公钥证书变更时开发者无感知，当前开放平台提供的SDK已基于该机制实现对应功能。若开发者未通过SDK接入，须自行实现该功能。
*/

// GetCertSN 获取证书序列号SN
// certPathOrData x509证书文件路径(appPublicCert.crt、alipayPublicCert.crt) 或证书 buffer
// 返回 sn：证书序列号(app_cert_sn、alipay_cert_sn)
// 返回 err：error 信息
func GetCertSN(certPathOrData any) (sn string, pub *rsa.PublicKey, err error) {
	var certData []byte
	var ok bool
	switch pathOrData := certPathOrData.(type) {
	case string:
		certData, err = os.ReadFile(pathOrData)
		if err != nil {
			return pkg.NULL, nil, err
		}
	case []byte:
		certData = pathOrData
	default:
		return pkg.NULL, nil, errors.New("certPathOrData 证书类型断言错误")
	}

	if block, _ := pem.Decode(certData); block != nil {
		cert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			return pkg.NULL, nil, err
		}
		pub, ok = cert.PublicKey.(*rsa.PublicKey)
		if ok == false {
			return pkg.NULL, nil, err
		}

		name := cert.Issuer.String()
		serialNumber := cert.SerialNumber.String()
		h := md5.New()
		h.Write([]byte(name))
		h.Write([]byte(serialNumber))
		sn = hex.EncodeToString(h.Sum(nil))
	}
	if sn == pkg.NULL {
		return pkg.NULL, nil, errors.New("failed to get sn,please check your cert")
	}
	return sn, pub, nil
}

// GetRootCertSN 获取root证书序列号SN
// rootCertPathOrData x509证书文件路径(alipayRootCert.crt) 或文件 buffer
// 返回 sn：证书序列号(alipay_root_cert_sn)
// 返回 err：error 信息
func GetRootCertSN(rootCertPathOrData any) (sn string, err error) {
	var (
		certData []byte
		certEnd  = `-----END CERTIFICATE-----`
	)
	switch pathOrData := rootCertPathOrData.(type) {
	case string:
		certData, err = os.ReadFile(pathOrData)
		if err != nil {
			return pkg.NULL, err
		}
	case []byte:
		certData = pathOrData
	default:
		return pkg.NULL, errors.New("rootCertPathOrData 断言异常")
	}

	pems := strings.Split(string(certData), certEnd)
	for _, c := range pems {
		if block, _ := pem.Decode([]byte(c + certEnd)); block != nil {
			cert, err := x509.ParseCertificate(block.Bytes)
			if err != nil {
				continue
			}
			if !allowSignatureAlgorithm[cert.SignatureAlgorithm.String()] {
				continue
			}
			name := cert.Issuer.String()
			serialNumber := cert.SerialNumber.String()
			h := md5.New()
			h.Write([]byte(name))
			h.Write([]byte(serialNumber))
			if sn == pkg.NULL {
				sn += hex.EncodeToString(h.Sum(nil))
			} else {
				sn += "_" + hex.EncodeToString(h.Sum(nil))
			}
		}
	}
	if sn == pkg.NULL {
		return pkg.NULL, errors.New("failed to get sn,please check your cert")
	}
	return sn, nil
}

// FormatAlipayPrivateKey 格式化支付宝普通应用秘钥
func FormatAlipayPrivateKey(privateKey string) (pKey string) {
	var buffer strings.Builder
	buffer.WriteString("-----BEGIN RSA PRIVATE KEY-----\n")
	rawLen := 64
	keyLen := len(privateKey)
	raws := keyLen / rawLen
	temp := keyLen % rawLen
	if temp > 0 {
		raws++
	}
	start := 0
	end := start + rawLen
	for i := 0; i < raws; i++ {
		if i == raws-1 {
			buffer.WriteString(privateKey[start:])
		} else {
			buffer.WriteString(privateKey[start:end])
		}
		buffer.WriteByte('\n')
		start += rawLen
		end = start + rawLen
	}
	buffer.WriteString("-----END RSA PRIVATE KEY-----\n")
	pKey = buffer.String()
	return
}

// FormatAlipayPublicKey 格式化支付宝普通支付宝公钥
func FormatAlipayPublicKey(publicKey string) (pKey string) {
	var buf strings.Builder
	buf.WriteString("-----BEGIN PUBLIC KEY-----\n")
	rawLen := 64
	keyLen := len(publicKey)
	raws := keyLen / rawLen
	temp := keyLen % rawLen
	if temp > 0 {
		raws++
	}
	start := 0
	end := start + rawLen
	for i := 0; i < raws; i++ {
		if i == raws-1 {
			buf.WriteString(publicKey[start:])
		} else {
			buf.WriteString(publicKey[start:end])
		}
		buf.WriteByte('\n')
		start += rawLen
		end = start + rawLen
	}
	buf.WriteString("-----END PUBLIC KEY-----\n")
	pKey = buf.String()
	return
}
