package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
)

var visitCount uint64

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// 打印出当前请求的路径
	log.Printf("Received request for path: %s", r.URL.Path)
	count := atomic.AddUint64(&visitCount, 1)
	log.Printf("Received request from %s, visit count: %d", r.RemoteAddr, count)
	fmt.Fprintf(w, "hello go world!\nvisit count: %d\n", count)
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request for path: %s", r.URL.Path)
		fmt.Fprintf(w, "Welcome to the home page!\n")
	})
	log.Println("Starting server on :8110")
	if err := http.ListenAndServe(":8110", nil); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
