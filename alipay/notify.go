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
 @Time    : 2024/8/28 -- 18:27
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: notify.go 主要处理回调通知的相关处理，接收 http 参数，返回解析后的通知结构体
*/

package alipay

import (
	"context"
	"github.com/Bishoptylaor/paypay/alipay/entity"
)

// https://opendocs.alipay.com/open/a5181a0b_alipay.trade.refund.depositback.completed?scene=common&pathHash=a5a8083d
func (c *Client) TradeRefundDepositbackCompleted(ctx context.Context, vals map[string][]string) (*entity.TradeRefundDepositbackCompletedReq, error) {
	// 解析 http req
	// sign 解码判断正确性
	// 比较 appid 等

	// 映射到 具体 entity
	// req := new(entity.TradeRefundDepositbackCompletedReq)

	// 返回
	return nil, nil
}

// https://opendocs.alipay.com/open/85544608_alipay.marketing.activity.delivery.changed?scene=common&pathHash=cfaaa86b
func (c *Client) MarketingActivityDeliveryChanged(ctx context.Context, vals map[string][]string) (*entity.MarketingActivityDeliveryChangedReq, error) {
	return nil, nil
}

// https://opendocs.alipay.com/open/33f45b5a_alipay.marketing.activity.deliverycreate.notify?pathHash=aea84d21
func (c *Client) MarketingActivityDeliverycreateNotify(ctx context.Context, vals map[string][]string) (*entity.MarketingActivityDeliverycreateNotifyReq, error) {
	return nil, nil
}
