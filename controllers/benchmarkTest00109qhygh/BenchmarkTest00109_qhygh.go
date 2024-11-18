// Semgrep original results: [89]
// Gosec original results: [89]
// CodeQL original results: []
// Snyk original results: [89]
// -------------
// Semgrep analysis results: [1004, 89, 79]
// Gosec analysis results: [89, 703]
// Snyk analysis results: [1004]
// CodeQL analysis results: []
// Original file name: controllers/benchmarkTest00109/BenchmarkTest00109.go
// Original file CWE's: [89]
// Original file kind: fail
// Mutation info: Insert template from templates-db/languages/go/sensitivity/collections/queue.tmt with name queue_contains_1_positive
// Used extensions: MACRO_Create_Queue -> ~[MACRO_QueueName]~ := list.New() | MACRO_Add_EXPR_ToQueue -> ~[MACRO_QueueName]~.PushBack(~[EXPR_~[TYPE@1]~]~) | MACRO_Add_Fixed_VAR_ToQueue -> ~[MACRO_QueueName]~.PushBack(~[VAR_~[TYPE@1]~@1]~) | MACRO_QueueName -> queue787231 | MACRO_QueueName -> queue787231 | MACRO_QueueName -> queue787231 | MACRO_QueueName -> queue787231
// Program:
package controllers

import (
	"container/list"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go-sec-code/models"
	"net/http"
	"net/url"
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

		queue787231 := list.New()
		queue787231.PushBack("ltjzs")
		queue787231.PushBack(param)
		value7845 := "nYEVm"
		for e := queue787231.Front(); e != nil; e = e.Next() {
			if e.Value == param {
				value7845 = e.Value.(string)
				break
			}
		}
		param = value7845

		bar := ""
		if param != "" {
			decodedParam, err := url.QueryUnescape(param)
			if err == nil {
				bar = string(decodedParam)
			}
		}

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
