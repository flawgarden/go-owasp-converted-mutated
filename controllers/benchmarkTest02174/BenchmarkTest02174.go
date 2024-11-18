package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

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

type BenchmarkTest02174Controller struct {
	beego.Controller
}

func (c *BenchmarkTest02174Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02174Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02174")
	if param == "" {
		param = ""
	}

	bar := doSomething(c.Ctx.Request, param)

	sqlStr := fmt.Sprintf("SELECT * FROM USERS WHERE USERNAME='foo' AND PASSWORD='%s'", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec(sqlStr)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
	c.Ctx.ResponseWriter.Write([]byte("No results can be displayed for query: " + sqlStr + "<br>"))
}

func doSomething(req *http.Request, param string) string {
	bar := ""
	num := 106

	if (7*18)+num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}

	return bar
}
