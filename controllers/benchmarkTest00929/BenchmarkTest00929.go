package controllers

import (
	"database/sql"
	"encoding/json"
	"go-sec-code/models"
	"net/http"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00929Controller struct {
	web.Controller
}

func (c *BenchmarkTest00929Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00929Controller) Post() {
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest00929")

	var bar string
	guess := "ABC"
	switchTarget := guess[2]

	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bobs_your_uncle"
	case 'C':
		bar = param
	case 'D':
		bar = param
	default:
		bar = "bobs_your_uncle"
	}

	sqlStr := "SELECT * from USERS where USERNAME=? and PASSWORD='" + bar + "'"

	db, err := sql.Open("mysql", source)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		return
	}
	defer db.Close()

	statement, err := db.Prepare(sqlStr)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		return
	}
	defer statement.Close()

	statement.Exec("foo")

	rows, err := statement.Query()
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Username, &user.Password); err != nil {
			c.Ctx.Output.SetStatus(http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}
	output, err := json.Marshal(users)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		return
	}
	c.Ctx.ResponseWriter.Write(output)
}
