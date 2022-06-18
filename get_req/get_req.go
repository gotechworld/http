package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)

func main() {

    resp, err := http.Get("https://cms-s1.altex.ro/sellers")

    if err != nil {
        log.Fatal(err)
    }

    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(body))
}
