package controllers

import (
	"net/http"
)

type BenchmarkTest00555 struct{}

func (bt *BenchmarkTest00555) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	var param string
	flag := true
	names := r.URL.Query()

	for name, values := range names {
		if !flag {
			break
		}
		for _, value := range values {
			if value == "BenchmarkTest00555" {
				param = name
				flag = false
				break
			}
		}
	}

	bar := ""
	num := 106

	if (7*42)-num > 200 {
		bar = "This should never happen"
	} else {
		bar = param
	}

	w.Header().Set("X-XSS-Protection", "0")
	w.Write([]byte(bar))
}
