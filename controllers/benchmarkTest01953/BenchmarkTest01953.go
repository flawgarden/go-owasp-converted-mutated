package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type BenchmarkTest01953 struct{}

func (b *BenchmarkTest01953) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.Header.Get("BenchmarkTest01953")
	param, _ = url.QueryUnescape(param)

	l := nextLong()
	rememberMeKey := strconv.FormatInt(l, 10)

	user := "SafeLogan"
	fullClassName := fmt.Sprintf("%T", b)
	testCaseNumber := fullClassName[len("main.BenchmarkTest"):]

	user += testCaseNumber
	cookieName := "rememberMe" + testCaseNumber

	cookies := r.Cookies()
	foundUser := false
	for _, cookie := range cookies {
		if cookieName == cookie.Name {
			if cookie.Value == r.Context().Value(cookieName).(string) {
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
			Path:     r.URL.Path,
		})
		r.Context().Value(cookieName) // Store the rememberMeKey in the context
		fmt.Fprintf(w, "%s has been remembered with cookie: %s whose value is: %s<br/>", user, cookieName, rememberMeKey)
	}

	fmt.Fprintln(w, "Weak Randomness Test executed")
}

func doSomething(r *http.Request, param string) string {
	bar := param
	if (7*42)-86 > 200 {
		bar = "This_should_always_happen"
	}
	return bar
}

func nextLong() int64 {
	return int64(42) // Replace with actual random generation logic
}

func main() {
	http.Handle("/weakrand-04/BenchmarkTest01953", &BenchmarkTest01953{})
	http.ListenAndServe(":8080", nil)
}
