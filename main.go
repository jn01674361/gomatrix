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
    
    // jsonInput := []byte(`[{"Elements": [1.0,2.0]},{"Elements": [3.0, 4.0]}]`)

    jsonArr := []byte(`[[1.0,2.0],[3.0,4.0]]`)

    var jArr [][]float64

    err := json.Unmarshal(jsonArr, &jArr)
    if err != nil{
        log.Println(err)
    }
    fmt.Println(jArr)

}