package paypay

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/Bishoptylaor/paypay/pkg"
	"github.com/Bishoptylaor/paypay/pkg/zutils"
	"io"
	"net/url"
	"sort"
	"strings"
)

type Payload map[string]any

type xmlMapMarshal struct {
	XMLName xml.Name
	Value   any `xml:",cdata"`
}

type xmlMapUnmarshal struct {
	XMLName xml.Name
	Value   string `xml:",cdata"`
}

type File struct {
	Name    string `json:"name"`
	Content []byte `json:"content"`
}

// Set 设置参数
func (pl Payload) Set(key string, value any) Payload {
	pl[key] = value
	return pl
}

func (pl Payload) SetPayload(key string, value func(b Payload)) Payload {
	_pl := make(Payload)
	value(_pl)
	pl[key] = _pl
	return pl
}

// SetFormFile 设置 FormFile
func (pl Payload) SetFormFile(key string, file *File) Payload {
	pl[key] = file
	return pl
}

// Get 获取参数，同 GetString()
func (pl Payload) Get(key string) string {
	return pl.GetString(key)
}

// GetString 获取参数转换string
func (pl Payload) GetString(key string) string {
	if pl == nil {
		return pkg.NULL
	}
	value, ok := pl[key]
	if !ok {
		return pkg.NULL
	}
	v, ok := value.(string)
	if !ok {
		return zutils.ConvertToString(value)
	}
	return v
}

// GetInterface 获取原始参数
func (pl Payload) GetInterface(key string) any {
	if pl == nil {
		return nil
	}
	return pl[key]
}

// Remove 删除参数
func (pl Payload) Remove(key string) {
	delete(pl, key)
}

// Reset 置空Payload
func (pl Payload) Reset() {
	for k := range pl {
		delete(pl, k)
	}
}

func (pl Payload) JsonBody() (jb string) {
	bs, err := json.Marshal(pl)
	if err != nil {
		return ""
	}
	jb = string(bs)
	return jb
}

// Unmarshal to struct or slice point
func (pl Payload) Unmarshal(ptr any) (err error) {
	bs, err := json.Marshal(pl)
	if err != nil {
		return err
	}
	return json.Unmarshal(bs, ptr)
}

func (pl Payload) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	if len(pl) == 0 {
		return nil
	}
	start.Name = xml.Name{Space: pkg.NULL, Local: "xml"}
	if err = e.EncodeToken(start); err != nil {
		return
	}
	for k := range pl {
		if v := pl.GetString(k); v != pkg.NULL {
			_ = e.Encode(xmlMapMarshal{XMLName: xml.Name{Local: k}, Value: v})
		}
	}
	return e.EncodeToken(start.End())
}

func (pl Payload) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	for {
		var e xmlMapUnmarshal
		err = d.Decode(&e)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		pl.Set(e.XMLName.Local, e.Value)
	}
}

// EncodeWeChatSignParams ("bar=baz&foo=foo") sorted by key.
func (pl Payload) EncodeWeChatSignParams(apiKey string) string {
	if pl == nil {
		return pkg.NULL
	}
	var (
		buf     strings.Builder
		keyList []string
	)
	for k := range pl {
		keyList = append(keyList, k)
	}
	sort.Strings(keyList)
	for _, k := range keyList {
		if v := pl.GetString(k); v != pkg.NULL {
			buf.WriteString(k)
			buf.WriteByte('=')
			buf.WriteString(v)
			buf.WriteByte('&')
		}
	}
	buf.WriteString("key")
	buf.WriteByte('=')
	buf.WriteString(apiKey)
	return buf.String()
}

// EncodeAliPaySignParams ("bar=baz&foo=foo") sorted by key.
func (pl Payload) EncodeAliPaySignParams() string {
	if pl == nil {
		return pkg.NULL
	}
	var (
		buf     strings.Builder
		keyList []string
	)
	for k := range pl {
		keyList = append(keyList, k)
	}
	sort.Strings(keyList)
	for _, k := range keyList {
		if v := pl.GetString(k); v != pkg.NULL {
			buf.WriteString(k)
			buf.WriteByte('=')
			buf.WriteString(v)
			buf.WriteByte('&')
		}
	}
	if buf.Len() <= 0 {
		return pkg.NULL
	}
	return buf.String()[:buf.Len()-1]
}

// EncodeURLParams ("bar=baz&foo=foo") sorted by key.
func (pl Payload) EncodeURLParams() string {
	if pl == nil {
		return pkg.NULL
	}
	var (
		buf  strings.Builder
		keys []string
	)
	for k := range pl {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		if v := pl.GetString(k); v != pkg.NULL {
			buf.WriteString(url.QueryEscape(k))
			buf.WriteByte('=')
			buf.WriteString(url.QueryEscape(v))
			buf.WriteByte('&')
		}
	}
	if buf.Len() <= 0 {
		return pkg.NULL
	}
	return buf.String()[:buf.Len()-1]
}

func (pl Payload) IntegrityCheck(ctx context.Context, rulers ...Ruler) error {
	var ok bool
	var err error
	for _, ruler := range rulers {
		ok, err = zutils.Expr(ctx, ruler.Rule, pl)
		if !ok || err != nil {
			return fmt.Errorf("[IntegrityCheck]: rule:[%s], err[%s], [%s]", ruler.Des, err, ruler.Alert)
		}
	}
	return nil
}

type PayloadRuler func(caller string) []Ruler
type PayloadPreSetter func(pl Payload)
type Ruler struct {
	Des   string
	Alert string
	Rule  string
}

func NewRuler(des, rule, alert string) Ruler {
	return Ruler{
		Des:   des,
		Alert: alert,
		Rule:  rule,
	}
}

func PreSetter(key, value string) PayloadPreSetter {
	return func(pl Payload) {
		pl.Set(key, value)
	}
}
