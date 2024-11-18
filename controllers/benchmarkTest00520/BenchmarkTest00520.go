package controllers

import (
	"encoding/xml"
	"net/http"
	"os"
	"path/filepath"
	// путь к вашей папке helpers
)

type Employees struct {
	Employees []Employee `xml:"Employee"`
}

type Employee struct {
	ID string `xml:"emplid,attr"`
}

func BenchmarkTest00520(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest00520")
	bar := ""

	num := 106
	if (7*18)+num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}

	filePath := filepath.Join("path", "to", "employees.xml") // укажите корректный путь к файлу
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "Error opening XML file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	var employees Employees
	decoder := xml.NewDecoder(file)
	if err := decoder.Decode(&employees); err != nil {
		http.Error(w, "Error decoding XML file", http.StatusInternalServerError)
		return
	}

	var result string
	for _, employee := range employees.Employees {
		if employee.ID == bar {
			result = employee.ID
			break
		}
	}

	_, err = w.Write([]byte("Your query results are: " + result + "<br/>"))
	if err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
	}
}
