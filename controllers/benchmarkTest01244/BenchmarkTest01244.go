package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-sec-code/models"
	"html"
	"log"
	"net/http"

	"github.com/beego/beego/v2/server/web"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

func init() {
	sql.Register("mysql", &mysql.MySQLDriver{})
}

type BenchmarkTest01244 struct {
	web.Controller
}

func (c *BenchmarkTest01244) Get() {
	c.Redirect("/hash-01/BenchmarkTest01244", http.StatusSeeOther)
}

func (c *BenchmarkTest01244) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest01244")
	if param == "" {
		param = ""
	}

	bar := new(Test).doSomething(c.Ctx.Request, param)

	db, err := sql.Open("mysql", source)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE username='%s'", bar)
	var user models.User
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		log.Fatal(err)
	}

	output, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

type Test struct{}

func (t *Test) doSomething(req *http.Request, param string) string {
	return html.EscapeString(param)
}
