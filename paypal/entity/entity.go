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
 @Time    : 2024/9/3 -- 17:10
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: entity.go
*/

package entity

import (
	"fmt"
)

type APIError struct {
	Code    int
	Message string
	Details interface{}
}

func (e *APIError) Error() string {
	return fmt.Sprintf("%d:%s:%+v", e.Code, e.Message, e.Details)
}

type ErrorResponse struct {
	Name    string        `json:"name,omitempty"`
	Message string        `json:"message,omitempty"`
	DebugId string        `json:"debug_id,omitempty"`
	Details []ErrorDetail `json:"details,omitempty"`
	Links   []Link        `json:"links,omitempty"`
}

type ErrorDetail struct {
	Issue       string `json:"issue,omitempty"`
	Field       string `json:"field,omitempty"`
	Value       string `json:"value,omitempty"`
	Description string `json:"description,omitempty"`
	Location    string `json:"location,omitempty"`
}

type Link struct {
	Href   string `json:"href,omitempty"`
	Rel    string `json:"rel,omitempty"`
	Method string `json:"method,omitempty"` // Possible values: GET,POST,PUT,DELETE,HEAD,CONNECT,OPTIONS,PATCH
}

type CommonResponse struct {
	Code          int            `json:"-"`
	Error         string         `json:"-"`
	ErrorResponse *ErrorResponse `json:"-"`
}

type AccessToken struct {
	Scope       string `json:"scope"`
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Appid       string `json:"app_id"`
	ExpiresIn   int    `json:"expires_in"`
	Nonce       string `json:"nonce"`
}

type Patch struct {
	Op    string `json:"op"` // The possible values are: add、remove、replace、move、copy、test
	Path  string `json:"path,omitempty"`
	Value any    `json:"value"` // The value to apply. The remove operation does not require a value.
	From  string `json:"from,omitempty"`
}

type EmptyRes struct {
	Code          int            `json:"-"`
	Error         string         `json:"-"`
	ErrorResponse *ErrorResponse `json:"-"`
}

// =========================================================分割=========================================================
