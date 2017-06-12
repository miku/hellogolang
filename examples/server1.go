// server1 serves a constant string.
package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	// Functions are first class objects.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello world!")
	})
	if err:= http.ListenAndServe("localhost:8000", nil); err != nil {
		log.Fatal(err)
	}
}

// func handler(w http.ResponseWriter, r *http.Request) {
// 	// fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
// 	fmt.Fprintf(w, time.Now().Format(time.Kitchen))
// }
