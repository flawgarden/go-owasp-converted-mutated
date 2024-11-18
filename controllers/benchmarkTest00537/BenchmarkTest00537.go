package controllers

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type BenchmarkTest00537 struct{}

func (b *BenchmarkTest00537) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		b.doPost(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (b *BenchmarkTest00537) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	flag := true
	names := r.URL.Query()

	for name, values := range names {
		if flag {
			for _, value := range values {
				if value == "BenchmarkTest00537" {
					param = name
					flag = false
					break
				}
			}
		}
	}

	bar := htmlEscape(param)

	var input []byte
	if inputParam := bar; inputParam != "" {
		input = []byte(inputParam)
	}

	hash := md5.New()
	hash.Write(input)
	result := hash.Sum(nil)

	fileTarget, err := os.OpenFile("passwordFile.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		http.Error(w, "Unable to open file", http.StatusInternalServerError)
		return
	}
	defer fileTarget.Close()

	if _, err := fileTarget.WriteString("hash_value=" + base64.StdEncoding.EncodeToString(result) + "\n"); err != nil {
		http.Error(w, "Unable to write to file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Sensitive value '%s' hashed and stored<br/>", htmlEscape(string(input)))

	fmt.Fprintln(w, "Hash Test executed")
}

func htmlEscape(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, "&", "&amp;"), "<", "&lt;")
}
