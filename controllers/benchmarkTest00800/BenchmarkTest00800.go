package controllers

import (
	"fmt"
	"net/http"
	"net/url"
)

type BenchmarkTest00800 struct{}

func (b *BenchmarkTest00800) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.RawQuery
	paramval := "BenchmarkTest00800="
	paramLoc := -1

	if queryString != "" {
		paramLoc = findParamLocation(queryString, paramval)
	}

	if paramLoc == -1 {
		http.Error(w, fmt.Sprintf("getQueryString() couldn't find expected parameter '%s' in query string.", "BenchmarkTest00800"), http.StatusBadRequest)
		return
	}

	param := extractParam(queryString, paramLoc, paramval)
	param, _ = url.QueryUnescape(param)

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

	w.Header().Set("X-XSS-Protection", "0")
	fmt.Fprintf(w, "Formatted like: %s and %s.", "a", bar)
}

func findParamLocation(queryString, paramval string) int {
	return -1 // Implement the logic to find the parameter location
}

func extractParam(queryString string, paramLoc int, paramval string) string {
	return "" // Implement the logic to extract the parameter value
}

func main() {
	http.Handle("/xss-01/BenchmarkTest00800", &BenchmarkTest00800{})
	http.ListenAndServe(":8080", nil)
}
