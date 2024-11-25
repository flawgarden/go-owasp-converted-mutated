//Original file region: null, null, null, null
//Mutated file region: null, null, null, null
//Semgrep original results: []
//Gosec original results: [89]
//CodeQL original results: [89]
//Snyk original results: []
//-------------
//Gosec analysis results: [89, 703]
//CodeQL analysis results: [570, 571, 563]
//Semgrep analysis results: []
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest00438/BenchmarkTest00438.go
//Original file CWE's: [89]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/conditional/conditionswitch.tmt with name switch_fallthrough_negative 
//Used extensions: 
//Program:
package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-sec-code/models"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00438Controller struct {
	web.Controller
}

func (c *BenchmarkTest00438Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00438Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest00438")

value := 0
switch {
case value < 0:
    param = "fixed_string"
case value == 0:
    param = param + "_suffix"
    fallthrough
case value > 0:
    param = "fixed_string"
default:
    param = "Unknown"
}

	if param == "" {
		param = ""
	}

	bar := param
	num := 106
	if (7*42)-num > 200 {
		bar = "This should never happen"
	}

	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME='foo' and PASSWORD='%s'", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
	defer db.Close()

	statement, err := db.Prepare(sqlStr)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
	defer statement.Close()

	rows, err := statement.Query()
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
