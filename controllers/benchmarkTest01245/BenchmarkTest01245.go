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

type BenchmarkTest01245 struct {
	web.Controller
}

func (c *BenchmarkTest01245) Get() {
	c.DoPost()
}

func (c *BenchmarkTest01245) Post() {
	c.DoPost()
}

func (c *BenchmarkTest01245) DoPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest01245")
	if param == "" {
		param = ""
	}

	bar := new(Test).doSomething(c.Ctx.Request, param)

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

type Test struct{}

func (t *Test) doSomething(request *http.Request, param string) string {
	bar := param
	num := 106

	if (7*42)-num > 200 {
		bar = "This should never happen"
	}

	return bar
}
