package tlsgo

import http "github.com/bogdanfinn/fhttp"

const OrderKey = http.HeaderOrderKey

type Header http.Header

func (h Header) Get(key string) string {
	return http.Header(h).Get(key)
}

func (h Header) Set(key, value string) {
	http.Header(h).Set(key, value)
}

func (h Header) Add(key, value string) {
	http.Header(h).Add(key, value)
}

func (h Header) Del(key string) {
	http.Header(h).Del(key)
}
