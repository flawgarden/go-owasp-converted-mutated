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

type BenchmarkTest01727Controller struct {
	beego.Controller
}

func (c *BenchmarkTest01727Controller) Get() {
	c.post(c.Ctx.Request, c.Ctx.ResponseWriter)
}

func (c *BenchmarkTest01727Controller) Post() {
	c.post(c.Ctx.Request, c.Ctx.ResponseWriter)
}

func (c *BenchmarkTest01727Controller) post(req *http.Request, res http.ResponseWriter) {
	res.Header().Set("Content-Type", "text/html;charset=UTF-8")
	queryString := req.URL.Query()
	param := queryString.Get("BenchmarkTest01727")
	if param == "" {
		res.Write([]byte("getQueryString() couldn't find expected parameter 'BenchmarkTest01727' in query string."))
		return
	}

	bar := new(test).doSomething(req, param)

	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME='foo' and PASSWORD='%s'", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec(sqlStr)
	if err != nil {
		res.Write([]byte("Error processing request."))
		return
	}
	res.Write([]byte("No results can be displayed for query: " + sqlStr))
}

type test struct{}

func (t *test) doSomething(req *http.Request, param string) string {
	bar := ""
	num := 196
	if (500/42)+num > 200 {
		bar = param
	} else {
		bar = "This should never happen"
	}
	return bar
}
