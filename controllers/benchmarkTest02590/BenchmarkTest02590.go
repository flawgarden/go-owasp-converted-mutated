package controllers

import (
	"net/http"
	"strings"
)

type BenchmarkTest02590 struct{}

func (b *BenchmarkTest02590) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	queryString := r.URL.RawQuery
	paramval := "BenchmarkTest02590="
	paramLoc := strings.Index(queryString, paramval)
	if paramLoc == -1 {
		w.Write([]byte("getQueryString() couldn't find expected parameter '" + "BenchmarkTest02590" + "' in query string."))
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := strings.Index(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : paramLoc+ampersandLoc]
	}
	param = strings.ReplaceAll(param, "%20", " ")

	bar := doSomething(param)

	w.Header().Set("X-XSS-Protection", "0")
	obj := []interface{}{"a", "b"}
	response := formatResponse(bar, obj)
	w.Write([]byte(response))
}

func doSomething(param string) string {
	bar := "safe!"
	map35520 := make(map[string]interface{})
	map35520["keyA-35520"] = "a_Value"
	map35520["keyB-35520"] = param
	map35520["keyC"] = "another_Value"
	bar = map35520["keyB-35520"].(string)
	bar = map35520["keyA-35520"].(string)
	return bar
}

func formatResponse(bar string, obj []interface{}) string {
	return bar // здесь можно добавить форматирование, если нужно
}
