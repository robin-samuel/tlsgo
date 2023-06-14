package tlsgo

import (
	tlsclient "github.com/bogdanfinn/tls-client"
)

type Profile string

const (
	// Chrome
	Chrome109 Profile = "Chrome 109"
	Chrome110 Profile = "Chrome 110"
	Chrome111 Profile = "Chrome 111"
	Chrome112 Profile = "Chrome 112"

	// Firefox
	Firefox105 Profile = "Firefox 105"
	Firefox106 Profile = "Firefox 106"
	Firefox108 Profile = "Firefox 108"
	Firefox110 Profile = "Firefox 110"

	// Safari
	Safari160    Profile = "Safari 160"
	SafariIOS160 Profile = "Safari IOS 160"

	// Android
	Okhttp4Android10 Profile = "Okhttp4 Android 10"
)

func getTlsClientProfile(profile Profile) tlsclient.ClientProfile {
	switch profile {
	// Chrome
	case Chrome109:
		return tlsclient.Chrome_109
	case Chrome110:
		return tlsclient.Chrome_110
	case Chrome111:
		return tlsclient.Chrome_111
	case Chrome112:
		return tlsclient.Chrome_112

	// Firefox
	case Firefox105:
		return tlsclient.Firefox_105
	case Firefox106:
		return tlsclient.Firefox_106
	case Firefox108:
		return tlsclient.Firefox_108
	case Firefox110:
		return tlsclient.Firefox_110

	// Safari
	case Safari160:
		return tlsclient.Safari_16_0
	case SafariIOS160:
		return tlsclient.Safari_IOS_16_0

	// Android
	case Okhttp4Android10:
		return tlsclient.Okhttp4Android10

	default:
		return tlsclient.Chrome_112
	}
}
