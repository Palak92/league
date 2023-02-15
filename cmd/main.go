package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"matrix"
	"net/http"
	"strings"
)

// Run with
//		go run .
// Send request with:
//		curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"

func main() {
	http.HandleFunc("/echo", EchoHandler)
	http.HandleFunc("/invert", InvertHandler)
	http.HandleFunc("/flatten", FlattenHandler)
	http.HandleFunc("/sum", SumHandler)
	http.HandleFunc("/multiply", MultiplyHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Echo operation
func EchoHandler(w http.ResponseWriter, r *http.Request) {
	m := parseCSV(r.Body)
	response := matrix.MatrixToString(m)
	fmt.Fprint(w, response)
}

// Invert operation
func InvertHandler(w http.ResponseWriter, r *http.Request) {
	m := parseCSV(r.Body)
	inv := matrix.InvertMatrix(m)
	response := matrix.matrixToString(inv)
	fmt.Fprint(w, response)
}

// Flatten operation
func FlattenHandler(w http.ResponseWriter, r *http.Request) {
	m := parseCSV(r.Body)
	flat := matrix.FlattenMatrix(m)
	fmt.Fprint(w, flat)
}

// Sum operation
func SumHandler(w http.ResponseWriter, r *http.Request) {
	m := parseCSV(r.Body)
	sum := matrix.SumMatrix(m)
	fmt.Fprint(w, sum)
}

// Multiply operation
func MultiplyHandler(w http.ResponseWriter, r *http.Request) {
	m := parseCSV(r.Body)
	product := matrix.MultiplyMatrix(m)
	fmt.Fprint(w, product)
}

// parseCSV reads a CSV file from an io.Reader and returns a 2D slice of strings
func parseCSV(reader io.Reader) [][]string {
	r := csv.NewReader(reader)
	var m [][]string
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		m = append(m, record)
	}
	return m
}

func echo(w http.ResponseWriter, req *http.Request) {
	file, _, err := req.FormFile("file")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
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
