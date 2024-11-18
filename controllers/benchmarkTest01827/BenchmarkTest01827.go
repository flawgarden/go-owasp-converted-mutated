package controllers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/crypto-02/BenchmarkTest01827", benchmarkTest01827)
	http.ListenAndServe(":8080", nil)
}

func benchmarkTest01827(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.SetCookie(w, &http.Cookie{
			Name:   "BenchmarkTest01827",
			Value:  "someSecret",
			Path:   r.RequestURI,
			Domain: r.URL.Hostname(),
			MaxAge: 60 * 3,
			Secure: true,
		})
		http.ServeFile(w, r, "crypto-02/BenchmarkTest01827.html")
		return
	}

	cookies := r.Cookies()
	param := "noCookieValueSupplied"
	for _, cookie := range cookies {
		if cookie.Name == "BenchmarkTest01827" {
			param = cookie.Value
			break
		}
	}

	bar := doSomething(param)

	block, err := aes.NewCipher([]byte("examplekey123456")) // 16 bytes key for AES-128
	if err != nil {
		http.Error(w, "Error initializing cipher", http.StatusInternalServerError)
		return
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		http.Error(w, "Error generating IV", http.StatusInternalServerError)
		return
	}

	ciphertext := encrypt(block, iv, []byte(bar))
	encoded := base64.StdEncoding.EncodeToString(ciphertext)

	fmt.Fprintf(w, "Sensitive value: '%s' encrypted and stored\n", encoded)
}

func doSomething(param string) string {
	switchTarget := 'B' // condition 'B', which is safe
	var bar string

	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bob"
	case 'C', 'D':
		bar = param
	default:
		bar = "bob's your uncle"
	}

	return bar
}

func encrypt(block cipher.Block, iv []byte, plaintext []byte) []byte {
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	copy(ciphertext[:aes.BlockSize], iv)

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	return ciphertext
}
