package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type BenchmarkTest02040 struct{}

func (b *BenchmarkTest02040) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := ""

	if headers := r.Header["BenchmarkTest02040"]; len(headers) > 0 {
		param = headers[0]
		param, _ = url.QueryUnescape(param)
	}

	value := randomValue()
	rememberMeKey := strings.TrimPrefix(fmt.Sprintf("%f", value), "0.")
	user := "Doug"
	testCaseNumber := "02040"
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
		fmt.Fprintf(w, "Welcome back: %s<br/>", user)
	} else {
		http.SetCookie(w, &http.Cookie{
			Name:     cookieName,
			Value:    rememberMeKey,
			Secure:   true,
			HttpOnly: true,
			Path:     r.RequestURI,
		})
		r.Context().Value(cookieName) // Set session attribute
		fmt.Fprintf(w, "%s has been remembered with cookie: %s whose value is: %s<br/>", user, cookieName, rememberMeKey)
	}
	fmt.Fprintln(w, "Weak Randomness Test executed")
}

func doSomething(r *http.Request, param string) string {
	bar := ""
	guess := "ABC"
	switch guess[1] {
	case 'A':
		bar = param
	case 'B':
		bar = "bob"
	case 'C', 'D':
		bar = param
	default:
		bar = "bob's your uncle"
	}
	return bar
}

func randomValue() float64 {
	return 0.123456789 // A placeholder for random value. Replace with actual random logic if needed.
}
