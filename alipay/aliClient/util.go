package aliClient

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/Bishoptylaor/paypay"
	"github.com/Bishoptylaor/paypay/alipay/entity"
	"github.com/Bishoptylaor/paypay/alipay/utils"
	"github.com/Bishoptylaor/paypay/pkg"
)

// PublicCertDownload
// alipay.open.app.alipaycert.download(应用支付宝公钥证书下载)
// 文档地址：https://opendocs.alipay.com/apis/api_9/alipay.open.app.alipaycert.download
func (c *Client) PublicCertDownload(ctx context.Context, pl paypay.Payload) (aliRes *entity.PublicCertDownloadResponse, err error) {
	c.EmptyChecker = func(method string) []paypay.Ruler {
		_map := map[string][]paypay.Ruler{
			"alipay.open.app.alipaycert.download": []paypay.Ruler{
				paypay.NewRuler("公钥证书下载", "alipay_cert_sn != nil", "缺少 alipay_cert_sn"),
			},
		}
		if rulers, ok := _map[method]; ok {
			return rulers
		} else {
			return []paypay.Ruler{}
		}
	}

	var bs []byte
	if bs, err = c.callAli(ctx, pl, "alipay.open.app.alipaycert.download"); err != nil {
		return nil, err
	}
	aliRes = new(entity.PublicCertDownloadResponse)
	if err = json.Unmarshal(bs, aliRes); err != nil || aliRes.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", pkg.ErrUnmarshal, string(bs))
	}
	if err = utils.ExtractBussErr(aliRes.Response.ErrorResponse); err != nil {
		return aliRes, err
	}
	certBs, err := base64.StdEncoding.DecodeString(aliRes.Response.AlipayCertContent)
	if err != nil {
		return nil, fmt.Errorf("AlipayCertContent(%s)_DecodeErr:%+v", aliRes.Response.AlipayCertContent, err)
	}
	aliRes.Response.AlipayCertContent = string(certBs)
	return aliRes, nil
}
