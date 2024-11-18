package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/url"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01089Controller struct {
	web.Controller
}

func (c *BenchmarkTest01089Controller) Get() {
	c.DoPost()
}

func (c *BenchmarkTest01089Controller) Post() {
	c.DoPost()
}

func (c *BenchmarkTest01089Controller) DoPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.Ctx.Request.Header.Get("BenchmarkTest01089")
	param, _ = url.QueryUnescape(param)

	bar := new(Test).doSomething(param)

	sqlStr := fmt.Sprintf("SELECT TOP 1 USERNAME from USERS where USERNAME='foo' and PASSWORD='%s'", bar)
	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var results string
	err = db.QueryRow(sqlStr).Scan(&results)
	if err == sql.ErrNoRows {
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("No results returned for query: %s", sqlStr)))
	} else if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
	} else {
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Your results are: %s", htmlEscape(results))))
	}
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	bar := "safe!"
	mapData := make(map[string]string)
	mapData["keyA-11607"] = "a_Value"
	mapData["keyB-11607"] = param
	mapData["keyC"] = "another_Value"

	bar = mapData["keyB-11607"]
	bar = mapData["keyA-11607"]

	return bar
}

func htmlEscape(str string) string {
	return fmt.Sprintf("%s", str) // заменить на безопасный метод экранирования, если нужно
}
