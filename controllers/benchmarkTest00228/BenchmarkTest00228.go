package controllers

import (
	"database/sql"
	"encoding/json"
	"go-sec-code/models"
	"net/http"

	"xorm.io/xorm"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type SqlInjectionVuln2Controller struct {
	xorm.Engine
}

func (c *SqlInjectionVuln2Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	sqlStr := "SELECT * FROM user WHERE id = ?"
	user := models.User{}
	err = db.QueryRow(sqlStr, id).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	output, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}
