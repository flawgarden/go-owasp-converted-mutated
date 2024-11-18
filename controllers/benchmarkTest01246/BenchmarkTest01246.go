package controllers

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"html"
	"net/http"
	"os"
)

type BenchmarkTest01246 struct{}

func (b *BenchmarkTest01246) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest01246")
	if param == "" {
		param = ""
	}

	bar := b.doSomething(param)

	hash := sha1.New()
	input := []byte(bar)
	_, err := hash.Write(input)
	if err != nil {
		http.Error(w, "Error hashing input", http.StatusInternalServerError)
		return
	}
	result := hash.Sum(nil)

	fileTarget := "passwordFile.txt"
	fw, err := os.OpenFile(fileTarget, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		http.Error(w, "Error opening file", http.StatusInternalServerError)
		return
	}
	defer fw.Close()

	if _, err := fw.WriteString("hash_value=" + base64.StdEncoding.EncodeToString(result) + "\n"); err != nil {
		http.Error(w, "Error writing to file", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Sensitive value '%s' hashed and stored<br/>", htmlEscape(string(input)))

	fmt.Fprintln(w, "Hash Test executed")
}

func (b *BenchmarkTest01246) doSomething(param string) string {
	bar := ""
	guess := "ABC"
	switchTarget := guess[2]

	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bobs_your_uncle"
	case 'C', 'D':
		bar = param
	default:
		bar = "bobs_your_uncle"
	}
	return bar
}

func htmlEscape(s string) string {
	return html.EscapeString(s)
}

func main() {
	http.Handle("/", &BenchmarkTest01246{})
	http.ListenAndServe(":8080", nil)
}
