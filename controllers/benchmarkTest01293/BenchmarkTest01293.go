package controllers

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
)

type BenchmarkTest struct{}

func (b *BenchmarkTest) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest01293")
	if param == "" {
		param = ""
	}

	randomBytes := make([]byte, 40)
	rand.Read(randomBytes)
	rememberMeKey := base64.StdEncoding.EncodeToString(randomBytes)

	user := "SafeByron"
	testCaseNumber := "BenchmarkTest01293"
	user += testCaseNumber
	cookieName := "rememberMe" + testCaseNumber

	cookies := r.Cookies()
	foundUser := false
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
			Secure:   true,
			HttpOnly: true,
			Path:     r.RequestURI,
		})
		r.Context().Value(cookieName)
		w.Write([]byte(user + " has been remembered with cookie: " + cookieName + " whose value is: " + rememberMeKey + "<br/>"))
	}
}

func (b *BenchmarkTest) doSomething(param string) string {
	bar := ""
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the param value
	}
	return bar
}

func main() {
	http.Handle("/weakrand-02/BenchmarkTest01293", &BenchmarkTest{})
	http.ListenAndServe(":8080", nil)
}
