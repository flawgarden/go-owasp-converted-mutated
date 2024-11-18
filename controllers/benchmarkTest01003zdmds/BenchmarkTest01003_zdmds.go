//Semgrep original results: [89]
//Gosec original results: [89]
//CodeQL original results: [89]
//Snyk original results: []
//-------------
//Semgrep analysis results: [1004, 89]
//Gosec analysis results: [89]
//CodeQL analysis results: []
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest01003/BenchmarkTest01003.go
//Original file CWE's: [89]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/virtuality/class.tmt with name base_binary_op_negative 
//Used extensions: 
//Program:
package controllers

import (
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01003 struct {
}

func (b *BenchmarkTest01003) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	http.SetCookie(w, &http.Cookie{
		Name:   "BenchmarkTest01003",
		Value:  "bar",
		MaxAge: 60 * 3,
		Secure: true,
		Path:   r.URL.Path,
		Domain: r.Host,
	})

	http.ServeFile(w, r, "sqli-02/BenchmarkTest01003.html")
	if r.Method == http.MethodPost {
		b.handlePost(w, r)
	}
}

func (b *BenchmarkTest01003) handlePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	cookie, err := r.Cookie("BenchmarkTest01003")
	param := "noCookieValueSupplied"
	if err == nil {
		param = cookie.Value
	}

	bar := new(Test).doSomething(param)

var a12341 BaseBinaryOpClass = &DerivedBinaryOpClassDefault{}
bar = a12341.VirtualCall(bar, bar)

	sqlStr := "SELECT * from USERS where USERNAME=? and PASSWORD='" + bar + "'"
	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	statement, err := db.Prepare(sqlStr)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer statement.Close()

	if _, err := statement.Exec("foo"); err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	bar := ""
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the param value
	}
	return bar
}
