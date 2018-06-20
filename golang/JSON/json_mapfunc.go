package main

import (
    "encoding/json"
    "fmt"
)

func main() {
    m := make(map[string]interface{})
    m["appVersion"] = "0.0.1"
    m["uploadStatus"] = "success"

    strJson, _ := json.Marshal(m)
    str := string(strJson)
    fmt.Println("---ENCODING--")
    fmt.Println(str)

    //put it back
    fmt.Println("---DECODING--")
    n := make(map[string]interface{})
    json.Unmarshal([]byte(str), &n)
    for key, val := range n {
        fmt.Println(key, ":", val)
    }
}