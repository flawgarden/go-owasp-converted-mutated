package controllers

import (
	"fmt"
	"html"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type BenchmarkTest01646 struct{}

func (b *BenchmarkTest01646) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	queryString := r.URL.RawQuery
	paramval := "BenchmarkTest01646="
	paramLoc := strings.Index(queryString, paramval)
	if paramLoc == -1 {
		http.Error(w, fmt.Sprintf("getQueryString() couldn't find expected parameter '%s' in query string.", "BenchmarkTest01646"), http.StatusBadRequest)
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := strings.Index(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : paramLoc+ampersandLoc]
	}
	param, _ = url.QueryUnescape(param)

	bar := new(Test).doSomething(r, param)

	var fileName string
	var fos *os.File

	fileName = fmt.Sprintf("/path/to/directory/%s", bar)

	fos, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Couldn't open FileOutputStream on file: '" + fileName + "'")
		return
	}
	defer fos.Close()

	fmt.Fprintf(w, "Now ready to write to file: %s", html.EscapeString(fileName))
}

type Test struct{}

func (t *Test) doSomething(r *http.Request, param string) string {
	bar := "safe!"
	map84260 := make(map[string]interface{})
	map84260["keyA-84260"] = "a_Value"
	map84260["keyB-84260"] = param
	map84260["keyC"] = "another_Value"
	bar = map84260["keyB-84260"].(string)
	bar = map84260["keyA-84260"].(string)

	return bar
}
