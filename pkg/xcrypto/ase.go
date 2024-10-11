package xcrypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/Bishoptylaor/paypay/pkg/xutils"
	"golang.org/x/crypto/pbkdf2"
	"hash"
)

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
 @Time    : 2024/7/13 -- 14:00
 @Author  : bishop
 @Description: 加解密函数
*/

const (
	pkc5SaltLen          = 8
	pkc5DefaultIter      = 2048
	pkc5DefaultMagicWord = "Salted__"
	maxIvLen             = 16
)

// aesCBCEncrypt 加密 由key的长度决定是128, 192 还是 256
func aesCBCEncrypt(origData, key, iv []byte, pad Pad) ([]byte, error) {
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
func aesCBCDecrypt(encrypted, key, iv []byte, pad Pad) ([]byte, error) {
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
func AESCBCEncrypt(origData, key, iv []byte, pad Pad) ([]byte, error) {
	return aesCBCEncrypt(origData, key, iv, pad)
}

// AESCBCEncryptWithBase64 加密 结果返回base64编码后的string
func AESCBCEncryptWithBase64(origData, key, iv []byte, pad Pad) (string, error) {
	bs, err := aesCBCEncrypt(origData, key, iv, pad)
	if err != nil {
		return "", err
	}
	return Base64.Encode(bs), nil
}

// AESCBCDecrypt 普通解密
func AESCBCDecrypt(encrypted, key, iv []byte, pad Pad) ([]byte, error) {
	return aesCBCDecrypt(encrypted, key, iv, pad)
}

// AESCBCDecryptWithBase64 base64编码后的加密串，返回原始数据
func AESCBCDecryptWithBase64(baseEncrypted, key, iv []byte, pad Pad) (string, error) {
	encrypted, err := Base64.Decode(string(baseEncrypted))
	if err != nil {
		return "", err
	}
	bs, err := aesCBCEncrypt(encrypted, key, iv, pad)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

func AESCBCEncryptWithSalt(origData, key []byte, iter int, magic string, pad Pad, h func() hash.Hash) ([]byte, error) {
	return AESEncryptWithSalt(origData, key, iter, magic, h, pad, AESCBCEncrypt)
}

func AESCBCDecryptWithSalt(data, key []byte, iter int, magic string, pad Pad, h func() hash.Hash) ([]byte, error) {
	return AESDecryptWithSalt(data, key, iter, magic, h, pad, AESCBCDecrypt)
}

func AESEncryptWithSalt(origData, key []byte, iter int, magic string, h func() hash.Hash, pad Pad, f func(origData, key, iv []byte, pad Pad) ([]byte, error)) ([]byte, error) {
	if iter <= 0 {
		iter = pkc5DefaultIter
	}

	if h == nil {
		h = md5.New
	}

	var salt = xutils.RandomString(pkc5SaltLen)
	var sKey = pbkdf2.Key(key, []byte(salt), iter, len(key), h)
	var sIV = pbkdf2.Key(sKey, []byte(salt), iter, maxIvLen, h)

	var encrypted, err = f(origData, sKey, sIV, pad)

	encrypted = append([]byte(salt), encrypted...)
	encrypted = append([]byte(magic), encrypted...)

	return encrypted, err
}

func AESDecryptWithSalt(encrypted, key []byte, iterCount int, magic string, h func() hash.Hash, pad Pad, f func(encrypted, key, iv []byte, pad Pad) ([]byte, error)) ([]byte, error) {
	if iterCount <= 0 {
		iterCount = pkc5DefaultIter
	}

	if h == nil {
		h = md5.New
	}

	var salt = encrypted[len(magic) : len(magic)+pkc5SaltLen]
	var sKey = pbkdf2.Key(key, salt, iterCount, len(key), h)
	var sIV = pbkdf2.Key(sKey, salt, iterCount, maxIvLen, h)

	var plaintext, err = f(encrypted[len(magic)+pkc5SaltLen:], sKey, sIV, pad)
	return plaintext, err
}

func AESCFBEncrypt(origData, key, iv []byte, pad Pad) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	origData = pad.Padding(origData, blockSize)
	blockMode := cipher.NewCFBEncrypter(block, iv[:blockSize])
	encrypted := make([]byte, len(origData))
	blockMode.XORKeyStream(encrypted, origData)
	return encrypted, nil
}

func AESCFBDecrypt(encrypted, key, iv []byte, pad Pad) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	blockMode := cipher.NewCFBDecrypter(block, iv[:blockSize]) // 初始向量的长度必须等于块block的长度
	origData := make([]byte, len(encrypted))
	blockMode.XORKeyStream(origData, encrypted)
	return pad.UnPadding(origData, blockSize)
}

func AESECBEncrypt(origData, key []byte, pad Pad) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	origData = pad.Padding(origData, blockSize)
	encrypted := make([]byte, len(origData))
	var start = 0
	var end = blockSize
	for start < len(origData) {
		block.Encrypt(encrypted[start:end], origData[start:end])
		start += blockSize
		end += blockSize
	}
	return encrypted, nil
}

func AESECBDecrypt(encrypted, key []byte, pad Pad) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	origData := make([]byte, len(encrypted))
	var start = 0
	var end = blockSize

	for start < len(encrypted) {
		block.Decrypt(origData[start:end], encrypted[start:end])
		start = start + blockSize
		end = end + blockSize
	}
	return pad.UnPadding(origData, blockSize)
}

func AESGCMEncrypt(origData, key, additional []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockMode, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := xutils.RandomString(blockMode.NonceSize())

	return blockMode.Seal([]byte(nonce), []byte(nonce), origData, additional), nil
}

func AESGCMEncryptWithNonce(origData, key, nonce, additional []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockMode, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	if len(nonce) != blockMode.NonceSize() {
		return nil, fmt.Errorf("invalid nonce size, must contain %d characters", blockMode.NonceSize())
	}

	return blockMode.Seal(nil, nonce, origData, additional), nil
}

func AESGCMDecrypt(encrypted, key, additional []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockMode, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := blockMode.NonceSize()
	if len(encrypted) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	var nonce []byte
	nonce, encrypted = encrypted[:nonceSize], encrypted[nonceSize:]
	return blockMode.Open(nil, nonce, encrypted, additional)
}

func AESGCMDecryptWithNonce(encrypted, key, nonce, additional []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockMode, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	if len(nonce) != blockMode.NonceSize() {
		return nil, fmt.Errorf("invalid nonce size, must contain %d characters", blockMode.NonceSize())
	}

	return blockMode.Open(nil, nonce, encrypted, additional)
}
