package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"text/template"
)

type BenchmarkTest01872 struct{}

func (b *BenchmarkTest01872) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.SetCookie(w, &http.Cookie{
			Name:   "BenchmarkTest01872",
			Value:  "color",
			MaxAge: 60 * 3,
			Secure: true,
			Path:   r.URL.Path,
			Domain: r.Host,
		})
		http.ServeFile(w, r, "trustbound-01/BenchmarkTest01872.html")
		return
	}

	if r.Method == http.MethodPost {
		cookies := r.Cookies()
		param := "noCookieValueSupplied"
		for _, cookie := range cookies {
			if cookie.Name == "BenchmarkTest01872" {
				param, _ = url.QueryUnescape(cookie.Value)
				break
			}
		}

		bar := doSomething(param)
		r.AddCookie(&http.Cookie{Name: bar, Value: "10340"})
		fmt.Fprintf(w, "Item: '%s' with value: 10340 saved in session.", htmlEscape(bar))
	}
}

func doSomething(param string) string {
	// Implementation of the doSomething logic.
	return param
}

func htmlEscape(s string) string {
	return template.HTMLEscapeString(s)
}
