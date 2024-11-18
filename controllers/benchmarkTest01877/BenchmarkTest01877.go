package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"go-sec-code/models"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

func main() {
	http.HandleFunc("/sqli-04/BenchmarkTest01877", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	cookie := http.Cookie{
		Name:   "BenchmarkTest01877",
		Value:  "verifyUserPassword%28%27foo%27%2C%27bar%27%29",
		MaxAge: 60 * 3,
		Secure: true,
		Path:   r.URL.Path,
		Domain: strings.Split(r.Host, ":")[0],
	}
	http.SetCookie(w, &cookie)

	if r.Method == http.MethodGet {
		http.ServeFile(w, r, "sqli-04/BenchmarkTest01877.html")
		return
	}

	param := "noCookieValueSupplied"
	cookies := r.Cookies()
	for _, c := range cookies {
		if c.Name == "BenchmarkTest01877" {
			param, _ = url.QueryUnescape(c.Value)
			break
		}
	}

	bar := doSomething(param)
	sqlQuery := fmt.Sprintf("CALL %s", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query(sqlQuery)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Username, &user.Password); err != nil {
			http.Error(w, "Error processing request.", http.StatusInternalServerError)
			return
		}
		output, err := json.Marshal(user)
		if err == nil {
			w.Write(output)
		}
	}
}

func doSomething(param string) string {
	bar := ""
	guess := "ABC"
	switchTarget := guess[1] // condition 'B'

	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bob"
	case 'C', 'D':
		bar = param
	default:
		bar = "bob's your uncle"
	}

	return bar
}
