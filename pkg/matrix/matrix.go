// Pacakge matrix has methods for modify a matrix.
package matrix

import (
	"log"
	"strconv"
)

// MatrixToString returns a string in matrix format for a given 2D slice of strings
func MatrixToString(m [][]string) string {
	response := ""
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			response += m[i][j]
			if j < len(m[0])-1 {
				response += ","
			}
		}
		if i < len(m)-1 {
			response += "\n"
		}
	}
	return response
}

// Invert returns a 2D slice of strings with the columns and rows inverted.
func Invert(m [][]string) [][]string {
	rows, cols := size(m)

	inv := make([][]string, cols)
	for j := 0; j < cols; j++ {
		irow := make([]string, rows)
		inv[j] = irow
		for i := 0; i < rows; i++ {
			inv[j][i] = m[i][j]
		}
	}
	return inv
}

// FlattenMatrix returns a string with the matrix flattened into 1 line
func FlattenMatrix(m [][]string) string {
	flat := ""
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			flat += m[i][j] + ","
		}
	}
	flat = flat[:len(flat)-1] // remove last comma
	return flat
}

// SumMatrix returns the sum of all integers in the matrix
func SumMatrix(m [][]string) int {
	sum := 0
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			val, err := strconv.Atoi(m[i][j])
			if err != nil {
				log.Fatal(err)
			}
			sum += val
		}
	}
	return sum
}

// MultiplyMatrix returns the product of all integers in the matrix
func MultiplyMatrix(m [][]string) int {
	product := 1
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			val, err := strconv.Atoi(m[i][j])
			if err != nil {
				log.Fatal(err)
			}
			product *= val
		}
	}
	return product
}

func size(m [][]string) (rows, cols int) {
	rows = len(m)
	if rows == 0 {
		return
	}
	cols = len(m[0])
	return
}
