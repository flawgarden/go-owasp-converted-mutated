package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01222 struct{}

func (b *BenchmarkTest01222) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	param := ""
	if header := r.Header.Get("BenchmarkTest01222"); header != "" {
		param = header

car := struct {
    Make  string
    Model string
    Specs struct {
        Year int
        Color string
    }
}{
    Make:  "Toyota",
    Model: "X5 AMG",
    Specs: struct {
        Year  int
        Color string
    }{
        Year:  2020,
        Color: param,
    },
}

param = car.Specs.Color

	}
	param, _ = url.QueryUnescape(param)

	bar := new(Test).doSomething(r, param)

	sqlStr := fmt.Sprintf("INSERT INTO users (username, password) VALUES ('foo','%s')", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec(sqlStr)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Update complete"))
}

type Test struct{}

func (t *Test) doSomething(r *http.Request, param string) string {
	bar := "safe!"
	map25738 := make(map[string]interface{})
	map25738["keyA-25738"] = "a-Value"
	map25738["keyB-25738"] = param
	map25738["keyC"] = "another-Value"
	bar = map25738["keyB-25738"].(string)
	return bar
}

func createPoint(x, y string) struct {
    X string
    Y string
} {
    return struct {
        X string
        Y string
    }{
        X: x,
        Y: y,
    }
}


