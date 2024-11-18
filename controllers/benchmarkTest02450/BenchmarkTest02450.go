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

type BenchmarkTest02450Controller struct {
	beego.Controller
}

func (c *BenchmarkTest02450Controller) Get() {
	c.doPost()
}

func (c *BenchmarkTest02450Controller) Post() {
	c.doPost()
}

func (c *BenchmarkTest02450Controller) doPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02450")
	if param == "" {
		param = ""
	}

	bar := doSomething(param)

	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME=? and PASSWORD='%s'", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
	defer db.Close()

	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query("foo")
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
	defer rows.Close()

	var results []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Username, &user.Password); err != nil {
			c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
			return
		}
		results = append(results, user)
	}

	output, err := json.Marshal(results)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
	c.Ctx.ResponseWriter.Write(output)
}

func doSomething(param string) string {
	bar := ""
	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}
	return bar
}
