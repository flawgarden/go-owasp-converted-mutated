package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"go-sec-code/models"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type SqlInjectionVuln1Handler struct{}

func (h *SqlInjectionVuln1Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("select * from user where id=%s", id)
	user := models.User{}
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
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func main() {
	http.Handle("/vuln", &SqlInjectionVuln1Handler{})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
