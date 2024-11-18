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

type BenchmarkTest00331Controller struct {
	web.Controller
}

func (c *BenchmarkTest00331Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00331Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	header := c.Ctx.Request.Header.Get("BenchmarkTest00331")
	if header != "" {
		param = header
	}

	param, _ = url.QueryUnescape(param)

	bar := "safe!"
	map59781 := make(map[string]interface{})
	map59781["keyA-59781"] = "a_Value"
	map59781["keyB-59781"] = param
	map59781["keyC"] = "another_Value"
	bar = map59781["keyB-59781"].(string)
	bar = map59781["keyA-59781"].(string)

	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME=? and PASSWORD='%s'", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec("foo")
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
		return
	}

	// Здесь можно добавить логику для возврата результата запроса
	// org.owasp.benchmark.helpers.DatabaseHelper.printResults(stmt, sqlStr, c.Ctx.ResponseWriter)
}
