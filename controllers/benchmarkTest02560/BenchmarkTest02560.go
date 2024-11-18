package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type BenchmarkTest02560Controller struct {
	http.Handler
}

func (c *BenchmarkTest02560Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	queryString := r.URL.RawQuery
	paramval := "BenchmarkTest02560="
	paramLoc := strings.Index(queryString, paramval)
	if paramLoc == -1 {
		http.Error(w, fmt.Sprintf("getQueryString() couldn't find expected parameter '%s' in query string.", "BenchmarkTest02560"), http.StatusBadRequest)
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := strings.Index(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : paramLoc+ampersandLoc]
	}
	param = r.URL.Query().Get("BenchmarkTest02560")

	bar := doSomething(param)

	fileName := filepath.Join("path/to/testfiles", bar)
	fis, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintf(w, "Couldn't open FileInputStream on file: '%s'", fileName)
		http.Error(w, fmt.Sprintf("Problem getting FileInputStream: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	defer fis.Close()

	b := make([]byte, 1000)
	size, err := fis.Read(b)
	if err != nil {
		fmt.Fprintf(w, "Error reading file: %s", err.Error())
		return
	}
	fmt.Fprintf(w, "The beginning of file: '%s' is:\n\n%s", fileName, string(b[:size]))
}

func doSomething(param string) string {
	bar := param
	num := 106
	if (7*42)-num > 200 {
		bar = "This should never happen"
	}
	return bar
}
