package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/Quieting/chowder/script/xerror"
)

var client = &http.Client{
	Timeout: 5 * time.Second, // 设置超时
}
var auth = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjgxMzcwMjIsImlhdCI6MTY2NTU0NTAyMiwidXNlcklkIjoxMDEyMjk1NX0.CfFi-RbWgYMdz5OezYgYZLeHyBMRW6tNZTqwDs6Jd0c"

func Post(url string, param interface{}, ret interface{}, authorization string) (err error) {
	if authorization != "" {
		SetAuth(authorization)
	}
	var p []byte
	if param == nil {
		p = []byte("{}")
	} else {
		p, err = json.Marshal(param)
		if err != nil {
			return
		}
	}

	respData, e := do(http.MethodPost, url, bytes.NewBuffer(p))
	if e != nil {
		return xerror.New(e, fmt.Sprintf("url:%s,  body:%+v, reply:%s\n", url, string(p), string(respData)))
	}

	if ret != nil {
		_ = json.Unmarshal(respData, ret)
	}

	return nil
}

// Get http GET 请求
// url:不包含参数
// param: path 请求参数
// ret: 返回数据映射
func Get(url string, param interface{}, ret interface{}, authorization string) (err error) {
	if authorization != "" {
		SetAuth(authorization)
	}

	if param != nil {
		vals := PathsValues(param)
		if len(vals) > 0 {
			url = url + "/?" + vals.Encode()
		}
	}

	reply, err := do(http.MethodGet, url, nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(reply, ret)
	if err != nil {
		return
	}

	return
}

// PathsValues 将 v 转换成get请求参数
// 仅支持结构体，不支持结构体嵌套
func PathsValues(v interface{}) url.Values {
	val := reflect.Indirect(reflect.ValueOf(v))
	if val.Kind() != reflect.Struct {
		return nil
	}
	typ := val.Type()

	values := make(url.Values, 0)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldKey := typ.Field(i).Tag.Get("form")
		fieldKey = strings.Split(fieldKey, ",")[0]
		fieldVal := ""
		switch field.Kind() {
		case reflect.String:
			fieldVal = field.String()
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			fieldVal = strconv.FormatInt(field.Int(), 10)
		default:

		}
		values.Set(fieldKey, fieldVal)
	}

	return values
}

func SetAuth(s string) {
	auth = s
}

func do(method, url string, body io.Reader) (data []byte, err error) {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return
	}
	request.Header.Set("Authorization", auth)
	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return
	}

	data, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	_ = response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, errors.New(string(data))
	}

	return
}
