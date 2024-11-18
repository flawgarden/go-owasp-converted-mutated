package controllers

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type BenchmarkTest02425 struct{}

func (b *BenchmarkTest02425) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest02425")
	if param == "" {
		param = ""
	}

	rememberMeKey := fmt.Sprint(rand.Int())
	user := "Ingrid"
	fullClassName := fmt.Sprintf("%T", b)
	testCaseNumber := fullClassName[len("benchmark."):]

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
		w.Write([]byte(fmt.Sprintf("Welcome back: %s<br/>", user)))
	} else {
		http.SetCookie(w, &http.Cookie{
			Name:     cookieName,
			Value:    rememberMeKey,
			Secure:   true,
			HttpOnly: true,
			Path:     r.URL.Path,
		})
		r.Context().Value(cookieName) // simulate setting it in context
		w.Write([]byte(fmt.Sprintf("%s has been remembered with cookie: %s whose value is: %s<br/>", user, cookieName, rememberMeKey)))
	}

	w.Write([]byte("Weak Randomness Test math/rand.Int() executed"))
}

func doSomething(param string) string {
	return param // Placeholder example, proper encoding should be applied if necessary
}
