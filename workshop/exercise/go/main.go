package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Webstep!")
}

func main() {
	portPtr := flag.Int("port", 8080, "port to run webserver on")
	flag.Parse()

	http.HandleFunc("/", handler)

	log.Printf("Listening on :%v", *portPtr)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", *portPtr), nil))
}
