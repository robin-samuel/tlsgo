package tlsgo

import (
	tlsclient "github.com/bogdanfinn/tls-client"
)

type Session struct {
	client tlsclient.HttpClient
	Header Header
	Proxy  string
	Usage  int

	Injection func(Response, error) (Response, error)
}

func NewSession(p Profile) *Session {
	jar := tlsclient.NewCookieJar()
	options := []tlsclient.HttpClientOption{
		tlsclient.WithTimeoutSeconds(30),
		tlsclient.WithClientProfile(getTlsClientProfile(p)),
		tlsclient.WithRandomTLSExtensionOrder(),
		tlsclient.WithCookieJar(jar),
	}
	client, _ := tlsclient.NewHttpClient(tlsclient.NewNoopLogger(), options...)
	session := &Session{client: client}
	session.Injection = func(r Response, e error) (Response, error) {
		return r, e
	}
	return session
}
