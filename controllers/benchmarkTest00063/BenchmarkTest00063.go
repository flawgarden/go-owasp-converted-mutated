package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type BenchmarkTest00063 struct{}

func (b *BenchmarkTest00063) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.SetCookie(w, &http.Cookie{
			Name:   "BenchmarkTest00063",
			Value:  "FileName",
			MaxAge: 60 * 3,
			Secure: true,
			Path:   r.RequestURI,
		})
		http.ServeFile(w, r, "pathtraver-00/BenchmarkTest00063.html")
		return
	}

	if r.Method == http.MethodPost {
		cookie, err := r.Cookie("BenchmarkTest00063")
		param := "noCookieValueSupplied"
		if err == nil {
			param = cookie.Value
		}

		bar := param
		num := 106
		if (7*18)+num > 200 {
			bar = "This_should_always_happen"
		}

		fileName := filepath.Join("testfiles", bar)
		fis, err := os.Open(fileName)
		if err != nil {
			fmt.Fprintf(w, "Couldn't open FileInputStream on file: '%s'", fileName)
			return
		}
		defer fis.Close()

		b := make([]byte, 1000)
		size, err := fis.Read(b)
		if err != nil {
			fmt.Fprintf(w, "Problem reading FileInputStream: %s", err.Error())
			return
		}

		w.Write([]byte(fmt.Sprintf("The beginning of file: '%s' is:\n\n%s",
			htmlEscape(fileName),
			htmlEscape(string(b[:size])))))
	}
}

func htmlEscape(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, "&", "&amp;"), "<", "&lt;")
}
