package controllers

import (
	"net/http"
	"net/url"
	"strings"
)

type BenchmarkTest02065 struct{}

func (b *BenchmarkTest02065) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		b.doPost(w, r)
		return
	}
	if r.Method == http.MethodPost {
		b.doPost(w, r)
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func (b *BenchmarkTest02065) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	var param string
	if headers := r.Header["BenchmarkTest02065"]; len(headers) > 0 {
		param = headers[0]
	}

	param, _ = url.QueryUnescape(param)
	bar := doSomething(r, param)

	cookie := http.Cookie{
		Name:     "SomeCookie",
		Value:    bar,
		Secure:   true,
		HttpOnly: true,
		Path:     r.URL.Path,
	}
	http.SetCookie(w, &cookie)

	responseMessage := "Created cookie: 'SomeCookie': with value: '" + htmlEscape(bar) + "' and secure flag set to: true"
	w.Write([]byte(responseMessage))
}

func doSomething(r *http.Request, param string) string {
	num := 106
	if (7*18)+num > 200 {
		return "This_should_always_happen"
	}
	return param
}

func htmlEscape(str string) string {
	return strings.ReplaceAll(strings.ReplaceAll(str, "&", "&amp;"), "<", "&lt;")
}
