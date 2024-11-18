package controllers

import (
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00171 struct {
}

func (b *BenchmarkTest00171) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.Header.Get("BenchmarkTest00171")
	param = strings.Join(strings.Fields(param), " ")

	bar := "safe!"
	map40534 := make(map[string]interface{})
	map40534["keyA-40534"] = "a_Value"
	map40534["keyB-40534"] = param
	map40534["keyC"] = "another_Value"
	bar = map40534["keyB-40534"].(string)
	bar = map40534["keyA-40534"].(string)

	cmd := ""
	a1 := ""
	a2 := ""
	var args []string
	if strings.Contains(strings.ToLower(r.UserAgent()), "windows") {
		a1 = "cmd.exe"
		a2 = "/c"
		cmd = "echo "
		args = []string{a1, a2, cmd + bar}
	} else {
		a1 = "sh"
		a2 = "-c"
		cmd = "ls "
		args = []string{a1, a2, cmd + bar}
	}

	argsEnv := []string{"foo=bar"}

	runtime := &Runtime{args, argsEnv}

	if err := runtime.Exec(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

type Runtime struct {
	args    []string
	argsEnv []string
}

func (r *Runtime) Exec() error {
	// Implement the exec logic here, potentially using os/exec
	return nil
}
