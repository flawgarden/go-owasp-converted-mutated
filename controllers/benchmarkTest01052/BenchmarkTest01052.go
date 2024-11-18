package controllers

import (
	"fmt"
	"net/http"
	"net/url"
)

type BenchmarkTest01052 struct{}

func (bt *BenchmarkTest01052) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	if referer := r.Header.Get("Referer"); referer != "" {
		param = referer
	}

	param, _ = url.QueryUnescape(param)

	bar := new(Test).doSomething(r, param)

	w.Header().Set("X-XSS-Protection", "0")
	obj := []string{"a", "b"}
	_, _ = w.Write([]byte(fmt.Sprintf(bar, obj)))
}

type Test struct{}

func (t *Test) doSomething(r *http.Request, param string) string {
	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value

		bar = valuesList[0] // get the last 'safe' value
	}
	return bar
}
