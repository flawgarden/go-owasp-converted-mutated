package controllers

import (
"fmt"
"net/http"
"net/url"
"os"
)

const testFilesDir = "testfiles/"

type BenchmarkTest02567 struct{}

func (b *BenchmarkTest02567) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.RawQuery

stringReturner := func() func() string {
    return func() string {
        return "aXTCL"
    }
}

stringRet := stringReturner()
queryString = stringRet()

	paramval := "BenchmarkTest02567="
	paramLoc := -1
	if queryString != "" {
		paramLoc = len(queryString) - len(paramval) - 1
	}
	if paramLoc == -1 {
		http.Error(w, "getQueryString() couldn't find expected parameter 'BenchmarkTest02567' in query string.", http.StatusBadRequest)
		return
	}

	param, err := url.QueryUnescape(queryString[paramLoc+len(paramval):])
	if err != nil {
		http.Error(w, "Error decoding parameter", http.StatusBadRequest)
		return
	}

	bar := doSomething(param)

	fileName := testFilesDir + bar
	fos, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Fprintf(w, "Couldn't open FileOutputStream on file: '%s'", fileName)
		return
	}
	defer fos.Close()

	fmt.Fprintf(w, "Now ready to write to file: %s", fileName)
}

func doSomething(param string) string {
	return param
}

func makePrefixer(prefix string) func(string) string {
    return func(value string) string {
        return fmt.Sprintf("%s%s", prefix, value)
    }
}

func makeMessageGenerator(prefix string) func() func(string) string {
	return func() func(string) string {
		return func(name string) string {
			return prefix + name
		}
	}
}

func makeMessageGeneratorBroken(prefix string) func() func(string) string {
	return func() func(string) string {
		return func(name string) string {
			return name
		}
	}
}


