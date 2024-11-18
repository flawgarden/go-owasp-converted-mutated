package controllers

import (
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"html"
	"io"
	"net/http"
	"os"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01822Controller struct {
	web.Controller
}

func (c *BenchmarkTest01822Controller) Get() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	http.SetCookie(c.Ctx.ResponseWriter, &http.Cookie{
		Name:   "BenchmarkTest01822",
		Value:  "someSecret",
		MaxAge: 60 * 3,
		Secure: true,
		Path:   c.Ctx.Request.URL.Path,
		Domain: c.Ctx.Request.URL.Hostname(),
	})
	http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, "crypto-02/BenchmarkTest01822.html")
}

func (c *BenchmarkTest01822Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	cookies := c.Ctx.Request.Cookies()
	param := "noCookieValueSupplied"
	for _, cookie := range cookies {
		if cookie.Name == "BenchmarkTest01822" {
			param = cookie.Value
			break
		}
	}

	bar := doSomething(c.Ctx.Request, param)

	iv := make([]byte, des.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error generating IV", http.StatusInternalServerError)
		return
	}

	block, err := des.NewCipher([]byte("somekey12")) // Must be 8 bytes for DES
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error creating cipher", http.StatusInternalServerError)
		return
	}

	ciphertext := make([]byte, len(bar))
	cipher.NewCFBEncrypter(block, iv).XORKeyStream(ciphertext, []byte(bar))

	fileTarget := "passwordFile.txt"
	fw, err := os.OpenFile(fileTarget, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error opening file", http.StatusInternalServerError)
		return
	}
	defer fw.Close()

	if _, err := fw.WriteString("secret_value=" + base64.StdEncoding.EncodeToString(ciphertext) + "\n"); err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error writing to file", http.StatusInternalServerError)
		return
	}

	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Sensitive value: '%s' encrypted and stored<br/>", html.EscapeString(bar))))
}

func doSomething(request *http.Request, param string) string {
	bar := "safe!"
	map20550 := map[string]interface{}{
		"keyA-20550": "a-Value",
		"keyB-20550": param,
		"keyC":       "another-Value",
	}
	bar = map20550["keyB-20550"].(string)
	return bar
}
