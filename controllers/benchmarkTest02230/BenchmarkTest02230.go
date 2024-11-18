package controllers

import (
	"fmt"
	"net/http"
	"strings"
)

type BenchmarkTest02230 struct{}

func (b *BenchmarkTest02230) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		b.doPost(w, r)
		return
	}
	b.doPost(w, r)
}

func (b *BenchmarkTest02230) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest02230")
	bar := doSomething(param)

	w.Header().Set("X-XSS-Protection", "0")
	output := fmt.Sprintf("Formatted like: %s and %s.", "a", bar)
	w.Write([]byte(output))
}

func doSomething(param string) string {
	bar := param
	if param != "" && len(param) > 1 {
		sbxyz51189 := strings.Builder{}
		sbxyz51189.WriteString(param)
		bar = sbxyz51189.String()[:len(param)-1] + "Z"
	}

	return bar
}

func main() {
	http.Handle("/xss-04/BenchmarkTest02230", &BenchmarkTest02230{})
	http.ListenAndServe(":8080", nil)
}
