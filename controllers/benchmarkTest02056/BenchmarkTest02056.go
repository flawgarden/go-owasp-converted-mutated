package controllers

import (
	"net/http"
	"net/url"
)

type BenchmarkTest02056 struct{}

func (b *BenchmarkTest02056) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := ""

	referer := r.Header.Get("Referer")
	if referer != "" {
		param = referer
	}

	decodedParam, _ := url.QueryUnescape(param)
	bar := doSomething(r, decodedParam)

	w.Header().Set("X-XSS-Protection", "0")
	w.Write([]byte(bar))
}

func doSomething(r *http.Request, param string) string {
	bar := "safe!"
	map88820 := make(map[string]interface{})
	map88820["keyA-88820"] = "a-Value"
	map88820["keyB-88820"] = param
	map88820["keyC"] = "another-Value"
	bar = map88820["keyB-88820"].(string)

	return bar
}
