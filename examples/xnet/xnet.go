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
 @Time    : 2024/8/28 -- 10:53
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: znet.go
*/

package main

import (
	"context"
	"fmt"
	"github.com/Bishoptylaor/paypay/pkg/xnet/xhttp"
)

func main() {
	client := xhttp.GetDefaultClient()
	res, bs, err := client.Call(context.Background(), "GET", "https://www.baidu.com", nil, nil)
	fmt.Printf("[res] %+v \n", res)
	fmt.Println("[body] ", string(bs))
	fmt.Println("[err] ", err)

	// todo post
	// todo 上传文件
	// todo header
	// todo cookie
}
