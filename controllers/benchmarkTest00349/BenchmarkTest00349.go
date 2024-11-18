package controllers

import (
	"context"
	"crypto/rand"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type BenchmarkTest00349 struct{}

func (b *BenchmarkTest00349) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	if r.Method == http.MethodGet {
		b.doPost(w, r)
	} else if r.Method == http.MethodPost {
		b.doPost(w, r)
	}
}

func (b *BenchmarkTest00349) doPost(w http.ResponseWriter, r *http.Request) {
	param := r.Body
	defer param.Close()

	randBytes := make([]byte, 8)
	_, _ = rand.Read(randBytes)
	randValue := float64(int64(randBytes[0])%100) / 100.0
	rememberMeKey := strconv.FormatFloat(randValue, 'f', -1, 64)[2:]

	user := "SafeDonna"
	fullClassName := "BenchmarkTest00349"
	testCaseNumber := strings.TrimPrefix(fullClassName, "BenchmarkTest")

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
		w.Write([]byte(fmt.Sprintf("Welcome back: %s<br/>", user)))
	} else {
		http.SetCookie(w, &http.Cookie{
			Name:     cookieName,
			Value:    rememberMeKey,
			Secure:   true,
			HttpOnly: true,
			Path:     r.RequestURI,
		})
		r.Context().Value(cookieName).(context.Context).Value(rememberMeKey)
		w.Write([]byte(fmt.Sprintf("%s has been remembered with cookie: %s whose value is: %s<br/>", user, cookieName, rememberMeKey)))
	}
	w.Write([]byte("Weak Randomness Test executed"))
}

func main() {
	http.Handle("/weakrand-00/BenchmarkTest00349", &BenchmarkTest00349{})
	http.ListenAndServe(":8080", nil)
}
