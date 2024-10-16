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
 @Time    : 2024/10/10 -- 17:46
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: transcation.go
*/

package paypal

import (
	"context"
	"github.com/Bishoptylaor/paypay"
	"github.com/Bishoptylaor/paypay/paypal/consts"
	"github.com/Bishoptylaor/paypay/paypal/entity"
)

// ListTransactions
// 交易列表 (List transactions)
// 文档：https://developer.paypal.com/docs/api/transaction-search/v1/#search_get
func (c *Client) ListTransactions(ctx context.Context, query paypay.Payload) (res *entity.ListTransactionsRes, err error) {
	method := ListTransactions
	c.EmptyChecker = method.Checker

	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"params": query.EncodeURLParams(),
	}), query, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.ListTransactionsRes{EmptyRes: emptyRes}
	res.Response = new(entity.TransactionInfos)
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, res.Response); err != nil {
		return res, err
	}
	return res, nil
}

// ListAllBalances
// 获取所有余额 (List all balances)
// 文档：https://developer.paypal.com/docs/api/transaction-search/v1/#balances_get
func (c *Client) ListAllBalances(ctx context.Context, query paypay.Payload) (res *entity.ListAllBalancesRes, err error) {
	method := ListAllBalances
	c.EmptyChecker = method.Checker

	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"params": query.EncodeURLParams(),
	}), query, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.ListAllBalancesRes{EmptyRes: emptyRes}
	res.Response = new(entity.AllBalances)
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, res.Response); err != nil {
		return res, err
	}
	return res, nil
}
