//Semgrep original results: [89]
//Gosec original results: [89]
//CodeQL original results: [89]
//Snyk original results: []
//-------------
//Semgrep analysis results: [89]
//Gosec analysis results: [89, 703]
//CodeQL analysis results: []
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest00681/BenchmarkTest00681.go
//Original file CWE's: [89]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/types/duck.tmt with name duck_typing_multiple_attributes_negative 
//Used extensions: 
//Program:
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

type BenchmarkTest00681Controller struct {
	beego.Controller
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", source)
	orm.Debug = true
}

func (c *BenchmarkTest00681Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00681Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest00681")
	if param == "" {
		param = ""
	}

	bar := param
	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME='foo' and PASSWORD='%s'", bar)

d := NewFakeDuckWithAttribute(sqlStr)
sqlStr = MakeItQuackFieldAttr(d, "tmp_str")

	db, err := sql.Open("mysql", source)
	if err != nil {
		c.Data["json"] = "Error processing request."
		c.ServeJSON()
		return
	}
	defer db.Close()

	rows, err := db.Query(sqlStr)
	if err != nil {
		c.Data["json"] = "Error processing request."
		c.ServeJSON()
		return
	}
	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Username, &user.Password); err != nil {
			c.Data["json"] = "Error processing request."
			c.ServeJSON()
			return
		}
		users = append(users, user)
	}

	output, err := json.Marshal(users)
	if err != nil {
		c.Data["json"] = "Error processing request."
		c.ServeJSON()
		return
	}

	c.Ctx.ResponseWriter.Write(output)
}

func MakeItQuack(duck interface{ Quack(string) string }, arg string) string {
    return duck.Quack(arg)
}

func MakeItQuackAttr(duck interface{}, arg string) string {
    if d, ok := duck.(interface{ Quack(string) string }); ok {
        return d.Quack(arg)
    }
    return "fixed_string"
}

func MakeItQuackFieldAttr(duck interface{}, arg string) string {
	if d, ok := duck.(DuckWithAttribute); ok && d.constant == 42 {
		return d.Quack(arg)
	}
	return "fixed_string"
}


