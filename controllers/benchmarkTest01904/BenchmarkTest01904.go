package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

type BenchmarkTest01904 struct{}

func (b *BenchmarkTest01904) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.Header.Get("BenchmarkTest01904")
	param, _ = url.QueryUnescape(param)

	bar := doSomething(param)

	fileTarget := bar
	w.Write([]byte(fmt.Sprintf("Access to file: '%s' created.", fileTarget)))

	if _, err := os.Stat(fileTarget); os.IsNotExist(err) {
		w.Write([]byte(" But file doesn't exist yet."))
	} else {
		w.Write([]byte(" And file already exists."))
	}
}

func doSomething(param string) string {
	bar := "safe!"
	values := map[string]interface{}{
		"keyA-13742": "a-Value",
		"keyB-13742": param,
		"keyC":       "another-Value",
	}
	bar = values["keyB-13742"].(string)
	return bar
}
