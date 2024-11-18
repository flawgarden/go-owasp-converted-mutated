package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type CryptoBenchmarkController struct {
	web.Controller
}

func (c *CryptoBenchmarkController) Get() {
	c.Post()
}

func (c *CryptoBenchmarkController) Post() {
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02663")
	bar := doSomething(param)

	// Simulate encryption step
	encryptedData := encryptData([]byte(bar))

	fileTarget := "passwordFile.txt"
	fw, err := os.OpenFile(fileTarget, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Could not open file", http.StatusInternalServerError)
		return
	}
	defer fw.Close()

	if _, err := fw.WriteString(fmt.Sprintf("secret_value=%s\n", encryptedData)); err != nil {
		http.Error(c.Ctx.ResponseWriter, "Could not write to file", http.StatusInternalServerError)
		return
	}

	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Sensitive value: '%s' encrypted and stored<br/>", htmlEncode(bar))))
}

func doSomething(param string) string {
	return htmlEscape(param)
}

func htmlEscape(input string) string {
	return strings.ReplaceAll(strings.ReplaceAll(input, "&", "&amp;"), "<", "&lt;")
}

func htmlEncode(input string) string {
	return htmlEscape(input)
}

func encryptData(input []byte) string {
	// Simulating encryption (replace with actual encryption logic)
	return fmt.Sprintf("Encrypted(%s)", string(input))
}
