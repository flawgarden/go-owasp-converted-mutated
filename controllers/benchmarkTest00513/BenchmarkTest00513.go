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

type BenchmarkTest00513Controller struct {
	beego.Controller
}

func (c *BenchmarkTest00513Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00513Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest00513")
	bar := "safe!"
	map63945 := make(map[string]interface{})
	map63945["keyA-63945"] = "a_Value"
	map63945["keyB-63945"] = param
	map63945["keyC"] = "another_Value"
	bar = map63945["keyB-63945"].(string)
	bar = map63945["keyA-63945"].(string)

	sqlStr := fmt.Sprintf("SELECT * FROM USERS WHERE USERNAME='foo' AND PASSWORD='%s'", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	statement, err := db.Prepare(sqlStr)
	if err != nil {
		c.Ctx.WriteString("Error processing request.")
		return
	}
	defer statement.Close()

	rows, err := statement.Query()
	if err != nil {
		c.Ctx.WriteString("Error processing request.")
		return
	}
	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Username, &user.Password); err != nil {
			c.Ctx.WriteString("Error processing request.")
			return
		}
		users = append(users, user)
	}

	output, err := json.Marshal(users)
	if err != nil {
		c.Ctx.WriteString("Error processing request.")
		return
	}
	c.Ctx.ResponseWriter.Write(output)
}
