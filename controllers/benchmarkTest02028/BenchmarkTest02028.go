package controllers

import (
	"fmt"
	"html"
	"net/http"
	"net/url"
	"os"
)

type BenchmarkTest02028 struct{}

func (b *BenchmarkTest02028) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	var param string
	if headerValues := r.Header["BenchmarkTest02028"]; len(headerValues) > 0 {
		param = headerValues[0]
	}

	param, _ = url.QueryUnescape(param)
	bar := b.doSomething(r, param)

	fileTarget := fmt.Sprintf("./testfiles/%s", bar)
	w.Write([]byte(fmt.Sprintf("Access to file: '%s' created.", htmlEscape(fileTarget))))
	if _, err := os.Stat(fileTarget); err == nil {
		w.Write([]byte(" And file already exists."))
	} else {
		w.Write([]byte(" But file doesn't exist yet."))
	}
}

func (b *BenchmarkTest02028) doSomething(r *http.Request, param string) string {
	var bar string
	guess := "ABC"
	switchTarget := guess[1] // condition 'B', which is safe

	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bob"
	case 'C', 'D':
		bar = param
	default:
		bar = "bob's your uncle"
	}

	return bar
}

func htmlEscape(str string) string {
	return html.EscapeString(str)
}

func main() {
	http.Handle("/pathtraver-02/BenchmarkTest02028", &BenchmarkTest02028{})
	http.ListenAndServe(":8080", nil)
}
