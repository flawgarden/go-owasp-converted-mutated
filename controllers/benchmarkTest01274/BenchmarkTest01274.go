package controllers

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01274 struct {
}

func (b *BenchmarkTest01274) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest01274")
	if param == "" {
		param = ""
	}


	rand.Seed(time.Now().UnixNano())
	randNumber := rand.Intn(99)
	rememberMeKey := fmt.Sprintf("%d", randNumber)

	user := "Inga"
	fullClassName := fmt.Sprintf("%T", b)
	testCaseNumber := fullClassName[strings.LastIndex(fullClassName, ".")+1+len("BenchmarkTest"):]

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
			Path:     r.URL.Path,
		})
		r.Context().Value(cookieName) // Store the value in the context
		fmt.Fprintf(w, "%s has been remembered with cookie: %s whose value is: %s<br/>", user, cookieName, rememberMeKey)
	}

	fmt.Fprintln(w, "Weak Randomness Test rand.Intn(int) executed")
}

func init() {
	http.Handle("/weakrand-02/BenchmarkTest01274", &BenchmarkTest01274{})
}
