package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/url"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type SqlInjectionVuln2Controller struct {
	web.Controller
}

func (c *SqlInjectionVuln2Controller) Get() {
	c.Post()
}

func (c *SqlInjectionVuln2Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	if header := c.Ctx.Request.Header.Get("BenchmarkTest00018"); header != "" {
		param = header
	}

	param, err := url.QueryUnescape(param)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error decoding parameter", http.StatusBadRequest)
		return
	}

	sqlStr := fmt.Sprintf("INSERT INTO users (username, password) VALUES ('foo','%s')", param)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error opening database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	statement, err := db.Prepare(sqlStr)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error preparing statement", http.StatusInternalServerError)
		return
	}
	defer statement.Close()

	_, err = statement.Exec()
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error executing statement", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(c.Ctx.ResponseWriter, "Inserted successfully")
}

func main() {
	web.Router("/sqli-00/BenchmarkTest00018", &SqlInjectionVuln2Controller{})
	web.Run()
}
