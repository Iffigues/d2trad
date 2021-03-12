package endpoint

import (
	"fmt"
	"net/http"
	"path/filepath"
	"polaroid/tool"
	"polaroid/types"
	"text/template"
)

func jointure(r *http.Request, w http.ResponseWriter, are interface{}, ar ...string) {
	var joins []string
	for _, ok := range ar {
		joins = append(joins, filepath.Join("template", ok))
	}
	tmpl, err := template.ParseFiles(joins...)
	fmt.Println(err)
	tmpl.ExecuteTemplate(w, "layout", are)
}

func Health(e *types.Data) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		head := tool.NewHeader(r, w, "gopiko", e)
		head.Jointure("layout.html", "home.html")
	})
}
