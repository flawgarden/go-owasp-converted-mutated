package controllers

import (
	"fmt"
	"net/http"
	"net/url"
)

type BenchmarkTest00169 struct{}

func (b *BenchmarkTest00169) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.Header.Get("BenchmarkTest00169")
	param, _ = url.QueryUnescape(param)

	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value

		bar = valuesList[0] // get the last 'safe' value
	}

	cookie := http.Cookie{
		Name:     "SomeCookie",
		Value:    bar,
		HttpOnly: true,
		Path:     r.URL.Path,
		Secure:   false,
	}
	http.SetCookie(w, &cookie)

	output := fmt.Sprintf("Created cookie: 'SomeCookie': with value: '%s' and secure flag set to: false", bar)
	w.Write([]byte(output))
}

func main() {
	http.Handle("/securecookie-00/BenchmarkTest00169", &BenchmarkTest00169{})
	http.ListenAndServe(":8080", nil)
}
