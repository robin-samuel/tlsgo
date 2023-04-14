package tlsgo

import (
	"bytes"
	"encoding/json"
	"io"
	"net/url"

	http "github.com/bogdanfinn/fhttp"
	"github.com/bogdanfinn/fhttp/httputil"
)

type Request struct {
	Method  Method
	Url     string
	URL     url.URL
	Params  url.Values
	Header  Header
	Cookies Cookies
	Data    io.Reader
	Json    interface{}
	Body    []byte
	Text    string
	Proxy   string
	Size    int
	http    *http.Request
}

type Options struct {
	Params  url.Values
	Header  Header
	Cookies Cookies
	Data    io.Reader
	Json    interface{}
	Proxy   string
}

type Method string

const (
	GET     Method = "GET"
	POST    Method = "POST"
	PUT     Method = "PUT"
	PATCH   Method = "PATCH"
	DELETE  Method = "DELETE"
	HEAD    Method = "HEAD"
	OPTIONS Method = "OPTIONS"
)

func NewRequest(method Method, url string, options *Options) Request {
	request := Request{Method: method, Url: url}
	if options != nil {
		request.Params = options.Params
		request.Header = options.Header
		request.Cookies = options.Cookies
		request.Data = options.Data
		request.Json = options.Json
		request.Proxy = options.Proxy
	}
	return request
}

func (r *Request) Raw() []byte {
	dump, _ := httputil.DumpRequest(r.http, true)
	raw := append(dump, r.Body...)
	return raw
}

func (s *Session) Do(request Request) (Response, error) {

	var response Response

	// Create URL
	u, err := url.Parse(request.Url)
	if err != nil {
		return response, err
	}

	// Set Params
	if len(request.Params) > 0 {
		q := u.Query()
		for k, v := range request.Params {
			for _, vv := range v {
				q.Add(k, vv)
			}
		}
		u.RawQuery = q.Encode()
	}

	// Set Proxy
	if request.Proxy != "" {
		s.Proxy = request.Proxy
		s.client.SetProxy(request.Proxy)
	}

	// Create Request
	req, err := http.NewRequest(string(request.Method), request.Url, nil)
	if err != nil {
		return response, err
	}

	// Set Headers
	req.Header = http.Header(s.Header)
	for k, v := range request.Header {
		for _, vv := range v {
			req.Header.Add(k, vv)
		}
	}

	// Set Cookies
	if len(request.Cookies) > 0 {
		s.client.SetCookies(u, request.Cookies)
	}

	// Set Body
	if request.Data != nil {
		req.Body = io.NopCloser(request.Data)
	}

	// Set Json
	if request.Json != nil {
		jsonBody, err := json.Marshal(request.Json)
		if err != nil {
			return response, err
		}
		req.Body = io.NopCloser(bytes.NewReader(jsonBody))
	}

	// Send Request
	resp, err := s.client.Do(req)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()

	// Dump Request & Response
	var dumpReq, dumpResp []byte
	if dumpReq, err = httputil.DumpRequestOut(req, true); err != nil {
		return response, err
	}
	if dumpResp, err = httputil.DumpResponse(resp, true); err != nil {
		return response, err
	}

	// Read Body Request
	bodyReq, err := io.ReadAll(req.Body)
	if err != nil {
		return response, err
	}

	// Read Body Response
	bodyRes, err := io.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}

	// Update Request
	request.Url = req.URL.String()
	request.Size = len(dumpReq)
	request.Header = Header(req.Header)
	request.Cookies = req.Cookies()
	request.Body = bodyReq
	request.Text = string(bodyReq)
	request.http = req

	// Update Response
	response.Request = request
	response.StatusCode = resp.StatusCode
	response.Status = resp.Status
	response.Url = resp.Request.URL.String()
	response.URL = *resp.Request.URL
	response.Params = resp.Request.URL.Query()
	response.Header = Header(resp.Header)
	response.Cookies = resp.Cookies()
	response.Body = bodyRes
	response.Text = string(bodyRes)
	response.Size = len(dumpResp)
	response.http = resp

	// Update Session
	s.addUrl(response.URL)
	s.Usage += request.Size + response.Size

	return s.Injection(response, err)
}

func (s *Session) Get(url string, options *Options) (Response, error) {
	return s.Do(NewRequest(GET, url, options))
}

func (s *Session) Post(url string, options *Options) (Response, error) {
	return s.Do(NewRequest(POST, url, options))
}

func (s *Session) Put(url string, options *Options) (Response, error) {
	return s.Do(NewRequest(PUT, url, options))
}

func (s *Session) Patch(url string, options *Options) (Response, error) {
	return s.Do(NewRequest(PATCH, url, options))
}

func (s *Session) Delete(url string, options *Options) (Response, error) {
	return s.Do(NewRequest(DELETE, url, options))
}

func (s *Session) Head(url string, options *Options) (Response, error) {
	return s.Do(NewRequest(HEAD, url, options))
}

func (s *Session) Options(url string, options *Options) (Response, error) {
	return s.Do(NewRequest(OPTIONS, url, options))
}
