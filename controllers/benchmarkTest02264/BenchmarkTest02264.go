package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	"go-sec-code/models"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02264Controller struct {
	web.Controller
}

func (c *BenchmarkTest02264Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02264Controller) Post() {
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest02264")
	bar := doSomething(param)

	sqlStatement := "{call " + bar + "}"
	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query(sqlStatement)
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
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func doSomething(param string) string {
	valuesList := []string{"safe", param, "moresafe"}
	valuesList = append(valuesList[:0], valuesList[1:]...) // remove the 1st safe value
	return strings.Join(valuesList[0:1], "")
}
