package controllers

import (
	"fmt"
	"net/http"
	"net/url"
)

type BenchmarkTest01187 struct{}

func (b *BenchmarkTest01187) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.Header.Get("BenchmarkTest01187")
	param, _ = url.QueryUnescape(param)

	bar := b.doSomething(r, param)

	cookie := http.Cookie{
		Name:     "SomeCookie",
		Value:    bar,
		HttpOnly: true,
		Path:     r.URL.Path,
	}
	http.SetCookie(w, &cookie)

	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	fmt.Fprintf(w, "Created cookie: 'SomeCookie': with value: '%s' and secure flag set to: false", bar)
}

func (b *BenchmarkTest01187) doSomething(r *http.Request, param string) string {
	return sanitize(param)
}

func sanitize(input string) string {
	return input // Здесь должна быть реализация экранирования
}
