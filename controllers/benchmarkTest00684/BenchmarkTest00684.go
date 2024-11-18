package controllers

import (
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"

	beego "github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00684Controller struct {
	beego.Controller
}

func (c *BenchmarkTest00684Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00684Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest00684")
	bar := sanitize(param)

	block, err := des.NewCipher(generateKey())
	if err != nil {
		handleError(c.Ctx.ResponseWriter, err)
		return
	}

	iv := make([]byte, block.BlockSize())
	if _, err := rand.Read(iv); err != nil {
		handleError(c.Ctx.ResponseWriter, err)
		return
	}

	cipherText := make([]byte, len(bar))
	stream := cipher.NewCBCEncrypter(block, iv)
	stream.CryptBlocks(cipherText, []byte(bar))

	err = storeInFile(cipherText)
	if err != nil {
		handleError(c.Ctx.ResponseWriter, err)
		return
	}

	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Sensitive value: '%s' encrypted and stored<br/>", htmlEscape(string(cipherText)))))
}

func sanitize(input string) string {
	return input // Add sanitation as per your requirements
}

func generateKey() []byte {
	key := make([]byte, 8) // DES requires an 8-byte key
	rand.Read(key)
	return key
}

func storeInFile(data []byte) error {
	file, err := os.OpenFile("passwordFile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString("secret_value=" + base64.StdEncoding.EncodeToString(data) + "\n")
	return err
}

func htmlEscape(input string) string {
	return input // Implement HTML escaping as needed
}

func handleError(w http.ResponseWriter, err error) {
	http.Error(w, "Problem executing crypto", http.StatusInternalServerError)
}
