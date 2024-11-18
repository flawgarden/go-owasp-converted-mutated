package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01490 struct {
}

func (b *BenchmarkTest01490) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		b.doPost(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (b *BenchmarkTest01490) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest01490")
	if param == "" {
		param = ""
	}

	bar := b.doSomething(param)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	filter := fmt.Sprintf("(&(objectclass=person))(|(uid=%s)(street={0}))", bar)
	// Псевдокод, так как Go не поддерживает LDAP из коробки
	// Вставьте код взаимодействия с LDAP здесь

	fmt.Fprintf(w, "LDAP query filter: %s", filter) // Замена реального LDAP-запроса на вывод фильтра
}

func (b *BenchmarkTest01490) doSomething(param string) string {
	bar := "safe!"
	map58491 := make(map[string]interface{})
	map58491["keyA-58491"] = "a-Value"
	map58491["keyB-58491"] = param
	map58491["keyC"] = "another-Value"
	bar = map58491["keyB-58491"].(string)

	return bar
}
