package controllers

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"html"
	"net/http"
	"os"
	"strings"
)

type BenchmarkTest01844 struct {
}

func (bt *BenchmarkTest01844) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.SetCookie(w, &http.Cookie{
			Name:   "BenchmarkTest01844",
			Value:  "someSecret",
			Path:   r.RequestURI,
			MaxAge: 60 * 3,
			Secure: true,
		})
		http.ServeFile(w, r, "hash-02/BenchmarkTest01844.html")
		return
	}

	if r.Method == http.MethodPost {
		cookie, err := r.Cookie("BenchmarkTest01844")
		param := "noCookieValueSupplied"
		if err == nil {
			param = cookie.Value
		}

		bar := doSomething(param)
		hashValue := hashSHA1(bar)

		fileTarget := "passwordFile.txt"
		fw, err := os.OpenFile(fileTarget, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err == nil {
			defer fw.Close()
			_, _ = fw.WriteString(fmt.Sprintf("hash_value=%s\n", base64.StdEncoding.EncodeToString(hashValue)))
		}
		fmt.Fprintf(w, "Sensitive value '%s' hashed and stored<br/>", html.EscapeString(bar))
		fmt.Fprint(w, "Hash Test executed")
	}
}

func doSomething(param string) string {
	b := strings.Builder{}
	b.WriteString(param)
	b.WriteString(" SafeStuff")
	safeContent := b.String()[:len(b.String())-1]
	return safeContent
}

func hashSHA1(input string) []byte {
	h := sha1.New()
	h.Write([]byte(input))
	return h.Sum(nil)
}
