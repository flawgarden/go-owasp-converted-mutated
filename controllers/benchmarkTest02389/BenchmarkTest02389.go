package controllers

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"html"
	"net/http"
	"os"

	"github.com/beego/beego/v2/server/web"
)

type HashTestController struct {
	web.Controller
}

func (c *HashTestController) Get() {
	c.Post()
}

func (c *HashTestController) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest02389")
	if param == "" {
		param = ""
	}

	bar := doSomething(param)

	hash := sha256.New()
	input := []byte(bar)
	hash.Write(input)
	result := hash.Sum(nil)

	file, err := os.OpenFile("passwordFile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error opening file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	if _, err := file.WriteString("hash_value=" + base64.StdEncoding.EncodeToString(result) + "\n"); err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error writing to file", http.StatusInternalServerError)
		return
	}

	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Sensitive value '%s' hashed and stored<br/>", html.EscapeString(string(input)))))
	c.Ctx.ResponseWriter.Write([]byte("Hash Test executed"))
}

func doSomething(param string) string {
	// Placeholder for actual functionality
	return param
}
