package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type BenchmarkTest02607 struct{}

func (bt *BenchmarkTest02607) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	queryString := r.URL.RawQuery
	paramval := "BenchmarkTest02607="
	paramLoc := strings.Index(queryString, paramval)

	if paramLoc == -1 {
		fmt.Fprintf(w, "getQueryString() couldn't find expected parameter '%s' in query string.", "BenchmarkTest02607")
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := strings.Index(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : paramLoc+ampersandLoc]
	}

	param, _ = url.QueryUnescape(param)
	bar := doSomething(r, param)

	cookie := &http.Cookie{
		Name:     "SomeCookie",
		Value:    bar,
		Secure:   true,
		HttpOnly: true,
		Path:     r.RequestURI,
	}
	http.SetCookie(w, cookie)

	fmt.Fprintf(w, "Created cookie: 'SomeCookie': with value: '%s' and secure flag set to: true", bar)
}

func doSomething(r *http.Request, param string) string {
	sbxyz75490 := strings.Builder{}
	sbxyz75490.WriteString(param)
	bar := sbxyz75490.String() + "_SafeStuff"
	return bar
}

func main() {
	http.Handle("/securecookie-00/BenchmarkTest02607", &BenchmarkTest02607{})
	http.ListenAndServe(":8080", nil)
}
