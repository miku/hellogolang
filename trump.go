package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func detect(link string, out chan string) {
	// start := time.Now()
	// defer func() {
	// 	fmt.Printf("%s took %s\n", link, time.Since(start))
	// }()
	resp, err := http.Get(link)
	if err != nil {
		out <- fmt.Sprintf("could not fetch %s: %s", link, err)
		return
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		out <- fmt.Sprintf("error reading body %s: %s", link, err)
		return
	}
	detected := strings.Contains(string(b), "Trump")
	if detected {
		out <- fmt.Sprintf("%s 👱", link)
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
		// go detect(link, ch)
		go detect(link, ch)
		fmt.Println(<-ch)
	}

	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-ch)
	// }
}
