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
			if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
			}	
			
			fmt.Println(r.Form, r.FormValue("login"))
			t.Login = r.Form["login"][0]
			t.Email = r.Form["email"][0]
			t.Pwd = r.Form["pwd"][0]
			t.Pwd1 = r.Form["pwd1"][0]
			db, err := e.Db.Connect()
			if err != nil {
				h := tool.NewHeader(r,w,"signup",e)
				h.Jointure("layout.html","signup.html");
				return
			}
			defer db.Close()
			err = pk.InsertUsers(db, t.Email, t.Login, t.Pwd)
			if err != nil {
				fmt.Println(err)
				h := tool.NewHeader(r,w,"signup",e)
				h.Jointure("layout.html","signup.html");
				return
			}
			session, _ := e.Store.Get(r, "user")
			session.Values["login"] = t.Login
			session.Values["co"] = true
			session.Values["role"] = 2
			session.Save(r, w)
			http.Redirect(w, r, "http://gopiko.fr", 303)
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
			if err := r.ParseForm(); err != nil {
				fmt.Fprintf(w, "ParseForm() err: %v", err)
				return
			}
			t.Login = r.Form["pwd"][0]
			t.Pwd = r.Form["pwd"][0]
			db, err := e.Db.Connect();
			fmt.Println(err)
			if err != nil {
				h := tool.NewHeader(r, w, "signin",e)
				h.Error("non")
				h.Jointure("layout.html", "signing.html")
				return
			}
			defer db.Close()
			var tt bool
			tt, err = pk. UserExists(db, t.Login, t.Pwd)
			fmt.Println(err)
			if err != nil {
				h := tool.NewHeader(r, w, "signin",e)
				h.Error("y a un prob");
				h.Jointure("layout.html", "signing.html")
				return
			}
			if tt == false {
				h := tool.NewHeader(r, w, "signin",e)
				h.Error("zazd");
				h.Jointure("layout.html", "signing.html")
				return
			}
			session, _ := e.Store.Get(r, "user")
			session.Values["login"] = t.Login
			session.Values["co"] = true
			session.Values["role"] = 2
			session.Save(r, w)
			http.Redirect(w,r, "http://gopiko.fr",303)
			return
		}
	})
}


func (a *User) Logout(e *types.Data) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		println("deco")
		session, _ := e.Store.Get(r, "user")
		session.Options.MaxAge = -1
		session.Save(r, w)
		http.Redirect(w, r, "http://gopiko.fr/", 303);
	})
}

func (a *User) WWW(s *server.Server) {
	s.NewR("/signup", "signup", []string{"GET","POST"}, a.Signup(s.Data), 1)
	s.NewR("/signin", "signin", []string{"GET", "POST"}, a.Loging(s.Data), 1)
	s.NewR("/signout", "signout", []string{"GET"}, a.Logout(s.Data), 2)
}
