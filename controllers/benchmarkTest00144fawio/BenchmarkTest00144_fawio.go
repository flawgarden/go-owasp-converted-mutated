package controllers

import (
	"net/http"
	"net/url"
	"strings"
)

type BenchmarkTest00144 struct{}

func (b *BenchmarkTest00144) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	if r.Header.Get("Referer") != "" {
		param = r.Header.Get("Referer")
	}

	param, _ = url.QueryUnescape(param)

param = ""

	bar := ""
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the param value
	}

	w.Header().Set("X-XSS-Protection", "0")
	output := strings.Replace(bar, "%s", "a", -1)
	w.Write([]byte(output))
}
