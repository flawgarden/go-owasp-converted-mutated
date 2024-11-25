//Original file region: null, null, null, null
//Mutated file region: null, null, null, null
//Semgrep original results: [89]
//Gosec original results: [89]
//CodeQL original results: []
//Snyk original results: [89]
//-------------
//Semgrep analysis results: [1004, 89, 79]
//Gosec analysis results: [89, 703]
//Snyk analysis results: [1004]
//CodeQL analysis results: []
//Original file name: controllers/benchmarkTest00109/BenchmarkTest00109.go
//Original file CWE's: [89]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/concurrency/concurrency.tmt with name null_and_restore_condvar_positive 
//Used extensions: 
//Program:
package controllers

import (
"database/sql"
"encoding/json"
"fmt"
"net/http"
"net/url"
"go-sec-code/models"
_ "github.com/go-sql-driver/mysql"
"sync"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00109 struct{}

func (b *BenchmarkTest00109) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.SetCookie(w, &http.Cookie{
			Name:   "BenchmarkTest00109",
			Value:  "bar",
			Path:   r.RequestURI,
			MaxAge: 60 * 3,
			Secure: true,
		})
		http.ServeFile(w, r, "sqli-00/BenchmarkTest00109.html")
		return
	}

	if r.Method == http.MethodPost {
		cookies := r.Cookies()
		param := "noCookieValueSupplied"
		for _, cookie := range cookies {
			if cookie.Name == "BenchmarkTest00109" {
				param = cookie.Value
				break
			}
		}

		bar := ""
		if param != "" {
			decodedParam, err := url.QueryUnescape(param)
			if err == nil {
				bar = string(decodedParam)
			}
		}

a := NewNullAndRestore(bar)
var wg sync.WaitGroup
wg.Add(2)
go func() {
    defer wg.Done()
	a.NullMethod()
}()
go func() {
    defer wg.Done()
	a.Restore()
}()
wg.Wait()

bar = a.Get()

		sqlStr := fmt.Sprintf("SELECT * FROM USERS WHERE USERNAME='foo' AND PASSWORD='%s'", bar)

		db, err := sql.Open("mysql", source)
		if err != nil {
			http.Error(w, "Error processing request.", http.StatusInternalServerError)
			return
		}
		defer db.Close()

		var user models.User
		err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
		if err != nil {
			http.Error(w, "Error processing request.", http.StatusInternalServerError)
			return
		}

		output, err := json.Marshal(user)
		if err == nil {
			w.Header().Set("Content-Type", "application/json")
			w.Write(output)
		}
	}
}
