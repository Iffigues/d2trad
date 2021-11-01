package main

import (
	"golang.org/x/net/html"
	"net/http"
	"log"
	"fmt"
	"bytes"
)

func uu(e string) (r *http.Response){
	var err error
	r, err = http.Get(e)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func collectText(n *html.Node, buf *bytes.Buffer) {
    if n.Type == html.TextNode {
        buf.WriteString(n.Data)
    }
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        collectText(c, buf)
    }
}

func collectImg(n *html.Node, buf *bytes.Buffer) {
    if n.Type == html.ElementNode && n.Data == "img" {
	for _,a := range n.Attr {
		if a.Key == "src" {
			buf.WriteString(a.Val)
		}
	}
    }
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        collectImg(c, buf)
    }
}

func h(e string, cc *html.Node) {
	i := 0
	var name string
	var char string
	var img string
	var f func(*html.Node)
	f = func(n *html.Node ) {
		if n.Type == html.ElementNode && n.Data == "td" {
			if i == 1 {
				text := &bytes.Buffer{}
				collectImg(n, text)
				img = text.String()
			} else  {
				text := &bytes.Buffer{}
				collectText(n, text)
				if i == 0 {
					name = text.String()
				}
				if i == 2 {
					char = text.String()
				}
			}
			i = i + 1
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(cc)
	fmt.Println(name, char, img)
}

func uniques(e Obj) {
	resp := uu(e.Url)
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var f func(*html.Node)
	var tt string
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "div" {
			var  bb bool
			for _,a := range n.Attr {
				if a.Key == "class" && a.Val == "et_pb_text_inner" {
					bb = true

				}
			}
			if bb {
				text := &bytes.Buffer{}
				collectText(n, text)
				if text.String() == "Normaux" || text.String() == "Normales" || text.String() == "normales" {
					tt = "Normaux"
					//print("normal\n")
				}
				if text.String() == "Exceptionnels" || text.String() == "Exceptionnelles" || text.String() == "EXCEPTIONNELleS"{
					tt = "Exceptionnels"
					//print("exceptionnel\n")
				}
				if text.String() == "Élites" || text.String() == "éLITES" {
					tt = "Élites"
					//print("elite\n")
				}
			}
		}
		if n.Type == html.ElementNode && n.Data == "table" {
			h(tt, n)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
}
