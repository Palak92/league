package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const validFilePath = "../test/data/matrix.csv"

func TestEchoHandler(t *testing.T) {
	body, ct := multipartBody(t, validFilePath)
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
	body, ct := multipartBody(t, validFilePath)
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
	body, ct := multipartBody(t, validFilePath)
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
	body, ct := multipartBody(t, validFilePath)
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

	body, ct := multipartBody(t, validFilePath)
	req := httptest.NewRequest(http.MethodPost, "/invert", body)
	req.Header.Add("Content-Type", ct)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(invertHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("invertHandler returned wrong status code: got %v, want %v", rr.Code, http.StatusOK)
	}

	expected := "1,4,7\n2,5,8\n3,6,9"
	if rr.Body.String() != expected {
		t.Errorf("invertHandler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestMatrixValidation(t *testing.T) {

	var tcs = []struct {
		desc     string
		filePath string
		wantErr  bool
	}{
		{
			desc:     "valid matrix",
			filePath: validFilePath,
			wantErr:  false,
		},
		{
			desc:     "non-square matrix",
			filePath: "../test/data/non_square.csv",
			wantErr:  true,
		},
		{
			desc:     "non-integer matrix elements",
			filePath: "../test/data/non_integer.csv",
			wantErr:  true,
		},
		{
			desc:     "Empty",
			filePath: "../test/data/empty.csv",
			wantErr:  false,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			body, ct := multipartBody(t, tc.filePath)
			req := httptest.NewRequest(http.MethodPost, "/sum", body)
			req.Header.Add("Content-Type", ct)

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(sumHandler)
			handler.ServeHTTP(rr, req)

			if rr.Code != http.StatusOK {
				t.Errorf("sumHandler returned wrong status code: got %v, want %v", rr.Code, http.StatusOK)
			}

			got := rr.Body.String()
			if strings.Contains(got, "error") && !tc.wantErr {
				t.Errorf("Output of file at path %q should not contain err", tc.filePath)
			}
		})
	}

}

func multipartBody(t *testing.T, p string) (io.Reader, string) {
	fieldName := "file"
	body := new(bytes.Buffer)
	absPath, _ := filepath.Abs(p)
	mw := multipart.NewWriter(body)

	file, err := os.Open(absPath)
	if err != nil {
		t.Fatal(err)
	}

	w, err := mw.CreateFormFile(fieldName, absPath)
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
