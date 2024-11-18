package controllers

import (
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01582 struct{}

func (b *BenchmarkTest01582) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	id := r.URL.Query().Get("BenchmarkTest01582")
	if id == "" {
		id = "safe!"
	}

	bar := b.doSomething(id)

	hashValue, err := b.hash(bar)
	if err != nil {
		http.Error(w, "Hash error", http.StatusInternalServerError)
		return
	}

	file, err := os.OpenFile("passwordFile.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		http.Error(w, "File error", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	if _, err := file.WriteString(fmt.Sprintf("hash_value=%s\n", hashValue)); err != nil {
		http.Error(w, "File write error", http.StatusInternalServerError)
		return
	}

	if _, err := w.Write([]byte(fmt.Sprintf("Sensitive value '%s' hashed and stored<br/>", bar))); err != nil {
		http.Error(w, "Response write error", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Hash Test executed")
}

func (b *BenchmarkTest01582) doSomething(param string) string {
	bar := "safe!"
	map17367 := map[string]string{
		"keyA-17367": "a_Value",
		"keyB-17367": param,
		"keyC":       "another_Value",
	}
	bar = map17367["keyB-17367"]
	bar = map17367["keyA-17367"]
	return bar
}

func (b *BenchmarkTest01582) hash(input string) (string, error) {
	// Simulating hashing (in a real case, use a proper hashing function)
	return input, nil
}

func main() {
	http.Handle("/hash-01/BenchmarkTest01582", &BenchmarkTest01582{})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
