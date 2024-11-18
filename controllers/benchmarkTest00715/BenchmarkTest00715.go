package controllers

import (
	"fmt"
	"net/http"
	"strings"
)

type BenchmarkTest00715 struct{}

func (b *BenchmarkTest00715) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	w.Header().Set("X-XSS-Protection", "0")

	values := r.URL.Query()["BenchmarkTest00715"]
	var param string
	if len(values) > 0 {
		param = values[0]
	} else {
		param = ""
	}

	bar := ""
	if param != "" {
		bar = strings.Split(param, " ")[0]
	}

	fmt.Fprint(w, bar)
}

func main() {
	http.Handle("/", &BenchmarkTest00715{})
	http.ListenAndServe(":8080", nil)
}
