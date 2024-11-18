package controllers

import (
	"database/sql"
	"encoding/json"
	"strings"

	"go-sec-code/models"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02215Controller struct {
	web.Controller
}

func (c *BenchmarkTest02215Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02215Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02215")
	bar := doSomething(param)

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStr := "SELECT * FROM user WHERE username='" + strings.Replace(bar, "'", "''", -1) + "'"
	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		panic(err)
	}

	output, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func doSomething(param string) string {
	num := 106
	if (7*42)-num > 200 {
		return "This should never happen"
	}
	return param
}
