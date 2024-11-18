package controllers

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type CryptoController struct{}

func (cc *CryptoController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		cc.doPost(w, r)
		return
	}
	cc.doPost(w, r)
}

func (cc *CryptoController) doPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest00353")
	if param == "" {
		param = ""
	}

	bar := "safe!"
	mapData := map[string]interface{}{
		"keyA-13001": "a_Value",
		"keyB-13001": param,
		"keyC":       "another_Value",
	}
	bar = mapData["keyB-13001"].(string)
	bar = mapData["keyA-13001"].(string)

	key, err := aes.NewCipher([]byte("a very secret key"))
	if err != nil {
		http.Error(w, "Error creating cipher", http.StatusInternalServerError)
		return
	}

	nonce := []byte("uniqueNonce12") // Example nonce, should be random and unique
	aesgcm, err := cipher.NewGCM(key)
	if err != nil {
		http.Error(w, "Error creating GCM", http.StatusInternalServerError)
		return
	}

	input := []byte(bar)
	result := aesgcm.Seal(nil, nonce, input, nil)

	file, err := os.OpenFile("passwordFile.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		http.Error(w, "Could not open file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	if _, err = file.WriteString(fmt.Sprintf("secret_value=%s\n", result)); err != nil {
		http.Error(w, "Error writing to file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Sensitive value: '%s' encrypted and stored<br/>", input)
}
