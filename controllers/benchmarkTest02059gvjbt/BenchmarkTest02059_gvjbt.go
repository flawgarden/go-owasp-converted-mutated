//Semgrep original results: []
//Gosec original results: [78]
//CodeQL original results: [78]
//Snyk original results: [78]
//-------------
//Gosec analysis results: [676, 78, 703]
//CodeQL analysis results: [78, 116, 79]
//Snyk analysis results: []
//Semgrep analysis results: []
//Original file name: controllers/benchmarkTest02059/BenchmarkTest02059.go
//Original file CWE's: [78]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/collections/queue.tmt with name queue_remove_if_positive 
//Used extensions: MACRO_Create_Queue -> ~[MACRO_QueueName]~ := list.New() | MACRO_Add_Fixed_CONST_ToQueue -> ~[MACRO_QueueName]~.PushBack(~[CONST_~[TYPE@1]~@1]~) | MACRO_Add_Fixed_VAR_ToQueue -> ~[MACRO_QueueName]~.PushBack(~[VAR_~[TYPE@1]~@1]~) | MACRO_QueueName -> queue787231 | MACRO_QueueName -> queue787231 | MACRO_QueueName -> queue787231 | MACRO_QueueName -> queue787231 | MACRO_QueueName -> queue787231 | MACRO_QueueName -> queue787231
//Program:
package controllers

import (
"net/http"
"os/exec"
"strings"
"container/list"
)

type BenchmarkTest02059 struct{}

func (b *BenchmarkTest02059) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.Header.Get("BenchmarkTest02059")

queue787231 := list.New()
queue787231.PushBack("iDfQe")
queue787231.PushBack(param)
for e := queue787231.Front(); e != nil; {
    next := e.Next()
    if e.Value == "iDfQe" {
        queue787231.Remove(e)
    }
    e = next
}
param = queue787231.Front().Value.(string)

	param = decode(param)

	bar := doSomething(r, param)

	var a1, a2 string
	if strings.Contains(strings.ToLower(r.UserAgent()), "windows") {
		a1 = "cmd.exe"
		a2 = "/c"
	} else {
		a1 = "sh"
		a2 = "-c"
	}
	args := []string{a1, a2, "echo " + bar}

	cmd := exec.Command(args[0], args[1:]...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		http.Error(w, "Error executing command", http.StatusInternalServerError)
		return
	}
	w.Write(output)
}

func doSomething(r *http.Request, param string) string {
	bar := "safe!"
	collection := map[string]interface{}{
		"keyA-36421": "a-Value",
		"keyB-36421": param,
		"keyC":       "another-Value",
	}
	bar = collection["keyB-36421"].(string)
	return bar
}

func decode(value string) string {
	if len(value) == 0 {
		return ""
	}
	decoded, err := exec.Command("echo", value).Output()
	if err != nil {
		return value
	}
	return string(decoded)
}

func main() {
	http.Handle("/cmdi-02/BenchmarkTest02059", &BenchmarkTest02059{})
	http.ListenAndServe(":8080", nil)
}
