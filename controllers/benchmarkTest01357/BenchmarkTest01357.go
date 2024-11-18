package controllers

import (
	"fmt"
	"math/rand"
	"net/http"
)

type BenchmarkTest01357 struct{}

func (b *BenchmarkTest01357) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	rand.Seed(rand.Int63())
	rememberMeKey := fmt.Sprintf("%d", rand.Int())

	user := "Ingrid"
	testCaseNumber := "01357"
	user += testCaseNumber
	cookieName := "rememberMe" + testCaseNumber

	cookie, err := r.Cookie(cookieName)
	foundUser := false

	if err == nil && cookie.Value == r.Context().Value(cookieName) {
		foundUser = true
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
		r.Context().Value(cookieName) // Simulating session attribute
		w.Write([]byte(user + " has been remembered with cookie: " + cookieName + " whose value is: " + rememberMeKey + "<br/>"))
	}

	w.Write([]byte("Weak Randomness Test using math/rand executed"))
}

func (b *BenchmarkTest01357) doSomething(r *http.Request, param string) string {
	bar := ""
	if param != "" {
		bar = string([]byte(param)) // Base64 decoding can be added here if needed.
	}
	return bar
}
