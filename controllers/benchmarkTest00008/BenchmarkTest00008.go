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

type BenchmarkTest00008Controller struct {
	web.Controller
}

func (c *BenchmarkTest00008Controller) Get() {
	c.handleRequest()
}

func (c *BenchmarkTest00008Controller) Post() {
	c.handleRequest()
}

func (c *BenchmarkTest00008Controller) handleRequest() {
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")
	param := c.Ctx.Request.Header.Get("BenchmarkTest00008")
	param, _ = url.QueryUnescape(param)

	sqlStr := fmt.Sprintf("{call %s}", param)

	db, err := sql.Open("mysql", source)
	if err != nil {
		c.Ctx.Output.Body([]byte("Error processing request."))
		return
	}
	defer db.Close()

	rows, err := db.Query(sqlStr)
	if err != nil {
		c.Ctx.Output.Body([]byte("Error processing request."))
		return
	}
	defer rows.Close()

	users := []models.User{}
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
