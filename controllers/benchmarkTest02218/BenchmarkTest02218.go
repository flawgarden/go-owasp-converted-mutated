package controllers

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
)

func BenchmarkTest02218(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		doPost(w, r)
		return
	}
	doPost(w, r)
}

func doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest02218")
	bar := doSomething(param)

	hash := sha256.New()
	var input []byte

	if len(bar) > 0 {
		input = []byte(bar)
	}

	hash.Write(input)
	result := hash.Sum(nil)

	fileTarget := "passwordFile.txt"
	fw, err := os.OpenFile(fileTarget, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		http.Error(w, "Error opening the file", http.StatusInternalServerError)
		return
	}
	defer fw.Close()

	if _, err := fw.WriteString("hash_value=" + base64.StdEncoding.EncodeToString(result) + "\n"); err != nil {
		http.Error(w, "Error writing to the file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Sensitive value '%s' hashed and stored<br/>", htmlEscape(string(input)))
	fmt.Fprintln(w, "Hash Test executed")
}

func doSomething(param string) string {
	return htmlEscape(param)
}

func htmlEscape(input string) string {
	return input // нужно использовать библиотеку для экранирования HTML
}
