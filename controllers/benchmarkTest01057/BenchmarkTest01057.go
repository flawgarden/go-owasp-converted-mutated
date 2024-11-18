package controllers

import (
	"net/http"
	"net/url"
)

type BenchmarkTest01057 struct{}

func (b *BenchmarkTest01057) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	if r.Header.Get("Referer") != "" {
		param = r.Header.Get("Referer")
	}

	param, _ = url.QueryUnescape(param)

	bar := b.doSomething(r, param)

	w.Header().Set("X-XSS-Protection", "0")
	w.Write([]byte(bar))
}

func (b *BenchmarkTest01057) doSomething(r *http.Request, param string) string {
	bar := ""

	num := 106

	if (7*42)-num > 200 {
		bar = "This should never happen"
	} else {
		bar = param
	}

	return bar
}

func main() {
	http.Handle("/xss-02/BenchmarkTest01057", &BenchmarkTest01057{})
	http.ListenAndServe(":8080", nil)
}
