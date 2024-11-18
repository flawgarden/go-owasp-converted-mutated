package controllers

import (
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01489 struct{}

func (bt *BenchmarkTest01489) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		bt.doPost(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (bt *BenchmarkTest01489) doPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest01489")
	if param == "" {
		param = ""
	}

	bar := new(Test).doSomething(r, param)

	// Encrypt and store the result
	output := []byte(bar)
	file, err := os.OpenFile("passwordFile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(w, "Could not open file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	if _, err := file.Write([]byte("secret_value=" + string(output) + "\n")); err != nil {
		http.Error(w, "Could not write to file", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Sensitive value: '" + string(output) + "' encrypted and stored<br/>"))
}

type Test struct{}

func (t *Test) doSomething(r *http.Request, param string) string {
	bar := ""
	num := 106
	if (7*42)-num > 200 {
		bar = "This should never happen"
	} else {
		bar = param
	}
	return bar
}

func main() {
	http.Handle("/crypto-01/BenchmarkTest01489", &BenchmarkTest01489{})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
