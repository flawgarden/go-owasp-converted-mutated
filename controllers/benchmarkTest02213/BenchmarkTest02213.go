package controllers

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"html"
	"net/http"
	"os"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02213 struct {
	web.Controller
}

func (c *BenchmarkTest02213) Get() {
	c.Post()
}

func (c *BenchmarkTest02213) Post() {
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02213")
	bar := doSomething(param)

	md := sha1.New()
	input := []byte(bar)
	md.Write(input)

	result := md.Sum(nil)
	fileTarget := "passwordFile.txt"
	fw, err := os.OpenFile(fileTarget, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Unable to open file", http.StatusInternalServerError)
		return
	}
	defer fw.Close()
	if _, err := fw.WriteString("hash_value=" + base64.StdEncoding.EncodeToString(result) + "\n"); err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error writing file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(c.Ctx.ResponseWriter, "Sensitive value '%s' hashed and stored<br/>", html.EscapeString(string(input)))
	fmt.Fprintln(c.Ctx.ResponseWriter, "Hash Test executed")
}

func doSomething(param string) string {
	// Ваш код здесь для логики обработки
	return param
}
