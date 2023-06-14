package tlsgo

import (
	"net/url"

	tlsclient "github.com/bogdanfinn/tls-client"
)

type Session struct {
	client tlsclient.HttpClient
	urls   []url.URL

	Header Header
	Proxy  string
	Usage  int

	Injection func(Response, error) (Response, error)
}

func NewSession(p Profile, forceHttp1 bool) *Session {
	jar := tlsclient.NewCookieJar()
	options := []tlsclient.HttpClientOption{
		tlsclient.WithTimeoutSeconds(30),
		tlsclient.WithClientProfile(getTlsClientProfile(p)),
		tlsclient.WithRandomTLSExtensionOrder(),
		tlsclient.WithCookieJar(jar),
	}
	if forceHttp1 {
		options = append(options, tlsclient.WithForceHttp1())
	}
	client, _ := tlsclient.NewHttpClient(tlsclient.NewNoopLogger(), options...)
	session := &Session{client: client}
	session.Header = make(Header)
	session.Injection = func(r Response, e error) (Response, error) {
		return r, e
	}
	return session
}

func (s *Session) GetCookiesDict() map[string]string {
	cookies := make(map[string]string)
	for _, u := range s.urls {
		for _, c := range s.client.GetCookies(&u) {
			cookies[c.Name] = c.Value
		}
	}
	return cookies
}
