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
 @Time    : 2024/10/14 -- 13:38
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: products.go
*/

package paypal

import (
	"context"
	"github.com/Bishoptylaor/paypay"
	"github.com/Bishoptylaor/paypay/paypal/consts"
	"github.com/Bishoptylaor/paypay/paypal/entity"
	"github.com/Bishoptylaor/paypay/pkg"
)

// CreateProduct
// 创建目录商品（Create product）
// 文档：https://developer.paypal.com/docs/api/catalog-products/v1/#products_create
func (c *Client) CreateProduct(ctx context.Context, pl paypay.Payload) (res *entity.CreateProductRes, err error) {
	method := CreateProduct
	c.EmptyChecker = method.Checker

	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, nil), pl, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.CreateProductRes{EmptyRes: emptyRes}
	res.Response = new(entity.ProductDetail)
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, res.Response); err != nil {
		return res, err
	}
	return res, nil
}

// ListProducts
// 商品列表（List products）
// 文档：https://developer.paypal.com/docs/api/catalog-products/v1/#products_list
func (c *Client) ListProducts(ctx context.Context, query paypay.Payload) (res *entity.ListProductsRes, err error) {
	method := ListProducts

	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"params": query.EncodeURLParams(),
	}), nil, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.ListProductsRes{EmptyRes: emptyRes}
	res.Response = new(entity.ProductList)
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, res.Response); err != nil {
		return res, err
	}
	return res, nil
}

// ShowProductDetails
// 查看商品详情（Show product details）
// 文档：https://developer.paypal.com/docs/api/catalog-products/v1/#products_get
func (c *Client) ShowProductDetails(ctx context.Context, productId string) (res *entity.ShowProductDetailsRes, err error) {
	method := ShowProductDetails
	if productId == pkg.NULL {
		return nil, pkg.ErrPaypalMissingProductId
	}

	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"product_id": productId,
	}), nil, nil)
	if err != nil {
		return nil, err
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.ShowProductDetailsRes{EmptyRes: emptyRes}
	res.Response = new(entity.ProductDetail)
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, res.Response); err != nil {
		return res, err
	}
	return res, nil
}

// UpdateProduct
// 更新订单（Update product）
// 文档：https://developer.paypal.com/docs/api/catalog-products/v1/#products_patch
func (c *Client) UpdateProduct(ctx context.Context, productId string, patches []*entity.Patch) (res *entity.UpdateProductRes, err error) {
	method := UpdateProduct
	if productId == pkg.NULL {
		return nil, pkg.ErrPaypalMissingProductId
	}

	httpRes, bs, err := method.Do(c)(ctx, method.Uri, c.GenUrl(ctx, map[string]string{
		"product_id": productId,
	}), nil, patches)
	if err != nil {
		return nil, pkg.WrapError("[UpdateProduct] do method: ", err)
	}
	emptyRes := entity.EmptyRes{Code: consts.Success}
	res = &entity.UpdateProductRes{EmptyRes: emptyRes}
	if err = c.handleResponse(ctx, method, httpRes, bs, &emptyRes, nil); err != nil {
		return res, pkg.WrapError("[UpdateProduct]: ", err)
	}
	return res, nil
}
