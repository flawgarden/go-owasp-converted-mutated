package controllers

import (
	"fmt"
	"html"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

type BenchmarkTest02566 struct{}

func (b *BenchmarkTest02566) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.RawQuery
	paramval := "BenchmarkTest02566="
	paramLoc := strings.Index(queryString, paramval)
	if paramLoc == -1 {
		http.Error(w, "getQueryString() couldn't find expected parameter 'BenchmarkTest02566' in query string.", http.StatusBadRequest)
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := strings.Index(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : paramLoc+ampersandLoc]
	}

	param, _ = url.QueryUnescape(param)
	bar := doSomething(param)
	fileName := filepath.Join("testfiles", bar)

	fos, err := os.Create(fileName)
	if err != nil {
		http.Error(w, fmt.Sprintf("Couldn't open FileOutputStream on file: '%s'", fileName), http.StatusInternalServerError)
		return
	}
	defer fos.Close()

	fmt.Fprintf(w, "Now ready to write to file: %s", html.EscapeString(fileName))
}

func doSomething(param string) string {
	var bar string
	guess := "ABC"
	switchTarget := guess[1] // condition 'B', which is safe

	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bob"
	case 'C', 'D':
		bar = param
	default:
		bar = "bob's your uncle"
	}

	return bar
}
