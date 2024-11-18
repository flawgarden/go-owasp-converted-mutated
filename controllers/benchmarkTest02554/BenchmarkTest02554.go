package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

type BenchmarkTest02554 struct{}

func (b *BenchmarkTest02554) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.RawQuery
	paramval := "BenchmarkTest02554="
	paramLoc := strings.Index(queryString, paramval)
	if paramLoc == -1 {
		http.Error(w, fmt.Sprintf("getQueryString() couldn't find expected parameter '%s' in query string.", "BenchmarkTest02554"), http.StatusBadRequest)
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := strings.Index(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : paramLoc+ampersandLoc]
	}

	bar := doSomething(param)

	fileTarget := fmt.Sprintf("%s/%s", os.Getenv("TESTFILES_DIR"), bar)
	w.Write([]byte(fmt.Sprintf("Access to file: '%s' created.", fileTarget)))
	if _, err := os.Stat(fileTarget); err == nil {
		w.Write([]byte(" And file already exists."))
	} else {
		w.Write([]byte(" But file doesn't exist yet."))
	}
}

func doSomething(param string) string {
	var bar string
	num := 106
	if (7*18)+num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}
	return bar
}

func main() {
	http.Handle("/pathtraver-03/BenchmarkTest02554", &BenchmarkTest02554{})
	http.ListenAndServe(":8080", nil)
}
