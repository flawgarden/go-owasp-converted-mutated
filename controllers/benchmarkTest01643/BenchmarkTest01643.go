package controllers

import (
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

type BenchmarkTest01643 struct{}

func (bt *BenchmarkTest01643) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.RawQuery
	paramval := "BenchmarkTest01643="
	paramLoc := -1

	if queryString != "" {
		paramLoc = findParamLoc(queryString, paramval)
	}

	if paramLoc == -1 {
		http.Error(w, "getQueryString() couldn't find expected parameter 'BenchmarkTest01643' in query string.", http.StatusBadRequest)
		return
	}

	param := extractParam(queryString, paramLoc, paramval)
	param, _ = url.QueryUnescape(param)

	bar := new(Test).doSomething(r, param)

	startURIslashes := ""
	if os.PathSeparator == '\\' {
		startURIslashes = "/"
	} else {
		startURIslashes = "//"
	}

	fileURI := "file:" + startURIslashes + filepath.Join("testfiles", bar)
	fileTarget := filepath.FromSlash(fileURI)

	w.Write([]byte("Access to file: '" + fileTarget + "' created."))
	if _, err := os.Stat(fileTarget); err == nil {
		w.Write([]byte(" And file already exists."))
	} else {
		w.Write([]byte(" But file doesn't exist yet."))
	}
}

func findParamLoc(queryString, paramval string) int {
	return stringIndex(queryString, paramval)
}

func stringIndex(s, substr string) int {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}

func extractParam(queryString string, paramLoc int, paramval string) string {
	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := stringIndex(queryString[paramLoc:], "&")

	if ampersandLoc != -1 {
		param = param[:ampersandLoc]
	}

	return param
}

type Test struct{}

func (t *Test) doSomething(r *http.Request, param string) string {
	num := 106
	if (7*42)-num > 200 {
		return "This should never happen"
	}
	return param
}
