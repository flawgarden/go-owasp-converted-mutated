package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

type BenchmarkTest02200 struct{}

func (b *BenchmarkTest02200) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		b.doPost(w, r)
	} else {
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
	}
}

func (b *BenchmarkTest02200) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := r.URL.Query().Get("BenchmarkTest02200")
	bar := doSomething(param)

	fileTarget := bar
	w.Write([]byte(fmt.Sprintf("Access to file: '%s' created.", htmlEscape(fileTarget))))
	if _, err := os.Stat(fileTarget); err == nil {
		w.Write([]byte(" And file already exists."))
	} else {
		w.Write([]byte(" But file doesn't exist yet."))
	}
}

func doSomething(param string) string {
	num := 106
	if (7*18)+num > 200 {
		return "This_should_always_happen"
	}
	return param
}

func htmlEscape(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, "&", "&amp;"), "<", "&lt;")
}

func main() {
	http.Handle("/pathtraver-02/BenchmarkTest02200", &BenchmarkTest02200{})
	http.ListenAndServe(":8080", nil)
}
