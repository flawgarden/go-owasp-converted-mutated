package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01492 struct {
}

func (b *BenchmarkTest01492) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest01492")
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

	filter := fmt.Sprintf("(&(objectclass=person)(uid=%s))", bar)
	rows, err := db.Query("SELECT uid, street FROM users WHERE uid = ?", bar)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	found := false
	for rows.Next() {
		var uid, street string
		if err := rows.Scan(&uid, &street); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte(fmt.Sprintf("LDAP query results:<br>Record found with name %s<br>Address: %s<br>", uid, street)))
		found = true
	}
	if !found {
		w.Write([]byte(fmt.Sprintf("LDAP query results: nothing found for query: %s", filter)))
	}
}

func (b *BenchmarkTest01492) doSomething(param string) string {
	var bar string
	num := 106
	if (7*18)+num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}
	return bar
}

func main() {
	http.Handle("/ldapi-00/BenchmarkTest01492", &BenchmarkTest01492{})
	http.ListenAndServe(":8080", nil)
}
