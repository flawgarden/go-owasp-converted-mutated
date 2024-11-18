package controllers

import (
	"fmt"
	"math/rand/v2"
	"net/http"
	"net/url"
	"strings"
)

type BenchmarkTest01183 struct{}

func (b *BenchmarkTest01183) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	if headers := r.Header["BenchmarkTest01183"]; len(headers) > 0 {
		param = headers[0]
	}

	param, _ = url.QueryUnescape(param)

	randKey := generateRandomKey()

	user := "Floyd"
	fullClassName := fmt.Sprintf("%T", b)
	testCaseNumber := strings.Split(fullClassName, ".")[1]

	user += testCaseNumber
	cookieName := "rememberMe" + testCaseNumber

	foundUser := false
	cookies := r.Cookies()
	for _, cookie := range cookies {
		if cookieName == cookie.Name && cookie.Value == r.Context().Value(cookieName) {
			foundUser = true
			break
		}
	}

	if foundUser {
		fmt.Fprintf(w, "Welcome back: %s<br/>", user)
	} else {
		http.SetCookie(w, &http.Cookie{
			Name:     cookieName,
			Value:    randKey,
			Secure:   true,
			HttpOnly: true,
			Path:     r.URL.Path,
		})
		r.Context().Value(cookieName) // Simulating session storage
		fmt.Fprintf(w, "%s has been remembered with cookie: %s whose value is: %s<br/>", user, cookieName, randKey)
	}

	fmt.Fprintln(w, "Weak Randomness Test generated executed")
}

func (b *BenchmarkTest01183) doSomething(r *http.Request, param string) string {
	var bar string
	guess := "ABC"
	switchTarget := guess[2]

	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bobs_your_uncle"
	case 'C', 'D':
		bar = param
	default:
		bar = "bobs_your_uncle"
	}

	return bar
}

func generateRandomKey() string {
	return fmt.Sprintf("%f", rand.Float64())[2:] // Trim off the 0.
}

func main() {
	http.Handle("/weakrand-02/BenchmarkTest01183", &BenchmarkTest01183{})
	http.ListenAndServe(":8080", nil)
}
