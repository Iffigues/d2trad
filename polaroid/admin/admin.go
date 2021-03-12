package admin

import (
	"net/http"
	"polaroid/server"
	"polaroid/tool"
	"polaroid/types"
)

type Admin struct {
	Data *types.Data
}

func NewAdmin(s *types.Data) (a *Admin) {
	a = new(Admin)
	a.Data = s
	return
}

func (a *Admin) IsAdmin(r *http.Request) (ok bool) {
	session, err := a.Data.Store.Get(r, "user")
	if err != nil {
		return false
	}
	if key, ok := session.Values["role"]; ok {
		if key == 3 {
			return true
		}
	}
	return false
}

func (a *Admin) isAdmin(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if a.IsAdmin(r) {
			h.ServeHTTP(w, r)
		} else {
			head := tool.NewHeader(r, w, "gopiko-admin", a.Data)
			head.Jointure("layout.html", "admin.html")
			return
		}
	})
}

func (a *Admin) Admins(e *types.Data) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		head := tool.NewHeader(r, w, "gopiko-admin", e)
		head.Jointure("layout.html", "admin.html")
	})
}

func (a *Admin) ConnectAdmins(e *types.Data) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if a.IsAdmin(r) {
			return
		}
		println("oui")
		if r.Method == "GET" {
			head := tool.NewHeader(r, w, "gopiko-admin", e)
			head.Jointure("layout.html", "admincon.html")
		}
		if r.Method == "POST" {
		   _, err := tool.NewForm(r)
		   if err != nil {
		   	return
		   }
		}
	})
}

func (a *Admin) WWW(s *server.Server) {
	s.NewR("/admins", "admins", []string{"GET"}, a.isAdmin(a.Admins(s.Data)), 1)
	s.NewR("/admin", "admin", []string{"GET", "POST"}, a.ConnectAdmins(s.Data), 1)
}
