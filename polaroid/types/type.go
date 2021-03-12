package types

import (
	"polaroid/config"
	"polaroid/pk"

	"github.com/gorilla/sessions"
)

type Data struct {
	Store *sessions.CookieStore
	Conf  config.Config
	Db    *pk.Pk
}
