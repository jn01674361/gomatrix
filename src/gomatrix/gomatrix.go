package gomatrix

import "fmt"

type Matrix struct{
	Matrix [][]float64
} 
// returns the number of rows of a matrix
// in: the matrix
/// out: number of rows
func Rows(A Matrix) int {
	return len(A.Matrix)
}
// returns the number of columns of a matrix
// in: the matrix
/// out: number of columns
func Cols(A Matrix) int {
	var cols int
	if len(A.Matrix) > 0 {
		cols = len(A.Matrix[0])
	} else {
		cols = 0
	}
	return cols
}
// returns a matrix with m rows and n columns, filled with zeros
// in: number of rows, columns
/// out: matrix object
func Zeros(m,n int) Matrix {
	mat := make([][]float64, m)
	for row:=0; row<m; row++{
		mat[row] = make([]float64, n)
	}
	ret := Matrix{
		Matrix: mat,
	}
	return ret
}
// returns a matrix object from a 2d-float array.
// in: 2d float64 array
// out: matrix object with same element as input array
func ToMatrix(A [][]float64) Matrix {
	var n int
	m := len(A);
	if m < 1 {
		return Zeros(m,n)
	} 	
	n = len(A[1])
	mat := make([][]float64, m)
	for i := 0; i < m; i++ {
		row := make([]float64, n)

		for j:=0; j < n; j++{
			row[j] = A[i][j] 
		}
		mat[i] = row
	}
	return Matrix{
		Matrix: mat,
	}
}
// returns the matrix product of two matrices
// in: two matrices 
// out: the matrix product of the matrices as a matrix object, 
// 		if the matrices can be multiplied in the supplied order.
func MatMul(A,B Matrix) Matrix {
	var m, n1, n2, p int
	m = Rows(A)
	n1 = Cols(A)
	n2 = Rows(B)
	p = Cols(B)
	if n1 != n2 {
		return Zeros(0,0)
	} else if !(m > 0 && n1 > 0 && n2 > 0 && p > 0) {
		return Zeros(0,0)
	} else{
		C := Zeros(m,p)
		for i := 0; i < m; i++ {
			for j := 0; j < p; j++ {
				for k := 0; k < n1; k++ {
					C.Matrix[i][j] += A.Matrix[i][k] * B.Matrix[k][j]
				} 
			}
		}
		return C
	}
}
// prints the values on each row of the matrix on a separate line
// in: the matrix
func PrintMat(A Matrix) {
	m := Rows(A)
	n := Cols(A)

	fmt.Println()
	for i:= 0; i < m; i++{
		for j:=0; j < n; j++{
			fmt.Printf("%g ", A.Matrix[i][j])
		}
		fmt.Print("\n")
	}
	fmt.Println()

}
