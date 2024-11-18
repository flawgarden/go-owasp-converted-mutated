package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"go-sec-code/models"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02186 struct {
	web.Controller
}

func (c *BenchmarkTest02186) Get() {
	c.Post()
}

func (c *BenchmarkTest02186) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02186")
	if param == "" {
		param = ""
	}

	bar := doSomething(param)

	sqlStr := fmt.Sprintf("SELECT * FROM USERS WHERE USERNAME='foo' AND PASSWORD='%s'", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec(sqlStr)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
		return
	}

	user := models.User{}
	err = db.QueryRow("SELECT * FROM USERS WHERE USERNAME='foo' AND PASSWORD=?", bar).Scan(&user.Id, &user.Username, &user.Password)
	if err == nil {
		output, err := json.Marshal(user)
		if err == nil {
			c.Ctx.ResponseWriter.Write(output)
			return
		}
	}
	http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
}

func doSomething(param string) string {
	// Логика обработки param
	return param
}
