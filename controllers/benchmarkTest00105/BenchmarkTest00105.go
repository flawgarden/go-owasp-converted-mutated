package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"go-sec-code/models"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00105 struct {
}

func (b *BenchmarkTest00105) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.SetCookie(w, &http.Cookie{
			Name:     "BenchmarkTest00105",
			Value:    "bar",
			Path:     r.URL.Path,
			Secure:   true,
			HttpOnly: true,
		})
		http.ServeFile(w, r, "sqli-00/BenchmarkTest00105.html")
		return
	}

	if r.Method == http.MethodPost {
		cookies := r.Cookies()
		param := "noCookieValueSupplied"

		for _, cookie := range cookies {
			if cookie.Name == "BenchmarkTest00105" {
				param, _ = url.QueryUnescape(cookie.Value)
				break
			}
		}

		bar := param
		num := 106
		if (7*18)+num > 200 {
			bar = "This_should_always_happen"
		}

		sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME='foo' and PASSWORD='%s'", bar)

		db, err := sql.Open("mysql", source)
		if err != nil {
			http.Error(w, "Error processing request.", http.StatusInternalServerError)
			return
		}
		defer db.Close()

		user := models.User{}
		err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)

		if err != nil {
			http.Error(w, "Error processing request.", http.StatusInternalServerError)
			return
		}

		output, err := json.Marshal(user)
		if err != nil {
			http.Error(w, "Error processing request.", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(output)
	}
}
