package controllers

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type BenchmarkTest01275 struct{}

func (b *BenchmarkTest01275) Get(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	response := "text/html;charset=UTF-8"
	w.Header().Set("Content-Type", response)

	param := r.URL.Query().Get("BenchmarkTest01275")
	if param == "" {
		param = ""
	}


	randNumber := rand.Intn(99)
	rememberMeKey := fmt.Sprint(randNumber)

	user := "Inga"
	fullClassName := fmt.Sprintf("%T", b)
	testCaseNumber := fullClassName[len("controllers."):]

	user += testCaseNumber
	cookieName := "rememberMe" + testCaseNumber

	cookies := r.Cookies()
	foundUser := false
	for _, cookie := range cookies {
		if cookieName == cookie.Name {
			if cookie.Value == r.Context().Value(cookieName).(string) {
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
		r.Context().Value(cookieName) // set session attribute
		fmt.Fprintf(w, "%s has been remembered with cookie: %s whose value is: %s<br/>", user, cookieName, rememberMeKey)
	}

	fmt.Fprintf(w, "Weak Randomness Test executed")
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	return param + "_SafeStuff"
}
