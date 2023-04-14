package tlsgo

import "net/url"

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
}
