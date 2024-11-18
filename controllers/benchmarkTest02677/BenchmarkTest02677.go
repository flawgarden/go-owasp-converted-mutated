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

type BenchmarkTest02677Controller struct {
	beego.Controller
}

func (c *BenchmarkTest02677Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02677Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02677")
	bar := doSomething(param)

	response, err := json.Marshal(bar)
	if err != nil {
		panic(err)
	}

	c.Ctx.ResponseWriter.Write(response)
}

func doSomething(param string) string {
	bar := param
	if param != "" && len(param) > 1 {
		sbxyz68995 := []rune(param)
		bar = string(append(sbxyz68995[:len(param)-1], 'Z'))
	}
	return bar
}

func (c *BenchmarkTest02677Controller) handleSQLInjection() {
	id := c.GetString("id")
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
