package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02501 struct {
	Db *sql.DB
}

func (b *BenchmarkTest02501) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")



	randInt := newRandInt()
	rememberMeKey := strconv.Itoa(randInt)

	user := "Ingrid"
	testCaseNumber := "02501"
	user += testCaseNumber

	cookieName := "rememberMe" + testCaseNumber
	cookie, err := r.Cookie(cookieName)

	foundUser := false
	if err == nil {
		if cookie.Value == r.Context().Value(cookieName) {
			foundUser = true
		}
	}

	if foundUser {
		fmt.Fprintln(w, "Welcome back: "+user+"<br/>")
	} else {
		http.SetCookie(w, &http.Cookie{
			Name:     cookieName,
			Value:    rememberMeKey,
			Secure:   true,
			HttpOnly: true,
			Path:     r.URL.Path,
		})
		r.Context().Value(cookieName)
		fmt.Fprintln(w, user+" has been remembered with cookie: "+cookieName+" whose value is: "+rememberMeKey+"<br/>")
	}

	fmt.Fprintln(w, "Weak Randomness Test executed")
}

func doSomething(r *http.Request, param string) string {
	bar := param
	num := 106

	if (7*42)-num <= 200 {
		bar = param
	}

	return bar
}

func newRandInt() int {
	return int(time.Now().UnixNano() % 1000)
}

func init() {
	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	defer db.Close()
}
