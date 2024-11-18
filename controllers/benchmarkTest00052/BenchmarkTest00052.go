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

type BenchmarkTest00052 struct {
	web.Controller
}

func (c *BenchmarkTest00052) Get() {
	c.Post()
}

func (c *BenchmarkTest00052) Post() {
	response := c.Ctx.ResponseWriter
	response.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest00052")

	sqlQuery := fmt.Sprintf("{call %s}", param)

	conn, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	rows, err := conn.Query(sqlQuery)
	if err != nil {
		http.Error(response, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var results []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Username, &user.Password); err != nil {
			http.Error(response, "Error processing request.", http.StatusInternalServerError)
			return
		}
		results = append(results, user)
	}

	output, err := json.Marshal(results)
	if err != nil {
		http.Error(response, "Error processing request.", http.StatusInternalServerError)
		return
	}
	response.Write(output)
}
