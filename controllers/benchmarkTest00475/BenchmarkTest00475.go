package controllers

import (
	"encoding/json"
	"net/http"

	"xorm.io/xorm"
)

type BenchmarkTest00475 struct {
	DB *xorm.Engine
}

func (b *BenchmarkTest00475) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	r.ParseForm()
	param := r.FormValue("BenchmarkTest00475")
	bar := ""

	num := 196
	if (500/42)+num > 200 {
		bar = param
	} else {
		bar = "This should never happen"
	}

	_, err := b.DB.Exec("INSERT INTO logs(message) VALUES(?)", bar)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	responseData := map[string]string{"message": bar}
	output, err := json.Marshal(responseData)
	if err != nil {
		http.Error(w, "JSON error", http.StatusInternalServerError)
		return
	}

	w.Write(output)
}

func (b *BenchmarkTest00475) InitDB(dataSource string) error {
	var err error
	b.DB, err = xorm.NewEngine("mysql", dataSource)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	b := &BenchmarkTest00475{}
	err := b.InitDB("root:password@tcp(127.0.0.1:3306)/goseccode")
	if err != nil {
		panic(err)
	}
	http.Handle("/xss-00/BenchmarkTest00475", b)
	http.ListenAndServe(":8080", nil)
}
