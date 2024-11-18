package controllers

import (
	"fmt"
	"net/http"
	"os"
)

type BenchmarkTest struct{}

func (b *BenchmarkTest) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := r.URL.Query().Get("BenchmarkTest01234")
	if param == "" {
		param = ""
	}

	bar := b.doSomething(param)

	fileTarget := fmt.Sprintf("path-to-testfiles-dir/%s", bar)
	w.Write([]byte(fmt.Sprintf("Access to file: '%s' created.", fileTarget)))
	if _, err := os.Stat(fileTarget); err == nil {
		w.Write([]byte(" And file already exists."))
	} else {
		w.Write([]byte(" But file doesn't exist yet."))
	}
}

func (b *BenchmarkTest) doSomething(param string) string {
	// Здесь можно подключить логику, аналогичную Test.doSomething из Java
	return param // или любая другая логика
}

func main() {
	http.Handle("/pathtraver-01/BenchmarkTest01234", &BenchmarkTest{})
	http.ListenAndServe(":8080", nil)
}
