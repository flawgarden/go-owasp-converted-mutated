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

type BenchmarkTest02092 struct {
	web.Controller
}

func (c *BenchmarkTest02092) Get() {
	c.DoPost()
}

func (c *BenchmarkTest02092) Post() {
	c.DoPost()
}

func (c *BenchmarkTest02092) DoPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	req := c.Ctx.Request
	param := ""
	if headers := req.Header["BenchmarkTest02092"]; len(headers) > 0 {
		param = headers[0] // just grab first element
	}

	param, _ = url.QueryUnescape(param)

	bar := doSomething(param)

	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME='foo' and PASSWORD='%s'", bar)
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

	var user models.User
	err = statement.QueryRow().Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}

	output, err := json.Marshal(user)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
	c.Ctx.ResponseWriter.Write(output)
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
