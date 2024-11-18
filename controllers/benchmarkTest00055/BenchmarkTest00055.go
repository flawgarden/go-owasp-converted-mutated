package controllers

import (
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/crypto-00/BenchmarkTest00055", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			http.SetCookie(w, &http.Cookie{
				Name:   "BenchmarkTest00055",
				Value:  "someSecret",
				Path:   r.URL.Path,
				MaxAge: 180, // Store cookie for 3 minutes
				Secure: true,
			})
			http.ServeFile(w, r, "crypto-00/BenchmarkTest00055.html")
		} else if r.Method == http.MethodPost {
			r.ParseForm()
			cookies := r.Cookies()
			param := "noCookieValueSupplied"
			for _, cookie := range cookies {
				if cookie.Name == "BenchmarkTest00055" {
					param = cookie.Value
					break
				}
			}

			bar := param
			block, err := des.NewCipher([]byte("12345678"))
			if err != nil {
				http.Error(w, "Error creating cipher", http.StatusInternalServerError)
				return
			}
			iv := make([]byte, block.BlockSize())
			if _, err := rand.Read(iv); err != nil {
				http.Error(w, "Error generating IV", http.StatusInternalServerError)
				return
			}

			ciphertext := make([]byte, len(bar))
			mode := cipher.NewCBCEncrypter(block, iv)
			mode.CryptBlocks(ciphertext, []byte(bar))

			encoded := base64.StdEncoding.EncodeToString(ciphertext)

			fileTarget := "passwordFile.txt"
			f, err := os.OpenFile(fileTarget, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
			if err != nil {
				http.Error(w, "Error writing to file", http.StatusInternalServerError)
				return
			}
			defer f.Close()
			if _, err := f.WriteString("secret_value=" + encoded + "\n"); err != nil {
				http.Error(w, "Error writing to file", http.StatusInternalServerError)
				return
			}

			fmt.Fprintf(w, "Sensitive value: '%s' encrypted and stored<br/>", bar)
		}
	})
	http.ListenAndServe(":8080", nil)
}
