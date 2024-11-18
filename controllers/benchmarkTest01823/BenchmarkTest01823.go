package controllers

import (
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"encoding/base64"
	"html"
	"net/http"
	"net/url"
	"os"
)

type BenchmarkTest01823Controller struct {
	http.Handler
}

func (c *BenchmarkTest01823Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		c.doGet(w, r)
	case http.MethodPost:
		c.doPost(w, r)
	}
}

func (c *BenchmarkTest01823Controller) doGet(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "BenchmarkTest01823",
		Value:  "someSecret",
		MaxAge: 180, // Store cookie for 3 minutes
		Secure: true,
		Path:   r.RequestURI,
		Domain: r.Host,
	})
	http.ServeFile(w, r, "crypto-02/BenchmarkTest01823.html")
}

func (c *BenchmarkTest01823Controller) doPost(w http.ResponseWriter, r *http.Request) {
	cookies := r.Cookies()

	param := "noCookieValueSupplied"
	for _, cookie := range cookies {
		if cookie.Name == "BenchmarkTest01823" {
			param, _ = url.QueryUnescape(cookie.Value)
			break
		}
	}

	bar := doSomething(param)

	key := make([]byte, 8) // DES requires 8 byte keys
	_, _ = rand.Read(key)

	block, err := des.NewCipher(key)
	if err != nil {
		http.Error(w, "Error creating cipher", http.StatusInternalServerError)
		return
	}

	iv := make([]byte, des.BlockSize)
	_, _ = rand.Read(iv)

	ciphertext := make([]byte, len(bar))
	encryptor := cipher.NewCBCEncrypter(block, iv)
	encryptor.CryptBlocks(ciphertext, []byte(bar))

	file, _ := os.OpenFile("passwordFile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
	file.WriteString("secret_value=" + base64.StdEncoding.EncodeToString(ciphertext) + "\n")

	output := "Sensitive value: '" + html.EscapeString(bar) + "' encrypted and stored<br/>"
	_, _ = w.Write([]byte(output))
}

func doSomething(param string) string {
	num := 86
	if (7*42)-num > 200 {
		return "This_should_always_happen"
	}
	return param
}
