package tool

import (
	"net/http"
	"net/url"
)

type Form struct {
	FormData url.Values
}

func NewForm(r *http.Request) (a *Form, err error) {
    err = r.ParseForm()
    if err != nil {
        return nil, err
    }
    return &Form{
    	FormData: r.Form,
    }, nil
}

func (a *Form)HaveField(bb ...string) (b map[string]string ,ok bool) {
	ok = true
	for _, val := range bb {
		if _,ok := a.FormData[val]; !ok {
			return b, false
		}
	}
	return
}
