package controllers

import (
"fmt"
"net/http"
"net/url"
"os"
"container/list"
)

const testFilesDir = "testfiles/"

type BenchmarkTest02567 struct{}

func (b *BenchmarkTest02567) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.RawQuery

queue787231 := list.New()
queue787231.PushBack("YyVxj")
queue787231.PushBack(queryString)
for e := queue787231.Front(); e != nil; {
    next := e.Next()
    if e.Value == "YyVxj" {
        queue787231.Remove(e)
    }
    e = next
}
queryString = queue787231.Front().Value.(string)

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
