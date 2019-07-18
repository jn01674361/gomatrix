package main

import(
    // "fmt"
    gomatrix "gomatrix"
    //"github.com/gorilla/mux"
    //"log"
    //"net/http"
    //"encoding/json"
)

func main() {
    //router := mux.NewRouter()
    test();
}

func test() {
    
    E := [][]float64{
        {1.0, 2.0},
        {3.0, 4.0},
    }

    F := gomatrix.ToMatrix(E)
    
    gomatrix.PrintMat(F)

}