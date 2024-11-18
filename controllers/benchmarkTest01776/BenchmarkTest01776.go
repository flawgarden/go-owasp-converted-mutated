package controllers

import (
	"database/sql"
	"net/http"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01776 struct {
	web.Controller
}

func (c *BenchmarkTest01776) Get() {
	c.Post()
}

func (c *BenchmarkTest01776) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	id := c.GetString("BenchmarkTest01776")

	bar := testDoSomething(c.Ctx.Request, id)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	if bar != "" {
		c.Ctx.ResponseWriter.Write([]byte(bar))
	}
}

func testDoSomething(r *http.Request, param string) string {
	bar := "safe!"
	map3531 := make(map[string]interface{})
	map3531["keyA-3531"] = "a_Value"
	map3531["keyB-3531"] = param
	map3531["keyC"] = "another_Value"
	bar = map3531["keyB-3531"].(string)
	bar = map3531["keyA-3531"].(string)

	return bar
}

func init() {
	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
