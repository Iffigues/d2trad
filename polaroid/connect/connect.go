package connect

import (
	"net/http"
	"polaroid/server"
	"polaroid/tool"
	"polaroid/types"
)

type Datas struct {
	User string `json:"user"`
	Pwd string `json:"pwd"`
}

type Connect struct {
	Data *types.Data
}

func NewConnect(s *types.Data) (a *Connect) {
	a = new(Connect)
	a.Data = s
	return
}

func (a *Connect) Connects(e *types.Data) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var t *Datas
		tool.Grap(r, t)
	})
}

func (a *Connect) WWW(s *server.Server) {
	s.NewR("/connect", "connect", []string{"POST"}, a.Connects(s.Data), 0)
}
