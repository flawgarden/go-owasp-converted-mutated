package controllers

import (
	"net/http"
	"net/url"
)

type BenchmarkTest00276 struct{}

func (b *BenchmarkTest00276) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	var param string

	headers := r.Header["Referer"]

	if len(headers) > 0 {
		param = headers[0]
	}

	param, _ = url.QueryUnescape(param)

	var bar string
	num := 106

	if (7*42)-num > 200 {
		bar = "This should never happen"
	} else {
		bar = param
	}

	w.Header().Set("X-XSS-Protection", "0")
	w.Write([]byte(bar))
}

func main() {
	http.Handle("/xss-00/BenchmarkTest00276", &BenchmarkTest00276{})
	http.ListenAndServe(":8080", nil)
}
