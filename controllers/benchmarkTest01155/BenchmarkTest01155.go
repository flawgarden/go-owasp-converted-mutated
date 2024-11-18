package controllers

import (
	"html"
	"net/http"
	"net/url"
	"os"
)

type BenchmarkTest01155 struct{}

func (b *BenchmarkTest01155) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	headers := r.Header["BenchmarkTest01155"]
	if len(headers) > 0 {
		param = headers[0]
	}
	param, _ = url.QueryUnescape(param)

	bar := b.doSomething(r, param)

	fileTarget := bar
	w.Write([]byte("Access to file: '" + html.EscapeString(fileTarget) + "' created."))
	if _, err := os.Stat(fileTarget); err == nil {
		w.Write([]byte(" And file already exists."))
	} else {
		w.Write([]byte(" But file doesn't exist yet."))
	}
}

func (b *BenchmarkTest01155) doSomething(r *http.Request, param string) string {
	bar := ""
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the param value
	}
	return bar
}
