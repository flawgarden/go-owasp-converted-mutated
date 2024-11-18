package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

type BenchmarkTest01835 struct{}

func (b *BenchmarkTest01835) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	userCookie := http.Cookie{
		Name:   "BenchmarkTest01835",
		Value:  "FileName",
		MaxAge: 60 * 3,
		Secure: true,
		Path:   r.URL.Path,
		Domain: getDomain(r),
	}
	http.SetCookie(w, &userCookie)

	http.ServeFile(w, r, "pathtraver-02/BenchmarkTest01835.html")
}

func (b *BenchmarkTest01835) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	cookies := r.Cookies()
	param := "noCookieValueSupplied"
	for _, cookie := range cookies {
		if cookie.Name == "BenchmarkTest01835" {
			param, _ = url.QueryUnescape(cookie.Value)
			break
		}
	}

	bar := doSomething(r, param)
	fileTarget := fmt.Sprintf("%s/%s", os.Getenv("TESTFILES_DIR"), bar)
	w.Write([]byte(fmt.Sprintf("Access to file: '%s' created.", fileTarget)))
	if _, err := os.Stat(fileTarget); os.IsNotExist(err) {
		w.Write([]byte(" But file doesn't exist yet."))
	} else {
		w.Write([]byte(" And file already exists."))
	}
}

func doSomething(r *http.Request, param string) string {
	bar := "This should never happen"
	num := 196
	if (500/42)+num > 200 {
		bar = param
	}
	return bar
}

func getDomain(r *http.Request) string {
	return r.Host
}
