package controllers

import (
	"fmt"
	"net/http"
	"net/url"
)

type BenchmarkTest00323 struct{}

func (b *BenchmarkTest00323) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	headers := r.Header["BenchmarkTest00323"]

	if len(headers) > 0 {
		param = headers[0]
	}

	param, _ = url.QueryUnescape(param)

	var bar string
	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}

	r.AddCookie(&http.Cookie{Name: "userid", Value: bar})

	fmt.Fprintf(w, "Item: 'userid' with value: '%s' saved in session.", bar)
}

func main() {
	http.Handle("/trustbound-00/BenchmarkTest00323", &BenchmarkTest00323{})
	http.ListenAndServe(":8080", nil)
}
