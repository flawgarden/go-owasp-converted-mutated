package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type BenchmarkTest02592 struct{}

func (bt *BenchmarkTest02592) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	queryString := r.URL.RawQuery
	paramval := "BenchmarkTest02592="
	paramLoc := -1

	if queryString != "" {
		paramLoc = len(queryString) - len(paramval)
		if queryString[paramLoc:] != paramval {
			http.Error(w, fmt.Sprintf("getQueryString() couldn't find expected parameter '%s' in query string.", "BenchmarkTest02592"), http.StatusBadRequest)
			return
		}
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := findAmpersand(queryString, paramLoc)
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : ampersandLoc]
	}
	param, _ = url.QueryUnescape(param)

	bar := doSomething(param)

	w.Header().Set("X-XSS-Protection", "0")
	response := map[string]string{"message": fmt.Sprintf("Formatted like: %s and b.", bar)}
	output, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func doSomething(param string) string {
	bar := param
	if param != "" && len(param) > 1 {
		sbxyz52014 := []rune(param)
		bar = string(sbxyz52014[:len(param)-1]) + "Z"
	}
	return bar
}

func findAmpersand(queryString string, paramLoc int) int {
	for i := paramLoc; i < len(queryString); i++ {
		if queryString[i] == '&' {
			return i
		}
	}
	return -1
}

func main() {
	http.Handle("/xss-05/BenchmarkTest02592", &BenchmarkTest02592{})
	http.ListenAndServe(":8080", nil)
}
