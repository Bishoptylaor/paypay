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
 @Time    : 2024/9/3 -- 15:37
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: consts.go
*/

package consts

const (
	HeaderAuthorization       = "Authorization" // 请求头Auth
	AuthorizationPrefixBasic  = "Basic "
	AuthorizationPrefixBearer = "Bearer "

	BaseUrl        = "https://api-m.paypal.com"         // 正式 URL
	SandboxBaseUrl = "https://api-m.sandbox.paypal.com" // 沙箱 URL

	// AccessTokenPath 获取链接路径
	AccessTokenPath = "/v1/oauth2/token" // 获取AccessToken POST

	Success = 0
)
