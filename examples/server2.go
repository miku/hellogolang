// server2 is a RNGAAS: random number generator as a service.
package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
)

// mu protects the count.
var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	if err := http.ListenAndServe("localhost:8000", nil); err != nil {
		log.Fatal(err)
	}
}

// handler keeps track of the count.
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "%d\n", rand.Intn(1000))
}

// counter emits the current value of count.
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "count %d\n", count)
	mu.Unlock()
}
