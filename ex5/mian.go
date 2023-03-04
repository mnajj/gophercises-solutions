package main

import (
	"encoding/xml"
	"flag"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"ex5/link"
)

type loc struct {
	Value string `xml:"loc"`
}

type urlSet struct {
	Urls []loc `xml:"url"`
}

func main() {
	urlFlag := flag.String("url", "https://gophercises.com/", "Specify url to build sitemap for")
	depth := flag.Int("depth", 3, "Specify sitemap max depth")
	flag.Parse()
	pages := bfs(*urlFlag, *depth)
	var toXml urlSet
	for _, p := range pages {
		toXml.Urls = append(toXml.Urls, loc{p})
	}
	enc := xml.NewEncoder(os.Stdout)
	if err := enc.Encode(toXml); err != nil {
		panic(err)
	}
}

func bfs(urlStr string, maxDepth int) []string {
	seen := make(map[string]struct{})
	var q map[string]struct{}
	nq := map[string]struct{}{
		urlStr: {},
	}
	for i := 0; i < maxDepth; i++ {
		q, nq = nq, make(map[string]struct{})
		for url, _ := range q {
			if _, ok := seen[url]; ok {
				continue
			}
			seen[url] = struct{}{}
			for _, l := range get(url) {
				nq[l] = struct{}{}
			}
		}
	}
	res := make([]string, 0, len(seen))
	for url, _ := range seen {
		res = append(res, url)
	}
	return res
}

func get(urlStr string) []string {
	res, err := http.Get(urlStr)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	reqUrl := res.Request.URL
	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}
	base := baseUrl.String()
	return filter(hrefs(res.Body, base), withPrefix(base))
}

func withPrefix(base string) func(link string) bool {
	return func(link string) bool {
		return strings.HasPrefix(link, base)
	}
}

func hrefs(r io.Reader, base string) []string {
	links, err := link.Parse(r)
	if err != nil {
		log.Fatal(err)
	}
	var hrefs []string
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			hrefs = append(hrefs, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			hrefs = append(hrefs, l.Href)
		}
	}
	return hrefs
}

func filter(links []string, keep func(string) bool) []string {
	var res []string
	for _, l := range links {
		if keep(l) {
			res = append(res, l)
		}
	}
	return res
}
