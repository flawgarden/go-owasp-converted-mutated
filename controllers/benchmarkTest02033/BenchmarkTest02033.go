package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

type BenchmarkTest02033 struct{}

func (b *BenchmarkTest02033) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	var param string

	if headers := r.Header["BenchmarkTest02033"]; len(headers) > 0 {
		param = headers[0]
	}

	param, _ = url.QueryUnescape(param)

	bar := doSomething(param)

	fileName := fmt.Sprintf("path/to/test/files/%s", bar)
	var fos *os.File

	defer func() {
		if fos != nil {
			fos.Close()
		}
	}()

	fos, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Fprintf(w, "Couldn't open FileOutputStream on file: '%s'", fileName)
		return
	}

	fmt.Fprintf(w, "Now ready to write to file: %s", fileName)
}

func doSomething(param string) string {
	bar := ""
	guess := "ABC"
	switchTarget := guess[1]

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
