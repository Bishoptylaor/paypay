package pkg

import "errors"

var (
	ErrUnmarshal = errors.New("json unmarshal error")

	ErrMissingInitParams = errors.New("缺少初始化必备参数")
	ErrPhrasePrivateKey  = errors.New("私钥格式有误，请检查")
)
