package controllers

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
)

type BenchmarkTest01539 struct{}

func (b *BenchmarkTest01539) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest01539")
	if param == "" {
		param = ""
	}

	randomBytes := make([]byte, 40)
	if _, err := rand.Read(randomBytes); err != nil {
		http.Error(w, "Error generating random bytes.", http.StatusInternalServerError)
		return
	}

	rememberMeKey := base64.StdEncoding.EncodeToString(randomBytes)

	user := "SafeByron"
	fullClassName := "BenchmarkTest01539"
	testCaseNumber := fullClassName[len("BenchmarkTest"):]
	user += testCaseNumber

	cookieName := "rememberMe" + testCaseNumber

	cookie, err := r.Cookie(cookieName)
	foundUser := false
	if err == nil {
		if cookie.Value == r.Context().Value(cookieName) {
			foundUser = true
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
		r.Context().Value(cookieName) // Simulate saving in session
		fmt.Fprintf(w, "%s has been remembered with cookie: %s whose value is: %s<br/>", user, cookieName, rememberMeKey)
	}

	fmt.Fprintln(w, "Randomness Test executed")
}

func (b *BenchmarkTest01539) doSomething(param string) string {
	bar := "bob's your uncle"
	switch 'B' {
	case 'A':
		bar = param
	case 'B':
		bar = "bob"
	case 'C', 'D':
		bar = param
	}
	return bar
}
