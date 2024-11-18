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

type BenchmarkTest00762Controller struct {
	beego.Controller
}

func (c *BenchmarkTest00762Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00762Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	values := c.GetStrings("BenchmarkTest00762")
	var param string
	if len(values) > 0 {
		param = values[0]
	}

	var bar string
	num := 196
	if (500/42)+num > 200 {
		bar = param
	} else {
		bar = "This should never happen"
	}

	sqlStr := fmt.Sprintf("{call %s}", bar)

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

	rows, err := stmt.Query()
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
	defer rows.Close()

	var user models.User
	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Username, &user.Password); err != nil {
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
