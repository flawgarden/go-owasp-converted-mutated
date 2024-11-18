package controllers

import (
	"fmt"
	"html"
	"net/http"
	"net/url"
	"strings"
)

type BenchmarkTest00820 struct{}

func (b *BenchmarkTest00820) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	queryString := r.URL.RawQuery
	paramval := "BenchmarkTest00820="
	paramLoc := -1
	if queryString != "" {
		paramLoc = findParamLocation(queryString, paramval)
	}
	if paramLoc == -1 {
		http.Error(w, fmt.Sprintf("getQueryString() couldn't find expected parameter '%s' in query string.", "BenchmarkTest00820"), http.StatusBadRequest)
		return
	}

	param := getQueryParamValue(queryString, paramLoc, paramval)

	bar := param
	cookie := http.Cookie{Name: "SomeCookie", Value: bar, Path: r.RequestURI, HttpOnly: true, Secure: false}
	http.SetCookie(w, &cookie)

	responseMsg := fmt.Sprintf("Created cookie: 'SomeCookie': with value: '%s' and secure flag set to: false", htmlEscape(bar))
	w.Write([]byte(responseMsg))
}

func findParamLocation(queryString, paramval string) int {
	return strings.Index(queryString, paramval)
}

func getQueryParamValue(queryString string, paramLoc int, paramval string) string {
	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := strings.Index(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : paramLoc+ampersandLoc]
	}
	decodedParam, _ := url.QueryUnescape(param)
	return decodedParam
}

func htmlEscape(str string) string {
	return html.EscapeString(str)
}

func main() {
	http.Handle("/", &BenchmarkTest00820{})
	http.ListenAndServe(":8080", nil)
}
