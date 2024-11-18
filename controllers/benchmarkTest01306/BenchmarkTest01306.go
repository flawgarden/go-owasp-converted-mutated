package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-sec-code/models"
	"net/http"

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

type BenchmarkTest01306Controller struct {
	beego.Controller
}

func (c *BenchmarkTest01306Controller) Get() {
	c.DoPost()
}

func (c *BenchmarkTest01306Controller) Post() {
	c.DoPost()
}

func (c *BenchmarkTest01306Controller) DoPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest01306")
	if param == "" {
		param = ""
	}

	bar := new(Test).doSomething(param)

	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME='%s' and PASSWORD='%s'", "foo", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query(sqlStr)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	user := models.User{}
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Username, &user.Password)
		if err != nil {
			http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
			return
		}
	}

	output, err := json.Marshal(user)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
		return
	}
	c.Ctx.ResponseWriter.Write(output)
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	if param != "" {
		return string(param) // Псевдокод для обработки параметра
	}
	return ""
}
