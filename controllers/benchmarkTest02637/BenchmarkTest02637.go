package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"go-sec-code/models"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02637Controller struct {
	web.Controller
}

func (c *BenchmarkTest02637Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02637Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	queryString := c.Ctx.Request.URL.RawQuery
	paramval := "BenchmarkTest02637="
	paramLoc := strings.Index(queryString, paramval)
	if paramLoc == -1 {
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("getQueryString() couldn't find expected parameter '%s' in query string.", "BenchmarkTest02637")))
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := strings.Index(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : paramLoc+ampersandLoc]
	}
	param, _ = url.QueryUnescape(param)

	bar := doSomething(param)

	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME=? and PASSWORD='%s'", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	statement, err := db.Prepare(sqlStr)
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	rows, err := statement.Query("foo")
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
	defer rows.Close()

	var results []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Username, &user.Password); err != nil {
			c.Ctx.ResponseWriter.Write([]byte("Error processing results."))
			return
		}
		results = append(results, user)
	}
	output, _ := json.Marshal(results)
	c.Ctx.ResponseWriter.Write(output)
}

func doSomething(param string) string {
	var bar string
	num := 106
	if (7*18)+num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}
	return bar
}
