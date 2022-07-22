package web_search

type Search struct {
	engine Engine
}

func New(engine Engine) *Search {
	return &Search{
		engine: engine,
	}
}
