package controllers

import (
	"crypto/des"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type BenchmarkTest00258 struct{}

func (b *BenchmarkTest00258) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.Header.Get("BenchmarkTest00258")
	if param == "" {
		http.Error(w, "Missing BenchmarkTest00258 header", http.StatusBadRequest)
		return
	}

	bar := htmlEscape(param)

	iv := make([]byte, des.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		http.Error(w, "Error generating IV", http.StatusInternalServerError)
		return
	}

	c, err := des.NewCipher([]byte("12345678")) // Example key, should not be hardcoded
	if err != nil {
		http.Error(w, "Error creating cipher", http.StatusInternalServerError)
		return
	}

	ciphertext := make([]byte, len(bar))
	c.Encrypt(ciphertext, []byte(bar))
	encoded := base64.StdEncoding.EncodeToString(ciphertext)

	file, err := os.OpenFile("passwordFile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(w, "Error writing to file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	if _, err := file.WriteString("secret_value=" + encoded + "\n"); err != nil {
		http.Error(w, "Error writing to file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Sensitive value: '%s' encrypted and stored<br/>", htmlEscape(string(ciphertext)))
}

func htmlEscape(input string) string {
	escaped := jsonEscape(input)
	return escaped
}

func jsonEscape(input string) string {
	b, _ := json.Marshal(input)
	return string(b[1 : len(b)-1])
}

func main() {
	http.Handle("/crypto-00/BenchmarkTest00258", &BenchmarkTest00258{})
	http.ListenAndServe(":8080", nil)
}
