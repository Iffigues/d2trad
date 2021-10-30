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

func uniques(e Obj) {
	resp := uu(e.Url)
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var f func(*html.Node)
	var tt string
	fmt.Println(e.Url)
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
					print("normal\n")
				}
				if text.String() == "Exceptionnels" || text.String() == "Exceptionnelles" || text.String() == "EXCEPTIONNELleS"{
					tt = "Exceptionnels"
					print("exceptionnel\n")
				}
				if text.String() == "Élites" || text.String() == "éLITES" {
					tt = "Élites"
					print("elite\n")
				}
			}
		}
		if n.Type == html.ElementNode && n.Data == "div" {
			//var b bool
			for _, a := range n.Attr {
				if a.Key == "class" && a.Val == "et_pb_text_inner" {
					//b = true
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
}
