package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

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

type SqlInjectionVuln2Controller struct {
	beego.Controller
}

func (c *SqlInjectionVuln2Controller) Get() {
	queryString := c.Ctx.Input.URI()
	paramLoc := strings.Index(queryString, "BenchmarkTest01725=")
	if paramLoc == -1 {
		c.Ctx.ResponseWriter.Write([]byte("getQueryString() couldn't find expected parameter 'BenchmarkTest01725' in query string."))
		return
	}

	param := queryString[paramLoc+len("BenchmarkTest01725="):]
	ampersandLoc := strings.Index(param, "&")
	if ampersandLoc != -1 {
		param = param[:ampersandLoc]
	}

	bar := new(Test).doSomething(c.Ctx.Request, param)

	sqlStr := "SELECT TOP 1 USERNAME FROM USERS WHERE USERNAME='foo' AND PASSWORD='" + bar + "'"
	results, err := executeQuery(sqlStr)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}

	output, err := json.Marshal(results)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error marshalling results."))
		return
	}
	c.Ctx.ResponseWriter.Write(output)
}

func executeQuery(sqlStr string) (interface{}, error) {
	db, err := sql.Open("mysql", source)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var result string
	err = db.QueryRow(sqlStr).Scan(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type Test struct{}

func (t *Test) doSomething(request *http.Request, param string) string {
	// Simulating some processing on param
	return param
}
