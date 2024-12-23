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

type BenchmarkTest00681Controller struct {
	beego.Controller
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", source)
	orm.Debug = true
}

func (c *BenchmarkTest00681Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00681Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest00681")
	if param == "" {
		param = ""
	}

	bar := param
	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME='foo' and PASSWORD='%s'", bar)

	db, err := sql.Open("mysql", source)

list787231 := make([] string, 0)
list787231 = append(list787231, sqlStr)
list787231 = nil
sqlStr = list787231[0]

	if err != nil {
		c.Data["json"] = "Error processing request."
		c.ServeJSON()
		return
	}
	defer db.Close()

	rows, err := db.Query(sqlStr)
	if err != nil {
		c.Data["json"] = "Error processing request."
		c.ServeJSON()
		return
	}
	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Username, &user.Password); err != nil {
			c.Data["json"] = "Error processing request."
			c.ServeJSON()
			return
		}
		users = append(users, user)
	}

	output, err := json.Marshal(users)
	if err != nil {
		c.Data["json"] = "Error processing request."
		c.ServeJSON()
		return
	}

	c.Ctx.ResponseWriter.Write(output)
}
