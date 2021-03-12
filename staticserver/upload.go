package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"mime/multipart"
	"static/loger"
)



func FileName(h *multipart.FileHeader) (a string , err error) {
	return 
}

func (a *Data) Upload(w http.ResponseWriter, r *http.Request) {
	var t Credentials
	grap(r, &t)
	_, err := loger.GetJwt(t.TokenString)
	if err != nil {
		return
	}
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	path, err := FileName(handler)
	if err != nil {
		return
	}
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
}
