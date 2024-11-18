package controllers

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01930 struct{}

func (b *BenchmarkTest01930) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.Header.Add("BenchmarkTest01930", "")
	param := r.Header.Get("BenchmarkTest01930")
	param, _ = url.QueryUnescape(param)


	rand.Seed(time.Now().UnixNano())
	value := rand.Float64()
	rememberMeKey := fmt.Sprintf("%.0f", value)[2:]

	user := "Donna"
	fullClassName := fmt.Sprintf("%T", b)
	testCaseNumber := fullClassName[len("controllers."):]

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
		w.Write([]byte("Welcome back: " + user + "<br/>"))
	} else {
		http.SetCookie(w, &http.Cookie{
			Name:     cookieName,
			Value:    rememberMeKey,
			Secure:   true,
			HttpOnly: true,
			Path:     r.URL.Path,
		})
		r.Context().Value(cookieName) // Storing in context as an example
		w.Write([]byte(user + " has been remembered with cookie: " + cookieName + " whose value is: " + rememberMeKey + "<br/>"))
	}

	w.Write([]byte("Weak Randomness Test executed"))
}

func doSomething(r *http.Request, param string) string {
	return param + "_SafeStuff"
}
