package controllers

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"html"
	"net/http"
	"os"
)

type BenchmarkTest00703 struct{}

func (b *BenchmarkTest00703) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	if r.Method != http.MethodPost {
		http.Error(w, "This endpoint only accepts POST requests", http.StatusMethodNotAllowed)
		return
	}

	values := r.FormValue("BenchmarkTest00703")
	var param string
	if values != "" {
		param = values
	}

	var bar string
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the param value
	}

	hasher := sha1.New()
	input := []byte(bar)
	hasher.Write(input)

	result := hasher.Sum(nil)
	fileTarget := "passwordFile.txt"
	f, err := os.OpenFile(fileTarget, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(w, "Failed to open file", http.StatusInternalServerError)
		return
	}
	defer f.Close()

	_, err = f.WriteString(fmt.Sprintf("hash_value=%s\n", base64.StdEncoding.EncodeToString(result)))
	if err != nil {
		http.Error(w, "Failed to write to file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Sensitive value '%s' hashed and stored<br/>", html.EscapeString(string(input)))
	fmt.Fprintln(w, "Hash Test executed")
}
