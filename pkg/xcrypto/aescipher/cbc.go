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
 @Time    : 2024/8/30 -- 16:34
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: cbc.go
*/

package aescipher

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"github.com/Bishoptylaor/paypay/pkg/xcrypto"
)

type CBCEncrypter struct{}

// aesCBCEncrypt 加密 由key的长度决定是128, 192 还是 256
func aesCBCEncrypt(_ context.Context, origData, key, iv []byte, pad xcrypto.Pad) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// AES分组长度为128位，所以blockSize=16，单位字节
	// 此处没有固定key的长度，可以由上层调用再封装常用功能
	blockSize := block.BlockSize()
	origData = pad.Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv[:blockSize]) // 初始向量的长度必须等于块block的长度16字节
	encrypted := make([]byte, len(origData))
	blockMode.CryptBlocks(encrypted, origData)
	return encrypted, nil
}

// aesCBCDecrypt 解密
func aesCBCDecrypt(_ context.Context, encrypted, key, iv []byte, pad xcrypto.Pad) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// AES分组长度为128位，所以blockSize=16，单位字节
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, iv[:blockSize]) // 初始向量的长度必须等于块block的长度
	origData := make([]byte, len(encrypted))
	blockMode.CryptBlocks(origData, encrypted)
	origData, _ = pad.UnPadding(origData, blockSize)
	return origData, nil
}

// AESCBCEncrypt 普通加密
func (CBCEncrypter) Encrypt(ctx context.Context, origData, key, iv []byte, pad xcrypto.Pad) ([]byte, error) {
	return aesCBCEncrypt(ctx, origData, key, iv, pad)
}

// AESCBCEncryptWithBase64 加密 结果返回base64编码后的string
func (CBCEncrypter) EncryptBase64(ctx context.Context, origData, key, iv []byte, pad xcrypto.Pad) (string, error) {
	bs, err := aesCBCEncrypt(ctx, origData, key, iv, pad)
	if err != nil {
		return "", err
	}
	return xcrypto.Base64.Encode(bs), nil
}

// AESCBCDecrypt 普通解密
func (CBCEncrypter) Decrypt(ctx context.Context, encrypted, key, iv []byte, pad xcrypto.Pad) ([]byte, error) {
	return aesCBCDecrypt(ctx, encrypted, key, iv, pad)
}

// AESCBCDecryptWithBase64 base64编码后的加密串，返回原始数据
func (CBCEncrypter) DecryptBase64(ctx context.Context, baseEncrypted, key, iv []byte, pad xcrypto.Pad) ([]byte, error) {
	encrypted, err := xcrypto.Base64.Decode(string(baseEncrypted))
	if err != nil {
		return []byte{}, err
	}
	bs, err := aesCBCEncrypt(ctx, encrypted, key, iv, pad)
	if err != nil {
		return []byte{}, err
	}
	return bs, nil
}
