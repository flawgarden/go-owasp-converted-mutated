package controllers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"os"
	"time"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00054Controller struct {
	web.Controller
}

func (c *BenchmarkTest00054Controller) Get() {
	http.SetCookie(c.Ctx.ResponseWriter, &http.Cookie{
		Name:    "BenchmarkTest00054",
		Value:   "someSecret",
		Expires: time.Now().Add(3 * time.Minute),
		Secure:  true,
		Path:    c.Ctx.Request.RequestURI,
		Domain:  c.Ctx.Request.URL.Hostname(),
	})
	c.Redirect("/crypto-00/BenchmarkTest00054.html", http.StatusFound)
}

func (c *BenchmarkTest00054Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	cookie, err := c.Ctx.Request.Cookie("BenchmarkTest00054")
	param := "noCookieValueSupplied"
	if err == nil {
		param = cookie.Value
	}

	bar := param
	iv := make([]byte, 16)
	if _, err := rand.Read(iv); err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error generating IV", http.StatusInternalServerError)
		return
	}
	block, err := aes.NewCipher([]byte("a very very very secret key")) // use a key of proper length
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error creating cipher", http.StatusInternalServerError)
		return
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error creating GCM", http.StatusInternalServerError)
		return
	}

	ciphertext := aesGCM.Seal(nil, iv, []byte(bar), nil)
	encoded := base64.StdEncoding.EncodeToString(ciphertext)

	// Write to a file (adjust the filepath accordingly)
	file, err := os.OpenFile("passwordFile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error writing to file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	if _, err := file.WriteString("secret_value=" + encoded + "\n"); err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error writing to file", http.StatusInternalServerError)
		return
	}

	c.Ctx.ResponseWriter.Write([]byte("Sensitive value: '" + bar + "' encrypted and stored<br/>"))
}
