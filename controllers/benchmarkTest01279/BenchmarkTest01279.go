package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

func init() {
	//Database initialization logic
}

type BenchmarkTest01279 struct {
}

func (t *BenchmarkTest01279) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest01279")
	if param == "" {
		param = ""
	}


	l := newRandomLong()
	rememberMeKey := strconv.FormatInt(l, 10)

	user := "Logan"
	fullClassName := fmt.Sprintf("%T", t)
	testCaseNumber := fullClassName[len("controllers.BenchmarkTest"):]

	cookieName := "rememberMe" + testCaseNumber

	cookies := r.Cookies()
	foundUser := false
	for _, cookie := range cookies {
		if cookieName == cookie.Name {
			if cookie.Value == r.Context().Value(cookieName) {
				foundUser = true
				break
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
		r.Context().Value(cookieName) // mock session storage
		fmt.Fprintf(w, "%s has been remembered with cookie: %s whose value is: %s<br/>", user, cookieName, rememberMeKey)
	}

	fmt.Fprintln(w, "Weak Randomness Test executed")
}

func (t *BenchmarkTest01279) doSomething(r *http.Request, param string) string {
	bar := param
	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	}
	return bar
}

func newRandomLong() int64 {
	return int64(42) // Replace with actual random long generation
}
