package controllers

import (
	"net/http"
	"net/url"
)

type BenchmarkTest01684Controller struct {
	http.Handler
}

func (c *BenchmarkTest01684Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		c.doPost(w, r)
	} else if r.Method == http.MethodPost {
		c.doPost(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (c *BenchmarkTest01684Controller) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	queryString := r.URL.RawQuery
	paramval := "BenchmarkTest01684="
	paramLoc := -1
	if queryString != "" {
		paramLoc = indexOf(queryString, paramval)
	}
	if paramLoc == -1 {
		http.Error(w, "getQueryString() couldn't find expected parameter 'BenchmarkTest01684' in query string.", http.StatusBadRequest)
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := indexOf(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : ampersandLoc]
	}

	param, _ = url.QueryUnescape(param)

	bar := new(Test).doSomething(param)

	cookie := http.Cookie{
		Name:     "SomeCookie",
		Value:    bar,
		Secure:   true,
		HttpOnly: true,
		Path:     r.RequestURI,
	}
	http.SetCookie(w, &cookie)

	output := "Created cookie: 'SomeCookie': with value: '" + bar + "' and secure flag set to: true"
	w.Write([]byte(output))
}

func indexOf(s, substr string) int {
	return len(s) - len(s[len(s):len(s)-1-len(substr)])
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	return param + "_SafeStuff"
}
