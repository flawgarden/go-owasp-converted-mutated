package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01079 struct {
}

func (b *BenchmarkTest01079) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		b.doPost(w, req)
	} else if req.Method == http.MethodPost {
		b.doPost(w, req)
	}
}

func (b *BenchmarkTest01079) doPost(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""

	if req.Header.Get("BenchmarkTest01079") != "" {
		param = req.Header.Get("BenchmarkTest01079")
	}

	param, _ = url.QueryUnescape(param)


	// Simulating secure random operation
	l := int64(1234567890) // replace with actual secure random
	rememberMeKey := strconv.FormatInt(l, 10)

	user := "SafeLogan"
	fullClassName := fmt.Sprintf("%T", b)
	testCaseNumber := strings.TrimPrefix(fullClassName, "*controllers.")

	user += testCaseNumber

	cookieName := "rememberMe" + testCaseNumber

	foundUser := false
	cookies := req.Cookies()
	for _, cookie := range cookies {
		if cookieName == cookie.Name {
			if cookie.Value == req.Context().Value(cookieName) {
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
			Path:     req.RequestURI,
		})
		req.Context().Value(cookieName) // Simulate setting session attribute
		w.Write([]byte(user + " has been remembered with cookie: " + cookieName + " whose value is: " + rememberMeKey + "<br/>"))
	}

	w.Write([]byte("Weak Randomness Test executed"))
}

func (b *BenchmarkTest01079) doSomething(req *http.Request, param string) string {
	// Simulated operation
	return param
}
