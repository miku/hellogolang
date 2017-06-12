// fetchall 
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
		"http://www.otz.de/",
		"http://www.sueddeutsche.de",
		"http://www.faz.net",
		"http://www.taz.de",
		"http://www.lvz.de",
		"http://www.l-iz.de",
		"http://www.zeit.de",
		"http://www.fr.de",
		"http://spiegel.de",
		"http://www.tagesspiegel.de",
		"http://news.google.com",
	}

	for _, url := range urls {
		go fetch(url, ch)
	}

	for range urls {
		fmt.Println(<-ch)
	}
	fmt.Printf("%0.4f\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading: %s: %v", url, err)
		return
	}
	ch <- fmt.Sprintf("%0.4f\t%7d\t%s", time.Since(start).Seconds(), nbytes, url)
}
