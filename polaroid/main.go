package main

import (
	"fmt"
	"polaroid/admin"
	"polaroid/connect"
	"polaroid/config"
	"polaroid/server"
)

func main() {
	conf := config.NewConf()
	conf["pk"] = config.Pk()
	srv := server.NewServer(conf)
	adm := admin.NewAdmin(srv.Data)
	srv.AddHH(adm)
	connect := connect.NewConnect(srv.Data);
	srv.AddHH(connect);
	serve := srv.Servers(conf)
	err := serve.ListenAndServe()
	fmt.Println(err)
}
