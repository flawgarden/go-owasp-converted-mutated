package controllers

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html"
	"net/http"
	"os"
)

type BenchmarkTest00022 struct{}

func (b *BenchmarkTest00022) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest00022")
	if param == "" {
		param = ""
	}

	hash := sha256.New()
	input := []byte(param)
	hash.Write(input)
	result := hash.Sum(nil)

	fileTarget := "passwordFile.txt"
	fw, err := os.OpenFile(fileTarget, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(w, "Unable to open file", http.StatusInternalServerError)
		return
	}
	defer fw.Close()

	if _, err := fw.WriteString("hash_value=" + base64.StdEncoding.EncodeToString(result) + "\n"); err != nil {
		http.Error(w, "Unable to write to file", http.StatusInternalServerError)
		return
	}

	output, err := json.Marshal(map[string]string{
		"message": fmt.Sprintf("Sensitive value '%s' hashed and stored<br/>", html.EscapeString(param)),
	})
	if err != nil {
		http.Error(w, "Unable to create JSON output", http.StatusInternalServerError)
		return
	}

	w.Write(output)
	w.Write([]byte("Hash Test executed"))
}
