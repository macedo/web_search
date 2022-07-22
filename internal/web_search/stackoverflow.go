package web_search

import "net/url"

type stackoverflow struct {
	engine
}

func NewStackOverflowEngine() Engine {
	return &stackoverflow{
		engine: engine{
			url: &url.URL{
				Scheme: "https",
				Host:   "stackoverflow.com",
				Path:   "search",
			},
		},
	}
}
