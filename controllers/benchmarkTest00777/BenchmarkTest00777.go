package controllers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

func main() {
	http.HandleFunc("/crypto-01/BenchmarkTest00777", BenchmarkTest00777)
	http.ListenAndServe(":8080", nil)
}

func BenchmarkTest00777(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	queryString := r.URL.RawQuery
	paramval := "BenchmarkTest00777="
	paramLoc := -1
	if queryString != "" {
		paramLoc = len(queryString) - len(paramval)
	}

	param := ""
	if paramLoc != -1 {
		param = queryString[paramLoc+len(paramval):]
	}

	param = url.QueryEscape(param)

	bar := "safe!"
	map4544 := make(map[string]interface{})
	map4544["keyA-4544"] = "a_Value"
	map4544["keyB-4544"] = param
	map4544["keyC"] = "another_Value"
	bar = map4544["keyB-4544"].(string)
	bar = map4544["keyA-4544"].(string)

	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		http.Error(w, "Problem generating random bytes", http.StatusInternalServerError)
		return
	}

	key := make([]byte, 32) // AES-256
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		http.Error(w, "Problem generating key", http.StatusInternalServerError)
		return
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		http.Error(w, "Problem creating cipher", http.StatusInternalServerError)
		return
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		http.Error(w, "Problem creating GCM", http.StatusInternalServerError)
		return
	}

	input := []byte(bar)
	result := aesGCM.Seal(nil, iv, input, nil)

	fileName := "passwordFile.txt"
	fw, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		http.Error(w, "Problem opening file", http.StatusInternalServerError)
		return
	}
	defer fw.Close()

	if _, err := fw.WriteString("secret_value=" + base64.StdEncoding.EncodeToString(result) + "\n"); err != nil {
		http.Error(w, "Problem writing to file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Sensitive value: '%s' encrypted and stored<br/>", input)
}
