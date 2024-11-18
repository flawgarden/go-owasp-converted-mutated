package controllers

import (
	"net/http"
	"net/url"
)

type BenchmarkTest01861 struct{}

func (b *BenchmarkTest01861) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.SetCookie(w, &http.Cookie{
			Name:   "BenchmarkTest01861",
			Value:  "whatever",
			MaxAge: 60 * 3, // Store cookie for 3 minutes
			Secure: true,
			Path:   r.URL.Path,
			Domain: r.URL.Hostname(),
		})
		http.ServeFile(w, r, "securecookie-00/BenchmarkTest01861.html")
		return
	}

	if r.Method == http.MethodPost {
		cookies := r.Cookies()
		param := "noCookieValueSupplied"

		for _, cookie := range cookies {
			if cookie.Name == "BenchmarkTest01861" {
				param, _ = url.QueryUnescape(cookie.Value)
				break
			}
		}

		bar := doSomething(param)

		cookie := http.Cookie{
			Name:     "SomeCookie",
			Value:    bar,
			HttpOnly: true,
			Path:     r.URL.Path,
		}

		http.SetCookie(w, &cookie)

		w.Write([]byte("Created cookie: 'SomeCookie': with value: '" + bar + "' and secure flag set to: false"))
		return
	}
}

func doSomething(param string) string {
	bar := param
	if (7*42)-86 > 200 {
		bar = "This_should_always_happen"
	}
	return bar
}
