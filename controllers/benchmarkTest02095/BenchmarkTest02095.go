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

type BenchmarkTest02095Controller struct {
	web.Controller
}

func (c *BenchmarkTest02095Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02095Controller) Post() {
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")

	param := ""

	headers := c.Ctx.Request.Header["BenchmarkTest02095"]
	if len(headers) > 0 {
		param = headers[0]
	}

	param, _ = url.PathUnescape(param)

	bar := doSomething(param)

	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME='foo' and PASSWORD='%s'", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Ctx.Output.Body([]byte("Error processing request."))
		return
	}
	defer db.Close()

	statement, err := db.Prepare(sqlStr)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Ctx.Output.Body([]byte("Error processing request."))
		return
	}
	defer statement.Close()

	rows, err := statement.Query()
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Ctx.Output.Body([]byte("Error processing request."))
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Username, &user.Password); err != nil {
			c.Ctx.Output.SetStatus(http.StatusInternalServerError)
			c.Ctx.Output.Body([]byte("Error processing request."))
			return
		}
		users = append(users, user)
	}

	output, err := json.Marshal(users)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Ctx.Output.Body([]byte("Error processing request."))
		return
	}
	c.Ctx.Output.Body(output)
}

func doSomething(param string) string {
	bar := "safe!"
	m := make(map[string]interface{})
	m["keyA-43923"] = "a_Value"
	m["keyB-43923"] = param
	m["keyC"] = "another_Value"
	bar = m["keyB-43923"].(string)
	bar = m["keyA-43923"].(string)
	return bar
}
