package controllers

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type BenchmarkTest02559 struct{}

func (b *BenchmarkTest02559) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.RawQuery
	paramval := "BenchmarkTest02559="
	paramLoc := strings.Index(queryString, paramval)
	if paramLoc == -1 {
		http.Error(w, "getQueryString() couldn't find expected parameter 'BenchmarkTest02559' in query string.", http.StatusBadRequest)
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := strings.Index(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : paramLoc+ampersandLoc]
	}

	bar := doSomething(param)

	startURIslashes := ""
	if strings.Contains(os.Getenv("OS"), "Windows") {
		startURIslashes = "/"
	} else {
		startURIslashes = "//"
	}

	fileURI := "file:" + startURIslashes + filepath.Join("path/to/testfiles", strings.ReplaceAll(bar, " ", "_"))
	fileTarget := filepath.Clean(fileURI)

	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	w.Write([]byte("Access to file: '" + fileTarget + "' created."))
	if _, err := os.Stat(fileTarget); os.IsNotExist(err) {
		w.Write([]byte(" But file doesn't exist yet."))
	} else {
		w.Write([]byte(" And file already exists."))
	}
}

func doSomething(param string) string {
	return param // Placeholder for actual logic
}
