package controllers

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"html"
	"net/http"
	"os"
)

type BenchmarkTest01766 struct{}

func (b *BenchmarkTest01766) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest01766")

	bar := b.doSomething(param)

	hash := sha512.Sum512([]byte(bar))
	hashValue := base64.StdEncoding.EncodeToString(hash[:])

	file, err := os.OpenFile("passwordFile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(w, "Unable to open file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	if _, err := file.WriteString(fmt.Sprintf("hash_value=%s\n", hashValue)); err != nil {
		http.Error(w, "Unable to write file", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("Sensitive value '%s' hashed and stored<br/>", html.EscapeString(bar))))
}

func (b *BenchmarkTest01766) doSomething(param string) string {
	if param != "" && len(param) > 1 {
		return param[:len(param)-1]
	}
	return param
}
