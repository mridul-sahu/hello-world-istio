package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		hostname := os.Getenv("HOSTNAME")

		log.Printf("Hostname: %s ", hostname)

		io.WriteString(w, "Hello, world! From " + hostname)
	}

	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
