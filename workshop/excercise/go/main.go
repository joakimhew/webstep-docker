package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, time.Now().Format("15:04:05"))
	fmt.Fprintf(w, "Hello, from Webstep presentation Go server! 🎉")
	//fmt.Fprintf(w, "The time is currently: %v", time.Now().Format("15:04:05"))
}

func main() {
	portPtr := flag.Int("port", 8080, "port to run webserver on")
	flag.Parse()

	http.HandleFunc("/", handler)

	log.Printf("Listening on :%v", *portPtr)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", *portPtr), nil))
}
