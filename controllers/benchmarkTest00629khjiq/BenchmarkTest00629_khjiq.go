package controllers

import (
	"net/http"
	"os"
	"path/filepath"
)

type BenchmarkTest00629 struct{}

func (b *BenchmarkTest00629) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest00629")

car := struct {
    Make  string
    Model string
    Specs struct {
        Year int
        Color string
    }
}{
    Make:  "Toyota",
    Model: "X5 AMG",
    Specs: struct {
        Year  int
        Color string
    }{
        Year:  2020,
        Color: param,
    },
}

param = car.Specs.Color

	if param == "" {
		param = ""
	}

	var bar string
	num := 196
	if (500/42)+num > 200 {
		bar = param
	} else {
		bar = "This should never happen"
	}

	fileName := filepath.Join(os.TempDir(), bar)
	file, err := os.Open(fileName)
	if err != nil {
		http.Error(w, "Couldn't open InputStream on file: "+fileName, http.StatusInternalServerError)
		return
	}
	defer file.Close()

	bContent := make([]byte, 1000)
	size, err := file.Read(bContent)
	if err != nil {
		http.Error(w, "Problem getting InputStream: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	w.Write([]byte("The beginning of file: '" + fileName + "' is:\n\n"))
	w.Write(bContent[:size])
}

func createPoint(x, y string) struct {
    X string
    Y string
} {
    return struct {
        X string
        Y string
    }{
        X: x,
        Y: y,
    }
}


