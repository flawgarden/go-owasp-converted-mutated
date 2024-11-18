package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type BenchmarkTest02557 struct{}

func (b *BenchmarkTest02557) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	queryString := r.URL.RawQuery
	paramval := "BenchmarkTest02557="
	paramLoc := strings.Index(queryString, paramval)

	if paramLoc == -1 {
		http.Error(w, fmt.Sprintf("getQueryString() couldn't find expected parameter '%s' in query string.", "BenchmarkTest02557"), http.StatusBadRequest)
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := strings.Index(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : ampersandLoc+paramLoc]
	}
	param = decode(param)

	bar := doSomething(param)

	fileTarget := fmt.Sprintf("%s/Test.txt", bar)
	responseMessage := fmt.Sprintf("Access to file: '%s' created.", fileTarget)
	if _, err := os.Stat(fileTarget); err == nil {
		responseMessage += " And file already exists."
	} else {
		responseMessage += " But file doesn't exist yet."
	}
	w.Write([]byte(responseMessage))
}

func doSomething(param string) string {
	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the last 'safe' value
	}
	return bar
}

func decode(param string) string {
	decoded, err := url.QueryUnescape(param)
	if err != nil {
		return param
	}
	return decoded
}

func main() {
	http.Handle("/pathtraver-03/BenchmarkTest02557", &BenchmarkTest02557{})
	http.ListenAndServe(":8080", nil)
}
