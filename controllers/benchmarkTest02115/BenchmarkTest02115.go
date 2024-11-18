package controllers

import (
	"database/sql"
	"encoding/json"
	"go-sec-code/models"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

func init() {
	sql.Register("mysql", nil)
}

type BenchmarkTest02115Controller struct {
	web.Controller
}

func (c *BenchmarkTest02115Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02115Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02115")
	if param == "" {
		param = ""
	}

	bar := doSomething(param)

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStr := "SELECT * FROM user WHERE uid=?"
	user := models.User{}

	err = db.QueryRow(sqlStr, bar).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("LDAP query results: nothing found for query: " + sqlStr))
		return
	}

	output, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func doSomething(param string) string {
	bar := "safe!"
	map89109 := make(map[string]interface{})
	map89109["keyA-89109"] = "a_Value"
	map89109["keyB-89109"] = param
	map89109["keyC"] = "another_Value"
	bar = map89109["keyB-89109"].(string)
	bar = map89109["keyA-89109"].(string)

	return bar
}
