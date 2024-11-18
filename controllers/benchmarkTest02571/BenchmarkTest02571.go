package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-sec-code/models"
	"net/url"

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

type BenchmarkTest02571Controller struct {
	beego.Controller
}

func (c *BenchmarkTest02571Controller) Get() {
	c.doPost()
}

func (c *BenchmarkTest02571Controller) Post() {
	c.doPost()
}

func (c *BenchmarkTest02571Controller) doPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	queryString := c.Ctx.Request.URL.RawQuery
	paramVal := "BenchmarkTest02571="
	paramLoc := -1
	if queryString != "" {
		paramLoc = len(queryString) - len(paramVal)
		if paramLoc < 0 || queryString[paramLoc:paramLoc+len(paramVal)] != paramVal {
			c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("getQueryString() couldn't find expected parameter '%s' in query string.", "BenchmarkTest02571")))
			return
		}
	}

	param := queryString[paramLoc+len(paramVal):]
	ampersandLoc := len(queryString)
	for i, ch := range queryString[paramLoc:] {
		if ch == '&' {
			ampersandLoc = i + paramLoc
			break
		}
	}
	param = queryString[paramLoc+len(paramVal) : ampersandLoc]
	// url decode
	param = url.QueryEscape(param)

	bar := doSomething(param)

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id=%s", bar)
	user := models.User{}
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

func doSomething(param string) string {
	bar := ""
	guess := "ABC"
	switchTarget := guess[1]

	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bob"
	case 'C':
		bar = param
	case 'D':
		bar = param
	default:
		bar = "bob's your uncle"
	}

	return bar
}
