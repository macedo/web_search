package web_search

import "net/url"

type google struct {
	engine
}

func NewGoogleEngine() Engine {
	return &google{
		engine: engine{
			url: &url.URL{
				Scheme: "https",
				Host:   "google.com",
				Path:   "search",
			},
		},
	}
}
