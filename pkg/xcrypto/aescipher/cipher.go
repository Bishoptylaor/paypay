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
 @Time    : 2024/8/30 -- 15:05
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: cipher.go
*/

package aescipher

import (
	"context"
	"github.com/Bishoptylaor/paypay/pkg/xcrypto"
)

// AesCipher 标准参数加密解密功能入口
type AesCipher interface {
	Encrypt(ctx context.Context, plainText, secretKey, ivAes []byte, pad xcrypto.Pad) (cipherText []byte, err error)
	EncryptBase64(ctx context.Context, plainText, secretKey, ivAes []byte, pad xcrypto.Pad) (cipherText string, err error)

	Decrypt(ctx context.Context, cipherText, secretKey, ivAes []byte, pad xcrypto.Pad) (plainText []byte, err error)
	DecryptBase64(ctx context.Context, cipherText, secretKey, ivAes []byte, pad xcrypto.Pad) (plainText []byte, err error)
}
