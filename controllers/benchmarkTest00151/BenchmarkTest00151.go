package controllers

import (
	"net/http"
	"net/url"
	"strings"
)

type BenchmarkTest00151 struct{}

func (b *BenchmarkTest00151) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	if r.Header.Get("Referer") != "" {
		param = r.Header.Get("Referer")
	}

	param, _ = url.QueryUnescape(param)

	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the last 'safe' value
	}

	w.Header().Set("X-XSS-Protection", "0")
	w.Write([]byte(strings.Replace(bar, "%s", " ", -1)))
}
