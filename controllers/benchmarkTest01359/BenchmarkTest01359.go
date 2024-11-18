package controllers

import (
	"fmt"
	"net/http"
	"strings"
)

type BenchmarkTest01359 struct{}

func (b *BenchmarkTest01359) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		b.doPost(w, r)
	} else if r.Method == http.MethodPost {
		b.doPost(w, r)
	}
}

func (b *BenchmarkTest01359) doPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest01359")
	bar := b.doSomething(r, param)

	cookie := http.Cookie{
		Name:     "SomeCookie",
		Value:    bar,
		Secure:   true,
		HttpOnly: true,
		Path:     r.RequestURI,
	}
	http.SetCookie(w, &cookie)

	response := fmt.Sprintf("Created cookie: 'SomeCookie': with value: '%s' and secure flag set to: true", bar)
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	w.Write([]byte(response))
}

func (b *BenchmarkTest01359) doSomething(r *http.Request, param string) string {
	return htmlEscape(param)
}

func htmlEscape(value string) string {
	return strings.ReplaceAll(strings.ReplaceAll(value, "&", "&amp;"), "<", "&lt;")
}

func main() {
	http.Handle("/securecookie-00/BenchmarkTest01359", &BenchmarkTest01359{})
	http.ListenAndServe(":8080", nil)
}
