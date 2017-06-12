package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func detect(link string, out chan string) {
	resp, err := http.Get(link)
	if err != nil {
		out <- fmt.Sprintf("could not fetch %s: %s", link, err)
		return
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		out <- fmt.Sprintf("error reading body %s: %s", link, err)
		return
	}
	if strings.Contains(string(b), "Donald Trump") {
		out <- fmt.Sprintf("%s 👱  x %v", link, strings.Count(string(b), "Donald Trump"))
	} else {
		out <- fmt.Sprintf("%s", link)
	}
}

func main() {
	links := []string{
		"http://www.otz.de/",
		"http://www.sueddeutsche.de",
		"http://www.faz.net",
		"http://www.taz.de",
		"http://www.lvz.de",
		"http://www.l-iz.de",
		"http://www.zeit.de",
	}
	ch := make(chan string)

	for _, link := range links {
		go detect(link, ch)
		// go detect(link, ch)
		// fmt.Println(<-ch)
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-ch)
	}
}
