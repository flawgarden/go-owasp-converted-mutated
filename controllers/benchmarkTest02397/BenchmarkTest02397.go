package controllers

import (
	"encoding/json"
	"net/http"
)

type BenchmarkTest02397 struct{}

func (b *BenchmarkTest02397) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "0")
	r.ParseForm()
	param := r.FormValue("BenchmarkTest02397")
	if param == "" {
		param = ""
	}

	bar := doSomething(param)

	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	output := struct {
		Message string `json:"message"`
	}{
		Message: bar,
	}
	json.NewEncoder(w).Encode(output)
}

func doSomething(param string) string {
	bar := param
	if len(param) > 1 {
		bar = param[:len(param)-1]
	}
	return bar
}
