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

type BenchmarkTest02717 struct {
	web.Controller
}

func (c *BenchmarkTest02717) Get() {
	c.Post()
}

func (c *BenchmarkTest02717) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	id := c.GetString("BenchmarkTest02717")

	bar := doSomething(id)

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("select * from user where id=%s", bar)
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
	bar := param
	if param != "" && len(param) > 1 {
		bar = param[:len(param)-1]
	}
	return bar
}

func main() {
	web.Router("/weakrand-06/BenchmarkTest02717", &BenchmarkTest02717{})
	web.AutoRouter(&BenchmarkTest02717{})
	web.Run()
}