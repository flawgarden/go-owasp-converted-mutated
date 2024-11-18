package controllers

import (
	"fmt"
	"math/rand/v2"
	"net/http"
	"reflect"
	"strings"
)

type BenchmarkTest01676 struct{}

func (b *BenchmarkTest01676) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	queryString := r.URL.RawQuery
	paramval := "BenchmarkTest01676="
	paramLoc := strings.Index(queryString, paramval)
	if paramLoc == -1 {
		http.Error(w, "getQueryString() couldn't find expected parameter 'BenchmarkTest01676' in query string.", http.StatusBadRequest)
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := strings.Index(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : paramLoc+ampersandLoc]
	}
	param = strings.TrimSpace(param)


	var value = rand.Float64()
	rememberMeKey := fmt.Sprintf("%.0f", value)[2:]

	user := "Donna"
	fullClassName := reflect.TypeOf(b).String()
	testCaseNumber := fullClassName[strings.LastIndex(fullClassName, ".")+1+len("BenchmarkTest"):]

	user += testCaseNumber
	cookieName := "rememberMe" + testCaseNumber

	foundUser := false
	cookies := r.Cookies()
	for _, cookie := range cookies {
		if cookieName == cookie.Name {
			if cookie.Value == r.Context().Value(cookieName) {
				foundUser = true
			}
		}
	}

	if foundUser {
		w.Write([]byte("Welcome back: " + user + "<br/>"))
	} else {
		http.SetCookie(w, &http.Cookie{
			Name:     cookieName,
			Value:    rememberMeKey,
			Secure:   true,
			HttpOnly: true,
			Path:     r.RequestURI,
		})
		r.Context().Value(cookieName).(http.ResponseWriter).Header().Set(cookieName, rememberMeKey)
		w.Write([]byte(user + " has been remembered with cookie: " + cookieName + " whose value is: " + rememberMeKey + "<br/>"))
	}

	w.Write([]byte("Weak Randomness Test executed"))
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	bar := param
	if len(param) > 1 {
		bar = param[:len(param)-1]
	}
	return bar
}
