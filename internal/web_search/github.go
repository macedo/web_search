package web_search

import "net/url"

type github struct {
	engine
}

func NewGithubEngine() Engine {
	return &github{
		engine: engine{
			url: &url.URL{
				Scheme: "https",
				Host:   "github.com",
				Path:   "search",
			},
		},
	}
}
