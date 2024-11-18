package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"go-sec-code/models"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type SqlInjectionVuln1Controller struct {
	web.Controller
}

func (c *SqlInjectionVuln1Controller) Get() {
	c.ServeHTTP(c.Ctx.ResponseWriter, c.Ctx.Request)
}

func (c *SqlInjectionVuln1Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, "Database connection failed", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var id string
	names := r.Header["Some-Header"]
	if len(names) > 0 {
		id = names[0]
	}

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id='%s'", strings.Replace(id, "'", "''", -1))
	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		http.Error(w, "Query execution failed", http.StatusInternalServerError)
		return
	}

	output, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "JSON marshaling failed", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}
