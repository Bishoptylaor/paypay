package zutils

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"time"
	"unicode/utf8"
	"unsafe"
)

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
 @Time    : 2024/7/13 -- 15:00
 @Author  : bishop ❤️ MONEY
 @Description: zstring.go
*/

var (
	sbc2dbcMap = map[rune]rune{
		'＋': '+',
		'－': '-',
		'０': '0',
		'１': '1',
		'２': '2',
		'３': '3',
		'４': '4',
		'５': '5',
		'６': '6',
		'７': '7',
		'８': '8',
		'９': '9',
		'‘': '\'',
		'’': '\'',
		'“': '"',
		'”': '"',
		'，': ',',
		'。': '.',
		'？': '?',
		'×': '*',
		'／': '/',
		'％': '%',
		'＃': '#',
		'＠': '@'}
)

// UCFirst 首写字母大写
func UCFirst(str string) string {
	runes := []rune(str)
	if len(runes) < 1 {
		return str
	}
	if runes[0] >= 97 && runes[0] <= 122 {
		runes[0] -= 32
	}
	return string(runes)
}

func SBC2DBC(str string) string {
	runes := []rune(str)
	var buf bytes.Buffer
	for i := 0; i < len(runes); i++ {
		r, ok := sbc2dbcMap[runes[i]]
		if ok {
			buf.WriteRune(r)
		} else {
			buf.WriteRune(runes[i])
		}
	}
	return buf.String()
}

func Concat(strings ...string) string {
	var buffer bytes.Buffer
	for _, s := range strings {
		buffer.WriteString(s)
	}
	return buffer.String()
}

func StrLen(str string) int {
	runes := []rune(str)
	return len(runes)
}

func Str2Bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]} // copy data, len, cap
	return *(*[]byte)(unsafe.Pointer(&h))
}

func Bytes2Str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func FilterEmoji(s string) string {
	dst := ""
	for _, value := range s {
		_, size := utf8.DecodeRuneInString(string(value))
		if size <= 3 {
			dst += string(value)
		}
	}
	return dst
}

func ConvertToString(v any) (str string) {
	if v == nil {
		return ""
	}
	var (
		bs  []byte
		err error
	)
	if bs, err = json.Marshal(v); err != nil {
		return ""
	}
	str = string(bs)
	return
}

func RandomString(l int) string {
	str := "0123456789AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz"
	_bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, _bytes[r.Intn(len(_bytes))])
	}
	return string(result)
}
