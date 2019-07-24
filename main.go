package main

import(
    // "fmt"
    gomatrix "gomatrix"
    "github.com/gorilla/mux"
    "log"
    "net/http"
    "encoding/json"
)
type matrixPair struct{
    A [][]float64
    B [][]float64
}
type product struct{
    result [][]float64
}
func main() {
    router := mux.NewRouter()
    router.HandleFunc("/matmul/{a}/{b}", matmul).Methods("GET")
    log.Fatal(http.ListenAndServe(":8000",router))
}
func matmul(w http.ResponseWriter, r *http.Request){
    params := mux.Vars(r)
    byteA := []byte(params["a"])
    byteB := []byte(params["b"])
    var mats matrixPair
    err1 := json.Unmarshal(byteA, &mats.A)
    if err1 != nil{
        log.Println(err1)
    }
    err2 := json.Unmarshal(byteB, &mats.B)
    if err2!= nil{
        log.Println(err2)
    }
    matA := gomatrix.ToMatrix(mats.A)
    matB := gomatrix.ToMatrix(mats.B)
    matC := gomatrix.MatMul(matA, matB)
    json.NewEncoder(w).Encode(matC.Matrix)
}