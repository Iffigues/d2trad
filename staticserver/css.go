package main

import (
	"errors"
	"net/http"
	"path/filepath"
)

func (a *Data) AddData(url string) (rr Give) {
	code := 200
	val, err := Asset(url[1:])
	ext := filepath.Ext(url[1:])
	types := "application/octet-stream"
	if len(ext) > 0 {
		ext = ext[1:]
	}
	if vals, ok := a.Types[ext]; ok {
		types = vals
	}
	if val == nil && err == nil {
		err = errors.New("empty data")
	}
	if err != nil {
		code = 404
		types = "text/html; charset=utf-8"
		val = a.Error
	}
	rr = Give{
		Code:  code,
		Types: types,
		Bytes: val,
	}
	a.Data[url] = rr
	return
}

func (a *Data) File(w http.ResponseWriter, r *http.Request) {
	if val, ok := a.Data[r.URL.Path]; ok {
		w.Header().Set("Content-type", val.Types)
		w.WriteHeader(val.Code)
		w.Write(val.Bytes)
		return
	}
	val := a.AddData(r.URL.Path)
	w.Header().Set("Content-type", val.Types)
	w.WriteHeader(val.Code)
	w.Write(val.Bytes)
	return
}
