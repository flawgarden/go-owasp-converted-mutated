//Original file region: null, null, null, null
//Mutated file region: null, null, null, null
//Semgrep original results: []
//Gosec original results: [89]
//CodeQL original results: [89]
//Snyk original results: []
//-------------
//Gosec analysis results: [89, 703]
//CodeQL analysis results: []
//Semgrep analysis results: []
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest02531/BenchmarkTest02531.go
//Original file CWE's: [89]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/virtuality/default.tmt with name binary_op_interface_default2_negative 
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

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", source)
	orm.Debug = true
}

type BenchmarkTest02531Controller struct {
	beego.Controller
}

func (c *BenchmarkTest02531Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02531Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html; charset=UTF-8")
	values := c.GetStrings("BenchmarkTest02531")
	param := ""
	if len(values) > 0 {
		param = values[0]
	}

var a12341 BinaryOpInterfaceDefault2 = &BinaryOpInterfaceDefault2Implementation{}
param = a12341.InterfaceCall(param, "")

	bar := doSomething(param)
	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME=? and PASSWORD='%s'", bar)

	connection, err := sql.Open("mysql", source)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
	defer connection.Close()

	statement, err := connection.Prepare(sqlStr)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
	defer statement.Close()

	rows, err := statement.Query("foo")
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Username, &user.Password); err != nil {
			c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
			return
		}
		users = append(users, user)
	}
	output, err := json.Marshal(users)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
	c.Ctx.ResponseWriter.Write(output)
}

func doSomething(param string) string {
	num := 106
	if (7*42)-num > 200 {
		return "This should never happen"
	}
	return param
}
