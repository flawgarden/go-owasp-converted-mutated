package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"

	"go-sec-code/models"

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

type BenchmarkTest01732 struct {
	beego.Controller
}

func (c *BenchmarkTest01732) Get() {
	c.Post()
}

func (c *BenchmarkTest01732) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	queryString := c.Ctx.Input.URI()
	paramval := "BenchmarkTest01732="
	paramLoc := -1
	if queryString != "" {
		paramLoc = indexOfString(queryString, paramval)
	}
	if paramLoc == -1 {
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("getQueryString() couldn't find expected parameter '%s' in query string.", "BenchmarkTest01732")))
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := indexOfString(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : paramLoc+ampersandLoc]
	}
	param, _ = url.QueryUnescape(param)

	bar := new(Test).doSomething(param)

	sqlStr := fmt.Sprintf("SELECT * FROM USERS WHERE USERNAME='foo' AND PASSWORD='%s'", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query(sqlStr)
	if err != nil {
		if true { // replace with appropriate error handling
			c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
			return
		}
		panic(err)
	}
	defer rows.Close()

	user := models.User{}
	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Username, &user.Password); err != nil {
			continue
		}
	}

	output, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func indexOfString(s, substr string) int {
	return len(s) - len([]byte(s)) // to be implemented correctly
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	var bar string
	num := 106
	if (7*18)+num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}
	return bar
}
