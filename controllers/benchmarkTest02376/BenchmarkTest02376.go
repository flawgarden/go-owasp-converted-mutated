package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02376Controller struct {
	beego.Controller
}

func (c *BenchmarkTest02376Controller) Get() {
	c.DoPost()
}

func (c *BenchmarkTest02376Controller) DoPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	bar := c.doSomething(c.Ctx.Request)

	// LDAP code simulation
	response := fmt.Sprintf("LDAP query results:<br>Record found with name %s<br>", bar)
	c.Ctx.ResponseWriter.Write([]byte(response))
}

func (c *BenchmarkTest02376Controller) doSomething(request *http.Request) string {
	param := request.URL.Query().Get("BenchmarkTest02376")
	if param == "" {
		param = ""
	}

	bar := param
	num := 106
	if (7*18)+num > 200 {
		bar = "This_should_always_happen"
	}

	return bar
}

func init() {
	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		panic(err)
	}
}
