package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-sec-code/models"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00183Controller struct {
	web.Controller
}

func (c *BenchmarkTest00183Controller) Get() {
	c.TplName = "index.tpl"
}

func (c *BenchmarkTest00183Controller) Post() {
	response := c.Ctx.ResponseWriter
	request := c.Ctx.Request

	response.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := request.Header.Get("BenchmarkTest00183")
	bar := ""

	guess := "ABC"
	switchTarget := guess[1] // condition 'B', which is safe

	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bob"
	case 'C', 'D':
		bar = param
	default:
		bar = "bob's your uncle"
	}

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var user models.User
	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE username='%s'", bar)
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		panic(err)
	}

	output, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	response.Write(output)
}
