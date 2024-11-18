package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

func main() {
	http.HandleFunc("/sqli-06/BenchmarkTest02648", BenchmarkTest02648)
	http.ListenAndServe(":8080", nil)
}

func BenchmarkTest02648(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest02648")
	bar := doSomething(param)

	sqlStr := fmt.Sprintf("SELECT * FROM USERS WHERE USERNAME='foo' AND PASSWORD='%s'", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec(sqlStr)
	if err != nil {
		http.Error(w, "Error processing request", http.StatusInternalServerError)
		return
	}

	// Simulated output
	json.NewEncoder(w).Encode(map[string]string{"result": "success"})
}

func doSomething(param string) string {
	bar := "safe!"
	map91488 := make(map[string]interface{})
	map91488["keyA-91488"] = "a_Value"
	map91488["keyB-91488"] = param
	map91488["keyC"] = "another_Value"
	bar = map91488["keyB-91488"].(string)
	bar = map91488["keyA-91488"].(string)

	return bar
}
