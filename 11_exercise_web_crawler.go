/*
In this exercise you'll use Go's concurrency features to parallelize a web crawler.
Modify the Crawl function to fetch URLs in parallel without fetching the same URL twice.
Hint: you can keep a cache of the URLs that have been fetched on a map, but maps alone are not safe for concurrent use! 
*/

package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type SafeFetchedUrls struct {
	v   map[string]bool
	mux sync.Mutex
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	fetchedUrls := SafeFetchedUrls{v: make(map[string]bool)}
	wg := &sync.WaitGroup{}
	c := make(chan string)
	wg.Add(1)
	go NightCrawler(url, depth, fetcher, fetchedUrls, c, wg)

	go func() {
		wg.Wait()
		close(c)
	}()
	for i := range c {
		fmt.Println(i)
	}
}
func (s SafeFetchedUrls) fetchedUrls(url string) bool {
	s.mux.Lock()
	defer s.mux.Unlock()
	_, ok := s.v[url]
	if ok == false {
		s.v[url] = true
		return false
	}
	return true

}
func NightCrawler(url string, depth int, fetcher Fetcher, fetchedUrls SafeFetchedUrls, c chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	if depth <= 0 || fetchedUrls.fetchedUrls(url) {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	c <- url + " " + body
	if err != nil {
		return
	}

	for _, u := range urls {
		go NightCrawler(u, depth-1, fetcher, fetchedUrls, c, wg)
		wg.Add(1)
	}
	return
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
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
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

/*
https://golang.org/ The Go Programming Language
https://golang.org/cmd/ 
https://golang.org/pkg/ Packages
https://golang.org/pkg/os/ Package os
https://golang.org/pkg/fmt/ Package fmt
*/
