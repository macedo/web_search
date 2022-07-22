package web_search

import (
	"net/url"
)

type Engine interface {
	SetQuery(string)
	URL() string
}

type engine struct {
	url *url.URL
}

func (e *engine) SetQuery(query string) {
	q := e.url.Query()
	q.Set("q", query)

	e.url.RawQuery = q.Encode()
}

func (e *engine) URL() string {
	return e.url.String()
}
