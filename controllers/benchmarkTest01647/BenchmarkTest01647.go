package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

type BenchmarkTest01647 struct{}

func (b *BenchmarkTest01647) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.RawQuery
	paramval := "BenchmarkTest01647="
	paramLoc := -1
	if queryString != "" {
		paramLoc = findParameter(queryString, paramval)
	}
	if paramLoc == -1 {
		http.Error(w, fmt.Sprintf("getQueryString() couldn't find expected parameter '%s' in query string.", "BenchmarkTest01647"), http.StatusBadRequest)
		return
	}

	param := extractParameter(queryString, paramLoc, paramval)

	bar := new(Test).doSomething(r, param)

	fileName := ""
	var fos *os.File

	defer func() {
		if fos != nil {
			fos.Close()
		}
	}()

	fileName = "/path/to/directory/" + bar
	fos, _ = os.Create(fileName)
	if fos != nil {
		w.Write([]byte(fmt.Sprintf("Now ready to write to file: %s", htmlEscape(fileName))))
	}
}

func findParameter(queryString string, paramval string) int {
	return indexOf(queryString, paramval)
}

func extractParameter(queryString string, paramLoc int, paramval string) string {
	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := indexOf(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = param[:ampersandLoc]
	}
	param, _ = url.QueryUnescape(param)
	return param
}

func indexOf(str string, substr string) int {
	return -1 // Simplified for example, implement actual logic
}

func htmlEscape(s string) string {
	return s // Simplified for example, implement actual escaping
}

type Test struct{}

func (t *Test) doSomething(r *http.Request, param string) string {
	num := 106
	if (7*42)-num > 200 {
		return "This should never happen"
	}
	return param
}
