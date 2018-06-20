//https://play.golang.org./p/7CtALgsjK3

package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
)

type Database struct {
    Address  string
    Port     int
    User     string
    Password string
    Database string
}
type Mysql struct {
    Database
}
type Postgres struct {
    Database
}

type Configuration struct {
    Mysql    Mysql
    Postgres Postgres
}

func main(){
    content, err := ioutil.ReadFile("./config.json")
    if err!=nil{
        fmt.Print("Error:",err)
    }

    var conf Configuration
    err=json.Unmarshal(content, &conf)
    if err!=nil{
        log.Print("Error:",err)
    }
    
    fmt.Println(conf)
    fmt.Println(conf.Mysql)
    fmt.Println(conf.Postgres)
    log.Println(conf.Postgres.Address)
}