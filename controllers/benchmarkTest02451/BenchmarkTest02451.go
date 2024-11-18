package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02451Controller struct {
	web.Controller
}

func init() {
	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()
}

func (c *BenchmarkTest02451Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02451Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest02451")
	if param == "" {
		param = ""
	}

	bar := doSomething(param)

	sqlStr := fmt.Sprintf("SELECT * FROM USERS WHERE USERNAME='foo' AND PASSWORD='%s'", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	if _, err := db.Exec(sqlStr); err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
	} else {
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("No results can be displayed for query: %s<br> because the database execute method doesn't return results.", sqlStr)))
	}
}

func doSomething(param string) string {
	return param // Здесь можно добавить логику для обработки параметра
}
