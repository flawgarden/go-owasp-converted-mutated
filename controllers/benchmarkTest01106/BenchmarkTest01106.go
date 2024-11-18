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

type BenchmarkTest01106 struct {
	web.Controller
}

func (c *BenchmarkTest01106) Get() {
	c.Post()
}

func (c *BenchmarkTest01106) Post() {
	response := c.Ctx.ResponseWriter
	response.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.Ctx.Request.Header.Get("Custom-Header")

	if param == "" {
		http.Error(response, "No custom header provided", http.StatusBadRequest)
		return
	}

	bar := encodeForHTML(param)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(response, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id='%s'", bar)
	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		http.Error(response, "Query error", http.StatusInternalServerError)
		return
	}

	output, err := json.Marshal(user)
	if err != nil {
		http.Error(response, "JSON encoding error", http.StatusInternalServerError)
		return
	}

	response.Write(output)
}

func encodeForHTML(input string) string {
	// Simulating HTML encoding
	return fmt.Sprintf("%s_encoded", input)
}
