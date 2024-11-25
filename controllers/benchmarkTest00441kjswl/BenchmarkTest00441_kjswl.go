//Original file region: null, null, null, null
//Mutated file region: null, null, null, null
//Semgrep original results: [89]
//Gosec original results: [89]
//CodeQL original results: [89]
//Snyk original results: []
//-------------
//Semgrep analysis results: [89]
//Gosec analysis results: [89, 703]
//CodeQL analysis results: []
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest00441/BenchmarkTest00441.go
//Original file CWE's: [89]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/assignment.tmt with name assignment_negative 
//Used extensions: 
//Program:
package controllers

import (
	"database/sql"
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00441Controller struct {
	beego.Controller
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", source)
	orm.Debug = true
}

func (c *BenchmarkTest00441Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest00441")

	if param == "" {
		param = ""
	}

tmpUnique42 := ""
param = tmpUnique42

	// Assume that createThing() returns a struct that has the method doSomething()
	thing := createThing()
	bar := thing.doSomething(param)

	sqlStr := fmt.Sprintf("INSERT INTO users (username, password) VALUES ('foo','%s')", bar)

	statement, err := sql.Open("mysql", source)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
	defer statement.Close()

	_, err = statement.Exec(sqlStr)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}

	c.Ctx.ResponseWriter.Write([]byte("Update complete"))
}

func createThing() *ThingInterface {
	// Dummy implementation for compilation
	return &ThingInterface{}
}

type ThingInterface struct{}

func (t *ThingInterface) doSomething(param string) string {
	return param // Dummy implementation
}

func main() {
	beego.Router("/sqli-00/BenchmarkTest00441", &BenchmarkTest00441Controller{})
	beego.Run()
}