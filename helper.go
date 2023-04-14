package tlsgo

import "net/url"

func (s *Session) addUrl(u url.URL) {
	// Check if URL is already in the list
	for _, url := range s.urls {
		if url == u {
			return
		}
	}
	s.urls = append(s.urls, u)
}
