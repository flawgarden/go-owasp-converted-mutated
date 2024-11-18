package controllers

import (
	"fmt"
	"net/http"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02703 struct {
}

func (b *BenchmarkTest02703) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")


	rand := float32(randomValue())
	rememberMeKey := fmt.Sprintf("%.5f", rand)[2:]

	user := "Floyd"
	testCaseNumber := "02703"
	user += testCaseNumber
	cookieName := "rememberMe" + testCaseNumber

	c, err := r.Cookie(cookieName)
	foundUser := false
	if err == nil && c.Value == r.Context().Value(cookieName) {
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
		r.Context().Value(cookieName) // simulate storing for context
		fmt.Fprintf(w, "%s has been remembered with cookie: %s whose value is: %s<br/>", user, cookieName, rememberMeKey)
	}

	fmt.Fprintln(w, "Weak Randomness Test executed")
}

func doSomething(param string) string {
	return fmt.Sprintf("%s_SafeStuff", param)
}

func randomValue() float64 {
	return float64(1) // Stub for random value generator
}
