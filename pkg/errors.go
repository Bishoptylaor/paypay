package pkg

import "errors"

func WrapError(info string, err error) error {
	if err == nil {
		return nil
	}
	return errors.New(info + ": " + err.Error())
}

var (
	ErrUnmarshal       = errors.New("json unmarshal error")
	ErrMissParam       = errors.New("miss param error")
	ErrGetSignData     = errors.New("get signature data error")
	ErrCertNotMatch    = errors.New("cert not match error")
	ErrVerifySignature = errors.New("verify signature error")

	ErrMissingInitParams = errors.New("缺少初始化必备参数")
	ErrMissingInitAppid  = errors.New("缺少 appid")
	ErrMissingInitLogger = errors.New("未设置 logger")
	ErrMissingInitHttp   = errors.New("未设置 http client")
	ErrPhrasePrivateKey  = errors.New("私钥格式有误，请检查")
)

var (
	ErrAliBadResponse       = errors.New("alipay: bad response")
	ErrAliSignNotFound      = errors.New("alipay: sign content not found")
	ErrAliPublicKeyNotFound = errors.New("alipay: alipay public key not found")
)

var (
	ErrPaypalNothingToChange             = errors.New("paypal: nothing to change")
	ErrPaypalMissingOrderId              = errors.New("paypal: missing order id")
	ErrPaypalMissingInvoiceId            = errors.New("paypal: missing invoice id")
	ErrPaypalMissingTransactionId        = errors.New("paypal: missing transaction id")
	ErrPaypalMissingTemplateId           = errors.New("paypal: missing template id")
	ErrPaypalMissingInvoiceTransactionId = errors.New("paypal: missing invoice or transaction id")
	ErrPaypalMissingRefundId             = errors.New("paypal: missing refund id")
	ErrPaypalMissingCaptureId            = errors.New("paypal: missing capture id")
	ErrPaypalMissingAuthorizeId          = errors.New("paypal: missing authorize id")
	ErrPaypalMissingPayoutBatchId        = errors.New("paypal: missing payout batch id")
	ErrPaypalMissingPayoutItemId         = errors.New("paypal: missing payout item id")
	ErrPaypalMissingSubscriptionId       = errors.New("paypal: missing subscription id")
	ErrPaypalMissingPlanId               = errors.New("paypal: missing plan id")
	ErrPaypalMissingProductId            = errors.New("paypal: missing product id")
	ErrPaypalMissingQueryId              = errors.New("paypal: missing query id")
)
