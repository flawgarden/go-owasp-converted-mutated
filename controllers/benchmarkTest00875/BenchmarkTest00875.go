package controllers

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00875Controller struct {
	web.Controller
}

func (c *BenchmarkTest00875Controller) Get() {
	c.doPost()
}

func (c *BenchmarkTest00875Controller) Post() {
	c.doPost()
}

func (c *BenchmarkTest00875Controller) doPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest00875")

	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = append(valuesList[1:])
		bar = valuesList[0]
	}

	input := []byte(bar)
	hasher := sha512.New()
	hasher.Write(input)

	result := hasher.Sum(nil)
	file, err := os.OpenFile("passwordFile.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Failed to open file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	encodedResult := base64.StdEncoding.EncodeToString(result)
	file.WriteString("hash_value=" + encodedResult + "\n")

	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Sensitive value '%s' hashed and stored<br/>", htmlEscape(string(input)))))
	c.Ctx.ResponseWriter.Write([]byte("Hash Test executed"))
}

func htmlEscape(s string) string {
	return template.HTMLEscapeString(s)
}
