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

type SqlInjectionVuln1Controller struct {
	beego.Controller
}

func (c *SqlInjectionVuln1Controller) Get() {
	camanize := "BenchmarkTest02547"

	queryString := c.Ctx.Request.URL.Query().Encode()
	paramval := camanize + "="
	paramLoc := -1
	if queryString != "" {
		paramLoc = indexOf(queryString, paramval)
	}
	if paramLoc == -1 {
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("getQueryString() couldn't find expected parameter '%s' in query string.", camanize)))
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := indexOf(queryString, "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : ampersandLoc]
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

func indexOf(s, substr string) int {
	for i := range s {
		if string(s[i]) == substr {
			return i
		}
	}
	return -1
}

func doSomething(param string) string {
	return param // здесь должна быть реализация кодирования
}

func main() {
	beego.Router("/sql-injection-vuln1", &SqlInjectionVuln1Controller{})
	beego.Run()
}
