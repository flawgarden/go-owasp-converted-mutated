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

type BenchmarkTest01506 struct {
	web.Controller
}

func (c *BenchmarkTest01506) Get() {
	c.Post()
}

func (c *BenchmarkTest01506) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	id := c.GetString("BenchmarkTest01506")
	if id == "" {
		id = ""
	}

	bar := new(testStruct).doSomething(c.Ctx.Request, id)
	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	_, _ = fmt.Fprintf(c.Ctx.ResponseWriter, bar, "a", "b")
}

type testStruct struct{}

func (t *testStruct) doSomething(r *http.Request, param string) string {
	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id=%s", param)
	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		panic(err)
	}

	output, _ := json.Marshal(user)
	return string(output)
}
