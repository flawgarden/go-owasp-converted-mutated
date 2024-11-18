package controllers

import (
	"fmt"
	"net/http"
)

type BenchmarkTest01445 struct{}

func (b *BenchmarkTest01445) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		b.doPost(w, r)
	} else if r.Method == http.MethodPost {
		b.doPost(w, r)
	}
}

func (b *BenchmarkTest01445) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	flag := true
	names := r.URL.Query()

	for name := range names {
		if flag {
			values := r.URL.Query()[name]
			for _, value := range values {
				if value == "BenchmarkTest01445" {
					param = name
					flag = false
				}
			}
		}
	}

	bar := b.doSomething(r, param)

	cmd := GetInsecureOSCommandString()

	argsEnv := []string{bar}
	runtime := Runtime{}

	if err := runtime.Exec(cmd, argsEnv); err != nil {
		fmt.Println("Problem executing cmdi - TestCase")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (b *BenchmarkTest01445) doSomething(r *http.Request, param string) string {
	bar := "safe!"
	map3083 := make(map[string]interface{})
	map3083["keyA-3083"] = "a_Value"
	map3083["keyB-3083"] = param
	map3083["keyC"] = "another_Value"
	bar = map3083["keyB-3083"].(string)
	bar = map3083["keyA-3083"].(string)

	return bar
}

func GetInsecureOSCommandString() string {
	// Implementation needed for your specific use case
	return "your-command-here"
}

type Runtime struct{}

func (r *Runtime) Exec(cmd string, argsEnv []string) error {
	// Replace this with actual command execution logic
	return nil
}
