package main

import (
	"net/http"
	"strings"
)

func (a *Data) HandleFunc(aa string, m []string, b func(w http.ResponseWriter, r *http.Request)) {
	method := []string{}
	var t []int
	for _, val := range m {
		if val == "GET" {
			method = append(method, http.MethodGet)
		}
	}
	e := strings.Split(aa, "/")
	for key, val := range e {
		if val == "*" {
			t = append(t, key)
		}
	}
	a.Url[aa] = H{
		H:      b,
		Method: m,
		Val:    t,
	}
}

func (a *Data) checkout(b string, w http.ResponseWriter, r *http.Request) (err error) {
	if val, ok := a.Url[b]; ok {
		method := r.Method
		for _, value := range val.Method {
			if method == value {
				val.H(w, r)
				return
			}
		}
	}
	return
}

func (a *Data) GetVal(b string) (aa []string) {
	e := strings.Split(b, "/")
	for _, val := range a.Url[b].Val {
		aa = append(aa, e[val])
	}
	return
}
