package controllers

import (
	"net/http"
	"os"
	"path/filepath"
)

type BenchmarkTest00955 struct{}

func (b *BenchmarkTest00955) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	userCookie := http.Cookie{Name: "BenchmarkTest00955", Value: "FileName", MaxAge: 180, Secure: true, Path: r.URL.Path, Domain: r.Host}
	http.SetCookie(w, &userCookie)
	http.ServeFile(w, r, "/pathtraver-01/BenchmarkTest00955.html")
}

func (b *BenchmarkTest00955) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	cookies := r.Cookies()
	param := "noCookieValueSupplied"
	for _, cookie := range cookies {
		if cookie.Name == "BenchmarkTest00955" {
			param = cookie.Value
			break
		}
	}

	bar := b.doSomething(r, param)
	fileName := filepath.Join("/path/to/testfiles", bar)

	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		http.Error(w, "Couldn't open FileOutputStream on file: '"+fileName+"'", http.StatusInternalServerError)
		return
	}
	defer f.Close()

	w.Write([]byte("Now ready to write to file: " + fileName))
}

func (b *BenchmarkTest00955) doSomething(r *http.Request, param string) string {
	bar := param
	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	}
	return bar
}
