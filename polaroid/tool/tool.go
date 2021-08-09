package tool

import (
	"net/http"
	"path/filepath"
	"polaroid/types"
	"text/template"
)

type Good struct {
	Favicon string
	Title   string
	CSS     []string
	JS      []string
	Meta    []string
	Isco    bool
	Role    int
	Error	string
	Data    interface{}
}

type Header struct {
	R    *http.Request
	W    http.ResponseWriter
	Data *Good
}

func (h *Header) Error(a string) {
	h.Data.Error = a
}

func MyRole(s *types.Data, r *http.Request) (role int) {
	session, err := s.Store.Get(r, "user")
	if err != nil {
		role = 1
	} else {
		if session.Values["role"] == nil {
			role = 1
		} else {
			role = session.Values["role"].(int)
		}
	}
	return
}

func NewHeader(r *http.Request, w http.ResponseWriter, title string, d *types.Data) (Headers *Header) {
	Headers = &Header{
		R: r,
		W: w,
		Data: &Good{
			Title: "gopiko-"+title,
			CSS:   nil,
			JS:    nil,
			Isco:  false,
			Role:  MyRole(d, r),
			Meta:  nil,
		},
	}
	if title == "" {
		Headers.Data.Title = "gopiko"
	}
	Headers.Data.Favicon = ""
	return
}

func (h *Header) AddCss(a ...string) {
	for _, val := range a {
		h.Data.CSS = append(h.Data.CSS, "/public/css/"+val)
	}
}

func (h *Header) SetData(a interface{}) {
	h.Data.Data = a
}

func (h *Header) Jointure(ar ...string) {
	var joins []string
	for _, ok := range ar {
		joins = append(joins, filepath.Join("template", ok))
	}
	tmpl, err := template.ParseFiles(joins...)
	if err != nil {
		return
	}
	tmpl.ExecuteTemplate(h.W, "layout", h.Data)
	h.Data.Error = "";
}
