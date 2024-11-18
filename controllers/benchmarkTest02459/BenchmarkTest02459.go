package controllers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"database/sql"
	"fmt"
	"io"
	"net/http"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type CryptoController struct {
	web.Controller
}

func (c *CryptoController) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02459")
	bar := doSomething(param)

	block, err := aes.NewCipher(generateKey())
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error creating cipher", http.StatusInternalServerError)
		return
	}

	ciphertext := encrypt(block, []byte(bar))

	err = storeEncryptedValue(ciphertext)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error storing value", http.StatusInternalServerError)
		return
	}

	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Sensitive value: '%s' encrypted and stored<br/>", param)))
}

func doSomething(param string) string {
	bar := "safe!"
	map56716 := map[string]string{"keyA-56716": "a-Value", "keyB-56716": param, "keyC": "another-Value"}
	bar = map56716["keyB-56716"]
	return bar
}

func generateKey() []byte {
	return []byte("examplekey12345") // 16 bytes for AES-128
}

func encrypt(block cipher.Block, plaintext []byte) []byte {
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)
	return ciphertext
}

func storeEncryptedValue(data []byte) error {
	db, err := sql.Open("mysql", source)
	if err != nil {
		return err
	}
	defer db.Close()

	sqlStr := "INSERT INTO secrets (data) VALUES (?)"
	_, err = db.Exec(sqlStr, data)
	return err
}
