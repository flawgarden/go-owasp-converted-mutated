package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"go-sec-code/models"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01411 struct{}

func (b *BenchmarkTest01411) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	param := ""
	for name, values := range params {
		for _, value := range values {
			if value == "BenchmarkTest01411" {
				param = name
				break
			}
		}
	}

	bar := b.doSomething(param)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id=%s", bar)
	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		http.Error(w, "Query error", http.StatusInternalServerError)
		return
	}

	output, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "JSON marshal error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func (b *BenchmarkTest01411) doSomething(param string) string {
	bar := param
	if (7*42)-86 > 200 {
		bar = "This_should_always_happen"
	}
	return bar
}

func TestBenchmark(t *testing.T) {
	handler := &BenchmarkTest01411{}
	req, err := http.NewRequest("GET", "/hash-01/BenchmarkTest01411?BenchmarkTest01411=value", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"Id":0,"Username":"","Password":""}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
