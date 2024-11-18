//Semgrep original results: [89]
//Gosec original results: [89]
//CodeQL original results: [89]
//Snyk original results: [89]
//-------------
//Semgrep analysis results: [89, 79, 319]
//Gosec analysis results: [676, 89, 703]
//CodeQL analysis results: []
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest00590/BenchmarkTest00590.go
//Original file CWE's: [89]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/anonstructures.tmt with name anonymous_structure_with_nested_fields_negative 
//Used extensions: 
//Program:
package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

func sqlInjectionHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

car := struct {
    Make  string
    Model string
    Specs struct {
        Year int
        Color string
    }
}{
    Make:  "Toyota",
    Model: "X5 AMG",
    Specs: struct {
        Year  int
        Color string
    }{
        Year:  2020,
        Color: id,
    },
}

id = car.Make

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	sqlStr := fmt.Sprintf("select * from user where id=%s", id)
	user := struct {
		Id       int
		Username string
		Password string
	}{}
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
	http.HandleFunc("/sqli-01/BenchmarkTest00590", sqlInjectionHandler)
	http.ListenAndServe(":8080", nil)
}

func createPoint(x, y string) struct {
    X string
    Y string
} {
    return struct {
        X string
        Y string
    }{
        X: x,
        Y: y,
    }
}


