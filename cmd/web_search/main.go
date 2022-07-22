package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/macedo/web_search/internal/web_search"
)

func main() {
	engine := flag.String("engine", "google", "select search engine (options: google/duckduckgo/stackoverflow/github)")
	flag.Parse()

	text := strings.Join(flag.Args(), " ")

	if err := run(*engine, text); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(engine string, text string) error {
	searchEngine, err := getSearchEngine(engine)
	if err != nil {
		return err
	}

	searchEngine.SetQuery(text)

	s := web_search.New(searchEngine)
	return s.Exec()
}

func getSearchEngine(engine string) (web_search.Engine, error) {
	switch engine {

	case "google":
		return web_search.NewGoogleEngine(), nil

	case "duckduckgo", "ddg":
		return web_search.NewDuckDuckGoEngine(), nil

	case "github":
		return web_search.NewGithubEngine(), nil

	case "stackoverflow":
		return web_search.NewStackOverflowEngine(), nil

	default:
		return nil, fmt.Errorf("search engine %q not supported", engine)
	}
}
