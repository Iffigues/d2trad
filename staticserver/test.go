package main

import (
	"net/http"
)

func (a *Data)hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	w.WriteHeader(200)
	w.Write(a.Home)
}
