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
 @Description: reader.go
*/

package xhttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Bishoptylaor/paypay"
	"github.com/Bishoptylaor/paypay/pkg/xutils"
	"io"
	"mime/multipart"
	"net/url"
	"sort"
	"strings"
)

var _ReqContentTypeReader = map[string]func(cfg *httpConfig) reader{
	TypeJSON:              readJson,
	TypeXML:               readXML,
	TypeFormData:          readForm,
	TypeMultipartFormData: readFile,
}

type reader func(any) (io.Reader, error)

func defaultReader() CfgOp {
	return Req(TypeJSON)
}

func readJson(c *httpConfig) reader {
	return func(v any) (io.Reader, error) {
		bs, _ := json.Marshal(v)
		return strings.NewReader(string(bs)), nil
	}
}

func readXML(c *httpConfig) reader {
	return func(v any) (io.Reader, error) {
		if bs, err := json.Marshal(v); err == nil {
			m := make(map[string]any)
			_ = json.Unmarshal(bs, &m)
			return strings.NewReader(FormatURLParam(m)), nil
		} else {
			return nil, err
		}
	}
}

func readForm(c *httpConfig) reader {
	return func(v any) (io.Reader, error) {
		if bs, err := json.Marshal(v); err == nil {
			m := make(map[string]any)
			_ = json.Unmarshal(bs, &m)
			return strings.NewReader(FormatURLParam(m)), nil
		} else {
			return nil, err
		}
	}
}

func readFile(c *httpConfig) reader {
	return func(v any) (io.Reader, error) {
		var (
			body        io.Reader
			fileContent *multipart.Writer
		)
		if bs, err := json.Marshal(v); err == nil {
			m := make(map[string]any)
			_ = json.Unmarshal(bs, &m)
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
					_ = fileContent.WriteField(k, xutils.Any2String(v))
				}
			}
			_ = fileContent.Close()
			c.headers.Set("Content-Type", fileContent.FormDataContentType())
			return body, nil
		} else {
			return nil, err
		}
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
			v = xutils.Any2String(body[k])
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
