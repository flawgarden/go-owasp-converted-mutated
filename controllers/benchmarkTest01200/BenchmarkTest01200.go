package controllers

import (
	"fmt"
	"net/http"
	"net/url"

	"xorm.io/xorm"
)

type BenchmarkTest01200 struct {
	engine *xorm.Engine
}

func NewBenchmarkTest01200(engine *xorm.Engine) *BenchmarkTest01200 {
	return &BenchmarkTest01200{engine: engine}
}

func (b *BenchmarkTest01200) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.Header.Get("BenchmarkTest01200")
	param, _ = url.QueryUnescape(param)


	rand, err := secureRandomFloat()
	if err != nil {
		http.Error(w, "Error generating random float", http.StatusInternalServerError)
		return
	}

	rememberMeKey := fmt.Sprintf("%f", rand)[2:]
	user := "SafeFloyd"
	testCaseNumber := "01200" // Update this according to your logic
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
		r.Context().Value(cookieName) // Update session or context
		fmt.Fprintf(w, "%s has been remembered with cookie: %s whose value is: %s<br/>",
			user, cookieName, rememberMeKey)
	}

	fmt.Fprintln(w, "Weak Randomness Test executed")
}

func (b *BenchmarkTest01200) doSomething(param string) string {
	bar := param
	if len(param) > 1 {
		bar = param[:len(param)-1]
	}
	return bar
}

func secureRandomFloat() (float64, error) {
	// Implement a secure random float generator
	return 0.0, nil
}
