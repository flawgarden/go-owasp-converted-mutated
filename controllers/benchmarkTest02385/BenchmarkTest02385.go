package controllers

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"html"
	"net/http"
	"os"
)

type BenchmarkTest02385Controller struct {
	http.Handler
}

func (c *BenchmarkTest02385Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		c.doPost(w, r)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (c *BenchmarkTest02385Controller) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest02385")
	if param == "" {
		param = ""
	}

	bar := doSomething(param)

	hash := sha1.New()
	input := []byte(bar)
	hash.Write(input)
	result := hash.Sum(nil)

	fileTarget, err := os.OpenFile("passwordFile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(w, "Could not open file", http.StatusInternalServerError)
		return
	}
	defer fileTarget.Close()

	if _, err := fileTarget.WriteString("hash_value=" + base64.StdEncoding.EncodeToString(result) + "\n"); err != nil {
		http.Error(w, "Could not write to file", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("Sensitive value '%s' hashed and stored<br/>", htmlEscape(string(input)))))
	w.Write([]byte("Hash Test executed"))
}

func doSomething(param string) string {
	bar := param
	num := 106
	if (7*18)+num > 200 {
		bar = "This_should_always_happen"
	}
	return bar
}

func htmlEscape(input string) string {
	return html.EscapeString(input)
}
