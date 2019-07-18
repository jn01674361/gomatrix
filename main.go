package main

import(
    "fmt"
    // gomatrix "gomatrix"
    //"github.com/gorilla/mux"
    "log"
    //"net/http"
    "encoding/json"
)

type jsonRow struct {
    Elements []float64
}

type jsonMatrix struct {
    Rows []jsonRow
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
    //router := mux.NewRouter()

    test();
}

func test() {
    
    jsonInput := []byte(`[{"Elements": [1.0,2.0]},{"Elements": [3.0, 4.0]}]`)

    var jMat jsonMatrix

    err := json.Unmarshal(jsonInput, &jMat)

    if err != nil{
        log.Println(err)
    }
    fmt.Println(jMat.Rows)
}