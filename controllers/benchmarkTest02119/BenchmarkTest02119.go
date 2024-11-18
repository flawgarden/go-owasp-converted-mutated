package controllers

import (
	"fmt"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

func init() {
	// Database initialization code
}

type BenchmarkTest02119Controller struct {
	web.Controller
}

func (c *BenchmarkTest02119Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02119Controller) Post() {
	response := c.Ctx.ResponseWriter
	response.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02119")
	if param == "" {
		param = ""
	}

	bar := doSomething(param)

	// Hashing and file writing logic
	// ...

	c.Ctx.Output.Body([]byte(fmt.Sprintf("Sensitive value '%s' hashed and stored<br/>", htmlentities(bar))))
}

func doSomething(param string) string {
	bar := "safe!"
	map70096 := make(map[string]interface{})
	map70096["keyA-70096"] = "a-Value"
	map70096["keyB-70096"] = param
	map70096["keyC"] = "another-Value"
	bar = map70096["keyB-70096"].(string)
	return bar
}

func htmlentities(str string) string {
	// Implementation of HTML entity encoding
	// ...
	return str
}
