package main

import (
	"net/http"
	"strings"
)

func (a *Data) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	url := strings.Split(r.URL.Path, "/")
	if len(url) > 1 && url[1] == "public" {
		if r.Method == "GET" {
			a.File(w, r)
			return
		}
	}
	if len(url) > 1 && url[1] == "import" {
		if r.Method == "GET" {
			a.File(w, r)
			return
		}
	}
	for key, val := range a.Url {
		urls := strings.Split(key, "/")
		if len(url) == len(urls) {
			yes := true
			for i := 0; i < len(url); i = i + 1 {
				if urls[i] != "*" {
					if urls[i] != url[i] {
						yes = false
						break
					}
				}
			}
			if yes {
				for _, valeur := range val.Method {
					if r.Method == valeur {
						val.H(w, r)
						return
					}
				}
				w.WriteHeader(404)
				w.Write([]byte("not found"))
				return
			}
		}
	}
	a.hello(w, r)
}

func main() {
	a := NewData()
	a.HandleFunc("/", []string{"GET"}, a.hello)
	a.HandleFunc("/upload", []string{"POST"}, a.Upload)
	a.HandleFunc("/public$", []string{"GET"}, a.File)
	a.HandleFunc("/loging", []string{"POST"}, a.Connecting)
	a.HandleFunc("/refresh", []string{"POST"}, a.Refresh)
	a.HandleFunc("/isco", []string{"POST"}, a.IsCo)
	http.Handle("/", a)
	http.ListenAndServe(":3006", nil)
}
