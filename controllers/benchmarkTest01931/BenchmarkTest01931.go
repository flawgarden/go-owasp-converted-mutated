package controllers

import (
	"fmt"
	"math/rand/v2"
	"net/http"
	"net/url"
)

type BenchmarkTest01931 struct {
	http.Handler
}

func (b *BenchmarkTest01931) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	b.doPost(w, r)
}

func (b *BenchmarkTest01931) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	if r.Header.Get("BenchmarkTest01931") != "" {
		param = r.Header.Get("BenchmarkTest01931")
	}

	param, _ = url.QueryUnescape(param)


	rand := rand.Float32()
	rememberMeKey := fmt.Sprintf("%.2f", rand)[2:]

	user := "Floyd"
	fullClassName := fmt.Sprintf("%T", b)
	testCaseNumber := fullClassName[len("controllers."):]

	user += testCaseNumber
	cookieName := "rememberMe" + testCaseNumber

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

		r.Context().Value(cookieName) // Simulate session storage
		fmt.Fprintf(w, "%s has been remembered with cookie: %s whose value is: %s<br/>", user, cookieName, rememberMeKey)
	}

	fmt.Fprintln(w, "Weak Randomness Test executed")
}

func doSomething(r *http.Request, param string) string {
	a40465 := param
	b40465 := a40465 + " SafeStuff"
	c40465 := b40465[:len(b40465)-1]
	var d40465 string
	d40465 = string([]byte(c40465))
	return d40465
}
