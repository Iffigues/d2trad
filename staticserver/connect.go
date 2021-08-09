package main

import (
	"crypto/rand"
	"fmt"
	"net/http"
	"static/loger"
)

type Credentials struct {
	TokenString string `json:"token"`
	Password    string `json:"password"`
	Username    string `json:"username"`
}

func TokenGenerator(size int) string {
	b := make([]byte, 8)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func (aa *Data) Connecting(w http.ResponseWriter, r *http.Request) {
	if aa.Connect == nil || !aa.Connect.connected {
		var t Credentials
		grap(r, &t)
		uid := TokenGenerator(16)
		a, b, c, d := loger.Connecings(uid, t.Username, t.Password)
		fmt.Println(b, d)
		if d == nil {
			aa.Connect = NewConnect(a, b, c)
			w.WriteHeader(202)
			sendJson(aa.Connect.TokenString, w)
			return
		}
		w.WriteHeader(404)
		w.Write([]byte("error"))
		return;
	}
	w.WriteHeader(404)
	w.Write([]byte("not connected"))
}

func (aa *Data) IsCo(w http.ResponseWriter, r *http.Request) {
	if aa.Connect != nil {
		var t Credentials
		grap(r, &t)
		c, err := loger.GetJwt(t.TokenString)
		fmt.Println(c, err)
	}
}

func (aa *Data) Refresh(w http.ResponseWriter, r *http.Request) {
	if aa.Connect != nil && aa.Connect.connected == true {
		var t Credentials
		grap(r, &t)
		a, b, c, err := loger.Refresh(t.TokenString)
		if err != nil {
			w.WriteHeader(404)
			w.Write([]byte("faux"))
		}
		aa.Connect.Token = a
		aa.Connect.Times = c
		aa.Connect.TokenString = b
		sendJson(aa.Connect.TokenString, w)
	}
}
