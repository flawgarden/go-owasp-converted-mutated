package controllers

import (
	"fmt"
	"net/http"

	"xorm.io/xorm"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01847 struct {
	engine *xorm.Engine
}

func NewBenchmarkTest01847() *BenchmarkTest01847 {
	engine, err := xorm.NewEngine("mysql", source)
	if err != nil {
		panic(err)
	}
	return &BenchmarkTest01847{engine: engine}
}

func (b *BenchmarkTest01847) Get(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "BenchmarkTest01847",
		Value:  "someSecret",
		Path:   r.URL.Path,
		Secure: true,
		MaxAge: 60 * 3,
	})

	http.ServeFile(w, r, "benchmark.html")
}

func (b *BenchmarkTest01847) Post(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("BenchmarkTest01847")
	param := "noCookieValueSupplied"
	if err == nil {
		param = cookie.Value
	}

	bar := doSomething(param)


	response := fmt.Sprintf("Sensitive value '%s' hashed and stored\n", bar)
	w.Write([]byte(response))
	w.Write([]byte("Hash Test executed"))
}

func hashData(input string) string {
	// Implement your hashing logic here
	return input
}

func doSomething(param string) string {
	bar := ""
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = append(valuesList[:0], valuesList[1:]...) // remove the 1st safe value
		bar = valuesList[0]                                    // get the param value
	}
	return bar
}
