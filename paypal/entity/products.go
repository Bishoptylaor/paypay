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

package entity

type CreateProductRes struct {
	EmptyRes
	Response *ProductDetail `json:"response,omitempty"`
}

type ListProductsRes struct {
	EmptyRes
	Response *ProductList `json:"response,omitempty"`
}

type ShowProductDetailsRes struct {
	EmptyRes
	Response *ProductDetail `json:"response,omitempty"`
}

type UpdateProductRes struct {
	EmptyRes
}

// =========================================================分割=========================================================

type ProductList struct {
	Products   []*ProductDetail `json:"products,omitempty"`
	TotalItems int              `json:"total_items,omitempty"`
	TotalPages int              `json:"total_pages,omitempty"`
	Links      []*Link          `json:"links,omitempty"`
}

type ProductDetail struct {
	Id          string  `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
	Type        string  `json:"type,omitempty"`
	Category    string  `json:"category,omitempty"`
	ImageUrl    string  `json:"image_url,omitempty"`
	HomeUrl     string  `json:"home_url,omitempty"`
	Links       []*Link `json:"links,omitempty"`
	CreateTime  string  `json:"create_time,omitempty"`
	UpdateTime  string  `json:"update_time,omitempty"`
}
