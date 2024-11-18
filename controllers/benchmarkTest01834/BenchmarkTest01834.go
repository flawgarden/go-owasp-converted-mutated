package controllers

import (
	"net/http"
	"net/url"
	"os"
)

type BenchmarkTest01834 struct{}

func (b *BenchmarkTest01834) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	cookie := http.Cookie{
		Name:   "BenchmarkTest01834",
		Value:  "FileName",
		MaxAge: 60 * 3,
		Secure: true,
		Path:   r.RequestURI,
		Domain: r.Host,
	}
	http.SetCookie(w, &cookie)

	http.ServeFile(w, r, "./pathtraver-02/BenchmarkTest01834.html")
}

func (b *BenchmarkTest01834) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	cookies := r.Cookies()
	param := "noCookieValueSupplied"
	for _, cookie := range cookies {
		if cookie.Name == "BenchmarkTest01834" {
			param, _ = url.QueryUnescape(cookie.Value)
			break
		}
	}

	bar := doSomething(param)
	fileTarget := "./testfiles/" + bar
	w.Write([]byte("Access to file: '" + fileTarget + "' created."))

	if _, err := os.Stat(fileTarget); err == nil {
		w.Write([]byte(" And file already exists."))
	} else {
		w.Write([]byte(" But file doesn't exist yet."))
	}
}

func doSomething(param string) string {
	bar := ""
	guess := "ABC"
	switchTarget := guess[1]

	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bob"
	case 'C', 'D':
		bar = param
	default:
		bar = "bob's your uncle"
	}

	return bar
}
