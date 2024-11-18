package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"

	"go-sec-code/models"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01694Controller struct {
	web.Controller
}

func (c *BenchmarkTest01694Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01694Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	queryString := c.Ctx.Request.URL.RawQuery
	paramval := "BenchmarkTest01694="
	paramLoc := -1
	if queryString != "" {
		paramLoc = len(queryString) - len(paramval)
	}
	if paramLoc == -1 {
		c.Ctx.ResponseWriter.Write([]byte("getQueryString() couldn't find expected parameter 'BenchmarkTest01694' in query string."))
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := len(queryString)

	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : ampersandLoc]
	}
	param, _ = url.QueryUnescape(param)

	user := models.User{}

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("select * from user where id=%s", param)
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		panic(err)
	}

	output, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}
