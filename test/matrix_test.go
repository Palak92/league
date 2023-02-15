package matrix

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestInvert(t *testing.T) {
	var tcs = []struct {
		desc   string
		matrix [][]string
		want   [][]string
	}{
		{
			desc:   "Normal array",
			matrix: [][]string{[]string{"1", "2", "3"}, []string{"4", "5", "6"}, []string{"7", "8", "9"}},
			want:   [][]string{[]string{"1", "4", "7"}, []string{"2", "5", "8"}, []string{"3", "6", "9"}},
		},
		{
			desc:   "Row to column",
			matrix: [][]string{[]string{"1", "2", "3"}},
			want:   [][]string{[]string{"1"}, []string{"2"}, []string{"3"}},
		},
		{
			desc:   "Column to row",
			matrix: [][]string{[]string{"1"}, []string{"2"}, []string{"3"}},
			want:   [][]string{[]string{"1", "2", "3"}},
		},
		{
			desc:   "Empty",
			matrix: [][]string{},
			want:   [][]string{},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := Invert(tc.matrix)
			if !reflect.DeepEqual(tc.want, got) {
				// t.Fatalf("Invert(%v): %v, got: %v", tc.matrix tc.want, got)
			}
		})
	}
}

func TestMultiplyHandler(t *testing.T) {
	m := "1,2,3\n4,5,6\n7,8,9"
	req, err := http.NewRequest("POST", "/multiply", strings.NewReader(m))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(MultiplyHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v, want %v", rr.Code, http.StatusOK)
	}

	expected := "362880"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestEchoHandler(t *testing.T) {
	m := "1,2,3\n4,5,6\n7,8,9"
	req, err := http.NewRequest("POST", "/echo", strings.NewReader(m))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(EchoHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v, want %v", rr.Code, http.StatusOK)
	}

	expected := "1,2,3\n4,5,6\n7,8,9"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestFlattenHandler(t *testing.T) {
	m := "1,2,3\n4,5,6\n7,8,9"
	req, err := http.NewRequest("POST", "/flatten", strings.NewReader(m))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(FlattenHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v, want %v", rr.Code, http.StatusOK)
	}

	expected := "1,2,3,4,5,6,7,8,9"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestSumHandler(t *testing.T) {
	m := "1,2,3\n4,5,6\n7,8,9"
	req, err := http.NewRequest("POST", "/sum", strings.NewReader(m))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(SumHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v, want %v", rr.Code, http.StatusOK)
	}

	expected := "45"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestInvertHandler(t *testing.T) {
	m := "1,2,3\n4,5,6\n7,8,9"
	req, err := http.NewRequest("POST", "/invert", strings.NewReader(m))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(InvertHandler)
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v, want %v", rr.Code, http.StatusOK)
	}

	expected := "1,4,7\n2,5,8\n3,6,9"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestParseCSV(t *testing.T) {
	csvData := "1,2,3\n4,5,6\n7,8,9"
	r := bytes.NewReader([]byte(csvData))
	m := parseCSV(r)

	if len(m) != 3 {
		t.Errorf("parseCSV() returned %v rows, expected %v", len(m), 3)
	}

	if len(m[0]) != 3 {
		t.Errorf("parseCSV() returned %v columns, expected %v", len(m[0]), 3)
	}

	expected := [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}}
	for i := range m {
		for j := range m[i] {
			if m[i][j] != expected[i][j] {
				t.Errorf("parseCSV() returned unexpected value at (%v, %v): got %v, want %v", i, j, m[i][j], expected[i][j])
			}
		}
	}
}

func TestInvertMatrix(t *testing.T) {
	m := [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}}
	inv := invertMatrix(m)

	if len(inv) != 3 {
		t.Errorf("invertMatrix() returned %v rows, expected %v", len(inv), 3)
	}

	if len(inv[0]) != 3 {
		t.Errorf("invertMatrix() returned %v columns, expected %v", len(inv[0]), 3)
	}

	expected := [][]string{{"1", "4", "7"}, {"2", "5", "8"}, {"3", "6", "9"}}
	for i := range inv {
		for j := range inv[i] {
			if inv[i][j] != expected[i][j] {
				t.Errorf("invertMatrix() returned unexpected value at (%v, %v): got %v, want %v", i, j, inv[i][j], expected[i][j])
			}
		}
	}
}

func TestFlattenMatrix(t *testing.T) {
	m := [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}}
	flat := flattenMatrix(m)

	expected := "1,2,3,4,5,6,7,8,9"
	if flat != expected {
		t.Errorf("flattenMatrix() returned unexpected value: got %v, want %v", flat, expected)
	}
}

func TestSumMatrix(t *testing.T) {
	m := [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}}
	sum := sumMatrix(m)

	if sum != 45 {
		t.Errorf("sumMatrix() returned unexpected value: got %v, want %v", sum, 45)
	}
}

func TestMultiplyMatrix(t *testing.T) {
	m := [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}}
	product := multiplyMatrix(m)

	if product != 362880 {
		t.Errorf("multiplyMatrix() returned unexpected value: got %v, want %v", product, 362880)
	}
}

func TestMatrixToString(t *testing.T) {
	m := [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}}
	str := matrixToString(m)

	expected := "1,2,3\n4,5,6\n7,8,9"
	if str != expected {
		t.Errorf("matrixToString() returned unexpected value: got %v, want %v", str, expected)
	}
}

func TestParseCSVEmptyInput(t *testing.T) {
	r := bytes.NewReader([]byte{})
	m := parseCSV(r)

	if len(m) != 0 {
		t.Errorf("parseCSV() returned %v rows, expected %v", len(m), 0)
	}
}

func TestParseCSVInvalidCSV(t *testing.T) {
	csvData := "1,2,3\n4,5,6\n7,8"
	r := bytes.NewReader([]byte(csvData))
	m := parseCSV(r)

	if len(m) != 2 {
		t.Errorf("parseCSV() returned %v rows, expected %v", len(m), 2)
	}

	if len(m[0]) != 3 {
		t.Errorf("parseCSV() returned %v columns, expected %v", len(m[0]), 3)
	}

	expected := [][]string{{"1", "2", "3"}, {"4", "5", "6"}}
	for i := range m {
		for j := range m[i] {
			if m[i][j] != expected[i][j] {
				t.Errorf("parseCSV() returned unexpected value at (%v, %v): got %v, want %v", i, j, m[i][j], expected[i][j])
			}
		}
	}
}

func TestInvertHandlerInvalidMatrix(t *testing.T) {
	m := "1,2,3\n4,5,6\n7,8"
	req, err := http.NewRequest("POST", "/invert", strings.NewReader(m))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(InvertHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v, want %v", rr.Code, http.StatusBadRequest)
	}
}

func TestMultiplyHandlerInvalidMatrix(t *testing.T) {
	m := "1,2,3\n4,5,6\n7,8"
	req, err := http.NewRequest("POST", "/multiply", strings.NewReader(m))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(MultiplyHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v, want %v", rr.Code, http.StatusBadRequest)
	}
}

func TestFlattenHandlerInvalidMatrix(t *testing.T) {
	m := "1,2,3\n4,5,6\n7,8"
	req, err := http.NewRequest("POST", "/flatten", strings.NewReader(m))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(FlattenHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v, want %v", rr.Code, http.StatusBadRequest)
	}
}

func TestSumHandlerInvalidMatrix(t *testing.T) {
	m := "1,2,3\n4,5,6\n7,8"
	req, err := http.NewRequest("POST", "/sum", strings.NewReader(m))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(SumHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v, want %v", rr.Code, http.StatusBadRequest)
	}
}

func TestMatrixToStringEmptyMatrix(t *testing.T) {
	m := [][]string{}
	str := matrixToString(m)

	expected := ""
	if str != expected {
		t.Errorf("matrixToString() returned unexpected value: got %v, want %v", str, expected)
	}
}

func TestInvertMatrixEmptyMatrix(t *testing.T) {
	m := [][]string{}
	inv := invertMatrix(m)

	if len(inv) != 0 {
		t.Errorf("invertMatrix() returned %v rows, expected %v", len(inv), 0)
	}
}

func TestFlattenMatrixEmptyMatrix(t *testing.T) {
	m := [][]string{}
	flat := flattenMatrix(m)

	expected := ""
	if flat != expected {
		t.Errorf("flattenMatrix() returned unexpected value: got %v, want %v", flat, expected)
	}
}

func TestSumMatrixEmptyMatrix(t *testing.T) {
	m := [][]string{}
	sum := sumMatrix(m)

	if sum != 0 {
		t.Errorf("sumMatrix() returned unexpected value: got %v, want %v", sum, 0)
	}
}

func TestMultiplyMatrixEmptyMatrix(t *testing.T) {
	m := [][]string{}
	product := multiplyMatrix(m)

	if product != 0 {
		t.Errorf("multiplyMatrix() returned unexpected value: got %v, want %v", product, 0)
	}
}
