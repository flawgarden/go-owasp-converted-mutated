package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

type BenchmarkTest01570 struct{}

func (b *BenchmarkTest01570) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	r.ParseForm()

	param := r.FormValue("BenchmarkTest01570")
	bar := b.doSomething(param)

	fileTarget := filepath.Join("/path/to/testfiles", bar)
	w.Write([]byte(fmt.Sprintf("Access to file: '%s' created.", fileTarget)))
	if _, err := os.Stat(fileTarget); err == nil {
		w.Write([]byte(" And file already exists."))
	} else {
		w.Write([]byte(" But file doesn't exist yet."))
	}
}

func (b *BenchmarkTest01570) doSomething(param string) string {
	num := 106
	if (7*18)+num > 200 {
		return "This_should_always_happen"
	}
	return param
}

func main() {
	http.Handle("/pathtraver-01/BenchmarkTest01570", &BenchmarkTest01570{})
	http.ListenAndServe(":8080", nil)
}
