package link

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

//Link represents a link in an anchor tag
type Link struct {
	Href string
	text string
}

var r io.Reader

var res []Link

//Parse is used to parse the document and search for anchor tags
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		fmt.Println(err)
	}
	dfs(doc, "")
	fmt.Println(res)
	return nil, nil
}

func dfs(n *html.Node, padding string) {
	if n.Type == html.ElementNode && n.Data == "a" {
		//fmt.Println(n.FirstChild.Data)
		//fmt.Printf("%+v", n.Attr[0].Val)
		var val string
		for _, t := range n.Attr {
			if t.Key == "href" {
				val = t.Val
				break
			}
		}
		var text string
		text = getText(n)
		l := Link{
			Href: val,
			text: text,
		}
		res = append(res, l)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		dfs(c, padding+" ")
	}
}

func getText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}
	var res string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		res += getText(c) + " "
	}
	return strings.Join(strings.Fields(res), " ")
}
