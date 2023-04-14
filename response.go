package tlsgo

import (
	"encoding/json"
	"net/url"

	http "github.com/bogdanfinn/fhttp"
	"github.com/bogdanfinn/fhttp/httputil"
)

type Response struct {
	StatusCode int
	Status     string
	Url        string
	URL        url.URL
	Params     url.Values
	Header     Header
	Cookies    Cookies
	Body       []byte
	Text       string
	Size       int
	Request    Request
	http       *http.Response
}

func (r *Response) Json(v interface{}) error {
	return json.Unmarshal(r.Body, v)
}

func (r *Response) Raw() []byte {
	dump, _ := httputil.DumpResponse(r.http, true)
	raw := append(dump, r.Body...)
	return raw
}
