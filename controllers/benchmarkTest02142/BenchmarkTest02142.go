package controllers

import (
	"fmt"
	"net/http"
)

type BenchmarkTest02142 struct{}

func (b *BenchmarkTest02142) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest02142")
	if param == "" {
		param = ""
	}

	bar := doSomething(r, param)

	cookie := http.Cookie{
		Name:     "SomeCookie",
		Value:    bar,
		Secure:   false,
		HttpOnly: true,
		Path:     r.RequestURI,
	}
	http.SetCookie(w, &cookie)

	response := fmt.Sprintf("Created cookie: 'SomeCookie': with value: '%s' and secure flag set to: %t", bar, cookie.Secure)
	w.Write([]byte(response))
}

func doSomething(r *http.Request, param string) string {
	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value

		bar = valuesList[0] // get the last 'safe' value
	}
	return bar
}

func main() {
	http.Handle("/securecookie-00/BenchmarkTest02142", &BenchmarkTest02142{})
	http.ListenAndServe(":8080", nil)
}
