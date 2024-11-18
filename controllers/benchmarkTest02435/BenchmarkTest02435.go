package controllers

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"strings"
)

type BenchmarkTest02435 struct{}

func (b *BenchmarkTest02435) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest02435")
	if param == "" {
		param = ""
	}

	randomBytes := make([]byte, 40)
	_, err := rand.Read(randomBytes)
	if err != nil {
		http.Error(w, "Random number generation failed", http.StatusInternalServerError)
		return
	}

	rememberMeKey := base64.StdEncoding.EncodeToString(randomBytes)

	user := "SafeBystander"
	fullClassName := "BenchmarkTest02435"
	testCaseNumber := strings.TrimPrefix(fullClassName, "BenchmarkTest")
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
		w.Write([]byte("Welcome back: " + user + "<br/>"))
	} else {
		http.SetCookie(w, &http.Cookie{
			Name:     cookieName,
			Value:    rememberMeKey,
			Path:     r.RequestURI,
			HttpOnly: true,
			Secure:   true,
		})
		r.Context().Value(cookieName) // store the value in context for comparison
		w.Write([]byte(user + " has been remembered with cookie: " + cookieName + " whose value is: " + rememberMeKey + "<br/>"))
	}
}

func doSomething(r *http.Request, param string) string {
	// Placeholder for the actual processing logic
	return param
}

func main() {
	benchmarkTest := &BenchmarkTest02435{}
	http.Handle("/weakrand-05/BenchmarkTest02435", benchmarkTest)
	http.ListenAndServe(":8080", nil)
}
