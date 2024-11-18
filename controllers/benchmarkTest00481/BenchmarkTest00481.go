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

type BenchmarkTest00481Controller struct {
	web.Controller
}

func (c *BenchmarkTest00481Controller) Get() {
	c.handleRequest()
}

func (c *BenchmarkTest00481Controller) Post() {
	c.handleRequest()
}

func (c *BenchmarkTest00481Controller) handleRequest() {
	id := c.GetString("BenchmarkTest00481")

	sqlStr := fmt.Sprintf("select * from user where id=%s", id)
	user := models.User{}

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		panic(err)
	}

	output, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	c.Ctx.Output.Body(output)
}

func main() {
	web.Router("/cmdi-00/BenchmarkTest00481", &BenchmarkTest00481Controller{})
	web.Run()
}
