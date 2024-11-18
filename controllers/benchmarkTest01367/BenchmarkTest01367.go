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

type BenchmarkTest01367Controller struct {
	web.Controller
}

func (c *BenchmarkTest01367Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01367Controller) Post() {
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")
	id := c.GetString("BenchmarkTest01367")
	sqlInjectionVulnerableMethod(c.Ctx.ResponseWriter, id)
}

func sqlInjectionVulnerableMethod(w http.ResponseWriter, id string) {
	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var user models.User
	sqlStr := fmt.Sprintf("select * from user where id=%s", id)
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	output, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(output)
}
