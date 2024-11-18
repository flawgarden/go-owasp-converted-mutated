package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-sec-code/models"
	"log"
	"net/http"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00675 struct{}

func (b *BenchmarkTest00675) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	r.ParseForm()
	param := r.FormValue("BenchmarkTest00675")
	if param == "" {
		param = ""
	}

	var bar string
	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}

	sqlStr := fmt.Sprintf("SELECT * FROM USERS WHERE USERNAME=? AND PASSWORD='%s'", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		log.Println("Error connecting to database:", err)
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		log.Println("Error preparing statement:", err)
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec("foo")
	if err != nil {
		log.Println("Error executing statement:", err)
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}

	results, err := stmt.Query()
	if err != nil {
		log.Println("Error querying results:", err)
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer results.Close()

	var users []models.User
	for results.Next() {
		var user models.User
		if err := results.Scan(&user.Id, &user.Username, &user.Password); err != nil {
			log.Println("Error scanning results:", err)
			http.Error(w, "Error processing request.", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	output, err := json.Marshal(users)
	if err != nil {
		log.Println("Error marshalling results:", err)
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
	w.Write(output)
}

func main() {
	http.Handle("/sqli-01/BenchmarkTest00675", &BenchmarkTest00675{})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
