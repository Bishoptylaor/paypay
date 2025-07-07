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
 @Time    : 2024/12/26 -- 09:35
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: interface.go
*/

package pkg

import (
	"context"
	"github.com/Bishoptylaor/paypay/pkg/xlog"
	"github.com/Bishoptylaor/paypay/pkg/xutils"
	"runtime"
	"time"
)

type AutoToken interface {
	GetToken(ctx context.Context) (interface{}, error)
	Wait()
	MaxRetry() int
}

func AutoRefreshToken(ctx context.Context, autoToken AutoToken, logger xlog.XLogger) {
	defer func() {
		if r := recover(); r != nil {
			buf := make([]byte, 64<<10)
			buf = buf[:runtime.Stack(buf, false)]
			logger.Errorf("paypal_goAuthRefreshToken: panic recovered: %s\n%s", r, buf)
		}
	}()
	for {
		autoToken.Wait()
		err := xutils.Retry(func() error {
			_, err := autoToken.GetToken(ctx)
			if err != nil {
				return err
			}
			return nil
		}, autoToken.MaxRetry(), time.Second)
		if err != nil {
			logger.Errorf("PayPal GetAccessToken Error: %s", err.Error())
		}
	}
}
