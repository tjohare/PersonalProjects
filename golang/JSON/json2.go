package main

import (
    "encoding/json"
    "fmt"
    "log"
)

var a = `{  
"sometext":[  
  {  
     "sha":"1234567",
     "message":"hello world",
     "author":"tjohare",
     "timestamp":1479445228
  },
  {  
     "sha":"00000000",
     "message":"hello TJ",
     "author":"tjohare",
     "timestamp":20180620
  }
]
}`

type Root struct {
    Text []*Object `json:"sometext"`
}

type Object struct {
    Sha       string `json:"sha"`
    Message   string `json:"message"`
    Author    string `json:"author"`
    Timestamp int    `json:"timestamp"`
}

func main() {
    var j Root
    err := json.Unmarshal([]byte(a), &j)
    if err != nil {
        log.Fatalf("error parsing JSON: %s\n", err.Error())
    }
    fmt.Printf("%+v\n", j.Text[0].Sha)
    fmt.Printf("%+v\n", j.Text[0].Message)
    fmt.Printf("%+v\n", j.Text[0].Author)
    fmt.Printf("%+v\n", j.Text[0].Timestamp)

    fmt.Printf("%+v\n", j.Text[1].Sha)
    fmt.Printf("%+v\n", j.Text[1].Message)
    fmt.Printf("%+v\n", j.Text[1].Author)
    fmt.Printf("%+v\n", j.Text[1].Timestamp)
}