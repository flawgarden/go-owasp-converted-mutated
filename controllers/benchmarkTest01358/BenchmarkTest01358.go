package controllers

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01358 struct{}

func (b *BenchmarkTest01358) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	rand.Seed(time.Now().UnixNano())
	rememberMeKey := strconv.Itoa(rand.Int())

	user := "Ingrid"
	fullClassName := fmt.Sprintf("%T", b)
	testCaseNumber := fullClassName[len("benchmark.BenchmarkTest"):]

	user += testCaseNumber
	cookieName := "rememberMe" + testCaseNumber

	cookie, err := r.Cookie(cookieName)
	foundUser := false
	if err == nil && cookie.Value == r.URL.Query().Get(cookieName) {
		foundUser = true
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
		http.SetCookie(w, &http.Cookie{
			Name:  cookieName,
			Value: rememberMeKey,
			Path:  r.RequestURI,
		})
		fmt.Fprintf(w, "%s has been remembered with cookie: %s whose value is: %s<br/>", user, cookieName, rememberMeKey)
	}

	fmt.Fprint(w, "Weak Randomness Test executed")
}

func (b *BenchmarkTest01358) doSomething(r *http.Request, param string) string {
	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = append(valuesList[:0], valuesList[1:]...) // remove the 1st safe value
		bar = valuesList[1]                                    // get the last 'safe' value
	}
	return bar
}

func main() {
	http.Handle("/weakrand-03/BenchmarkTest01358", &BenchmarkTest01358{})
	http.ListenAndServe(":8080", nil)
}
