package controllers

import (
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"os"
)

type BenchmarkTest00707 struct{}

func (b *BenchmarkTest00707) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest00707")
	bar := ""

	num := 106

	if (7*18)+num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}

	input := []byte(bar)
	hash := sha512.Sum512(input)

	file, err := os.OpenFile("passwordFile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(w, "Unable to open file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	_, err = file.WriteString("hash_value=" + base64.StdEncoding.EncodeToString(hash[:]) + "\n")
	if err != nil {
		http.Error(w, "Unable to write to file", http.StatusInternalServerError)
		return
	}

	output := map[string]string{"message": "Sensitive value '" + bar + "' hashed and stored"}
	json.NewEncoder(w).Encode(output)
}

func main() {
	http.Handle("/hash-00/BenchmarkTest00707", &BenchmarkTest00707{})
	http.ListenAndServe(":8080", nil)
}
