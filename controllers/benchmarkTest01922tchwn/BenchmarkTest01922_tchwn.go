package controllers

import (
	"net/http"
	"net/url"
)

type BenchmarkTest01922 struct{}

func (b *BenchmarkTest01922) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := ""

	if referrer := r.Header.Get("Referer"); referrer != "" {
		param = referrer
	}

	decodedParam, _ := url.QueryUnescape(param)

tmpArrayUnique42 := []string{"", "", ""}
tmpArrayUnique42[0] = decodedParam
ah := NewArrayHolderWithValues(tmpArrayUnique42)
decodedParam = ah.Values[1]

	bar := doSomething(decodedParam)

	w.Header().Set("X-XSS-Protection", "0")
	_, _ = w.Write([]byte(bar))
}

func doSomething(param string) string {
	bar := ""
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the param value
	}
	return bar
}
