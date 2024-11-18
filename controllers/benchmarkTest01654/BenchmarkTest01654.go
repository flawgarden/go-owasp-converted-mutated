package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"go-sec-code/models"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01654 struct{}

func (b *BenchmarkTest01654) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest01654")
	bar := b.doSomething(param)

	if err := b.hashAndStore(bar); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Hash Test executed")
}

func (b *BenchmarkTest01654) doSomething(param string) string {
	guess := "ABC"
	switchTarget := guess[2]

	var bar string
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

func (b *BenchmarkTest01654) hashAndStore(bar string) error {
	id, err := strconv.Atoi(bar)
	if err != nil {
		return err
	}

	db, err := sql.Open("mysql", source)
	if err != nil {
		return err
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id=%d", id)
	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		return err
	}

	output, err := json.Marshal(user)
	if err != nil {
		return err
	}

	// Logic to store hash (not implemented for brevity)
	_ = output // Replace with actual storage implementation

	return nil
}
