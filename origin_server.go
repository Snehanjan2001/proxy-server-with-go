package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	originServerhandler := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		fmt.Printf("[origin server ] recieved request at %s\n", time.Now())
		_, _ = fmt.Fprint(rw, "origin server response")

	})
	fmt.Println("Server Started")
	log.Fatal(http.ListenAndServe(":8081", originServerhandler))
	fmt.Println("Server running")

}
