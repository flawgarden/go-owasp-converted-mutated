package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

type BenchmarkTest01027 struct{}

func (b *BenchmarkTest01027) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.Header.Get("BenchmarkTest01027")
	param, _ = url.QueryUnescape(param)

	bar := b.doSomething(r, param)

	fileTarget := bar
	w.Write([]byte(fmt.Sprintf("Access to file: '%s' created.", fileTarget)))

	if _, err := os.Stat(fileTarget); os.IsNotExist(err) {
		w.Write([]byte(" But file doesn't exist yet."))
	} else {
		w.Write([]byte(" And file already exists."))
	}
}

func (b *BenchmarkTest01027) doSomething(r *http.Request, param string) string {
	bar := ""
	num := 106
	if (7*18)+num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}
	return bar
}

func main() {
	http.Handle("/pathtraver-01/BenchmarkTest01027", &BenchmarkTest01027{})
	http.ListenAndServe(":8080", nil)
}
