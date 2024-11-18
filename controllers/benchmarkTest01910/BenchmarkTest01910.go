package controllers

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

type BenchmarkTest01910Controller struct{}

func (c *BenchmarkTest01910Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")

	param := ""
	if r.Header.Get("BenchmarkTest01910") != "" {
		param = r.Header.Get("BenchmarkTest01910")
	}
	param, _ = url.QueryUnescape(param)


	rand.Seed(time.Now().UnixNano())
	value := rand.Float64()
	rememberMeKey := fmt.Sprintf("%.0f", value)[2:]

	user := "Doug"
	fullClassName := "controllers.BenchmarkTest01910Controller"
	testCaseNumber := fullClassName[len(fullClassName)-12:]

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
			Path:     r.URL.Path,
		})
		r.Context().Value(cookieName) // Mocking session storage
		fmt.Fprintf(w, "%s has been remembered with cookie: %s whose value is: %s<br/>", user, cookieName, rememberMeKey)
	}
	fmt.Fprint(w, "Weak Randomness Test executed")
}

func doSomething(param string) string {
	var bar string
	guess := "ABC"
	switchTarget := guess[1] // condition 'B', which is safe

	switch switchTarget {
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
