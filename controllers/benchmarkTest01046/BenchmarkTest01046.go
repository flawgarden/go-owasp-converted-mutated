package controllers

import (
	"fmt"
	"net/http"
	"net/url"
)

type BenchmarkTest01046 struct{}

func (b *BenchmarkTest01046) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	if r.Header.Get("Referer") != "" {
		param = r.Header.Get("Referer")
	}

	decodedParam, _ := url.QueryUnescape(param)
	bar := new(Test).doSomething(r, decodedParam)

	w.Header().Set("X-XSS-Protection", "0")
	obj := []interface{}{"a", bar}
	output := "<!DOCTYPE html>\n<html>\n<body>\n<p>"
	output += fmt.Sprintf("Formatted like: %s and %s.", obj[0], obj[1])
	output += "\n</p>\n</body>\n</html>"
	w.Write([]byte(output))
}

type Test struct{}

func (t *Test) doSomething(r *http.Request, param string) string {
	bar := "safe!"
	mapData := make(map[string]interface{})
	mapData["keyA-95803"] = "a-Value"
	mapData["keyB-95803"] = param
	mapData["keyC"] = "another-Value"
	bar = mapData["keyB-95803"].(string)

	return bar
}
