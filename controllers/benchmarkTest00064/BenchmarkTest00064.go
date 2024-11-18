package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func BenchmarkTest00064(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	cookie := http.Cookie{
		Name:   "BenchmarkTest00064",
		Value:  "FileName",
		Path:   r.RequestURI,
		MaxAge: 60 * 3,
		Secure: true,
	}
	http.SetCookie(w, &cookie)

	http.ServeFile(w, r, "./pathtraver-00/BenchmarkTest00064.html")
}

func BenchmarkTest00064Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	cookies := r.Cookies()
	param := "noCookieValueSupplied"
	for _, cookie := range cookies {
		if cookie.Name == "BenchmarkTest00064" {
			param = cookie.Value
			break
		}
	}

	var bar string
	num := 106
	if (7*18)+num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}

	fileName := filepath.Join("path/to/test/files", bar)
	fos, err := os.Create(fileName)
	if err != nil {
		fmt.Fprintf(w, "Couldn't open FileOutputStream on file: '%s'", fileName)
		return
	}
	defer fos.Close()

	fmt.Fprintf(w, "Now ready to write to file: %s", escapeHTML(fileName))
}

func escapeHTML(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, "&", "&amp;"), "<", "&lt;")
}

func main() {
	http.HandleFunc("/pathtraver-00/BenchmarkTest00064", BenchmarkTest00064)
	http.HandleFunc("/pathtraver-00/BenchmarkTest00064Post", BenchmarkTest00064Post)
	http.ListenAndServe(":8080", nil)
}
