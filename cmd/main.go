package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/palak92/league/pkg/matrix"
)

// Run with
//		go run .
// Send request with:
//		curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"

func main() {
	http.HandleFunc("/echo", echoHandler)
	http.HandleFunc("/invert", invertHandler)
	http.HandleFunc("/flatten", flattenHandler)
	http.HandleFunc("/sum", sumHandler)
	http.HandleFunc("/multiply", multiplyHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	records, err := csvRecords(r)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	var response string
	for _, row := range records {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}
	fmt.Fprint(w, response)
}

// invertHandler inverts matrix.
func invertHandler(w http.ResponseWriter, r *http.Request) {
	records, err := csvRecords(r)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	inv := matrix.Invert(records)
	response := matrix.String(inv)
	fmt.Fprint(w, response)
}

// flattenHandler performs flats the matrix.
func flattenHandler(w http.ResponseWriter, r *http.Request) {
	records, err := csvRecords(r)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	flat := matrix.Flatten(records)
	fmt.Fprint(w, flat)
}

// Sum operation performs sum operation.
func sumHandler(w http.ResponseWriter, r *http.Request) {
	records, err := csvRecords(r)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	sum := matrix.Sum(records)
	fmt.Fprint(w, sum)
}

// multiplyHandler performs multiply operation.
func multiplyHandler(w http.ResponseWriter, r *http.Request) {
	records, err := csvRecords(r)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	product := matrix.Multiply(records)
	fmt.Fprint(w, product)
}

// csvRecords reads a CSV file from an io.Reader and returns a 2D slice of strings
func csvRecords(req *http.Request) ([][]string, error) {
	file, _, err := req.FormFile("file")
	if err != nil {
		return nil, fmt.Errorf("while getting file from request :%v", err)
	}
	defer file.Close()
	return csv.NewReader(file).ReadAll()
}
