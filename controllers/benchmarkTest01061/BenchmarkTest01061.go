package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type BenchmarkTest01061 struct{}

func (b *BenchmarkTest01061) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		b.doPost(w, r)
		return
	}
	if r.Method == http.MethodPost {
		b.doPost(w, r)
		return
	}
	http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
}

func (b *BenchmarkTest01061) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.Header.Get("BenchmarkTest01061")
	param = strings.TrimSpace(param)

	bar := new(Test).doSomething(param)

	cookie := http.Cookie{
		Name:     "SomeCookie",
		Value:    bar,
		Secure:   false,
		HttpOnly: true,
		Path:     r.URL.Path,
	}
	http.SetCookie(w, &cookie)

	fmt.Fprintf(w, "Created cookie: 'SomeCookie': with value: '%s' and secure flag set to: false", htmlEscape(bar))
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	bar := param
	if len(param) > 1 {
		sbxyz37461 := strings.Builder{}
		sbxyz37461.WriteString(param)
		bar = sbxyz37461.String()[:len(param)-1] + "Z"
	}
	return bar
}

func htmlEscape(str string) string {
	escaped, _ := json.Marshal(str)
	return string(escaped[1 : len(escaped)-1])
}

func main() {
	http.Handle("/", &BenchmarkTest01061{})
	http.ListenAndServe(":8080", nil)
}
