package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01812Controller struct {
	web.Controller
}

func (c *BenchmarkTest01812Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01812Controller) Post() {
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest01812")
	bar := testDoSomething(param)

	sqlStr := fmt.Sprintf("SELECT username FROM users WHERE username='foo' AND password='%s'", bar)
	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var results string
	err = db.QueryRow(sqlStr).Scan(&results)
	if err == sql.ErrNoRows {
		fmt.Fprintf(c.Ctx.ResponseWriter, "No results returned for query: %s", sqlStr)
	} else if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
	} else {
		fmt.Fprintf(c.Ctx.ResponseWriter, "Your results are: %s", results)
	}
}

func testDoSomething(param string) string {
	bar := "safe!"
	map20703 := make(map[string]interface{})
	map20703["keyA-20703"] = "a_Value"
	map20703["keyB-20703"] = param
	map20703["keyC"] = "another_Value"
	bar = map20703["keyB-20703"].(string)
	bar = map20703["keyA-20703"].(string)
	return bar
}
