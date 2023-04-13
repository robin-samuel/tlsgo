package tlsgo

import (
	tlsclient "github.com/bogdanfinn/tls-client"
)

type Profile string

const (
	Chrome109 Profile = "Chrome 109"
	Chrome110 Profile = "Chrome 110"
	Chrome111 Profile = "Chrome 111"
	Chrome112 Profile = "Chrome 112"
)

func getTlsClientProfile(profile Profile) tlsclient.ClientProfile {
	switch profile {
	case Chrome109:
		return tlsclient.Chrome_109
	case Chrome110:
		return tlsclient.Chrome_110
	case Chrome111:
		return tlsclient.Chrome_111
	case Chrome112:
		return tlsclient.Chrome_112
	default:
		return tlsclient.Chrome_112
	}
}
