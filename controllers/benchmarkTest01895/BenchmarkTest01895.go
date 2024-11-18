package controllers

import (
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01895Controller struct {
	http.ResponseWriter
	*http.Request
}

func (c *BenchmarkTest01895Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.ResponseWriter = w
	c.Request = r
	if r.Method == http.MethodGet {
		c.doPost()
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (c *BenchmarkTest01895Controller) doPost() {
	w := c.ResponseWriter
	r := c.Request

	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := r.Header.Get("BenchmarkTest01895")
	param, _ = url.QueryUnescape(param)

	bar := doSomething(param)

	block, err := des.NewCipher([]byte("12345678"))
	if err != nil {
		http.Error(w, "Problem executing crypto", http.StatusInternalServerError)
		return
	}

	iv := make([]byte, block.BlockSize())
	if _, err := rand.Read(iv); err != nil {
		http.Error(w, "Problem executing crypto", http.StatusInternalServerError)
		return
	}
	ciphertext := make([]byte, len(bar))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, []byte(bar))

	file, err := os.OpenFile("passwordFile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(w, "Problem opening file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	if _, err := file.WriteString("secret_value=" + base64.StdEncoding.EncodeToString(ciphertext) + "\n"); err != nil {
		http.Error(w, "Problem writing to file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Sensitive value: '%s' encrypted and stored<br/>", bar)
}

func doSomething(param string) string {
	bar := ""
	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}
	return bar
}
