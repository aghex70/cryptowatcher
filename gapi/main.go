package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"gapi-agp/konfig"
)

func main() {
	konfig.LoadConfig(konfig.CONFIG_PATH)
	// Hello world, the web server
	fmt.Println("Hello, universe!!!!!")
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
