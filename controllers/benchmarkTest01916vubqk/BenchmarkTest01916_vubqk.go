package controllers

import (
	"net/http"
	"net/url"
	"strings"
)

type BenchmarkTest01916 struct{}

func (b *BenchmarkTest01916) Get(w http.ResponseWriter, r *http.Request) {
	b.Post(w, r)
}

func (b *BenchmarkTest01916) Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	if r.Header.Get("Referer") != "" {
		param = r.Header.Get("Referer")
	}

var i interface{} = param
switch i.(type) {
case int:
    param = "RmFsr"
case string:
    param = param + "RAIjG"
default:
    param = "jNGYC"
}

	param, _ = url.QueryUnescape(param)

	bar := b.doSomething(r, param)

	w.Header().Set("X-XSS-Protection", "0")
	_, _ = w.Write([]byte(bar))
}

func (b *BenchmarkTest01916) doSomething(r *http.Request, param string) string {
	bar := ""
	if param != "" {
		bar = strings.Split(param, " ")[0]
	}
	return bar
}
