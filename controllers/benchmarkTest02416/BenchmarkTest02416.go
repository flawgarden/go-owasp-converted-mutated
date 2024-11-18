package controllers

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type BenchmarkTest02416 struct{}

func (b *BenchmarkTest02416) Get(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest02416")
	if param == "" {
		param = ""
	}


	rand.Seed(time.Now().UnixNano())
	value := rand.Float64()
	rememberMeKey := strconv.FormatFloat(value, 'f', -1, 64)[2:]

	user := "Donna"
	fullClassName := "BenchmarkTest02416"

	cookieName := "rememberMe" + fullClassName

	cookies := r.Cookies()
	foundUser := false
	for _, cookie := range cookies {
		if cookieName == cookie.Name {
			if cookie.Value == r.Context().Value(cookieName).(string) {
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
			Path:     r.URL.Path,
		})
		r.Context().Value(cookieName)
		w.Write([]byte(user + " has been remembered with cookie: " + cookieName + " whose value is: " + rememberMeKey + "<br/>"))
	}

	w.Write([]byte("Weak Randomness Test math/rand executed"))
}

func doSomething(r *http.Request, param string) string {
	bar := "safe!"
	map59148 := make(map[string]interface{})
	map59148["keyA-59148"] = "a_Value"
	map59148["keyB-59148"] = param
	map59148["keyC"] = "another_Value"
	bar = map59148["keyB-59148"].(string)
	bar = map59148["keyA-59148"].(string)

	return bar
}
