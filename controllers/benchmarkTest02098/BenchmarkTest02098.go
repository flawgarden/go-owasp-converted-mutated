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

type BenchmarkTest02098Controller struct {
	web.Controller
}

func (c *BenchmarkTest02098Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02098Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	headers := c.Ctx.Request.Header["BenchmarkTest02098"]

	if len(headers) > 0 {
		param = headers[0]
	}

	param, _ = url.QueryUnescape(param)
	bar := doSomething(c.Ctx.Request, param)

	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME='foo' and PASSWORD='%s'", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
	defer db.Close()

	rows, err := db.Query(sqlStr)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
	defer rows.Close()

	user := models.User{}
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Password)
		if err != nil {
			c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
			return
		}
	}

	output, err := json.Marshal(user)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
	c.Ctx.ResponseWriter.Write(output)
}

func doSomething(req *http.Request, param string) string {
	bar := ""
	guess := "ABC"
	switchTarget := guess[1] // condition 'B', which is safe

	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bob"
	case 'C':
	case 'D':
		bar = param
	default:
		bar = "bob's your uncle"
	}

	return bar
}
