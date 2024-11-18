package controllers

import (
	"net/http"
	"net/url"
	"strings"
)

type BenchmarkTest00290Controller struct {
	http.ResponseWriter
	*http.Request
}

func (c *BenchmarkTest00290Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	headers := r.Header["Referer"]

	if len(headers) > 0 {
		param = headers[0]
	}

	param, _ = url.QueryUnescape(param)

	sbxyz73567 := strings.Builder{}
	sbxyz73567.WriteString(param)
	bar := sbxyz73567.String() + "_SafeStuff"

	w.Header().Set("X-XSS-Protection", "0")
	w.Write([]byte(bar))
}
