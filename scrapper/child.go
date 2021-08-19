package main

import "fmt"
import "github.com/gocolly/colly/v2"
import "strings"

func set(a string) {
	t := &Collector{
		url: a,
		c: colly.NewCollector(),
	}
	t.c.OnHTML("table", func(e *colly.HTMLElement) {
		fmt.Println(e.Attr("style"))
	})
	t.c.Visit(t.url)
}

func unique(a string) {

}

func standart(a string) {
}


func haveStr(a []string, c string) (b bool) {
	for _, g := range a {
		if g == c {
			return true
		}
	}
	return false
}

func setClass(a string) (b bool) {
	t := []string{"amazone", "assassin", "barbare", "necromancien", "druide", "sorciere", "paladin"}
	for _,c := range t {
		if a == c {
			return true
		}
	}
	return false
}

func child(a []string) {
	for _, b := range a {
		l := strings.Split(b[31:len(b) - 1],"-")
		e := len(l)
		if l[0] == "set" {
			if setClass(l[1]) {
			} else {
			}
		}
		if l[2] == "specifiques" {
		}
		if l[1] == "uniques" {
			fmt.Println(l[1], l)
		}
		z := l[e - 1]
		if z == "standards" {
			fmt.Println("stand", l)
		}
		if z == "uniques" {
			fmt.Println("unique", l)
		}
	}
}
