package controllers

import (
	"fmt"
	"net/http"
	"strconv"
)

type BenchmarkTest01520 struct{}

func (b *BenchmarkTest01520) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest01520")
	if param == "" {
		param = ""
	}

	l := randomLong()
	rememberMeKey := strconv.FormatInt(l, 10)

	user := "Logan"
	testCaseNumber := "01520"
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
		fmt.Fprintf(w, "Welcome back: %s<br/>", user)
	} else {
		http.SetCookie(w, &http.Cookie{
			Name:     cookieName,
			Value:    rememberMeKey,
			Secure:   true,
			HttpOnly: true,
			Path:     r.RequestURI,
		})
		r.Context().Value(cookieName) // Storing in context (for example purposes)
		fmt.Fprintf(w, "%s has been remembered with cookie: %s whose value is: %s<br/>", user, cookieName, rememberMeKey)
	}

	fmt.Fprintln(w, "Weak Randomness Test randomLong() executed")
}

func (b *BenchmarkTest01520) doSomething(r *http.Request, param string) string {
	bar := "safe!"
	m := map[string]interface{}{
		"keyA-44642": "a_Value",
		"keyB-44642": param,
		"keyC":       "another_Value",
	}
	bar = m["keyB-44642"].(string)
	bar = m["keyA-44642"].(string)

	return bar
}

func randomLong() int64 {
	// Placeholder for an actual random long generator
	return 1234567890
}
