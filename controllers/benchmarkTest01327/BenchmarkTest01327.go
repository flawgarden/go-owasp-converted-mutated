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

type BenchmarkTest01327 struct {
	web.Controller
}

func (c *BenchmarkTest01327) Get() {
	c.Post()
}

func (c *BenchmarkTest01327) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest01327")

	bar := new(Test).doSomething(c.Ctx.Request, param)

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("select * from user where username='%s'", bar)
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

func (t *Test) doSomething(req *http.Request, param string) string {
	bar := "safe!"
	map41804 := make(map[string]interface{})
	map41804["keyA-41804"] = "a_Value"
	map41804["keyB-41804"] = param
	map41804["keyC"] = "another_Value"
	bar = map41804["keyB-41804"].(string)
	bar = map41804["keyA-41804"].(string)

	return bar
}
