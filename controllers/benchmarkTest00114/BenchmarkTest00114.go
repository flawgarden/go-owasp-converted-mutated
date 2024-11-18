package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00114 struct {
}

func (b *BenchmarkTest00114) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.SetCookie(w, &http.Cookie{
			Name:   "BenchmarkTest00114",
			Value:  "bar",
			MaxAge: 60 * 3,
			Secure: true,
			Path:   r.RequestURI,
			Domain: r.Host,
		})
		http.ServeFile(w, r, "sqli-00/BenchmarkTest00114.html")
		return
	}

	if r.Method == http.MethodPost {
		cookies := r.Cookies()
		param := "noCookieValueSupplied"
		for _, cookie := range cookies {
			if cookie.Name == "BenchmarkTest00114" {
				param = cookie.Value
				break
			}
		}

		bar := param
		num := 86
		if (7*42)-num > 200 {
			bar = "This_should_always_happen"
		}

		sqlStmt := fmt.Sprintf("INSERT INTO users (username, password) VALUES ('foo','%s')", bar)

		db, err := sql.Open("mysql", source)
		if err != nil {
			http.Error(w, "Error processing request.", http.StatusInternalServerError)
			return
		}
		defer db.Close()

		_, err = db.Exec(sqlStmt)
		if err != nil {
			http.Error(w, "Error processing request.", http.StatusInternalServerError)
			return
		}

		w.Write([]byte("Update complete"))
	}
}

func main() {
	http.Handle("/sqli-00/BenchmarkTest00114", &BenchmarkTest00114{})
	http.ListenAndServe(":8080", nil)
}
