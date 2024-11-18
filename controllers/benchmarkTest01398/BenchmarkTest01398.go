package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-sec-code/models"
	"net/http"
)

type BenchmarkTest01398 struct {
}

func (bt *BenchmarkTest01398) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		bt.doPost(w, r)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (bt *BenchmarkTest01398) doPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := ""

	for name, values := range r.Form {
		for _, value := range values {
			if value == "BenchmarkTest01398" {
				param = name
				break
			}
		}
	}

	bar := new(Test).doSomething(r, param)

	source := "root:password@tcp(127.0.0.1:3306)/goseccode"
	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id=%s", bar)
	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	output, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Error marshaling JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

type Test struct {
}

func (t *Test) doSomething(request *http.Request, param string) string {
	a37227 := param
	b37227 := a37227 + " SafeStuff"
	b37227 = b37227[:len(b37227)-5] + "Chars"
	map37227 := make(map[string]interface{})
	map37227["key37227"] = b37227
	c37227 := map37227["key37227"].(string)
	d37227 := c37227[:len(c37227)-1]
	e37227 := string([]byte(d37227)) // Simulating B64 encode/decode
	f37227 := e37227[:len(e37227)-len(e37227)%5]
	return f37227
}
