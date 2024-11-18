package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02076 struct{}

func (b *BenchmarkTest02076) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.Header.Get("BenchmarkTest02076")
	param = decode(param)


	rand := getNextNumber()

	rememberMeKey := fmt.Sprintf("%f", rand)[2:]

	user := "SafeDonatella"
	testCaseNumber := "02076"
	user += testCaseNumber

	cookieName := "rememberMe" + testCaseNumber
	cookieValue := r.Context().Value(cookieName)

	if cookieValue != nil && cookieValue == rememberMeKey {
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

	fmt.Fprint(w, "Weak Randomness Test executed")
}

func decode(value string) string {
	decoded, _ := url.QueryUnescape(value)
	return decoded
}

func doSomething(param string) string {
	bar := "safe!"
	map15760 := make(map[string]interface{})
	map15760["keyA-15760"] = "a-Value"
	map15760["keyB-15760"] = param
	map15760["keyC"] = "another-Value"
	bar = map15760["keyB-15760"].(string)

	return bar
}

func getNextNumber() float64 {
	return float64(time.Now().UnixNano()%10000) / 10000.0
}
