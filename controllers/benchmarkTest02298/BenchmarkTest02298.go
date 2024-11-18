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

type BenchmarkTest02298Controller struct {
	beego.Controller
}

func (c *BenchmarkTest02298Controller) Get() {
	c.DoPost()
}

func (c *BenchmarkTest02298Controller) Post() {
	c.DoPost()
}

func (c *BenchmarkTest02298Controller) DoPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	var param string
	names := c.Ctx.Request.URL.Query()
	for name, values := range names {
		if values != nil {
			for _, value := range values {
				if value == "BenchmarkTest02298" {
					param = name
					break
				}
			}
		}
	}

	bar := doSomething(param)

	id := bar
	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id=%s", id)
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
	// Simulate a function that modifies the parameter
	return param
}
