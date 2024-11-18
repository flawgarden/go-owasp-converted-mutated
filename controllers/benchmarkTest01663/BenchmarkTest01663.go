package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type BenchmarkTest01663 struct{}

func (b *BenchmarkTest01663) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	queryString := r.URL.RawQuery
	paramval := "BenchmarkTest01663="
	paramLoc := strings.Index(queryString, paramval)
	if paramLoc == -1 {
		http.Error(w, fmt.Sprintf("getQueryString() couldn't find expected parameter '%s' in query string.", "BenchmarkTest01663"), http.StatusBadRequest)
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := strings.Index(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : paramLoc+ampersandLoc]
	}
	param, _ = url.QueryUnescape(param)

	bar := new(Test).doSomething(r, param)

	w.Header().Set("X-XSS-Protection", "0")
	_, _ = fmt.Fprintf(w, bar, "a", "b")
}

type Test struct{}

func (t *Test) doSomething(r *http.Request, param string) string {
	bar := "safe!"
	mapVar := map[string]interface{}{
		"keyA-36601": "a_Value",
		"keyB-36601": param,
		"keyC":       "another_Value",
	}
	bar = mapVar["keyB-36601"].(string)
	bar = mapVar["keyA-36601"].(string)
	return bar
}

func main() {
	http.Handle("/xss-03/BenchmarkTest01663", &BenchmarkTest01663{})
	http.ListenAndServe(":8080", nil)
}
