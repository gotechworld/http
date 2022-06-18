package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	resp, err := http.Get("https://cms-s1.altex.ro/sellers")

	if err != nil {

		log.Fatal(err)
	}

	fmt.Println(resp.Status)
	fmt.Println(resp.StatusCode)
}
