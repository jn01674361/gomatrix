package main

import(
    //"fmt"
    gomatrix "gomatrix"
)

func main() {
    
    A := gomatrix.Matrix{
        Matrix: [][]float64{
            {1.0, 2.0},
            {3.0, 4.0},
        },
    }
    B := gomatrix.Matrix{
        Matrix: [][]float64{
            {5.0, 6.0, 7.0},
            {8.0, 9.0, 10.0},
        },
    }

    C := gomatrix.MatMul(A,B)

    gomatrix.PrintMat(A)
    gomatrix.PrintMat(B)
    gomatrix.PrintMat(C)

}