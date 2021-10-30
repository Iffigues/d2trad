package main

import "fmt"
import "github.com/gocolly/colly/v2"
import "strings"

type Obj struct {
	Url		string
	Unique		bool
	Standar		bool
	Set		bool
	Specifique	bool
	Ring		bool
	Amu		bool
	Charme		bool
	Joyeau		bool
	Gemme		bool
	Class		string
}

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

func setObj(a string) (b bool) {
	t := []string{"anneaux", "amulettes", "charmes", "joyaux", "gemmes"}
	for _,c := range t {
		if a == c {
			return true
		}
	}
	return false
}

func child(a []string) (o []Obj) {
	var obj []Obj
	for _, b := range a {
		var oj Obj
		l := strings.Split(b[31:len(b) - 1],"-")
		oj.Url = b
		for _, val := range l {
			if val == "set" {
				oj.Set = true
			}
			if val == "specifiques" {
				oj.Specifique = true
			}
			if val == "uniques" {
				oj.Unique  = true
			}
			if val == "standards" {
				oj.Standar = true
			}
			if val == "anneaux" {
				oj.Ring = true
			}
			if val == "amulettes" {
				oj.Amu = true
			}
			if val == "charmes" {
				oj.Charme = true
			}
			if "joyaux" == val {
				oj.Joyeau = true
			}
			if "gemmes"  == val {
				oj.Gemme = true
			}
			if setClass(val) {
				oj.Class  = val
			}
		}
		obj = append(obj, oj)
	}
	return obj
}
