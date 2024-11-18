package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/url"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01934 struct {
}

func (b *BenchmarkTest01934) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := r.Header.Get("BenchmarkTest01934")
	param, _ = url.QueryUnescape(param)


	l := time.Now().UnixNano()
	rememberMeKey := fmt.Sprintf("%d", l)

	user := "Logan"
	testCaseNumber := "01934"
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
		fmt.Fprintf(w, "Welcome back: %s<br/>", user)
	} else {
		http.SetCookie(w, &http.Cookie{
			Name:     cookieName,
			Value:    rememberMeKey,
			Secure:   true,
			HttpOnly: true,
			Path:     r.URL.Path,
		})
		r.Context().Value(cookieName)
		fmt.Fprintf(w, "%s has been remembered with cookie: %s whose value is: %s<br/>", user, cookieName, rememberMeKey)
	}
	fmt.Fprintf(w, "Weak Randomness Test executed")
}

func doSomething(param string) string {
	bar := sql.NullString{String: param, Valid: true}
	return bar.String
}
