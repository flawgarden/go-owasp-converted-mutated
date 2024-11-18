package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-sec-code/models"

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

type BenchmarkTest02269Controller struct {
	beego.Controller
}

func (c *BenchmarkTest02269Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02269Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02269")
	bar := doSomething(param)

	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME=? and PASSWORD='%s'", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	statement, err := db.Prepare(sqlStr)
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	rows, err := statement.Query("foo")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var user models.User
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Password)
		if err != nil {
			panic(err)
		}
	}

	output, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func doSomething(param string) string {
	bar := ""
	num := 196
	if (500/42)+num > 200 {
		bar = param
	} else {
		bar = "This should never happen"
	}
	return bar
}
