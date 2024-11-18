package controllers

import (
	"fmt"
	"html"
	"net/http"
	"os"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01247Controller struct {
	web.Controller
}

func (c *BenchmarkTest01247Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01247Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest01247")
	if param == "" {
		param = ""
	}

	bar := newTest().doSomething(c.Ctx.Request, param)

	// Симуляция операции с MD5 (просто для примера, не безопасно)
	hashOutput := []byte(fmt.Sprintf("hashed_value_of_%s", bar))
	filePath := "passwordFile.txt"
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err := f.WriteString(fmt.Sprintf("hash_value=%s\n", hashOutput)); err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Sensitive value '%s' hashed and stored<br/>", html.EscapeString(bar))))
	c.Ctx.ResponseWriter.Write([]byte("Hash Test executed"))
}

type test struct{}

func newTest() *test {
	return &test{}
}

func (t *test) doSomething(req *http.Request, param string) string {
	bar := param
	if len(param) > 1 {
		bar = param[:len(param)-1]
	}
	return bar
}
