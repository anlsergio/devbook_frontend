package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
)

func main() {
	fmt.Println("Running Webapp...")

	r := router.Create()
	log.Fatal(http.ListenAndServe(":3000", r))
}