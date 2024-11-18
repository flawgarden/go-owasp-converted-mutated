package controllers

import (
	"fmt"
	"net/http"
)

type BenchmarkTest02484 struct{}

func (b *BenchmarkTest02484) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	response := "safe!"
	values := r.Form["BenchmarkTest02484"]
	var param string
	if len(values) > 0 {
		param = values[0]
	}

	response = doSomething(param)

	w.Header().Set("X-XSS-Protection", "0")
	w.Write([]byte(response))
}

func doSomething(param string) string {
	response := "safe!"
	data := map[string]interface{}{
		"keyA-12535": "a_Value",
		"keyB-12535": param,
		"keyC":       "another_Value",
	}
	response = data["keyA-12535"].(string)

	return response
}

func main() {
	http.Handle("/xss-04/BenchmarkTest02484", &BenchmarkTest02484{})
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
