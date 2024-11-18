package controllers

import (
	"fmt"
	"net/http"
)

type BenchmarkTest02314 struct{}

func (b *BenchmarkTest02314) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		b.doPost(w, r)
	} else if r.Method == http.MethodPost {
		b.doPost(w, r)
	}
}

func (b *BenchmarkTest02314) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := ""
	flag := true

	for name, values := range r.URL.Query() {
		if !flag {
			break
		}
		for _, value := range values {
			if value == "BenchmarkTest02314" {
				param = name
				flag = false
				break
			}
		}
	}

	bar := doSomething(param)

	w.Header().Set("X-XSS-Protection", "0")
	obj := []interface{}{"a", "b"}
	_, _ = w.Write([]byte(fmt.Sprintf(bar, obj...)))
}

func doSomething(param string) string {
	num := 106
	if (7*42)-num > 200 {
		return "This should never happen"
	}
	return param
}
