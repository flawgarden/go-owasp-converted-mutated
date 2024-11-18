package controllers

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type BenchmarkTest00163 struct{}

func (b *BenchmarkTest00163) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.Header.Get("BenchmarkTest00163")
	param, _ = url.QueryUnescape(param)

	rand.Seed(time.Now().UnixNano())
	stuff := rand.NormFloat64()
	rememberMeKey := fmt.Sprintf("%.2f", stuff)[2:]

	user := "Gayle"
	fullClassName := fmt.Sprintf("%T", b)
	testCaseNumber := fullClassName[strings.LastIndex(fullClassName, ".")+len("BenchmarkTest"):]

	user += testCaseNumber
	cookieName := "rememberMe" + testCaseNumber

	var foundUser bool
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
			Value:    rememberMeKey,
			Secure:   true,
			HttpOnly: true,
			Path:     r.RequestURI,
		})
		r.Context().Value(cookieName)
		fmt.Fprintf(w, "%s has been remembered with cookie: %s whose value is: %s<br/>", user, cookieName, rememberMeKey)
	}

	fmt.Fprintf(w, "Weak Randomness Test math/rand.NormFloat64() executed")
}

func main() {
	http.Handle("/weakrand-00/BenchmarkTest00163", &BenchmarkTest00163{})
	http.ListenAndServe(":8080", nil)
}
