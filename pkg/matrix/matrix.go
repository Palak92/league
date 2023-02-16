// Pacakge matrix has methods for modify a matrix.
package matrix

import (
	"fmt"
	"log"
	"strconv"
)

// ContainsAllIntegerElements returns true if  matrix contains all int elements.
func ContainsAllIntegerElements(m [][]string) (bool, error) {
	rows, cols := size(m)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			_, err := strconv.Atoi(m[i][j])
			if err != nil {
				return false, fmt.Errorf("error while converting element %q at row %d and column %d to integer: %w", m[i][j], i, j, err)
			}
		}
	}
	return true, nil
}

// IsSquare returns true if  matrix is square matrix.
func IsSquare(m [][]string) bool {
	rows, cols := size(m)
	return rows == cols
}

// String returns a string in matrix format for a given 2D slice of strings.
func String(m [][]string) string {
	rows, cols := size(m)
	response := ""
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			response += m[i][j]
			if j < cols-1 {
				response += ","
			}
		}
		if i < rows-1 {
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

// Flatten returns a string with the matrix flattened into 1 line.
func Flatten(m [][]string) string {
	flat := ""
	rows, cols := size(m)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			flat += m[i][j] + ","
		}
	}
	if flat != "" {
		flat = flat[:len(flat)-1] // remove last comma
	}
	return flat
}

// Sum returns the sum of all integers in the matrix.
func Sum(m [][]string) string {
	sum := 0
	rows, cols := size(m)
	if rows == 0 && cols == 0 {
		return ""
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			val, err := strconv.Atoi(m[i][j])
			if err != nil {
				log.Fatal(err)
			}
			sum += val
		}
	}
	return strconv.Itoa(sum)
}

// Multiply returns the product of all integers in the matrix.
func Multiply(m [][]string) string {
	product := 1
	rows, cols := size(m)
	if rows == 0 && cols == 0 {
		return ""
	}
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			val, err := strconv.Atoi(m[i][j])
			if err != nil {
				log.Fatal(err)
			}
			product *= val
		}
	}
	return strconv.Itoa(product)
}

func size(m [][]string) (rows, cols int) {
	rows = len(m)
	if rows == 0 {
		return
	}
	cols = len(m[0])
	return
}
