package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type BenchmarkTest01007 struct {
	DB *sql.DB
}

func (b *BenchmarkTest01007) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		b.handleGet(w, r)
	case http.MethodPost:
		b.handlePost(w, r)
	}
}

func (b *BenchmarkTest01007) handleGet(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "BenchmarkTest01007",
		Value:  "bar",
		MaxAge: 180,
		Secure: true,
		Path:   r.RequestURI,
		Domain: r.Host,
	})
	http.ServeFile(w, r, "sqli-02/BenchmarkTest01007.html")
}

func (b *BenchmarkTest01007) handlePost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := "noCookieValueSupplied"
	if cookies := r.Cookies(); cookies != nil {
		for _, cookie := range cookies {
			if cookie.Name == "BenchmarkTest01007" {
				param = cookie.Value
				break
			}
		}
	}

	bar := b.doSomething(r, param)

	sqlStr := fmt.Sprintf("SELECT userid FROM USERS WHERE USERNAME='foo' AND PASSWORD='%s'", bar)
	results, err := b.queryForMap(sqlStr)
	if err != nil {
		http.Error(w, "No results returned for query", http.StatusInternalServerError)
		return
	}

	output, _ := json.Marshal(results)
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func (b *BenchmarkTest01007) doSomething(r *http.Request, param string) string {
	var bar string
	guess := "ABC"
	switchTarget := guess[2]

	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bobs_your_uncle"
	case 'C', 'D':
		bar = param
	default:
		bar = "bobs_your_uncle"
	}

	return bar
}

func (b *BenchmarkTest01007) queryForMap(sqlStr string) (map[string]interface{}, error) {
	row := b.DB.QueryRow(sqlStr)
	var userid interface{}
	err := row.Scan(&userid)
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{"userid": userid}, nil
}

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/goseccode")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	b := &BenchmarkTest01007{DB: db}
	http.Handle("/sqli-02/BenchmarkTest01007", b)
	http.ListenAndServe(":8080", nil)
}
