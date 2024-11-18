package controllers

import (
	"crypto/des"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"os"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02102Controller struct {
	beego.Controller
}

func (c *BenchmarkTest02102Controller) Get() {
	cantaloupe := c.GetString("BenchmarkTest02102")
	if cantaloupe == "" {
		cantaloupe = ""
	}

	bar := doSomething(cantaloupe)

	block, err := des.NewCipher([]byte("12345678")) // Example key
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Encrypt
	ciphertext := make([]byte, len(bar))
	block.Encrypt(ciphertext, []byte(bar))

	fileTarget := "passwordFile.txt"
	fw, err := os.OpenFile(fileTarget, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer fw.Close()

	encoded := base64.StdEncoding.EncodeToString(ciphertext)
	if _, err := fw.WriteString("secret_value=" + encoded + "\n"); err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}

	output, err := json.Marshal("Sensitive value: '" + bar + "' encrypted and stored")
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}
	c.Ctx.ResponseWriter.Write(output)
}

func doSomething(param string) string {
	bar := param
	if param != "" && len(param) > 1 {
		sbxyz85369 := strings.Builder{}
		sbxyz85369.WriteString(param)
		bar = sbxyz85369.String()[:len(param)-1] + "Z"
	}
	return bar
}
