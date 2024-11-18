package controllers

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"html"
	"os"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01333Controller struct {
	web.Controller
}

func (c *BenchmarkTest01333Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01333Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest01333")

	bar := doSomething(param)

	md := md5.New()
	input := []byte(bar)
	md.Write(input)
	result := md.Sum(nil)

	fileTarget, _ := os.OpenFile("passwordFile.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	defer fileTarget.Close()
	fileTarget.WriteString(fmt.Sprintf("hash_value=%s\n", base64.StdEncoding.EncodeToString(result)))

	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Sensitive value '%s' hashed and stored<br/>", html.EscapeString(string(input)))))
	c.Ctx.ResponseWriter.Write([]byte("Hash Test executed"))
}

func doSomething(param string) string {
	return html.EscapeString(param)
}
