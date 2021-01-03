package helpers

import (
	"net/http"
	"reflect"
	"unicode"

	"github.com/gorilla/sessions"
	viper "github.com/spf13/viper"
)

// MaxDepth ...
const MaxDepth = 32

// MergeMap ...
func MergeMap(dst, src map[string]interface{}) map[string]interface{} {
	return merge(dst, src, 0)
}

func merge(dst, src map[string]interface{}, depth int) map[string]interface{} {
	if depth > MaxDepth {
		panic("too deep!")
	}
	for key, srcVal := range src {
		if dstVal, ok := dst[key]; ok {
			srcMap, srcMapOk := mapify(srcVal)
			dstMap, dstMapOk := mapify(dstVal)
			if srcMapOk && dstMapOk {
				srcVal = merge(dstMap, srcMap, depth+1)
			}
		}
		dst[key] = srcVal
	}
	return dst
}

// UcFirst ..
func UcFirst(s string) string {
	for index, value := range s {
		return string(unicode.ToUpper(value)) + s[index+1:]
	}
	return ""
}

// LcFirst ..
func LcFirst(s string) string {
	for index, value := range s {
		return string(unicode.ToLower(value)) + s[index+1:]
	}
	return ""
}

// StoreSessionData ..
func StoreSessionData(request *http.Request, writer http.ResponseWriter, nameSession string, data interface{}) {
	sessionStore := sessions.NewCookieStore([]byte(viper.GetString("app.key")))
	sessionNow, err := sessionStore.Get(request, nameSession)
	if err != nil {
		panic(err)
	}

	sessionNow.Values["data"] = data
	sessionNow.Save(request, writer)
}

// GetSessionDataString ..
func GetSessionDataString(request *http.Request, writer http.ResponseWriter, nameSession string) string {
	sessionStore := sessions.NewCookieStore([]byte(viper.GetString("app.key")))
	sessionNow, err := sessionStore.Get(request, nameSession)
	if err != nil {
		return ""
	}

	return sessionNow.Values["data"].(string)
}

// GetSessionDataString ..
func GetSessionData(request *http.Request, writer http.ResponseWriter, nameSession string) interface{} {
	sessionStore := sessions.NewCookieStore([]byte(viper.GetString("app.key")))
	sessionNow, err := sessionStore.Get(request, nameSession)
	if err != nil {
		return ""
	}

	return sessionNow.Values["data"]
}

func mapify(i interface{}) (map[string]interface{}, bool) {
	value := reflect.ValueOf(i)
	if value.Kind() == reflect.Map {
		m := map[string]interface{}{}
		for _, k := range value.MapKeys() {
			m[k.String()] = value.MapIndex(k).Interface()
		}
		return m, true
	}
	return map[string]interface{}{}, false
}
