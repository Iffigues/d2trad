package user

import (
	"net/http"
	"polaroid/server"
	"polaroid/tool"
	"polaroid/types"
)

type User struct {
	Data *types.Data
}

func NewUser(s *types.Data) (a *User) {
	a = new(User)
	a.Data = s
	return
}

func (a *User) Signup(e *types.Data) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			h := tool.NewHeader(r,w,"signup",e)
			h.Jointure("layout.html", "signup.html")
			return
		}
		if r.Method == "POST" {
			return
		}
	})
}

func (a *User) Loging(e *types.Data) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})
}

func (a *User) WWW(s *server.Server) {
	s.NewR("/signup", "admins", []string{"GET"}, a.Signup(s.Data), 1)
	s.NewR("/signin", "admin", []string{"GET", "POST"}, a.Loging(s.Data), 1)
}
