package controllers

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"html"
	"net/http"
	"os"
	"strings"
)

type BenchmarkTest01578 struct{}

func (b *BenchmarkTest01578) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	if r.Method == http.MethodGet {
		b.doPost(w, r)
		return
	}
	b.doPost(w, r)
}

func (b *BenchmarkTest01578) doPost(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()["BenchmarkTest01578"]
	var param string
	if len(values) > 0 {
		param = values[0]
	} else {
		param = ""
	}

	bar := b.doSomething(param)

	hash := sha256.New()
	input := []byte(bar)
	hash.Write(input)
	result := hash.Sum(nil)

	fileTarget := "passwordFile.txt"
	fw, err := os.OpenFile(fileTarget, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(w, "Error opening file", http.StatusInternalServerError)
		return
	}
	defer fw.Close()

	_, err = fw.WriteString("hash_value=" + base64.StdEncoding.EncodeToString(result) + "\n")
	if err != nil {
		http.Error(w, "Error writing to file", http.StatusInternalServerError)
		return
	}

	output := fmt.Sprintf("Sensitive value '%s' hashed and stored<br/>", html.EscapeString(string(input)))
	_, _ = w.Write([]byte(output))
	_, _ = w.Write([]byte("Hash Test executed"))
}

func (b *BenchmarkTest01578) doSomething(param string) string {
	var sbxyz85952 strings.Builder
	sbxyz85952.WriteString(param)
	bar := sbxyz85952.String() + "_SafeStuff"
	return bar
}
