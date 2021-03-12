package uploader

import (
	"io"
	"net/http"
)

type Up struct {
	Username string
	Pwd string
	Client *http.Client

}

func NewUp(a,b string) (r *Up) {
	return &Up{
		Username: a,
		Pwd: b,
		Client: &http.Client{},
	}
}

func (r *Up)MakeRequest(m, url string , i io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(m, url, i)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(r.Username, r.Pwd)
	return r.Client.Do(req);
}
