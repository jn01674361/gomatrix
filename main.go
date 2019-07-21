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
// func fromJsonArr(jsonData []byte) interface{
//     var arr interface
//     err := json.Unmarshal(jsonData, &arr)
//     if err != nil {
//         log.Println(err)
//     }
//     return arr
// }
func main() {
    router := mux.NewRouter()
    router.HandleFunc("/matmul", matmul).Methods("GET")
    log.Fatal(http.ListenAndServe(":8000",router))
}
func matmul(w http.ResponseWriter, r *http.Request){
    
    params := mux.Vars(r)

    log.Println("params")
    log.Println(params)

    // params will be what holds the matrix!
    // matmul/{matrix} will need to be routed!

    var mats matrixPair
    _ = json.NewDecoder(r.Body).Decode(&mats)

    
    // err := json.Unmarshal(jsonInput, &mats)
    // if err != nil{
    //     log.Println(err)
    // }

    matA := gomatrix.ToMatrix(mats.A)
    matB := gomatrix.ToMatrix(mats.B)
    matC := gomatrix.MatMul(matA, matB)
    
    json.NewEncoder(w).Encode(matC.Matrix)
    return 
    // gomatrix.PrintMat(matA)
    // gomatrix.PrintMat(matB)
    // gomatrix.PrintMat(gomatrix.MatMul(matA,matB))

    
}
func test() {
    

    jsonInput := []byte(`{"A":[[1.0,2.0],[3.0,4.0]], "B":[[1.0,2.0],[3.0,4.0]]}`)

    var mats matrixPair

    err := json.Unmarshal(jsonInput, &mats)
    if err != nil{
        log.Println(err)
    }

    matA := gomatrix.ToMatrix(mats.A)
    matB := gomatrix.ToMatrix(mats.B)

    gomatrix.PrintMat(matA)
    gomatrix.PrintMat(matB)
    gomatrix.PrintMat(gomatrix.MatMul(matA,matB))

    //jsonOutput
}