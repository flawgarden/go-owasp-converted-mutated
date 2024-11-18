package controllers

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type BenchmarkTest01062 struct{}

func (b *BenchmarkTest01062) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		b.doPost(w, r)
		return
	}
	b.doPost(w, r)
}

func (b *BenchmarkTest01062) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.Header.Get("BenchmarkTest01062")
	param, _ = url.QueryUnescape(param)

	bar := b.test().doSomething(r, param)

	var str string
	if bar == "" {
		str = "No cookie value supplied"
	} else {
		str = bar
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "SomeCookie",
		Value:    str,
		Secure:   true,
		HttpOnly: true,
		Path:     r.RequestURI,
	})

	_, _ = w.Write([]byte("Created cookie: 'SomeCookie': with value: '" + str + "' and secure flag set to: true"))
}

func (b *BenchmarkTest01062) test() *Test {
	return &Test{}
}

type Test struct{}

func (t *Test) doSomething(r *http.Request, param string) string {
	if param != "" {
		decoded, _ := json.Marshal(param)
		return string(decoded)
	}
	return ""
}
