package main

import (
	"github.com/gocolly/colly/v2"
)

type Collector struct {
	url string
	child []string
	c  *colly.Collector
}
