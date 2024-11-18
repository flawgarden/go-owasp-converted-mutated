package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type SqlInjectionVuln1Controller struct {
	web.Controller
}

func (c *SqlInjectionVuln1Controller) Get() {
	c.doGet()
}

func (c *SqlInjectionVuln1Controller) Post() {
	c.doPost()
}

func (c *SqlInjectionVuln1Controller) doGet() {
	c.doPost()
}

func (c *SqlInjectionVuln1Controller) doPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json")

	id := c.GetString("BenchmarkTest00764")
	sqlStr := fmt.Sprintf("SELECT userid FROM USERS WHERE USERNAME='foo' AND PASSWORD='%s'", id)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var results int64
	err = db.QueryRow(sqlStr).Scan(&results)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error processing request", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{"results": results}
	output, err := json.Marshal(response)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error marshaling response", http.StatusInternalServerError)
		return
	}
	c.Ctx.ResponseWriter.Write(output)
}
