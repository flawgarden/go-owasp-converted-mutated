package controllers

import (
	"fmt"
	"net/http"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01281 struct{}

func (b *BenchmarkTest01281) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest01281")
	if param == "" {
		param = ""
	}

	bar := b.doSomething(param)

	cookie := http.Cookie{
		Name:     "SomeCookie",
		Value:    bar,
		Secure:   false,
		HttpOnly: true,
		Path:     r.URL.Path,
	}
	http.SetCookie(w, &cookie)

	response := fmt.Sprintf("Created cookie: 'SomeCookie': with value: '%s' and secure flag set to: false", bar)
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	w.Write([]byte(response))
}

func (b *BenchmarkTest01281) doSomething(param string) string {
	bar := ""
	if param != "" {
		bar = param
	}
	return bar
}
