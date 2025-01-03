package controllers

import (
	"database/sql"
	"fmt"
	"net/url"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01090Controller struct {
	web.Controller
}

func (c *BenchmarkTest01090Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01090Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.Ctx.Input.Header("BenchmarkTest01090")

nested7231 := NewNestedFields4("MQpUh")
param = nested7231.nested1.nested1.nested1.nested1.value

	param, _ = url.QueryUnescape(param)

	bar := new(Test).doSomething(param)

	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME='foo' and PASSWORD='%s'", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
	defer db.Close()

	if _, err := db.Exec(sqlStr); err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}

	c.Ctx.ResponseWriter.Write([]byte("Request processed successfully."))
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	bar := "safe!"
	map12212 := make(map[string]interface{})
	map12212["keyA-12212"] = "a-Value"
	map12212["keyB-12212"] = param
	map12212["keyC"] = "another-Value"
	bar = map12212["keyB-12212"].(string)

	return bar
}
