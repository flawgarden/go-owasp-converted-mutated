package controllers

import (
	"fmt"
	"net/http"
)

type BenchmarkTest01455 struct{}

func (b *BenchmarkTest01455) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		b.doPost(w, r)
	} else if r.Method == http.MethodPost {
		b.doPost(w, r)
	}
}

func (b *BenchmarkTest01455) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	var param string
	flag := true
	names := r.URL.Query()

	for name, values := range names {
		if flag {
			for _, value := range values {
				if value == "BenchmarkTest01455" {
					param = name
					flag = false
					break
				}
			}
		}
	}

	bar := b.doSomething(param)

	r.AddCookie(&http.Cookie{Name: "userid", Value: bar})

	fmt.Fprintf(w, "Item: 'userid' with value: '%s' saved in session.", bar)
}

func (b *BenchmarkTest01455) doSomething(param string) string {
	bar := param
	if len(param) > 1 {
		bar = param[:len(param)-1]
	}
	return bar
}
