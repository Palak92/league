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
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/invert", InvertHandler)
	http.HandleFunc("/flatten", FlattenHandler)
	http.HandleFunc("/sum", SumHandler)
	http.HandleFunc("/multiply", MultiplyHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Echo operation
func EchoHandler(w http.ResponseWriter, r *http.Request) {
	records, err := csvRecords(r)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	response := matrix.MatrixToString(records)
	fmt.Fprint(w, response)
}

// Invert operation
func InvertHandler(w http.ResponseWriter, r *http.Request) {
	records, err := csvRecords(r)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	inv := matrix.Invert(records)
	response := matrix.MatrixToString(inv)
	fmt.Fprint(w, response)
}

// Flatten operation
func FlattenHandler(w http.ResponseWriter, r *http.Request) {
	records, err := csvRecords(r)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	flat := matrix.FlattenMatrix(records)
	fmt.Fprint(w, flat)
}

// Sum operation
func SumHandler(w http.ResponseWriter, r *http.Request) {
	records, err := csvRecords(r)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	sum := matrix.SumMatrix(records)
	fmt.Fprint(w, sum)
}

// Multiply operation
func MultiplyHandler(w http.ResponseWriter, r *http.Request) {
	records, err := csvRecords(r)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	product := matrix.MultiplyMatrix(records)
	fmt.Fprint(w, product)
}

// csvRecords reads a CSV file from an io.Reader and returns a 2D slice of strings
func csvRecords(req *http.Request) ([][]string, error) {
	file, _, err := req.FormFile("file")
	if err != nil {
		return nil, fmt.Errorf("while getting file from req")
	}
	defer file.Close()
	return csv.NewReader(file).ReadAll()
}

func echo(w http.ResponseWriter, r *http.Request) {
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
