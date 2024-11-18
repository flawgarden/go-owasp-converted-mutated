package controllers

import (
	"fmt"
	"net/http"
	"strings"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01141Controller struct {
}

func (c *BenchmarkTest01141Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := getHeaderParam(r)

	bar := new(Test).doSomething(r, param)

	trySetCookie(w, r, bar)
	http.Error(w, "Weak Randomness Test executed", http.StatusOK)
}

func getHeaderParam(r *http.Request) string {
	var param string
	for name := range r.Header {
		if !isCommonHeader(name) {
			param = name
			break
		}
	}
	return param
}

func isCommonHeader(name string) bool {
	return strings.Contains("Accept,Content-Type,User-Agent", name)
}

func trySetCookie(w http.ResponseWriter, r *http.Request, bar string) {
	rememberMeKey := fmt.Sprintf("%d", getSecureRandomLong())
	cookieName := "rememberMe" + bar

	var foundUser bool
	for _, cookie := range r.Cookies() {
		if cookie.Name == cookieName && cookie.Value == r.Context().Value(cookieName) {
			foundUser = true
			break
		}
	}

	if foundUser {
		w.Write([]byte("Welcome back: " + "SafeUser" + "<br/>"))
	} else {
		http.SetCookie(w, &http.Cookie{
			Name:     cookieName,
			Value:    rememberMeKey,
			Secure:   true,
			HttpOnly: true,
			Path:     r.RequestURI,
		})
		r.Context().Value(cookieName) // Store in context
		w.Write([]byte("SafeUser has been remembered with cookie: " + cookieName + " whose value is: " + rememberMeKey + "<br/>"))
	}
}

func getSecureRandomLong() int64 {
	// Simulated SecureRandom logic for example purposes.
	return 1234567890
}

type Test struct{}

func (t *Test) doSomething(r *http.Request, param string) string {
	bar := param
	if len(param) > 1 {
		bar = param[:len(param)-1] + "Z"
	}
	return bar
}
