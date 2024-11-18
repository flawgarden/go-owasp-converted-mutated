package controllers

import (
	"fmt"
	"net/http"
	"net/url"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01824 struct {
}

func (bt *BenchmarkTest01824) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		bt.handleGet(w, r)
	case http.MethodPost:
		bt.handlePost(w, r)
	}
}

func (bt *BenchmarkTest01824) handleGet(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "BenchmarkTest01824",
		Value:  "someSecret",
		Path:   r.URL.Path,
		MaxAge: 180,
		Secure: true,
	})
	http.Redirect(w, r, "/crypto-02/BenchmarkTest01824.html", http.StatusSeeOther)
}

func (bt *BenchmarkTest01824) handlePost(w http.ResponseWriter, r *http.Request) {
	cookies := r.Cookies()
	param := "noCookieValueSupplied"
	for _, cookie := range cookies {
		if cookie.Name == "BenchmarkTest01824" {
			param, _ = url.QueryUnescape(cookie.Value)
			break
		}
	}

	bar := doSomething(r, param)

	// ... (AES encryption logic should be here)

	fmt.Fprintf(w, "Sensitive value: '%s' encrypted and stored<br/>", bar)
}

func doSomething(r *http.Request, param string) string {
	return param
}
