package controllers

import (
	"encoding/json"

	"database/sql"
	"fmt"
	"go-sec-code/models"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00813Controller struct {
	web.Controller
}

func (c *BenchmarkTest00813Controller) Get() {
	queryString := c.Ctx.Request.URL.RawQuery
	if queryString == "" {
		c.Ctx.WriteString("getQueryString() couldn't find expected parameter 'BenchmarkTest00813' in query string.")
		return
	}
	paramval := "BenchmarkTest00813="
	paramLoc := -1
	paramLoc = len(queryString) - len(paramval) - 1

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := len(queryString)
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : ampersandLoc]
	}

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	id := param
	sqlStr := fmt.Sprintf("select * from user where id=%s", id)
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

func main() {
	web.Router("/xss-01/BenchmarkTest00813", &BenchmarkTest00813Controller{})
	web.Run()
}
