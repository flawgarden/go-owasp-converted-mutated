package controllers

import (
	"database/sql"
	"fmt"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01465Controller struct {
	beego.Controller
}

func (c *BenchmarkTest01465Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01465Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	flag := true
	names := c.Ctx.Request.URL.Query()

	for name, values := range names {
		if flag {
			for _, value := range values {
				if value == "BenchmarkTest01465" {
					param = name
					flag = false
					break
				}
			}
		}
	}

	bar := doSomething(param)

	sqlQuery := fmt.Sprintf("SELECT * from USERS where USERNAME='foo' and PASSWORD='%s'", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		c.Data["json"] = "Error processing request."
		c.ServeJSON()
		return
	}
	defer db.Close()

	_, err = db.Exec(sqlQuery)
	if err != nil {
		c.Data["json"] = "Error processing request."
		c.ServeJSON()
		return
	}

	c.Ctx.ResponseWriter.Write([]byte("No results can be displayed for query: " + htmlEscape(sqlQuery)))
}

func doSomething(param string) string {
	bar := ""
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the param value
	}

	return bar
}

func htmlEscape(input string) string {
	return strings.ReplaceAll(strings.ReplaceAll(input, "<", "&lt;"), ">", "&gt;")
}
