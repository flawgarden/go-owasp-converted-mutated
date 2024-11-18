package controllers

import (
	"crypto/sha512"
	"encoding/base64"
	"net/http"
	"os"
	"strings"
)

type BenchmarkTest00074 struct{}

func (b *BenchmarkTest00074) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.SetCookie(w, &http.Cookie{
			Name:   "BenchmarkTest00074",
			Value:  "someSecret",
			Path:   r.URL.Path,
			MaxAge: 60 * 3,
			Secure: true,
		})
		http.ServeFile(w, r, "hash-00/BenchmarkTest00074.html")
	} else if r.Method == http.MethodPost {
		cookies := r.Cookies()
		param := "noCookieValueSupplied"
		for _, cookie := range cookies {
			if cookie.Name == "BenchmarkTest00074" {
				param = cookie.Value
				break
			}
		}

		bar := "safe!"
		map98790 := map[string]interface{}{
			"keyA-98790": "a-Value",
			"keyB-98790": param,
			"keyC":       "another-Value",
		}
		bar = map98790["keyB-98790"].(string)

		input := []byte("?")
		if len(bar) > 0 {
			input = []byte(bar)
		}
		hash := sha512.Sum512(input)

		filePath := "passwordFile.txt"
		f, _ := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		defer f.Close()
		hashValue := base64.StdEncoding.EncodeToString(hash[:])
		f.WriteString("hash_value=" + hashValue + "\n")

		response := "Sensitive value '" + htmlEscape(string(input)) + "' hashed and stored<br/>"
		w.Write([]byte(response))
	}
}

func htmlEscape(str string) string {
	escaped := strings.ReplaceAll(str, "&", "&amp;")
	escaped = strings.ReplaceAll(escaped, "<", "&lt;")
	escaped = strings.ReplaceAll(escaped, ">", "&gt;")
	escaped = strings.ReplaceAll(escaped, "\"", "&quot;")
	escaped = strings.ReplaceAll(escaped, "'", "&#39;")
	return escaped
}
