package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-sec-code/models"
	"net/http"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01364 struct {
	web.Controller
}

func (c *BenchmarkTest01364) Get() {
	c.Post()
}

func (c *BenchmarkTest01364) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	id := c.GetString("BenchmarkTest01364")

	bar := c.doSomething(id)

	sqlStr := fmt.Sprintf("select * from user where id=%s", bar)
	user := models.User{}

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
	output, err := json.Marshal(user)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
	c.Ctx.ResponseWriter.Write(output)
}

func (c *BenchmarkTest01364) doSomething(param string) string {
	bar := ""

	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}
	return bar
}

func main() {
	web.Router("/cmdi-01/BenchmarkTest01364", &BenchmarkTest01364{})
	web.Run()
}
