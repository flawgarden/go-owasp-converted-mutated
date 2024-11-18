package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/url"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", source)
	orm.Debug = true
}

type BenchmarkTest01084Controller struct {
	beego.Controller
}

func (c *BenchmarkTest01084Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01084Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.Ctx.Request.Header.Get("BenchmarkTest01084")
	param, _ = url.QueryUnescape(param)

	bar := new(Test).doSomething(c.Ctx.Request, param)

	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME=? and PASSWORD='%s'", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
	defer db.Close()

	statement, err := db.Prepare(sqlStr)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}

	_, err = statement.Exec("foo")
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}

	c.Ctx.ResponseWriter.Write([]byte("Query executed successfully."))
}

type Test struct{}

func (t *Test) doSomething(request *http.Request, param string) string {
	bar := ""
	num := 196
	if (500/42)+num > 200 {
		bar = param
	} else {
		bar = "This should never happen"
	}
	return bar
}
