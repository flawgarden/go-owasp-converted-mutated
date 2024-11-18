package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02538Controller struct {
	web.Controller
}

func init() {
	web.Router("/sqli-05/BenchmarkTest02538", &BenchmarkTest02538Controller{})
}

func (c *BenchmarkTest02538Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02538Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html; charset=UTF-8")

	param := c.GetString("BenchmarkTest02538")
	bar := doSomething(c.Ctx.Request, param)

	sqlStr := fmt.Sprintf("SELECT TOP 1 USERNAME from USERS where USERNAME='foo' and PASSWORD='%s'", bar)
	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var result string
	err = db.QueryRow(sqlStr).Scan(&result)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "No results returned for query: "+sqlStr, http.StatusNotFound)
		return
	}
	c.Ctx.ResponseWriter.Write([]byte("Your results are: " + result))
}

func doSomething(req *http.Request, param string) string {
	a := param
	b := a + " SafeStuff"
	b = b[:len(b)-5] + "Chars"
	return b // Simplified for demonstration
}
