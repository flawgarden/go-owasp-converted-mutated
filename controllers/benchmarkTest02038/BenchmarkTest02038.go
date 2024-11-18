package controllers

import (
	"net/http"
	"net/url"
	"strconv"
)

type BenchmarkTest02038 struct{}

func (bt *BenchmarkTest02038) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	if header := r.Header.Get("BenchmarkTest02038"); header != "" {
		param = header
		param, _ = url.QueryUnescape(param)
	}

	value := strconv.FormatFloat(randomFloat(), 'f', -1, 64)
	rememberMeKey := value[2:]

	user := "Doug"
	fullClassName := "benchmarks.BenchmarkTest02038"
	testCaseNumber := fullClassName[len("benchmarks.BenchmarkTest"):]

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
			Secure:   true,
			HttpOnly: true,
			Path:     r.RequestURI,
		})
		r.Context().Value(cookieName) // Assume we set the session value
		w.Write([]byte(user + " has been remembered with cookie: " + cookieName + " whose value is: " + rememberMeKey + "<br/>"))
	}
	w.Write([]byte("Weak Randomness Test executed"))
}

func doSomething(r *http.Request, param string) string {
	bar := ""
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = append(valuesList[:0], valuesList[1:]...) // remove the 1st safe value
		bar = valuesList[0]
	}
	return bar
}

func randomFloat() float64 {
	return 0.123456 // Replace with a proper random number implementation
}
