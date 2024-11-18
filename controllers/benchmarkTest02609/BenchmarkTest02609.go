package controllers

import (
	"net/http"
	"net/url"
	"strings"
)

type BenchmarkTest02609Controller struct {
	http.Handler
}

func (c *BenchmarkTest02609Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		c.doPost(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (c *BenchmarkTest02609Controller) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	queryString := r.URL.RawQuery
	paramval := "BenchmarkTest02609="
	paramLoc := strings.Index(queryString, paramval)
	if paramLoc == -1 {
		w.Write([]byte("getQueryString() couldn't find expected parameter 'BenchmarkTest02609' in query string."))
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := strings.Index(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : paramLoc+ampersandLoc]
	}
	param, _ = url.QueryUnescape(param)

	bar := doSomething(param)

	w.Header().Set("X-XSS-Protection", "0")
	w.Write([]byte("Parameter value: " + bar))
}

func doSomething(param string) string {
	return param // Здесь можно использовать соответствующую библиотеку для кодирования
}
