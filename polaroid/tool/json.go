package tool

import (
	"net/http"
	"encoding/json"
)

func Grap(r *http.Request,ou interface{})(err error){
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&ou)
	if err != nil {
		println(err)
	}
	defer r.Body.Close()
	return
}

func SendJson(ar interface{},w http.ResponseWriter)(err error){
	w.Header().Set("Content-Type", "application/json;")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(ar);
	if err != nil {
		println(err)
	}
	return
}
