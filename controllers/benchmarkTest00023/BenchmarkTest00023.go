package controllers

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00023Controller struct {
	http.Handler
}

func (c *BenchmarkTest00023Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano())).Float64()
	rememberMeKey := fmt.Sprintf("%.6f", rnd)[2:]

	user := "Floyd"
	testCaseNumber := "00023"
	user += testCaseNumber

	cookieName := fmt.Sprintf("rememberMe%s", testCaseNumber)
	cookie, err := r.Cookie(cookieName)
	foundUser := false

	if err == nil && cookie.Value == r.Context().Value(cookieName).(string) {
		foundUser = true
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
		r.Context().Value(cookieName) // Save the key in context (mocking session storage)
		fmt.Fprintf(w, "%s has been remembered with cookie: %s whose value is: %s<br/>", user, cookieName, rememberMeKey)
	}

	fmt.Fprintf(w, "Weak Randomness Test executed")
}
