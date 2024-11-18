//Semgrep original results: []
//Gosec original results: [78]
//CodeQL original results: [78]
//Snyk original results: [78]
//-------------
//Gosec analysis results: [78, 703]
//CodeQL analysis results: []
//Snyk analysis results: [79, 78]
//Semgrep analysis results: []
//Original file name: controllers/benchmarkTest02137/BenchmarkTest02137.go
//Original file CWE's: [78]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/closure.tmt with name simple_closure_counter_negative 
//Used extensions: 
//Program:
package controllers

import (
"net/http"
"os"
"os/exec"
"fmt"
)

type BenchmarkTest02137 struct{}

func (b *BenchmarkTest02137) Handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest02137")
	if param == "" {
		param = ""
	}

	bar := doSomething(param)

counter := func() func(str string) string {
    count := 0
    return func(str string) string {
        count++
        if count == 1 {
            return str
        } else {
            return "fixed_string"
        }

    }
}()
counter(bar)
bar = counter(bar)

	var cmd []string
	if isWindows() {
		cmd = []string{"cmd.exe", "/c", "echo " + bar}
	} else {
		cmd = []string{"sh", "-c", "echo " + bar}
	}

	output, err := exec.Command(cmd[0], cmd[1:]...).Output()
	if err != nil {
		http.Error(w, "Error executing command", http.StatusInternalServerError)
		return
	}

	w.Write(output)
}

func doSomething(param string) string {
	bar := ""
	num := 196
	if (500/42)+num > 200 {
		bar = param
	} else {
		bar = "This should never happen"
	}
	return bar
}

func isWindows() bool {
	return os.Getenv("OS") == "Windows_NT"
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


