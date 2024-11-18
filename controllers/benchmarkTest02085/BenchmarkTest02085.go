package controllers

import (
	"html"
	"net/http"
	"net/url"
)

type BenchmarkTest02085 struct{}

func (b *BenchmarkTest02085) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	var param string
	if headers := r.Header["BenchmarkTest02085"]; len(headers) > 0 {
		param = headers[0]
	}

	param, _ = url.QueryUnescape(param)

	bar := doSomething(r, param)

	r.Header.Set("userid", bar)

	w.Write([]byte("Item: 'userid' with value: '" + encodeForHTML(bar) + "' saved in session."))
}

func doSomething(r *http.Request, param string) string {
	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = append(valuesList[1:], valuesList[2:]...) // remove the 1st safe value
		bar = valuesList[1]                                    // get the last 'safe' value
	}
	return bar
}

func encodeForHTML(value string) string {
	return html.EscapeString(value)
}
