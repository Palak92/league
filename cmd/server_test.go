package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestEchoHandler(t *testing.T) {
	body, ct := multipartBody(t)
	req := httptest.NewRequest(http.MethodPost, "/echo", body)
	req.Header.Add("Content-Type", ct)

	rr := httptest.NewRecorder()
	echoHandler(rr, req)
	res := rr.Result()
	defer res.Body.Close()

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v, want %v", rr.Code, http.StatusOK)
	}
	want := "1,2,3\n4,5,6\n7,8,9\n"
	got, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if diff := cmp.Diff(want, string(got)); diff != "" {
		t.Errorf("echoHandler() diff (-want +got):\n%s", diff)
	}
}

func TestMultiplyHandler(t *testing.T) {
	body, ct := multipartBody(t)
	req := httptest.NewRequest(http.MethodPost, "/echo", body)
	req.Header.Add("Content-Type", ct)

	rr := httptest.NewRecorder()
	multiplyHandler(rr, req)
	res := rr.Result()
	defer res.Body.Close()

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v, want %v", rr.Code, http.StatusOK)
	}
	want := "362880"
	got, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if string(got) != want {
		t.Errorf("echoHandler(): got %q, want %q", string(got), want)
	}
}

func TestFlattenHandler(t *testing.T) {
	body, ct := multipartBody(t)
	req := httptest.NewRequest(http.MethodPost, "/flatten", body)
	req.Header.Add("Content-Type", ct)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(flattenHandler)
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
	body, ct := multipartBody(t)
	req := httptest.NewRequest(http.MethodPost, "/sum", body)
	req.Header.Add("Content-Type", ct)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(sumHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("sumHandler returned wrong status code: got %v, want %v", rr.Code, http.StatusOK)
	}

	expected := "45"
	if rr.Body.String() != expected {
		t.Errorf("sumHandler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestInvertHandler(t *testing.T) {

	body, ct := multipartBody(t)
	req := httptest.NewRequest(http.MethodPost, "/invert", body)
	req.Header.Add("Content-Type", ct)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(invertHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("invertHandler returned wrong status code: got %v, want %v", rr.Code, http.StatusOK)
	}

	expected := "45"
	if rr.Body.String() != expected {
		t.Errorf("invertHandler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}

	m := "1,4,7\n2,5,8\n3,6,9"
	req, err := http.NewRequest("POST", "/invert", strings.NewReader(m))
	if err != nil {
		t.Fatal(err)
	}
}

// func TestMatrixToString(t *testing.T) {
// 	m := [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}}
// 	str := MatrixToString(m)

// 	expected := "1,2,3\n4,5,6\n7,8,9"
// 	if str != expected {
// 		t.Errorf("matrixToString() returned unexpected value: got %v, want %v", str, expected)
// 	}
// }

// func TestParseCSVEmptyInput(t *testing.T) {
// 	r := bytes.NewReader([]byte{})
// 	m := parseCSV(r)

// 	if len(m) != 0 {
// 		t.Errorf("parseCSV() returned %v rows, expected %v", len(m), 0)
// 	}
// }

func multipartBody(t *testing.T) (io.Reader, string) {
	filePath := "../test/test_data/matrix.csv"
	fieldName := "file"
	body := new(bytes.Buffer)

	mw := multipart.NewWriter(body)

	file, err := os.Open(filePath)
	if err != nil {
		t.Fatal(err)
	}

	w, err := mw.CreateFormFile(fieldName, filePath)
	if err != nil {
		t.Fatal(err)
	}

	if _, err := io.Copy(w, file); err != nil {
		t.Fatal(err)
	}
	// close the writer before making the request
	mw.Close()
	return body, mw.FormDataContentType()
}
