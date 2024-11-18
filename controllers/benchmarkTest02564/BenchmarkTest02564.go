package controllers

import (
	"fmt"
	"net/http"
	"os"
)

type BenchmarkTest02564 struct{}

func (b *BenchmarkTest02564) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.RawQuery
	paramVal := "BenchmarkTest02564="
	paramLoc := -1
	if queryString != "" {
		paramLoc = findParamLoc(queryString, paramVal)
	}
	if paramLoc == -1 {
		http.Error(w, fmt.Sprintf("getQueryString() couldn't find expected parameter '%s' in query string.", "BenchmarkTest02564"), http.StatusBadRequest)
		return
	}

	param := extractParam(queryString, paramLoc, paramVal)

	bar := doSomething(param)

	fileName := ""
	fos, err := os.Create(bar)
	if err != nil {
		fmt.Printf("Couldn't open FileOutputStream on file: '%s'\n", fileName)
		return
	}
	defer fos.Close()

	w.Write([]byte(fmt.Sprintf("Now ready to write to file: %s", fileName)))
}

func findParamLoc(queryString, paramVal string) int {
	return -1 // Implement your logic to find the parameter location
}

func extractParam(queryString string, paramLoc int, paramVal string) string {
	return "" // Implement your logic to extract the parameter value
}

func doSomething(param string) string {
	var bar string
	guess := "ABC"
	switchTarget := guess[2]

	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bobs_your_uncle"
	case 'C':
	case 'D':
		bar = param
	default:
		bar = "bobs_your_uncle"
	}

	return bar
}

func main() {
	http.Handle("/pathtraver-03/BenchmarkTest02564", &BenchmarkTest02564{})
	http.ListenAndServe(":8080", nil)
}
