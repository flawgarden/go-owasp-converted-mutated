package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-sec-code/models"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

func main() {
	http.HandleFunc("/crypto-01/BenchmarkTest01016", BenchmarkTest01016)
	http.ListenAndServe(":8080", nil)
}

func BenchmarkTest01016(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.Header.Get("BenchmarkTest01016")

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("select * from user where id='%s'", param)
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
