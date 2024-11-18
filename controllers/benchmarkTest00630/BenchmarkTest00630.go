package controllers

import (
	"database/sql"
	"fmt"

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

type BenchmarkTest00630Controller struct {
	beego.Controller
}

func (c *BenchmarkTest00630Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00630Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest00630")
	if param == "" {
		param = ""
	}

	var bar string
	guess := "ABC"
	switchTarget := guess[2]

	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bobs_your_uncle"
	case 'C', 'D':
		bar = param
	default:
		bar = "bobs_your_uncle"
	}

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	filter := fmt.Sprintf("(&(objectclass=person)(uid=%s))", bar)

	// Simulated LDAP query results for demonstration purposes
	found := false
	if bar != "" { // Simulate a found result if bar is not empty
		// Normally, LDAP querying code would go here
		response := fmt.Sprintf("LDAP query results:<br>Record found with name %s<br>Address: Some Address<br>", bar)
		c.Ctx.ResponseWriter.Write([]byte(response))
		found = true
	}

	if !found {
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("LDAP query results: nothing found for query: %s", filter)))
	}
}
