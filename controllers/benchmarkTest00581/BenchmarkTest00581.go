package controllers

import (
	"fmt"
	"net/http"
	"strconv"
)

type BenchmarkTest00581Controller struct{}

func (c *BenchmarkTest00581Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")

	param := ""
	flag := true
	for name, values := range r.Form {
		for _, value := range values {
			if value == "BenchmarkTest00581" {
				param = name
				flag = false
				break
			}
		}
		if !flag {
			break
		}
	}

	bar := fmt.Sprintf("%s", param) // Simple escape for demo

	rand := float64(len(bar)) / 1000 // Dummy randomness
	rememberMeKey := strconv.FormatFloat(rand, 'f', -1, 64)[2:]

	user := "SafeDonna"
	testCaseNumber := "00581"
	user += testCaseNumber

	cookieName := "rememberMe" + testCaseNumber

	var foundUser bool
	cookies := r.Cookies()
	for _, cookie := range cookies {
		if cookieName == cookie.Name {
			if cookie.Value == r.Context().Value(cookieName) {
				foundUser = true
				break
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
		r.Context().Value(cookieName) // Assuming a place to set the session here
		fmt.Fprintf(w, "%s has been remembered with cookie: %s whose value is: %s<br/>", user, cookieName, rememberMeKey)
	}
	fmt.Fprintf(w, "Weak Randomness Test executed")
}
