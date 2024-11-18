package controllers

import (
	"fmt"
	"net/http"
)

type BenchmarkTest01282 struct{}

func (b *BenchmarkTest01282) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest01282")
	if param == "" {
		param = ""
	}

	bar := b.doSomething(param)

	cookie := http.Cookie{
		Name:     "SomeCookie",
		Value:    bar,
		Secure:   false,
		HttpOnly: true,
		Path:     r.RequestURI,
	}
	http.SetCookie(w, &cookie)

	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	fmt.Fprintf(w, "Created cookie: 'SomeCookie': with value: '%s' and secure flag set to: false", bar)
}

func (b *BenchmarkTest01282) doSomething(param string) string {
	bar := param
	if len(param) > 1 {
		bar = param[:len(param)-1]
	}
	return bar
}

func main() {
	http.Handle("/securecookie-00/BenchmarkTest01282", &BenchmarkTest01282{})
	http.ListenAndServe(":8080", nil)
}
