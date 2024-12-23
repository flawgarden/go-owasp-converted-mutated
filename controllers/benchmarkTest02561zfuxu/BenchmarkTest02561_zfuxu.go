package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

type BenchmarkTest02561 struct{}

func (bt *BenchmarkTest02561) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	queryString := r.URL.Query()
	param := queryString.Get("BenchmarkTest02561")

tmpArrayUnique42 := []string{"", "", ""}
tmpArrayUnique42[0] = param
ah := NewArrayHolderWithValues(tmpArrayUnique42)
param = ah.Values[1]

	if param == "" {
		http.Error(w, "getQueryString() couldn't find expected parameter 'BenchmarkTest02561'", http.StatusBadRequest)
		return
	}

	bar := doSomething(param)

	fileName := filepath.Join("path/to/testfiles", bar)
	fis, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Couldn't open FileInputStream on file: '%s'\n", fileName)
		return
	}
	defer fis.Close()

	b := make([]byte, 1000)
	size, _ := fis.Read(b)
	fmt.Fprintf(w, "The beginning of file: '%s' is:\n\n", fileName)
	w.Write(b[:size])
}

func doSomething(param string) string {
	bar := ""
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the param value
	}
	return bar
}

func main() {
	http.Handle("/pathtraver-03/BenchmarkTest02561", &BenchmarkTest02561{})
	http.ListenAndServe(":8080", nil)
}
