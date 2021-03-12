package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func writeError(err error) {
	if err != nil {
		println(err)
	}
}

func grap(r *http.Request, ou interface{}) (err error) {
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&ou)
	writeError(err)
	defer r.Body.Close()
	return
}

func sendJson(ar interface{}, w http.ResponseWriter) (err error) {
	w.Header().Set("Content-Type", "application/json;")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(ar)
	writeError(err)
	return
}

func newJson(method, url string, data interface{}) {
	client := &http.Client{}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(data)
	req, err := http.NewRequest(method, url, b)
	if err != nil {
		println(err)
		return
	}
	client.Do(req)
}
