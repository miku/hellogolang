package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello world!")
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// func handler(w http.ResponseWriter, r *http.Request) {
// 	// fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
// 	fmt.Fprintf(w, time.Now().Format(time.Kitchen))
// }
