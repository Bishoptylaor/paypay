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
 @Time    : 2024/10/14 -- 17:36
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: updateOrder.go
*/

package orders

import (
	"context"
	"github.com/Bishoptylaor/paypay/paypal"
	"github.com/Bishoptylaor/paypay/paypal/consts"
	"github.com/Bishoptylaor/paypay/paypal/entity"
	"github.com/Bishoptylaor/paypay/pkg/xlog"
)

func UpdateOrder(ctx context.Context, client *paypal.Client) {
	ps := []*entity.Patch{
		&entity.Patch{
			Op:   "replace",
			Path: "/purchase_units/@reference_id=='HqSGj0ZW89jkzeUJ'/shipping/address", // reference_id is yourself set when create order
			Value: &entity.Address{
				AddressLine1: "321 Townsend St",
				AddressLine2: "Floor 7",
				AdminArea1:   "San Francisco",
				AdminArea2:   "CA",
				PostalCode:   "94107",
				CountryCode:  "US",
			},
		},
		// &entity.Patch{
		// 	Op:    "add",
		// 	Path:  "/purchase_units/@reference_id=='HqSGj0ZW89jkzeUJ'/description",
		// 	Value: "I am patch info",
		// },
	}

	res, err := client.UpdateOrder(ctx, "4HD07929GA974304A", ps)
	if err != nil {
		xlog.Error(err)
		return
	}
	if res.Code != consts.Success {
		xlog.Infof(ctx, "res.Code: %+v", res.Code)
		xlog.Infof(ctx, "res.Error: %+v", res.Error)
		xlog.Infof(ctx, "res.ErrorResponse: %+v", res.ErrorResponse)
		return
	}
	xlog.Infof(ctx, "res: %+v", res)
}
