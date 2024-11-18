package controllers

import (
	"fmt"
	"net/http"
	"strings"
)

type BenchmarkTest01703 struct{}

func (bt *BenchmarkTest01703) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.RawQuery
	paramval := "BenchmarkTest01703="
	paramLoc := strings.Index(queryString, paramval)
	if paramLoc == -1 {
		http.Error(w, fmt.Sprintf("getQueryString() couldn't find expected parameter '%s' in query string.", "BenchmarkTest01703"), http.StatusBadRequest)
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := strings.Index(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : paramLoc+ampersandLoc]
	}
	param = decodeURIComponent(param)


	randNum := generateSecureRandom()
	rememberMeKey := fmt.Sprintf("%.5f", randNum)[2:]

	user := "SafeFloyd"
	testCaseNumber := "01703"
	cookieName := "rememberMe" + testCaseNumber

	foundUser := false
	for _, cookie := range r.Cookies() {
		if cookieName == cookie.Name {
			if cookie.Value == r.Context().Value(cookieName) {
				foundUser = true
			}
		}
	}

	if foundUser {
		http.Error(w, "Welcome back: "+user, http.StatusOK)
	} else {
		http.SetCookie(w, &http.Cookie{
			Name:     cookieName,
			Value:    rememberMeKey,
			Secure:   true,
			HttpOnly: true,
			Path:     r.URL.Path,
		})
		r.Context().Value(cookieName) // Simulating session attribute
		fmt.Fprintf(w, "%s has been remembered with cookie: %s whose value is: %s<br/>", user, cookieName, rememberMeKey)
	}

	fmt.Fprintln(w, "Weak Randomness Test java.security.SecureRandom.nextFloat() executed")
}

type Test struct{}

func (t *Test) doSomething(r *http.Request, param string) string {
	var bar string
	if param != "" {
		bar = string(param)
	}
	return bar
}

func decodeURIComponent(encoded string) string {
	// Decode the URL-encoded string
	return strings.ReplaceAll(encoded, "%20", " ") // Simplified for demonstration
}

func generateSecureRandom() float64 {
	// Placeholder for secure random generation
	return 0.12345 // Replace with actual random generation logic
}
