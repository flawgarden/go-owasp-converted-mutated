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

type BenchmarkTest02529 struct {
	beego.Controller
}

func (c *BenchmarkTest02529) Get() {
	c.Post()
}

func (c *BenchmarkTest02529) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	values := c.GetStrings("BenchmarkTest02529")
	var param string
	if len(values) > 0 {
		param = values[0]
	} else {
		param = ""
	}

	bar := doSomething(param)

	sqlStr := fmt.Sprintf("{call %s}", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query(sqlStr)
	if err != nil {
		c.Ctx.WriteString("Error processing request.")
		return
	}
	defer rows.Close()

	user := models.User{}
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Password)
		if err != nil {
			c.Ctx.WriteString("Error processing request.")
			return
		}
		output, err := json.Marshal(user)
		if err == nil {
			c.Ctx.ResponseWriter.Write(output)
		}
	}
}

func doSomething(param string) string {
	bar := "safe!"
	map32022 := make(map[string]interface{})
	map32022["keyA-32022"] = "a_Value"
	map32022["keyB-32022"] = param
	map32022["keyC"] = "another_Value"
	bar = map32022["keyB-32022"].(string)
	bar = map32022["keyA-32022"].(string)

	return bar
}
