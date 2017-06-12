package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	urls := []string{
		"http://www.otz.de",
		"http://www.lvz.de",
		"http://faz.net",
		"http://zeit.de",
		"https://uni-jena.de",
		"http://bahn.de",
	}

	for _, url := range urls {
		go fetch(url, ch)
	}

	for range urls {
		fmt.Println(<-ch)
	}
	fmt.Printf("%s elapsed\n", time.Since(start))
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("%s", err)
		return
	}
	n, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, n, url)
}
