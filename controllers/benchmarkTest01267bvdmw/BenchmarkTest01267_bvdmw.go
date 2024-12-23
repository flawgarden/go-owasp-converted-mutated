package controllers

import (
"fmt"
"net/http"
)

type BenchmarkTest01267 struct{}

func (b *BenchmarkTest01267) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()
	param := r.FormValue("BenchmarkTest01267")
	if param == "" {
		param = ""
	}

generator := makeMessageGeneratorBroken(param)()
param = generator("fixed_string")

	bar := new(Test).doSomething(r, param)

	w.Header().Set("X-XSS-Protection", "0")
	w.Write([]byte(bar))
}

type Test struct{}

func (t *Test) doSomething(r *http.Request, param string) string {
	bar := "safe!"
	map53210 := make(map[string]interface{})
	map53210["keyA-53210"] = "a-Value"
	map53210["keyB-53210"] = param
	map53210["keyC"] = "another-Value"
	bar = map53210["keyB-53210"].(string)

	return bar
}

func main() {
	http.Handle("/xss-02/BenchmarkTest01267", &BenchmarkTest01267{})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func makePrefixer(prefix string) func(string) string {
    return func(value string) string {
        return fmt.Sprintf("%s%s", prefix, value)
    }
}

func makeMessageGenerator(prefix string) func() func(string) string {
	return func() func(string) string {
		return func(name string) string {
			return prefix + name
		}
	}
}

func makeMessageGeneratorBroken(prefix string) func() func(string) string {
	return func() func(string) string {
		return func(name string) string {
			return name
		}
	}
}


