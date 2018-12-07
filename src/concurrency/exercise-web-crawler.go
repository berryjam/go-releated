package main

import (
	"fmt"
	"sync"
)

var cachedUrls map[string]bool = make(map[string]bool)
var mux sync.Mutex
var workNums int

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	quit := make(chan bool)
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	cachedUrls[url] = true
	for _, u := range urls {
		go paralleCrawl(u, depth-1, fetcher, quit)
	}
	<-quit
	return
}

func paralleCrawl(url string, depth int, fetcher Fetcher, quit chan bool) {
	mux.Lock()
	workNums++
	if cachedUrls[url] || depth <= 0 {
		workNums--
		if workNums == 0 {
			quit <- true
		}
		mux.Unlock()
		return
	} else {
		mux.Unlock()
		body, urls, err := fetcher.Fetch(url)
		mux.Lock()
		cachedUrls[url] = true
		mux.Unlock()
		if err != nil {
			fmt.Println(err)
			mux.Lock()
			workNums--
			if workNums == 0 {
				quit <- true
			}
			mux.Unlock()
			return
		}
		fmt.Printf("found: %s %q\n", url, body)
		for _, u := range urls {
			go paralleCrawl(u, depth-1, fetcher, quit)
		}
		mux.Lock()
		workNums--
		mux.Unlock()
	}
}

func main() {
	Crawl("http://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
