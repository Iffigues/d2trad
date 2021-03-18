package user

import (
	"net/http"
	"polaroid/server"
	"polaroid/tool"
	"polaroid/types"
	"polaroid/pk"
	"fmt"
)

type User struct {
	Data *types.Data
}

type NewUsers struct {
	Email string `json:"email"`
	Login string `json:"login"`
	Pwd string `json:"pwd"`
	Pwd1 string `json:"pwd1"`

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
			var t NewUsers
			tool.Grap(r, &t);
			db, err := e.Db.Connect()
			if err != nil {
				h := tool.NewHeader(r,w,"home",e)
				h.Jointure("layout.html","signup.html");
				return
			}
			err = pk.InsertUsers(db, t.Email, t.Login, t.Pwd)
			if err != nil {
				fmt.Println(err)
				h := tool.NewHeader(r,w,"zz",e)
				h.Jointure("layout.html","signup.html");
				return
			}
			h := tool.NewHeader(r,w,"home",e);
			h.Jointure("layout.html","home.html")
			return
		}
	})
}

func (a *User) Loging(e *types.Data) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			h := tool.NewHeader(r,w,"signin",e)
			h.Jointure("layout.html", "signing.html");
			return
		}
		if r.Method == "POST" {
			var t NewUsers
			tool.Grap(r, &t);
			return
		}
	})
}

func (a *User) WWW(s *server.Server) {
	s.NewR("/signup", "signup", []string{"GET","POST"}, a.Signup(s.Data), 1)
	s.NewR("/signin", "signin", []string{"GET", "POST"}, a.Loging(s.Data), 1)
}
