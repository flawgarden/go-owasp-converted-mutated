package controllers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"io"
	"net/http"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02460Controller struct {
	web.Controller
}

func (c *BenchmarkTest02460Controller) Get() {
	c.DoPost()
}

func (c *BenchmarkTest02460Controller) Post() {
	c.DoPost()
}

func (c *BenchmarkTest02460Controller) DoPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02460")
	bar := doSomething(param)

	iv := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		http.Error(c.Ctx.ResponseWriter, "Failed to generate IV", http.StatusInternalServerError)
		return
	}

	key := []byte("examplekey123456") // 16 bytes for AES-128
	block, err := aes.NewCipher(key)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Failed to create cipher", http.StatusInternalServerError)
		return
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Failed to create GCM", http.StatusInternalServerError)
		return
	}

	ciphertext := gcm.Seal(nil, iv, []byte(bar), nil)

	encoded := base64.StdEncoding.EncodeToString(ciphertext)
	_, err = c.Ctx.ResponseWriter.Write([]byte("Sensitive value: '" + encoded + "' encrypted and stored<br/>"))
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Failed to write response", http.StatusInternalServerError)
		return
	}
}

func doSomething(param string) string {
	bar := ""
	guess := "ABC"
	switchTarget := guess[2]

	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bobs_your_uncle"
	case 'C', 'D':
		bar = param
	default:
		bar = "bobs_your_uncle"
	}

	return bar
}

func init() {
	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Any other initialization logic can go here
}
