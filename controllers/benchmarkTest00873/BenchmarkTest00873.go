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

type BenchmarkTest00873 struct {
	web.Controller
}

func (c *BenchmarkTest00873) Get() {
	c.Post()
}

func (c *BenchmarkTest00873) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest00873")

	bar := "safe!"
	map58847 := map[string]interface{}{
		"keyA-58847": "a_Value",
		"keyB-58847": param,
		"keyC":       "another_Value",
	}
	bar = map58847["keyB-58847"].(string)
	bar = map58847["keyA-58847"].(string)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM users WHERE username='%s'", bar)
	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "User query error", http.StatusInternalServerError)
		return
	}

	output, err := json.Marshal(user)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "JSON marshal error", http.StatusInternalServerError)
		return
	}
	c.Ctx.ResponseWriter.Write(output)
}
