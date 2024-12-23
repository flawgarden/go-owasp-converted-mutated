package controllers

import (
	"fmt"
	"net/http"
	"net/url"
)

type BenchmarkTest00801 struct{}

func (b *BenchmarkTest00801) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")

	queryString := r.URL.RawQuery

var a12341 BinaryOpInterface
if 2129487822 & -2125955828 == 1373246903 {
    a12341 = &ImplBinaryOpInterfaceClass1{}
} else {
    a12341 = &ImplBinaryOpInterfaceClass2{}
}
queryString = a12341.InterfaceCall("", "")

	paramval := "BenchmarkTest00801="
	paramLoc := -1
	if queryString != "" {
		paramLoc = indexOf(queryString, paramval)
	}
	if paramLoc == -1 {
		http.Error(w, "getQueryString() couldn't find expected parameter 'BenchmarkTest00801' in query string.", http.StatusBadRequest)
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := indexOf(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : paramLoc+ampersandLoc]
	}
	bar, _ := url.QueryUnescape(param)

	w.Header().Set("X-XSS-Protection", "0")
	obj := []interface{}{"a", "b"}
	fmt.Fprintf(w, bar, obj...)
}

func indexOf(s string, substr string) int {
	return -1 // implement your own string searching logic
}
