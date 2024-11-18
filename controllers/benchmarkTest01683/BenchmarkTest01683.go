package controllers

import (
	"fmt"
	"net/http"
	"net/url"
)

type BenchmarkTest01683 struct{}

func (b *BenchmarkTest01683) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	queryString := r.URL.RawQuery
	paramval := "BenchmarkTest01683="
	paramLoc := -1

	if queryString != "" {
		paramLoc = len(queryString) - len(paramval)
	}

	if paramLoc == -1 {
		http.Error(w, fmt.Sprintf("getQueryString() couldn't find expected parameter '%s' in query string.", "BenchmarkTest01683"), http.StatusBadRequest)
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := len(queryString)
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : ampersandLoc]
	}
	param, _ = url.QueryUnescape(param)

	bar := new(Test).doSomething(r, param)

	cookie := http.Cookie{
		Name:     "SomeCookie",
		Value:    bar,
		Path:     r.URL.Path,
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)

	fmt.Fprintf(w, "Created cookie: 'SomeCookie': with value: '%s' and secure flag set to: false", bar)
}

type Test struct{}

func (t *Test) doSomething(r *http.Request, param string) string {
	return param
}

func main() {
	http.Handle("/securecookie-00/BenchmarkTest01683", &BenchmarkTest01683{})
	http.ListenAndServe(":8080", nil)
}
