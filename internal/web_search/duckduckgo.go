package web_search

import "net/url"

type duckduckgo struct {
	engine
}

func NewDuckDuckGoEngine() Engine {
	return &duckduckgo{
		engine: engine{
			url: &url.URL{
				Scheme: "https",
				Host:   "duckduckgo.com",
			},
		},
	}
}
