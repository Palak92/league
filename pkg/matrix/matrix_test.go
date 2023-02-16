package matrix

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestInvert(t *testing.T) {
	var tcs = []struct {
		desc   string
		matrix [][]string
		want   [][]string
	}{
		{
			desc:   "Normal array",
			matrix: [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}},
			want:   [][]string{{"1", "4", "7"}, {"2", "5", "8"}, {"3", "6", "9"}},
		},
		{
			desc:   "Row to column",
			matrix: [][]string{{"1", "2", "3"}},
			want:   [][]string{{"1"}, {"2"}, {"3"}},
		},
		{
			desc:   "Column to row",
			matrix: [][]string{{"1"}, {"2"}, {"3"}},
			want:   [][]string{{"1", "2", "3"}},
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
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("Invert(%v) mismatch (-want +got):\n%s", tc.matrix, diff)
			}
		})
	}
}

func TestFlatten(t *testing.T) {

	var tcs = []struct {
		desc   string
		matrix [][]string
		want   string
	}{
		{
			desc:   "Normal array",
			matrix: [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}},
			want:   "1,2,3,4,5,6,7,8,9",
		},
		{
			desc:   "Row matrix",
			matrix: [][]string{{"1", "-2", "3"}},
			want:   "1,-2,3",
		},
		{
			desc:   "Column matrix",
			matrix: [][]string{{"1"}, {"2"}, {"3"}},
			want:   "1,2,3",
		},
		{
			desc:   "Empty",
			matrix: [][]string{},
			want:   "",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := Flatten(tc.matrix)
			if tc.want != got {
				t.Errorf("Flatten(%s): got %v, want %v", tc.matrix, got, tc.want)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	var tcs = []struct {
		desc   string
		matrix [][]string
		want   string
	}{
		{
			desc:   "Normal array",
			matrix: [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}},
			want:   "362880",
		},
		{
			desc:   "Negative",
			matrix: [][]string{{"1", "-2", "3"}},
			want:   "-6",
		},
		{
			desc:   "Zero",
			matrix: [][]string{{"1", "0", "3"}},
			want:   "0",
		},
		{
			desc:   "Empty",
			matrix: [][]string{},
			want:   "",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := Multiply(tc.matrix)
			if tc.want != got {
				t.Errorf("Multiply(%s): got %v, want %v", tc.matrix, got, tc.want)
			}
		})
	}
}

func TestSum(t *testing.T) {
	var tcs = []struct {
		desc   string
		matrix [][]string
		want   string
	}{
		{
			desc:   "Normal array",
			matrix: [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}},
			want:   "45",
		},
		{
			desc:   "Row with sum zero",
			matrix: [][]string{{"1", "-1"}},
			want:   "0",
		},
		{
			desc:   "Empty",
			matrix: [][]string{},
			want:   "",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := Sum(tc.matrix)
			if tc.want != got {
				t.Errorf("Sum(%v): %v, got: %v", tc.matrix, tc.want, got)
			}
		})
	}
}

func TestString(t *testing.T) {
	var tcs = []struct {
		desc   string
		matrix [][]string
		want   string
	}{
		{
			desc:   "Normal array",
			matrix: [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}},
			want:   "1,2,3\n4,5,6\n7,8,9",
		},
		{
			desc:   "Negative",
			matrix: [][]string{{"1", "-2", "3"}},
			want:   "1,-2,3",
		},
		{
			desc:   "One element",
			matrix: [][]string{{"1"}},
			want:   "1",
		},
		{
			desc:   "Empty",
			matrix: [][]string{},
			want:   "",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := String(tc.matrix)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("String(%v) mismatch (-want +got):\n%s", tc.matrix, diff)
			}
		})
	}
}

func TestIsSquare(t *testing.T) {
	var tcs = []struct {
		desc   string
		matrix [][]string
		want   bool
	}{
		{
			desc:   "Normal array",
			matrix: [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}},
			want:   true,
		},
		{
			desc:   "Negative",
			matrix: [][]string{{"1", "-2", "3"}},
			want:   false,
		},
		{
			desc:   "One element",
			matrix: [][]string{{"1"}},
			want:   true,
		},
		{
			desc:   "Empty",
			matrix: [][]string{},
			want:   true,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := IsSquare(tc.matrix)
			if tc.want != got {
				t.Errorf("IsSquare(%v)= got %v, want %v", tc.matrix, got, tc.want)
			}
		})
	}
}
