package main

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"time"
)

type Give struct {
	Types string
	Code  int
	Bytes []byte
}

type Connect struct {
	connected   bool
	Token       *jwt.Token
	Times       time.Time
	TokenString string
}

func NewConnect(token *jwt.Token, tokenString string, expirationTime time.Time) (a *Connect) {
	a = &Connect{
		connected:   true,
		Token:       token,
		Times:       expirationTime,
		TokenString: tokenString,
	}
	return
}

type H struct {
	H      func(w http.ResponseWriter, r *http.Request)
	Method []string
	Val    []int
}

type Oauth struct {
}

type Data struct {
	Url     map[string]H
	Data    map[string]Give
	Types   map[string]string
	Connect *Connect
	Oauth   *Oauth
	Error   []byte
	Home []byte
}

func existe(e interface{}) {
	if e != nil {
	}
}

func NewData() (D *Data) {
	html, err := Asset("public/html/404.html")
	if err != nil {
		log.Fatal(err)
	}
	home, err := Asset("public/html/home.html")
	if err != nil {
		log.Fatal(err)
	}
	return &Data{
		Data: make(map[string]Give),
		Url:  make(map[string]H),
		Types: map[string]string{
			"css":  "text/css",
			"json": "application/json",
			"png":  "image/png",
		},
		Error: html,
		Home: home,
	}
}
