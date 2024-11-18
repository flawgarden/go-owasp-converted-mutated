package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-sec-code/models"
	"net/http"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02284Controller struct {
	web.Controller
}

func (c *BenchmarkTest02284Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02284Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02284")
	bar := doSomething(c.Ctx.Request, param)

	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME='foo' and PASSWORD='%s'", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	statement, err := db.Prepare(sqlStr)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer statement.Close()

	rows, err := statement.Query()
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Username, &user.Password); err != nil {
			http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}
	output, err := json.Marshal(users)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
		return
	}
	c.Ctx.ResponseWriter.Write(output)
}

func doSomething(request *http.Request, param string) string {
	num := 106
	bar := param
	if (7*42)-num > 200 {
		bar = "This should never happen"
	}
	return bar
}
