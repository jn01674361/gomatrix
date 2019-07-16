package goMatrix

type Matrix struct{
	Matrix [][]float64
} 

func Initialize(m,n int) Matrix {

	mat := make([][]float64, m)

	for row:=0; row<m; row++{
		mat[row] = make([]float64, n)
	}

	ret := Matrix{
		Matrix: mat,
	}
	return ret
}

// func ones(m,n int) Matrix {
// 	return Matrix{rows:1,cols:1, Matrix:{{1}}}
// }