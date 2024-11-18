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

type BenchmarkTest02293Controller struct {
	beego.Controller
}

func (c *BenchmarkTest02293Controller) Get() {
	c.DoPost()
}

func (c *BenchmarkTest02293Controller) Post() {
	c.DoPost()
}

func (c *BenchmarkTest02293Controller) DoPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := ""
	flag := true
	names := c.Ctx.Request.URL.Query()
	for name, values := range names {
		if flag {
			for _, value := range values {
				if value == "BenchmarkTest02293" {
					param = name
					flag = false
				}
			}
		}
	}

	bar := doSomething(param)

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	sqlStr := fmt.Sprintf("select * from user where id='%s'", bar)
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
	a := param
	b := a + " SafeStuff"
	b = b[:len(b)-1] + "Chars"
	m := map[string]interface{}{"key": b}
	c := m["key"].(string)
	d := c[:len(c)-1]
	e := d // Simplified for demonstration
	return e
}
