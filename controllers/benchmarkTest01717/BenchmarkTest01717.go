package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

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

type BenchmarkTest01717Controller struct {
	beego.Controller
}

func (c *BenchmarkTest01717Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01717Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	queryString := c.Ctx.Request.URL.Query()
	paramval := "BenchmarkTest01717"

	if values, ok := queryString[paramval]; ok {
		param := values[0]
		bar := doSomething(c.Ctx.Request, param)

		sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME=? and PASSWORD='%s'", bar)

		db, err := sql.Open("mysql", source)
		if err != nil {
			c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
			return
		}
		defer db.Close()

		statement, err := db.Prepare(sqlStr)
		if err != nil {
			c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
			return
		}
		defer statement.Close()

		result, err := statement.Exec("foo")
		if err != nil {
			c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
			return
		}
		output, _ := json.Marshal(result)
		c.Ctx.ResponseWriter.Write(output)
	} else {
		c.Ctx.ResponseWriter.Write([]byte("getQueryString() couldn't find expected parameter '" + paramval + "' in query string."))
	}
}

func doSomething(request *http.Request, param string) string {
	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[1]         // get the last 'safe' value
	}
	return bar
}
