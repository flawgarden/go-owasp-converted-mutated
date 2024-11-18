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

type BenchmarkTest02271Controller struct {
	beego.Controller
}

func (c *BenchmarkTest02271Controller) Get() {
	c.doPost()
}

func (c *BenchmarkTest02271Controller) Post() {
	c.doPost()
}

func (c *BenchmarkTest02271Controller) doPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02271")
	bar := doSomething(param)

	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME=? and PASSWORD='%s'", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		c.Ctx.WriteString("Error processing request.")
		return
	}
	defer db.Close()

	statement, err := db.Prepare(sqlStr)
	if err != nil {
		c.Ctx.WriteString("Error processing request.")
		return
	}
	defer statement.Close()

	rows, err := statement.Query("foo")
	if err != nil {
		c.Ctx.WriteString("Error processing request.")
		return
	}
	defer rows.Close()

	var user models.User
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Username, &user.Password)
		if err != nil {
			c.Ctx.WriteString("Error processing request.")
			return
		}
	}

	output, err := json.Marshal(user)
	if err != nil {
		c.Ctx.WriteString("Error processing request.")
		return
	}
	c.Ctx.ResponseWriter.Write(output)
}

func doSomething(param string) string {
	var bar string
	num := 106
	if (7*18)+num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}
	return bar
}
