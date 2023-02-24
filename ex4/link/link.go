package link

import (
	"golang.org/x/net/html"
	"io"
	"strings"
)

type Link struct {
	Href string
	Text string
}

func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	nodes := linkNodes(doc)
	var links []Link
	for _, n := range nodes {
		links = append(links, buildLink(n))
	}
	return links, nil
}

func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, linkNodes(c)...)
	}
	return ret
}

func buildLink(in *html.Node) Link {
	var out Link
	for _, a := range in.Attr {
		if a.Key == "href" {
			out.Href = a.Val
			break
		}
	}
	out.Text = text(in)
	return out
}

func text(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}
	var res string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		res += text(c)
	}
	return strings.Join(strings.Fields(res), " ")
}
