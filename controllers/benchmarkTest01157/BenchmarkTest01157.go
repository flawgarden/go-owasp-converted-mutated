package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

type BenchmarkTest01157 struct{}

func (b *BenchmarkTest01157) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := ""

	if headers := r.Header["BenchmarkTest01157"]; len(headers) > 0 {
		param = headers[0]
	}

	param, _ = url.QueryUnescape(param)
	bar := new(Test).doSomething(r, param)

	startURIslashes := ""
	if os.Getenv("OS") == "Windows_NT" {
		startURIslashes = "/"
	} else {
		startURIslashes = "//"
	}

	fileURI, err := url.Parse(fmt.Sprintf("file:%s%s", startURIslashes, filepath.Clean("/path/to/testfiles/"+bar)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fileTarget := fileURI.Path
	fmt.Fprintf(w, "Access to file: '%s' created.", fileTarget)

	if _, err := os.Stat(fileTarget); err == nil {
		fmt.Fprintln(w, " And file already exists.")
	} else {
		fmt.Fprintln(w, " But file doesn't exist yet.")
	}
}

type Test struct{}

func (t *Test) doSomething(r *http.Request, param string) string {
	var bar string
	guess := "ABC"
	switchTarget := guess[2]

	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bobs_your_uncle"
	case 'C', 'D':
		bar = param
	default:
		bar = "bobs_your_uncle"
	}

	return bar
}
