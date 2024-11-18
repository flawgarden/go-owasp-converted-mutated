package controllers

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/hash-00/BenchmarkTest00704", BenchmarkTest00704)
	http.ListenAndServe(":8080", nil)
}

func BenchmarkTest00704(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		doPost(w, r)
	} else if r.Method == http.MethodPost {
		doPost(w, r)
	}
}

func doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")

	values := r.URL.Query()["BenchmarkTest00704"]
	var param string
	if len(values) > 0 {
		param = values[0]
	} else {
		param = ""
	}

	bar := param
	if len(param) > 1 {
		bar = param[:len(param)-1]
	}

	var input []byte
	if len(bar) > 0 {
		input = []byte(bar)
	} else {
		input = []byte("?")
	}

	hash := sha1.New()
	hash.Write(input)
	result := hash.Sum(nil)

	file, err := os.OpenFile("passwordFile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(w, "Could not open file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	if _, err = file.WriteString("hash_value=" + base64.StdEncoding.EncodeToString(result) + "\n"); err != nil {
		http.Error(w, "Could not write to file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Sensitive value '%s' hashed and stored<br/>", htmlEscape(string(input)))
	fmt.Fprintf(w, "Hash Test executed")
}

func htmlEscape(s string) string {
	return fmt.Sprintf("%q", s)
}
