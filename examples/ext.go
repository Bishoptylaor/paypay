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
 @Time    : 2024/10/14 -- 16:47
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: ext.go
*/

package examples

import (
	"context"
	"github.com/Bishoptylaor/paypay/pkg/xlog"
	"reflect"
)

func Equal(prefix func([]byte) map[string]interface{}, res, example func(func([]byte) map[string]interface{}) map[string]interface{}) bool {
	ctx := context.TODO()
	xlog.Infof(ctx, "res: %+v", res(prefix))
	xlog.Infof(ctx, "exa: %+v", example(prefix))
	equal := reflect.DeepEqual(res(prefix), example(prefix))
	xlog.Infof(ctx, "exampleJson and resJson are equal: %t", equal)
	return equal
}
