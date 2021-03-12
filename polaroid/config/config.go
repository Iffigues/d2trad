package config

import (
	_ "github.com/lib/pq"
)

type Conf map[string]interface{}

type Config map[string]Conf

func NewConf() (a Config) {
	return make(map[string]Conf)
}

func Pk() (ar Conf) {
	ar = make(map[string]interface{})
	ar["host"] = "psql"
	ar["port"] = 5432
	ar["user"] = "iffigues"
	ar["password"] = "Marie1426"
	ar["dbname"] = "polaroid"
	return
}
