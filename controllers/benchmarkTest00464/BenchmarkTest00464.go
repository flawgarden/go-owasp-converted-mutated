package controllers

import (
	"crypto/sha1"
	"encoding/base64"
	"html"
	"net/http"
	"os"

	beego "github.com/beego/beego/v2/server/web"
)

type HashTestController struct {
	beego.Controller
}

func (c *HashTestController) Get() {
	c.DoPost()
}

func (c *HashTestController) Post() {
	c.DoPost()
}

func (c *HashTestController) DoPost() {
	response := c.Ctx.ResponseWriter
	response.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest00464", "")
	bar := param
	if len(param) > 1 {
		bar = param[:len(param)-1]
	}

	h := sha1.New()
	h.Write([]byte(bar))
	result := h.Sum(nil)

	fileTarget := "passwordFile.txt"
	fw, err := os.OpenFile(fileTarget, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(response, "Unable to open file", http.StatusInternalServerError)
		return
	}
	defer fw.Close()

	if _, err := fw.WriteString("hash_value=" + base64.StdEncoding.EncodeToString(result) + "\n"); err != nil {
		http.Error(response, "Unable to write to file", http.StatusInternalServerError)
		return
	}

	encodedInput := html.EscapeString(string([]byte(bar)))
	response.Write([]byte("Sensitive value '" + encodedInput + "' hashed and stored<br/>"))
	response.Write([]byte("Hash Test executed"))
}
