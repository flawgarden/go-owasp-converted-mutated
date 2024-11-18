package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"go-sec-code/models"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01715Controller struct {
	web.Controller
}

func (c *BenchmarkTest01715Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01715Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	queryString := c.Ctx.Request.URL.RawQuery
	paramval := "BenchmarkTest01715="
	paramLoc := -1
	if queryString != "" {
		paramLoc = findParamLocation(queryString, paramval)
	}
	if paramLoc == -1 {
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("getQueryString() couldn't find expected parameter '%s' in query string.", "BenchmarkTest01715")))
		return
	}
	param := extractParam(queryString, paramLoc, paramval)
	bar := new(Test).doSomething(c.Ctx.Request, param)

	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME=? and PASSWORD='%s'", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec("foo")
	if err != nil {
		panic(err)
	}

	rows, err := stmt.Query()
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	user := models.User{}
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Password)
		if err != nil {
			panic(err)
		}
	}

	output, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func findParamLocation(queryString, paramval string) int {
	return -1 // Implement logic to find parameter location
}

func extractParam(queryString string, paramLoc int, paramval string) string {
	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := indexOfAmpersand(queryString, paramLoc)
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : ampersandLoc]
	}
	decodedParam, _ := url.QueryUnescape(param)
	return decodedParam
}

func indexOfAmpersand(queryString string, paramLoc int) int {
	return -1 // Implement logic to find index of '&'
}

type Test struct{}

func (t *Test) doSomething(req *http.Request, param string) string {
	// Implement your logic here
	return param
}
