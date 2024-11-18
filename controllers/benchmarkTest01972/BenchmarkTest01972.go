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

type BenchmarkTest01972 struct {
	web.Controller
}

func (c *BenchmarkTest01972) Get() {
	c.Post()
}

func (c *BenchmarkTest01972) Post() {
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")

	param := ""
	if c.Ctx.Input.Header("BenchmarkTest01972") != "" {
		param = c.Ctx.Input.Header("BenchmarkTest01972")
	}

	param, _ = url.QueryUnescape(param)

	bar := doSomething(param)

	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME='foo' and PASSWORD='%s'", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		c.Ctx.Output.Body([]byte("Error processing request."))
		return
	}
	defer db.Close()

	statement, err := db.Prepare(sqlStr)
	if err != nil {
		if err != nil {
			c.Ctx.Output.Body([]byte("Error processing request."))
			return
		}
	}

	rows, err := statement.Query()
	if err != nil {
		c.Ctx.Output.Body([]byte("Error processing request."))
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Username, &user.Password); err != nil {
			c.Ctx.Output.Body([]byte("Error processing request."))
			return
		}
		users = append(users, user)
	}

	output, err := json.Marshal(users)
	if err != nil {
		c.Ctx.Output.Body([]byte("Error processing request."))
		return
	}
	c.Ctx.Output.Body(output)
}

func doSomething(param string) string {
	bar := param
	return bar
}
