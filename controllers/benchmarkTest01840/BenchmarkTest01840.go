package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type BenchmarkTest01840 struct{}

func (b *BenchmarkTest01840) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	http.SetCookie(w, &http.Cookie{
		Name:   "BenchmarkTest01840",
		Value:  "FileName",
		MaxAge: 60 * 3, // Store cookie for 3 minutes
		Secure: true,
		Path:   r.RequestURI,
		Domain: r.Host,
	})

	http.ServeFile(w, r, "pathtraver-02/BenchmarkTest01840.html")
}

func (b *BenchmarkTest01840) Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	var param string = "noCookieValueSupplied"
	cookies := r.Cookies()
	for _, cookie := range cookies {
		if cookie.Name == "BenchmarkTest01840" {
			param = cookie.Value
			break
		}
	}

	bar := doSomething(r, param)
	var fileName string
	var fos *os.File

	defer func() {
		if fos != nil {
			fos.Close()
		}
	}()

	fileName = filepath.Join("testfiles", bar)
	fos, _ = os.Create(fileName)
	if fos != nil {
		fmt.Fprintf(w, "Now ready to write to file: %s", htmlEscape(fileName))
	} else {
		fmt.Println("Couldn't open FileOutputStream on file:", fileName)
	}
}

func doSomething(r *http.Request, param string) string {
	return param
}

func htmlEscape(str string) string {
	return strings.ReplaceAll(strings.ReplaceAll(str, "&", "&amp;"), "<", "&lt;") // Simple HTML escaping
}
