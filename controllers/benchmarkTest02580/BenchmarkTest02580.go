package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02580 struct {
}

func (b *BenchmarkTest02580) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	queryString := r.URL.RawQuery
	paramval := "BenchmarkTest02580="
	paramLoc := strings.Index(queryString, paramval)

	if paramLoc == -1 {
		http.Error(w, "getQueryString() couldn't find expected parameter 'BenchmarkTest02580' in query string.", http.StatusBadRequest)
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := strings.Index(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : ampersandLoc+paramLoc]
	}
	param, err := url.QueryUnescape(param)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bar := doSomething(param)

	w.Header().Set("X-XSS-Protection", "0")
	output := fmt.Sprintf(bar, "a", "b")
	w.Write([]byte(output))
}

func doSomething(param string) string {
	var sbxyz1658 strings.Builder
	sbxyz1658.WriteString(param)
	bar := sbxyz1658.String() + "_SafeStuff"
	return bar
}
