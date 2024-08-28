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
 @Time    : 2024/8/27 -- 18:42
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: readfunc.go
*/

package zhttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Bishoptylaor/paypay"
	"github.com/Bishoptylaor/paypay/pkg/zutils"
	"io"
	"mime/multipart"
	"net/url"
	"sort"
	"strings"
)

var _ReqContentTypeReader = map[string]func(cfg *httpConfig) readFunc{
	TypeJSON:              readJson,
	TypeXML:               readXML,
	TypeFormData:          readForm,
	TypeMultipartFormData: readFile,
}

type readFunc func(map[string]any) (io.Reader, error)

func defaultReadFunc(c *httpConfig) readFunc {
	return readJson(c)
}

func readJson(c *httpConfig) readFunc {
	return func(m map[string]any) (io.Reader, error) {
		bs, _ := json.Marshal(m)
		return strings.NewReader(string(bs)), nil
	}
}

func readXML(c *httpConfig) readFunc {
	return func(m map[string]any) (io.Reader, error) {
		return strings.NewReader(FormatURLParam(m)), nil
	}
}

func readForm(c *httpConfig) readFunc {
	return func(m map[string]any) (io.Reader, error) {
		return strings.NewReader(FormatURLParam(m)), nil
	}
}

func readFile(c *httpConfig) readFunc {
	return func(m map[string]any) (io.Reader, error) {
		var (
			body        io.Reader
			fileContent *multipart.Writer
		)

		if c.requestType == TypeMultipartFormData {
			body = &bytes.Buffer{}
			fileContent = multipart.NewWriter(body.(io.Writer))
		}
		for k, v := range m {
			// file 参数
			if file, ok := v.(*paypay.File); ok {
				fw, e := fileContent.CreateFormFile(k, file.Name)
				if e != nil {
					return body, fmt.Errorf("create form file error: %v", e)
				}
				_, _ = fw.Write(file.Content)
				continue
			}
			// text 参数
			switch vs := v.(type) {
			case string:
				_ = fileContent.WriteField(k, vs)
			default:
				_ = fileContent.WriteField(k, zutils.Any2String(v))
			}
		}
		_ = fileContent.Close()
		c.headers.Set("Content-Type", fileContent.FormDataContentType())
		return body, nil
	}
}

func FormatURLParam(body map[string]any) (urlParam string) {
	var (
		buf  strings.Builder
		keys []string
	)
	for k := range body {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		v, ok := body[k].(string)
		if !ok {
			v = zutils.Any2String(body[k])
		}
		if v != "" {
			buf.WriteString(url.QueryEscape(k))
			buf.WriteByte('=')
			buf.WriteString(url.QueryEscape(v))
			buf.WriteByte('&')
		}
	}
	if buf.Len() <= 0 {
		return ""
	}
	return buf.String()[:buf.Len()-1]
}
