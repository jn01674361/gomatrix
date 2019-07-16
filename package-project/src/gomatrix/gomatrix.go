package goMatrix

type Matrix struct{
	Matrix [][]float64
} 

func NewMatrix(m,n int) Matrix {

	mat := make([][]float64, m)

	for row:=0; row<m; row++{
		mat[row] = make([]float64, n)
	}

	ret := Matrix{
		Matrix: mat,
	}
	return ret
}