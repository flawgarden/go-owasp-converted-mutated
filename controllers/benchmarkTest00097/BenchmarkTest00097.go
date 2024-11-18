package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type BenchmarkTest00097 struct{}

func (b *BenchmarkTest00097) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	http.SetCookie(w, &http.Cookie{
		Name:   "BenchmarkTest00097",
		Value:  "color",
		MaxAge: 60 * 3,
		Secure: true,
		Path:   r.URL.Path,
		Domain: strings.Split(r.Host, ":")[0],
	})

	http.ServeFile(w, r, "trustbound-00/BenchmarkTest00097.html")
}

func (b *BenchmarkTest00097) handlePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	cookies := r.Cookies()
	param := "noCookieValueSupplied"
	for _, cookie := range cookies {
		if cookie.Name == "BenchmarkTest00097" {
			value, _ := url.QueryUnescape(cookie.Value)
			param = value
			break
		}
	}

	num := 106
	bar := ""
	if (7*18)+num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}

	r.Context().Value("session").(map[string]interface{})[bar] = "10340"
	fmt.Fprintf(w, "Item: '%s' with value: 10340 saved in session.", htmlEscape(bar))
}

func htmlEscape(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, "&", "&amp;"), "<", "&lt;")
}

func main() {
	benchmark := &BenchmarkTest00097{}
	http.Handle("/trustbound-00/BenchmarkTest00097", benchmark)
	http.ListenAndServe(":8080", nil)
}
