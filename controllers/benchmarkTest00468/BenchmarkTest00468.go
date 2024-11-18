package controllers

import (
	"fmt"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00468Controller struct {
	web.Controller
}

func (c *BenchmarkTest00468Controller) Get() {
	c.DoPost()
}

func (c *BenchmarkTest00468Controller) Post() {
	c.DoPost()
}

func (c *BenchmarkTest00468Controller) DoPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest00468")
	bar := "safe!"
	map50379 := make(map[string]interface{})
	map50379["keyA-50379"] = "a_Value"
	map50379["keyB-50379"] = param
	map50379["keyC"] = "another_Value"
	bar = map50379["keyB-50379"].(string)
	bar = map50379["keyA-50379"].(string)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	_, err := fmt.Fprintf(c.Ctx.ResponseWriter, bar, "a", "b")
	if err != nil {
		panic(err)
	}
}
