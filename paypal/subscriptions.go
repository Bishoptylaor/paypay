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
 @Time    : 2024/10/8 -- 16:41
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: subscriptions.go
*/

package paypal

import (
	"context"
	"github.com/Bishoptylaor/paypay"
	"github.com/Bishoptylaor/paypay/paypal/consts"
	"github.com/Bishoptylaor/paypay/paypal/entity"
	"github.com/Bishoptylaor/paypay/pkg"
)

// CreatePlan
// 创建计划（Create Plan）
// 文档：https://developer.paypal.com/docs/api/subscriptions/v1/#plans_create
func (c *Client) CreatePlan(ctx context.Context, pl paypay.Payload) (res *entity.CreatePlanRes, err error) {
	method := CreatePlan
	c.EmptyChecker = method.Checker

	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, nil), pl, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.CreatePlanRes{EmptyRes: emptyRes}
	res.Response = new(entity.BillingDetail)
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, res.Response); err != nil {
		return res, err
	}
	return res, nil
}

// ListPlans
// 展示订阅计划（List plans）
// 文档：https://developer.paypal.com/docs/api/subscriptions/v1/#plans_list
func (c *Client) ListPlans(ctx context.Context, query paypay.Payload) (res *entity.ListPlansRes, err error) {
	method := ListPlans
	c.EmptyChecker = method.Checker

	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"params": query.EncodeURLParams(),
	}), query, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.ListPlansRes{EmptyRes: emptyRes}
	res.Response = new(entity.PlanDetails)
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, res.Response); err != nil {
		return res, err
	}
	return res, nil
}

// ShowPlanDetails
// 展示计划（show plan details）
// 文档：https://developer.paypal.com/docs/api/subscriptions/v1/#plans_get
func (c *Client) ShowPlanDetails(ctx context.Context, planId string) (res *entity.ShowPlanDetailsRes, err error) {
	method := ShowPlanDetails
	if planId == pkg.NULL {
		return nil, pkg.ErrPaypalMissingPlanId
	}

	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"plan_id": planId,
	}), nil, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.ShowPlanDetailsRes{EmptyRes: emptyRes}
	res.Response = new(entity.BillingDetail)
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, res.Response); err != nil {
		return res, err
	}
	return res, nil
}

// UpdatePlan
// 更新计划（Update plan）
// 文档：https://developer.paypal.com/docs/api/subscriptions/v1/#plans_patch
func (c *Client) UpdatePlan(ctx context.Context, planId string, patches []*entity.Patch) (res *entity.UpdatePlanRes, err error) {
	method := UpdatePlan
	if planId == pkg.NULL {
		return nil, pkg.ErrPaypalMissingPlanId
	}
	if len(patches) == 0 {
		return nil, pkg.ErrPaypalNothingToChange
	}

	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"plan_id": planId,
	}), nil, patches)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.UpdatePlanRes{EmptyRes: emptyRes}
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, nil); err != nil {
		return res, err
	}
	return res, nil
}

// ActivePlan
// 激活计划（Active plan）
// 文档：https://developer.paypal.com/docs/api/subscriptions/v1/#plans_activate
func (c *Client) ActivePlan(ctx context.Context, planId string) (res *entity.ActivePlanRes, err error) {
	method := ActivePlan
	if planId == pkg.NULL {
		return nil, pkg.ErrPaypalMissingPlanId
	}

	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"plan_id": planId,
	}), nil, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.ActivePlanRes{EmptyRes: emptyRes}
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, nil); err != nil {
		return res, err
	}
	return res, nil
}

// DeactivePlan
// 暂停计划（Deactive plan）
// 文档：https://developer.paypal.com/docs/api/subscriptions/v1/#plans_deactivate
func (c *Client) DeactivePlan(ctx context.Context, planId string) (res *entity.DeactivePlanRes, err error) {
	method := DeactivePlan
	if planId == pkg.NULL {
		return nil, pkg.ErrPaypalMissingPlanId
	}

	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"plan_id": planId,
	}), nil, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.DeactivePlanRes{EmptyRes: emptyRes}
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, nil); err != nil {
		return res, err
	}
	return res, nil
}

// UpdatePricing
// 更新计划价格方案（Update pricing）
// 文档：https://developer.paypal.com/docs/api/subscriptions/v1/#plans_update-pricing-schemes
func (c *Client) UpdatePricing(ctx context.Context, planId string) (res *entity.UpdatePricingRes, err error) {
	method := DeactivePlan
	if planId == pkg.NULL {
		return nil, pkg.ErrPaypalMissingPlanId
	}

	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"plan_id": planId,
	}), nil, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.UpdatePricingRes{EmptyRes: emptyRes}
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, nil); err != nil {
		return res, err
	}
	return res, nil
}

// CreateSubscription
// 创建订阅（Create subscription）
// 文档：https://developer.paypal.com/docs/api/subscriptions/v1/#subscriptions_create
func (c *Client) CreateSubscription(ctx context.Context, pl paypay.Payload) (res *entity.CreateSubscriptionRes, err error) {
	method := CreatePlan
	c.EmptyChecker = method.Checker

	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, nil), pl, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.CreateSubscriptionRes{EmptyRes: emptyRes}
	res.Response = new(entity.SubscriptionDetail)
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, res.Response); err != nil {
		return res, err
	}
	return res, nil
}

// ShowSubscriptionDetails
// 查看订阅详情（Show subscription details）
// 文档：https://developer.paypal.com/docs/api/subscriptions/v1/#subscriptions_get
func (c *Client) ShowSubscriptionDetails(ctx context.Context, subscriptionId string) (res *entity.ShowSubscriptionDetailsRes, err error) {
	method := ShowSubscriptionDetails
	if subscriptionId == pkg.NULL {
		return nil, pkg.ErrPaypalMissingSubscriptionId
	}

	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"subscription_id": subscriptionId,
	}), nil, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.ShowSubscriptionDetailsRes{EmptyRes: emptyRes}
	res.Response = new(entity.SubscriptionDetail)
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, res.Response); err != nil {
		return res, err
	}
	return res, nil
}

// UpdateSubscription
// 更新订阅（Update subscription）
// 文档：https://developer.paypal.com/docs/api/subscriptions/v1/#subscriptions_patch
func (c *Client) UpdateSubscription(ctx context.Context, subscriptionId string, patches []*entity.Patch) (res *entity.UpdateSubscriptionRes, err error) {
	method := UpdateSubscription
	if subscriptionId == pkg.NULL {
		return nil, pkg.ErrPaypalMissingSubscriptionId
	}
	if len(patches) == 0 {
		return nil, pkg.ErrPaypalNothingToChange
	}

	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"subscription_id": subscriptionId,
	}), nil, patches)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.UpdateSubscriptionRes{EmptyRes: emptyRes}
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, nil); err != nil {
		return res, err
	}
	return res, nil
}

// ShowSubscriptionDetails
// 更新计划或者数量（Revise plan or quantity of subscription）
// 文档：https://developer.paypal.com/docs/api/subscriptions/v1/#subscriptions_revise
func (c *Client) RevisePlanOrQuantityOfSubsription(ctx context.Context, subscriptionId string) (res *entity.RevisePlanOrQuantityOfSubsriptionRes, err error) {
	method := RevisePlanOrQuantityOfSubsription
	if subscriptionId == pkg.NULL {
		return nil, pkg.ErrPaypalMissingSubscriptionId
	}

	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"subscription_id": subscriptionId,
	}), nil, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.RevisePlanOrQuantityOfSubsriptionRes{EmptyRes: emptyRes}
	res.Response = new(entity.SubscriptionDetail)
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, res.Response); err != nil {
		return res, err
	}
	return res, nil
}

// SuspendSubscription
// 暂定订阅（Suspend subscription）
// 文档：https://developer.paypal.com/docs/api/subscriptions/v1/#subscriptions_suspend
func (c *Client) SuspendSubscription(ctx context.Context, subscriptionId string, pl paypay.Payload) (res *entity.SuspendSubscriptionRes, err error) {
	method := SuspendSubscription
	c.EmptyChecker = method.Checker
	if subscriptionId == pkg.NULL {
		return nil, pkg.ErrPaypalMissingSubscriptionId
	}

	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"subscription_id": subscriptionId,
	}), pl, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.SuspendSubscriptionRes{EmptyRes: emptyRes}
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, nil); err != nil {
		return res, err
	}
	return res, nil
}

// CancelSubscription
// 取消订阅（Cancel subscription）
// 文档：https://developer.paypal.com/docs/api/subscriptions/v1/#subscriptions_cancel
func (c *Client) CancelSubscription(ctx context.Context, subscriptionId string, pl paypay.Payload) (res *entity.CancelSubscriptionRes, err error) {
	method := CancelSubscription
	c.EmptyChecker = method.Checker
	if subscriptionId == pkg.NULL {
		return nil, pkg.ErrPaypalMissingSubscriptionId
	}

	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"subscription_id": subscriptionId,
	}), pl, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.CancelSubscriptionRes{EmptyRes: emptyRes}
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, nil); err != nil {
		return res, err
	}
	return res, nil
}

// ActivateSubscription
// 取消订阅（Activate subscription）
// 文档：https://developer.paypal.com/docs/api/subscriptions/v1/#subscriptions_activate
func (c *Client) ActivateSubscription(ctx context.Context, subscriptionId string, pl paypay.Payload) (res *entity.ActivateSubscriptionRes, err error) {
	method := ActivateSubscription
	c.EmptyChecker = method.Checker
	if subscriptionId == pkg.NULL {
		return nil, pkg.ErrPaypalMissingSubscriptionId
	}

	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"subscription_id": subscriptionId,
	}), pl, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.ActivateSubscriptionRes{EmptyRes: emptyRes}
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, nil); err != nil {
		return res, err
	}
	return res, nil
}

// CaptureAuthoriedPaymentOnSubscription
// 捕获订阅的授权支付信息（Capture authorized payment on subscription）
// 文档：https://developer.paypal.com/docs/api/subscriptions/v1/#subscriptions_capture
func (c *Client) CaptureAuthoriedPaymentOnSubscription(ctx context.Context, subscriptionId string, pl paypay.Payload) (res *entity.CaptureAuthoriedPaymentOnSubscriptionRes, err error) {
	method := ActivateSubscription
	c.EmptyChecker = method.Checker
	if subscriptionId == pkg.NULL {
		return nil, pkg.ErrPaypalMissingSubscriptionId
	}

	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"subscription_id": subscriptionId,
	}), pl, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.CaptureAuthoriedPaymentOnSubscriptionRes{EmptyRes: emptyRes}
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, nil); err != nil {
		return res, err
	}
	return res, nil
}

// ListTransactions4Subscription
// 列出一个订阅的所有交易记录 (List transactions for subscription)
// 文档：https://developer.paypal.com/docs/api/subscriptions/v1/#subscriptions_transactions
func (c *Client) ListTransactions4Subscription(ctx context.Context, subscriptionId string, query paypay.Payload) (res *entity.ListTransactions4SubscriptionRes, err error) {
	method := ListTransactions4Subscription
	c.EmptyChecker = method.Checker
	if subscriptionId == pkg.NULL {
		return nil, pkg.ErrPaypalMissingSubscriptionId
	}

	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"subscription_id": subscriptionId,
		"params":          query.EncodeURLParams(),
	}), query, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.ListTransactions4SubscriptionRes{EmptyRes: emptyRes}
	res.Response = new(entity.SubscriptionDetail)
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, res.Response); err != nil {
		return res, err
	}
	return res, nil
}
