package controllers

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01294 struct{}

func (b *BenchmarkTest01294) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest01294")
	if param == "" {
		param = ""
	}


	rand.Seed(time.Now().UnixNano())
	rememberMeKey := fmt.Sprintf("%f", rand.Float64())[2:]

	user := "SafeFloyd"
	fullClassName := "BenchmarkTest01294"
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
		r.Context().Value(cookieName) // Temporarily store the rememberMeKey
		fmt.Fprintf(w, "%s has been remembered with cookie: %s whose value is: %s<br/>", user, cookieName, rememberMeKey)
	}

	fmt.Fprintf(w, "Weak Randomness Test executed")
}

func testDoSomething(r *http.Request, param string) string {
	bar := ""
	if param != "" {
		bar = string(param) // Simulating decode & encode
	}
	return bar
}
