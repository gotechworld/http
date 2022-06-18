package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {

    resp, err := http.Head("https://cms-s1.altex.ro/sellers")

    if err != nil {

        log.Fatal(err)
    }

    for k, v := range resp.Header {

        fmt.Printf("%s %s\n", k, v)
    }
}